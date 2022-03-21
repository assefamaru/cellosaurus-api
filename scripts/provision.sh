#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

echo "== Starting provision =="

echo "== Unregistering submodules =="
git submodule deinit --all

echo "== Initializing submodules =="
git submodule update --init --recursive

./parse.sh
./build.sh
./setup-db.sh

echo "== Finished provision =="
