package main

import (
	"fmt"
	"log"
	"os"

	"github.com/assefamaru/cellosaurus-api/pkg/api"
	"github.com/getsentry/sentry-go"
)

func main() {
	var ctx api.Context

	ctx.Mode = "release"
	ctx.Port = os.Getenv("PORT")
	if ctx.Port == "" {
		log.Fatal("PORT must be set")
	}

	sentryDsn := os.Getenv("CELLOSAURUS_SENTRY_DSN")
	if sentryDsn == "" {
		logWarning("CELLOSAURUS_SENTRY_DSN env must be set")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
	})
	if err != nil {
		logWarning(fmt.Sprintf("sentry.Init: %s", err))
	}

	api.Init(ctx)
}

func logWarning(message string) {
	warning := log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
	warning.Println(message)
}
