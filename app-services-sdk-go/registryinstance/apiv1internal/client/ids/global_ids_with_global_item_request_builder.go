package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// GlobalIdsWithGlobalItemRequestBuilder access artifact content utilizing an artifact version's globally unique identifier.
type GlobalIdsWithGlobalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GlobalIdsWithGlobalItemRequestBuilderGetQueryParameters gets the content for an artifact version in the registry using its globally uniqueidentifier.This operation may fail for one of the following reasons:* No artifact version with this `globalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
type GlobalIdsWithGlobalItemRequestBuilderGetQueryParameters struct {
    // Allows the user to specify if the content should be dereferenced when being returned
    Dereference *bool
}
// GlobalIdsWithGlobalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GlobalIdsWithGlobalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GlobalIdsWithGlobalItemRequestBuilderGetQueryParameters
}
// NewGlobalIdsWithGlobalItemRequestBuilderInternal instantiates a new WithGlobalItemRequestBuilder and sets the default values.
func NewGlobalIdsWithGlobalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsWithGlobalItemRequestBuilder) {
    m := &GlobalIdsWithGlobalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/ids/globalIds/{globalId}{?dereference*}", pathParameters),
    }
    return m
}
// NewGlobalIdsWithGlobalItemRequestBuilder instantiates a new WithGlobalItemRequestBuilder and sets the default values.
func NewGlobalIdsWithGlobalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsWithGlobalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGlobalIdsWithGlobalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the content for an artifact version in the registry using its globally uniqueidentifier.This operation may fail for one of the following reasons:* No artifact version with this `globalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *GlobalIdsWithGlobalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GlobalIdsWithGlobalItemRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *GlobalIdsWithGlobalItemRequestBuilder) References()(*GlobalIdsItemReferencesRequestBuilder) {
    return NewGlobalIdsItemReferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets the content for an artifact version in the registry using its globally uniqueidentifier.This operation may fail for one of the following reasons:* No artifact version with this `globalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *GlobalIdsWithGlobalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GlobalIdsWithGlobalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
