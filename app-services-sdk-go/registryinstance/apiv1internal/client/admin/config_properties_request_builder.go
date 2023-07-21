package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ConfigPropertiesRequestBuilder manage configuration properties.
type ConfigPropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConfigPropertiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigPropertiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigPropertiesRequestBuilderInternal instantiates a new PropertiesRequestBuilder and sets the default values.
func NewConfigPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigPropertiesRequestBuilder) {
    m := &ConfigPropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/config/properties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewConfigPropertiesRequestBuilder instantiates a new PropertiesRequestBuilder and sets the default values.
func NewConfigPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of all configuration properties that have been set.  The list is not paged.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ConfigPropertiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigPropertiesRequestBuilderGetRequestConfiguration)([]i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ConfigurationPropertyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateConfigurationPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ConfigurationPropertyable, len(res))
    for i, v := range res {
        val[i] = v.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.ConfigurationPropertyable)
    }
    return val, nil
}
// ToGetRequestInformation returns a list of all configuration properties that have been set.  The list is not paged.This operation may fail for one of the following reasons:* A server error occurred (HTTP error `500`)
func (m *ConfigPropertiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigPropertiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
