SHELL := /bin/bash
BuildDir = build
APP = garage

release:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o ${BuildDir}/${APP} ./bin/