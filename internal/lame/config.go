package lame

import (
	"context"

	"github.com/fdully/calljournal/internal/logging"
	"github.com/sethvargo/go-envconfig"
)

type lameConfig struct {
	TempDir string `env:"CJ_TEMP_DIR, default=."`
}

var config lameConfig

func Init(ctx context.Context) {
	logger := logging.FromContext(ctx)

	if err := envconfig.Process(ctx, &config); err != nil {
		logger.Fatalf("failed to process audio config: %v", err)
	}

	if err := pingLame(); err != nil {
		logger.Fatal(err)
	}
}
