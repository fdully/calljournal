package logging

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type loggerConfig struct {
	LoggerName  string `env:"CJ_LOGGER_NAME, required"`
	Severity    string `env:"CJ_LOGGER_LEVEL, default=debug"`
	LogFileName string `env:"CJ_LOGGER_FILE"`
	DevelopMode bool   `env:"CJ_LOGGER_DEVELOP_MODE, default=true"`
}

func Init(ctx context.Context) {
	var config loggerConfig
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatalf("failed to process logging config: %v\n", err)
	}

	err := createLogger(&config)
	if err != nil {
		log.Fatalf("failed to create logger: %v\n", err)
	}
}
