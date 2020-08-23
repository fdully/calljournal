package calljournal

import (
	"bytes"
	"context"
	"io"

	cjdb "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBaseCallServer(env *serverenv.ServerEnv, config *Config) *BaseCallServer {
	return &BaseCallServer{
		db:        cjdb.NewCallJournalDB(env.Database()),
		blobStore: env.Blobstore(),
		config:    config,
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

	// TODO send basecall to NSQ queue

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

	res := &pb.SaveAudioRecordResponse{Uuid: info.Uuid, Size: uint32(recordSize)}

	// TODO send recordInfo to NSQ queue

	pth := util.CreateRecordPath(recordInfo)

	// storing audio record in wav to temporary file storage
	err = c.blobStore.CreateObject(context.Background(), c.config.Bucket, pth, recordData)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to store record: %v", err))
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response saving record: %v", err))
	}

	return nil
}
