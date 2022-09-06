package logging

import (
	"fmt"
	"log"
	"os"
)

type LocalLogger struct {
	*log.Logger
}

// NewLocalLogger initializes a new LocalLogger with
// specific severity.
func NewLocalLogger(severity Severity) *LocalLogger {
	var prefix string
	switch severity {
	case DEFAULT:
		prefix = fmt.Sprintf("%-8v ", "")
	default:
		prefix = fmt.Sprintf("%-8v ", severity)
	}
	return &LocalLogger{
		Logger: log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Close is a no-op.
func (l *LocalLogger) Close() {}
