#!/bin/bash


if [ ! -f ./dockerapi/integration_tests/docker-latest ]; then
	wget -O ./dockerapi/integration_tests/docker-latest https://get.docker.io/builds/Linux/x86_64/docker-latest
fi

docker build -t dind ./dockerapi/integration_tests/
docker run -privileged -t -i -v `pwd`:/opt dind