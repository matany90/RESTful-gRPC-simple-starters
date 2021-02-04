#!/bin/bash

echo "compiling matan-service server gRPC instance.\n"

# compile matan-service.proto to compiled files
python -m \
grpc_tools.protoc -I./proto \
--python_out=./proto \
--grpc_python_out=./proto \
./proto/matan-service.proto

sed -i.bak '5s/.*/import proto.matan_service_pb2 as matan__service__pb2/' ./proto/matan_service_pb2_grpc.py