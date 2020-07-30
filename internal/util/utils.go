package util

import (
	"regexp"

	"github.com/google/uuid"
)

var uuidRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

func GetUUIDFromString(str string) (uuid.UUID, error) {
	s := uuidRegex.FindString(str)
	return uuid.Parse(s)
}
