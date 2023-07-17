package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ArtifactsRequestBuilder search for artifacts in the registry.
type ArtifactsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ArtifactsRequestBuilderGetQueryParameters returns a paginated list of all artifacts that match the provided filter criteria.
type ArtifactsRequestBuilderGetQueryParameters struct {
    // Filter by contentId.
    ContentId *int64
    // Filter by description.
    Description *string
    // Filter by globalId.
    GlobalId *int64
    // Filter by artifact group.
    Group *string
    // Filter by label.  Include one or more label to only return artifacts containing all of thespecified labels.
    Labels []string
    // The number of artifacts to return.  Defaults to 20.
    Limit *int32
    // Filter by artifact name.
    Name *string
    // The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
    Offset *int32
    // Sort order, ascending (`asc`) or descending (`desc`).
    Order *string
    // The field to sort by.  Can be one of:* `name`* `createdOn`
    Orderby *string
    // Filter by one or more name/value property.  Separate each name/value pair using a colon.  Forexample `properties=foo:bar` will return only artifacts with a custom property named `foo`and value `bar`.
    Properties []string
}
// ArtifactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ArtifactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ArtifactsRequestBuilderGetQueryParameters
}
// ArtifactsRequestBuilderPostQueryParameters returns a paginated list of all artifacts with at least one version that matches theposted content.
type ArtifactsRequestBuilderPostQueryParameters struct {
    // Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the `canonical` query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts.
    ArtifactType *string
    // Parameter that can be set to `true` to indicate that the server should "canonicalize" the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the `artifactType` query parameter.
    Canonical *bool
    // The number of artifacts to return.  Defaults to 20.
    Limit *int32
    // The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
    Offset *int32
    // Sort order, ascending (`asc`) or descending (`desc`).
    Order *string
    // The field to sort by.  Can be one of:* `name`* `createdOn`
    Orderby *string
}
// ArtifactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ArtifactsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ArtifactsRequestBuilderPostQueryParameters
}
// NewArtifactsRequestBuilderInternal instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewArtifactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ArtifactsRequestBuilder) {
    m := &ArtifactsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/artifacts{?name*,offset*,limit*,order*,orderby*,labels*,properties*,description*,group*,globalId*,contentId*,canonical*,artifactType*}", pathParameters),
    }
    return m
}
// NewArtifactsRequestBuilder instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewArtifactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ArtifactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewArtifactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a paginated list of all artifacts that match the provided filter criteria.
func (m *ArtifactsRequestBuilder) Get(ctx context.Context, requestConfiguration *ArtifactsRequestBuilderGetRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactSearchResultsable, error) {
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
// Post returns a paginated list of all artifacts with at least one version that matches theposted content.
func (m *ArtifactsRequestBuilder) Post(ctx context.Context, body []byte, requestConfiguration *ArtifactsRequestBuilderPostRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.ArtifactSearchResultsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation returns a paginated list of all artifacts that match the provided filter criteria.
func (m *ArtifactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ArtifactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation returns a paginated list of all artifacts with at least one version that matches theposted content.
func (m *ArtifactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body []byte, requestConfiguration *ArtifactsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
