package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApisRequestBuilder builds and executes requests for operations under \apis
type ApisRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Beta the beta property
func (m *ApisRequestBuilder) Beta()(*BetaRequestBuilder) {
    return NewBetaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewApisRequestBuilderInternal instantiates a new ApisRequestBuilder and sets the default values.
func NewApisRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApisRequestBuilder) {
    m := &ApisRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewApisRequestBuilder instantiates a new ApisRequestBuilder and sets the default values.
func NewApisRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApisRequestBuilderInternal(urlParams, requestAdapter)
}
// Organizations the organizations property
func (m *ApisRequestBuilder) Organizations()(*OrganizationsRequestBuilder) {
    return NewOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Service_accounts the service_accounts property
func (m *ApisRequestBuilder) Service_accounts()(*Service_accountsRequestBuilder) {
    return NewService_accountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
