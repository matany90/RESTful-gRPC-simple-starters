#!/bin/bash

echo "compiling matan-service server gRPC instance.\n"

PATH="$HOME/go/bin:$PATH"; export PATH;
# GOPATH="$HOME/devel/googleapis:$HOME/go:$GOPATH"; export GOPATH;

mkdir -p ./pkg/matan-service

protoc \
  -I ./proto \
  -I $GOPATH \
  -I $GOPATH/src/github.com/envoyproxy/protoc-gen-validate \
  --go_out ./pkg/matan-service \
  --go_opt paths=source_relative \
  --go-grpc_out ./pkg/matan-service \
  --go-grpc_opt paths=source_relative \
  --validate_out=lang=go,paths=source_relative:./pkg/matan-service \
  "matan-service.proto"
