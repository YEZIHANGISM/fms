#!/bin/bash

docker rmi fms:latest
docker build -t fms .
