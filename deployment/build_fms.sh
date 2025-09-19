#!/bin/bash

# 获取当前文件的路径
CURDIR=$(dirname "$0")
CURDIR=$(cd "$CURDIR" && pwd)

PROJECT_DIR=$CURDIR/../

cd $PROJECT_DIR

docker build -t fms .

cd $CURDIR
helm uninstall fms
helm upgrade --install fms ./fms
