#!/bin/bash

if [ $# -lt 1 ]; then
    echo "usage: $0 proto_files"
    exit 1
fi

protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. $@
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. $@
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. $@
