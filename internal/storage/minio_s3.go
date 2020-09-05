package storage

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var _ Blobstore = (*Minio)(nil)

type minioConfig struct {
	Endpoint  string `env:"CJ_MINIO_ENDPOINT, default=localhost:9000"`
	AccessKey string `env:"CJ_MINIO_ACCESS_KEY, required"`
	SecretKey string `env:"CJ_MINIO_SECRET_KEY, required"`
}

func newMinio(ctx context.Context, config *minioConfig) (Blobstore, error) {
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

func (m *Minio) CreateObject(ctx context.Context, bucket, objectName string, contents *bytes.Buffer) error {
	if err := m.makeBucketIfNotExist(ctx, bucket); err != nil {
		return err
	}

	_, err := m.client.PutObject(ctx, bucket, objectName, contents,
		int64(contents.Len()), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("failed CreateObject storage: %w", err)
	}

	return nil
}

func (m *Minio) CreateFObject(ctx context.Context, bucket, objectName, fromFilename string) error {
	if err := m.makeBucketIfNotExist(ctx, bucket); err != nil {
		return err
	}

	_, err := m.client.FPutObject(ctx, bucket, objectName, fromFilename,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("failed CreateObject storage: %w", err)
	}

	return nil
}

func (m *Minio) GetObject(ctx context.Context, bucket, objectName string) (*bytes.Buffer, error) {
	o, err := m.client.GetObject(context.Background(), bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s %s, - %w", bucket, objectName, err)
	}
	defer o.Close()

	buffer := new(bytes.Buffer)

	_, err = buffer.ReadFrom(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s from storage: %w", objectName, err)
	}

	return buffer, nil
}

func (m *Minio) DeleteObject(ctx context.Context, bucket, objectName string) error {
	return m.client.RemoveObject(context.Background(), bucket, objectName, minio.RemoveObjectOptions{})
}

func (m *Minio) makeBucketIfNotExist(ctx context.Context, bucket string) error {
	ok, err := m.client.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}

	if !ok {
		if err := m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}

	return nil
}
