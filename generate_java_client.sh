#!/bin/bash

export SRC_DIR=./interface DST_DIR=./java_api GO111MODULE=on

mkdir -p $DST_DIR

protoc -I=$SRC_DIR \
--plugin=protoc-gen-grpc-java=grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
--java_out=$DST_DIR \
--grpc-java_out=$DST_DIR \
$SRC_DIR/*.proto
