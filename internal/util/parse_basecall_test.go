package util_test

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestParseCall(t *testing.T) {
	baseCallJSON, err := ioutil.ReadFile("../../testdata/d8dd0d47-0dad-4be0-868b-a710b8f94d84.json")
	require.NoError(t, err)

	stti, err := time.Parse("2006-01-02 15:04:05", "2020-07-18 19:59:26")
	require.NoError(t, err)

	testBaseCall := &model.BaseCall{
		UUID: uuid.UUID{0xd8, 0xdd, 0xd, 0x47, 0xd, 0xad, 0x4b, 0xe0, 0x86, 0x8b, 0xa7, 0x10, 0xb8, 0xf9, 0x4d, 0x84},
		CLID: "79002800421", CLNA: "79002800421", DEST: "84957899090", DIRC: "inc", STTI: stti, DURS: 18,
		BILS: 17, RECD: false, RECS: 0, RECL: "", RTAG: "",
		EPOS: 1595102366, EPOA: 1595102367, EPOE: 1595102384, WBYE: "recv_bye", HANG: "NORMAL_CLEARING", CODE: "200",
	}

	t.Run("plain parse of call", func(t *testing.T) {
		c, err := util.ParseCall(context.Background(), baseCallJSON)
		require.NoError(t, err)
		require.Equal(t, testBaseCall, c)
	})
}
