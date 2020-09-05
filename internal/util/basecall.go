package util

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

func BaseCallToProtobufBaseCall(bc *model.BaseCall) (*pb.BaseCall, error) {
	var pbBC pb.BaseCall

	stti, err := ptypes.TimestampProto(bc.STTI)
	if err != nil {
		return nil, fmt.Errorf("failed to make protobuf timestamp from timestamp: %w", err)
	}

	pbBC = pb.BaseCall{
		Uuid: bc.UUID.String(),
		Clid: bc.CLID,
		Clna: bc.CLNA,
		Dest: bc.DEST,
		Dirc: bc.DIRC,
		Stti: stti,
		Durs: bc.DURS,
		Bils: bc.BILS,
		Recd: bc.RECD,
		Recs: bc.RECS,
		Rnam: bc.RNAM,
		Rtag: bc.RTAG,
		Epos: bc.EPOS,
		Epoa: bc.EPOA,
		Epoe: bc.EPOE,
		Tag1: bc.TAG1,
		Tag2: bc.TAG2,
		Tag3: bc.TAG3,
		Wbye: bc.WBYE,
		Hang: bc.HANG,
		Code: bc.CODE,
	}

	return &pbBC, nil
}

func ProtobufBaseCallToBaseCall(pbBC *pb.BaseCall) (*model.BaseCall, error) {
	var bc model.BaseCall

	stti, err := ptypes.Timestamp(pbBC.Stti)
	if err != nil {
		return nil, fmt.Errorf("failed to make timestamp from protobuf times: %w", err)
	}

	id, err := uuid.Parse(pbBC.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uuid from %s: %w", pbBC.Uuid, err)
	}

	bc = model.BaseCall{
		UUID: id,
		CLID: pbBC.Clid,
		CLNA: pbBC.Clna,
		DEST: pbBC.Dest,
		DIRC: pbBC.Dirc,
		STTI: stti,
		DURS: pbBC.Durs,
		BILS: pbBC.Bils,
		RECD: pbBC.Recd,
		RECS: pbBC.Recs,
		RNAM: pbBC.Rnam,
		RTAG: pbBC.Rtag,
		EPOS: pbBC.Epos,
		EPOA: pbBC.Epoa,
		EPOE: pbBC.Epoe,
		TAG1: pbBC.Tag1,
		TAG2: pbBC.Tag2,
		TAG3: pbBC.Tag3,
		WBYE: pbBC.Wbye,
		HANG: pbBC.Hang,
		CODE: pbBC.Code,
	}

	return &bc, nil
}

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
		RNAM string
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
		RNAM: bc.RNAM,
		RTAG: bc.RTAG,
		EPOS: bc.EPOS,
		EPOA: bc.EPOA,
		EPOE: bc.EPOE,
		WBYE: bc.WBYE,
		HANG: bc.HANG,
		CODE: bc.CODE,
	}, nil
}
