########################################################################################################################
# Constants
########################################################################################################################


# Golang compilation options
export GOOS=darwin
export GOARCH=amd64

PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

VERSION=1.0.0
BUILD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Main gopath
export GOPATH=/Users/mstepan/repo/go-workspace

# go version go1.13.7 darwin/amd64 location
export GOROOT=/usr/local/go

########################################################################################################################
# Commands
########################################################################################################################
.PHONY: build build_all clean default install run vendor

default: build

build: clean vendor
	go build ${LDFLAGS} -o bin/main src/main.go
	go build ${LDFLAGS} -o bin/copy_files src/copy_files.go
	go build ${LDFLAGS} -o bin/folder_stat src/folder_stat.go

# Build executables for all platforms and architectures.
build_all: clean vendor
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/main-$(GOOS)-$(GOARCH) src/main.go)))

vendor:
	go mod vendor

run:
	./bin/main

install:
	go install ${LDFLAGS} src/main.go
	go install ${LDFLAGS} src/folder_stat.go

clean:
	go clean
	rm -rf bin

