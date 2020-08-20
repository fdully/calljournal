package util_test

import (
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/pb"
	"github.com/fdully/calljournal/internal/util"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestBaseCall(t *testing.T) {
	timeLayout := "2006-01-02 15:04:05"
	stti, err := time.Parse(timeLayout, "2020-07-13 19:59:26")
	require.NoError(t, err)

	id, err := uuid.Parse("03fb24ea-3a81-4469-8522-7753d643dcfe")
	require.NoError(t, err)

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
		RECL: "03fb24ea-3a81-4469-8522-7753d643dcfe.wav",
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

	pbStti, err := ptypes.TimestampProto(stti)
	require.NoError(t, err)

	pbBC := pb.BaseCall{
		Uuid: id.String(),
		Clid: "79002800421",
		Clna: "79002800421",
		Dest: "9000",
		Dirc: "inc",
		Stti: pbStti,
		Durs: 188,
		Bils: 178,
		Recd: true,
		Recs: 178,
		Recl: "03fb24ea-3a81-4469-8522-7753d643dcfe.wav",
		Rtag: "public",
		Epos: 1595102366,
		Epoa: 1595102367,
		Epoe: 1595102384,
		Tag1: "",
		Tag2: "",
		Tag3: "",
		Wbye: "recv_bye",
		Hang: "NORMAL_CLEARING",
		Code: "200",
	}

	t.Run("basecall to pb basecall", func(t *testing.T) {
		p, err := util.BaseCallToProtobufBaseCall(&bc)
		require.NoError(t, err)
		require.Equal(t, &pbBC, p)
	})

	t.Run("pb basecall to basecall", func(t *testing.T) {
		b, err := util.ProtobufBaseCallToBaseCall(&pbBC)
		require.NoError(t, err)
		require.Equal(t, &bc, b)
	})
}
