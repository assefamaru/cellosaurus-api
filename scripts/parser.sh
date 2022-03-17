#!/bin/bash

cd "$(dirname "$0")"

rm -rf ../data
mkdir -p ../data

go run ../cmd/scanner/main.go
