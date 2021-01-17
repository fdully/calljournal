package search

import (
	"context"
	"errors"
	"net/http"

	"github.com/fdully/calljournal/internal/calljournal/model"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/google/uuid"
)

func (h *handler) searchCallByUUID(ctx context.Context, id uuid.UUID, search string) *response {
	bc, err := h.getCallByUUID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &response{
				status: http.StatusOK,
				err:    nil,
				search: search,
			}
		}

		return &response{
			status: http.StatusInternalServerError,
			err:    err,
		}
	}

	return &response{
		status: http.StatusOK,
		err:    nil,
		calls:  []*model.Call{bc},
		search: search,
	}
}

func (h *handler) getCallByUUID(ctx context.Context, id uuid.UUID) (*model.Call, error) {
	a, rn, err := h.cdrdb.GetBaseCallByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Call{
		BaseCall:    a,
		RecordName:  rn,
		ConnectTime: a.Duration - a.Billsec,
		Disconnect:  cdrutil.WhoDisconnect(a),
	}, nil
}
