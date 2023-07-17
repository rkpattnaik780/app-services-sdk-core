package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ContentIdsItemReferencesRequestBuilder builds and executes requests for operations under \ids\contentIds\{contentId}\references
type ContentIdsItemReferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ContentIdsItemReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentIdsItemReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContentIdsItemReferencesRequestBuilderInternal instantiates a new ReferencesRequestBuilder and sets the default values.
func NewContentIdsItemReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentIdsItemReferencesRequestBuilder) {
    m := &ContentIdsItemReferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/ids/contentIds/{contentId}/references", pathParameters),
    }
    return m
}
// NewContentIdsItemReferencesRequestBuilder instantiates a new ReferencesRequestBuilder and sets the default values.
func NewContentIdsItemReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentIdsItemReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentIdsItemReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list containing all the artifact references using the artifact content ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ContentIdsItemReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentIdsItemReferencesRequestBuilderGetRequestConfiguration)([]ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactReferenceable, error) {
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
// ToGetRequestInformation returns a list containing all the artifact references using the artifact content ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ContentIdsItemReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContentIdsItemReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
