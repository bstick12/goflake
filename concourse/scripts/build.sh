#!/bin/bash

set -e -x

echo "List whats in the current directory"
ls -lat 
echo ""

export GOPATH=$PWD

mkdir -p src/github.com/bstick12/
cp -R ./goflake src/github.com/bstick12/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
echo ""
cd src/github.com/bstick12/goflake

go get -t -d -v ./...

# Put the binary hello-go filename in /dist
go build -v

ls -lat