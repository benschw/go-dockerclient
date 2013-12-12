#!/bin/sh
set -e

export GOPATH="${PWD}/build"

. ./build.sh

go test -i ./dockerapi
go test -v ./dockerapi

