# \DefaultApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthenticationPolicy**](DefaultApi.md#GetAuthenticationPolicy) | **Get** /apis/organizations/v1/{id}/authentication-policy | Get current authentication policy information
[**SetAuthenticationPolicy**](DefaultApi.md#SetAuthenticationPolicy) | **Post** /apis/organizations/v1/{id}/authentication-policy | Update current authentication policy information



## GetAuthenticationPolicy

> AuthenticationPolicy GetAuthenticationPolicy(ctx, id).Execute()

Get current authentication policy information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAuthenticationPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationPolicy`: AuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

[authFlow](../README.md#authFlow), [bearerAuth](../README.md#bearerAuth), [serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAuthenticationPolicy

> AuthenticationPolicy SetAuthenticationPolicy(ctx, id).AuthenticationPolicy(authenticationPolicy).Execute()

Update current authentication policy information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    authenticationPolicy := *openapiclient.NewAuthenticationPolicy() // AuthenticationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetAuthenticationPolicy(context.Background(), id).AuthenticationPolicy(authenticationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAuthenticationPolicy`: AuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationPolicy** | [**AuthenticationPolicy**](AuthenticationPolicy.md) |  | 

### Return type

[**AuthenticationPolicy**](AuthenticationPolicy.md)

### Authorization

[authFlow](../README.md#authFlow), [bearerAuth](../README.md#bearerAuth), [serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

