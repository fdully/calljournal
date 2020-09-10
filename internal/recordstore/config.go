package recordstore

import (
	"github.com/fdully/calljournal/internal/queue"
	"github.com/fdully/calljournal/internal/storage"
)

type Config struct {
	Blobstore         storage.Config
	Subscriber        queue.Config
	WorkersNum        int    `env:"CJ_RECORDSTORE_WORKERS_NUM, default=5"`
	Bucket            string `env:"CJ_STORAGE_BUCKET, required"`
	RecordInfoTopic   string `env:"CJ_RECORDINFO_TOPIC, default=recordinfo"`
	RecordInfoChannel string `env:"CJ_RECORDINFO_CHANNEL, default=recorduploader"`
	NsqdAddr          string `env:"CJ_NSQD_ADDR, default=localhost:4150"`
	ConvertToMp3      bool   `env:"CJ_CONVERT_MP3, default=true"`
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}

func (c *Config) SubscriberConfig() *queue.Config {
	return &c.Subscriber
}
