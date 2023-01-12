# rhoas_smart_events_mgmt_sdk.SourceConnectorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**source_connectors_api_create_source_connector**](SourceConnectorsApi.md#source_connectors_api_create_source_connector) | **POST** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources | Create a Source Connector for a Bridge instance
[**source_connectors_api_delete_source_connector**](SourceConnectorsApi.md#source_connectors_api_delete_source_connector) | **DELETE** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Delete a Source Connector of a Bridge instance
[**source_connectors_api_get_source_connector**](SourceConnectorsApi.md#source_connectors_api_get_source_connector) | **GET** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Get a Source Connector of a Bridge instance
[**source_connectors_api_get_source_connectors**](SourceConnectorsApi.md#source_connectors_api_get_source_connectors) | **GET** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources | Get the list of Sink Connectors for a Bridge
[**source_connectors_api_update_source_connector**](SourceConnectorsApi.md#source_connectors_api_update_source_connector) | **PUT** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Update a Source Connector instance.


# **source_connectors_api_create_source_connector**
> SourceConnectorResponse source_connectors_api_create_source_connector(bridge_id)

Create a Source Connector for a Bridge instance

Create a Source Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import source_connectors_api
from rhoas_smart_events_mgmt_sdk.model.source_connector_response import SourceConnectorResponse
from rhoas_smart_events_mgmt_sdk.model.connector_request import ConnectorRequest
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
    api_instance = source_connectors_api.SourceConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    connector_request = ConnectorRequest(
        name="my-connector",
        connector_type_id="slack_sink_0.1",
        connector={},
    ) # ConnectorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Source Connector for a Bridge instance
        api_response = api_instance.source_connectors_api_create_source_connector(bridge_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_create_source_connector: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Source Connector for a Bridge instance
        api_response = api_instance.source_connectors_api_create_source_connector(bridge_id, connector_request=connector_request)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_create_source_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **connector_request** | [**ConnectorRequest**](ConnectorRequest.md)|  | [optional]

### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

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

# **source_connectors_api_delete_source_connector**
> source_connectors_api_delete_source_connector(bridge_id, source_id)

Delete a Source Connector of a Bridge instance

Delete a Source Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import source_connectors_api
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
    api_instance = source_connectors_api.SourceConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    source_id = "sourceId_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete a Source Connector of a Bridge instance
        api_instance.source_connectors_api_delete_source_connector(bridge_id, source_id)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_delete_source_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **source_id** | **str**|  |

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

# **source_connectors_api_get_source_connector**
> SourceConnectorResponse source_connectors_api_get_source_connector(bridge_id, source_id)

Get a Source Connector of a Bridge instance

Get a Source Connector of a Bridge instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import source_connectors_api
from rhoas_smart_events_mgmt_sdk.model.source_connector_response import SourceConnectorResponse
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
    api_instance = source_connectors_api.SourceConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    source_id = "sourceId_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a Source Connector of a Bridge instance
        api_response = api_instance.source_connectors_api_get_source_connector(bridge_id, source_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_get_source_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **source_id** | **str**|  |

### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

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

# **source_connectors_api_get_source_connectors**
> SourceConnectorListResponse source_connectors_api_get_source_connectors(bridge_id)

Get the list of Sink Connectors for a Bridge

Get the list of Source Connector instances of a Bridge instance instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import source_connectors_api
from rhoas_smart_events_mgmt_sdk.model.source_connector_list_response import SourceConnectorListResponse
from rhoas_smart_events_mgmt_sdk.model.managed_resource_status import ManagedResourceStatus
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
    api_instance = source_connectors_api.SourceConnectorsApi(api_client)
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
        api_response = api_instance.source_connectors_api_get_source_connectors(bridge_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_get_source_connectors: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get the list of Sink Connectors for a Bridge
        api_response = api_instance.source_connectors_api_get_source_connectors(bridge_id, name=name, page=page, size=size, status=status)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_get_source_connectors: %s\n" % e)
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

[**SourceConnectorListResponse**](SourceConnectorListResponse.md)

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

# **source_connectors_api_update_source_connector**
> SourceConnectorResponse source_connectors_api_update_source_connector(bridge_id, source_id)

Update a Source Connector instance.

Update a Source Connector instance for the authenticated user.

### Example

* Bearer Authentication (bearer):

```python
import time
import rhoas_smart_events_mgmt_sdk
from rhoas_smart_events_mgmt_sdk.api import source_connectors_api
from rhoas_smart_events_mgmt_sdk.model.source_connector_response import SourceConnectorResponse
from rhoas_smart_events_mgmt_sdk.model.connector_request import ConnectorRequest
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
    api_instance = source_connectors_api.SourceConnectorsApi(api_client)
    bridge_id = "bridgeId_example" # str | 
    source_id = "sourceId_example" # str | 
    connector_request = ConnectorRequest(
        name="my-connector",
        connector_type_id="slack_sink_0.1",
        connector={},
    ) # ConnectorRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Source Connector instance.
        api_response = api_instance.source_connectors_api_update_source_connector(bridge_id, source_id)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_update_source_connector: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Source Connector instance.
        api_response = api_instance.source_connectors_api_update_source_connector(bridge_id, source_id, connector_request=connector_request)
        pprint(api_response)
    except rhoas_smart_events_mgmt_sdk.ApiException as e:
        print("Exception when calling SourceConnectorsApi->source_connectors_api_update_source_connector: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridge_id** | **str**|  |
 **source_id** | **str**|  |
 **connector_request** | [**ConnectorRequest**](ConnectorRequest.md)|  | [optional]

### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

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

