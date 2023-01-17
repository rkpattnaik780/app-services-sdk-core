# rhoas_service_accounts_mgmt_sdk.DefaultApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_authentication_policy**](DefaultApi.md#get_authentication_policy) | **GET** /apis/organizations/v1/{id}/authentication-policy | Get current authentication policy information
[**set_authentication_policy**](DefaultApi.md#set_authentication_policy) | **POST** /apis/organizations/v1/{id}/authentication-policy | Update current authentication policy information


# **get_authentication_policy**
> AuthenticationPolicy get_authentication_policy(id)

Get current authentication policy information

### Example

* OAuth Authentication (authFlow):
* Bearer (JWT) Authentication (bearerAuth):
* OAuth Authentication (serviceAccounts):

```python
import time
import rhoas_service_accounts_mgmt_sdk
from rhoas_service_accounts_mgmt_sdk.api import default_api
from rhoas_service_accounts_mgmt_sdk.model.error import Error
from rhoas_service_accounts_mgmt_sdk.model.authentication_policy import AuthenticationPolicy
from pprint import pprint
# Defining the host is optional and defaults to https://sso.redhat.com/auth/realms/redhat-external
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure OAuth2 access token for authorization: authFlow
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Configure Bearer authorization (JWT): bearerAuth
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Configure OAuth2 access token for authorization: serviceAccounts
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Enter a context with an instance of the API client
with rhoas_service_accounts_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    id = "id_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get current authentication policy information
        api_response = api_instance.get_authentication_policy(id)
        pprint(api_response)
    except rhoas_service_accounts_mgmt_sdk.ApiException as e:
        print("Exception when calling DefaultApi->get_authentication_policy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**|  |

### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

[authFlow](../README.md#authFlow), [bearerAuth](../README.md#bearerAuth), [serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_authentication_policy**
> AuthenticationPolicy set_authentication_policy(id)

Update current authentication policy information

### Example

* OAuth Authentication (authFlow):
* Bearer (JWT) Authentication (bearerAuth):
* OAuth Authentication (serviceAccounts):

```python
import time
import rhoas_service_accounts_mgmt_sdk
from rhoas_service_accounts_mgmt_sdk.api import default_api
from rhoas_service_accounts_mgmt_sdk.model.error import Error
from rhoas_service_accounts_mgmt_sdk.model.authentication_policy import AuthenticationPolicy
from pprint import pprint
# Defining the host is optional and defaults to https://sso.redhat.com/auth/realms/redhat-external
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure OAuth2 access token for authorization: authFlow
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Configure Bearer authorization (JWT): bearerAuth
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Configure OAuth2 access token for authorization: serviceAccounts
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Enter a context with an instance of the API client
with rhoas_service_accounts_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = default_api.DefaultApi(api_client)
    id = "id_example" # str | 
    authentication_policy = AuthenticationPolicy(
        authentication_factors=AuthenticationFactors(
            otp=Otp(
                required=True,
            ),
        ),
    ) # AuthenticationPolicy |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update current authentication policy information
        api_response = api_instance.set_authentication_policy(id)
        pprint(api_response)
    except rhoas_service_accounts_mgmt_sdk.ApiException as e:
        print("Exception when calling DefaultApi->set_authentication_policy: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update current authentication policy information
        api_response = api_instance.set_authentication_policy(id, authentication_policy=authentication_policy)
        pprint(api_response)
    except rhoas_service_accounts_mgmt_sdk.ApiException as e:
        print("Exception when calling DefaultApi->set_authentication_policy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**|  |
 **authentication_policy** | [**AuthenticationPolicy**](AuthenticationPolicy.md)|  | [optional]

### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

[authFlow](../README.md#authFlow), [bearerAuth](../README.md#bearerAuth), [serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

