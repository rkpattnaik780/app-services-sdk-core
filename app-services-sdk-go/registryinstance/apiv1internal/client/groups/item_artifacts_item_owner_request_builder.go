package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ItemArtifactsItemOwnerRequestBuilder manage the ownership of a single artifact.
type ItemArtifactsItemOwnerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemArtifactsItemOwnerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemOwnerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemArtifactsItemOwnerRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemOwnerRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemArtifactsItemOwnerRequestBuilderInternal instantiates a new OwnerRequestBuilder and sets the default values.
func NewItemArtifactsItemOwnerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemOwnerRequestBuilder) {
    m := &ItemArtifactsItemOwnerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{groupId}/artifacts/{artifactId}/owner";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemArtifactsItemOwnerRequestBuilder instantiates a new OwnerRequestBuilder and sets the default values.
func NewItemArtifactsItemOwnerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemOwnerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemArtifactsItemOwnerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the owner of an artifact in the registry.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemOwnerRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemArtifactsItemOwnerRequestBuilderGetRequestConfiguration)(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactOwnerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateArtifactOwnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactOwnerable), nil
}
// Put changes the ownership of an artifact.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemOwnerRequestBuilder) Put(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactOwnerable, requestConfiguration *ItemArtifactsItemOwnerRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets the owner of an artifact in the registry.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemOwnerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsItemOwnerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation changes the ownership of an artifact.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemOwnerRequestBuilder) ToPutRequestInformation(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactOwnerable, requestConfiguration *ItemArtifactsItemOwnerRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
