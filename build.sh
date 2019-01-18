#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w"
upx searchcode-server-highlighter.exe
mv searchcode-server-highlighter.exe searchcode-server-highlighter-x86_64-pc-windows.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
upx searchcode-server-highlighter
mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-unknown-linux
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w"
upx searchcode-server-highlighter
mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-apple-darwin