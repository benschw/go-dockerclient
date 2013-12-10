#!/bin/sh
set -e


. ./build


go test -i ./dockerapi
go test -v ./dockerapi

