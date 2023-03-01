#!/usr/bin/env bash
set -euxo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

endpoints=(
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml"
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/connector_mgmt.yaml"
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/srs-fleet-manager/main/core/src/main/resources/srs-fleet-manager.json"
    "https://raw.githubusercontent.com/5733d9e2be6485d52ffa08870cabdee0/sandbox/main/openapi/smartevents_mgmt_v2.yaml"
    "https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kafka-admin-api/main/kafka-admin/.openapi/kafka-admin-rest.yaml"
    "https://raw.githubusercontent.com/Apicurio/apicurio-registry/main/app/src/main/resources-unfiltered/META-INF/resources/api-specifications/registry/v2/openapi-gen.json"
    "https://sso.redhat.com/auth/realms/redhat-external/apis/openapi.yaml"
)

for endpoint in ${endpoints[*]}; do
    filename=$(echo "${endpoint##*/}")

    if [[ ${filename} == "openapi.yaml" ]]; then
        filename="service-accounts.yaml"
    fi
    if [[ ${filename} == "openapi-gen.json" ]]; then
        filename="registry-instance.json"
    fi

    output_filename="${SCRIPT_DIR}/../.openapi/${filename}"

    curl -sL $endpoint --output $output_filename
done
