#!/bin/bash
set -e

. /opt/dockerapi/integration_tests/wrapdocker 

cd /opt

export GOPATH="${PWD}/build"


sleep 2

docker pull benschw/etcd


echo building...
. ./clean.sh
. ./build.sh


echo 
echo Unit Tests:
echo 
./test_unit.sh

echo 
echo Integration Tests:
echo
./test_integration.sh

