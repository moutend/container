#!/bin/bash

set -e

docker build -t demo1:hello-secret-env-v2 .
docker image prune --filter label=stage=temporary
