package logging

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LoggerName  string `env:"CJ_LOGGER_NAME, required"`
	Severity    string `env:"CJ_LOGGER_LEVEL, default=debug"`
	LogFileName string `env:"CJ_LOGGER_FILE"`
}

type loggerKey struct{}

var fallbackLogger *zap.SugaredLogger

func CreateLogger(c *Config) error {
	var l zapcore.Level
	switch c.Severity {
	case "debug":
		l = zapcore.DebugLevel
	case "info":
		l = zapcore.InfoLevel
	case "warn":
		l = zapcore.WarnLevel
	case "error":
		l = zapcore.ErrorLevel
	case "fatal":
		l = zapcore.FatalLevel
	default:
		l = zapcore.DebugLevel
	}

	z := zap.NewProductionConfig()
	z.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	z.EncoderConfig.MessageKey = "message"
	z.EncoderConfig.LevelKey = "severity"
	z.Level = zap.NewAtomicLevelAt(l)

	if c.LogFileName != "" {
		z.OutputPaths = append(z.OutputPaths, c.LogFileName)
	}

	logger, err := z.Build()
	if err != nil {
		return fmt.Errorf("failed to build zapp logger: %w", err)
	}
	fallbackLogger = logger.Named(c.LoggerName).Sugar()

	return nil
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}
	return fallbackLogger
}
