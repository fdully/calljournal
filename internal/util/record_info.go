package util

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/google/uuid"
)

func CreateRecordInfo(bc *model.BaseCall, addr string) model.RecordInfo {
	return model.RecordInfo{
		UUID: bc.UUID,
		ADDR: addr,
		DIRC: bc.DIRC,
		YEAR: strconv.Itoa(bc.STTI.Year()),
		MONT: fmt.Sprintf("%02d", int(bc.STTI.Month())),
		RDAY: fmt.Sprintf("%02d", bc.STTI.Day()),
		RNAM: bc.RNAM,
	}
}

// Creates http record path: /inc/2020/04/24/record_file_name.wav.
func CreateHTTPRecordPath(r model.RecordInfo) string {
	return r.DIRC + "/" + r.YEAR + "/" + r.MONT + "/" + r.RDAY + "/" + r.RNAM
}

func CreateFileRecordPath(r model.RecordInfo) string {
	return filepath.Join(r.DIRC, r.YEAR, r.MONT, r.RDAY, r.RNAM)
}

func ProtobufRecordInfoToRecordInfo(p *pb.RecordInfo) (model.RecordInfo, error) {
	var recordInfo model.RecordInfo

	id, err := uuid.Parse(p.Uuid)
	if err != nil {
		return recordInfo, fmt.Errorf("failed to parse uuid %s: %w", p.Uuid, err)
	}

	return model.RecordInfo{
		UUID: id,
		ADDR: p.Addr,
		DIRC: p.Dirc,
		YEAR: p.Year,
		MONT: p.Mont,
		RDAY: p.Rday,
		RNAM: p.Rnam,
	}, nil
}

func RecordInfoToProtobufRecordInfo(p model.RecordInfo) *pb.RecordInfo {
	return &pb.RecordInfo{
		Uuid: p.UUID.String(),
		Addr: p.ADDR,
		Dirc: p.DIRC,
		Year: p.YEAR,
		Mont: p.MONT,
		Rday: p.RDAY,
		Rnam: p.RNAM,
	}
}
