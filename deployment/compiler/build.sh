#!/bin/bash

# 启动postgres容器，用于生成model
docker run -d \
           --name model-gen-pg \
           -p 5432:5432 \
           fms-postgres:latest

docker run -it \
           --rm \
           --network host \
           -v /mnt/d/projects/fms/:/root/fms/ \
           fms-compiler:latest \
           /bin/bash -c "cd /root/fms/service/fms && make"

docker stop model-gen-pg:latest
docker rm model-gen-pg:latest