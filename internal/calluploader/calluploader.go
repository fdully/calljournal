package calluploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/calluploader/callfiles"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/util"
	"golang.org/x/sync/errgroup"
)

type CallUploader struct {
	cf               callfiles.FilesInterface
	fileChan         chan os.FileInfo
	baseCallClient   pb.BaseCallServiceClient
	recordInfoClient pb.RecordInfoServiceClient
	recordDataClient pb.RecordDataServiceClient
	config           *Config
}

func NewCallUploader(config *Config, bcClient pb.BaseCallServiceClient,
	recordInfoClient pb.RecordInfoServiceClient, recordDataClient pb.RecordDataServiceClient) *CallUploader {
	return &CallUploader{
		cf:               callfiles.NewCallFiles(config.BaseCallDir, config.ReadDirPeriod),
		fileChan:         make(chan os.FileInfo),
		baseCallClient:   bcClient,
		recordInfoClient: recordInfoClient,
		recordDataClient: recordDataClient,
		config:           config,
	}
}

// Worker gets base call file name from channel, reads it and sends it to remote server with grpc. If base call has
// audio record then sends it too. If there is error, worker will process same file again on next
// iteration.
func (c *CallUploader) Worker(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	logger.Info("starting worker")

	for fname := range c.fileChan {
		bc, err := c.baseCallFromFile(ctx, fname.Name())
		if err != nil {
			c.cf.DoItAgainLater(ctx, fname.Name(), err)

			continue
		}

		switch {
		// process only basecall
		case !bc.RECD && bc.RECS == 0:
			{
				if err := c.UploadBaseCall(ctx, bc); err != nil {
					c.cf.DoItAgainLater(ctx, fname.Name(), err)
				} else {
					c.DeleteBaseCallFile(ctx, fname.Name())
				}
			}
		// process base call and record
		case bc.RECD && bc.RNAM != "":
			{
				recordInfo := util.CreateRecordInfo(bc, c.config.StorageAddr)

				recordData, err := c.cf.OpenFile(recordInfo.RNAM)
				if err != nil {
					c.cf.DoItAgainLater(ctx, fname.Name(), err)

					continue
				}

				if err := c.UploadBaseCallAndRecord(ctx, bc, recordInfo, recordData); err != nil {
					c.cf.DoItAgainLater(ctx, fname.Name(), err)

					continue
				}

				c.DeleteBaseCallFileAndRecordFile(ctx, fname.Name(), recordInfo.RNAM)
			}
		default:
			logger.DPanicf("No case for basecall file upload: %s", fname.Name())
		}
	}

	logger.Info("call file channel is closed, exiting from worker")

	return nil
}

func (c *CallUploader) RunCallFilesReader(ctx context.Context) error {
	return c.cf.ReadBaseCallsFromDir(ctx, c.fileChan)
}

func (c *CallUploader) UploadBaseCallAndRecord(ctx context.Context, bc *model.BaseCall,
	info model.RecordInfo, recordData []byte) error {
	ctx = context.Background()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return c.saveBaseCall(ctx, bc)
	})

	g.Go(func() error {
		return c.saveRecordData(ctx, info, recordData)
	})

	g.Go(func() error {
		return c.saveRecordInfo(ctx, info)
	})

	return g.Wait()
}

func (c *CallUploader) UploadBaseCall(ctx context.Context, bc *model.BaseCall) error {
	if err := c.saveBaseCall(ctx, bc); err != nil {
		return err
	}

	return nil
}

func (c *CallUploader) DeleteBaseCallFile(ctx context.Context, baseCallFileName string) {
	logger := logging.FromContext(ctx)

	if err := c.cf.DeleteFile(baseCallFileName); err != nil {
		logger.DPanic(err)
	}
}

func (c *CallUploader) DeleteBaseCallFileAndRecordFile(ctx context.Context, baseCallFileName, recordInfoFileName string) {
	logger := logging.FromContext(ctx)

	if err := c.cf.DeleteFile(baseCallFileName); err != nil {
		logger.DPanic(err)
	}

	if err := c.cf.DeleteFile(recordInfoFileName); err != nil {
		logger.DPanic(err)
	}
}

func (c *CallUploader) saveBaseCall(ctx context.Context, bc *model.BaseCall) error {
	pbBC, err := util.BaseCallToProtobufBaseCall(bc)
	if err != nil {
		return err
	}

	req := &pb.SaveBaseCallRequest{BaseCall: pbBC}

	res, err := c.baseCallClient.SaveBaseCall(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to save basecall %s: %w", pbBC.Uuid, err)
	}

	logger := logging.FromContext(ctx)
	logger.Infof("basecall with id - %s is saved", res.GetUuid())

	return nil
}

func (c *CallUploader) saveRecordInfo(ctx context.Context, info model.RecordInfo) error {
	pbInfo := util.RecordInfoToProtobufRecordInfo(info)

	req := &pb.SaveRecordInfoRequest{Info: pbInfo}

	res, err := c.recordInfoClient.SaveRecordInfo(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to save record info %s: %w", pbInfo.Uuid, err)
	}

	logger := logging.FromContext(ctx)
	logger.Infof("record info with id - %s is saved", res.GetUuid())

	return nil
}

func (c *CallUploader) saveRecordData(ctx context.Context, recordInfo model.RecordInfo, recordData []byte) error {
	stream, err := c.recordDataClient.SaveRecordData(ctx)
	if err != nil {
		return fmt.Errorf("failed to create stream for audio record: %w", err)
	}

	req := &pb.SaveRecordDataRequest{
		Data: &pb.SaveRecordDataRequest_Info{
			Info: util.RecordInfoToProtobufRecordInfo(recordInfo),
		},
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Errorf("failed to send record info: %w", err)
	}

	reader := bytes.NewReader(recordData)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read audio record: %w", err)
		}

		req := &pb.SaveRecordDataRequest{Data: &pb.SaveRecordDataRequest_RecordChunk{RecordChunk: buffer[:n]}}

		err = stream.Send(req)
		if err != nil {
			return fmt.Errorf("failed to send audio record: %w", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to send audio record: %w", err)
	}

	logger := logging.FromContext(ctx)
	logger.Infof("audio record is saved to server id - %s, size - %d", res.GetUuid(), res.GetSize())

	return nil
}

func (c *CallUploader) baseCallFromFile(ctx context.Context, fname string) (*model.BaseCall, error) {
	f, err := c.cf.OpenFile(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open basecall file %s: %w", fname, err)
	}

	bc, err := util.ParseCall(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to parse basecall file %s: %w", fname, err)
	}

	return bc, nil
}
