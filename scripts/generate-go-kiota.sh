#!/usr/bin/env bash

echo "Hello World!"

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

    echo "inside generate sdk"

    local file_name=$1
    local output_path=$2
    local package_name=$3

    echo "Generating source code based on ${file_name}"

    # remove old generated models
    rm -Rf $OUTPUT_PATH
    
    kiota generate -l go -c RegistryInstance -n kiota_posts/client -d .openapi/registry-instance.json -o ./app-services-sdk-go/registryinstance/apiv1internal/client
}


if [ "$INPUT_FILE" = "registry-instance.json" ] || [ "$INPUT_FILE" = "" ];
then
    echo "generating registry instance SDK "

    # cd .openapi
    # echo "Removing codegen "
    # cat registry-instance.json | jq 'del(.paths."x-codegen-contextRoot")' > registry-instance-tmp.json
    # mv -f registry-instance-tmp.json registry-instance.json

    # cat registry-instance.json | sed "s/create.extended+json/json/" > registry-instance-tmp.json
    # mv -f registry-instance-tmp.json registry-instance.json

    # echo "Ensuring only single tag is created "
    # cat registry-instance.json | jq 'walk( if type == "object" and has("tags") 
    #     then .tags |= select(.[0])
    #     else . end )' > registry-instance-tmp.json
    # mv -f registry-instance-tmp.json registry-instance.json

    # echo "Removing invalid datetime definitions"
    # sed -i '' 's/date-time/utc-date/' registry-instance.json

    # cd ..

    OPENAPI_FILENAME=".openapi/registry-instance.json"
    PACKAGE_NAME="registryinstanceclient"
    OUTPUT_PATH="app-services-sdk-go/registryinstance/apiv1internal/client"

    generate_sdk $OPENAPI_FILENAME $OUTPUT_PATH $PACKAGE_NAME
fi

#  kiota generate -l go -c RegistryInstance -n kiota_posts/client -d ./schema.json -o ./client