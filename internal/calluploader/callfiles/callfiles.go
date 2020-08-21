package callfiles

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fdully/calljournal/internal/logging"
)

const (
	callFilesExt = ".json"
)

type FilesInterface interface {
	ReadBaseCallsFromDir(ctx context.Context, fileChan chan os.FileInfo) error
	OpenFile(fname string) ([]byte, error)
	DeleteFile(fname string) error
	DoItAgainLater(ctx context.Context, fname string, err error)
}

// callFiles is periodically reading basecalls json files from file system and send to workers.
// Keeps in map "work" currently processed files.
// When basecall and record is successfully uploaded to grpc server deletes basecall json file and record.
// If fails then process same files on next iteration.
type callFiles struct {
	mu         sync.Mutex
	dir        string
	work       map[string]struct{}
	readPeriod time.Duration
}

func NewCallFiles(baseCallDir string, readDirPeriod time.Duration) FilesInterface {
	return &callFiles{
		mu:         sync.Mutex{},
		dir:        baseCallDir,
		work:       make(map[string]struct{}),
		readPeriod: readDirPeriod,
	}
}

// Periodically reads json basecall files and send them to chan for further processing.
func (c *callFiles) ReadBaseCallsFromDir(ctx context.Context, fileChan chan os.FileInfo) error {
	defer close(fileChan)

	ticker := time.NewTicker(c.readPeriod)

	for {
		select {
		case <-ticker.C:
			files, err := ioutil.ReadDir(c.dir)
			if err != nil {
				return fmt.Errorf("failed to read dir %s: %w", c.dir, err)
			}

			for _, v := range files {
				if !v.Mode().IsRegular() {
					continue
				}
				// Check if file has correct extension
				if filepath.Ext(v.Name()) != callFilesExt {
					continue
				}

				c.mu.Lock()

				_, ok := c.work[v.Name()]
				if !ok {
					c.work[v.Name()] = struct{}{}
				}

				c.mu.Unlock()

				// sending file further to worker
				if !ok {
					fileChan <- v
				}
			}
		case <-ctx.Done():
			ticker.Stop()

			return nil
		}
	}
}

func (c *callFiles) OpenFile(fname string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return ioutil.ReadFile(filepath.Join(c.dir, fname))
}

func (c *callFiles) DeleteFile(fname string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// delete file from work
	delete(c.work, fname)

	if err := os.Remove(filepath.Join(c.dir, fname)); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete %s: %w", fname, err)
	}

	return nil
}

func (c *callFiles) DoItAgainLater(ctx context.Context, fname string, err error) {
	if err != nil {
		logger := logging.FromContext(ctx)
		logger.DPanicf("failed to process basecall file %s: %v", fname, err)
	}

	c.removeFileFromWork(fname)
}

func (c *callFiles) removeFileFromWork(fname string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.work, fname)
}
