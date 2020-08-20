package calljournal

import (
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/storage"
)

var _ setup.BlobstoreConfigProvider = (*Config)(nil)

var _ setup.DatabaseConfigProvider = (*Config)(nil)

type Config struct {
	Database          database.Config
	Blobstore         storage.Config
	Bucket            string                `env:"CJ_STORAGE_BUCKET, required"`
	BlobstoreType     storage.BlobstoreType `env:"CJ_BLOBSTORE, default=MINIO"`
	GrpcServerAddress string                `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
}

func (c *Config) DatabaseConfig() *database.Config {
	return &c.Database
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}
