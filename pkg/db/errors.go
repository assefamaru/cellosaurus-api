package db

import "errors"

var (
	errMissingEnv = errors.New("missing environment variable")
)
