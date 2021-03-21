package basecall

import (
	"fmt"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/util"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

func ToProtobufBaseCall(bc *model.BaseCall) (*pb.BaseCall, error) {
	var pbBC pb.BaseCall

	if bc == nil {
		return nil, fmt.Errorf("basecall: %w", util.ErrIsNil)
	}

	stti, err := ptypes.TimestampProto(bc.StartStamp)
	if err != nil {
		return nil, fmt.Errorf("failed to make protobuf timestamp: %w", err)
	}

	pbBC = pb.BaseCall{
		Uuid:                 bc.UUID.String(),
		Username:             bc.Username,
		CallerIdNumber:       bc.CallerIDNumber,
		CallerIdName:         bc.CallerIDName,
		DestinationNumber:    bc.DestinationNumber,
		Direction:            bc.Direction,
		StartStamp:           stti,
		Duration:             bc.Duration,
		Billsec:              bc.Billsec,
		RecordSeconds:        bc.RecordSeconds,
		RecordName:           bc.RecordName,
		StartEpoch:           bc.StartEpoch,
		AnswerEpoch:          bc.AnswerEpoch,
		EndEpoch:             bc.EndEpoch,
		SipHangupDisposition: bc.SIPHangupDisposition,
		HangupCause:          bc.HangupCause,
		SipTermStatus:        bc.SIPTermStatus,
	}

	return &pbBC, nil
}

func ProtobufBaseCallToBaseCall(pbBC *pb.BaseCall) (*model.BaseCall, error) {
	var bc model.BaseCall

	if pbBC == nil {
		return nil, fmt.Errorf("protobuf basecall: %w", util.ErrIsNil)
	}

	stti, err := ptypes.Timestamp(pbBC.StartStamp)
	if err != nil {
		return nil, fmt.Errorf("failed to make timestamp from protobuf times: %w", err)
	}

	id, err := uuid.Parse(pbBC.Uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uuid from %s: %w", pbBC.Uuid, err)
	}

	bc = model.BaseCall{
		UUID:                 id,
		Username:             pbBC.Username,
		CallerIDNumber:       pbBC.CallerIdNumber,
		CallerIDName:         pbBC.CallerIdName,
		DestinationNumber:    pbBC.DestinationNumber,
		Direction:            pbBC.Direction,
		StartStamp:           stti,
		Duration:             pbBC.Duration,
		Billsec:              pbBC.Billsec,
		RecordSeconds:        pbBC.RecordSeconds,
		RecordName:           pbBC.RecordName,
		StartEpoch:           pbBC.StartEpoch,
		AnswerEpoch:          pbBC.AnswerEpoch,
		EndEpoch:             pbBC.EndEpoch,
		SIPHangupDisposition: pbBC.SipHangupDisposition,
		HangupCause:          pbBC.HangupCause,
		SIPTermStatus:        pbBC.SipTermStatus,
	}

	return &bc, nil
}
