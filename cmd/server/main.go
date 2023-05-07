package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/assefamaru/cellosaurus-api/internal/apiserver"
	"github.com/assefamaru/cellosaurus-api/pkg/db"
	"github.com/assefamaru/cellosaurus-api/pkg/logging"
)

const (
	defaultVersion = "v42"
)

var (
	mode      = flag.String("mode", envOrDefault("MODE", "release"), "Gin server mode")
	port      = flag.String("port", envOrDefault("PORT", "8080"), "API server port")
	version   = flag.String("version", envOrDefault("VERSION", defaultVersion), "The current Cellosaurus data version")
	sentryDSN = flag.String("sentry-dsn", envOrDefault("CELLOSAURUS_SENTRY_DSN", ""), "Sentry DSN")
)

func main() {
	flag.Parse()
	if err := logging.NewSentryLogger(*sentryDSN); err != nil {
		logging.NewLocalLogger().Errorf("initialize Sentry SDK: %v", err)
	}
	ctx := context.Background()
	client, err := db.NewMySQLClient(ctx)
	if err != nil {
		logging.NewLocalLogger().Fatalf("create new MySQL client: %v", err)
	}
	defer client.Close()
	addr, mode, version := fmt.Sprintf(":%s", *port), *mode, *version
	server := apiserver.New(ctx, client, addr, mode, version)
	server.Run()
}

func envOrDefault(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}
	return defaultVal
}
