package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ContentHashesWithContentHashItemRequestBuilder access artifact content utilizing the SHA-256 hash of the content.
type ContentHashesWithContentHashItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ContentHashesWithContentHashItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentHashesWithContentHashItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContentHashesWithContentHashItemRequestBuilderInternal instantiates a new WithContentHashItemRequestBuilder and sets the default values.
func NewContentHashesWithContentHashItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesWithContentHashItemRequestBuilder) {
    m := &ContentHashesWithContentHashItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/ids/contentHashes/{contentHash}", pathParameters),
    }
    return m
}
// NewContentHashesWithContentHashItemRequestBuilder instantiates a new WithContentHashItemRequestBuilder and sets the default values.
func NewContentHashesWithContentHashItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesWithContentHashItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentHashesWithContentHashItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the content for an artifact version in the registry using the SHA-256 hash of the content.  This content hash may be shared by multiple artifactversions in the case where the artifact versions have identical content.This operation may fail for one of the following reasons:* No content with this `contentHash` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ContentHashesWithContentHashItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentHashesWithContentHashItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// References the references property
func (m *ContentHashesWithContentHashItemRequestBuilder) References()(*ContentHashesItemReferencesRequestBuilder) {
    return NewContentHashesItemReferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets the content for an artifact version in the registry using the SHA-256 hash of the content.  This content hash may be shared by multiple artifactversions in the case where the artifact versions have identical content.This operation may fail for one of the following reasons:* No content with this `contentHash` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ContentHashesWithContentHashItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContentHashesWithContentHashItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
