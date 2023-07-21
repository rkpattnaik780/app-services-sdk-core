package apis

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrganizationsV1RequestBuilder builds and executes requests for operations under \apis\organizations\v1
type OrganizationsV1RequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewOrganizationsV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewOrganizationsV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1RequestBuilder) {
    m := &OrganizationsV1RequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/organizations/v1";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrganizationsV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewOrganizationsV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationsV1RequestBuilderInternal(urlParams, requestAdapter)
}
