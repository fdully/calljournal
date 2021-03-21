package basecall

import (
	"bytes"
	"context"
	"io"

	"github.com/fdully/calljournal/internal/calljournal"
	cdrstoredb "github.com/fdully/calljournal/internal/cdrstore/database"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	GRPCMsgByteChunk = 64 * 1024
)

func NewServer(env *serverenv.ServerEnv, config *calljournal.Config) *Server {
	return &Server{
		blobStore: env.Blobstore(),
		publisher: env.Publisher(),
		db:        cdrstoredb.New(env.Database()),
		config:    config,
	}
}

var _ pb.BaseCallServiceServer = (*Server)(nil)

type Server struct {
	blobStore storage.Blobstore
	publisher queue.Publisher
	db        *cdrstoredb.CDRStoreDB
	config    *calljournal.Config
}

func (c *Server) SaveBaseCall(ctx context.Context, pbBC *pb.SaveBaseCallRequest) (*pb.SaveBaseCallResponse, error) {
	pbBaseCall := pbBC.GetBasecall()

	_, err := uuid.Parse(pbBaseCall.Uuid)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.InvalidArgument, "basecall %s is not a valid UUID: %v", pbBaseCall.Uuid, err))
	}

	bc, err := ProtobufBaseCallToBaseCall(pbBaseCall)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	err = c.db.AddBaseCall(ctx, *bc)
	if err != nil {
		return nil, util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	return &pb.SaveBaseCallResponse{Uuid: bc.UUID.String()}, nil
}

func (c *Server) SaveBaseCallRecord(stream pb.BaseCallService_SaveBaseCallRecordServer) error {
	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to receive basecall record callpath %v", err))
	}

	cpPB := req.GetCallpath()
	if cpPB == nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to get basecall record callpath, it is nil"))
	}

	cp, err := cdrutil.ProtobufCallPathToCallPath(cpPB)
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
			return util.LogError(status.Errorf(codes.Unknown, "failed to receive basecall record for %s: %v", cp.UUID.String(), err))
		}

		chunk := req.GetRecordChunk()
		recordSize += len(chunk)

		_, err = recordData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to write chunk record: %v", err))
		}
	}

	err = c.db.AddRecordPath(stream.Context(), c.config.StorageAddress, &cp)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	res := &pb.SaveBaseCallResponse{Uuid: cp.UUID.String(), Size: uint32(recordSize)}
	pth := cdrutil.CreateHTTPCallPath(cp)

	// storing audio record into blob storage
	err = c.blobStore.CreateObject(context.Background(), c.config.Bucket, pth, recordData)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to store record into %s: %v", c.config.Bucket, err))
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response record %s: %v", cp.UUID.String(), err))
	}

	return nil
}
