package cdrutil_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/fdully/calljournal/internal/cdrclient/callfiles"
	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestCDR(t *testing.T) {
	glbCtx := context.Background()

	logger := zap.NewNop()
	glbCtx = logging.WithLogger(glbCtx, logger.Sugar())

	testFolder := "../../../testdata"

	var testCDR []model.CDR

	cdrFiles, err := ioutil.ReadDir(testFolder)
	require.NoError(t, err)

	t.Run("parse cdr xml", func(t *testing.T) {
		for _, v := range cdrFiles {
			if !v.Mode().IsRegular() {
				continue
			}

			// Check if file has correct extension
			if filepath.Ext(v.Name()) != callfiles.CDRFilesExt {
				continue
			}

			cdrData, err := ioutil.ReadFile(filepath.Join(testFolder, v.Name()))
			require.NoError(t, err)

			cdr, err := cdrutil.ParseCDR(cdrData)
			require.NoError(t, err)

			testCDR = append(testCDR, cdr)
		}

		require.Len(t, testCDR, 4)
	})

	t.Run("cdr path info from cdr", func(t *testing.T) {
		for _, cdr := range testCDR {
			uuidStr := cdr.Variables.UUID
			id, err := uuid.Parse(uuidStr)
			require.NoError(t, err)

			dirc := cdr.Variables.CJCallDirection
			name := cdr.Variables.UUID + ".xml"

			stti, err := time.Parse("2006-01-02 15:04:05", string(cdr.Variables.StartStamp))
			require.NoError(t, err)

			year := strconv.Itoa(stti.Year())
			mont := fmt.Sprintf("%02d", int(stti.Month()))
			dayx := fmt.Sprintf("%02d", stti.Day())

			expected := model.CallPath{
				UUID: id,
				DIRC: dirc,
				YEAR: year,
				MONT: mont,
				DAYX: dayx,
				NAME: name,
			}

			cdrPathInfo, err := cdrutil.CDRPathInfoFromCDR(glbCtx, cdr)
			require.NoError(t, err)

			require.Equal(t, expected, cdrPathInfo)
		}
	})

	t.Run("record path info from cdr", func(t *testing.T) {
		for _, cdr := range testCDR {
			uuidStr := cdr.Variables.UUID
			id, err := uuid.Parse(uuidStr)
			require.NoError(t, err)

			dirc := cdr.Variables.CJCallDirection
			name := cdr.Variables.CJRecordName

			stti, err := time.Parse("2006-01-02 15:04:05", string(cdr.Variables.StartStamp))
			require.NoError(t, err)

			year := strconv.Itoa(stti.Year())
			mont := fmt.Sprintf("%02d", int(stti.Month()))
			dayx := fmt.Sprintf("%02d", stti.Day())

			expected := model.CallPath{
				UUID: id,
				DIRC: dirc,
				YEAR: year,
				MONT: mont,
				DAYX: dayx,
				NAME: string(name),
			}

			cdrPathInfo, err := cdrutil.RecordPathInfoFromCDR(glbCtx, cdr)
			require.NoError(t, err)

			require.Equal(t, expected, cdrPathInfo)
		}
	})

	t.Run("cdr to basecall", func(t *testing.T) {
		for _, cdr := range testCDR {
			bc, err := cdrutil.CDRToBaseCall(cdr)
			require.NoError(t, err)

			require.Equal(t, cdr.Variables.UUID, bc.UUID.String())
		}
	})
}
