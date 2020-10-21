#!/bin/bash

export SRC_DIR=./src/server/interface DST_DIR=./protobuf GO111MODULE=on

mkdir -p $DST_DIR

# is easiest to make it so that the go_package option path matches
# the path to the generated files
# use paths=source_relative for the grpc services and message definitions
# so that all generated files end up in api/protobuf and can import them via
# api/protobuf

protoc -I=$SRC_DIR \
--go-grpc_opt=paths=source_relative \
--go_opt=paths=source_relative \
--go_out=$DST_DIR \
--go-grpc_out=$DST_DIR  \
$SRC_DIR/*.proto
