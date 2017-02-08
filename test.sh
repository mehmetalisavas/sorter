#!/bin/bash

set -o errexit

# ./test.sh runs the tests with name of the tests (use regular expressions)
go test -race -v -test.run="$1" .
