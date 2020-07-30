package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseCall struct {
	UUID uuid.UUID `json:"uuid"`
	CLID string    `json:"clid"`
	CLNA string    `json:"clna"`
	DEST string    `json:"dest"`
	DIRC string    `json:"dirc"`
	STTI time.Time `json:"stti"`
	DURS int32     `json:"durs"`
	BILS int32     `json:"bils"`
	RECD bool      `json:"recd"`
	RECS int32     `json:"recs"`
	RECL string    `json:"recl"`
	RTAG string    `json:"rtag"`
	EPOS int64     `json:"epos"`
	EPOA int64     `json:"epoa"`
	EPOE int64     `json:"epoe"`
	WBYE string    `json:"wbye"`
	HANG string    `json:"hang"`
	CODE string    `json:"code"`
}

type AudioRecord struct {
	UUID       uuid.UUID
	StartTime  time.Time
	ObjectName string
	Data       []byte
}
