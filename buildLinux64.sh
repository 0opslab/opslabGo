#!/bin/bahs
export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64 
go build -o build/gosearch src/opslabgo/gosearch/main.go