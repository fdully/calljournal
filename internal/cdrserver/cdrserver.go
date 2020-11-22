package cdrserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	GRPCMsgByteChunk = 64 * 1024
)

func NewCDRServer(env *serverenv.ServerEnv, config *Config) *CDRServer {
	return &CDRServer{
		blobStore: env.Blobstore(),
		publisher: env.Publisher(),
		config:    config,
	}
}

var _ pb.CDRServiceServer = (*CDRServer)(nil)

type CDRServer struct {
	blobStore storage.Blobstore
	publisher queue.Publisher
	config    *Config
}

func (c *CDRServer) GetFile(cp *pb.CallPath, stream pb.CDRService_GetFileServer) error {
	buf, err := c.readDataFromStorage(cp)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to read data from storage: %v", err))
	}

	buffer := make([]byte, GRPCMsgByteChunk)

	for {
		n, err := buf.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to read data: %v", err))
		}

		res := &pb.GetFileResponse{FileChunk: buffer[:n]}

		err = stream.Send(res)
		if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "failed to send data response: %v", err))
		}
	}

	return nil
}

func (c *CDRServer) SaveCDR(stream pb.CDRService_SaveCDRServer) error {
	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to receive cdr %v", err))
	}

	cpPB := req.GetCallpath()
	if cpPB == nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to get callpath, it is nil"))
	}

	cp, err := cdrutil.ProtobufCallPathToCallPath(cpPB)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "%v", err))
	}

	cdrData := new(bytes.Buffer)
	cdrSize := 0

	for {
		if err := util.ContextError(stream.Context()); err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return util.LogError(status.Errorf(codes.Unknown, "failed to receive cdr data for %s: %v", cpPB.Uuid, err))
		}

		chunk := req.GetCdrChunk()
		cdrSize += len(chunk)

		_, err = cdrData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to write chunk cdr %s: %v", cp.UUID.String(), err))
		}
	}

	res := &pb.SaveCDRResponse{Uuid: cpPB.Uuid, Size: uint32(cdrSize)}
	pth := cdrutil.CreateFileCallPath(cp)

	// storing audio record in wav to temporary file storage
	err = c.blobStore.CreateObject(context.Background(), c.config.Bucket, pth, cdrData)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to store cdr: %v", err))
	}

	if err := c.publishCallPath(c.config.CDRTopic, cp); err != nil {
		return util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response cdr %s: %v", cp.UUID.String(), err))
	}

	return nil
}

func (c *CDRServer) SaveRecord(stream pb.CDRService_SaveRecordServer) error {
	req, err := stream.Recv()
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to receive record callpath %v", err))
	}

	cpPB := req.GetCallpath()
	if cpPB == nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to get record callpath, it is nil"))
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
			return util.LogError(status.Errorf(codes.Unknown, "failed to receive record data for %s: %v", cp.UUID.String(), err))
		}

		chunk := req.GetRecordChunk()
		recordSize += len(chunk)

		_, err = recordData.Write(chunk)
		if err != nil {
			return util.LogError(status.Errorf(codes.Internal, "failed to write chunk record: %v", err))
		}
	}

	res := &pb.SaveCDRResponse{Uuid: cp.UUID.String(), Size: uint32(recordSize)}
	pth := cdrutil.CreateFileCallPath(cp)

	// storing audio record in wav to temporary file storage
	err = c.blobStore.CreateObject(context.Background(), c.config.Bucket, pth, recordData)
	if err != nil {
		return util.LogError(status.Errorf(codes.Internal, "failed to store record into %s: %v", c.config.Bucket, err))
	}

	if err := c.publishCallPath(c.config.RecordTopic, cp); err != nil {
		return util.LogError(status.Errorf(codes.Internal, "%v", err))
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return util.LogError(status.Errorf(codes.Unknown, "failed to send and close response record %s: %v", cp.UUID.String(), err))
	}

	return nil
}

func (c *CDRServer) publishCallPath(topic string, cp model.CallPath) error {
	msg, err := json.Marshal(cp)
	if err != nil {
		return fmt.Errorf("failed to marshal json callpath %s: %w", cp.NAME, err)
	}

	if err := c.publisher.Publish(topic, msg); err != nil {
		return fmt.Errorf("failed to publish callpath %s: %w", cp.NAME, err)
	}

	return nil
}

func (c *CDRServer) readDataFromStorage(cp *pb.CallPath) (*bytes.Buffer, error) {
	callpath, err := cdrutil.ProtobufCallPathToCallPath(cp)
	if err != nil {
		return nil, err
	}

	pth := cdrutil.CreateFileCallPath(callpath)

	buf, err := c.blobStore.GetObject(context.Background(), c.config.Bucket, pth)
	if err != nil {
		return nil, err
	}

	return buf, err
}
