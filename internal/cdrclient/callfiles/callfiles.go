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
	CDRFilesExt = ".xml"
)

type FilesInterface interface {
	ReadCDRFiles(ctx context.Context, fileChan chan string) error
	OpenFile(fname string) ([]byte, error)
	DeleteFile(ctx context.Context, fname string) error
	AgainLater(ctx context.Context, fname string, err error)
}

// callFiles is periodically reading cdr xml files from file system and send to workers.
// Keeps in map "work" currently processed files.
// When cdr and record is successfully uploaded to grpc server deletes cdr file and record.
// If fails then process same files on next iteration.
type callFiles struct {
	dir        string
	readPeriod time.Duration
	mu         sync.Mutex
	work       map[string]struct{}
}

func NewCallFiles(baseCallDir string, readDirPeriod time.Duration) FilesInterface {
	return &callFiles{
		dir:        baseCallDir,
		readPeriod: readDirPeriod,
		mu:         sync.Mutex{},
		work:       make(map[string]struct{}),
	}
}

// Periodically reads cdr files and send them to chan for further processing.
func (c *callFiles) ReadCDRFiles(ctx context.Context, fileChan chan string) error {
	logger := logging.FromContext(ctx)

	defer close(fileChan)

	ticker := time.NewTicker(c.readPeriod)

	for {
		select {
		case <-ticker.C:
			logger.Infof("reading cdr directory: %s", c.dir)

			cdrFiles, err := c.getCDRFiles()
			if err != nil {
				return err
			}

			for _, v := range cdrFiles {
				// sending file further to worker
				logger.Infof("passing cdr %s", v)
				fileChan <- v
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

func (c *callFiles) DeleteFile(ctx context.Context, fname string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// delete file from work
	delete(c.work, fname)

	if _, err := os.Stat(filepath.Join(c.dir, fname)); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(filepath.Join(c.dir, fname))
}

func (c *callFiles) AgainLater(ctx context.Context, fname string, err error) {
	if err != nil {
		logger := logging.FromContext(ctx)
		logger.DPanicf("failed to process cdr file %s: %v", fname, err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.work, fname)
}

func (c *callFiles) getCDRFiles() ([]string, error) {
	files, err := ioutil.ReadDir(c.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %s: %w", c.dir, err)
	}

	c.mu.Lock()

	defer c.mu.Unlock()

	cdrFiles := make([]string, 0, 200)

	for _, v := range files {
		if !v.Mode().IsRegular() {
			continue
		}

		// Check if file has correct extension
		if filepath.Ext(v.Name()) != CDRFilesExt {
			continue
		}

		_, ok := c.work[v.Name()]
		if ok {
			continue
		}

		c.work[v.Name()] = struct{}{}

		cdrFiles = append(cdrFiles, v.Name())
	}

	return cdrFiles, nil
}
