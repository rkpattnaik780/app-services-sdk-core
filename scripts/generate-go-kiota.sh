#!/usr/bin/env bash
## OPENAPI_FILENAME=yourapi generate.sh 

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# if input file is a empty string that means this is a workflow run
# manualy and we want to generate all sdks 
INPUT_FILE=`basename $1`
INPUT_FILE=$(${SCRIPT_DIR}/get-filename.sh ${INPUT_FILE})

echo "========================="
echo "Input file is $INPUT_FILE"
echo "========================="

# generate an API client for a service
generate_sdk() {
    local file_name=$1
    local output_path=$2
    local class_name=$3
     
    echo "Validating OpenAPI ${file_name}"
    npx @openapitools/openapi-generator-cli validate -i "$file_name"

    echo "Generating source code based on ${file_name}"

    remove old generated models
    rm -Rf $output_path
    
    # npx @openapitools/openapi-generator-cli generate -g go -i \
    # "$file_name" -o "$output_path" \
    # --package-name="${package_name}" \
    # --additional-properties=$additional_properties \
    # --ignore-file-override=.openapi-generator-ignore

    PACKAGE_NAME="linux-x64"
    if [[ $OSTYPE == 'darwin'* ]]; then
    PACKAGE_NAME="osx-x64"
    fi
    URL="https://github.com/microsoft/kiota/releases/download/v1.0.1/${PACKAGE_NAME}.zip"

    COMMAND="kiota"
    if ! command -v $COMMAND &> /dev/null
    then
    echo "System wide kiota could not be found, using local version"
    if [[ ! -f $SCRIPT_DIR/kiota ]]
    then
        echo "Local kiota could not be found, downloading"
        rm -rf $SCRIPT_DIR/tmp-kiota
        mkdir -p $SCRIPT_DIR/tmp-kiota
        curl -sL $URL > $SCRIPT_DIR/tmp-kiota/kiota.zip
        unzip $SCRIPT_DIR/tmp-kiota/kiota.zip -d $SCRIPT_DIR/tmp-kiota

        mkdir -p $SCRIPT_DIR/tmp-kiota/bin
        cp $SCRIPT_DIR/tmp-kiota/*/kiota $SCRIPT_DIR/kiota
        chmod a+x $SCRIPT_DIR/kiota
        rm -rf $SCRIPT_DIR/tmp-kiota
    fi
    COMMAND="$SCRIPT_DIR/kiota"
    fi

    $COMMAND generate \
    --language go \
    --class-name "$class_name" \
    --namespace-name "$output_path" \
    --openapi "$file_name" \
    --output "$output_path" \

}

npx @openapitools/openapi-generator-cli version-manager set 5.2.0
echo "Generating Go SDKs"
# additional_properties="generateInterfaces=true,enumClassPrefix=true"

# if [ "$INPUT_FILE" = "kas-fleet-manager.yaml" ] || [ "$INPUT_FILE" = "" ];
# then
#     OPENAPI_FILENAME=".openapi/kas-fleet-manager.yaml"
#     PACKAGE_NAME="kafkamgmtclient"
#     OUTPUT_PATH="app-services-sdk-go/kafkamgmt/apiv1/client"

#     generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
# fi

# if [ "$INPUT_FILE" = "srs-fleet-manager.json" ] || [ "$INPUT_FILE" = "" ];
# then
# OPENAPI_FILENAME=".openapi/srs-fleet-manager.json"
# PACKAGE_NAME="registrymgmtclient"
# OUTPUT_PATH="app-services-sdk-go/registrymgmt/apiv1/client"

# generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
# fi

# if [ "$INPUT_FILE" = "connector_mgmt.yaml" ] || [ "$INPUT_FILE" = "" ];
# then
#     OPENAPI_FILENAME=".openapi/connector_mgmt.yaml"
#     PACKAGE_NAME="connectormgmtclient"
#     OUTPUT_PATH="app-services-sdk-go/connectormgmt/apiv1/client"

#     generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
# fi

# if [ "$INPUT_FILE" = "kafka-admin-rest.yaml" ] || [ "$INPUT_FILE" = "" ];
# then
#     OPENAPI_FILENAME=".openapi/kafka-admin-rest.yaml"
#     PACKAGE_NAME="kafkainstanceclient"
#     OUTPUT_PATH="app-services-sdk-go/kafkainstance/apiv1/client"

#     generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
# fi

# if [ "$INPUT_FILE" = "ams.json" ] || [ "$INPUT_FILE" = "" ];
# then
#     OPENAPI_FILENAME=".openapi/ams.json"
#     PATCH_FILE=".openapi/ams.patch" 
#     PACKAGE_NAME="accountmgmtclient"
#     OUTPUT_PATH="app-services-sdk-go/accountmgmt/apiv1/client"

#     patch $OPENAPI_FILENAME < $PATCH_FILE

#     rm -Rf $OUTPUT_PATH
#     npx @openapitools/openapi-generator-cli generate -g go -i \
#         "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
#         --package-name="${PACKAGE_NAME}" \
#         --additional-properties=$additional_properties \
#         --ignore-file-override=./accountmgmt/.openapi-generator-ignore

#     git checkout -- $OPENAPI_FILENAME
# fi

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
    CLASS_NAME="RegistryInstance"
    OUTPUT_PATH="app-services-sdk-go/registryinstance/apiv1internal/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $CLASS_NAME
fi

if [ "$INPUT_FILE" = "service-accounts.yaml" ] || [ "$INPUT_FILE" = "" ];
then
    OPENAPI_FILENAME=".openapi/service-accounts.yaml"
    CLASS_NAME="ServiceAccount"
    OUTPUT_PATH="app-services-sdk-go/serviceaccountmgmt/apiv1/client"
    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $CLASS_NAME
fi


# if [ "$INPUT_FILE" = "smartevents_mgmt_v2.yaml" ] || [ "$INPUT_FILE" = "" ];
# then
#     OPENAPI_FILENAME=".openapi/smartevents_mgmt_v2.yaml"
#     PACKAGE_NAME="smarteventsmgmtclient"
#     OUTPUT_PATH="app-services-sdk-go/smarteventsmgmt/apiv1alpha/client"
#     generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
# fi

# this hack is due to the api generator not ignoring the correct files
# so we need to revert the changes it does on these files. It ignores
# the other mod files but for some reason it decides to change this
# git restore app-services-sdk-go/accountmgmt/apiv1/client/go.mod
# git restore app-services-sdk-go/accountmgmt/apiv1/client/go.sum
