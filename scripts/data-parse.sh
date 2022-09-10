#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

rm -rf data
mkdir -p data

# Parse Cellosaurus raw data from text files, and
# output parsed data to "data" directory.
go run ./cmd/scanner
