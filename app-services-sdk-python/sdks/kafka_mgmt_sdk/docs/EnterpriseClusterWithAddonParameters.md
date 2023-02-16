# EnterpriseClusterWithAddonParameters

Enterprise cluster with addon parameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**kind** | **str** |  | 
**href** | **str** |  | 
**access_kafkas_via_private_network** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**multi_az** | **bool** | A flag indicating whether this cluster is available on multiple availability zones or not | 
**cluster_id** | **str** | The OCM&#39;s cluster id of the registered Enterprise cluster. | [optional] 
**status** | **str** | The status of Enterprise cluster registration | [optional] 
**cloud_provider** | **str** | The cloud provider for this cluster. This valus will be used as the Kafka&#39;s cloud provider value when a Kafka is created on this cluster | [optional] 
**region** | **str** | The region of this cluster. This valus will be used as the Kafka&#39;s region value when a Kafka is created on this cluster | [optional] 
**fleetshard_parameters** | **[bool, date, datetime, dict, float, int, list, str, none_type]** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


