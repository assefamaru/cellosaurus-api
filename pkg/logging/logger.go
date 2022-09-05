package logging

import (
	"log"
	"time"
)

const (
	datetimeFormat = "2006-01-02 15:04:05"
)

// Defaultf writes Default severity log messages.
func Defaultf(format string, v ...any) {
	writef(Default, format, v...)
}

// Infof writes Info severity log messages.
func Infof(format string, v ...any) {
	writef(Info, format, v...)
}

// Warningf writes Warning severity log messages.
func Warningf(format string, v ...any) {
	writef(Warning, format, v...)
}

// Errorf writes Error severity log messages.
func Errorf(format string, v ...any) {
	writef(Error, format, v...)
}

// Criticalf writes Critical severity log messages.
func Criticalf(format string, v ...any) {
	writef(Critical, format, v...)
}

// writef prefixes all log events with the right severity and
// current timestamp.
func writef(sev Severity, format string, v ...any) {
	var severity string
	switch sev {
	case Default:
		severity = "DEFAULT"
	case Info:
		severity = "INFO"
	case Warning:
		severity = "WARNING"
	case Error:
		severity = "ERROR"
	case Critical:
		severity = "CRITICAL"
	}
	log.Printf("%-8s %s ", severity, time.Now().Format(datetimeFormat))
	log.Printf(format, v...)
}
