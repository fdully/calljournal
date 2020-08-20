package callfiles_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/calluploader"
	"github.com/fdully/calljournal/internal/calluploader/callfiles"
	"github.com/fdully/calljournal/internal/util"
	"github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/require"
)

func TestCallFiles(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	filesChan := make(chan os.FileInfo)

	const (
		baseCallDir = "../../../testdata"
	)

	cf := callfiles.NewCallFiles(baseCallDir, 100*time.Millisecond)

	go func() {
		<-time.After(300 * time.Millisecond)
		cancel()
	}()

	go func() {
		for v := range filesChan {
			require.NotEqual(t, v, "")
		}
	}()

	go func() {
		defer cancel()

		err := cf.ReadBaseCallsFromDir(ctx, filesChan)
		require.NoError(t, err)
	}()

	<-ctx.Done()
}

func TestParseCalls(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	filesChan := make(chan os.FileInfo)

	lookuper := envconfig.MapLookuper(map[string]string{
		"CJ_BASECALL_DIR":    "../../../testdata",
		"CJ_READ_DIR_PERIOD": "100ms",
		"CJ_STORAGE_ADDR":    "localhost:8080",
	})

	var config calluploader.Config
	err := envconfig.ProcessWith(ctx, &config, lookuper)
	require.NoError(t, err)

	cf := callfiles.NewCallFiles(config.BaseCallDir, config.ReadDirPeriod)

	go func() {
		<-time.After(300 * time.Millisecond)
		cancel()
	}()

	go func() {
		defer cancel()

		err := cf.ReadBaseCallsFromDir(ctx, filesChan)
		require.NoError(t, err)
	}()

	for v := range filesChan {
		id, err := util.GetUUIDFromString(v.Name())
		require.NoError(t, err)

		b, err := cf.OpenFile(v.Name())
		require.NoError(t, err)

		bc, err := util.ParseCall(ctx, b)
		require.NoError(t, err)
		require.Equal(t, id, bc.UUID)
	}

	<-ctx.Done()
}
