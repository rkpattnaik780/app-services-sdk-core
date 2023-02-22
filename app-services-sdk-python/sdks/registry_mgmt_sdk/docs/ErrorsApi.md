# rhoas_service_registry_mgmt_sdk.ErrorsApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_error**](ErrorsApi.md#get_error) | **GET** /api/serviceregistry_mgmt/v1/errors/{id} | 
[**get_errors**](ErrorsApi.md#get_errors) | **GET** /api/serviceregistry_mgmt/v1/errors | 


# **get_error**
> Error get_error(id)



Get information about a specific error type

### Example


```python
import time
import rhoas_service_registry_mgmt_sdk
from rhoas_service_registry_mgmt_sdk.api import errors_api
from rhoas_service_registry_mgmt_sdk.model.empty import Empty
from rhoas_service_registry_mgmt_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_service_registry_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)


# Enter a context with an instance of the API client
with rhoas_service_registry_mgmt_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = errors_api.ErrorsApi(api_client)
    id = "id_example" # str | The id of the object you wish to interact with.

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_error(id)
        pprint(api_response)
    except rhoas_service_registry_mgmt_sdk.ApiException as e:
        print("Exception when calling ErrorsApi->get_error: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| The id of the object you wish to interact with. |

### Return type

[**Error**](Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response - returns a single &#x60;Error&#x60;. |  -  |
**401** | Authentication was not successful. Make sure the token is valid. |  -  |
**403** | User is not authorized to perform the operation. |  -  |
**404** | No Service Registry with the specified id exists. |  -  |
**500** | Unexpected error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_errors**
> ErrorList get_errors()



Get the list of all errors

### Example


```python
import time
import rhoas_service_registry_mgmt_sdk
from rhoas_service_registry_mgmt_sdk.api import errors_api
from rhoas_service_registry_mgmt_sdk.model.empty import Empty
from rhoas_service_registry_mgmt_sdk.model.error_list import ErrorList
from rhoas_service_registry_mgmt_sdk.model.error import Error
from pprint import pprint
# Defining the host is optional and defaults to https://api.openshift.com
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_service_registry_mgmt_sdk.Configuration(
    host = "https://api.openshift.com"
)


# Enter a context with an instance of the API client
with rhoas_service_registry_mgmt_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = errors_api.ErrorsApi(api_client)
    page = 0 # int | Page index. (optional)
    size = 100 # int | Number of items in each page. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_errors(page=page, size=size)
        pprint(api_response)
    except rhoas_service_registry_mgmt_sdk.ApiException as e:
        print("Exception when calling ErrorsApi->get_errors: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page index. | [optional]
 **size** | **int**| Number of items in each page. | [optional]

### Return type

[**ErrorList**](ErrorList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**400** | Invalid request content or parameters. |  -  |
**401** | Authentication was not successful. Make sure the token is valid. |  -  |
**403** | User is not authorized to perform the operation. |  -  |
**500** | Unexpected error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

