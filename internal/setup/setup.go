package setup

import (
	"context"
	"fmt"
	"time"

	"github.com/fdully/calljournal/internal/cdrstore"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	grpcpool "github.com/processout/grpc-go-pool"
	"github.com/sethvargo/go-envconfig"
	"google.golang.org/grpc"
)

// BlobstoreConfigProvider provides the information about current storage configuration.
type BlobstoreConfigProvider interface {
	BlobstoreConfig() *storage.Config
}

// DatabaseConfigProvider ensures that the environment config can provide a DB config.
type DatabaseConfigProvider interface {
	DatabaseConfig() *database.Config
}

type PublisherConfigProvider interface {
	PublisherConfig() *queue.Config
}

type SubscriberConfigProvider interface {
	SubscriberConfig() *queue.Config
}

type GRPCPoolConfigProvider interface {
	GRPCPoolConfig() (serverAddr string, dialOpt []grpc.DialOption, numWorkers int)
}

func Setup(ctx context.Context, config interface{}) (*serverenv.ServerEnv, error) {
	return WithSetup(ctx, config, envconfig.OsLookuper())
}

//nolint:funlen
func WithSetup(ctx context.Context, config interface{}, l envconfig.Lookuper) (*serverenv.ServerEnv, error) {
	logger := logging.FromContext(ctx)

	if err := envconfig.ProcessWith(ctx, config, l); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	if recordStoreConfig, ok := config.(cdrstore.Config); ok {
		if recordStoreConfig.ConvertToMp3 {
			logger.Debug("ping lame app")

			if err := cdrstore.PingLame(); err != nil {
				logger.Errorf("failed to find lame app: %v", err)

				return nil, err
			}
		}
	}

	var serverEnvOpts []serverenv.Option

	if provider, ok := config.(BlobstoreConfigProvider); ok {
		logger.Info("configuring blobstore")

		storageConfig := provider.BlobstoreConfig()

		blobStore, err := storage.BlobstoreFor(ctx, storageConfig.BlobstoreType)
		if err != nil {
			return nil, fmt.Errorf("failed to create blob storage: %w", err)
		}

		serverEnvOpts = append(serverEnvOpts, serverenv.WithBlobstore(blobStore))
	}

	if provider, ok := config.(DatabaseConfigProvider); ok {
		logger.Info("configuring database")

		dbConfig := provider.DatabaseConfig()

		db, err := database.NewDB(ctx, dbConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}

		// Update serverEnv setup.
		serverEnvOpts = append(serverEnvOpts, serverenv.WithDatabase(db))
	}

	if provider, ok := config.(PublisherConfigProvider); ok {
		logger.Info("configuring publisher")

		pubConfig := provider.PublisherConfig()

		p, err := queue.NewPub(ctx, pubConfig)
		if err != nil {
			return nil, err
		}

		serverEnvOpts = append(serverEnvOpts, serverenv.WithPublisher(p))
	}

	if provider, ok := config.(SubscriberConfigProvider); ok {
		logger.Info("configuring subscriber")

		subConfig := provider.SubscriberConfig()
		sub := queue.NewSubscriber(subConfig)

		serverEnvOpts = append(serverEnvOpts, serverenv.WithSubscriber(sub))
	}

	if provider, ok := config.(GRPCPoolConfigProvider); ok {
		logger.Info("configuring grpc pool")

		serverAddr, dialOpt, numWorkers := provider.GRPCPoolConfig()

		factory := func() (*grpc.ClientConn, error) {
			conn, err := grpc.DialContext(ctx, serverAddr, dialOpt...)
			if err != nil {
				return nil, fmt.Errorf("failed on grpc connection: %w", err)
			}

			return conn, err
		}

		pool, err := grpcpool.New(factory, numWorkers, numWorkers, time.Second)
		if err != nil {
			return nil, fmt.Errorf("failed to create gRPC pool: %w", err)
		}

		serverEnvOpts = append(serverEnvOpts, serverenv.WithGRPCPool(pool))
	}

	return serverenv.NewServerEnv(ctx, serverEnvOpts...), nil
}
