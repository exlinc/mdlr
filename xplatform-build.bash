#!/usr/bin/env bash
rm -rf ./build
GOOS=darwin
GOARCH=amd64
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=darwin
GOARCH=386
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=linux
GOARCH=amd64
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=linux
GOARCH=386
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=windows
GOARCH=amd64
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr.exe)
GOOS=windows
GOARCH=386
go build -o $(echo ./build/$GOOS-$GOARCH-mdlr.exe)
