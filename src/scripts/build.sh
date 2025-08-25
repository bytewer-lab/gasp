#!/bin/bash

set -e
set -x

echo "Building for linux (amd64)"
env GOOS=linux GOARCH=amd64 go build -o ./build/gasp-linux-amd64

echo "Building for darwin (amd64)"
env GOOS=darwin GOARCH=amd64 go build -o ./build/gasp-darwin-amd64

echo "Building for windows (amd64)"
env GOOS=windows GOARCH=amd64 go build -o ./build/gasp-windows-amd64.exe
