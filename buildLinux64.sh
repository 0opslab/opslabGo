#!/bin/bahs
export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64 
go build -o build/weixin src/base/web/04web/main.go 