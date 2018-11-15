#!/bin/sh
GOARCH=wasm GOOS=js go build -o main.wasm main.go

go build -o server server.go

./server
