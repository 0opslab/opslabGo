set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o build/rysnc.exe src/opslabgo/uploadRsync2/main.go