package home

import (
	"context"
	"net/http"

	"github.com/fdully/calljournal/internal/templates"
)

func Handle(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = templates.Execute(w, "home.page.gohtml", nil)
	})
}
