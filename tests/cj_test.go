// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/fdully/calljournal/internal/logging"

	"github.com/google/uuid"

	cdrstoredb "github.com/fdully/calljournal/internal/cdrstore/database"

	"github.com/fdully/calljournal/internal/storage"

	"github.com/fdully/calljournal/internal/database"

	"github.com/fdully/calljournal/internal/setup"
	"github.com/stretchr/testify/require"
)

type Integration struct {
	Database  database.Config
	Blobstore storage.Config
}

func (c *Integration) BlobstoreConfig() *storage.Config {
	return &c.Blobstore
}

func (c *Integration) DatabaseConfig() *database.Config {
	return &c.Database
}

func TestCDRIntergration(t *testing.T) {
	ctx := context.Background()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting integration tests")

	var config Integration

	env, err := setup.Setup(ctx, &config)
	require.NoError(t, err)

	defer env.Close(ctx)

	db := cdrstoredb.New(env.Database())

	t.Run("cdr", func(t *testing.T) {
		uuids := []string{"3f78ec6c-54d0-4b0a-991b-9e5eecf6934d", "4b818455-0f50-4a97-b9e2-23ec5b3a2e27",
			"622c01e5-7b1e-46a8-b1c3-62f7c531ebaa", "a5c191b2-734d-4450-ae23-9062b1226718"}

		for _, v := range uuids {
			vv := uuid.MustParse(v)
			_, _, err := db.GetBaseCallByUUID(ctx, vv)
			require.NoError(t, err)
		}

		uu := uuid.MustParse("622c01e5-7b1e-46a8-b1c3-62f7c531ebaa")
		_, err = db.GetRecordPathByUUID(ctx, uu)
		require.NoError(t, err)

		uu = uuid.MustParse("000c01e5-7b1e-46a8-b1c3-62f7c531ebaa")
		_, err = db.GetRecordPathByUUID(ctx, uu)
		require.Equal(t, database.ErrNotFound, err)
	})

}
