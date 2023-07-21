package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BetaRequestBuilder builds and executes requests for operations under \apis\beta
type BetaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Acs the acs property
func (m *BetaRequestBuilder) Acs()(*BetaAcsRequestBuilder) {
    return NewBetaAcsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewBetaRequestBuilderInternal instantiates a new BetaRequestBuilder and sets the default values.
func NewBetaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaRequestBuilder) {
    m := &BetaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/beta";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBetaRequestBuilder instantiates a new BetaRequestBuilder and sets the default values.
func NewBetaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBetaRequestBuilderInternal(urlParams, requestAdapter)
}
