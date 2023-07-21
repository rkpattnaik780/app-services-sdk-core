package system

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SystemRequestBuilder builds and executes requests for operations under \system
type SystemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewSystemRequestBuilderInternal instantiates a new SystemRequestBuilder and sets the default values.
func NewSystemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemRequestBuilder) {
    m := &SystemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/system";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSystemRequestBuilder instantiates a new SystemRequestBuilder and sets the default values.
func NewSystemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SystemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSystemRequestBuilderInternal(urlParams, requestAdapter)
}
// Info retrieve system information
func (m *SystemRequestBuilder) Info()(*InfoRequestBuilder) {
    return NewInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Limits retrieve resource limits information
func (m *SystemRequestBuilder) Limits()(*LimitsRequestBuilder) {
    return NewLimitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
