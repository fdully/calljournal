package calljournal

import (
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/storage"
)

var (
	_ setup.BlobstoreConfigProvider = (*Config)(nil)
	_ setup.DatabaseConfigProvider  = (*Config)(nil)
)

type Config struct {
	Database   database.Config
	Blobstore  storage.Config
	Bucket     string `env:"CJ_STORAGE_BUCKET, required"`
	Addr       string `env:"CJ_CALLJOURNAL_ADDR, default=:8080"`
	CloudAddr  string `env:"CJ_CLOUD_ADDR, required"`
	TLSEnabled bool   `env:"CJ_TLS_ON, default=false"`
}

func (c *Config) DatabaseConfig() *database.Config {
	return &c.Database
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}
