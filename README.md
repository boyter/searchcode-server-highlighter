A very simple Go HTTP based Syntax highlighter. Run it, then post some code to the default port and it will return 
CSS + HTML syntax highlighted code.

Build step to create smallest builds

```
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w"
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w"
```

Then apply upx over the resulting files

```
upx searchcode-server-highlighter
upx searchcode-server-highlighter.exe
```

The resulting build for each file should be ~3 MB