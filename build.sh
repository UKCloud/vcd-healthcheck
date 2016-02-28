#!/bin/bash

set -ev

OS="linux windows"
ARCH="amd64"

echo "Getting build dependencies"
go get . 
go get -u github.com/golang/lint/golint

echo "Ensuring code quality"
go vet ./...
golint ./...
go test -v -check.v ./...
go build


if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
	# go get github.com/laher/goxc
	# goxc -t 
	# goxc bump
	# goxc
	for GOOS in $OS; do
	    for GOARCH in $ARCH; do
	        arch="$GOOS-$GOARCH"
	        binary="bin/vcd-healthcheck.$arch"
	        echo "Building $binary"
	        GOOS=$GOOS GOARCH=$GOARCH go build -o $binary
	    done
	done
fi
pwd
set
ls -lRa