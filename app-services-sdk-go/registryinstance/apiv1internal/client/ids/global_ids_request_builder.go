package ids

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GlobalIdsRequestBuilder builds and executes requests for operations under \ids\globalIds
type GlobalIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewGlobalIdsRequestBuilderInternal instantiates a new GlobalIdsRequestBuilder and sets the default values.
func NewGlobalIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsRequestBuilder) {
    m := &GlobalIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/ids/globalIds";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGlobalIdsRequestBuilder instantiates a new GlobalIdsRequestBuilder and sets the default values.
func NewGlobalIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GlobalIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGlobalIdsRequestBuilderInternal(urlParams, requestAdapter)
}
