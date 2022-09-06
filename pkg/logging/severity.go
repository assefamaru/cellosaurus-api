package logging

type Severity string

const (
	DEFAULT  Severity = "DEFAULT"
	INFO     Severity = "INFO"
	WARNING  Severity = "WARNING"
	ERROR    Severity = "ERROR"
	CRITICAL Severity = "CRITICAL"
)
