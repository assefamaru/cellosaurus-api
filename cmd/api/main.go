package main

import (
	"flag"
	"os"
	"time"

	"github.com/assefamaru/cellosaurus-api/pkg/api"
	"github.com/assefamaru/cellosaurus-api/pkg/logging"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
)

const (
	sentryDsnEnv = "CELLOSAURUS_SENTRY_DSN"
)

func main() {
	mode := flag.String("mode", fromEnvOrDefault("MODE", "release"), "Gin server mode")
	port := flag.String("port", fromEnvOrDefault("PORT", "8080"), "API server port")
	flag.Parse()

	sentryDsn := os.Getenv(sentryDsnEnv)
	if sentryDsn == "" {
		logging.Warningf("missing environment variable: %v", sentryDsnEnv)
	}
	options := sentry.ClientOptions{Dsn: sentryDsn}
	if err := sentry.Init(options); err != nil {
		logging.Warningf("sentry.Init: %v", err)
	}

	cors := &cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	server := api.NewServer(*mode, *port, cors)
	server.Run()
}

// fromEnvOrDefault returns the environment variable value, if present,
// or the specified default value.
func fromEnvOrDefault(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}
	return defaultVal
}
