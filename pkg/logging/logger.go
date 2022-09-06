package logging

// Defaultf writes DEFAULT severity log events.
func Defaultf(format string, v ...any) {
	writef(DEFAULT, format, v...)
}

// Infof writes INFO severity log events.
func Infof(format string, v ...any) {
	writef(INFO, format, v...)
}

// Warningf writes WARNING severity log events.
func Warningf(format string, v ...any) {
	writef(WARNING, format, v...)
}

// Errorf writes ERROR severity log events.
func Errorf(format string, v ...any) {
	writef(ERROR, format, v...)
}

// Criticalf writes CRITICAL severity log events.
func Criticalf(format string, v ...any) {
	writef(CRITICAL, format, v...)
}

// writef writes log events with the appropriate severity.
func writef(severity Severity, format string, v ...any) {
	logger := NewLocalLogger(severity)
	defer logger.Close()

	logger.Printf(format, v...)
}
