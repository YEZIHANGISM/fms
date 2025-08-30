#!/bin/bash

# 从secret中获取密码，设置环境变量
export POSTGRES_PASSWORD=$(kubectl get secret --namespace default fms-postgres-postgresql -o jsonpath="{.data.postgres-password}" | base64 --decode)
kubectl run my-postgres-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:latest --env="PGPASSWORD=$POSTGRES_PASSWORD" --command -- psql --host fms-postgres-postgresql -U postgres -d postgres -p 5432