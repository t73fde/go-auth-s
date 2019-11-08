#!/bin/sh

export GOARCH=amd64
export LDFLAGS="-s -w"
TAG_VERSION=$(git tag | sort | tail -n 1)
for GOOS in "linux" "darwin" "windows"
do
  EXECUTABLE="go-auth-s-$TAG_VERSION-$GOOS-$GOARCH"
  GOOS="$GOOS" go build -o "$EXECUTABLE" -ldflags "$LDFLAGS"
  if [ "$GOOS" = "windows" ]; then
    mv "$EXECUTABLE" "$EXECUTABLE.exe"
    rm -f "$EXECUTABLE.zip"
    zip -9q "$EXECUTABLE.zip" "$EXECUTABLE.exe"
    rm "$EXECUTABLE.exe"
  else
    rm -f "$EXECUTABLE.bz2"
    bzip2 -9 "$EXECUTABLE"
  fi
done
