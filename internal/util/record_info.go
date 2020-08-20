package util

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/google/uuid"
)

var ErrCantBeNil = errors.New("can't be nil")

func CreateRecordInfo(bc *model.BaseCall) model.RecordInfo {
	return model.RecordInfo{
		UUID:  bc.UUID,
		Name:  bc.RECL,
		Dirc:  bc.DIRC,
		Year:  strconv.Itoa(bc.STTI.Year()),
		Month: fmt.Sprintf("%02d", int(bc.STTI.Month())),
		Day:   fmt.Sprintf("%02d", bc.STTI.Day()),
	}
}

// Creates record path: /inc/2020/04/24/record_file_name.wav.
func CreateRecordPath(r model.RecordInfo) string {
	return r.Dirc + "/" + r.Year + "/" + r.Month + "/" + r.Day + "/" + r.Name
}

func ProtobufRecordInfoToRecordInfo(p *pb.RecordInfo) (model.RecordInfo, error) {
	var recordInfo model.RecordInfo

	if p == nil {
		return recordInfo, ErrCantBeNil
	}

	id, err := uuid.Parse(p.Uuid)
	if err != nil {
		return recordInfo, fmt.Errorf("failed to parse uuid %s: %w", p.Uuid, err)
	}

	return model.RecordInfo{
		UUID:  id,
		Name:  p.Name,
		Dirc:  p.Dirc,
		Year:  p.Year,
		Month: p.Month,
		Day:   p.Day,
	}, nil
}

func RecordInfoToProtobufRecordInfo(p model.RecordInfo) *pb.RecordInfo {
	return &pb.RecordInfo{
		Uuid:  p.UUID.String(),
		Name:  p.Name,
		Dirc:  p.Dirc,
		Year:  p.Year,
		Month: p.Month,
		Day:   p.Day,
	}
}
