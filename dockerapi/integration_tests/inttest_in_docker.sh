#!/bin/bash
set -e

. /opt/dockerapi/integration_tests/wrapdocker 

cd /opt

export GOPATH="${PWD}/build"


sleep 2



echo building...
. ./clean.sh
. ./build.sh


docker pull benschw/etcd


echo 
echo Unit Tests:
echo 
./test_unit.sh

echo 
echo Integration Tests:
echo
./test_integration.sh

