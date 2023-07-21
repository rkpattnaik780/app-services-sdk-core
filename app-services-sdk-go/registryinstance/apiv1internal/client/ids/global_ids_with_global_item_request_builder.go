package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// GlobalIdsWithGlobalItemRequestBuilder access artifact content utilizing an artifact version's globally unique identifier.
type GlobalIdsWithGlobalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GlobalIdsWithGlobalItemRequestBuilderGetQueryParameters gets the content for an artifact version in the registry using its globally uniqueidentifier.This operation may fail for one of the following reasons:* No artifact version with this `globalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
type GlobalIdsWithGlobalItemRequestBuilderGetQueryParameters struct {
    // Allows the user to specify how references in the content should be treated.
    References *string
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
    }
    m.urlTemplate = "{+baseurl}/ids/globalIds/{globalId}{?references*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
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
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
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
    return NewGlobalIdsItemReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation gets the content for an artifact version in the registry using its globally uniqueidentifier.This operation may fail for one of the following reasons:* No artifact version with this `globalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *GlobalIdsWithGlobalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GlobalIdsWithGlobalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
