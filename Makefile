SHELL := /bin/bash
BASEDIR = $(shell pwd)
APP = garage
BuildDIR = build

gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0 | sed 's/v//g'; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
versionDir = "garage/command"
ldflags= "-X ${versionDir}.gitTag=${gitTag} \
-X ${versionDir}.buildDate=${buildDate} \
-X ${versionDir}.gitCommit=${gitCommit} \
-X ${versionDir}.gitTreeState=${gitTreeState}"

release:
	# Build for linux
	go clean
	GOOS=linux GOARCH=amd64 go build -v -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-amd64 ./bin/
	GOOS=linux GOARCH=arm64 go build -v -ldflags ${ldflags} -o ${BuildDIR}/${APP}-linux64-arm64 ./bin/
	# Build for win
	go clean
	GOOS=windows GOARCH=amd64 go build -v -ldflags ${ldflags} -o ${BuildDIR}/${APP}-windows-amd64.exe ./bin/
	# Build for mac
	go clean
	GOOS=darwin GOARCH=amd64 go build -v -ldflags ${ldflags} -o ${BuildDIR}/${APP}-darwin-amd64 ./bin/
	GOOS=darwin GOARCH=arm64 go build -v -ldflags ${ldflags} -o ${BuildDIR}/${APP}-darwin-arm64 ./bin/
	# list
	ls -al ${BuildDIR}

clean:
	@go clean --cache
	@rm -rvf build/*

.PHONY: release clean
