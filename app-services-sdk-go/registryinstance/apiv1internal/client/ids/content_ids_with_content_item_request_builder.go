package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ContentIdsWithContentItemRequestBuilder access artifact content utilizing the unique content identifier for that content.
type ContentIdsWithContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ContentIdsWithContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentIdsWithContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContentIdsWithContentItemRequestBuilderInternal instantiates a new WithContentItemRequestBuilder and sets the default values.
func NewContentIdsWithContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentIdsWithContentItemRequestBuilder) {
    m := &ContentIdsWithContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/ids/contentIds/{contentId}", pathParameters),
    }
    return m
}
// NewContentIdsWithContentItemRequestBuilder instantiates a new WithContentItemRequestBuilder and sets the default values.
func NewContentIdsWithContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentIdsWithContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentIdsWithContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the content for an artifact version in the registry using the unique contentidentifier for that content.  This content ID may be shared by multiple artifactversions in the case where the artifact versions are identical.This operation may fail for one of the following reasons:* No content with this `contentId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ContentIdsWithContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentIdsWithContentItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *ContentIdsWithContentItemRequestBuilder) References()(*ContentIdsItemReferencesRequestBuilder) {
    return NewContentIdsItemReferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets the content for an artifact version in the registry using the unique contentidentifier for that content.  This content ID may be shared by multiple artifactversions in the case where the artifact versions are identical.This operation may fail for one of the following reasons:* No content with this `contentId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ContentIdsWithContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContentIdsWithContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
