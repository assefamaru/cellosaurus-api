#!/bin/bash

echo "== initialize submodules =="
git submodule update --init --recursive

echo "== parse cellosaurus text files =="
./parser.sh

echo "== build executables =="
./build.sh

echo "== setup mysql =="
./db.sh
