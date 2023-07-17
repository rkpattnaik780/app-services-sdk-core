package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ItemArtifactsRequestBuilder manage the collection of artifacts within a single group in the registry.
type ItemArtifactsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemArtifactsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemArtifactsRequestBuilderGetQueryParameters returns a list of all artifacts in the group.  This list is paged.
type ItemArtifactsRequestBuilderGetQueryParameters struct {
    // The number of artifacts to return.  Defaults to 20.
    Limit *int32
    // The number of artifacts to skip before starting the result set.  Defaults to 0.
    Offset *int32
    // Sort order, ascending (`asc`) or descending (`desc`).
    Order *string
    // The field to sort by.  Can be one of:* `name`* `createdOn`
    Orderby *string
}
// ItemArtifactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemArtifactsRequestBuilderGetQueryParameters
}
// ItemArtifactsRequestBuilderPostQueryParameters creates a new artifact by posting the artifact content.  The body of the request shouldbe the raw content of the artifact.  This is typically in JSON format for *most* of the supported types, but may be in another format for a few (for example, `PROTOBUF`).The registry attempts to figure out what kind of artifact is being added from thefollowing supported list:* Avro (`AVRO`)* Protobuf (`PROTOBUF`)* JSON Schema (`JSON`)* Kafka Connect (`KCONNECT`)* OpenAPI (`OPENAPI`)* AsyncAPI (`ASYNCAPI`)* GraphQL (`GRAPHQL`)* Web Services Description Language (`WSDL`)* XML Schema (`XSD`)Alternatively, you can specify the artifact type using the `X-Registry-ArtifactType` HTTP request header, or include a hint in the request's `Content-Type`.  For example:```Content-Type: application/json; artifactType=AVRO```An artifact is created using the content provided in the body of the request.  Thiscontent is created under a unique artifact ID that can be provided in the requestusing the `X-Registry-ArtifactId` request header.  If not provided in the request,the server generates a unique ID for the artifact.  It is typically recommendedthat callers provide the ID, because this is typically a meaningful identifier, and for most use cases should be supplied by the caller.If an artifact with the provided artifact ID already exists, the default behavioris for the server to reject the content with a 409 error.  However, the caller cansupply the `ifExists` query parameter to alter this default behavior. The `ifExists`query parameter can have one of the following values:* `FAIL` (*default*) - server rejects the content with a 409 error* `UPDATE` - server updates the existing artifact and returns the new metadata* `RETURN` - server does not create or add content to the server, but instead returns the metadata for the existing artifact* `RETURN_OR_UPDATE` - server returns an existing **version** that matches the provided content if such a version exists, otherwise a new version is createdThis operation may fail for one of the following reasons:* An invalid `ArtifactType` was indicated (HTTP error `400`)* No `ArtifactType` was indicated and the server could not determine one from the content (HTTP error `400`)* Provided content (request body) was empty (HTTP error `400`)* An artifact with the provided ID already exists (HTTP error `409`)* The content violates one of the configured global rules (HTTP error `409`)* A server error occurred (HTTP error `500`)
type ItemArtifactsRequestBuilderPostQueryParameters struct {
    // Used only when the `ifExists` query parameter is set to `RETURN_OR_UPDATE`, this parameter can be set to `true` to indicate that the server should "canonicalize" the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner.
    Canonical *bool
    // Set this option to instruct the server on what to do if the artifact already exists.
    IfExists *string
}
// ItemArtifactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemArtifactsRequestBuilderPostQueryParameters
}
// ByArtifactId manage a single artifact.
func (m *ItemArtifactsRequestBuilder) ByArtifactId(artifactId string)(*ItemArtifactsWithArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if artifactId != "" {
        urlTplParams["artifactId"] = artifactId
    }
    return NewItemArtifactsWithArtifactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemArtifactsRequestBuilderInternal instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewItemArtifactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsRequestBuilder) {
    m := &ItemArtifactsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{groupId}/artifacts{?limit*,offset*,order*,orderby*,ifExists*,canonical*}", pathParameters),
    }
    return m
}
// NewItemArtifactsRequestBuilder instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewItemArtifactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemArtifactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes all of the artifacts that exist in a given group.
func (m *ItemArtifactsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemArtifactsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns a list of all artifacts in the group.  This list is paged.
func (m *ItemArtifactsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemArtifactsRequestBuilderGetRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactSearchResultsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateArtifactSearchResultsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactSearchResultsable), nil
}
// Post creates a new artifact by posting the artifact content.  The body of the request shouldbe the raw content of the artifact.  This is typically in JSON format for *most* of the supported types, but may be in another format for a few (for example, `PROTOBUF`).The registry attempts to figure out what kind of artifact is being added from thefollowing supported list:* Avro (`AVRO`)* Protobuf (`PROTOBUF`)* JSON Schema (`JSON`)* Kafka Connect (`KCONNECT`)* OpenAPI (`OPENAPI`)* AsyncAPI (`ASYNCAPI`)* GraphQL (`GRAPHQL`)* Web Services Description Language (`WSDL`)* XML Schema (`XSD`)Alternatively, you can specify the artifact type using the `X-Registry-ArtifactType` HTTP request header, or include a hint in the request's `Content-Type`.  For example:```Content-Type: application/json; artifactType=AVRO```An artifact is created using the content provided in the body of the request.  Thiscontent is created under a unique artifact ID that can be provided in the requestusing the `X-Registry-ArtifactId` request header.  If not provided in the request,the server generates a unique ID for the artifact.  It is typically recommendedthat callers provide the ID, because this is typically a meaningful identifier, and for most use cases should be supplied by the caller.If an artifact with the provided artifact ID already exists, the default behavioris for the server to reject the content with a 409 error.  However, the caller cansupply the `ifExists` query parameter to alter this default behavior. The `ifExists`query parameter can have one of the following values:* `FAIL` (*default*) - server rejects the content with a 409 error* `UPDATE` - server updates the existing artifact and returns the new metadata* `RETURN` - server does not create or add content to the server, but instead returns the metadata for the existing artifact* `RETURN_OR_UPDATE` - server returns an existing **version** that matches the provided content if such a version exists, otherwise a new version is createdThis operation may fail for one of the following reasons:* An invalid `ArtifactType` was indicated (HTTP error `400`)* No `ArtifactType` was indicated and the server could not determine one from the content (HTTP error `400`)* Provided content (request body) was empty (HTTP error `400`)* An artifact with the provided ID already exists (HTTP error `409`)* The content violates one of the configured global rules (HTTP error `409`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsRequestBuilder) Post(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactContentable, requestConfiguration *ItemArtifactsRequestBuilderPostRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactMetaDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "409": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateArtifactMetaDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactMetaDataable), nil
}
// ToDeleteRequestInformation deletes all of the artifacts that exist in a given group.
func (m *ItemArtifactsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of all artifacts in the group.  This list is paged.
func (m *ItemArtifactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation creates a new artifact by posting the artifact content.  The body of the request shouldbe the raw content of the artifact.  This is typically in JSON format for *most* of the supported types, but may be in another format for a few (for example, `PROTOBUF`).The registry attempts to figure out what kind of artifact is being added from thefollowing supported list:* Avro (`AVRO`)* Protobuf (`PROTOBUF`)* JSON Schema (`JSON`)* Kafka Connect (`KCONNECT`)* OpenAPI (`OPENAPI`)* AsyncAPI (`ASYNCAPI`)* GraphQL (`GRAPHQL`)* Web Services Description Language (`WSDL`)* XML Schema (`XSD`)Alternatively, you can specify the artifact type using the `X-Registry-ArtifactType` HTTP request header, or include a hint in the request's `Content-Type`.  For example:```Content-Type: application/json; artifactType=AVRO```An artifact is created using the content provided in the body of the request.  Thiscontent is created under a unique artifact ID that can be provided in the requestusing the `X-Registry-ArtifactId` request header.  If not provided in the request,the server generates a unique ID for the artifact.  It is typically recommendedthat callers provide the ID, because this is typically a meaningful identifier, and for most use cases should be supplied by the caller.If an artifact with the provided artifact ID already exists, the default behavioris for the server to reject the content with a 409 error.  However, the caller cansupply the `ifExists` query parameter to alter this default behavior. The `ifExists`query parameter can have one of the following values:* `FAIL` (*default*) - server rejects the content with a 409 error* `UPDATE` - server updates the existing artifact and returns the new metadata* `RETURN` - server does not create or add content to the server, but instead returns the metadata for the existing artifact* `RETURN_OR_UPDATE` - server returns an existing **version** that matches the provided content if such a version exists, otherwise a new version is createdThis operation may fail for one of the following reasons:* An invalid `ArtifactType` was indicated (HTTP error `400`)* No `ArtifactType` was indicated and the server could not determine one from the content (HTTP error `400`)* Provided content (request body) was empty (HTTP error `400`)* An artifact with the provided ID already exists (HTTP error `409`)* The content violates one of the configured global rules (HTTP error `409`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactContentable, requestConfiguration *ItemArtifactsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
