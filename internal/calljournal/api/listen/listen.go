package listen

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/fdully/calljournal/internal/calljournal"
	cdrstoredb "github.com/fdully/calljournal/internal/cdrstore/database"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/serverenv"
	"github.com/fdully/calljournal/internal/storage"
	"github.com/fdully/calljournal/internal/util/cdrutil"
	"github.com/google/uuid"
)

func Handle(ctx context.Context, c *calljournal.Config, env *serverenv.ServerEnv) http.Handler {
	h := newHandler(ctx, c, env)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.FromContext(ctx)
		response := h.handleRequest(w, r)

		if response.err != nil {
			logger.Error(response.err)
			w.WriteHeader(response.status)
			_, _ = w.Write([]byte(http.StatusText(response.status)))

			return
		}

		if response.status == http.StatusNotFound {
			w.WriteHeader(response.status)
			_, _ = w.Write([]byte(http.StatusText(response.status)))

			return
		}

		w.Header().Add("Content-Type", "audio/mpeg")
		w.Header().Add("accept-ranges", "bytes")
		w.Header().Add("Content-Length", strconv.Itoa(response.data.Len()))
		// Content-Range: "bytes 0-18900/18900"
		w.Header().Add("Content-Range", "bytes 0-"+strconv.Itoa(response.data.Len())+"/"+strconv.Itoa(response.data.Len()))
		w.WriteHeader(response.status)

		_, _ = response.data.WriteTo(w)
	})
}

type handler struct {
	config    *calljournal.Config
	blobstore storage.Blobstore
	database  *cdrstoredb.CDRStoreDB
}

func newHandler(ctx context.Context, config *calljournal.Config, env *serverenv.ServerEnv) *handler {
	return &handler{
		config:    config,
		blobstore: env.Blobstore(),
		database:  cdrstoredb.New(env.Database()),
	}
}

type response struct {
	status int
	err    error
	data   *bytes.Buffer
}

func (h *handler) handleRequest(w http.ResponseWriter, r *http.Request) *response {
	idStr := r.URL.Query().Get("uuid")
	if idStr == "" {
		return &response{
			status: http.StatusNotFound,
			err:    nil,
			data:   nil,
		}
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return &response{
			status: http.StatusNotFound,
			err:    nil,
			data:   nil,
		}
	}

	return h.process(r.Context(), id)
}

func (h *handler) process(ctx context.Context, id uuid.UUID) *response {
	recPath, err := h.database.GetRecordPathByUUID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &response{
				status: http.StatusNotFound,
				err:    nil,
				data:   nil,
			}
		}

		return &response{
			status: http.StatusInternalServerError,
			err:    err,
			data:   nil,
		}
	}

	pth := cdrutil.CreateHTTPCallPath(*recPath)

	b, err := h.blobstore.GetObject(ctx, h.config.Bucket, pth)
	if err != nil {
		return &response{
			status: http.StatusInternalServerError,
			err:    err,
			data:   nil,
		}
	}

	return &response{
		status: http.StatusOK,
		err:    nil,
		data:   b,
	}
}
