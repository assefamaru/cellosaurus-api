#!/bin/bash

rm -rf ../data
mkdir -p ../data

go run ../cmd/scanner/main.go
