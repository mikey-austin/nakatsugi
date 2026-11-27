#!/bin/sh
#
# This relies on the protoc compiler being installed and the golang
# generation backend, which can be installed with (installs to $GOPATH/bin
# or $GOBIN if set):
#
#   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#
protoc -I=/usr/include/ -I=./ --go_out=../ ./messages.proto
