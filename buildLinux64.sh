#!/bin/bash
export CGO_ENABLED=0 
export GOOS=linux 
export GOARCH=amd64 
#go build -o build/UploadRysnc src/opslabgo/main.go
go build -o build/cmdHpler.exe cmdhepler/main.go
go build -o build/httpcmd httpcmd/main.go