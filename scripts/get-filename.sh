#!/usr/bin/env bash
set -euo pipefail

ENDPOINT=${1}

filename=$(echo "${ENDPOINT##*/}")

if [[ ${filename} == "openapi.yaml" ]]; then
    filename="service-accounts.yaml"
fi
if [[ ${filename} == "openapi" ]]; then
    filename="account-management.json"
fi
if [[ ${filename} == "openapi-gen.json" ]]; then
    filename="registry-instance.json"
fi

echo ${filename}
