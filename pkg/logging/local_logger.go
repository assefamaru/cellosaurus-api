package logging

import (
	"fmt"
	"log"
	"os"
)

type LocalLogger struct {
}

// NewLocalLogger creates a new LocalLogger.
func NewLocalLogger() *LocalLogger {
	return &LocalLogger{}
}

// Close is a no-op.
func (l *LocalLogger) Close() {}

// Debugf writes DEBUG severity logs to STDERR.
func (l *LocalLogger) Debugf(format string, a ...any) {
	logBySeverity(DEBUG, format, a...)
}

// Infof writes INFO severity logs to STDERR.
func (l *LocalLogger) Infof(format string, a ...any) {
	logBySeverity(INFO, format, a...)
}

// Warningf writes WARNING severity logs to STDERR.
func (l *LocalLogger) Warningf(format string, a ...any) {
	logBySeverity(WARNING, format, a...)
}

// Errorf writes ERROR severity logs to STDERR.
func (l *LocalLogger) Errorf(format string, a ...any) {
	logBySeverity(ERROR, format, a...)
}

// Criticalf writes CRITICAL severity logs to STDERR.
func (l *LocalLogger) Criticalf(format string, a ...any) {
	logBySeverity(CRITICAL, format, a...)
}

// Fatalf writes CRITICAL severity log to STDERR,
// and aborts the caller process by exiting 1.
func (l *LocalLogger) Fatalf(format string, a ...any) {
	l.Criticalf(format, a...)
	os.Exit(1)
}

// logBySeverity emits log events with specified severity.
func logBySeverity(severity Severity, format string, a ...any) {
	switch severity {
	case WARNING, ERROR, CRITICAL:
		// We will internally emit higher
		// severity logs to Sentry. This
		// is a no-op if Sentry has not
		// been initialized on server
		// startup.
		LogSentry(fmt.Errorf(format, a...))
	}
	prefix := fmt.Sprintf("%-8v ", severity)
	logger := log.New(os.Stderr, prefix, log.Ldate|log.Ltime)
	logger.Printf(format, a...)
}
