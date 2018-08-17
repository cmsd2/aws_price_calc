#!/bin/bash
set -eux
go generate
export GOOS=linux
export GOARCH=amd64
go build -o main main.go
