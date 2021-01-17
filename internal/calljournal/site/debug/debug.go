package debug

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/fdully/calljournal/internal/calljournal"
	cjdatabase "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/calljournal/model"
	cdrdatabase "github.com/fdully/calljournal/internal/cdrstore/database"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/templates"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/google/uuid"
)

type handler struct {
	config    *calljournal.Config
	blobstore storage.Blobstore
	cdrdb     *cdrdatabase.CDRStoreDB
	cjdb      *cjdatabase.CJDB
}

func new(ctx context.Context, config *calljournal.Config, env *serverenv.ServerEnv) *handler {
	return &handler{
		config:    config,
		blobstore: env.Blobstore(),
		cdrdb:     cdrdatabase.New(env.Database()),
		cjdb:      cjdatabase.New(env.Database()),
	}
}

func Handle(ctx context.Context, c *calljournal.Config, env *serverenv.ServerEnv) http.Handler {
	h := new(ctx, c, env)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.FromContext(ctx)
		response := h.handleRequest(w, r)

		if response.err != nil {
			logger.Error(response.err)
			w.WriteHeader(response.status)
			_, _ = w.Write([]byte(http.StatusText(response.status)))

			return
		}

		if response.status != http.StatusOK {
			w.WriteHeader(response.status)
			_, _ = w.Write([]byte(http.StatusText(response.status)))

			return
		}

		var tmplData templateData

		if response.uuid != uuid.Nil {
			tmplData.UUID = response.uuid.String()
		}

		tmplData.Search = response.search
		tmplData.RecordName = response.recordName
		tmplData.Call = response.call
		if response.call != nil {
			tmplData.ConnectionTime = response.call.Duration - response.call.Billsec
			tmplData.Disconnect = cdrutil.WhoDisconnect(response.call)
		}

		if err := templates.Check("debug.page.gohtml", tmplData); err != nil {
			logger.Errorf("failed to check template: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

			return
		}

		w.WriteHeader(response.status)
		_ = templates.Execute(w, "debug.page.gohtml", tmplData)
	})
}

type response struct {
	status     int
	err        error
	uuid       uuid.UUID
	search     string
	call       *model.BaseCall
	recordName string
}

type templateData struct {
	UUID           string
	Search         string
	Call           *model.BaseCall
	ConnectionTime int64
	RecordName     string
	Disconnect     string
}

func (h *handler) handleRequest(w http.ResponseWriter, r *http.Request) *response {
	searchStr := strings.TrimSpace(r.URL.Query().Get("search"))
	if searchStr == "" {
		return &response{
			status: http.StatusOK,
			err:    nil,
		}
	}

	id, err := uuid.Parse(searchStr)
	if err != nil {
		return &response{
			status: http.StatusOK,
			err:    nil,
			search: searchStr,
		}
	}

	return h.process(r.Context(), id, searchStr)
}

func (h *handler) process(ctx context.Context, id uuid.UUID, search string) *response {
	a, rn, err := h.cdrdb.GetBaseCallByUUID(ctx, id)
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
		status:     http.StatusOK,
		err:        nil,
		uuid:       id,
		search:     search,
		call:       a,
		recordName: rn,
	}
}
