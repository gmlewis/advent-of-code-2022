#!/bin/bash -ex
go run main.go ../example1.txt
if [ -f ../input.txt ]; then
  go run main.go ../input.txt
fi
