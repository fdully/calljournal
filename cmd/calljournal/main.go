package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fdully/calljournal/bindata/staticfs"
	"github.com/fdully/calljournal/internal/calljournal"
	apicdrxml "github.com/fdully/calljournal/internal/calljournal/api/cdrxml"
	apidownload "github.com/fdully/calljournal/internal/calljournal/api/download"
	apilisten "github.com/fdully/calljournal/internal/calljournal/api/listen"
	sitecall "github.com/fdully/calljournal/internal/calljournal/site/call"
	sitedebug "github.com/fdully/calljournal/internal/calljournal/site/debug"
	sitehelp "github.com/fdully/calljournal/internal/calljournal/site/help"
	sitehome "github.com/fdully/calljournal/internal/calljournal/site/home"
	sitesearch "github.com/fdully/calljournal/internal/calljournal/site/search"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/server"
	"github.com/fdully/calljournal/internal/setup"
	"github.com/fdully/calljournal/internal/templates"
	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()

	logging.Init(ctx)

	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting calljournal")

	// parse all templates and layouts and create store
	templates.Init(ctx)

	err := realMain(ctx)

	done()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("exiting calljournal")
}

func realMain(ctx context.Context) error {
	var config calljournal.Config

	env, err := setup.Setup(ctx, &config)
	if err != nil {
		return fmt.Errorf("failed setup %w", err)
	}

	mux := http.DefaultServeMux
	mux.Handle("/static/", http.FileServer(staticfs.AssetFile()))
	mux.Handle("/api/v1/listen", apilisten.Handle(ctx, &config, env))
	mux.Handle("/api/v1/cdrxml", apicdrxml.Handle(ctx, &config, env))
	mux.Handle("/api/v1/call/download", apidownload.Handle(ctx, &config, env))
	mux.Handle("/way188/debug", sitedebug.Handle(ctx, &config, env))
	mux.Handle("/way188/home", sitehome.Handle(ctx))
	mux.Handle("/way188/search", sitesearch.Handle(ctx, &config, env))
	mux.Handle("/way188/help", sitehelp.Handle(ctx, &config))
	mux.Handle("/way188/call", sitecall.Handle(ctx, &config, env))

	var srv *server.Server

	if config.TLSEnabled {
		srv, err = server.NewTLSServer(config.Addr)
	} else {
		srv, err = server.NewServer(config.Addr)
	}

	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	return srv.ServeHTTPHandler(ctx, mux)
}
