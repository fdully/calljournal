package main

import (
	"context"
	"fmt"

	"github.com/fdully/calljournal/internal/cdrserver"
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

	logger.Info("starting cdrserver")

	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting cdrserver")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var config cdrserver.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	defer env.Close(ctx)

	baseCallServer := cdrserver.NewCDRServer(env, &config)

	var options []grpc.ServerOption
	options = append(options, grpc.MaxSendMsgSize(cdrserver.GRPCMaxMsgSize), grpc.MaxRecvMsgSize(cdrserver.GRPCMaxMsgSize))
	grpcServer := grpc.NewServer(options...)

	pb.RegisterCDRServiceServer(grpcServer, baseCallServer)

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
