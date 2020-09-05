package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/recordstore"
	"github.com/fdully/calljournal/internal/server"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/nsqio/go-nsq"
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
	if err := recordstore.PingLame(); err != nil {
		return err
	}

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

	// Instantiate a consumer that will subscribe to the provided channel.
	cnf := nsq.NewConfig()
	cnf.MaxAttempts = 1000

	consumer, err := nsq.NewConsumer(config.RecordInfoTopic, config.RecordInfoChannel, cnf)
	if err != nil {
		return err
	}

	rs := recordstore.NewStoreServer(env, &config)

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(rs.HandleMessageWithContext(ctx))

	err = consumer.ConnectToNSQD(config.NsqdAddr)
	if err != nil {
		return err
	}

	defer consumer.Stop()

	<-ctx.Done()
	time.Sleep(time.Second * 1)

	return nil
}
