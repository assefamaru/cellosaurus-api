package logging

type Severity string

const (
	DEBUG    Severity = "DEBUG"
	INFO     Severity = "INFO"
	WARNING  Severity = "WARNING"
	ERROR    Severity = "ERROR"
	CRITICAL Severity = "CRITICAL"
)
