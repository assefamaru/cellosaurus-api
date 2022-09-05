package logging

type Severity int32

const (
	Default Severity = iota
	Info
	Warning
	Error
	Critical
)
