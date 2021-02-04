#!/bin/bash

HOST="127.0.0.1"
PORT="50006"

GRPC_ROOT_PATH="$HOME/Desktop/Comp/RESTful-gRPC-simple-starters/python/grpc"

grpcurl \
  -plaintext \
  -import-path $GRPC_ROOT_PATH/proto \
  -proto matan-service.proto \
  -d '{"name": "Matan"}' \
  $HOST:$PORT \
  matanservice.MatanService.SayHello
