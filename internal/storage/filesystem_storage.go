package storage

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ Blobstore = (*FilesystemStorage)(nil)

// FilesystemStorage implements Blobstore and provides the ability write files to the filesystem.
type FilesystemStorage struct{}

// NewFilesystemStorage creates a Blobsstore compatible storage for the filesystem.
func NewFilesystemStorage(ctx context.Context) (Blobstore, error) {
	return &FilesystemStorage{}, nil
}

// CreateObject creates a new object on the filesystem or overwrites an existing one.
func (s *FilesystemStorage) CreateObject(ctx context.Context, folder, filename string, contents []byte) error {
	pth := filepath.Join(folder, filename)
	if err := ioutil.WriteFile(pth, contents, 0644); err != nil {
		return fmt.Errorf("failed to create object: %w", err)
	}
	return nil
}

// GetObject download object from storage system
func (s *FilesystemStorage) GetObject(ctx context.Context, folder, filename string) ([]byte, error) {
	pth := filepath.Join(folder, filename)
	return ioutil.ReadFile(pth)
}

// DeleteObject deletes an object from the filesystem. It returns nil if the object was deleted or if the object no longer exists.
func (s *FilesystemStorage) DeleteObject(ctx context.Context, folder, filename string) error {
	pth := filepath.Join(folder, filename)
	if err := os.Remove(pth); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}
