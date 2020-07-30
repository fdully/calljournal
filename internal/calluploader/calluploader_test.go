package calluploader

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/util"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestParseCall(t *testing.T) {
	baseCallJson, err := ioutil.ReadFile("../../testdata/base_call_inc.json")
	require.NoError(t, err)

	stti, err := time.Parse("2006-01-02 15:04:05", "2020-07-18 19:59:26")
	require.NoError(t, err)

	testBaseCall := &model.BaseCall{UUID: uuid.UUID{0xd8, 0xdd, 0xd, 0x47, 0xd, 0xad, 0x4b, 0xe0, 0x86, 0x8b, 0xa7, 0x10, 0xb8, 0xf9, 0x4d, 0x84},
		CLID: "79002800421", CLNA: "79002800421", DEST: "84957899090", DIRC: "inc", STTI: stti, DURS: 18, BILS: 17, RECD: false, RECS: 0, RECL: "", RTAG: "",
		EPOS: 1595102366, EPOA: 1595102367, EPOE: 1595102384, WBYE: "recv_bye", HANG: "NORMAL_CLEARING", CODE: "200"}

	t.Run("plain parse of call", func(t *testing.T) {
		c, err := ParseCall(context.Background(), baseCallJson)
		require.NoError(t, err)
		require.Equal(t, testBaseCall, c)
	})
}

func TestCreateAudioRecord(t *testing.T) {
	testUUID, err := uuid.Parse("03fb24ea-3a81-4469-8522-7753d643dcfe")
	require.NoError(t, err)

	startTime, err := time.Parse(time.RFC3339, "2020-07-18T19:59:26+03:00")
	require.NoError(t, err)

	files, err := LookForFiles("../../testdata", ".wav", 5*time.Second)
	require.NoError(t, err)
	require.Len(t, files, 1)

	fileName := files[0]

	u, err := util.GetUUIDFromString(fileName)
	require.Equal(t, testUUID, u)
	require.NoError(t, err)

	data, err := ioutil.ReadFile(fileName)
	require.NoError(t, err)

	testAudio := &model.AudioRecord{
		UUID:       testUUID,
		ObjectName: fileName,
		StartTime:  startTime,
		Data:       data,
	}

	audio := CreateAudioRecord(u, fileName, startTime, data)
	require.Equal(t, testAudio, audio)
}
