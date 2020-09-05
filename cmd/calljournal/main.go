package main

import (
	"context"
	"fmt"

	"github.com/fdully/calljournal/internal/calljournal"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/server"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/sethvargo/go-signalcontext"
	"google.golang.org/grpc"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()
	defer done()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting calljournal")

	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting calljournal")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var config calljournal.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	defer env.Close(ctx)

	baseCallServer := calljournal.NewBaseCallServer(env, &config)

	grpcServer := grpc.NewServer()
	pb.RegisterBaseCallServiceServer(grpcServer, baseCallServer)
	pb.RegisterRecordInfoServiceServer(grpcServer, baseCallServer)
	pb.RegisterRecordDataServiceServer(grpcServer, baseCallServer)

	srv, err := server.NewServer(config.GrpcServerAddress)
	if err != nil {
		return fmt.Errorf("failed to create NewServer: %w", err)
	}

	logger.Infof("listen on %s", config.GrpcServerAddress)

	if err := server.ServeMetricsIfPrometheus(ctx); err != nil {
		return fmt.Errorf("failed to serve metrics: %w", err)
	}

	return srv.ServeGRPC(ctx, grpcServer)
}
