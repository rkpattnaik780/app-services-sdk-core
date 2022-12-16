#!/usr/bin/env bash
## OPENAPI_FILENAME=yourapi generate.sh 

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

# generate an API client for a service
generate_sdk() {
    local file_name=$1
    local output_path=$2
    local package_name=$3
     
    echo "Validating OpenAPI ${file_name}"
    npx @openapitools/openapi-generator-cli validate -i "$file_name"

    echo "Generating source code based on ${file_name}"

    # remove old generated models
    rm -Rf $OUTPUT_PATH
    
    npx @openapitools/openapi-generator-cli generate -g go -i \
    "$file_name" -o "$output_path" \
    --package-name="${package_name}" \
    --additional-properties=$additional_properties \
    --ignore-file-override=.openapi-generator-ignore
}

npx @openapitools/openapi-generator-cli version-manager set 6.2.1
echo "Generating Go SDKs"
additional_properties="generateInterfaces=true,enumClassPrefix=true"

if [ "$INPUT_FILE" = "kas-fleet-manager.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/kas-fleet-manager.yaml"
    PACKAGE_NAME="kafkamgmtclient"
    OUTPUT_PATH="app-services-sdk-go/kafkamgmt/apiv1/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

if [ "$INPUT_FILE" = "srs-fleet-manager.json" ] || [ "$INPUT_FILE" = "" ];
then
OPENAPI_FILENAME=".openapi/srs-fleet-manager.json"
PACKAGE_NAME="registrymgmtclient"
OUTPUT_PATH="app-services-sdk-go/registrymgmt/apiv1/client"

generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

if [ "$INPUT_FILE" = "connector_mgmt.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/connector_mgmt.yaml"
    PACKAGE_NAME="connectormgmtclient"
    OUTPUT_PATH="app-services-sdk-go/connectormgmt/apiv1/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

if [ "$INPUT_FILE" = "kafka-admin-rest.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/kafka-admin-rest.yaml"
    PACKAGE_NAME="kafkainstanceclient"
    OUTPUT_PATH="app-services-sdk-go/kafkainstance/apiv1/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

if [ "$INPUT_FILE" = "ams.json" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/ams.json"
    PATCH_FILE=".openapi/ams.patch" 
    PACKAGE_NAME="accountmgmtclient"
    OUTPUT_PATH="app-services-sdk-go/accountmgmt/apiv1/client"

    patch $OPENAPI_FILENAME < $PATCH_FILE

    rm -Rf $OUTPUT_PATH
    npx @openapitools/openapi-generator-cli generate -g go -i \
        "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
        --package-name="${PACKAGE_NAME}" \
        --additional-properties=$additional_properties \
        --ignore-file-override=./accountmgmt/.openapi-generator-ignore

    git checkout -- $OPENAPI_FILENAME
fi

if [ "$INPUT_FILE" = "registry-instance.json" ] || [ "$INPUT_FILE" = "" ];
then
    echo "generating registry instance SDK "

    cd .openapi
    echo "Removing codegen "
    cat registry-instance.json | jq 'del(.paths."x-codegen-contextRoot")' > registry-instance-tmp.json
    mv -f registry-instance-tmp.json registry-instance.json

    cat registry-instance.json | sed "s/create.extended+json/json/" > registry-instance-tmp.json
    mv -f registry-instance-tmp.json registry-instance.json

    echo "Ensuring only single tag is created "
    cat registry-instance.json | jq 'walk( if type == "object" and has("tags") 
        then .tags |= select(.[0])
        else . end )' > registry-instance-tmp.json
    mv -f registry-instance-tmp.json registry-instance.json

    echo "Removing invalid datetime definitions"
    sed -i '' 's/date-time/utc-date/' registry-instance.json

    cd ..

    OPENAPI_FILENAME=".openapi/registry-instance.json"
    PACKAGE_NAME="registryinstanceclient"
    OUTPUT_PATH="app-services-sdk-go/registryinstance/apiv1internal/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

if [ "$INPUT_FILE" = "service-accounts.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/service-accounts.yaml"
    PACKAGE_NAME="serviceaccountsclient"
    OUTPUT_PATH="app-services-sdk-go/serviceaccountmgmt/apiv1/client"
    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi


if [ "$INPUT_FILE" = "smartevents_mgmt_v2.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/smartevents_mgmt_v2.yaml"
    PACKAGE_NAME="smarteventsmgmtclient"
    OUTPUT_PATH="app-services-sdk-go/smarteventsmgmt/apiv1alpha/client"
    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi
