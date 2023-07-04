SHELL := /bin/bash
BASEDIR = $(shell pwd)
APP = garage
BuildDIR = build

gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0 | sed 's/v//g'; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
versionDir = "github.com/gsxhnd/garage"
ldflags= "-s -w -X ${versionDir}.gitTag=${gitTag} \
-X ${versionDir}.buildDate=${buildDate} \
-X ${versionDir}.gitCommit=${gitCommit} \
-X ${versionDir}.gitTreeState=${gitTreeState}"

all:
	# Build for linux
	go clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath  -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-amd64 ./src
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-arm64 ./src
	# Build for win
	go clean
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-windows-amd64.exe ./src
	CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-windows-arm.exe ./src
	# Build for mac
	go clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-darwin-amd64 ./src
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -trimpath -ldflags ${ldflags} -o ${BuildDIR}/${APP}-darwin-arm64 ./src

test:
	goreleaser release --clean --skip-publish --skip-validate

clean:
	@go clean --cache
	@rm -rvf build/*
	@rm -rvf testdata/
	@rm -rvf javbus/

mock_data:
	rm -rvf testdata
	mkdir testdata
	touch ./testdata/{1,2,3,4,5}.{mkv,mp4,ass,ttf}
	ls -al ./testdata

.PHONY: release_linux release_win release_mac clean