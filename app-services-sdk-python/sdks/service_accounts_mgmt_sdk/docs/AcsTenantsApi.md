# rhoas_service_accounts_mgmt_sdk.AcsTenantsApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_acs_client**](AcsTenantsApi.md#create_acs_client) | **POST** /apis/beta/acs/v1 | Create ACS managed central client
[**delete_acs_client**](AcsTenantsApi.md#delete_acs_client) | **DELETE** /apis/beta/acs/v1/{clientId} | Delete ACS managed central client


# **create_acs_client**
> AcsClientResponseData create_acs_client(acs_client_request_data)

Create ACS managed central client

Create an ACS managed central client. Created ACS managed central clients are associated with the supplied organization id.

### Example

* OAuth Authentication (serviceAccounts):

```python
import time
import rhoas_service_accounts_mgmt_sdk
from rhoas_service_accounts_mgmt_sdk.api import acs_tenants_api
from rhoas_service_accounts_mgmt_sdk.model.acs_client_request_data import AcsClientRequestData
from rhoas_service_accounts_mgmt_sdk.model.error import Error
from rhoas_service_accounts_mgmt_sdk.model.acs_client_response_data import AcsClientResponseData
from rhoas_service_accounts_mgmt_sdk.model.red_hat_error_representation import RedHatErrorRepresentation
from rhoas_service_accounts_mgmt_sdk.model.validation_exception_data import ValidationExceptionData
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

# Configure OAuth2 access token for authorization: serviceAccounts
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Enter a context with an instance of the API client
with rhoas_service_accounts_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = acs_tenants_api.AcsTenantsApi(api_client)
    acs_client_request_data = AcsClientRequestData(
        name="name_example",
        redirect_uris=[
            "redirect_uris_example",
        ],
        org_id="4",
    ) # AcsClientRequestData | The name, redirect URIs and the organization id of the ACS managed central client

    # example passing only required values which don't have defaults set
    try:
        # Create ACS managed central client
        api_response = api_instance.create_acs_client(acs_client_request_data)
        pprint(api_response)
    except rhoas_service_accounts_mgmt_sdk.ApiException as e:
        print("Exception when calling AcsTenantsApi->create_acs_client: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acs_client_request_data** | [**AcsClientRequestData**](AcsClientRequestData.md)| The name, redirect URIs and the organization id of the ACS managed central client |

### Return type

[**AcsClientResponseData**](AcsClientResponseData.md)

### Authorization

[serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | OK |  -  |
**400** | All fields did not pass validation. |  -  |
**401** | Unauthorized |  -  |
**403** | Exceeded maximum number of ACS managed central clients per tenant. |  -  |
**405** | Not allowed, API Currently Disabled |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_acs_client**
> delete_acs_client(client_id)

Delete ACS managed central client

Delete ACS managed central client by clientId. Throws not found exception if the client is not found

### Example

* OAuth Authentication (serviceAccounts):

```python
import time
import rhoas_service_accounts_mgmt_sdk
from rhoas_service_accounts_mgmt_sdk.api import acs_tenants_api
from rhoas_service_accounts_mgmt_sdk.model.error import Error
from rhoas_service_accounts_mgmt_sdk.model.red_hat_error_representation import RedHatErrorRepresentation
from rhoas_service_accounts_mgmt_sdk.model.validation_exception_data import ValidationExceptionData
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

# Configure OAuth2 access token for authorization: serviceAccounts
configuration = rhoas_service_accounts_mgmt_sdk.Configuration(
    host = "https://sso.redhat.com/auth/realms/redhat-external"
)
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# Enter a context with an instance of the API client
with rhoas_service_accounts_mgmt_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = acs_tenants_api.AcsTenantsApi(api_client)
    client_id = "clientId_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Delete ACS managed central client
        api_instance.delete_acs_client(client_id)
    except rhoas_service_accounts_mgmt_sdk.ApiException as e:
        print("Exception when calling AcsTenantsApi->delete_acs_client: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client_id** | **str**|  |

### Return type

void (empty response body)

### Authorization

[serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request if clientId is invalid |  -  |
**401** | Unauthorized |  -  |
**404** | Not Found |  -  |
**405** | Not allowed, API Currently Disabled |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

