#!/bin/sh
set -e

export GOPATH="${PWD}/build"

. ./build.sh

go test -i ./dockerapi/integration_tests
go test -v ./dockerapi/integration_tests

