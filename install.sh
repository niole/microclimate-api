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

echo "Cancel if tests do not succeed"

cd grpc-java/compiler
../gradlew java_pluginExecutable -PskipAndroid=true
../gradlew test -PskipAndroid=true

echo "Install go stuff"

go install all
