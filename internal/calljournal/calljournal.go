package calljournal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	cjdb "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/queue"
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
		publisher: env.Publisher(),
		config:    config,
	}
}

var (
	_ pb.BaseCallServiceServer   = (*BaseCallServer)(nil)
	_ pb.RecordInfoServiceServer = (*BaseCallServer)(nil)
	_ pb.RecordDataServiceServer = (*BaseCallServer)(nil)
)

type BaseCallServer struct {
	db        *cjdb.CallJournalDB
	blobStore storage.Blobstore
	publisher queue.Publisher
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

	if err := c.db.AddBaseCall(ctx, bc); err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	msg, err := json.Marshal(bc)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "failed to make json basecall: %v", err))
	}

	if err := c.publisher.Publish(c.config.BaseCallTopic, msg); err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	return &pb.SaveBaseCallResponse{Uuid: bc.UUID.String()}, nil
}

func (c *BaseCallServer) SaveRecordData(stream pb.RecordDataService_SaveRecordDataServer) error {
	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to receive record info %v", err))
	}

	info := req.GetInfo()
	if info == nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to get record info, it is nil"))
	}

	ri, err := util.ProtobufRecordInfoToRecordInfo(info)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "%v", err))
	}

	recordData := new(bytes.Buffer)
	recordSize := 0

	for {
		if err := util.ContextError(stream.Context()); err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "failed to receive record data for %s: %v", info.Uuid, err))
		}

		chunk := req.GetRecordChunk()
		recordSize += len(chunk)

		_, err = recordData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to write chunk record: %v", err))
		}
	}

	res := &pb.SaveRecordDataResponse{Uuid: info.Uuid, Size: uint32(recordSize)}
	pth := util.CreateFileRecordPath(ri)

	// storing audio record in wav to temporary file storage
	err = c.blobStore.CreateObject(context.Background(), c.config.Bucket, pth, recordData)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to store record: %v", err))
	}

	if err := c.publishRecordInfo(ri); err != nil {
		return util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response saving record: %v", err))
	}

	return nil
}

func (c *BaseCallServer) SaveRecordInfo(ctx context.Context, req *pb.SaveRecordInfoRequest) (*pb.SaveRecordInfoResponse, error) {
	ri := req.GetInfo()

	_, err := uuid.Parse(ri.Uuid)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.InvalidArgument,
			"record info %s is not a valid UUID: %v", ri.Uuid, err))
	}

	r, err := util.ProtobufRecordInfoToRecordInfo(ri)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	if err := c.db.AddRecordInfo(ctx, r); err != nil {
		return nil, err
	}

	return &pb.SaveRecordInfoResponse{Uuid: ri.Uuid}, nil
}

func (c *BaseCallServer) publishRecordInfo(ri model.RecordInfo) error {
	msg, err := json.Marshal(ri)
	if err != nil {
		return fmt.Errorf("failed to make json record info: %w", err)
	}

	if err := c.publisher.Publish(c.config.RecordInfoTopic, msg); err != nil {
		return fmt.Errorf("failed to publish record info %s: %w", ri.UUID.String(), err)
	}

	return nil
}
