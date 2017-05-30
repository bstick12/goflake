#!/bin/bash
# goflake unit-test.sh

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

go test -cover ./... | tee test_coverage.txt
sed -i -e 's/^/     /' test_coverage.txt

mv test_coverage.txt $GOPATH/coverage-results/.


