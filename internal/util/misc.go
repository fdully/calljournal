package util

import (
	"context"
	"regexp"
	"strings"

	"github.com/fdully/calljournal/internal/logging"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var uuidRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

func GetUUIDFromString(str string) (uuid.UUID, error) {
	s := uuidRegex.FindString(str)
	return uuid.Parse(s)
}

func LogError(err error) error {
	logger := logging.FromContext(context.Background())

	if err != nil {
		logger.Error(err)
	}

	return err
}

func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return LogError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return LogError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func ChangeWavExtToMp3(wavName string) string {
	return strings.TrimSuffix(wavName, ".wav") + ".mp3"
}
