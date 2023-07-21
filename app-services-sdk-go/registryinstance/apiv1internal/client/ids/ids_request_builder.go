package ids

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdsRequestBuilder builds and executes requests for operations under \ids
type IdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdsRequestBuilderInternal instantiates a new IdsRequestBuilder and sets the default values.
func NewIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdsRequestBuilder) {
    m := &IdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/ids";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIdsRequestBuilder instantiates a new IdsRequestBuilder and sets the default values.
func NewIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentHashes the contentHashes property
func (m *IdsRequestBuilder) ContentHashes()(*ContentHashesRequestBuilder) {
    return NewContentHashesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentHashesById access artifact content utilizing the SHA-256 hash of the content.
func (m *IdsRequestBuilder) ContentHashesById(id string)(*ContentHashesWithContentHashItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentHash"] = id
    }
    return NewContentHashesWithContentHashItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ContentIds the contentIds property
func (m *IdsRequestBuilder) ContentIds()(*ContentIdsRequestBuilder) {
    return NewContentIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentIdsById access artifact content utilizing the unique content identifier for that content.
func (m *IdsRequestBuilder) ContentIdsById(id string)(*ContentIdsWithContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentId"] = id
    }
    return NewContentIdsWithContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// GlobalIds the globalIds property
func (m *IdsRequestBuilder) GlobalIds()(*GlobalIdsRequestBuilder) {
    return NewGlobalIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GlobalIdsById access artifact content utilizing an artifact version's globally unique identifier.
func (m *IdsRequestBuilder) GlobalIdsById(id string)(*GlobalIdsWithGlobalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["globalId"] = id
    }
    return NewGlobalIdsWithGlobalItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
