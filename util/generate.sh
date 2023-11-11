#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


chmod +x ../vendor/k8s.io/code-generator/generate-groups.sh
../vendor/k8s.io/code-generator/generate-groups.sh \
  all github.com/edgelive/def   github.com/edgelive/def/apis \
  "recorder:v1 heartbeat:v1" \
  --go-header-file $(pwd)/boilerplate.go.txt \
#  --output-base /Users/zh/go/src