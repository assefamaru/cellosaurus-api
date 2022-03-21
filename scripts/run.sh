#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

PORT=8080 go run cmd/api/main.go
