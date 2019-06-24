A very simple Go HTTP based Syntax highlighter. Run it, then post some code to the default port and it will return 
CSS + HTML syntax highlighted code.

[![Build Status](https://travis-ci.org/boyter/searchcode-server-highlighter.svg?branch=master)](https://travis-ci.org/boyter/searchcode-server-highlighter)
[![Scc Count Badge](https://sloc.xyz/github/boyter/searchcode-server-highlighter/)](https://github.com/boyter/searchcode-server-highlighter/)

Build step to create smallest builds

```
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter.exe && mv searchcode-server-highlighter.exe searchcode-server-highlighter-x86_64-pc-windows.exe
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter && mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-unknown-linux
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" && upx searchcode-server-highlighter && mv searchcode-server-highlighter searchcode-server-highlighter-x86_64-apple-darwin
```

The resulting build for each file should be ~3 MB

## Sample Usage

Install [httpie](https://httpie.org/) and [jq](https://stedolan.github.io/jq/), then run the following snippet to generate a local HTML file.

```shell
FILE=$GOPATH/src/github.com/boyter/searchcode-server-highlighter/main.go
LANG=go
STYLE=tango

# send up the file
http POST localhost:8089/v1/highlight/ languageName=$LANG fileName=$(basename $FILE) style=$STYLE content=@$FILE > res

# munch results
CSS=$(cat res| jq --raw-output .css  | sed 's/\\n/\n/g')
HTML=$(cat res| jq --raw-output .html | sed 's/\\n/\n/g')

cat  << EOF > $(basename $FILE).html
<!DOCTYPE html>
<html>
<head>
<style>
$CSS
</style>
</head>
<body>
$HTML
</body>
</html>
EOF

echo "Done: $(basename $FILE).html"
```
