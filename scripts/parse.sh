#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

rm -rf data
mkdir -p data

echo "== Parsing Cellosaurus raw data =="
go run cmd/scanner/main.go
