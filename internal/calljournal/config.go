package cdrserver

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
	Database  database.Config
	Blobstore storage.Config
	Bucket    string `env:"CJ_STORAGE_BUCKET, required"`
}

func (c *Config) DatabaseConfig() *database.Config {
	return &c.Database
}

func (c *Config) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}
