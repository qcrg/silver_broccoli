#!/bin/bash
set -e
cd "$(dirname "$0")"

  # -v ./init:/docker-entrypoint-initdb.d \
docker run \
  --rm \
  --name silver_broccoli_postgres \
  -e POSTGRES_USER=guest \
  -e POSTGRES_PASSWORD='asdf;lkj' \
  -e POSTGRES_DB=test \
  -v /home/qcrg/develop/silver_broccoli/testdata/dev/postgres/init:/docker-entrypoint-initdb.d \
  -v ./data:/csv_import \
  -p 5432:5432 \
  postgres:latest
