#!/bin/bash

export GOOS=linux
# 目标操作系统架构
export GOARCH=amd64
FUNCTION_DIR=serverless

cd $FUNCTION_DIR

rm -rf main.zip
go clean
go build -o main main.go
