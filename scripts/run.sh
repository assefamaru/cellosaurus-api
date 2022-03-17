#!/bin/bash

cd "$(dirname "$0")"

PORT=8080 go run ../cmd/api/main.go
