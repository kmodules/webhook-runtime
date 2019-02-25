#!/bin/bash

set -x

GOPATH=$(go env GOPATH)
PACKAGE_NAME=kmodules.xyz/webhook-runtime
REPO_ROOT="$GOPATH/src/$PACKAGE_NAME"
DOCKER_REPO_ROOT="/go/src/$PACKAGE_NAME"

pushd $REPO_ROOT

# Generate deep copies
docker run --rm -ti -u $(id -u):$(id -g) \
    -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
    -w "$DOCKER_REPO_ROOT" \
    appscode/gengo:release-1.13 deepcopy-gen \
    --go-header-file "hack/gengo/boilerplate.go.txt" \
    --input-dirs "$PACKAGE_NAME/apis/workload/v1" \
    --output-file-base zz_generated.deepcopy

popd
