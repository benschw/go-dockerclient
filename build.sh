#!/bin/sh
set -e

PACKAGE=github.com/benschw/go-dockerclient/dockerapi

export GOPATH="${PWD}/build"

mkdir -p $GOPATH
SRC_DIR="$GOPATH/src"
PACKAGE_DIR="$SRC_DIR/$PACKAGE"

PACKAGE_BASE=$(dirname "${PACKAGE_DIR}")
mkdir -p  "${PACKAGE_BASE}"

ln -sf ${PWD}/dockerapi "${PACKAGE_DIR}"


#./scripts/release-version > server/release_version.go
go build "${PACKAGE}"
