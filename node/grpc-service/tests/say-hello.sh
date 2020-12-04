#!/bin/bash

HOST="127.0.0.1"
PORT="50001"

GRPC_ROOT_PATH="$HOME/Desktop/Comp/RESTful-gRPC-simple-starters/node/grpc-service"

grpcurl \
  -plaintext \
  -import-path $GRPC_ROOT_PATH/proto \
  -d '{"name": "Matan"}' \
  -proto matan.proto \
  $HOST:$PORT \
  matan.MatanService.SayHello
