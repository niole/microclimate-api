#!/bin/bash

export SRC_DIR=./interface DST_DIR=./python_api GO111MODULE=on PLUGIN_PATH=

mkdir -p $DST_DIR

python -m grpc_tools.protoc \
-I$SRC_DIR \
--python-grpc_opt=paths=source_relative \
--python_opt=paths=source_relative \
--python_out=$DST_DIR \
--grpc_python_out=$DST_DIR  \
$SRC_DIR/*.proto


