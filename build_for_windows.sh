#!/bin/bash

clear
FILEMAIN=./internal/main.go
FILEAPP=./bin/whatsapp_chatgpt
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc
go build -o $FILEAPP.exe $FILEMAIN

