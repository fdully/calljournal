package storage

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilesystemStorage_CreateObject(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = os.RemoveAll(tmp) })

	cases := []struct {
		name     string
		folder   string
		filepath string
		contents *bytes.Buffer
		err      bool
	}{
		{
			name:     "default",
			folder:   tmp,
			filepath: "myfile",
			contents: bytes.NewBuffer([]byte{}),
		},
		{
			name:     "bad_path",
			folder:   "/badpath/doesn't/exist",
			filepath: "myfile",
			contents: new(bytes.Buffer),
			err:      true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			storage, err := newFilesystemStorage(ctx)
			if err != nil {
				t.Fatal(err)
			}

			err = storage.CreateObject(ctx, tc.folder, tc.filepath, tc.contents)
			if (err != nil) != tc.err {
				t.Fatal(err)
			}

			if !tc.err {
				contents, err := storage.GetObject(ctx, tc.folder, tc.filepath)
				if err != nil {
					t.Fatal(err)
				}

				require.Equal(t, tc.contents, contents)
			}
		})
	}
}

func TestFilesystemStorage_DeleteObject(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		folder   string
		filepath string
	}{
		{
			name:     "default",
			folder:   filepath.Dir(f.Name()),
			filepath: filepath.Base(f.Name()),
		},
		{
			name:     "not_exist",
			folder:   filepath.Dir(f.Name()),
			filepath: "not-exist",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			storage, err := newFilesystemStorage(ctx)
			if err != nil {
				t.Fatal(err)
			}

			if err = storage.DeleteObject(ctx, tc.folder, tc.filepath); err != nil {
				t.Fatal(err)
			}
		})
	}
}
