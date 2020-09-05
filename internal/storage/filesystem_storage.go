package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/tools/apidiff/ioext"
)

var _ Blobstore = (*FilesystemStorage)(nil)

// FilesystemStorage implements Blobstore and provides the ability write files to the filesystem.
type FilesystemStorage struct{}

// newFilesystemStorage creates a Blobsstore compatible storage for the filesystem.
func newFilesystemStorage(ctx context.Context) (Blobstore, error) {
	return &FilesystemStorage{}, nil
}

// CreateObject creates a new object on the filesystem or overwrites an existing one.
func (s *FilesystemStorage) CreateObject(ctx context.Context, folder, filename string, contents *bytes.Buffer) error {
	pth := filepath.Join(folder, filename)
	if err := os.MkdirAll(filepath.Dir(pth), 0755); err != nil {
		return fmt.Errorf("failed to create object: %w", err)
	}

	f, err := os.Create(pth)
	if err != nil {
		return fmt.Errorf("failed to create object: %w", err)
	}

	_, err = contents.WriteTo(f)
	if err != nil {
		return fmt.Errorf("failed to write record to file: %w", err)
	}

	return nil
}

func (s *FilesystemStorage) CreateFObject(ctx context.Context, folder, filename, fromFilename string) error {
	return ioext.CopyFile(fromFilename, filename, true)
}

// GetObject download object from storage system.
func (s *FilesystemStorage) GetObject(ctx context.Context, folder, filename string) (*bytes.Buffer, error) {
	pth := filepath.Join(folder, filename)

	b, err := ioutil.ReadFile(pth)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return bytes.NewBuffer(b), nil
}

// DeleteObject deletes an object from the filesystem. It returns nil if the object was deleted or if the object no longer exists.
func (s *FilesystemStorage) DeleteObject(ctx context.Context, folder, filename string) error {
	pth := filepath.Join(folder, filename)
	if err := os.Remove(pth); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}
