package cdrutil

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	cjmodel "github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/google/uuid"
)

func CDRPathInfoFromCDR(ctx context.Context, cdr model.CDR) (model.CallPath, error) {
	logger := logging.FromContext(ctx)

	var (
		callPath model.CallPath
		dirc     = "inc"
		name     = cdr.Variables.UUID + ".xml"
	)

	id, err := uuid.Parse(cdr.Variables.UUID)
	if err != nil {
		return callPath, fmt.Errorf("failed to parse cdr uuid %s: %w", cdr.Variables.UUID, err)
	}

	stti := time.Unix(int64(cdr.Variables.StartEpoch), 0)

	if cdr.Variables.CJCallDirection != "" {
		dirc = cdr.Variables.CJCallDirection
	} else {
		logger.Debugf("cdr %s has no dirc, setting default", cdr.Variables.UUID)
	}

	return model.CallPath{
		UUID: id,
		DIRC: dirc,
		YEAR: strconv.Itoa(stti.UTC().Year()),
		MONT: fmt.Sprintf("%02d", int(stti.UTC().Month())),
		DAYX: fmt.Sprintf("%02d", stti.UTC().Day()),
		NAME: name,
	}, nil
}

func RecordPathInfoFromCDR(ctx context.Context, cdr model.CDR) (model.CallPath, error) {
	logger := logging.FromContext(ctx)

	var (
		callPathInfo model.CallPath
		dirc         = "inc"
	)

	id, err := uuid.Parse(cdr.Variables.UUID)
	if err != nil {
		return callPathInfo, fmt.Errorf("failed to parse cdr uuid %s: %w", cdr.Variables.UUID, err)
	}

	stti := time.Unix(int64(cdr.Variables.StartEpoch), 0)

	if cdr.Variables.CJCallDirection != "" {
		dirc = cdr.Variables.CJCallDirection
	} else {
		logger.Debugf("cdr %s has no dirc, setting default", cdr.Variables.UUID)
	}

	return model.CallPath{
		UUID: id,
		DIRC: dirc,
		YEAR: strconv.Itoa(stti.UTC().Year()),
		MONT: fmt.Sprintf("%02d", int(stti.UTC().Month())),
		DAYX: fmt.Sprintf("%02d", stti.UTC().Day()),
		NAME: string(cdr.Variables.CJRecordName),
	}, nil
}

// Creates http record path: /inc/2020/04/24/record_file_name.wav.
func CreateHTTPCallPath(r model.CallPath) string {
	return r.DIRC + "/" + r.YEAR + "/" + r.MONT + "/" + r.DAYX + "/" + r.NAME
}

func CreateFileCallPath(r model.CallPath) string {
	return filepath.Join(r.DIRC, r.YEAR, r.MONT, r.DAYX, r.NAME)
}

func ProtobufCallPathToCallPath(p *pb.CallPath) (model.CallPath, error) {
	var info model.CallPath

	id, err := uuid.Parse(p.Uuid)
	if err != nil {
		return info, fmt.Errorf("failed to parse uuid %s: %w", p.Uuid, err)
	}

	return model.CallPath{
		UUID: id,
		DIRC: p.Dirc,
		YEAR: p.Year,
		MONT: p.Mont,
		DAYX: p.Dayx,
		NAME: p.Name,
	}, nil
}

func CallPathToProtobufCallPath(p model.CallPath) *pb.CallPath {
	return &pb.CallPath{
		Uuid: p.UUID.String(),
		Dirc: p.DIRC,
		Year: p.YEAR,
		Mont: p.MONT,
		Dayx: p.DAYX,
		Name: p.NAME,
	}
}

func BasecallToCallPath(bc *cjmodel.BaseCall) model.CallPath {
	return model.CallPath{
		UUID: bc.UUID,
		DIRC: bc.Direction,
		YEAR: strconv.Itoa(bc.StartStamp.UTC().Year()),
		MONT: fmt.Sprintf("%02d", int(bc.StartStamp.UTC().Month())),
		DAYX: fmt.Sprintf("%02d", bc.StartStamp.UTC().Day()),
		NAME: bc.RecordName,
	}
}

func BasecallToCDRPath(bc *cjmodel.BaseCall) model.CallPath {
	return model.CallPath{
		UUID: bc.UUID,
		DIRC: bc.Direction,
		YEAR: strconv.Itoa(bc.StartStamp.UTC().Year()),
		MONT: fmt.Sprintf("%02d", int(bc.StartStamp.UTC().Month())),
		DAYX: fmt.Sprintf("%02d", bc.StartStamp.UTC().Day()),
		NAME: bc.UUID.String() + ".xml",
	}
}
