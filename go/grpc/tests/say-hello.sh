#!/bin/bash

HOST="127.0.0.1"
PORT="50004"

GRPC_ROOT_PATH="$HOME/Desktop/Comp/RESTful-gRPC-simple-starters/go/grpc"

grpcurl \
  -plaintext \
  -import-path $GRPC_ROOT_PATH/proto \
  -proto matan-service.proto \
  -d '{"name": "Matan"}' \
  $HOST:$PORT \
  matanservice.MatanService.SayHello
