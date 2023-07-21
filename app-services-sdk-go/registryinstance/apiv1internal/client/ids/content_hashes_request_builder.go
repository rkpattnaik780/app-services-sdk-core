package ids

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContentHashesRequestBuilder builds and executes requests for operations under \ids\contentHashes
type ContentHashesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewContentHashesRequestBuilderInternal instantiates a new ContentHashesRequestBuilder and sets the default values.
func NewContentHashesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesRequestBuilder) {
    m := &ContentHashesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/ids/contentHashes";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewContentHashesRequestBuilder instantiates a new ContentHashesRequestBuilder and sets the default values.
func NewContentHashesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentHashesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentHashesRequestBuilderInternal(urlParams, requestAdapter)
}
