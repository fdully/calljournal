package database

import (
	"context"
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v4"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNotFound    = Error("record not found")
	ErrKeyConflict = Error("key conflict")
	ErrNotEffeced  = Error("no lines were effected")
)

func (db *DB) NullableTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}

	return &t
}

// InTx runs the given function f within a transaction with isolation level isoLevel.
func (db *DB) InTx(ctx context.Context, isoLevel pgx.TxIsoLevel, f func(tx pgx.Tx) error) error {
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquiring connection: %w", err)
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: isoLevel})
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err := f(tx); err != nil {
		if err1 := tx.Rollback(ctx); err1 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err1, err)
		}

		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
