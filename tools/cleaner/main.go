package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sethvargo/go-signalcontext"
)

var (
	basecallFolderFlag = flag.String("folder", "", "basecall folder to clean")
	daysToCleanFlag    = flag.Int("days", 7, "older then these days to clean")

	ErrFolderFlagRequired = errors.New("--folder is required and cannot be empty")
)

func main() {
	ctx, done := signalcontext.OnInterrupt()

	err := realMain(ctx)

	done()

	if err != nil {
		log.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	flag.Parse()

	if *basecallFolderFlag == "" {
		return ErrFolderFlagRequired
	}

	bc := *basecallFolderFlag

	err := cleanBasecallFolder(ctx, bc)
	if err != nil {
		return err
	}

	return nil
}

func cleanBasecallFolder(ctx context.Context, bc string) error {
	objects := make([]string, 0, 500)

	err := filepath.Walk(bc, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			objects = append(objects, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	for _, v := range objects {
		if isDirOld(v, *daysToCleanFlag) {
			if err := os.RemoveAll(v); err != nil {
				log.Fatal(err)
			}
		}

		if isDirEmpty(v) {
			if err := os.RemoveAll(v); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

func isDirEmpty(dir string) bool {
	elems, err := ioutil.ReadDir(dir)
	if err != nil {
		return false
	}

	if len(elems) < 1 {
		return true
	}

	return false
}

func isDirOld(dir string, days int) bool {
	const (
		minDirLength = 10
		minElems     = 3
	)

	if len(dir) < minDirLength {
		return false
	}

	elems := strings.Split(dir[len(dir)-minDirLength:], "/")
	if len(elems) < minElems {
		return false
	}

	date, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", elems[0], elems[1], elems[2]))
	if err != nil {
		return false
	}

	if date.AddDate(0, 0, days).Before(time.Now()) {
		return true
	}

	return false
}
