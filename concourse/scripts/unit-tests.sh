#!/bin/bash

set -e -x

echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

export GOPATH=$PWD

mkdir -p src/github.com/bstick12/
cp -R ./goflake src/github.com/bstick12/.

echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/bstick12/goflake
ls -lat

go get -t -d -v ./...

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

mv coverage.txt $GOPATH/coverage-results/.


