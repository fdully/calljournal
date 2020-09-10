package recordstore

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util"
	"github.com/nsqio/go-nsq"
)

func NewStoreServer(env *serverenv.ServerEnv, config *Config) *StoreServer {
	return &StoreServer{
		config:     config,
		blobstore:  env.Blobstore(),
		subscriber: env.Subscriber(),
	}
}

type StoreServer struct {
	config     *Config
	blobstore  storage.Blobstore
	subscriber queue.Subscribe
}

// HandleMessage implements the Handler interface.
func (r *StoreServer) UploadMp3(ctx context.Context) nsq.HandlerFunc {
	return func(m *nsq.Message) error {
		logger := logging.FromContext(ctx)

		if len(m.Body) == 0 {
			// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
			// In this case, a message with an empty body is simply ignored/discarded.
			return nil
		}

		var ri model.RecordInfo
		if err := json.Unmarshal(m.Body, &ri); err != nil {
			return fmt.Errorf("failed unmarshaling record info msg: %w", err)
		}

		wavPath := filepath.Join(r.config.Bucket, ri.DIRC, ri.YEAR, ri.MONT, ri.RDAY, ri.RNAM)

		mp3Path := util.ChangeWavExtToMp3(wavPath)
		if err := WavFileToMp3File(ctx, wavPath, mp3Path); err != nil {
			return util.LogError(err)
		}

		ri.RNAM = util.ChangeWavExtToMp3(ri.RNAM)
		pth := util.CreateHTTPRecordPath(ri)

		if err := r.blobstore.CreateFObject(ctx, r.config.Bucket, pth, mp3Path); err != nil {
			return util.LogError(err)
		}

		logger.Infof("deleting mp3 file after uploading to storage: %s", ri.UUID.String())

		if err := os.Remove(mp3Path); err != nil {
			logger.DPanicf("failed to remove mp3 file %s: %v", ri.UUID.String(), err)
		}

		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return nil
	}
}

func (r *StoreServer) UploadWav(ctx context.Context) nsq.HandlerFunc {
	return func(m *nsq.Message) error {
		logger := logging.FromContext(ctx)

		if len(m.Body) == 0 {
			// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
			// In this case, a message with an empty body is simply ignored/discarded.
			return nil
		}

		var ri model.RecordInfo
		if err := json.Unmarshal(m.Body, &ri); err != nil {
			return fmt.Errorf("failed unmarshaling record info msg: %w", err)
		}

		wavPath := filepath.Join(r.config.Bucket, ri.DIRC, ri.YEAR, ri.MONT, ri.RDAY, ri.RNAM)

		pth := util.CreateHTTPRecordPath(ri)

		if err := r.blobstore.CreateFObject(ctx, r.config.Bucket, pth, wavPath); err != nil {
			return util.LogError(err)
		}

		logger.Infof("uploaded wav file: %s", ri.UUID.String())

		return nil
	}
}

func (r *StoreServer) Subscribe(ctx context.Context) (queue.Stop, error) {
	if r.config.ConvertToMp3 {
		return r.subscriber(r.config.RecordInfoTopic, r.config.RecordInfoChannel, r.UploadMp3(ctx))
	}

	return r.subscriber(r.config.RecordInfoTopic, r.config.RecordInfoChannel, r.UploadWav(ctx))
}
