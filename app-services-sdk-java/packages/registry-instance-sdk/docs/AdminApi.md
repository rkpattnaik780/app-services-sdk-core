# AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createGlobalRule**](AdminApi.md#createGlobalRule) | **POST** /admin/rules | Create global rule
[**createRoleMapping**](AdminApi.md#createRoleMapping) | **POST** /admin/roleMappings | Create a new role mapping
[**deleteAllGlobalRules**](AdminApi.md#deleteAllGlobalRules) | **DELETE** /admin/rules | Delete all global rules
[**deleteGlobalRule**](AdminApi.md#deleteGlobalRule) | **DELETE** /admin/rules/{rule} | Delete global rule
[**deleteRoleMapping**](AdminApi.md#deleteRoleMapping) | **DELETE** /admin/roleMappings/{principalId} | Delete a role mapping
[**exportData**](AdminApi.md#exportData) | **GET** /admin/export | Export registry data
[**getConfigProperty**](AdminApi.md#getConfigProperty) | **GET** /admin/config/properties/{propertyName} | Get configuration property value
[**getGlobalRuleConfig**](AdminApi.md#getGlobalRuleConfig) | **GET** /admin/rules/{rule} | Get global rule configuration
[**getLogConfiguration**](AdminApi.md#getLogConfiguration) | **GET** /admin/loggers/{logger} | Get a single logger configuration
[**getRoleMapping**](AdminApi.md#getRoleMapping) | **GET** /admin/roleMappings/{principalId} | Return a single role mapping
[**importData**](AdminApi.md#importData) | **POST** /admin/import | Import registry data
[**listArtifactTypes**](AdminApi.md#listArtifactTypes) | **GET** /admin/artifactTypes | List artifact types
[**listConfigProperties**](AdminApi.md#listConfigProperties) | **GET** /admin/config/properties | List all configuration properties
[**listGlobalRules**](AdminApi.md#listGlobalRules) | **GET** /admin/rules | List global rules
[**listLogConfigurations**](AdminApi.md#listLogConfigurations) | **GET** /admin/loggers | List logging configurations
[**listRoleMappings**](AdminApi.md#listRoleMappings) | **GET** /admin/roleMappings | List all role mappings
[**removeLogConfiguration**](AdminApi.md#removeLogConfiguration) | **DELETE** /admin/loggers/{logger} | Removes logger configuration
[**resetConfigProperty**](AdminApi.md#resetConfigProperty) | **DELETE** /admin/config/properties/{propertyName} | Reset a configuration property
[**setLogConfiguration**](AdminApi.md#setLogConfiguration) | **PUT** /admin/loggers/{logger} | Set a logger&#39;s configuration
[**updateConfigProperty**](AdminApi.md#updateConfigProperty) | **PUT** /admin/config/properties/{propertyName} | Update a configuration property
[**updateGlobalRuleConfig**](AdminApi.md#updateGlobalRuleConfig) | **PUT** /admin/rules/{rule} | Update global rule configuration
[**updateRoleMapping**](AdminApi.md#updateRoleMapping) | **PUT** /admin/roleMappings/{principalId} | Update a role mapping



## createGlobalRule

> createGlobalRule(rule)

Create global rule

Adds a rule to the list of globally configured rules.  This operation can fail for the following reasons:  * The rule type is unknown (HTTP error &#x60;400&#x60;) * The rule already exists (HTTP error &#x60;409&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        Rule rule = new Rule(); // Rule | 
        try {
            apiInstance.createGlobalRule(rule);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#createGlobalRule");
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
 **rule** | [**Rule**](Rule.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | The global rule was added. |  -  |
| **400** | Common response for all operations that can return a &#x60;400&#x60; error. |  -  |
| **409** | Common response used when an input conflicts with existing data. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## createRoleMapping

> createRoleMapping(roleMapping)

Create a new role mapping

Creates a new mapping between a user/principal and a role.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;)  

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        RoleMapping roleMapping = new RoleMapping(); // RoleMapping | 
        try {
            apiInstance.createRoleMapping(roleMapping);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#createRoleMapping");
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
 **roleMapping** | [**RoleMapping**](RoleMapping.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Returned when the role mapping was successfully created. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## deleteAllGlobalRules

> deleteAllGlobalRules()

Delete all global rules

Deletes all globally configured rules.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            apiInstance.deleteAllGlobalRules();
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#deleteAllGlobalRules");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

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
| **204** | All global rules have been removed successfully. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## deleteGlobalRule

> deleteGlobalRule(rule)

Delete global rule

Deletes a single global rule.  If this is the only rule configured, this is the same as deleting **all** rules.  This operation can fail for the following reasons:  * Invalid rule name/type (HTTP error &#x60;400&#x60;) * No rule with name/type &#x60;rule&#x60; exists (HTTP error &#x60;404&#x60;) * Rule cannot be deleted (HTTP error &#x60;409&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        RuleType rule = RuleType.fromValue("VALIDITY"); // RuleType | The unique name/type of a rule.
        try {
            apiInstance.deleteGlobalRule(rule);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#deleteGlobalRule");
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
 **rule** | [**RuleType**](.md)| The unique name/type of a rule. | [enum: VALIDITY, COMPATIBILITY]

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
| **204** | The global rule was successfully deleted. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## deleteRoleMapping

> deleteRoleMapping(principalId)

Delete a role mapping

Deletes a single role mapping, effectively denying access to a user/principal.  This operation can fail for the following reasons:  * No role mapping for the principalId exists (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String principalId = "principalId_example"; // String | Unique id of a principal (typically either a user or service account).
        try {
            apiInstance.deleteRoleMapping(principalId);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#deleteRoleMapping");
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
 **principalId** | **String**| Unique id of a principal (typically either a user or service account). |

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
| **204** | Response returned when the delete was successful. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## exportData

> File exportData(forBrowser)

Export registry data

Exports registry data as a ZIP archive.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        Boolean forBrowser = true; // Boolean | Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called `href`.  This `href` will be a single-use, naked download link suitable for use by a web browser to download the content.
        try {
            File result = apiInstance.exportData(forBrowser);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#exportData");
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
 **forBrowser** | **Boolean**| Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called &#x60;href&#x60;.  This &#x60;href&#x60; will be a single-use, naked download link suitable for use by a web browser to download the content. | [optional]

### Return type

[**File**](File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Response when the export is successful. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## getConfigProperty

> ConfigurationProperty getConfigProperty(propertyName)

Get configuration property value

Returns the value of a single configuration property.  This operation may fail for one of the following reasons:  * Property not found or not configured (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String propertyName = "propertyName_example"; // String | The name of a configuration property.
        try {
            ConfigurationProperty result = apiInstance.getConfigProperty(propertyName);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#getConfigProperty");
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
 **propertyName** | **String**| The name of a configuration property. |

### Return type

[**ConfigurationProperty**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The configuration property value. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## getGlobalRuleConfig

> Rule getGlobalRuleConfig(rule)

Get global rule configuration

Returns information about the named globally configured rule.  This operation can fail for the following reasons:  * Invalid rule name/type (HTTP error &#x60;400&#x60;) * No rule with name/type &#x60;rule&#x60; exists (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        RuleType rule = RuleType.fromValue("VALIDITY"); // RuleType | The unique name/type of a rule.
        try {
            Rule result = apiInstance.getGlobalRuleConfig(rule);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#getGlobalRuleConfig");
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
 **rule** | [**RuleType**](.md)| The unique name/type of a rule. | [enum: VALIDITY, COMPATIBILITY]

### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The global rule&#39;s configuration. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## getLogConfiguration

> NamedLogConfiguration getLogConfiguration(logger)

Get a single logger configuration

Returns the configured logger configuration for the provided logger name, if no logger configuration is persisted it will return the current default log configuration in the system.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String logger = "logger_example"; // String | The name of a single logger.
        try {
            NamedLogConfiguration result = apiInstance.getLogConfiguration(logger);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#getLogConfiguration");
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
 **logger** | **String**| The name of a single logger. |

### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The logger configuration for the named logger. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## getRoleMapping

> RoleMapping getRoleMapping(principalId)

Return a single role mapping

Gets the details of a single role mapping (by &#x60;principalId&#x60;).  This operation can fail for the following reasons:  * No role mapping for the &#x60;principalId&#x60; exists (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String principalId = "principalId_example"; // String | Unique id of a principal (typically either a user or service account).
        try {
            RoleMapping result = apiInstance.getRoleMapping(principalId);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#getRoleMapping");
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
 **principalId** | **String**| Unique id of a principal (typically either a user or service account). |

### Return type

[**RoleMapping**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | When successful, returns the details of a role mapping. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## importData

> importData(body, xRegistryPreserveGlobalId, xRegistryPreserveContentId)

Import registry data

Imports registry data that was previously exported using the &#x60;/admin/export&#x60; operation.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        File body = new File("/path/to/file"); // File | The ZIP file representing the previously exported registry data.
        Boolean xRegistryPreserveGlobalId = true; // Boolean | If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict.
        Boolean xRegistryPreserveContentId = true; // Boolean | If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict.
        try {
            apiInstance.importData(body, xRegistryPreserveGlobalId, xRegistryPreserveContentId);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#importData");
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
 **body** | **File**| The ZIP file representing the previously exported registry data. |
 **xRegistryPreserveGlobalId** | **Boolean**| If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict. | [optional]
 **xRegistryPreserveContentId** | **Boolean**| If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict. | [optional]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/zip
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Indicates that the import was successful. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listArtifactTypes

> List&lt;ArtifactTypeInfo&gt; listArtifactTypes()

List artifact types

Gets a list of all the configured artifact types.  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            List<ArtifactTypeInfo> result = apiInstance.listArtifactTypes();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#listArtifactTypes");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**List&lt;ArtifactTypeInfo&gt;**](ArtifactTypeInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The list of available artifact types. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listConfigProperties

> List&lt;ConfigurationProperty&gt; listConfigProperties()

List all configuration properties

Returns a list of all configuration properties that have been set.  The list is not paged.  This operation may fail for one of the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            List<ConfigurationProperty> result = apiInstance.listConfigProperties();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#listConfigProperties");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**List&lt;ConfigurationProperty&gt;**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | On a successful response, returns a list of configuration properties. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listGlobalRules

> List&lt;RuleType&gt; listGlobalRules()

List global rules

Gets a list of all the currently configured global rules (if any).  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            List<RuleType> result = apiInstance.listGlobalRules();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#listGlobalRules");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**List&lt;RuleType&gt;**](RuleType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The list of names of the globally configured rules. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listLogConfigurations

> List&lt;NamedLogConfiguration&gt; listLogConfigurations()

List logging configurations

List all of the configured logging levels.  These override the default logging configuration.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            List<NamedLogConfiguration> result = apiInstance.listLogConfigurations();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#listLogConfigurations");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**List&lt;NamedLogConfiguration&gt;**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The list of logging configurations. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## listRoleMappings

> List&lt;RoleMapping&gt; listRoleMappings()

List all role mappings

Gets a list of all role mappings configured in the registry (if any).  This operation can fail for the following reasons:  * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        try {
            List<RoleMapping> result = apiInstance.listRoleMappings();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#listRoleMappings");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**List&lt;RoleMapping&gt;**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | A successful response will return the list of role mappings. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## removeLogConfiguration

> NamedLogConfiguration removeLogConfiguration(logger)

Removes logger configuration

Removes the configured logger configuration (if any) for the given logger.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String logger = "logger_example"; // String | The name of a single logger.
        try {
            NamedLogConfiguration result = apiInstance.removeLogConfiguration(logger);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#removeLogConfiguration");
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
 **logger** | **String**| The name of a single logger. |

### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The default logger configuration (now that the configuration for this logger has been removed, the  default configuration is applied). |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## resetConfigProperty

> resetConfigProperty(propertyName)

Reset a configuration property

Resets the value of a single configuration property.  This will return the property to its default value (see external documentation for supported properties and their default values).  This operation may fail for one of the following reasons:  * Property not found or not configured (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String propertyName = "propertyName_example"; // String | The name of a configuration property.
        try {
            apiInstance.resetConfigProperty(propertyName);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#resetConfigProperty");
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
 **propertyName** | **String**| The name of a configuration property. |

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
| **204** | The configuration property was deleted. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## setLogConfiguration

> NamedLogConfiguration setLogConfiguration(logger, logConfiguration)

Set a logger&#39;s configuration

Configures the logger referenced by the provided logger name with the given configuration.

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String logger = "logger_example"; // String | The name of a single logger.
        LogConfiguration logConfiguration = new LogConfiguration(); // LogConfiguration | The new logger configuration.
        try {
            NamedLogConfiguration result = apiInstance.setLogConfiguration(logger, logConfiguration);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#setLogConfiguration");
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
 **logger** | **String**| The name of a single logger. |
 **logConfiguration** | [**LogConfiguration**](LogConfiguration.md)| The new logger configuration. |

### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The new configuration for the given logger. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## updateConfigProperty

> updateConfigProperty(propertyName, updateConfigurationProperty)

Update a configuration property

Updates the value of a single configuration property.  This operation may fail for one of the following reasons:  * Property not found or not configured (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String propertyName = "propertyName_example"; // String | The name of a configuration property.
        UpdateConfigurationProperty updateConfigurationProperty = new UpdateConfigurationProperty(); // UpdateConfigurationProperty | 
        try {
            apiInstance.updateConfigProperty(propertyName, updateConfigurationProperty);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#updateConfigProperty");
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
 **propertyName** | **String**| The name of a configuration property. |
 **updateConfigurationProperty** | [**UpdateConfigurationProperty**](UpdateConfigurationProperty.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | The configuration property was updated. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## updateGlobalRuleConfig

> Rule updateGlobalRuleConfig(rule, rule2)

Update global rule configuration

Updates the configuration for a globally configured rule.  This operation can fail for the following reasons:  * Invalid rule name/type (HTTP error &#x60;400&#x60;) * No rule with name/type &#x60;rule&#x60; exists (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        RuleType rule = RuleType.fromValue("VALIDITY"); // RuleType | The unique name/type of a rule.
        Rule rule2 = new Rule(); // Rule | 
        try {
            Rule result = apiInstance.updateGlobalRuleConfig(rule, rule2);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#updateGlobalRuleConfig");
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
 **rule** | [**RuleType**](.md)| The unique name/type of a rule. | [enum: VALIDITY, COMPATIBILITY]
 **rule2** | [**Rule**](Rule.md)|  |

### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | The global rule&#39;s configuration was successfully updated. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |


## updateRoleMapping

> updateRoleMapping(principalId, updateRole)

Update a role mapping

Updates a single role mapping for one user/principal.  This operation can fail for the following reasons:  * No role mapping for the principalId exists (HTTP error &#x60;404&#x60;) * A server error occurred (HTTP error &#x60;500&#x60;) 

### Example

```java
// Import classes:
import com.openshift.cloud.api.registry.instance.invoker.ApiClient;
import com.openshift.cloud.api.registry.instance.invoker.ApiException;
import com.openshift.cloud.api.registry.instance.invoker.Configuration;
import com.openshift.cloud.api.registry.instance.invoker.models.*;
import com.openshift.cloud.api.registry.instance.AdminApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AdminApi apiInstance = new AdminApi(defaultClient);
        String principalId = "principalId_example"; // String | Unique id of a principal (typically either a user or service account).
        UpdateRole updateRole = new UpdateRole(); // UpdateRole | 
        try {
            apiInstance.updateRoleMapping(principalId, updateRole);
        } catch (ApiException e) {
            System.err.println("Exception when calling AdminApi#updateRoleMapping");
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
 **principalId** | **String**| Unique id of a principal (typically either a user or service account). |
 **updateRole** | [**UpdateRole**](UpdateRole.md)|  |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Response when the update is successful. |  -  |
| **404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
| **500** | Common response for all operations that can fail with an unexpected server error. |  -  |

