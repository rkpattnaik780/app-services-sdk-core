# rhoas_registry_instance_sdk.GroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_group**](GroupsApi.md#create_group) | **POST** /groups | Create a new group
[**delete_group_by_id**](GroupsApi.md#delete_group_by_id) | **DELETE** /groups/{groupId} | Delete a group by the specified ID.
[**get_group_by_id**](GroupsApi.md#get_group_by_id) | **GET** /groups/{groupId} | Get a group by the specified ID.
[**list_groups**](GroupsApi.md#list_groups) | **GET** /groups | List groups


# **create_group**
> GroupMetaData create_group(create_group_meta_data)

Create a new group

Creates a new group.  This operation can fail for the following reasons:  * A server error occurred (HTTP error `500`) * The group already exist (HTTP error `409`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import groups_api
from rhoas_registry_instance_sdk.model.group_meta_data import GroupMetaData
from rhoas_registry_instance_sdk.model.error import Error
from rhoas_registry_instance_sdk.model.create_group_meta_data import CreateGroupMetaData
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = groups_api.GroupsApi(api_client)
    create_group_meta_data = CreateGroupMetaData(
        description="description_example",
        properties=Properties(
            key="key_example",
        ),
        id="id_example",
    ) # CreateGroupMetaData | 

    # example passing only required values which don't have defaults set
    try:
        # Create a new group
        api_response = api_instance.create_group(create_group_meta_data)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling GroupsApi->create_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_group_meta_data** | [**CreateGroupMetaData**](CreateGroupMetaData.md)|  |

### Return type

[**GroupMetaData**](GroupMetaData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The group has been successfully created. |  -  |
**409** | Common response used when an input conflicts with existing data. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_group_by_id**
> delete_group_by_id(group_id)

Delete a group by the specified ID.

Deletes a group by identifier.  This operation can fail for the following reasons:  * A server error occurred (HTTP error `500`) * The group does not exist (HTTP error `404`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import groups_api
from rhoas_registry_instance_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = groups_api.GroupsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.

    # example passing only required values which don't have defaults set
    try:
        # Delete a group by the specified ID.
        api_instance.delete_group_by_id(group_id)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling GroupsApi->delete_group_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Empty content indicates a successful deletion. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_group_by_id**
> GroupMetaData get_group_by_id(group_id)

Get a group by the specified ID.

Returns a group using the specified id.  This operation can fail for the following reasons:  * No group exists with the specified ID (HTTP error `404`) * A server error occurred (HTTP error `500`)

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import groups_api
from rhoas_registry_instance_sdk.model.group_meta_data import GroupMetaData
from rhoas_registry_instance_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = groups_api.GroupsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.

    # example passing only required values which don't have defaults set
    try:
        # Get a group by the specified ID.
        api_response = api_instance.get_group_by_id(group_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling GroupsApi->get_group_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |

### Return type

[**GroupMetaData**](GroupMetaData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The group&#39;s metadata. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_groups**
> GroupSearchResults list_groups()

List groups

Returns a list of all groups.  This list is paged.

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import groups_api
from rhoas_registry_instance_sdk.model.sort_order import SortOrder
from rhoas_registry_instance_sdk.model.group_search_results import GroupSearchResults
from rhoas_registry_instance_sdk.model.error import Error
from rhoas_registry_instance_sdk.model.sort_by import SortBy
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = groups_api.GroupsApi(api_client)
    limit = 1 # int | The number of groups to return.  Defaults to 20. (optional)
    offset = 1 # int | The number of groups to skip before starting the result set.  Defaults to 0. (optional)
    order = SortOrder("asc") # SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby = SortBy("name") # SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List groups
        api_response = api_instance.list_groups(limit=limit, offset=offset, order=order, orderby=orderby)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling GroupsApi->list_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int**| The number of groups to return.  Defaults to 20. | [optional]
 **offset** | **int**| The number of groups to skip before starting the result set.  Defaults to 0. | [optional]
 **order** | **SortOrder**| Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | [optional]
 **orderby** | **SortBy**| The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | [optional]

### Return type

[**GroupSearchResults**](GroupSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | On a successful response, returns a bounded set of groups. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

