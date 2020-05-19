#!/bin/bash

kubectl create secret generic cloudsql-instance-credentials --from-file=/tmp/credentials.json
