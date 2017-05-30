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

# Put the binary hello-go filename in /dist
go build -o dist/goflake ./main.go

# cp the Dockerfile into /dist
cp ci/Dockerfile dist/Dockerfile

# Check
echo "List whats in the /dist directory"
ls -lat dist
echo ""

# Move what you need to $GOPATH/dist
# BECAUSE the resource type docker-image works in /dist.
cp -R ./dist $GOPATH/.

cd $GOPATH
# Check whats here
echo "List whats in top directory"
ls -lat 
echo ""

# Check whats in /dist
echo "List whats in /dist"
ls -lat dist
echo ""