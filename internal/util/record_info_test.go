package util_test

import (
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRecordInfo(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	stti, err := time.Parse(timeLayout, "2020-07-18 19:59:26")
	require.NoError(t, err)

	id, err := uuid.Parse("03fb24ea-3a81-4469-8522-7753d643dcfe")
	require.NoError(t, err)

	storageAddr := "localhost:8080"

	bc := model.BaseCall{
		UUID: id,
		CLID: "79002800421",
		CLNA: "79002800421",
		DEST: "9000",
		DIRC: "inc",
		STTI: stti,
		DURS: 188,
		BILS: 178,
		RECD: true,
		RECS: 178,
		RNAM: "inc_2020_07_18_03fb24ea-3a81-4469-8522-7753d643dcfe.wav",
		RTAG: "public",
		EPOS: 1595102366,
		EPOA: 1595102367,
		EPOE: 1595102384,
		TAG1: "",
		TAG2: "",
		TAG3: "",
		WBYE: "recv_bye",
		HANG: "NORMAL_CLEARING",
		CODE: "200",
	}

	recordInfo := model.RecordInfo{
		UUID: id,
		ADDR: storageAddr,
		RNAM: "inc_2020_07_18_03fb24ea-3a81-4469-8522-7753d643dcfe.wav",
		DIRC: "inc",
		YEAR: "2020",
		MONT: "07",
		RDAY: "18",
	}

	recordPath := "inc/2020/07/18/inc_2020_07_18_03fb24ea-3a81-4469-8522-7753d643dcfe.wav"

	t.Run("create record info", func(t *testing.T) {
		ri := util.CreateRecordInfo(&bc, storageAddr)
		require.Equal(t, recordInfo, ri)
	})

	t.Run("create record path", func(t *testing.T) {
		path := util.CreateFileRecordPath(recordInfo)
		require.Equal(t, recordPath, path)
	})
}
