package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/fdully/calljournal/internal/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeMetricsIfPrometheus(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	metricsAddr := os.Getenv("METRICS_ADDR")
	if metricsAddr == "" {
		logger.Error("'prometheus' METRICS_ADDR is not set")

		return nil
	}

	srv := http.Server{Addr: metricsAddr}

	// to shutdown metrics server
	go func(ctx context.Context) {
		<-ctx.Done()

		const shutdownTimeout = 200

		shutdownCtx, done := context.WithTimeout(ctx, time.Millisecond*shutdownTimeout)
		defer done()

		_ = srv.Shutdown(shutdownCtx)

		logger.Info("metrics server shut down.")
	}(ctx)

	// serving http prometheus metrics
	go func(ctx context.Context) {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		logger.Infof("Metrics endpoint listening on %s", metricsAddr)

		srv.Handler = mux
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error while serving metrics endpoint: %v", err)
		}

		logger.Info("Metrics is done.")
	}(ctx)

	return nil
}
