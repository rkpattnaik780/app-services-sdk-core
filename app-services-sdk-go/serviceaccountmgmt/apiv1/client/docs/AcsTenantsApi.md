# \AcsTenantsApi

All URIs are relative to *https://sso.redhat.com/auth/realms/redhat-external*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAcsClient**](AcsTenantsApi.md#CreateAcsClient) | **Post** /apis/beta/acs/v1 | Create ACS managed central client
[**DeleteAcsClient**](AcsTenantsApi.md#DeleteAcsClient) | **Delete** /apis/beta/acs/v1/{clientId} | Delete ACS managed central client



## CreateAcsClient

> AcsClientResponseData CreateAcsClient(ctx).AcsClientRequestData(acsClientRequestData).Execute()

Create ACS managed central client



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
    acsClientRequestData := *openapiclient.NewAcsClientRequestData([]string{"RedirectUris_example"}, "OrgId_example") // AcsClientRequestData | The name, redirect URIs and the organization id of the ACS managed central client

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AcsTenantsApi.CreateAcsClient(context.Background()).AcsClientRequestData(acsClientRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsTenantsApi.CreateAcsClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAcsClient`: AcsClientResponseData
    fmt.Fprintf(os.Stdout, "Response from `AcsTenantsApi.CreateAcsClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAcsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acsClientRequestData** | [**AcsClientRequestData**](AcsClientRequestData.md) | The name, redirect URIs and the organization id of the ACS managed central client | 

### Return type

[**AcsClientResponseData**](AcsClientResponseData.md)

### Authorization

[serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAcsClient

> DeleteAcsClient(ctx, clientId).Execute()

Delete ACS managed central client



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
    clientId := "clientId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AcsTenantsApi.DeleteAcsClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcsTenantsApi.DeleteAcsClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAcsClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[serviceAccounts](../README.md#serviceAccounts)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

