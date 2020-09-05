package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

var (
	ErrNotFound      = errors.New("storage object not found")
	ErrUknownStorage = errors.New("unknown blob storage")
)

// Blobstore defines the minimum interface for a blob storage system.
type Blobstore interface {
	// CreateObject creates or overwrites an object in the storage system.
	CreateObject(ctx context.Context, bucket, objectName string, contents *bytes.Buffer) error

	// CreateFObject creates from file or overwrites an object in the storage system.
	CreateFObject(ctx context.Context, bucket, objectName string, fileName string) error

	// GetObject download object from storage system
	GetObject(ctx context.Context, bucket, objectName string) (*bytes.Buffer, error)

	// DeleteObject deltes an object or does nothing if the object doesn't exist.
	DeleteObject(ctx context.Context, bucket, objectName string) error
}

// BlobstoreFor returns the blob store for the given type, or an error if one does not exist.
func BlobstoreFor(ctx context.Context, typ BlobstoreType) (Blobstore, error) {
	switch typ {
	case BlobstoreTypeMinio:
		var config minioConfig

		err := envconfig.Process(ctx, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to process minio config %w", err)
		}

		return newMinio(ctx, &config)
	case BlobstoreTypeFilesystem:
		return newFilesystemStorage(ctx)
	default:
		return nil, fmt.Errorf("%w: %v", ErrUknownStorage, typ)
	}
}
