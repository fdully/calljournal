package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var _ Blobstore = (*Minio)(nil)

type MinioConfig struct {
	Endpoint  string `env:"CJ_MINIO_ENDPOINT, default=localhost:9000"`
	AccessKey string `env:"CJ_MINIO_ACCESS_KEY, required"`
	SecretKey string `env:"CJ_MINIO_SECRET_KEY, required"`
}

func NewMinio(ctx context.Context, config *MinioConfig) (Blobstore, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio blob storage client: %w", err)
	}

	return &Minio{client}, nil
}

type Minio struct {
	client *minio.Client
}

func (m *Minio) CreateObject(ctx context.Context, bucket, objectName string, contents []byte) error {
	_, err := m.client.PutObject(ctx, bucket, objectName, bytes.NewReader(contents), int64(len(contents)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("failed CreateObject storage: %w", err)
	}
	return nil
}

func (m *Minio) GetObject(ctx context.Context, bucket, objectName string) ([]byte, error) {
	o, err := m.client.GetObject(context.Background(), bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s %s, - %v", bucket, objectName, err)
	}
	return ioutil.ReadAll(o)
}

func (m *Minio) DeleteObject(ctx context.Context, bucket, objectName string) error {
	return m.client.RemoveObject(context.Background(), bucket, objectName, minio.RemoveObjectOptions{})
}
