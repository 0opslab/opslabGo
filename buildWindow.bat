set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o build/cmdHpler.exe cmdhepler/main.go
go build -o build/httpcmd.exe httpcmd/main.go