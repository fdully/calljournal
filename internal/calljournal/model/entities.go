package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseCall struct {
	UUID                 uuid.UUID
	Username             string
	CallerIDName         string
	CallerIDNumber       string
	DestinationNumber    string
	Direction            string
	StartStamp           time.Time
	Duration             int64
	Billsec              int64
	RecordSeconds        int64
	RecordName           string
	StartEpoch           int64
	AnswerEpoch          int64
	EndEpoch             int64
	SIPHangupDisposition string
	HangupCause          string
	SIPTermStatus        string
}

type Call struct {
	BaseCall    *BaseCall
	RecordName  string
	Disconnect  string
	ConnectTime int64
}
