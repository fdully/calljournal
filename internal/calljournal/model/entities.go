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
	RNAM string    `json:"rnam"`
	RTAG string    `json:"rtag"`
	EPOS int64     `json:"epos"`
	EPOA int64     `json:"epoa"`
	EPOE int64     `json:"epoe"`
	TAG1 string    `json:"tag1"`
	TAG2 string    `json:"tag2"`
	TAG3 string    `json:"tag3"`
	WBYE string    `json:"wbye"`
	HANG string    `json:"hang"`
	CODE string    `json:"code"`
}

type RecordInfo struct {
	UUID uuid.UUID `json:"uuid"`
	ADDR string    `json:"addr"`
	DIRC string    `json:"dirc"`
	YEAR string    `json:"year"`
	MONT string    `json:"mont"`
	RDAY string    `json:"rday"`
	RNAM string    `json:"rnam"`
}
