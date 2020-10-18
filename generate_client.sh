#!/bin/bash

export SRC_DIR=./src/server/interface DST_DIR=. GO111MODULE=on

protoc -I=$SRC_DIR --go_out=$DST_DIR --go-grpc_out=$DST_DIR $SRC_DIR/*.proto
