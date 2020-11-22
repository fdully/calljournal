package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fdully/calljournal/internal/cdrclient"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/server"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()
	defer done()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting cdrclient")

	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting cdrclient")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var config cdrclient.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	defer env.Close(ctx)

	cdr := cdrclient.NewCDRClient(env, &config)

	if err := server.ServeMetricsIfPrometheus(ctx); err != nil {
		return fmt.Errorf("failed to serve metrics: %w", err)
	}

	// Reader reads cdr from directory and sends it to workers for upload to cdr server.
	go func() {
		err := cdr.RunCallFilesReader(ctx)
		if err != nil {
			logger.Errorf("failed to run cdr files reader: %v", err)
		}
	}()

	var wg sync.WaitGroup

	// Run cdr upload workers in parallel
	logger.Infof("running %d workers", config.NumWorkers)

	wg.Add(config.NumWorkers)

	for i := 0; i < config.NumWorkers; i++ {
		i := i

		go func() {
			logger.Infof("run worker %d", i)

			defer wg.Done()

			// Worker uploads cdr and audio record to grpc server.
			// If failed then files will be read again later.
			err := cdr.Worker(ctx)
			if err != nil {
				logger.Errorf("failed to run worker: %v", err)
			}

			logger.Infof("exiting worker %d", i)
		}()
	}

	// Wait for all goroutines to finish their work.
	wg.Wait()

	return nil
}
