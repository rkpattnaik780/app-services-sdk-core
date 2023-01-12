# rhoas_smart_events_mgmt_sdk.SinkConnectorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sink_connectors_api_create_sink_connector**](SinkConnectorsApi.md#sink_connectors_api_create_sink_connector) | **POST** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks | Create a Sink Connector for a Bridge instance
[**sink_connectors_api_delete_sink_connector**](SinkConnectorsApi.md#sink_connectors_api_delete_sink_connector) | **DELETE** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Delete a Sink Connector of a Bridge instance
[**sink_connectors_api_get_sink_connector**](SinkConnectorsApi.md#sink_connectors_api_get_sink_connector) | **GET** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Get a Sink Connector of a Bridge instance
[**sink_connectors_api_get_sink_connectors**](SinkConnectorsApi.md#sink_connectors_api_get_sink_connectors) | **GET** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks | Get the list of Sink Connectors for a Bridge
[**sink_connectors_api_update_sink_connector**](SinkConnectorsApi.md#sink_connectors_api_update_sink_connector) | **PUT** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Update a Sink Connector instance.


# **sink_connectors_api_create_sink_connector**
> SinkConnectorResponse sink_connectors_api_create_sink_connector(bridge_id)

Create a Sink Connector for a Bridge instance

Create a Sink Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import sink_connectors_api
from rhoas_smart_events_mgmt_sdk.model.connector_request import ConnectorRequest
from rhoas_smart_events_mgmt_sdk.model.sink_connector_response import SinkConnectorResponse
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from pprint import pprint
# Defining the host is optional and defaults to https://api.stage.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    host = "https://api.stage.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearer
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_smart_events_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sink_connectors_api.SinkConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    connector_request = ConnectorRequest(
        name="my-connector",
        connector_type_id="slack_sink_0.1",
        connector={},
    ) # ConnectorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Sink Connector for a Bridge instance
        api_response = api_instance.sink_connectors_api_create_sink_connector(bridge_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_create_sink_connector: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Sink Connector for a Bridge instance
        api_response = api_instance.sink_connectors_api_create_sink_connector(bridge_id, connector_request=connector_request)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_create_sink_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **connector_request** | [**ConnectorRequest**](ConnectorRequest.md)|  | [optional]

### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Accepted. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthorized. |  -  |
**402** | Not enough quota. |  -  |
**403** | Forbidden. |  -  |
**404** | Not found. |  -  |
**500** | Internal error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sink_connectors_api_delete_sink_connector**
> sink_connectors_api_delete_sink_connector(bridge_id, sink_id)

Delete a Sink Connector of a Bridge instance

Delete a Sink Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import sink_connectors_api
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from pprint import pprint
# Defining the host is optional and defaults to https://api.stage.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    host = "https://api.stage.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearer
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_smart_events_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sink_connectors_api.SinkConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    sink_id = "sinkId_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete a Sink Connector of a Bridge instance
        api_instance.sink_connectors_api_delete_sink_connector(bridge_id, sink_id)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_delete_sink_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **sink_id** | **str**|  |

### Return type

void (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Accepted. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Not found. |  -  |
**500** | Internal error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sink_connectors_api_get_sink_connector**
> SinkConnectorResponse sink_connectors_api_get_sink_connector(bridge_id, sink_id)

Get a Sink Connector of a Bridge instance

Get a Sink Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import sink_connectors_api
from rhoas_smart_events_mgmt_sdk.model.sink_connector_response import SinkConnectorResponse
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from pprint import pprint
# Defining the host is optional and defaults to https://api.stage.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    host = "https://api.stage.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearer
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_smart_events_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sink_connectors_api.SinkConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    sink_id = "sinkId_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a Sink Connector of a Bridge instance
        api_response = api_instance.sink_connectors_api_get_sink_connector(bridge_id, sink_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_get_sink_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **sink_id** | **str**|  |

### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Not found. |  -  |
**500** | Internal error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sink_connectors_api_get_sink_connectors**
> SinkConnectorListResponse sink_connectors_api_get_sink_connectors(bridge_id)

Get the list of Sink Connectors for a Bridge

Get the list of Sink Connector instances of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import sink_connectors_api
from rhoas_smart_events_mgmt_sdk.model.managed_resource_status import ManagedResourceStatus
from rhoas_smart_events_mgmt_sdk.model.sink_connector_list_response import SinkConnectorListResponse
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from pprint import pprint
# Defining the host is optional and defaults to https://api.stage.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    host = "https://api.stage.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearer
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_smart_events_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sink_connectors_api.SinkConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    name = "name_example" # str |  (optional)
    page = 0 # int |  (optional) if omitted the server will use the default value of 0
    size = 100 # int |  (optional) if omitted the server will use the default value of 100
    status = [
        ManagedResourceStatus("accepted"),
    ] # [ManagedResourceStatus] |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get the list of Sink Connectors for a Bridge
        api_response = api_instance.sink_connectors_api_get_sink_connectors(bridge_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_get_sink_connectors: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get the list of Sink Connectors for a Bridge
        api_response = api_instance.sink_connectors_api_get_sink_connectors(bridge_id, name=name, page=page, size=size, status=status)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_get_sink_connectors: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **name** | **str**|  | [optional]
 **page** | **int**|  | [optional] if omitted the server will use the default value of 0
 **size** | **int**|  | [optional] if omitted the server will use the default value of 100
 **status** | [**[ManagedResourceStatus]**](ManagedResourceStatus.md)|  | [optional]

### Return type

[**SinkConnectorListResponse**](SinkConnectorListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Not found. |  -  |
**500** | Internal error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sink_connectors_api_update_sink_connector**
> SinkConnectorResponse sink_connectors_api_update_sink_connector(bridge_id, sink_id)

Update a Sink Connector instance.

Update a Sink Connector instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import sink_connectors_api
from rhoas_smart_events_mgmt_sdk.model.connector_request import ConnectorRequest
from rhoas_smart_events_mgmt_sdk.model.sink_connector_response import SinkConnectorResponse
from rhoas_smart_events_mgmt_sdk.model.errors_list import ErrorsList
from pprint import pprint
# Defining the host is optional and defaults to https://api.stage.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    host = "https://api.stage.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearer
configuration = rhoas_smart_events_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_smart_events_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = sink_connectors_api.SinkConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    sink_id = "sinkId_example" # str | 
    connector_request = ConnectorRequest(
        name="my-connector",
        connector_type_id="slack_sink_0.1",
        connector={},
    ) # ConnectorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Sink Connector instance.
        api_response = api_instance.sink_connectors_api_update_sink_connector(bridge_id, sink_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_update_sink_connector: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Sink Connector instance.
        api_response = api_instance.sink_connectors_api_update_sink_connector(bridge_id, sink_id, connector_request=connector_request)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SinkConnectorsApi->sink_connectors_api_update_sink_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **sink_id** | **str**|  |
 **connector_request** | [**ConnectorRequest**](ConnectorRequest.md)|  | [optional]

### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Accepted. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Not found. |  -  |
**500** | Internal error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

