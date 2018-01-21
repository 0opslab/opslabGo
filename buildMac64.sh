export CGO_ENABLED=0
export GOOS=darwin
export GOARCH=amd64
go build -o build/weixin src/base/web/03web/main.go 