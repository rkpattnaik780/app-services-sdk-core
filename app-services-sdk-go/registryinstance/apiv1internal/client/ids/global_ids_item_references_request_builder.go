package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// GlobalIdsItemReferencesRequestBuilder builds and executes requests for operations under \ids\globalIds\{globalId}\references
type GlobalIdsItemReferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GlobalIdsItemReferencesRequestBuilderGetQueryParameters returns a list containing all the artifact references using the artifact global ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
type GlobalIdsItemReferencesRequestBuilderGetQueryParameters struct {
    // Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND.
    RefType *string
}
// GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GlobalIdsItemReferencesRequestBuilderGetQueryParameters
}
// NewGlobalIdsItemReferencesRequestBuilderInternal instantiates a new ReferencesRequestBuilder and sets the default values.
func NewGlobalIdsItemReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsItemReferencesRequestBuilder) {
    m := &GlobalIdsItemReferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/ids/globalIds/{globalId}/references{?refType*}", pathParameters),
    }
    return m
}
// NewGlobalIdsItemReferencesRequestBuilder instantiates a new ReferencesRequestBuilder and sets the default values.
func NewGlobalIdsItemReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsItemReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGlobalIdsItemReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list containing all the artifact references using the artifact global ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *GlobalIdsItemReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration)([]ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateArtifactReferenceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactReferenceable, len(res))
    for i, v := range res {
        val[i] = v.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactReferenceable)
    }
    return val, nil
}
// ToGetRequestInformation returns a list containing all the artifact references using the artifact global ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *GlobalIdsItemReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
