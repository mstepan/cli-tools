########################################################################################################################
# Constants
########################################################################################################################

# Golang compilation
GOOS=darwin
GOARCH=amd64
GOPATH=/Users/mstepan/repo/go-workspace
GOROOT=/usr/local/bin/go

# Get the version number from VERSION file.
CODE_VERSION = $(strip $(shell cat VERSION))

ifndef CODE_VERSION
$(error You need to create a VERSION file to build a release)
endif

# Get the latest commit.
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

########################################################################################################################
# Commands
########################################################################################################################
.PHONY: default build clean

default: build

build:
	go build -o bin/main main.go
	go build -o bin/folder_stat folder_stat.go

run:
	./bin/main

clean:
	rm -rf main

