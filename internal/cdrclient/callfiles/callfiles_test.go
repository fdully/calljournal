package callfiles_test

import (
	"context"
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/cdrclient"
	"github.com/fdully/calljournal/internal/cdrclient/callfiles"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/util"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestCallFiles(t *testing.T) {
	glbCtx := context.Background()

	logger := zap.NewNop()
	glbCtx = logging.WithLogger(glbCtx, logger.Sugar())

	lookuper := envconfig.MapLookuper(map[string]string{
		"CJ_BASECALL_DIR":    "../../../testdata",
		"CJ_READ_DIR_PERIOD": "100ms",
	})

	var config cdrclient.Config
	err := envconfig.ProcessWith(glbCtx, &config, lookuper)
	require.NoError(t, err)

	cf := callfiles.NewCallFiles(config.BaseCallDir, config.ReadDirPeriod)

	t.Run("plain read", func(t *testing.T) {
		filesChan := make(chan string)
		ctx, cancel := context.WithCancel(glbCtx)

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

			err := cf.ReadCDRFiles(ctx, filesChan)
			require.NoError(t, err)
		}()

		<-ctx.Done()
	})

	t.Run("parse calls", func(t *testing.T) {
		filesChan := make(chan string)
		ctx, cancel := context.WithCancel(glbCtx)

		go func() {
			<-time.After(300 * time.Millisecond)
			cancel()
		}()

		go func() {
			defer cancel()

			err := cf.ReadCDRFiles(ctx, filesChan)
			require.NoError(t, err)
		}()

		for v := range filesChan {
			id, err := util.GetUUIDFromString(v)
			require.NoError(t, err)

			b, err := cf.OpenFile(v)
			require.NoError(t, err)

			bc, err := cdrutil.ParseCDR(b)
			require.NoError(t, err)
			require.Equal(t, id, bc.Variables.UUID)
		}

		<-ctx.Done()
	})
}
