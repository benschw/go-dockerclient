#!/bin/sh
set -e


sudo ./clean.sh
. ./build.sh

echo 
echo Unit Tests:
echo 
. ./test_unit.sh

echo 
echo Integration Tests:
echo
sudo ./test_integration.sh


