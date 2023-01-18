# rhoas_kafka_mgmt_sdk.EnterpriseDataplaneClustersApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_enterprise_cluster_by_id**](EnterpriseDataplaneClustersApi.md#delete_enterprise_cluster_by_id) | **DELETE** /api/kafkas_mgmt/v1/clusters/{id} | 
[**get_enterprise_cluster_by_id**](EnterpriseDataplaneClustersApi.md#get_enterprise_cluster_by_id) | **GET** /api/kafkas_mgmt/v1/clusters/{id} | 
[**get_enterprise_cluster_with_addon_parameters**](EnterpriseDataplaneClustersApi.md#get_enterprise_cluster_with_addon_parameters) | **GET** /api/kafkas_mgmt/v1/clusters/{id}/addon_parameters | 
[**get_enterprise_osd_clusters**](EnterpriseDataplaneClustersApi.md#get_enterprise_osd_clusters) | **GET** /api/kafkas_mgmt/v1/clusters | 
[**register_enterprise_osd_cluster**](EnterpriseDataplaneClustersApi.md#register_enterprise_osd_cluster) | **POST** /api/kafkas_mgmt/v1/clusters | 


# **delete_enterprise_cluster_by_id**
> Error delete_enterprise_cluster_by_id(_async, id)



### Example

* Bearer (JWT) Authentication (Bearer):

```python
import time
import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.api import enterprise_dataplane_clusters_api
from rhoas_kafka_mgmt_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): Bearer
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_kafka_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = enterprise_dataplane_clusters_api.EnterpriseDataplaneClustersApi(api_client)
    _async = True # bool | Perform the action in an asynchronous manner
    id = "id_example" # str | ID of the enterprise data plane cluster
    force = True # bool | When provided with value: true - enterprise cluster will be deleted alongside all kafkas present on the cluster. When skipped and enterprise cluster has any kafkas associated with it, the request will fail. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.delete_enterprise_cluster_by_id(_async, id)
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->delete_enterprise_cluster_by_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.delete_enterprise_cluster_by_id(_async, id, force=force)
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->delete_enterprise_cluster_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **_async** | **bool**| Perform the action in an asynchronous manner |
 **id** | **str**| ID of the enterprise data plane cluster |
 **force** | **bool**| When provided with value: true - enterprise cluster will be deleted alongside all kafkas present on the cluster. When skipped and enterprise cluster has any kafkas associated with it, the request will fail. | [optional]

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Enterprise cluster deletion accepted |  -  |
**400** | Validation errors occurred |  -  |
**401** | Auth token is invalid |  -  |
**403** | User not authorized to access the service |  -  |
**404** | No Enterprise cluster with specified ID exists |  -  |
**500** | Unexpected error occurred |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_enterprise_cluster_by_id**
> EnterpriseCluster get_enterprise_cluster_by_id(id)



Returns enterprise data plane cluster by ID

### Example

* Bearer (JWT) Authentication (Bearer):

```python
import time
import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.api import enterprise_dataplane_clusters_api
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster import EnterpriseCluster
from rhoas_kafka_mgmt_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): Bearer
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_kafka_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = enterprise_dataplane_clusters_api.EnterpriseDataplaneClustersApi(api_client)
    id = "id_example" # str | ID of the enterprise data plane cluster

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_enterprise_cluster_by_id(id)
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->get_enterprise_cluster_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the enterprise data plane cluster |

### Return type

[**EnterpriseCluster**](EnterpriseCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Found enterprise data plane cluster with provided ID |  -  |
**401** | Auth token is invalid |  -  |
**403** | User not authorized to access the service |  -  |
**404** | No Enterprise data plane cluster with specified ID exists |  -  |
**500** | Unexpected error occurred |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_enterprise_cluster_with_addon_parameters**
> EnterpriseClusterWithAddonParameters get_enterprise_cluster_with_addon_parameters(id)



Returns enterprise data plane cluster by ID along with its addon parameters

### Example

* Bearer (JWT) Authentication (Bearer):

```python
import time
import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.api import enterprise_dataplane_clusters_api
from rhoas_kafka_mgmt_sdk.model.error import Error
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster_with_addon_parameters import EnterpriseClusterWithAddonParameters
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): Bearer
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_kafka_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = enterprise_dataplane_clusters_api.EnterpriseDataplaneClustersApi(api_client)
    id = "id_example" # str | ID of the enterprise data plane cluster

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_enterprise_cluster_with_addon_parameters(id)
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->get_enterprise_cluster_with_addon_parameters: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the enterprise data plane cluster |

### Return type

[**EnterpriseClusterWithAddonParameters**](EnterpriseClusterWithAddonParameters.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Returns the enterprise data plane cluster with addon parameters for an existing enterprise data plane cluster with provided ID |  -  |
**401** | Auth token is invalid |  -  |
**403** | User not authorized to access the service |  -  |
**404** | No Enterprise data plane cluster with specified ID exists |  -  |
**500** | Unexpected error occurred |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_enterprise_osd_clusters**
> EnterpriseClusterList get_enterprise_osd_clusters()



List all Enterprise data plane clusters

### Example

* Bearer (JWT) Authentication (Bearer):

```python
import time
import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.api import enterprise_dataplane_clusters_api
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster_list import EnterpriseClusterList
from rhoas_kafka_mgmt_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): Bearer
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_kafka_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = enterprise_dataplane_clusters_api.EnterpriseDataplaneClustersApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.get_enterprise_osd_clusters()
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->get_enterprise_osd_clusters: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**EnterpriseClusterList**](EnterpriseClusterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List Enterprise data plane clusters |  -  |
**400** | Bad request |  -  |
**401** | Auth token is invalid |  -  |
**403** | User not authorized to access the service |  -  |
**500** | Unexpected error occurred |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **register_enterprise_osd_cluster**
> EnterpriseClusterWithAddonParameters register_enterprise_osd_cluster(enterprise_osd_cluster_payload)



Register enterprise data plane cluster

### Example

* Bearer (JWT) Authentication (Bearer):

```python
import time
import rhoas_kafka_mgmt_sdk
from rhoas_kafka_mgmt_sdk.api import enterprise_dataplane_clusters_api
from rhoas_kafka_mgmt_sdk.model.error import Error
from rhoas_kafka_mgmt_sdk.model.enterprise_cluster_with_addon_parameters import EnterpriseClusterWithAddonParameters
from rhoas_kafka_mgmt_sdk.model.enterprise_osd_cluster_payload import EnterpriseOsdClusterPayload
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): Bearer
configuration = rhoas_kafka_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with rhoas_kafka_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = enterprise_dataplane_clusters_api.EnterpriseDataplaneClustersApi(api_client)
    enterprise_osd_cluster_payload = EnterpriseOsdClusterPayload(
        access_kafkas_via_private_network=True,
        cluster_id="cluster_id_example",
        cluster_external_id="cluster_external_id_example",
        cluster_ingress_dns_name="cluster_ingress_dns_name_example",
        kafka_machine_pool_node_count=1,
    ) # EnterpriseOsdClusterPayload | Enterprise data plane cluster details

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.register_enterprise_osd_cluster(enterprise_osd_cluster_payload)
        pprint(api_response)
    except rhoas_kafka_mgmt_sdk.ApiException as e:
        print("Exception when calling EnterpriseDataplaneClustersApi->register_enterprise_osd_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enterprise_osd_cluster_payload** | [**EnterpriseOsdClusterPayload**](EnterpriseOsdClusterPayload.md)| Enterprise data plane cluster details |

### Return type

[**EnterpriseClusterWithAddonParameters**](EnterpriseClusterWithAddonParameters.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Enterprise cluster registered |  -  |
**400** | Validation errors occurred |  -  |
**401** | Auth token is invalid |  -  |
**403** | User is not authorized to access the service |  -  |
**409** | A conflict has been detected in the creation of this resource |  -  |
**500** | An unexpected error occurred while registering Enterprise cluster |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

