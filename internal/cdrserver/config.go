package cdrserver

import (
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/storage"
)

var (
	_ setup.BlobstoreConfigProvider = (*Config)(nil)
	_ setup.PublisherConfigProvider = (*Config)(nil)
)

const (
	GRPCMaxMsgSize = 5 * 1024 * 1024 * 1024 * 1024
)

type Config struct {
	Blobstore         storage.Config
	Publisher         queue.Config
	Bucket            string `env:"CJ_STORAGE_BUCKET, required"`
	GrpcServerAddress string `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
	CDRTopic          string `env:"CJ_CDR_TOPIC, default=cdr"`
	RecordTopic       string `env:"CJ_RECORD_TOPIC, default=record"`
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}

func (c *Config) PublisherConfig() *queue.Config {
	return &c.Publisher
}
