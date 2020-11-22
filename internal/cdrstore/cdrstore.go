package cdrstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/fdully/calljournal/internal/cdrserver/model"
	cdrstoredb "github.com/fdully/calljournal/internal/cdrstore/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/nsqio/go-nsq"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewStoreServer(env *serverenv.ServerEnv, config *Config) *StoreServer {
	return &StoreServer{
		config:     config,
		pool:       env.GRPCPool(),
		blobstore:  env.Blobstore(),
		db:         cdrstoredb.New(env.Database()),
		subscriber: env.Subscriber(),
	}
}

type StoreServer struct {
	config     *Config
	blobstore  storage.Blobstore
	subscriber queue.Subscribe
	db         *cdrstoredb.CDRStoreDB
	pool       *grpcpool.Pool
}

func (r *StoreServer) UploadMp3(ctx context.Context, grpcClient pb.CDRServiceClient) nsq.HandlerFunc {
	return func(m *nsq.Message) error {
		logger := logging.FromContext(ctx)

		if len(m.Body) == 0 {
			// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
			// In this case, a message with an empty body is simply ignored/discarded.
			return nil
		}

		var cp model.CallPath
		if err := json.Unmarshal(m.Body, &cp); err != nil {
			return fmt.Errorf("failed unmarshaling record info msg: %w", err)
		}

		stream, err := grpcClient.GetFile(ctx, cdrutil.CallPathToProtobufCallPath(cp))
		if err != nil {
			return err
		}

		data := new(bytes.Buffer)
		size := 0

		for {
			if err := util.ContextError(stream.Context()); err != nil {
				return err
			}

			req, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "failed to receive file for %s: %v", cp.NAME, err))
			}

			chunk := req.GetFileChunk()
			size += len(chunk)

			if _, err := data.Write(chunk); err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "failed to receive file for %s: %v", cp.NAME, err))
			}
		}

		mp3Data, err := Wav2Mp3(ctx, data)
		if err != nil {
			return util.LogError(err)
		}

		cp.NAME = util.ChangeWavExtToMp3(cp.NAME)
		pth := cdrutil.CreateHTTPCallPath(cp)

		logger.Infof("inserting record path into db %s", cp.UUID.String())

		if err := r.insertRecordPathToDatabase(ctx, cp); err != nil {
			logger.Errorf("failed to insert record path %s: %w", cp.NAME, err)

			return err
		}

		logger.Infof("uploading %s", cp.NAME)

		if err := r.blobstore.CreateObject(ctx, r.config.Bucket, pth, mp3Data); err != nil {
			return util.LogError(err)
		}

		return nil
	}
}

func (r *StoreServer) UploadCallFileToStorage(ctx context.Context, grpcClient pb.CDRServiceClient) nsq.HandlerFunc {
	return func(m *nsq.Message) error {
		logger := logging.FromContext(ctx)

		if len(m.Body) == 0 {
			// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
			// In this case, a message with an empty body is simply ignored/discarded.
			return nil
		}

		var cp model.CallPath
		if err := json.Unmarshal(m.Body, &cp); err != nil {
			return fmt.Errorf("failed unmarshaling callpath msg: %w", err)
		}

		stream, err := grpcClient.GetFile(ctx, cdrutil.CallPathToProtobufCallPath(cp))
		if err != nil {
			return err
		}

		data := new(bytes.Buffer)
		size := 0

		for {
			if err := util.ContextError(stream.Context()); err != nil {
				return err
			}

			req, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "failed to receive file for %s: %v", cp.NAME, err))
			}

			chunk := req.GetFileChunk()
			size += len(chunk)

			if _, err := data.Write(chunk); err != nil {
				return util.LogError(status.Errorf(codes.Unknown, "failed to receive file for %s: %v", cp.NAME, err))
			}
		}

		logger.Infof("inserting cdr into db %s", cp.UUID.String())

		if err := r.insertBaseCallToDatabase(ctx, data.String()); err != nil {
			logger.Errorf("failed to insert basecall %s: %w", cp.UUID.String(), err)

			return err
		}

		httpPth := cdrutil.CreateHTTPCallPath(cp)

		logger.Infof("uploading %s", cp.NAME)

		if err := r.blobstore.CreateObject(ctx, r.config.Bucket, httpPth, data); err != nil {
			logger.Errorf("failed to upload cdr %s: %w", cp.UUID.String(), err)

			return err
		}

		return nil
	}
}

func (r *StoreServer) SubscribeToRecord(ctx context.Context, grpcClient pb.CDRServiceClient) (queue.Stop, error) {
	if r.config.ConvertToMp3 {
		return r.subscriber(r.config.RecordTopic, r.config.RecordChannel, r.UploadMp3(ctx, grpcClient))
	}

	return r.subscriber(r.config.RecordTopic, r.config.RecordChannel, r.UploadCallFileToStorage(ctx, grpcClient))
}

func (r *StoreServer) SubscribeToCDR(ctx context.Context, grpcClient pb.CDRServiceClient) (queue.Stop, error) {
	return r.subscriber(r.config.CDRTopic, r.config.CDRChannel, r.UploadCallFileToStorage(ctx, grpcClient))
}

func (r *StoreServer) GetGRPCPool() *grpcpool.Pool {
	return r.pool
}

func (r *StoreServer) insertBaseCallToDatabase(ctx context.Context, cdrXML string) error {
	cdr, err := cdrutil.ParseCDR([]byte(cdrXML))
	if err != nil {
		return err
	}

	bc, err := cdrutil.CDRToBaseCall(cdr)
	if err != nil {
		return err
	}

	if err := r.db.AddBaseCall(ctx, bc); err != nil {
		return err
	}

	return nil
}

func (r *StoreServer) insertRecordPathToDatabase(ctx context.Context, cp model.CallPath) error {
	if err := r.db.AddRecordPath(ctx, r.config.StorageAddress, &cp); err != nil {
		return err
	}

	return nil
}
