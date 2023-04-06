


module.exports ={
    go : {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-go/kafkamgmt/apiv1/error/errors.go"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-go/registrymgmt/apiv1/error/errors.go"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-go/connectormgmt/apiv1/error/errors.go"
        }, 
    },
    ts : {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-ts/packages/kafka-management-sdk/src/errors.ts"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-ts/packages/registry-management-sdk/src/errors.ts"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-ts/packages/connector-management-sdk/src/errors.ts"
        }, 
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-ts/packages/kafka-instance-sdk/src/errors.ts"
        }, 
    },
    python: {
        kafkamgmt: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-python/sdks/kafka_mgmt_sdk/errors.py"
        },
        registrymgmt:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-python/sdks/registry_mgmt_sdk/errors.py"
        },
        connectormgmt: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-python/sdks/connector_mgmt_sdk/errors.py"
        }, 
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-python/sdks/kafka_instance_sdk/errors.py"
        }, 
    }
}

