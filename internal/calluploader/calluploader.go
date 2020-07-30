package calluploader

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/google/uuid"
)

func ParseCall(ctx context.Context, call []byte) (*model.BaseCall, error) {
	bc := struct {
		UUID uuid.UUID
		CLID string
		CLNA string
		DEST string
		DIRC string
		STTI string
		DURS int32
		BILS int32
		RECD bool
		RECS int32
		RECL string
		RTAG string
		EPOS int64
		EPOA int64
		EPOE int64
		WBYE string
		HANG string
		CODE string
	}{}
	err := json.Unmarshal(call, &bc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal base call %w", err)
	}

	stti, err := time.Parse("2006-01-02 15:04:05", bc.STTI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stti time %w", err)
	}

	return &model.BaseCall{
		UUID: bc.UUID,
		CLID: bc.CLID,
		CLNA: bc.CLNA,
		DEST: bc.DEST,
		DIRC: bc.DIRC,
		STTI: stti,
		DURS: bc.DURS,
		BILS: bc.BILS,
		RECD: bc.RECD,
		RECS: bc.RECS,
		RECL: bc.RECL,
		RTAG: bc.RTAG,
		EPOS: bc.EPOS,
		EPOA: bc.EPOA,
		EPOE: bc.EPOE,
		WBYE: bc.WBYE,
		HANG: bc.HANG,
		CODE: bc.CODE,
	}, nil
}

func CreateAudioRecord(uuid uuid.UUID, name string, startTime time.Time, data []byte) *model.AudioRecord {
	return &model.AudioRecord{
		UUID:       uuid,
		StartTime:  startTime,
		ObjectName: name,
		Data:       data,
	}
}

func LookForFiles(dir, ext string, old time.Duration) ([]string, error) {
	var result []string

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if !f.Mode().IsRegular() {
			return nil
		}
		if filepath.Ext(path) == ext && f.ModTime().Add(old).Before(time.Now()) {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed look for files in %s with extension %s, - %v", dir, ext, err)
	}

	return result, nil
}
