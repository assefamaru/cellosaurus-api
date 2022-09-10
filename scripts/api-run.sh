#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

go run cmd/api/main.go
