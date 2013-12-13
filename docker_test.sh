#!/bin/bash


if [ ! -f ./docker-testing-container/docker-latest ]; then
	wget -O ./docker-testing-container/docker-latest https://get.docker.io/builds/Linux/x86_64/docker-latest
fi

ln -sf `pwd` ./docker-testing-container/wc

docker build -t dind ./docker-testing-container/
docker run -privileged -t -i dind