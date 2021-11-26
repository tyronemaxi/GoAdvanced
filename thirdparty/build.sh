#!/bin/bash
set -ex

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

#$GO build -o model-unpack-tool-linux main.go util.go
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $GO build -o model-unpack-tool-mac main.go util.go
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $GO build -o model-unpack-tool-win.exe main.go util.go