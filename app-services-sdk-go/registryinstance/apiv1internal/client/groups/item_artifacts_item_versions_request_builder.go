package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ItemArtifactsItemVersionsRequestBuilder manage all the versions of an artifact in the registry.
type ItemArtifactsItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemArtifactsItemVersionsRequestBuilderGetQueryParameters returns a list of all versions of the artifact.  The result set is paged.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
type ItemArtifactsItemVersionsRequestBuilderGetQueryParameters struct {
    // The number of versions to return.  Defaults to 20.
    Limit *int32
    // The number of versions to skip before starting to collect the result set.  Defaults to 0.
    Offset *int32
}
// ItemArtifactsItemVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemArtifactsItemVersionsRequestBuilderGetQueryParameters
}
// ItemArtifactsItemVersionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemVersionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVersion manage a single version of a single artifact in the registry.
func (m *ItemArtifactsItemVersionsRequestBuilder) ByVersion(version string)(*ItemArtifactsItemVersionsWithVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if version != "" {
        urlTplParams["version"] = version
    }
    return NewItemArtifactsItemVersionsWithVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemArtifactsItemVersionsRequestBuilderInternal instantiates a new VersionsRequestBuilder and sets the default values.
func NewItemArtifactsItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemVersionsRequestBuilder) {
    m := &ItemArtifactsItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{groupId}/artifacts/{artifactId}/versions{?offset*,limit*}", pathParameters),
    }
    return m
}
// NewItemArtifactsItemVersionsRequestBuilder instantiates a new VersionsRequestBuilder and sets the default values.
func NewItemArtifactsItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemArtifactsItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of all versions of the artifact.  The result set is paged.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemArtifactsItemVersionsRequestBuilderGetRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.VersionSearchResultsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateVersionSearchResultsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.VersionSearchResultsable), nil
}
// Post creates a new version of the artifact by uploading new content.  The configured rules forthe artifact are applied, and if they all pass, the new content is added as the most recent version of the artifact.  If any of the rules fail, an error is returned.The body of the request can be the raw content of the new artifact version, or the raw content and a set of references pointing to other artifacts, and the typeof that content should match the artifact's type (for example if the artifact type is `AVRO`then the content of the request should be an Apache Avro document).This operation can fail for the following reasons:* Provided content (request body) was empty (HTTP error `400`)* No artifact with this `artifactId` exists (HTTP error `404`)* The new content violates one of the rules configured for the artifact (HTTP error `409`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemVersionsRequestBuilder) Post(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactContentable, requestConfiguration *ItemArtifactsItemVersionsRequestBuilderPostRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.VersionMetaDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "409": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateVersionMetaDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.VersionMetaDataable), nil
}
// ToGetRequestInformation returns a list of all versions of the artifact.  The result set is paged.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsItemVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a new version of the artifact by uploading new content.  The configured rules forthe artifact are applied, and if they all pass, the new content is added as the most recent version of the artifact.  If any of the rules fail, an error is returned.The body of the request can be the raw content of the new artifact version, or the raw content and a set of references pointing to other artifacts, and the typeof that content should match the artifact's type (for example if the artifact type is `AVRO`then the content of the request should be an Apache Avro document).This operation can fail for the following reasons:* Provided content (request body) was empty (HTTP error `400`)* No artifact with this `artifactId` exists (HTTP error `404`)* The new content violates one of the rules configured for the artifact (HTTP error `409`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemVersionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactContentable, requestConfiguration *ItemArtifactsItemVersionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
