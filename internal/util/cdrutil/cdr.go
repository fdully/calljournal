package cdrutil

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	cjmodel "github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/google/uuid"
)

func ParseCDR(cdr []byte) (model.CDR, error) {
	var result model.CDR
	if err := xml.Unmarshal(cdr, &result); err != nil {
		return result, err
	}

	return result, nil
}

func CDRToBaseCall(cdr model.CDR) (cjmodel.BaseCall, error) {
	var bc cjmodel.BaseCall

	const (
		profileIndex = "1"
		location     = "Europe/Moscow"
	)

	id, err := uuid.Parse(cdr.Variables.UUID)
	if err != nil {
		return bc, fmt.Errorf("failed to convert cdr to basecall, uuid %s: %w", cdr.Variables.UUID, err)
	}

	loc, _ := time.LoadLocation(location)

	t, err := time.ParseInLocation("2006-01-02 15:04:05", string(cdr.Variables.StartStamp), loc)
	if err != nil {
		return bc, fmt.Errorf("failed to convert cdr to basecall, uuid %s: %w", cdr.Variables.UUID, err)
	}

	for _, v := range cdr.Callflow {
		if v.ProfileIndex == profileIndex {
			bc.Username = string(v.CallerProfile.Username)
			bc.CallerIDName = string(v.CallerProfile.CallerIDName)
			bc.CallerIDNumber = string(v.CallerProfile.CallerIDNumber)
			bc.DestinationNumber = string(v.CallerProfile.DestinationNumber)
		}
	}

	bc.UUID = id
	bc.Direction = cdr.Variables.CJCallDirection
	bc.StartStamp = t
	bc.Duration = int64(cdr.Variables.Duration)
	bc.Billsec = int64(cdr.Variables.Billsec)
	bc.RecordSeconds = int64(cdr.Variables.RecordSeconds)
	bc.RecordName = string(cdr.Variables.CJRecordName)
	bc.StartEpoch = int64(cdr.Variables.StartEpoch)
	bc.AnswerEpoch = int64(cdr.Variables.AnswerEpoch)
	bc.EndEpoch = int64(cdr.Variables.EndEpoch)
	bc.SIPHangupDisposition = cdr.Variables.SIPHangupDisposition
	bc.HangupCause = cdr.Variables.HangupCause
	bc.SIPTermStatus = cdr.Variables.SIPTermStatus

	return bc, nil
}

func WhoDisconnect(call *cjmodel.BaseCall) string {
	const (
		worker = "сотрудник"
		client = "клиент"
	)

	if call.Direction == "inc" {
		if strings.Contains(call.SIPHangupDisposition, "send") {
			return worker
		}

		return client
	}

	if strings.Contains(call.SIPHangupDisposition, "send") {
		return client
	}

	return worker
}
