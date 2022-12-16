# DefaultApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAuthenticationPolicy**](DefaultApi.md#getAuthenticationPolicy) | **GET** /apis/organizations/v1/{id}/authentication-policy | Get current authentication policy information
[**setAuthenticationPolicy**](DefaultApi.md#setAuthenticationPolicy) | **POST** /apis/organizations/v1/{id}/authentication-policy | Update current authentication policy information



## getAuthenticationPolicy

> AuthenticationPolicy getAuthenticationPolicy(id)

Get current authentication policy information

### Example

```java
// Import classes:
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.auth.*;
import com.openshift.cloud.api.serviceaccounts.invoker.models.*;
import com.openshift.cloud.api.serviceaccounts.DefaultApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("https://sso.redhat.com/auth/realms/redhat-external");
        
        // Configure OAuth2 access token for authorization: authFlow
        OAuth authFlow = (OAuth) defaultClient.getAuthentication("authFlow");
        authFlow.setAccessToken("YOUR ACCESS TOKEN");

        // Configure HTTP bearer authorization: bearerAuth
        HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
        bearerAuth.setBearerToken("BEARER TOKEN");

        // Configure OAuth2 access token for authorization: serviceAccounts
        OAuth serviceAccounts = (OAuth) defaultClient.getAuthentication("serviceAccounts");
        serviceAccounts.setAccessToken("YOUR ACCESS TOKEN");

        DefaultApi apiInstance = new DefaultApi(defaultClient);
        String id = "id_example"; // String | 
        try {
            AuthenticationPolicy result = apiInstance.getAuthenticationPolicy(id);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DefaultApi#getAuthenticationPolicy");
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
 **id** | **String**|  |

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
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |


## setAuthenticationPolicy

> AuthenticationPolicy setAuthenticationPolicy(id, authenticationPolicy)

Update current authentication policy information

### Example

```java
// Import classes:
import com.openshift.cloud.api.serviceaccounts.invoker.ApiClient;
import com.openshift.cloud.api.serviceaccounts.invoker.ApiException;
import com.openshift.cloud.api.serviceaccounts.invoker.Configuration;
import com.openshift.cloud.api.serviceaccounts.invoker.auth.*;
import com.openshift.cloud.api.serviceaccounts.invoker.models.*;
import com.openshift.cloud.api.serviceaccounts.DefaultApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("https://sso.redhat.com/auth/realms/redhat-external");
        
        // Configure OAuth2 access token for authorization: authFlow
        OAuth authFlow = (OAuth) defaultClient.getAuthentication("authFlow");
        authFlow.setAccessToken("YOUR ACCESS TOKEN");

        // Configure HTTP bearer authorization: bearerAuth
        HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
        bearerAuth.setBearerToken("BEARER TOKEN");

        // Configure OAuth2 access token for authorization: serviceAccounts
        OAuth serviceAccounts = (OAuth) defaultClient.getAuthentication("serviceAccounts");
        serviceAccounts.setAccessToken("YOUR ACCESS TOKEN");

        DefaultApi apiInstance = new DefaultApi(defaultClient);
        String id = "id_example"; // String | 
        AuthenticationPolicy authenticationPolicy = new AuthenticationPolicy(); // AuthenticationPolicy | 
        try {
            AuthenticationPolicy result = apiInstance.setAuthenticationPolicy(id, authenticationPolicy);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DefaultApi#setAuthenticationPolicy");
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
 **id** | **String**|  |
 **authenticationPolicy** | [**AuthenticationPolicy**](AuthenticationPolicy.md)|  | [optional]

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
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |

