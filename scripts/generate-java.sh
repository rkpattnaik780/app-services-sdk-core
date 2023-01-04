#!/usr/bin/env bash
## OPENAPI_FILENAME=yourapi generate_js.sh 

# if input file is a empty string that means this is a workflow run
# manualy and we want to generate all sdks 
INPUT_FILE=`basename $1`

# if the input file is openapi.yaml that means it should be 
# service accounts, change the INPUT_FILE and make the service
# accounts file up to date with the new changes
if [ "$INPUT_FILE" = "openapi.yaml" ];
then
    INPUT_FILE="service-accounts.yaml"
fi

echo "========================="
echo "Input file is $INPUT_FILE"
echo "========================="

set -e

npx @openapitools/openapi-generator-cli version-manager set 5.4.0
echo "Generating SDKs"

TEMPLATES_DIR="$(dirname $0)/templates"

if [ "$INPUT_FILE" = "kas-fleet-manager.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="kafka-management-sdk"
    OPENAPI_FILENAME=".openapi/kas-fleet-manager.yaml"
    PACKAGE_NAME="com.openshift.cloud.api.kas"
    OUTPUT_PATH="app-services-sdk-java/packages/kafka-management-sdk"

    echo "Generating based on ${OPENAPI_FILENAME}"
    yq e 'del(.. | select(has("deprecated")))' "${OPENAPI_FILENAME}" > "${OPENAPI_FILENAME}.processed"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    npx @openapitools/openapi-generator-cli generate -g java --library resteasy -t "$TEMPLATES_DIR"  -i \
        "$OPENAPI_FILENAME.processed" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "srs-fleet-manager.json" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="registry-management-sdk"
    OPENAPI_FILENAME=".openapi/srs-fleet-manager.json"
    PACKAGE_NAME="com.openshift.cloud.api.srs"
    OUTPUT_PATH="app-services-sdk-java/packages/registry-management-sdk/"

    echo "Generating based on ${OPENAPI_FILENAME}"

    yq e 'del(.. | select(has("deprecated")))' "${OPENAPI_FILENAME}" > "${OPENAPI_FILENAME}.processed"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    npx @openapitools/openapi-generator-cli generate -g java --library resteasy -t "$TEMPLATES_DIR"  -i \
        "$OPENAPI_FILENAME.processed" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "kafka-admin-rest.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="kafka-instance-sdk"
    OPENAPI_FILENAME=".openapi/kafka-admin-rest.yaml"
    PACKAGE_NAME="com.openshift.cloud.api.kas.auth"
    OUTPUT_PATH="app-services-sdk-java/packages/kafka-instance-sdk/"
    yq e 'del(.. | select(has("deprecated")))' "${OPENAPI_FILENAME}" > "${OPENAPI_FILENAME}.processed"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    echo "Generating based on ${OPENAPI_FILENAME}"

    npx @openapitools/openapi-generator-cli generate -g java --library resteasy  -t "$TEMPLATES_DIR"  -i \
        "$OPENAPI_FILENAME.processed" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="openApiNullable=false,apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "connector_mgmt.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="connector-management-sdk"
    OPENAPI_FILENAME=".openapi/connector_mgmt.yaml"
    PACKAGE_NAME="com.openshift.cloud.api.connector"
    OUTPUT_PATH="app-services-sdk-java/packages/connector-management-sdk/"
    yq e 'del(.. | select(has("deprecated")))' "${OPENAPI_FILENAME}" > "${OPENAPI_FILENAME}.processed"
    yq e 'del(.. | select(has("deprecated")))' "${OPENAPI_FILENAME}" > "${OPENAPI_FILENAME}.processed"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    echo "Generating based on ${OPENAPI_FILENAME}"

    npx @openapitools/openapi-generator-cli generate -g java --library resteasy  -t "$TEMPLATES_DIR"  -i \
        "$OPENAPI_FILENAME.processed" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "service-accounts.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="service-accounts-sdk"
    OPENAPI_FILENAME=".openapi/service-accounts.yaml"
    PACKAGE_NAME="com.openshift.cloud.api.serviceaccounts"
    OUTPUT_PATH="app-services-sdk-java/packages/service-accounts-sdk/"

    echo "Generating based on ${OPENAPI_FILENAME}"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    npx @openapitools/openapi-generator-cli generate -g java \
        --library resteasy -t "$TEMPLATES_DIR" \
        -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "smartevents_mgmt_v2.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    ARTIFACT_ID="smartevents-management-sdk"
    OPENAPI_FILENAME=".openapi/smartevents_mgmt_v2.yaml"
    PACKAGE_NAME="com.openshift.cloud.api.smartevents"
    OUTPUT_PATH="app-services-sdk-java/packages/smartevents-management-sdk/"

    echo "Generating based on ${OPENAPI_FILENAME}"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    npx @openapitools/openapi-generator-cli generate -g java \
        --library resteasy -t "$TEMPLATES_DIR" \
        -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

if [ "$INPUT_FILE" = "registry-instance.json" ] || [ "$INPUT_FILE" = "" ];
then
    GROUP_ID="com.redhat.cloud"
    OPENAPI_FILENAME=".openapi/registry-instance.json"
    ARTIFACT_ID="registry-instance-sdk"
    PACKAGE_NAME="com.openshift.cloud.api.registry.instance"
    OUTPUT_PATH="app-services-sdk-java/packages/registry-instance-sdk/"

    echo "Generating based on ${OPENAPI_FILENAME}"
    rm -Rf $OUTPUT_PATH/src $OUTPUT_PATH/target

    npx @openapitools/openapi-generator-cli generate -g java \
        --library resteasy -t "$TEMPLATES_DIR" \
        -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties="apiTests=false,modelTests=false,hideGenerationTimestamp=true,groupId=${GROUP_ID},artifactId=${ARTIFACT_ID},modelPackage=${PACKAGE_NAME}.models,invokerPackage=${PACKAGE_NAME}.invoker,apiPackage=${PACKAGE_NAME},dateLibrary=java8,licenseName=Apache-2.0,licenseUrl=https://www.apache.org/licenses/LICENSE-2.0.txt" \
        --ignore-file-override=.openapi-generator-ignore
fi

rm .openapi/*.processed