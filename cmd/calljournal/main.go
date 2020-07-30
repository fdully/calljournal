package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fdully/calljournal/internal/calluploader"

	cjdb "github.com/fdully/calljournal/internal/calljournal/database"
	"github.com/fdully/calljournal/internal/database"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/sethvargo/go-envconfig"
	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, done := signalcontext.OnInterrupt()
	defer done()

	var loggerConfig logging.Config
	if err := envconfig.Process(ctx, &loggerConfig); err != nil {
		log.Fatalf("failed to parse logger config %v\n", err)
	}
	if err := logging.CreateLogger(&loggerConfig); err != nil {
		log.Fatalf("failed to create logger %v\n", err)
	}
	logger := logging.FromContext(ctx)
	ctx = logging.WithLogger(ctx, logger)

	logger.Info("starting application")
	if err := realMain(ctx); err != nil {
		logger.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

	var dbConfig database.Config
	if err := envconfig.Process(ctx, &dbConfig); err != nil {
		logger.Fatalf("failed to parse database config %v\n", err)
	}
	db, err := database.NewDB(ctx, &dbConfig)
	if err != nil {
		logger.Fatalf("failed to start database connection %v\n", err)
	}
	defer db.Close(ctx)

	cjDB := cjdb.NewCallJournalDB(db)

	bcData, err := ioutil.ReadFile("./testdata/base_call_inc.json")
	if err != nil {
		logger.Fatal(err)
	}

	bc, err := calluploader.ParseCall(ctx, bcData)
	if err != nil {
		logger.Fatal(err)
	}

	err = cjDB.AddBaseCall(ctx, bc)
	if err != nil {
		logger.Fatal(err)
	}

	d, err := cjDB.GetBaseCall(ctx, bc.UUID)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(d)

	err = cjDB.DeleteBaseCall(ctx, bc.UUID)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debugf("getting out from app")
	return nil
}
