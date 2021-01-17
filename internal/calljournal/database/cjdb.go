package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	cjmodel "github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/jackc/pgx/v4"
)

type CJDB struct {
	db *database.DB
}

func New(db *database.DB) *CJDB {
	return &CJDB{db: db}
}

func (db *CJDB) SearchCallsByNumber(ctx context.Context, phone string) ([]*cjmodel.Call, error) {
	calls := make([]*cjmodel.Call, 0, 2000)

	if err := db.db.InTx(ctx, pgx.Serializable, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, `
			SELECT
				u.uuid, u.username, u.caller_id_name, u.caller_id_number, u.destination_number, u.cj_direction, u.start_stamp,
				u.duration, u.billsec, u.record_seconds, u.record_name, u.start_epoch, u.answer_epoch, u.end_epoch,
				u.sip_hangup_disposition, u.hangup_cause, u.sip_term_status, d."name"
			FROM
				base_calls as u
			LEFT JOIN records_path as d ON u.uuid = d.uuid
			WHERE
				substr(u.caller_id_number,length(u.caller_id_number)-9,10) = $1
			OR
				substr(u.destination_number,length(u.destination_number)-9,10) = $1
			ORDER BY start_stamp DESC LIMIT 2000
		`, phone)
		if err != nil {
			return fmt.Errorf("failed to get from db: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Err(); err != nil {
				return fmt.Errorf("failed to iterate: %w", err)
			}

			bc, record, err := scanOneBaseCall(rows)
			if err != nil {
				return fmt.Errorf("failed to parse: %w", err)
			}

			calls = append(calls, &cjmodel.Call{
				BaseCall:    bc,
				RecordName:  record,
				Disconnect:  cdrutil.WhoDisconnect(bc),
				ConnectTime: bc.Duration - bc.Billsec,
			})
		}

		return nil
	}); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, database.ErrNotFound
		}

		return nil, fmt.Errorf("failed to get calls by number: %w", err)
	}

	return calls, nil
}

func scanOneBaseCall(row pgx.Row) (*cjmodel.BaseCall, string, error) {
	var (
		bc                   cjmodel.BaseCall
		record               sql.NullString
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
		&bc.AnswerEpoch, &bc.EndEpoch, &sipHangupDisposition, &hangupCause, &sipTermStatus, &record); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, record.String, database.ErrNotFound
		}

		return nil, record.String, fmt.Errorf("failed to scan base call: %w", err)
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

	return &bc, record.String, nil
}
