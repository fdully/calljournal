package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type CallJournalDB struct {
	db *database.DB
}

func NewCallJournalDB(db *database.DB) *CallJournalDB {
	return &CallJournalDB{db: db}
}

func (c *CallJournalDB) AddBaseCall(ctx context.Context, bc *model.BaseCall) error {
	return c.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		var recl, rtag *string
		var recs *int32
		if bc.RECD {
			recl = &bc.RECL
			rtag = &bc.RTAG
			recs = &bc.RECS
		}
		q := `
			INSERT INTO
				base_calls
				(uuid, clid, clna, dest, dirc, stti, durs, bils, recd, recs, recl, rtag, epos, epoa, epoe, wbye, hang, code)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
			ON CONFLICT DO NOTHING
		`
		_, err := tx.Exec(ctx, q, bc.UUID.String(), bc.CLID, bc.CLNA, bc.DEST, bc.DIRC, bc.STTI.Format(time.RFC3339), bc.DURS, bc.BILS,
			bc.RECD, recs, recl, rtag, bc.EPOS, bc.EPOA, bc.EPOE, bc.WBYE, bc.HANG, bc.CODE)
		if err != nil {
			return fmt.Errorf("failed to add call: %w", err)
		}
		return nil
	})
}

func (c *CallJournalDB) GetBaseCall(ctx context.Context, uuid uuid.UUID) (*model.BaseCall, error) {
	var bc model.BaseCall

	var recs sql.NullInt32
	var recl, rtag sql.NullString

	if err := c.db.InTx(ctx, pgx.ReadCommitted, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
			SELECT
				uuid, clid, clna, dest, dirc, stti, durs, bils, recd, recs, recl, rtag, epos, epoa, epoe, wbye, hang, code
			FROM
				base_calls
			WHERE
				uuid = $1
		`, uuid)

		if err := row.Scan(&bc.UUID, &bc.CLID, &bc.CLNA, &bc.DEST, &bc.DIRC,
			&bc.STTI, &bc.DURS, &bc.BILS, &bc.RECD, &recs, &recl, &rtag,
			&bc.EPOS, &bc.EPOA, &bc.EPOE, &bc.WBYE, &bc.HANG, &bc.CODE); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return database.ErrNotFound
			}
			return fmt.Errorf("failed to scan base call: %w", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to get base call by uuid: %w", err)
	}

	if recs.Valid {
		bc.RECS = recs.Int32
	}
	if recl.Valid {
		bc.RECL = recl.String
	}
	if rtag.Valid {
		bc.RTAG = rtag.String
	}

	return &bc, nil
}

func (c *CallJournalDB) DeleteBaseCall(ctx context.Context, uuid uuid.UUID) error {
	var count int64
	err := c.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		result, err := tx.Exec(ctx, `
			DELETE FROM
				base_calls
			WHERE
				uuid = $1
			`, uuid)
		if err != nil {
			return fmt.Errorf("failed to delete base call: %w", err)
		}
		count = result.RowsAffected()
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete base call: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("no rows were deleted")
	}
	return nil
}
