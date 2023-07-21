package ids

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// GlobalIdsItemReferencesRequestBuilder builds and executes requests for operations under \ids\globalIds\{globalId}\references
type GlobalIdsItemReferencesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
    }
    m.urlTemplate = "{+baseurl}/ids/globalIds/{globalId}/references{?refType*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGlobalIdsItemReferencesRequestBuilder instantiates a new ReferencesRequestBuilder and sets the default values.
func NewGlobalIdsItemReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsItemReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGlobalIdsItemReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list containing all the artifact references using the artifact global ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *GlobalIdsItemReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration)([]i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactReferenceable, error) {
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
// ToGetRequestInformation returns a list containing all the artifact references using the artifact global ID.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *GlobalIdsItemReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GlobalIdsItemReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
