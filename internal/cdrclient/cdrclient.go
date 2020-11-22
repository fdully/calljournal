package cdrclient

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/fdully/calljournal/internal/cdrclient/callfiles"
	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	grpcpool "github.com/processout/grpc-go-pool"
	"golang.org/x/sync/errgroup"
)

type Defer func()

type CDRClient struct {
	cf       callfiles.FilesInterface
	fileChan chan string
	pool     *grpcpool.Pool
	config   *Config
}

func NewCDRClient(env *serverenv.ServerEnv, config *Config) *CDRClient {
	return &CDRClient{
		cf:       callfiles.NewCallFiles(config.BaseCallDir, config.ReadDirPeriod),
		fileChan: make(chan string),
		pool:     env.GRPCPool(),
		config:   config,
	}
}

// Worker gets cdr filename from channel, reads it and sends it to remote server with grpc. If base call has
// audio record then sends it too. If there is error, worker will process same file again on next
// iteration.
//nolint:gocognit
func (c *CDRClient) Worker(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	gCDR, closeConn, err := c.createGRPCClients(ctx)
	if err != nil {
		return err
	}

	defer closeConn()

	for fname := range c.fileChan {
		cdrData, cdr, err := c.parseCDRFile(ctx, fname)
		if err != nil {
			c.cf.AgainLater(ctx, fname, err)

			continue
		}

		isRecordExist := c.IsRecordExist(ctx, cdr)

		// Delete small record file, less than 100KB
		if !isRecordExist && cdr.Variables.CJRecordName != "" {
			if err := c.cf.DeleteFile(ctx, string(cdr.Variables.CJRecordName)); err != nil {
				logger.Errorf("failed to delete record file %s: %w", cdr.Variables.CJRecordName, err)
			}
		}

		// ignore some files
		if !c.config.AllCDR && cdr.Variables.CJCdr != "true" {
			logger.Infof("ignoring and deleting cdr: %s", cdr.Variables.UUID)

			if isRecordExist {
				c.DeleteCDRFileAndRecordFile(ctx, fname, string(cdr.Variables.CJRecordName))

				continue
			}

			if !isRecordExist {
				c.DeleteCDRFile(ctx, fname)

				continue
			}
		}

		cdrPathInfo, err := cdrutil.CDRPathInfoFromCDR(ctx, cdr)
		if err != nil {
			logger.Debugf("failed to create cdr path info: %s", cdr.Variables.UUID)
			c.cf.AgainLater(ctx, fname, err)

			continue
		}

		// process only cdr data
		if !isRecordExist {
			if err := c.UploadCDR(ctx, cdrPathInfo, cdrData, gCDR); err != nil {
				c.cf.AgainLater(ctx, fname, err)

				continue
			} else {
				c.DeleteCDRFile(ctx, fname)

				continue
			}
		}

		// process cdr and record
		if isRecordExist {
			recordPathInfo, err := cdrutil.RecordPathInfoFromCDR(ctx, cdr)
			if err != nil {
				c.cf.AgainLater(ctx, fname, err)

				continue
			}

			if err := c.UploadCDRAndRecord(ctx, cdrPathInfo, recordPathInfo, cdrData,
				gCDR); err != nil {
				c.cf.AgainLater(ctx, fname, err)

				continue
			}

			c.DeleteCDRFileAndRecordFile(ctx, fname, recordPathInfo.NAME)
		}
	}

	return nil
}

func (c *CDRClient) RunCallFilesReader(ctx context.Context) error {
	return c.cf.ReadCDRFiles(ctx, c.fileChan)
}

func (c *CDRClient) UploadCDRAndRecord(ctx context.Context, cdr, record model.CallPath, cdrData []byte,
	gCDR pb.CDRServiceClient) error {
	recordData, err := c.cf.OpenFile(record.NAME)
	if err != nil {
		return err
	}

	ctx = context.Background()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return c.saveCDRData(ctx, cdr, cdrData, gCDR)
	})

	g.Go(func() error {
		return c.saveRecordData(ctx, record, recordData, gCDR)
	})

	return g.Wait()
}

func (c *CDRClient) UploadCDR(ctx context.Context, callpath model.CallPath,
	cdr []byte, gCDR pb.CDRServiceClient) error {
	if err := c.saveCDRData(ctx, callpath, cdr, gCDR); err != nil {
		return err
	}

	return nil
}

func (c *CDRClient) DeleteCDRFile(ctx context.Context, cdrFName string) {
	logger := logging.FromContext(ctx)
	logger.Infof("deleting cdr %s", cdrFName)

	if cdrFName != "" {
		if err := c.cf.DeleteFile(ctx, cdrFName); err != nil {
			logger.DPanic(err)
		}
	}
}

func (c *CDRClient) DeleteCDRFileAndRecordFile(ctx context.Context, cdrFName, recordFName string) {
	logger := logging.FromContext(ctx)
	logger.Infof("deleting cdr %s, record %s", cdrFName, recordFName)

	if cdrFName != "" {
		if err := c.cf.DeleteFile(ctx, cdrFName); err != nil {
			logger.DPanic(err)
		}
	}

	if recordFName != "" {
		if err := c.cf.DeleteFile(ctx, recordFName); err != nil {
			logger.DPanic(err)
		}
	}
}

func (c *CDRClient) saveCDRData(ctx context.Context, cp model.CallPath,
	cdrData []byte, gCDR pb.CDRServiceClient) error {
	stream, err := gCDR.SaveCDR(ctx)
	if err != nil {
		return fmt.Errorf("failed to create stream for cdr: %w", err)
	}

	req := &pb.SaveCDRRequest{
		Data: &pb.SaveCDRRequest_Callpath{
			Callpath: cdrutil.CallPathToProtobufCallPath(cp),
		},
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Errorf("failed to send record info: %w", err)
	}

	reader := bytes.NewReader(cdrData)
	buffer := make([]byte, GRPCMsgByteChunk)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read cdr: %w", err)
		}

		req := &pb.SaveCDRRequest{Data: &pb.SaveCDRRequest_CdrChunk{CdrChunk: buffer[:n]}}

		err = stream.Send(req)
		if err != nil {
			return fmt.Errorf("failed to send cdr %s: %w", cp.UUID.String(), err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to send cdr %s: %w", cp.UUID.String(), err)
	}

	logger := logging.FromContext(ctx)
	logger.Infof("saving cdr %s, size - %d", res.GetUuid(), res.GetSize())

	return nil
}

func (c *CDRClient) saveRecordData(ctx context.Context, cp model.CallPath,
	recordData []byte, gRecord pb.CDRServiceClient) error {
	stream, err := gRecord.SaveRecord(ctx)
	if err != nil {
		return fmt.Errorf("failed to create stream for audio record: %w", err)
	}

	req := &pb.SaveRecordRequest{
		Data: &pb.SaveRecordRequest_Callpath{
			Callpath: cdrutil.CallPathToProtobufCallPath(cp),
		},
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Errorf("failed to send record callpath %s: %w", cp.UUID.String(), err)
	}

	reader := bytes.NewReader(recordData)
	buffer := make([]byte, GRPCMsgByteChunk)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read audio record %s: %w", cp.UUID.String(), err)
		}

		req := &pb.SaveRecordRequest{Data: &pb.SaveRecordRequest_RecordChunk{RecordChunk: buffer[:n]}}

		err = stream.Send(req)
		if err != nil {
			return fmt.Errorf("failed to send audio record %s: %w", cp.UUID.String(), err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to send audio record %s: %w", cp.UUID.String(), err)
	}

	logger := logging.FromContext(ctx)
	logger.Infof("saving record %s, size - %d", res.GetUuid(), res.GetSize())

	return nil
}

func (c *CDRClient) parseCDRFile(ctx context.Context, fname string) ([]byte, model.CDR, error) {
	var cdr model.CDR

	cdrData, err := c.cf.OpenFile(fname)
	if err != nil {
		return nil, cdr, fmt.Errorf("failed to open fname %s: %w", fname, err)
	}

	cdr, err = cdrutil.ParseCDR(cdrData)
	if err != nil {
		return nil, cdr, fmt.Errorf("failed to parse fname %s: %w", fname, err)
	}

	return cdrData, cdr, err
}

func (c *CDRClient) createGRPCClients(ctx context.Context) (pb.CDRServiceClient, Defer, error) {
	conn, err := c.pool.Get(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get grpc conn from pool: %w", err)
	}

	return pb.NewCDRServiceClient(conn),
		func() {
			if conn != nil {
				conn.Close()
			}
		}, nil
}

func (c *CDRClient) IsRecordExist(ctx context.Context, cdr model.CDR) bool {
	const smallRecordFile = 100000

	return cdr.Variables.RecordFileSize > smallRecordFile
}
