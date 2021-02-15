package search

import (
	"context"
	"net/http"
	"strings"

	"github.com/fdully/calljournal/internal/calljournal"
	cjdatabase "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/calljournal/model"
	cdrdatabase "github.com/fdully/calljournal/internal/cdrstore/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/templates"
	"github.com/google/uuid"
)

func Handle(ctx context.Context, c *calljournal.Config, env *serverenv.ServerEnv) http.Handler {
	h := newHandler(ctx, c, env)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := h.handleRequest(w, r)

		if resp.err != nil {
			logger := logging.FromContext(ctx)

			logger.Error(resp.err)
			w.WriteHeader(resp.status)
			_, _ = w.Write([]byte(http.StatusText(resp.status)))

			return
		}

		if resp.status != http.StatusOK {
			w.WriteHeader(resp.status)
			_, _ = w.Write([]byte(http.StatusText(resp.status)))

			return
		}

		var tmplData templateData
		tmplData.Search = resp.search
		if resp.calls != nil {
			tmplData.Calls = resp.calls
		}

		if err := templates.Check("search.page.gohtml", tmplData); err != nil {
			logger := logging.FromContext(ctx)
			logger.Errorf("failed to check template: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

			return
		}

		w.WriteHeader(http.StatusOK)
		_ = templates.Execute(w, "search.page.gohtml", tmplData)
	})
}

func newHandler(ctx context.Context, config *calljournal.Config, env *serverenv.ServerEnv) *handler {
	return &handler{
		config:    config,
		blobstore: env.Blobstore(),
		cjdb:      cjdatabase.New(env.Database()),
		cdrdb:     cdrdatabase.New(env.Database()),
	}
}

type handler struct {
	config    *calljournal.Config
	blobstore storage.Blobstore
	cdrdb     *cdrdatabase.CDRStoreDB
	cjdb      *cjdatabase.CJDB
}

type response struct {
	status int
	err    error
	calls  []*model.Call
	search string
}

type templateData struct {
	Search string
	Calls  []*model.Call
}

func (h *handler) handleRequest(w http.ResponseWriter, r *http.Request) *response {
	searchStr := strings.TrimSpace(r.URL.Query().Get("search"))
	if searchStr == "" {
		return &response{
			status: http.StatusOK,
		}
	}

	// search by uuid
	id, err := uuid.Parse(searchStr)
	if err == nil {
		return h.searchCallByUUID(r.Context(), id, searchStr)
	}

	// search by telephone number
	if isSearchByNumber(searchStr) {
		return h.searchCallByNumber(r.Context(), searchStr)
	}

	// no search
	return &response{
		status: http.StatusOK,
		err:    nil,
		search: searchStr,
	}
}
