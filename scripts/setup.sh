#!/bin/bash

echo "== initialize submodules =="
git submodule update --init --recursive

echo "== parse cellosaurus text files =="
./parser.sh

echo "== setup mysql =="
./db.sh
