package calljournal

import (
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/storage"
)

var (
	_ setup.BlobstoreConfigProvider = (*Config)(nil)
	_ setup.DatabaseConfigProvider  = (*Config)(nil)
	_ setup.PublisherConfigProvider = (*Config)(nil)
)

type Config struct {
	Database          database.Config
	Blobstore         storage.Config
	Publisher         queue.Config
	Bucket            string `env:"CJ_STORAGE_BUCKET, required"`
	GrpcServerAddress string `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
	BaseCallTopic     string `env:"CJ_BASECALL_TOPIC, default=basecall"`
	RecordInfoTopic   string `env:"CJ_RECORDINFO_TOPIC, default=recordinfo"`
}

func (c *Config) DatabaseConfig() *database.Config {
	return &c.Database
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}

func (c *Config) PublisherConfig() *queue.Config {
	return &c.Publisher
}
