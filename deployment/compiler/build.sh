#!/bin/bash

# 启动postgres容器，用于生成model
function run_pg() {
    docker run -d \
               --name model-gen-pg \
               -p 5432:5432 \
               fms-postgres:latest
    sleep 1
}

function docker_make() {
    local make_cmd=$1
    docker run -it \
               --rm \
               --network host \
               -v /mnt/d/projects/fms/:/root/fms/ \
               fms-compiler:latest \
               /bin/bash -c "cd /root/fms/service/fms && make ${make_cmd}"
}

function clean () {
    # 构建后清理一些中间数据，如容器
    echo "Starting to delete pg container"

    docker stop model-gen-pg &> /dev/null
    docker rm model-gen-pg &> /dev/null

    echo "delete pg container success"
}

function main() {
    local command=$1
    case "${command}" in
        codegen)
            echo "Starting to generate code"
            docker_make codegen
            ;;
        modelgen)
            echo "Starting to generate model"
            run_pg
            docker_make modelgen
            ;;
        all)
            echo "Starting to build all"
            run_pg
            docker_make all
            ;;
        clean)
            echo "Starting to clean"
            clean
            return 0
            ;;
        *)
            echo "Unsupport command"
            return 1
            ;;
    esac

    clean
}

main "$@"