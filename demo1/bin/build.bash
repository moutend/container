#!/bin/bash

set -e

docker build -t demo1:hello-secret-env .
docker image prune --filter label=stage=temporary
