package logging

import (
	"context"
	"testing"

	"github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Run("init logging", func(t *testing.T) {
		lookuper := envconfig.MapLookuper(map[string]string{
			"CJ_LOGGER_NAME":  "test logger",
			"CJ_LOGGER_LEVEL": "info",
		})
		var config Config
		err := envconfig.ProcessWith(context.Background(), &config, lookuper)
		require.NoError(t, err)

		err = CreateLogger(&config)
		require.NoError(t, err)

		ctx := context.Background()
		logger := FromContext(ctx)

		require.Equal(t, fallbackLogger, logger)

		ctx = WithLogger(ctx, logger)
		require.Equal(t, logger, FromContext(ctx))
	})
}
