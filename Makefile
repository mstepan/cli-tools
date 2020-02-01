########################################################################################################################
# Constants
########################################################################################################################

# Golang compilation options
export GOOS=darwin
export GOARCH=amd64

# Main gopath
export GOPATH=/Users/mstepan/repo/go-workspace

# go version go1.13.7 darwin/amd64 location
export GOROOT=/usr/local/go

########################################################################################################################
# Commands
########################################################################################################################
.PHONY: default build clean

default: build

build:
	go mod vendor
	go build -o bin/main src/main.go
	go build -o bin/folder_stat src/folder_stat.go

run:
	./bin/main

clean:
	rm -rf bin

