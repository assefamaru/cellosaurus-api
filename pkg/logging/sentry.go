package logging

import "github.com/getsentry/sentry-go"

// Sentry writes error events to Sentry.
// Ensure sentry.Init() is called first
// to initialize the SDK with options.
func Sentry(err error) {
	sentry.CaptureException(err)
}
