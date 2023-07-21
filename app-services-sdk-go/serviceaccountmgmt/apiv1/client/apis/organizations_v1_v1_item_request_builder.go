package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrganizationsV1V1ItemRequestBuilder builds and executes requests for operations under \apis\organizations\v1\{id}
type OrganizationsV1V1ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationPolicy the authenticationPolicy property
func (m *OrganizationsV1V1ItemRequestBuilder) AuthenticationPolicy()(*OrganizationsV1ItemAuthenticationPolicyRequestBuilder) {
    return NewOrganizationsV1ItemAuthenticationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewOrganizationsV1V1ItemRequestBuilderInternal instantiates a new V1ItemRequestBuilder and sets the default values.
func NewOrganizationsV1V1ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1V1ItemRequestBuilder) {
    m := &OrganizationsV1V1ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/organizations/v1/{id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrganizationsV1V1ItemRequestBuilder instantiates a new V1ItemRequestBuilder and sets the default values.
func NewOrganizationsV1V1ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1V1ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationsV1V1ItemRequestBuilderInternal(urlParams, requestAdapter)
}
