package storage

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

// BlobstoreType defines a specific blobstore.
type BlobstoreType string

const (
	BlobstoreTypeFilesystem BlobstoreType = "FILESYSTEM"
	BlobstoreTypeMinio      BlobstoreType = "MINIO"
)

// Config defines the configuration for a blobstore.
type Config struct {
	BlobstoreType BlobstoreType `env:"CJ_BLOBSTORE, default=MINIO"`
}

// Blobstore defines the minimum interface for a blob storage system.
type Blobstore interface {
	// CreateObject creates or overwrites an object in the storage system.
	CreateObject(ctx context.Context, bucket, objectName string, contents []byte) error

	// GetObject download object from storage system
	GetObject(ctx context.Context, bucket, objectName string) ([]byte, error)

	// DeleteObject deltes an object or does nothing if the object doesn't exist.
	DeleteObject(ctx context.Context, bucket, objectName string) error
}

// BlobstoreFor returns the blob store for the given type, or an error if one does not exist.
func BlobstoreFor(ctx context.Context, typ BlobstoreType) (Blobstore, error) {
	switch typ {
	case BlobstoreTypeMinio:
		var config MinioConfig
		err := envconfig.Process(ctx, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to process minio config %v", err)
		}
		return NewMinio(ctx, &config)
	case BlobstoreTypeFilesystem:
		return NewFilesystemStorage(ctx)
	default:
		return nil, fmt.Errorf("unknown blob store: %v", typ)
	}
}
