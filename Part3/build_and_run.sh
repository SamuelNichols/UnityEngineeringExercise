#!/bin/bash

# Building Webservice Dockerfile
pushd WebService
docker build -t webservice-container .
popd
# Starting kube script
pushd Kubernetes
./kube-manage.sh create
popd
