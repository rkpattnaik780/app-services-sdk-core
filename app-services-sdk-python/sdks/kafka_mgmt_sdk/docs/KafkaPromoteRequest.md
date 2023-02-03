# KafkaPromoteRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**desired_kafka_billing_model** | **str** | The desired Kafka billing model to promote the kafka instance to. Promotion is performed asynchronously. Accepted values: [&#39;marketplace&#39;, &#39;standard&#39;] | 
**desired_marketplace** | **str** | The desired billing marketplace to promote the kafka instance to. Accepted values: [&#39;aws&#39;, &#39;rhm&#39;]. Only considered when desired_kafka_billing_model is &#39;marketplace&#39; | [optional] 
**desired_billing_cloud_account_id** | **str** | The desired Kafka billing cloud account ID to promote the kafka instance to. Only considered when desired_kafka_billing_model is &#39;marketplace&#39; | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


