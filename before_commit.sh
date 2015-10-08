#!/bin/bash

echo "go fmt"
find . -name "*.go" -exec go fmt {} \;

echo "golint"
golint
# golint palintodo

echo "gom test"
go test
