// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/fdully/calljournal/internal/calljournal"
	"github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUploader(t *testing.T) {
	ctx := context.Background()
	logging.Init(ctx)

	var config calljournal.Config
	env, err := setup.Setup(ctx, &config)
	require.NoError(t, err)

	defer env.Close(ctx)

	cjDB := database.NewCallJournalDB(env.Database())

	id, err := uuid.Parse("03fb24ea-3a81-4469-8522-7753d643dcfe")
	require.NoError(t, err)

	_, err = cjDB.GetBaseCall(ctx, id)
	require.NoError(t, err)

	id, err = uuid.Parse("d8dd0d47-0dad-4be0-868b-a710b8f94d84")
	require.NoError(t, err)

	_, err = cjDB.GetBaseCall(ctx, id)
	require.NoError(t, err)

	err = cjDB.DeleteBaseCall(ctx, id)
	require.NoError(t, err)
}
