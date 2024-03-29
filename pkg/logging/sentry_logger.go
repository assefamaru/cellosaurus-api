package logging

import (
	"github.com/getsentry/sentry-go"
)

// NewSentryLogger initializes Sentry SDK
// with options.
func NewSentryLogger(sentryDSN string) error {
	return sentry.Init(sentry.ClientOptions{Dsn: sentryDSN})
}

// LogSentry writes errors to Sentry.
// Ensure NewSentryLogger() is called
// first to initialize the Sentry SDK
// with options.
func LogSentry(err error) {
	sentry.CaptureException(err)
}
