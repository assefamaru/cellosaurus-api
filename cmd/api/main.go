package main

import (
	"flag"
	"os"
	"time"

	"github.com/assefamaru/cellosaurus-api/pkg/api"
	"github.com/assefamaru/cellosaurus-api/pkg/logging"
	"github.com/gin-contrib/cors"
)

const (
	latestVersion = "42"
)

func main() {
	mode := flag.String("mode", fromEnvOrDefaultVal("MODE", "release"), "Gin server mode")
	port := flag.String("port", fromEnvOrDefaultVal("PORT", "8080"), "API server port")
	version := flag.String("version", fromEnvOrDefaultVal("VERSION", latestVersion), "The current Cellosaurus data version")
	sentryDSN := flag.String("sentry-dsn", fromEnvOrDefaultVal("CELLOSAURUS_SENTRY_DSN", ""), "Sentry DSN")
	flag.Parse()

	if err := logging.NewSentryLogger(*sentryDSN); err != nil {
		logging.Errorf("initialize Sentry SDK: %v", err)
	}

	cors := &cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	server := api.NewServer(*mode, *port, *version, cors)
	server.Run()
}

// fromEnvOrDefaultVal returns an environment variable value if it exists,
// or the specified default value.
func fromEnvOrDefaultVal(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}
	return defaultVal
}
