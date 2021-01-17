package search

import (
	"context"
	"errors"
	"net/http"

	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/util/phoneutil"
)

const (
	maxPhoneLen = 30
	minPhoneLen = 10
)

func isSearchByNumber(search string) bool {
	if len(search) > maxPhoneLen {
		return false
	}

	phone := phoneutil.NormalizePhone(search)

	return phoneutil.HasOnlyDigits(phone)
}

func (h *handler) searchCallByNumber(ctx context.Context, phone string) *response {
	phone = phoneutil.NormalizePhone(phone)
	if len(phone) < minPhoneLen {
		return &response{
			status: http.StatusOK,
			err:    nil,
			search: phone,
		}
	}

	calls, err := h.cjdb.SearchCallsByNumber(ctx, phone[len(phone)-10:])
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &response{
				status: http.StatusOK,
				err:    nil,
				search: phone,
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
		calls:  calls,
		search: phone,
	}
}
