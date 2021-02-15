package call

import (
	"context"
	"errors"
	"net/http"

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

func Handle(ctx context.Context, config *calljournal.Config, env *serverenv.ServerEnv) http.Handler {
	h := newHandler(ctx, config, env)

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
		if resp.call != nil {
			tmplData.Call = resp.call
		}

		if err := templates.Check("call.page.gohtml", tmplData); err != nil {
			logger := logging.FromContext(ctx)
			logger.Errorf("failed to check template: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

			return
		}

		w.WriteHeader(http.StatusOK)
		_ = templates.Execute(w, "call.page.gohtml", tmplData)
	})
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
	call   *model.Call
}

type templateData struct {
	Call *model.Call
}

func newHandler(ctx context.Context, config *calljournal.Config, env *serverenv.ServerEnv) *handler {
	return &handler{
		config:    config,
		blobstore: env.Blobstore(),
		cjdb:      cjdatabase.New(env.Database()),
		cdrdb:     cdrdatabase.New(env.Database()),
	}
}

func (h handler) handleRequest(w http.ResponseWriter, r *http.Request) *response {
	idStr := r.URL.Query().Get("uuid")
	if idStr == "" {
		return &response{
			status: http.StatusBadRequest,
			err:    nil,
		}
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return &response{
			status: http.StatusBadRequest,
			err:    nil,
		}
	}

	return h.process(r.Context(), id)
}

func (h handler) process(ctx context.Context, id uuid.UUID) *response {
	bc, err := h.getCallByUUID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &response{
				status: http.StatusOK,
				err:    nil,
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
		call:   bc,
	}
}

func (h handler) getCallByUUID(ctx context.Context, id uuid.UUID) (*model.Call, error) {
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
