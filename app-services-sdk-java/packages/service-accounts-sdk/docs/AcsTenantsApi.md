# AcsTenantsApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAcsClient**](AcsTenantsApi.md#createAcsClient) | **POST** /apis/beta/acs/v1 | Create ACS managed central client
[**deleteAcsClient**](AcsTenantsApi.md#deleteAcsClient) | **DELETE** /apis/beta/acs/v1/{clientId} | Delete ACS managed central client



## createAcsClient

> AcsClientResponseData createAcsClient(acsClientRequestData)

Create ACS managed central client

Create an ACS managed central client. Created ACS managed central clients are associated with the supplied organization id.

### Example

```java
// Import classes:
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.auth.*;
import com.openshift.cloud.api.serviceaccounts.invoker.models.*;
import com.openshift.cloud.api.serviceaccounts.AcsTenantsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("https://sso.redhat.com/auth/realms/redhat-external");
        
        // Configure OAuth2 access token for authorization: serviceAccounts
        OAuth serviceAccounts = (OAuth) defaultClient.getAuthentication("serviceAccounts");
        serviceAccounts.setAccessToken("YOUR ACCESS TOKEN");

        AcsTenantsApi apiInstance = new AcsTenantsApi(defaultClient);
        AcsClientRequestData acsClientRequestData = new AcsClientRequestData(); // AcsClientRequestData | The name, redirect URIs and the organization id of the ACS managed central client
        try {
            AcsClientResponseData result = apiInstance.createAcsClient(acsClientRequestData);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AcsTenantsApi#createAcsClient");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acsClientRequestData** | [**AcsClientRequestData**](AcsClientRequestData.md)| The name, redirect URIs and the organization id of the ACS managed central client |

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
| **201** | OK |  -  |
| **400** | All fields did not pass validation. |  -  |
| **401** | Unauthorized |  -  |
| **403** | Exceeded maximum number of ACS managed central clients per tenant. |  -  |
| **405** | Not allowed, API Currently Disabled |  -  |


## deleteAcsClient

> deleteAcsClient(clientId)

Delete ACS managed central client

Delete ACS managed central client by clientId. Throws not found exception if the client is not found

### Example

```java
// Import classes:
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.auth.*;
import com.openshift.cloud.api.serviceaccounts.invoker.models.*;
import com.openshift.cloud.api.serviceaccounts.AcsTenantsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("https://sso.redhat.com/auth/realms/redhat-external");
        
        // Configure OAuth2 access token for authorization: serviceAccounts
        OAuth serviceAccounts = (OAuth) defaultClient.getAuthentication("serviceAccounts");
        serviceAccounts.setAccessToken("YOUR ACCESS TOKEN");

        AcsTenantsApi apiInstance = new AcsTenantsApi(defaultClient);
        String clientId = "clientId_example"; // String | 
        try {
            apiInstance.deleteAcsClient(clientId);
        } catch (ApiException e) {
            System.err.println("Exception when calling AcsTenantsApi#deleteAcsClient");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **String**|  |

### Return type

null (empty response body)

### Authorization

[serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **404** | Not Found |  -  |
| **405** | Not allowed, API Currently Disabled |  -  |

