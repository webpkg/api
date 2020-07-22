#!/bin/bash
set -e

# gofmt -w .

make clean

make

docker-compose build
docker-compose up -d