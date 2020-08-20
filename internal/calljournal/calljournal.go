package calljournal

import (
	"bytes"
	"context"
	"io"

	cjdb "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/lame"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBaseCallServer(ctx context.Context, env *serverenv.ServerEnv, config *Config) *BaseCallServer {
	return &BaseCallServer{
		db:        cjdb.NewCallJournalDB(env.Database()),
		blobStore: env.Blobstore(),
		config:    config,
		ctx:       ctx,
	}
}

var (
	_ pb.BaseCallServiceServer    = (*BaseCallServer)(nil)
	_ pb.AudioRecordServiceServer = (*BaseCallServer)(nil)
)

type BaseCallServer struct {
	db        *cjdb.CallJournalDB
	blobStore storage.Blobstore
	config    *Config
	ctx       context.Context
}

func (c *BaseCallServer) SaveBaseCall(ctx context.Context, req *pb.SaveBaseCallRequest) (*pb.SaveBaseCallResponse, error) {
	pbBC := req.GetBaseCall()

	_, err := uuid.Parse(pbBC.Uuid)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.InvalidArgument, "basecall %s is not a valid UUID: %v", pbBC.Uuid, err))
	}

	bc, err := util.ProtobufBaseCallToBaseCall(pbBC)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	err = c.db.AddBaseCall(ctx, bc)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	return &pb.SaveBaseCallResponse{Uuid: bc.UUID.String()}, nil
}

func (c *BaseCallServer) SaveAudioRecord(stream pb.AudioRecordService_SaveAudioRecordServer) error {
	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to receive record info %v", err))
	}

	info := req.GetInfo()
	if info == nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to get record info, it is nil"))
	}

	recordInfo, err := util.ProtobufRecordInfoToRecordInfo(info)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "%v", err))
	}

	recordData := new(bytes.Buffer)
	recordSize := 0

	for {
		err := util.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "failed to receive record data for %s: %v", info.Uuid, err))
		}

		chunk := req.GetRecordChunk()
		recordSize += len(chunk)

		_, err = recordData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to write chunk record: %v", err))
		}
	}

	res := &pb.SaveAudioRecordResponse{
		Uuid: info.Uuid,
		Size: uint32(recordSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response saving record: %v", err))
	}

	go c.processAudioRecord(c.ctx, recordInfo, recordData)

	return nil
}

func (c *BaseCallServer) processAudioRecord(ctx context.Context, recordInfo model.RecordInfo,
	recordData *bytes.Buffer) {
	c.createMp3AndSavetoStorage(ctx, recordInfo, recordData)
}

func (c *BaseCallServer) createMp3AndSavetoStorage(ctx context.Context, recordInfo model.RecordInfo,
	recordData *bytes.Buffer) {
	logger := logging.FromContext(ctx)

	mp3Data, err := lame.Wav2Mp3(ctx, recordData)
	if err != nil {
		logger.Fatal(err)
	}

	recordInfo.Name = util.ChangeWavExtToMp3(recordInfo.Name)
	pth := util.CreateRecordPath(recordInfo)

	logger.Info("saving record to storage")

	err = c.blobStore.CreateObject(ctx, c.config.Bucket, pth, mp3Data)
	if err != nil {
		logger.DPanic(err)
	}
}
