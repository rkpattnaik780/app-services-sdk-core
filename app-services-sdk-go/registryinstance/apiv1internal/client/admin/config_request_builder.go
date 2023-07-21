package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ConfigRequestBuilder builds and executes requests for operations under \admin\config
type ConfigRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewConfigRequestBuilderInternal instantiates a new ConfigRequestBuilder and sets the default values.
func NewConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigRequestBuilder) {
    m := &ConfigRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/config";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewConfigRequestBuilder instantiates a new ConfigRequestBuilder and sets the default values.
func NewConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Properties manage configuration properties.
func (m *ConfigRequestBuilder) Properties()(*ConfigPropertiesRequestBuilder) {
    return NewConfigPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PropertiesById manage a single configuration property (by name).
func (m *ConfigRequestBuilder) PropertiesById(id string)(*ConfigPropertiesWithPropertyNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["propertyName"] = id
    }
    return NewConfigPropertiesWithPropertyNameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
