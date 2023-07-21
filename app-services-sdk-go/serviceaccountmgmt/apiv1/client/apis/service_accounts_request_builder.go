package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Service_accountsRequestBuilder builds and executes requests for operations under \apis\service_accounts
type Service_accountsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewService_accountsRequestBuilderInternal instantiates a new Service_accountsRequestBuilder and sets the default values.
func NewService_accountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsRequestBuilder) {
    m := &Service_accountsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/service_accounts";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewService_accountsRequestBuilder instantiates a new Service_accountsRequestBuilder and sets the default values.
func NewService_accountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewService_accountsRequestBuilderInternal(urlParams, requestAdapter)
}
// V1 the v1 property
func (m *Service_accountsRequestBuilder) V1()(*Service_accountsV1RequestBuilder) {
    return NewService_accountsV1RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// V1ById gets an item from the app-services-sdk-go/serviceaccountmgmt/apiv1/client.apis.service_accounts.v1.item collection
func (m *Service_accountsRequestBuilder) V1ById(id string)(*Service_accountsV1V1ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewService_accountsV1V1ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
