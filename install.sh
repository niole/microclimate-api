#!/bin/bash

export GO111MODULE=on

if [ ! -d "grpc-java" ]
then
    echo "Cloning grpc-java"
    git clone git@github.com:grpc/grpc-java.git
    git checkout a43ae54c5
fi
if [ ! -d "grpc-java" ]
then
    echo "Failed to clone grpc-java"
    exit 1
fi

echo "Install java stuff"

./client-generators/java/install_grpc_java.sh

echo "Install go stuff"

go install all

echo "Install python stuff"

cd ./client-generators/python && ./env_setup.sh && pip install -r requirements.txt
