package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BetaAcsRequestBuilder builds and executes requests for operations under \apis\beta\acs
type BetaAcsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewBetaAcsRequestBuilderInternal instantiates a new AcsRequestBuilder and sets the default values.
func NewBetaAcsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsRequestBuilder) {
    m := &BetaAcsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/beta/acs";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBetaAcsRequestBuilder instantiates a new AcsRequestBuilder and sets the default values.
func NewBetaAcsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBetaAcsRequestBuilderInternal(urlParams, requestAdapter)
}
// V1 the v1 property
func (m *BetaAcsRequestBuilder) V1()(*BetaAcsV1RequestBuilder) {
    return NewBetaAcsV1RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// V1ById gets an item from the app-services-sdk-go/serviceaccountmgmt/apiv1/client.apis.beta.acs.v1.item collection
func (m *BetaAcsRequestBuilder) V1ById(id string)(*BetaAcsV1WithClientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["clientId"] = id
    }
    return NewBetaAcsV1WithClientItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
