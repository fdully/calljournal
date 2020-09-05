package setup

import (
	"context"
	"fmt"

	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/sethvargo/go-envconfig"
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

func Setup(ctx context.Context, config interface{}) (*serverenv.ServerEnv, error) {
	return WithSetup(ctx, config, envconfig.OsLookuper())
}

func WithSetup(ctx context.Context, config interface{}, l envconfig.Lookuper) (*serverenv.ServerEnv, error) {
	logger := logging.FromContext(ctx)

	if err := envconfig.ProcessWith(ctx, config, l); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
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

	return serverenv.NewServerEnv(ctx, serverEnvOpts...), nil
}
