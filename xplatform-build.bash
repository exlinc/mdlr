#!/usr/bin/env bash
if [ ! -d "./build" ]; then
  echo "Use mdlr to import the ./build directory with mdlr import"
  exit 1
fi

GOOS=darwin
GOARCH=amd64
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=darwin
GOARCH=386
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=linux
GOARCH=amd64
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=linux
GOARCH=386
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr)
GOOS=windows
GOARCH=amd64
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr.exe)
GOOS=windows
GOARCH=386
GOOS=$GOOS GOARCH=$GOARCH go build -o $(echo ./build/$GOOS-$GOARCH-mdlr.exe)
