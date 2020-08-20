package calluploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fdully/calljournal/internal/calljournal"
	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/calluploader/callfiles"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/util"
)

type CallUploader struct {
	cf             callfiles.FilesInterface
	fileChan       chan os.FileInfo
	baseCallClient pb.BaseCallServiceClient
	recordClient   pb.AudioRecordServiceClient
	config         *Config
}

func NewCallUploader(config *Config, bcClient pb.BaseCallServiceClient,
	recordClient pb.AudioRecordServiceClient) *CallUploader {
	return &CallUploader{
		cf:             callfiles.NewCallFiles(config.BaseCallDir, config.ReadDirPeriod),
		fileChan:       make(chan os.FileInfo),
		baseCallClient: bcClient,
		recordClient:   recordClient,
		config:         config,
	}
}

// Worker gets base call file name from channel, reads it and sends it to remote server with grpc. If base call has
// audio record then sends it too. If there is error, worker will process same file again on next
// iteration.
func (c *CallUploader) Worker(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	logger.Info("starting worker")

	for fname := range c.fileChan {
		f, err := c.cf.OpenFile(fname.Name())
		if err != nil {
			logger.DPanicf("failed to open basecall file %s: %v", fname.Name(), err)
			c.cf.DoItAgainLater(fname.Name())

			continue
		}

		bc, err := util.ParseCall(ctx, f)
		if err != nil {
			logger.DPanic("failed to parse basecall file %s: %v", fname.Name(), err)
			c.cf.DoItAgainLater(fname.Name())

			continue
		}

		// process only basecall
		if !bc.RECD && bc.RECL == "" {
			err := c.saveBaseCall(ctx, bc)
			if err != nil {
				logger.DPanic(err)
				c.cf.DoItAgainLater(fname.Name())

				continue
			}

			err = c.cf.DeleteFile(fname.Name())
			if err != nil {
				logger.DPanic(err)
			}

			continue
		}

		// process basecall and recording
		recordInfo := util.CreateRecordInfo(bc)
		bc.RECL = c.config.StorageAddr + "/" + util.ChangeWavExtToMp3(util.CreateRecordPath(recordInfo))

		err = c.saveBaseCall(ctx, bc)
		if err != nil {
			logger.DPanic(err)
			c.cf.DoItAgainLater(fname.Name())

			continue
		}

		err = c.saveRecord(ctx, recordInfo)
		if err != nil {
			logger.DPanic(err)
			c.cf.DoItAgainLater(fname.Name())

			continue
		}

		// remove json file
		err = c.cf.DeleteFile(fname.Name())
		if err != nil {
			logger.DPanic(err)
		}

		// remove audio file
		err = c.cf.DeleteFile(recordInfo.Name)
		if err != nil {
			logger.DPanic(err)
		}
	}

	return nil
}

func (c *CallUploader) RunCallFilesReader(ctx context.Context) error {
	return c.cf.ReadBaseCallsFromDir(ctx, c.fileChan)
}

func (c *CallUploader) saveBaseCall(ctx context.Context, bc *model.BaseCall) error {
	pbBC, err := calljournal.BaseCallToProtobufBaseCall(bc)
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

func (c *CallUploader) saveRecord(ctx context.Context, recordInfo model.RecordInfo) error {
	record, err := c.cf.OpenFile(recordInfo.Name)
	if err != nil {
		return fmt.Errorf("failed to open audio file %s: %w", recordInfo.Name, err)
	}

	stream, err := c.recordClient.SaveAudioRecord(ctx)
	if err != nil {
		return fmt.Errorf("failed to create stream for audio record: %w", err)
	}

	req := &pb.SaveAudioRecordRequest{
		Data: &pb.SaveAudioRecordRequest_Info{
			Info: util.RecordInfoToProtobufRecordInfo(recordInfo),
		},
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Errorf("failed to send record info: %w", err)
	}

	reader := bytes.NewReader(record)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read audio record: %w", err)
		}

		req := &pb.SaveAudioRecordRequest{Data: &pb.SaveAudioRecordRequest_RecordChunk{RecordChunk: buffer[:n]}}

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
