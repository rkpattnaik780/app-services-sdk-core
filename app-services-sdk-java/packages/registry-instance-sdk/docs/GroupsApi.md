# GroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createGroup**](GroupsApi.md#createGroup) | **POST** /groups | Create a new group
[**deleteGroupById**](GroupsApi.md#deleteGroupById) | **DELETE** /groups/{groupId} | Delete a group by the specified ID.
[**getGroupById**](GroupsApi.md#getGroupById) | **GET** /groups/{groupId} | Get a group by the specified ID.
[**listGroups**](GroupsApi.md#listGroups) | **GET** /groups | List groups



## createGroup

> GroupMetaData createGroup(createGroupMetaData)

Create a new group

Creates a new group.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) * The group already exist (HTTP error &#x60;409&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.GroupsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GroupsApi apiInstance = new GroupsApi(defaultClient);
        CreateGroupMetaData createGroupMetaData = new CreateGroupMetaData(); // CreateGroupMetaData | 
        try {
            GroupMetaData result = apiInstance.createGroup(createGroupMetaData);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupsApi#createGroup");
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
 **createGroupMetaData** | [**CreateGroupMetaData**](CreateGroupMetaData.md)|  |

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
| **200** | The group has been successfully created. |  -  |
| **409** | Common response used when an input conflicts with existing data. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## deleteGroupById

> deleteGroupById(groupId)

Delete a group by the specified ID.

Deletes a group by identifier.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) * The group does not exist (HTTP error &#x60;404&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.GroupsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GroupsApi apiInstance = new GroupsApi(defaultClient);
        String groupId = "groupId_example"; // String | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
        try {
            apiInstance.deleteGroupById(groupId);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupsApi#deleteGroupById");
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
 **groupId** | **String**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Empty content indicates a successful deletion. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## getGroupById

> GroupMetaData getGroupById(groupId)

Get a group by the specified ID.

Returns a group using the specified id.  This operation can fail for the following reasons:  * No group exists with the specified ID (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;)

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.GroupsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GroupsApi apiInstance = new GroupsApi(defaultClient);
        String groupId = "groupId_example"; // String | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
        try {
            GroupMetaData result = apiInstance.getGroupById(groupId);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupsApi#getGroupById");
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
 **groupId** | **String**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |

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
| **200** | The group&#39;s metadata. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listGroups

> GroupSearchResults listGroups(limit, offset, order, orderby)

List groups

Returns a list of all groups.  This list is paged.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.GroupsApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GroupsApi apiInstance = new GroupsApi(defaultClient);
        Integer limit = 56; // Integer | The number of groups to return.  Defaults to 20.
        Integer offset = 56; // Integer | The number of groups to skip before starting the result set.  Defaults to 0.
        SortOrder order = SortOrder.fromValue("asc"); // SortOrder | Sort order, ascending (`asc`) or descending (`desc`).
        SortBy orderby = SortBy.fromValue("name"); // SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn` 
        try {
            GroupSearchResults result = apiInstance.listGroups(limit, offset, order, orderby);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GroupsApi#listGroups");
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
 **limit** | **Integer**| The number of groups to return.  Defaults to 20. | [optional]
 **offset** | **Integer**| The number of groups to skip before starting the result set.  Defaults to 0. | [optional]
 **order** | [**SortOrder**](.md)| Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | [optional] [enum: asc, desc]
 **orderby** | [**SortBy**](.md)| The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | [optional] [enum: name, createdOn]

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
| **200** | On a successful response, returns a bounded set of groups. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |

