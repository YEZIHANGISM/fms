#!/bin/bash

set -x

docker compose down

docker compose build

docker compose up -d