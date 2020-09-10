package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/recordstore"
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

	logger.Info("starting recordstore")

	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting recordstore")
}

func realMain(ctx context.Context) error {
	var config recordstore.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	defer env.Close(ctx)

	// run prometheus metrics
	if err := server.ServeMetricsIfPrometheus(ctx); err != nil {
		return fmt.Errorf("failed to serve metrics: %w", err)
	}

	rs := recordstore.NewStoreServer(env, &config)

	stops := make([]queue.Stop, 0, config.WorkersNum)

	for i := 0; i < config.WorkersNum; i++ {
		s, err := rs.Subscribe(ctx)
		if err != nil {
			return err
		}

		stops = append(stops, s)
	}

	<-ctx.Done()

	var wg sync.WaitGroup

	for _, stop := range stops {
		if stop == nil {
			break
		}

		wg.Add(1)

		go func(stop func()) {
			defer wg.Done()
			stop()
		}(stop)
	}

	wg.Wait()

	return nil
}
