# rhoas_registry_instance_sdk.ArtifactsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_artifact**](ArtifactsApi.md#create_artifact) | **POST** /groups/{groupId}/artifacts | Create artifact
[**delete_artifact**](ArtifactsApi.md#delete_artifact) | **DELETE** /groups/{groupId}/artifacts/{artifactId} | Delete artifact
[**delete_artifacts_in_group**](ArtifactsApi.md#delete_artifacts_in_group) | **DELETE** /groups/{groupId}/artifacts | Delete artifacts in group
[**get_content_by_global_id**](ArtifactsApi.md#get_content_by_global_id) | **GET** /ids/globalIds/{globalId} | Get artifact by global ID
[**get_content_by_hash**](ArtifactsApi.md#get_content_by_hash) | **GET** /ids/contentHashes/{contentHash}/ | Get artifact content by SHA-256 hash
[**get_content_by_id**](ArtifactsApi.md#get_content_by_id) | **GET** /ids/contentIds/{contentId}/ | Get artifact content by ID
[**get_latest_artifact**](ArtifactsApi.md#get_latest_artifact) | **GET** /groups/{groupId}/artifacts/{artifactId} | Get latest artifact
[**list_artifacts_in_group**](ArtifactsApi.md#list_artifacts_in_group) | **GET** /groups/{groupId}/artifacts | List artifacts in group
[**references_by_content_hash**](ArtifactsApi.md#references_by_content_hash) | **GET** /ids/contentHashes/{contentHash}/references | List artifact references by hash
[**references_by_content_id**](ArtifactsApi.md#references_by_content_id) | **GET** /ids/contentIds/{contentId}/references | List artifact references by content ID
[**references_by_global_id**](ArtifactsApi.md#references_by_global_id) | **GET** /ids/globalIds/{globalId}/references | List artifact references by global ID
[**update_artifact**](ArtifactsApi.md#update_artifact) | **PUT** /groups/{groupId}/artifacts/{artifactId} | Update artifact
[**update_artifact_state**](ArtifactsApi.md#update_artifact_state) | **PUT** /groups/{groupId}/artifacts/{artifactId}/state | Update artifact state


# **create_artifact**
> ArtifactMetaData create_artifact(group_id, body)

Create artifact

Creates a new artifact by posting the artifact content.  The body of the request should be the raw content of the artifact.  This is typically in JSON format for *most* of the  supported types, but may be in another format for a few (for example, `PROTOBUF`).  The registry attempts to figure out what kind of artifact is being added from the following supported list:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`)  Alternatively, you can specify the artifact type using the `X-Registry-ArtifactType`  HTTP request header, or include a hint in the request's `Content-Type`.  For example:  ``` Content-Type: application/json; artifactType=AVRO ```  An artifact is created using the content provided in the body of the request.  This content is created under a unique artifact ID that can be provided in the request using the `X-Registry-ArtifactId` request header.  If not provided in the request, the server generates a unique ID for the artifact.  It is typically recommended that callers provide the ID, because this is typically a meaningful identifier,  and for most use cases should be supplied by the caller.  If an artifact with the provided artifact ID already exists, the default behavior is for the server to reject the content with a 409 error.  However, the caller can supply the `ifExists` query parameter to alter this default behavior. The `ifExists` query parameter can have one of the following values:  * `FAIL` (*default*) - server rejects the content with a 409 error * `UPDATE` - server updates the existing artifact and returns the new metadata * `RETURN` - server does not create or add content to the server, but instead  returns the metadata for the existing artifact * `RETURN_OR_UPDATE` - server returns an existing **version** that matches the  provided content if such a version exists, otherwise a new version is created  This operation may fail for one of the following reasons:  * An invalid `ArtifactType` was indicated (HTTP error `400`) * No `ArtifactType` was indicated and the server could not determine one from the content (HTTP error `400`) * Provided content (request body) was empty (HTTP error `400`) * An artifact with the provided ID already exists (HTTP error `409`) * The content violates one of the configured global rules (HTTP error `409`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.rule_violation_error import RuleViolationError
from rhoas_registry_instance_sdk.model.if_exists import IfExists
from rhoas_registry_instance_sdk.model.error import Error
from rhoas_registry_instance_sdk.model.artifact_content import ArtifactContent
from rhoas_registry_instance_sdk.model.artifact_meta_data import ArtifactMetaData
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    body = open('{"openapi":"3.0.2","info":{"title":"Empty API","version":"1.0.7","description":"An example API design using OpenAPI."},"paths":{"/widgets":{"get":{"responses":{"200":{"content":{"application/json":{"schema":{"type":"array","items":{"type":"string"}}}},"description":"All widgets"}},"summary":"Get widgets"}}},"components":{"schemas":{"Widget":{"title":"Root Type for Widget","description":"A sample data type.","type":"object","properties":{"property-1":{"type":"string"},"property-2":{"type":"boolean"}},"example":{"property-1":"value1","property-2":true}}}}}', 'rb') # file_type | The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    x_registry_artifact_type = "AVRO" # str | Specifies the type of the artifact being added. Possible values include:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) (optional)
    x_registry_artifact_id = "X-Registry-ArtifactId_example" # str | A client-provided, globally unique identifier for the new artifact. (optional)
    x_registry_version = "3.1.6" # str | Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version `1`). (optional)
    if_exists = IfExists("FAIL") # IfExists | Set this option to instruct the server on what to do if the artifact already exists. (optional)
    canonical = True # bool | Used only when the `ifExists` query parameter is set to `RETURN_OR_UPDATE`, this parameter can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. (optional)
    x_registry_description = "Artifact description" # str | Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    x_registry_description_encoded = "QXJ0aWZhY3QgZGVzY3JpcHRpb24K" # str | Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)
    x_registry_name = "Artifact name" # str | Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    x_registry_name_encoded = "QXJ0aWZhY3QgbmFtZQo=" # str | Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    x_registry_content_hash = "X-Registry-Content-Hash_example" # str | Specifies the (optional) hash of the artifact to be verified. (optional)
    x_registry_hash_algorithm = "SHA256" # str | The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) (optional)
    content_type = "Content-Type_example" # str | This header is explicit so clients using the OpenAPI Generator are able select the content type. Ignore otherwise. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create artifact
        api_response = api_instance.create_artifact(group_id, body)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->create_artifact: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create artifact
        api_response = api_instance.create_artifact(group_id, body, x_registry_artifact_type=x_registry_artifact_type, x_registry_artifact_id=x_registry_artifact_id, x_registry_version=x_registry_version, if_exists=if_exists, canonical=canonical, x_registry_description=x_registry_description, x_registry_description_encoded=x_registry_description_encoded, x_registry_name=x_registry_name, x_registry_name_encoded=x_registry_name_encoded, x_registry_content_hash=x_registry_content_hash, x_registry_hash_algorithm=x_registry_hash_algorithm, content_type=content_type)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->create_artifact: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **body** | **file_type**| The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  |
 **x_registry_artifact_type** | **str**| Specifies the type of the artifact being added. Possible values include:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;) | [optional]
 **x_registry_artifact_id** | **str**| A client-provided, globally unique identifier for the new artifact. | [optional]
 **x_registry_version** | **str**| Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version &#x60;1&#x60;). | [optional]
 **if_exists** | **IfExists**| Set this option to instruct the server on what to do if the artifact already exists. | [optional]
 **canonical** | **bool**| Used only when the &#x60;ifExists&#x60; query parameter is set to &#x60;RETURN_OR_UPDATE&#x60;, this parameter can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. | [optional]
 **x_registry_description** | **str**| Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | [optional]
 **x_registry_description_encoded** | **str**| Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | [optional]
 **x_registry_name** | **str**| Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | [optional]
 **x_registry_name_encoded** | **str**| Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | [optional]
 **x_registry_content_hash** | **str**| Specifies the (optional) hash of the artifact to be verified. | [optional]
 **x_registry_hash_algorithm** | **str**| The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) | [optional]
 **content_type** | **str**| This header is explicit so clients using the OpenAPI Generator are able select the content type. Ignore otherwise. | [optional]

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/vnd.json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Artifact was successfully created. |  -  |
**400** | Common response for all operations that can return a &#x60;400&#x60; error. |  -  |
**409** | Common response used when an input conflicts with existing data. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_artifact**
> delete_artifact(group_id, artifact_id)

Delete artifact

Deletes an artifact completely, resulting in all versions of the artifact also being deleted.  This may fail for one of the following reasons:  * No artifact with the `artifactId` exists (HTTP error `404`) * A server error occurred (HTTP error `500`)

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifact_id = "example-artifact" # str | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.

    # example passing only required values which don't have defaults set
    try:
        # Delete artifact
        api_instance.delete_artifact(group_id, artifact_id)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->delete_artifact: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **artifact_id** | **str**| The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. |

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
**204** | Returned when the artifact was successfully deleted. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_artifacts_in_group**
> delete_artifacts_in_group(group_id)

Delete artifacts in group

Deletes all of the artifacts that exist in a given group.

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.

    # example passing only required values which don't have defaults set
    try:
        # Delete artifacts in group
        api_instance.delete_artifacts_in_group(group_id)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->delete_artifacts_in_group: %s\n" % e)
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
**204** | When the delete operation is successful, a simple 204 is returned. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_content_by_global_id**
> file_type get_content_by_global_id(global_id)

Get artifact by global ID

Gets the content for an artifact version in the registry using its globally unique identifier.  This operation may fail for one of the following reasons:  * No artifact version with this `globalId` exists (HTTP error `404`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.handle_references_type import HandleReferencesType
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    global_id = 1 # int | Global identifier for an artifact version.
    references = HandleReferencesType("PRESERVE") # HandleReferencesType | Allows the user to specify how references in the content should be treated. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get artifact by global ID
        api_response = api_instance.get_content_by_global_id(global_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_content_by_global_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get artifact by global ID
        api_response = api_instance.get_content_by_global_id(global_id, references=references)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_content_by_global_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **global_id** | **int**| Global identifier for an artifact version. |
 **references** | **HandleReferencesType**| Allows the user to specify how references in the content should be treated. | [optional]

### Return type

**file_type**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The content of one version of one artifact. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_content_by_hash**
> file_type get_content_by_hash(content_hash)

Get artifact content by SHA-256 hash

Gets the content for an artifact version in the registry using the  SHA-256 hash of the content.  This content hash may be shared by multiple artifact versions in the case where the artifact versions have identical content.  This operation may fail for one of the following reasons:  * No content with this `contentHash` exists (HTTP error `404`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    content_hash = "contentHash_example" # str | SHA-256 content hash for a single artifact content.

    # example passing only required values which don't have defaults set
    try:
        # Get artifact content by SHA-256 hash
        api_response = api_instance.get_content_by_hash(content_hash)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_content_by_hash: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content_hash** | **str**| SHA-256 content hash for a single artifact content. |

### Return type

**file_type**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The content of one version of one artifact. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_content_by_id**
> file_type get_content_by_id(content_id)

Get artifact content by ID

Gets the content for an artifact version in the registry using the unique content identifier for that content.  This content ID may be shared by multiple artifact versions in the case where the artifact versions are identical.  This operation may fail for one of the following reasons:  * No content with this `contentId` exists (HTTP error `404`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    content_id = 1 # int | Global identifier for a single artifact content.

    # example passing only required values which don't have defaults set
    try:
        # Get artifact content by ID
        api_response = api_instance.get_content_by_id(content_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_content_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content_id** | **int**| Global identifier for a single artifact content. |

### Return type

**file_type**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The content of one version of one artifact. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_latest_artifact**
> file_type get_latest_artifact(group_id, artifact_id)

Get latest artifact

Returns the latest version of the artifact in its raw form.  The `Content-Type` of the response depends on the artifact type.  In most cases, this is `application/json`, but  for some types it may be different (for example, `PROTOBUF`). If the latest version of the artifact is marked as `DISABLED`, the next available non-disabled version will be used.  This operation may fail for one of the following reasons:  * No artifact with this `artifactId` exists or all versions are `DISABLED` (HTTP error `404`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.handle_references_type import HandleReferencesType
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifact_id = "example-artifact" # str | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    references = HandleReferencesType("PRESERVE") # HandleReferencesType | Allows the user to specify how references in the content should be treated. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get latest artifact
        api_response = api_instance.get_latest_artifact(group_id, artifact_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_latest_artifact: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get latest artifact
        api_response = api_instance.get_latest_artifact(group_id, artifact_id, references=references)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->get_latest_artifact: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **artifact_id** | **str**| The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. |
 **references** | **HandleReferencesType**| Allows the user to specify how references in the content should be treated. | [optional]

### Return type

**file_type**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The content of one version of one artifact. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_artifacts_in_group**
> ArtifactSearchResults list_artifacts_in_group(group_id)

List artifacts in group

Returns a list of all artifacts in the group.  This list is paged.

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.sort_order import SortOrder
from rhoas_registry_instance_sdk.model.artifact_search_results import ArtifactSearchResults
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    limit = 1 # int | The number of artifacts to return.  Defaults to 20. (optional)
    offset = 1 # int | The number of artifacts to skip before starting the result set.  Defaults to 0. (optional)
    order = SortOrder("asc") # SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby = SortBy("name") # SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    # example passing only required values which don't have defaults set
    try:
        # List artifacts in group
        api_response = api_instance.list_artifacts_in_group(group_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->list_artifacts_in_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List artifacts in group
        api_response = api_instance.list_artifacts_in_group(group_id, limit=limit, offset=offset, order=order, orderby=orderby)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->list_artifacts_in_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **limit** | **int**| The number of artifacts to return.  Defaults to 20. | [optional]
 **offset** | **int**| The number of artifacts to skip before starting the result set.  Defaults to 0. | [optional]
 **order** | **SortOrder**| Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | [optional]
 **orderby** | **SortBy**| The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | [optional]

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | On a successful response, returns a bounded set of artifacts. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **references_by_content_hash**
> [ArtifactReference] references_by_content_hash(content_hash)

List artifact references by hash

Returns a list containing all the artifact references using the artifact content hash.  This operation may fail for one of the following reasons:  * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.artifact_reference import ArtifactReference
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = artifacts_api.ArtifactsApi(api_client)
    content_hash = "contentHash_example" # str | SHA-256 content hash for a single artifact content.

    # example passing only required values which don't have defaults set
    try:
        # List artifact references by hash
        api_response = api_instance.references_by_content_hash(content_hash)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->references_by_content_hash: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content_hash** | **str**| SHA-256 content hash for a single artifact content. |

### Return type

[**[ArtifactReference]**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A list containing all the references for the artifact with the given content hash. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **references_by_content_id**
> [ArtifactReference] references_by_content_id(content_id)

List artifact references by content ID

Returns a list containing all the artifact references using the artifact content ID.  This operation may fail for one of the following reasons:  * A server error occurred (HTTP error `500`)

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.artifact_reference import ArtifactReference
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = artifacts_api.ArtifactsApi(api_client)
    content_id = 1 # int | Global identifier for a single artifact content.

    # example passing only required values which don't have defaults set
    try:
        # List artifact references by content ID
        api_response = api_instance.references_by_content_id(content_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->references_by_content_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content_id** | **int**| Global identifier for a single artifact content. |

### Return type

[**[ArtifactReference]**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A list containing all the references for the artifact with the given content id. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **references_by_global_id**
> [ArtifactReference] references_by_global_id(global_id)

List artifact references by global ID

Returns a list containing all the artifact references using the artifact global ID.  This operation may fail for one of the following reasons:  * A server error occurred (HTTP error `500`)

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.reference_type import ReferenceType
from rhoas_registry_instance_sdk.model.artifact_reference import ArtifactReference
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = artifacts_api.ArtifactsApi(api_client)
    global_id = 1 # int | Global identifier for an artifact version.
    ref_type = ReferenceType("INBOUND") # ReferenceType | Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND. (optional)

    # example passing only required values which don't have defaults set
    try:
        # List artifact references by global ID
        api_response = api_instance.references_by_global_id(global_id)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->references_by_global_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List artifact references by global ID
        api_response = api_instance.references_by_global_id(global_id, ref_type=ref_type)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->references_by_global_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **global_id** | **int**| Global identifier for an artifact version. |
 **ref_type** | **ReferenceType**| Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND. | [optional]

### Return type

[**[ArtifactReference]**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A list containing all the references for the artifact with the given global id. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_artifact**
> ArtifactMetaData update_artifact(group_id, artifact_id, body)

Update artifact

Updates an artifact by uploading new content.  The body of the request can be the raw content of the artifact or a JSON object containing both the raw content and a set of references to other artifacts..  This is typically in JSON format for *most* of the supported types, but may be in another format for a few (for example, `PROTOBUF`). The type of the content should be compatible with the artifact's type (it would be an error to update an `AVRO` artifact with new `OPENAPI` content, for example).  The update could fail for a number of reasons including:  * Provided content (request body) was empty (HTTP error `400`) * No artifact with the `artifactId` exists (HTTP error `404`) * The new content violates one of the rules configured for the artifact (HTTP error `409`) * A server error occurred (HTTP error `500`)  When successful, this creates a new version of the artifact, making it the most recent (and therefore official) version of the artifact.

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.error import Error
from rhoas_registry_instance_sdk.model.artifact_content import ArtifactContent
from rhoas_registry_instance_sdk.model.artifact_meta_data import ArtifactMetaData
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = rhoas_registry_instance_sdk.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with rhoas_registry_instance_sdk.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifact_id = "example-artifact" # str | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    body = open('{"openapi":"3.0.2","info":{"title":"Empty API","version":"1.0.7","description":"An example API design using OpenAPI."},"paths":{"/widgets":{"get":{"responses":{"200":{"content":{"application/json":{"schema":{"type":"array","items":{"type":"string"}}}},"description":"All widgets"}},"summary":"Get widgets"}}},"components":{"schemas":{"Widget":{"title":"Root Type for Widget","description":"A sample data type.","type":"object","properties":{"property-1":{"type":"string"},"property-2":{"type":"boolean"}},"example":{"property-1":"value1","property-2":true}}}}}', 'rb') # file_type | The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    x_registry_version = "3.1.6" # str | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. (optional)
    x_registry_name = "Artifact name" # str | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    x_registry_name_encoded = "QXJ0aWZhY3QgbmFtZQo=" # str | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    x_registry_description = "Artifact description" # str | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    x_registry_description_encoded = "QXJ0aWZhY3QgZGVzY3JpcHRpb24K" # str | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)
    content_type = "Content-Type_example" # str | This header is explicit so clients using the OpenAPI Generator are able select the content type. Ignore otherwise. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update artifact
        api_response = api_instance.update_artifact(group_id, artifact_id, body)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->update_artifact: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update artifact
        api_response = api_instance.update_artifact(group_id, artifact_id, body, x_registry_version=x_registry_version, x_registry_name=x_registry_name, x_registry_name_encoded=x_registry_name_encoded, x_registry_description=x_registry_description, x_registry_description_encoded=x_registry_description_encoded, content_type=content_type)
        pprint(api_response)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->update_artifact: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **artifact_id** | **str**| The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. |
 **body** | **file_type**| The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  |
 **x_registry_version** | **str**| Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. | [optional]
 **x_registry_name** | **str**| Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | [optional]
 **x_registry_name_encoded** | **str**| Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | [optional]
 **x_registry_description** | **str**| Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | [optional]
 **x_registry_description_encoded** | **str**| Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | [optional]
 **content_type** | **str**| This header is explicit so clients using the OpenAPI Generator are able select the content type. Ignore otherwise. | [optional]

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/vnd.json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | When successful, returns the updated artifact metadata. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**409** | Common response used when an input conflicts with existing data. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_artifact_state**
> update_artifact_state(group_id, artifact_id, update_state)

Update artifact state

Updates the state of the artifact.  For example, you can use this to mark the latest version of an artifact as `DEPRECATED`. The operation changes the state of the latest version of the artifact, even if this version is `DISABLED`. If multiple versions exist, only the most recent is changed.  This operation can fail for the following reasons:  * No artifact with this `artifactId` exists (HTTP error `404`) * A server error occurred (HTTP error `500`) 

### Example


```python
import time
import rhoas_registry_instance_sdk
from rhoas_registry_instance_sdk.api import artifacts_api
from rhoas_registry_instance_sdk.model.update_state import UpdateState
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
    api_instance = artifacts_api.ArtifactsApi(api_client)
    group_id = "my-group" # str | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifact_id = "example-artifact" # str | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    update_state = UpdateState(
        state=ArtifactState("ENABLED"),
    ) # UpdateState | 

    # example passing only required values which don't have defaults set
    try:
        # Update artifact state
        api_instance.update_artifact_state(group_id, artifact_id, update_state)
    except rhoas_registry_instance_sdk.ApiException as e:
        print("Exception when calling ArtifactsApi->update_artifact_state: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. |
 **artifact_id** | **str**| The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. |
 **update_state** | [**UpdateState**](UpdateState.md)|  |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Returned when the operation was successful. |  -  |
**400** | Common response for all operations that can return a &#x60;400&#x60; error. |  -  |
**404** | Common response for all operations that can return a &#x60;404&#x60; error. |  -  |
**500** | Common response for all operations that can fail with an unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

