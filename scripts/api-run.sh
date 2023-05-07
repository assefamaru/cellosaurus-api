#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

go run cmd/server/main.go
