# SinkConnectorResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**kind** | **str** | The kind (type) of this resource | 
**id** | **str** | The unique identifier of this resource | 
**name** | **str** | The name of this resource | 
**href** | **str** | The URL of this resource, without the protocol | 
**submitted_at** | **datetime** |  | 
**status** | **bool, date, datetime, dict, float, int, list, str, none_type** |  | 
**owner** | **str** | The user that owns this resource | 
**connector_type_id** | **str** | The connector type | 
**connector** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | The Connector configuration payload | 
**uri_dsl** | **str** | The URI to be used in Camel DSL to send data to this sink | 
**published_at** | **datetime** |  | [optional] 
**modified_at** | **datetime** |  | [optional] 
**status_message** | **str** | A detailed status message in case there is a problem with the connector | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


