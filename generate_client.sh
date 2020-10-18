#!/bin/bash

export SRC_DIR=./src/server/interface DST_DIR=./src/server/interface GO111MODULE=on

protoc -I=$SRC_DIR --go_out=$DST_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$DST_DIR --go-grpc_opt=paths=source_relative \
    $SRC_DIR/*.proto
