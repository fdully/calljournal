package cdrstore

import (
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/storage"
	"google.golang.org/grpc"
)

const GRPCMaxMsgSize = 5 * 1024 * 1024 * 1024 * 1024

var (
	_ setup.BlobstoreConfigProvider  = (*Config)(nil)
	_ setup.DatabaseConfigProvider   = (*Config)(nil)
	_ setup.SubscriberConfigProvider = (*Config)(nil)
)

type Config struct {
	Blobstore      storage.Config
	Database       database.Config
	Subscriber     queue.Config
	GrpcServerAddr string `env:"CJ_GRPC_SERVER_ADDR, default=localhost:9111"`
	WorkersNum     int    `env:"CJ_CDRSTORE_WORKERS, default=5"`
	StorageAddress string `env:"CJ_STORAGE_ADDR, required"`
	Bucket         string `env:"CJ_STORAGE_BUCKET, required"`
	RecordTopic    string `env:"CJ_RECORD_TOPIC, default=record"`
	RecordChannel  string `env:"CJ_RECORD_CHANNEL, default=recordstore"`
	CDRTopic       string `env:"CJ_CDR_TOPIC, default=cdr"`
	CDRChannel     string `env:"CJ_CDR_CHANNEL, default=cdrstore"`
	NsqdAddr       string `env:"CJ_NSQD_ADDR, default=localhost:4150"`
	ConvertToMp3   bool   `env:"CJ_CONVERT_MP3, default=true"`
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}

func (c *Config) SubscriberConfig() *queue.Config {
	return &c.Subscriber
}

func (c *Config) DatabaseConfig() *database.Config {
	return &c.Database
}

func (c *Config) GRPCPoolConfig() (string, []grpc.DialOption, int) {
	var dialOpt []grpc.DialOption
	dialOpt = append(dialOpt, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(GRPCMaxMsgSize),
		grpc.MaxCallSendMsgSize(GRPCMaxMsgSize)), grpc.WithInsecure())

	return c.GrpcServerAddr, dialOpt, c.WorkersNum
}
