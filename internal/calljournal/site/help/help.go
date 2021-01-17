package help

import (
	"context"
	"net/http"
	"net/url"

	"github.com/fdully/calljournal/internal/calljournal"
	"github.com/fdully/calljournal/internal/calljournal/model/cjerrors"
	"github.com/fdully/calljournal/internal/templates"
)

func Handle(ctx context.Context, config *calljournal.Config) http.Handler {
	h := newHandler(ctx, config)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := h.handleRequest(w, r)

		if resp == nil {
			w.WriteHeader(http.StatusOK)
			_ = templates.Execute(w, "help.page.gohtml", nil)

			return
		}

		var tmpl templateData
		tmpl.Err = resp.err

		w.WriteHeader(http.StatusOK)
		_ = templates.Execute(w, "help.page.gohtml", tmpl)
	})
}

func newHandler(ctx context.Context, c *calljournal.Config) *handler {
	return &handler{
		config: c,
	}
}

type handler struct {
	config *calljournal.Config
}

type response struct {
	status int
	err    error
}

type templateData struct {
	Err error
}

func (h *handler) handleRequest(w http.ResponseWriter, r *http.Request) *response {
	var msg string

	switch r.Method {
	case "GET":
		msg = r.URL.Query().Get("message")
	case "POST":
		msg = r.PostFormValue("message")
	default:
		return nil
	}

	if msg == "" {
		return nil
	}

	return h.process(r.Context(), msg)
}

func (h *handler) process(ctx context.Context, msg string) *response {
	urlData := url.Values{}
	urlData.Set("message", msg)

	//nolint:noctx
	resp, err := http.PostForm(h.config.CloudAddr+"/api/v1/send/support", urlData)
	if err != nil {
		return &response{
			status: http.StatusOK,
			err:    err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &response{
			status: http.StatusOK,
			err:    cjerrors.ErrEmpty,
		}
	}

	return &response{
		status: http.StatusOK,
		err:    nil,
	}
}
