#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


chmod +x ../vendor/k8s.io/code-generator/generate-groups.sh
../vendor/k8s.io/code-generator/generate-groups.sh \
  all github.com/edgelive/pkg   github.com/edgelive/pkg/apis \
  "rule:v1 service:v1 heartbeat:v1" \
  --go-header-file $(pwd)/boilerplate.go.txt \
#  --output-base /Users/zh/go/src