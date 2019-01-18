A very simple Go HTTP based Syntax highlighter. Run it, then post some code to the default port and it will return 
CSS + HTML syntax highlighted code.

[![Build Status](https://travis-ci.org/boyter/searchcode-server-highlighter.svg?branch=master)](https://travis-ci.org/boyter/searchcode-server-highlighter)

Build step to create smallest builds

```
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter.exe && mv searchcode-server-highlighter.exe searchcode-server-highlighter-x86_64-pc-windows.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter && mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-unknown-linux
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter && mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-apple-darwin
```

Then apply upx over the resulting files

```
upx searchcode-server-highlighter
upx searchcode-server-highlighter.exe
```

The resulting build for each file should be ~3 MB