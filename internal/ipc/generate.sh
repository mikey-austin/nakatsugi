#!/bin/sh
#
# This relies on the protoc compiler being installed and the golang
# generation backend, which can be installed with (installs to $GOPATH/bin
# or $GOBIN if set):
#
#   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#
INC=
if [ -d /usr/include/protobuf/ ]; then
    INC="${INC} -I/usr/include/"
elif [ -d /opt/homebrew/opt/protobuf/ ]; then
    INC="${INC} -I/opt/homebrew/opt/protobuf/include/"
fi
protoc $INC -I=./ --go_out=../ ./messages.proto
