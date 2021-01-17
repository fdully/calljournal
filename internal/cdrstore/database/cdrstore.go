package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	cjmodel "github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/cdrserver/model"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type CDRStoreDB struct {
	db *database.DB
}

func New(db *database.DB) *CDRStoreDB {
	return &CDRStoreDB{db: db}
}

func (db *CDRStoreDB) AddBaseCall(ctx context.Context, bc cjmodel.BaseCall) error {
	var username, callerName, callerNumber, recordName, sipHangupDisposition, hangupCause, sipTermStatus *string

	if bc.Username != "" {
		username = &bc.Username
	}

	if bc.CallerIDName != "" {
		callerName = &bc.CallerIDName
	}

	if bc.CallerIDNumber != "" {
		callerNumber = &bc.CallerIDNumber
	}

	if bc.RecordName != "" {
		recordName = &bc.RecordName
	}

	if bc.SIPHangupDisposition != "" {
		sipHangupDisposition = &bc.SIPHangupDisposition
	}

	if bc.HangupCause != "" {
		hangupCause = &bc.HangupCause
	}

	if bc.SIPTermStatus != "" {
		sipTermStatus = &bc.SIPTermStatus
	}

	return db.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		q := `
			INSERT INTO
				base_calls
				(uuid, username, caller_id_name, caller_id_number, destination_number, cj_direction, start_stamp,
				duration, billsec, record_seconds, record_name, start_epoch, answer_epoch, end_epoch,
				sip_hangup_disposition, hangup_cause, sip_term_status)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
			ON CONFLICT DO NOTHING
		`
		_, err := tx.Exec(ctx, q, bc.UUID, username, callerName, callerNumber, bc.DestinationNumber,
			bc.Direction, bc.StartStamp.Format(time.RFC3339), bc.Duration, bc.Billsec, bc.RecordSeconds,
			recordName, bc.StartEpoch, bc.AnswerEpoch, bc.EndEpoch, sipHangupDisposition, hangupCause, sipTermStatus)
		if err != nil {
			return fmt.Errorf("failed to add record path %s: %w", bc.UUID.String(), err)
		}

		return nil
	})
}

func (db *CDRStoreDB) AddRecordPath(ctx context.Context, storageAddress string, cp *model.CallPath) error {
	if storageAddress == "" {
		return fmt.Errorf("failed, storage address: %w", util.ErrIsEmpty)
	}

	if cp == nil {
		return fmt.Errorf("failed, call path: %w", util.ErrIsNil)
	}

	return db.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		q := `
			INSERT INTO
				records_path
				(uuid, storage_address, direction, year, month, day, name)
			VALUES
				($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT DO NOTHING
		`

		_, err := tx.Exec(ctx, q, cp.UUID, storageAddress, cp.DIRC, cp.YEAR, cp.MONT, cp.DAYX, cp.NAME)
		if err != nil {
			return fmt.Errorf("failed to add record path %s: %w", cp.UUID, err)
		}

		return nil
	})
}

func (db *CDRStoreDB) GetRecordPathByUUID(ctx context.Context, uuid uuid.UUID) (*model.CallPath, error) {
	var cp model.CallPath

	if err := db.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
			SELECT
				uuid, direction, year, month, day, name
			FROM
				records_path
			WHERE
				uuid = $1
			`, uuid.String())

		if err := row.Scan(&cp.UUID, &cp.DIRC, &cp.YEAR, &cp.MONT, &cp.DAYX, &cp.NAME); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return database.ErrNotFound
			}

			return fmt.Errorf("failed to scan row: %w", err)
		}

		return nil
	}); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, database.ErrNotFound
		}

		return nil, fmt.Errorf("failed to get record path by uuid: %w", err)
	}

	return &cp, nil
}

func (db *CDRStoreDB) GetBaseCallByUUID(ctx context.Context, uuid uuid.UUID) (*cjmodel.BaseCall, string, error) {
	var (
		bc *cjmodel.BaseCall
		rn string
	)

	if err := db.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
			SELECT
				u.uuid, u.username, u.caller_id_name, u.caller_id_number, u.destination_number, u.cj_direction, u.start_stamp,
				u.duration, u.billsec, u.record_seconds, u.record_name, u.start_epoch, u.answer_epoch, u.end_epoch,
				u.sip_hangup_disposition, u.hangup_cause, u.sip_term_status, d."name"
			FROM
				base_calls as u
			LEFT JOIN records_path as d ON u.uuid = d.uuid	
			WHERE
				u.uuid = $1
			`, uuid.String())

		var err error
		bc, rn, err = scanOneBaseCall(row)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, rn, database.ErrNotFound
		}

		return nil, rn, fmt.Errorf("failed to get base call by uuid: %w", err)
	}

	return bc, rn, nil
}

func scanOneBaseCall(row pgx.Row) (*cjmodel.BaseCall, string, error) {
	var (
		bc                   cjmodel.BaseCall
		rn                   sql.NullString
		username             sql.NullString
		callerName           sql.NullString
		callerNumber         sql.NullString
		recordName           sql.NullString
		sipHangupDisposition sql.NullString
		hangupCause          sql.NullString
		sipTermStatus        sql.NullString
	)

	if err := row.Scan(&bc.UUID, &username, &callerName, &callerNumber, &bc.DestinationNumber, &bc.Direction,
		&bc.StartStamp, &bc.Duration, &bc.Billsec, &bc.RecordSeconds, &recordName, &bc.StartEpoch,
		&bc.AnswerEpoch, &bc.EndEpoch, &sipHangupDisposition, &hangupCause, &sipTermStatus, &rn); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rn.String, database.ErrNotFound
		}

		return nil, rn.String, fmt.Errorf("failed to scan base call: %w", err)
	}

	if username.Valid {
		bc.Username = username.String
	}

	if callerName.Valid {
		bc.CallerIDName = callerName.String
	}

	if callerNumber.Valid {
		bc.CallerIDNumber = callerNumber.String
	}

	if recordName.Valid {
		bc.RecordName = recordName.String
	}

	if sipHangupDisposition.Valid {
		bc.SIPHangupDisposition = sipHangupDisposition.String
	}

	if sipTermStatus.Valid {
		bc.SIPTermStatus = sipTermStatus.String
	}

	return &bc, rn.String, nil
}
