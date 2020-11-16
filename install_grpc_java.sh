#!/bin/bash

export ANDROID_SDK_ROOT="$HOME/Library/Android/sdk"

echo "Installing grpc java compiler. Cancel if tests do not succeed"

cd grpc-java/compiler
../gradlew java_pluginExecutable
../gradlew test

