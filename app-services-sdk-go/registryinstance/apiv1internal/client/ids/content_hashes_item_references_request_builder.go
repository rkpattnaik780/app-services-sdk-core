package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ContentHashesItemReferencesRequestBuilder builds and executes requests for operations under \ids\contentHashes\{contentHash}\references
type ContentHashesItemReferencesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentHashesItemReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentHashesItemReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewContentHashesItemReferencesRequestBuilderInternal instantiates a new ReferencesRequestBuilder and sets the default values.
func NewContentHashesItemReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesItemReferencesRequestBuilder) {
    m := &ContentHashesItemReferencesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/ids/contentHashes/{contentHash}/references";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewContentHashesItemReferencesRequestBuilder instantiates a new ReferencesRequestBuilder and sets the default values.
func NewContentHashesItemReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesItemReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentHashesItemReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list containing all the artifact references using the artifact content hash.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ContentHashesItemReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentHashesItemReferencesRequestBuilderGetRequestConfiguration)([]i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateArtifactReferenceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactReferenceable, len(res))
    for i, v := range res {
        val[i] = v.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactReferenceable)
    }
    return val, nil
}
// ToGetRequestInformation returns a list containing all the artifact references using the artifact content hash.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ContentHashesItemReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContentHashesItemReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
