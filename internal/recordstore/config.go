package recordstore

import (
	"github.com/fdully/calljournal/internal/storage"
)

type Config struct {
	Blobstore         storage.Config
	Bucket            string `env:"CJ_STORAGE_BUCKET, required"`
	RecordInfoTopic   string `env:"CJ_RECORDINFO_TOPIC, default=recordinfo"`
	RecordInfoChannel string `env:"CJ_RECORDINFO_CHANNEL, default=recorduploader"`
	NsqdAddr          string `env:"CJ_NSQD_ADDR, default=localhost:4150"`
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}
