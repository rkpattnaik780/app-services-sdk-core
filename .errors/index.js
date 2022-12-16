


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
    js : {
        kafka: {
            definition: require("./errors_kafka_mgmt.json"),
            file: "app-services-sdk-js/packages/kafka-management-sdk/src/errors.ts"
        },
        srs:  {
            definition: require("./errors_srs_mgmt.json"),
            file: "app-services-sdk-js/packages/registry-management-sdk/src/errors.ts"
        },
        connector: {
            definition: require("./errors_connector_mgmt.json"),
            file: "app-services-sdk-js/packages/connector-management-sdk/src/errors.ts"
        }, 
        kafkainstance: {
            definition: require("./errors_kafka_instance.json"),
            file: "app-services-sdk-js/packages/kafka-instance-sdk/src/errors.ts"
        }, 
    }
}

