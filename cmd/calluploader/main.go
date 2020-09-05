package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fdully/calljournal/internal/calluploader"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/server"
	"github.com/sethvargo/go-envconfig"
	"github.com/sethvargo/go-signalcontext"
	"google.golang.org/grpc"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()
	defer done()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting calluploader")

	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting calluploader")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var config calluploader.Config

	err := envconfig.Process(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed to process config: %w", err)
	}

	const workerCount = 5

	transportOption := grpc.WithInsecure()

	clientConn, err := grpc.DialContext(ctx, config.GrpcServerAddr, transportOption)
	if err != nil {
		return fmt.Errorf("failed on grpc connection: %w", err)
	}

	bcClient := pb.NewBaseCallServiceClient(clientConn)
	recordInfoClient := pb.NewRecordInfoServiceClient(clientConn)
	recordDataClient := pb.NewRecordDataServiceClient(clientConn)

	cu := calluploader.NewCallUploader(&config, bcClient, recordInfoClient, recordDataClient)

	if err := server.ServeMetricsIfPrometheus(ctx); err != nil {
		return fmt.Errorf("failed to serve metrics: %w", err)
	}

	// Reader reads basecalls from directory and sends it to channel.
	go func() {
		err := cu.RunCallFilesReader(ctx)
		if err != nil {
			logger.Errorf("failed to run basecall files reader: %v", err)
		}
	}()

	var wg sync.WaitGroup

	// Run call upload workers in parallel
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()

			// Worker uploads basecall and audio record to grpc server.
			// If failed then files will be read again later.
			err := cu.Worker(ctx)
			if err != nil {
				logger.Errorf("failed to run worker: %v", err)
			}
		}()
	}

	// Wait for all goroutines to finish their work.
	wg.Wait()

	return nil
}
