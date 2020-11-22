package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fdully/calljournal/internal/cdrstore"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/server"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting cdrstore")

	err := realMain(ctx)

	done()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting cdrstore")
}

func realMain(ctx context.Context) error {
	var config cdrstore.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	defer env.Close(ctx)

	// run prometheus metrics
	if err := server.ServeMetricsIfPrometheus(ctx); err != nil {
		return fmt.Errorf("failed to serve metrics: %w", err)
	}

	rs := cdrstore.NewStoreServer(env, &config)

	stops := make([]queue.Stop, 0, config.WorkersNum)

	for i := 0; i < config.WorkersNum; i++ {
		conn, err := rs.GetGRPCPool().Get(ctx)
		if err != nil {
			return err
		}

		grpcClient := pb.NewCDRServiceClient(conn)

		s, err := rs.SubscribeToRecord(ctx, grpcClient)
		if err != nil {
			return err
		}

		c, err := rs.SubscribeToCDR(ctx, grpcClient)
		if err != nil {
			return err
		}

		stops = append(stops, s, c)
	}

	defer rs.GetGRPCPool().Close()

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
