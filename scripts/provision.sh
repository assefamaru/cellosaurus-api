#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

CELLOSAURUS_DATA_REPO="https://github.com/calipho-sib/cellosaurus"

echo "== starting provision =="
rm -rf bin cellosaurus data

echo "== building binaries =="
./scripts/api-build.sh

echo "== cloning Cellosaurus repo =="
git clone "$CELLOSAURUS_DATA_REPO" 1> /dev/null 2>&1

echo "== parsing Cellosaurus raw data =="
./scripts/data-parse.sh

echo "== setting up database locally =="
./scripts/data-setup-db.sh

echo "== provision complete =="
