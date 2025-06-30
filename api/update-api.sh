#!/bin/bash

cd "$(dirname $0)"

sed -n -i -re '0,/\/\/ ###/p' api.go
echo >> api.go
go run ./gen >> api.go
gofmt -w .
