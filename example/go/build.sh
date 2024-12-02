#!/bin/bash

# 设置环境变量
export CGO_ENABLED=1
export GOARCH=arm64
export GOOS=darwin

# 进入go目录
cd "$(dirname "$0")"

# 编译共享库
echo "Building shared library..."
go build -buildmode=c-shared -o bridge.so ./cmd/main.go 