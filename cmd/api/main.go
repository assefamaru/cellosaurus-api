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
		log.Print("[WARNING] CELLOSAURUS_SENTRY_DSN env missing")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
	})
	if err != nil {
		log.Print(fmt.Sprintf("[WARNING] sentry.Init: %s", err))
	}

	api.Init(ctx)
}
