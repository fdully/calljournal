package storage

// BlobstoreType defines a specific blobstore.
type BlobstoreType string

const (
	BlobstoreTypeFilesystem BlobstoreType = "FILESYSTEM"
	BlobstoreTypeMinio      BlobstoreType = "MINIO"
)

// Config defines the configuration for a blobstore.
type Config struct {
	BlobstoreType BlobstoreType `env:"CJ_BLOBSTORE_TYPE,default=FILESYSTEM"`
}
