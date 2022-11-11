


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
    }
}

