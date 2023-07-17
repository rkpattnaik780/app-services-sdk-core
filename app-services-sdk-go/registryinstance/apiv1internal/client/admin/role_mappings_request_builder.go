package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// RoleMappingsRequestBuilder collection to manage role mappings for authenticated principals
type RoleMappingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleMappingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleMappingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleMappingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleMappingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrincipalId manage the configuration of a single role mapping.
func (m *RoleMappingsRequestBuilder) ByPrincipalId(principalId string)(*RoleMappingsWithPrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if principalId != "" {
        urlTplParams["principalId"] = principalId
    }
    return NewRoleMappingsWithPrincipalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoleMappingsRequestBuilderInternal instantiates a new RoleMappingsRequestBuilder and sets the default values.
func NewRoleMappingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleMappingsRequestBuilder) {
    m := &RoleMappingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/roleMappings", pathParameters),
    }
    return m
}
// NewRoleMappingsRequestBuilder instantiates a new RoleMappingsRequestBuilder and sets the default values.
func NewRoleMappingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleMappingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleMappingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a list of all role mappings configured in the registry (if any).This operation can fail for the following reasons:* A server error occurred (HTTP error `500`)
func (m *RoleMappingsRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleMappingsRequestBuilderGetRequestConfiguration)([]ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.RoleMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateRoleMappingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.RoleMappingable, len(res))
    for i, v := range res {
        val[i] = v.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.RoleMappingable)
    }
    return val, nil
}
// Post creates a new mapping between a user/principal and a role.This operation can fail for the following reasons:* A server error occurred (HTTP error `500`)
func (m *RoleMappingsRequestBuilder) Post(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.RoleMappingable, requestConfiguration *RoleMappingsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation gets a list of all role mappings configured in the registry (if any).This operation can fail for the following reasons:* A server error occurred (HTTP error `500`)
func (m *RoleMappingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleMappingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation creates a new mapping between a user/principal and a role.This operation can fail for the following reasons:* A server error occurred (HTTP error `500`)
func (m *RoleMappingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.RoleMappingable, requestConfiguration *RoleMappingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
