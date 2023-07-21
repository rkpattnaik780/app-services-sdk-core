package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ItemArtifactsItemMetaRequestBuilder manage the metadata of a single artifact.
type ItemArtifactsItemMetaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemArtifactsItemMetaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemMetaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemArtifactsItemMetaRequestBuilderPostQueryParameters gets the metadata for an artifact that matches the raw content.  Searches the registryfor a version of the given artifact matching the content provided in the body of thePOST.This operation can fail for the following reasons:* Provided content (request body) was empty (HTTP error `400`)* No artifact with the `artifactId` exists (HTTP error `404`)* No artifact version matching the provided content exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
type ItemArtifactsItemMetaRequestBuilderPostQueryParameters struct {
    // Parameter that can be set to `true` to indicate that the server should "canonicalize" the content when searching for a matching version.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.
    Canonical *bool
}
// ItemArtifactsItemMetaRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemMetaRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemArtifactsItemMetaRequestBuilderPostQueryParameters
}
// ItemArtifactsItemMetaRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemMetaRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemArtifactsItemMetaRequestBuilderInternal instantiates a new MetaRequestBuilder and sets the default values.
func NewItemArtifactsItemMetaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemMetaRequestBuilder) {
    m := &ItemArtifactsItemMetaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{groupId}/artifacts/{artifactId}/meta{?canonical*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemArtifactsItemMetaRequestBuilder instantiates a new MetaRequestBuilder and sets the default values.
func NewItemArtifactsItemMetaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemMetaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemArtifactsItemMetaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the metadata for an artifact in the registry, based on the latest version. If the latest version of the artifact is marked as `DISABLED`, the next available non-disabled version will be used. The returned metadata includesboth generated (read-only) and editable metadata (such as name and description).This operation can fail for the following reasons:* No artifact with this `artifactId` exists  or all versions are `DISABLED` (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemArtifactsItemMetaRequestBuilderGetRequestConfiguration)(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactMetaDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateArtifactMetaDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactMetaDataable), nil
}
// Post gets the metadata for an artifact that matches the raw content.  Searches the registryfor a version of the given artifact matching the content provided in the body of thePOST.This operation can fail for the following reasons:* Provided content (request body) was empty (HTTP error `400`)* No artifact with the `artifactId` exists (HTTP error `404`)* No artifact version matching the provided content exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) Post(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactContentable, requestConfiguration *ItemArtifactsItemMetaRequestBuilderPostRequestConfiguration)(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.VersionMetaDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateVersionMetaDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.VersionMetaDataable), nil
}
// Put updates the editable parts of the artifact's metadata.  Not all metadata fields canbe updated.  For example, `createdOn` and `createdBy` are both read-only properties.This operation can fail for the following reasons:* No artifact with the `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) Put(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.EditableMetaDataable, requestConfiguration *ItemArtifactsItemMetaRequestBuilderPutRequestConfiguration)(error) {
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
// ToGetRequestInformation gets the metadata for an artifact in the registry, based on the latest version. If the latest version of the artifact is marked as `DISABLED`, the next available non-disabled version will be used. The returned metadata includesboth generated (read-only) and editable metadata (such as name and description).This operation can fail for the following reasons:* No artifact with this `artifactId` exists  or all versions are `DISABLED` (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsItemMetaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation gets the metadata for an artifact that matches the raw content.  Searches the registryfor a version of the given artifact matching the content provided in the body of thePOST.This operation can fail for the following reasons:* Provided content (request body) was empty (HTTP error `400`)* No artifact with the `artifactId` exists (HTTP error `404`)* No artifact version matching the provided content exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) ToPostRequestInformation(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ArtifactContentable, requestConfiguration *ItemArtifactsItemMetaRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/get.extended+json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation updates the editable parts of the artifact's metadata.  Not all metadata fields canbe updated.  For example, `createdOn` and `createdBy` are both read-only properties.This operation can fail for the following reasons:* No artifact with the `artifactId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemMetaRequestBuilder) ToPutRequestInformation(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.EditableMetaDataable, requestConfiguration *ItemArtifactsItemMetaRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
