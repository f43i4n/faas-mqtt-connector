#!/bin/bash

TAG=${TAG:-latest}
NAMESPACE=${NAMESPACE:-f43i4n}
IMAGE=${IMAGE:-openfaas-mqtt-connector}

PLATFORMS=${PLATFORMS:-linux/amd64,linux/arm/v7}

# use docker buildkit with experimental features to target multiple architectures
# requires buildkit configuration that is able to build the provided platforms
docker buildx build \
    -t "${NAMESPACE}/${IMAGE}:${TAG}" \
    --platform "${PLATFORMS}" \
    --push \
    .
