package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrganizationsRequestBuilder builds and executes requests for operations under \apis\organizations
type OrganizationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewOrganizationsRequestBuilderInternal instantiates a new OrganizationsRequestBuilder and sets the default values.
func NewOrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsRequestBuilder) {
    m := &OrganizationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/organizations";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrganizationsRequestBuilder instantiates a new OrganizationsRequestBuilder and sets the default values.
func NewOrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// V1 the v1 property
func (m *OrganizationsRequestBuilder) V1()(*OrganizationsV1RequestBuilder) {
    return NewOrganizationsV1RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// V1ById gets an item from the app-services-sdk-go/serviceaccountmgmt/apiv1/client.apis.organizations.v1.item collection
func (m *OrganizationsRequestBuilder) V1ById(id string)(*OrganizationsV1V1ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewOrganizationsV1V1ItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
