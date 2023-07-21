package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586 "app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// RoleMappingsWithPrincipalItemRequestBuilder manage the configuration of a single role mapping.
type RoleMappingsWithPrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleMappingsWithPrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleMappingsWithPrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleMappingsWithPrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleMappingsWithPrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleMappingsWithPrincipalItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleMappingsWithPrincipalItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleMappingsWithPrincipalItemRequestBuilderInternal instantiates a new WithPrincipalItemRequestBuilder and sets the default values.
func NewRoleMappingsWithPrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleMappingsWithPrincipalItemRequestBuilder) {
    m := &RoleMappingsWithPrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/roleMappings/{principalId}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewRoleMappingsWithPrincipalItemRequestBuilder instantiates a new WithPrincipalItemRequestBuilder and sets the default values.
func NewRoleMappingsWithPrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleMappingsWithPrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleMappingsWithPrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a single role mapping, effectively denying access to a user/principal.This operation can fail for the following reasons:* No role mapping for the principalId exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get gets the details of a single role mapping (by `principalId`).This operation can fail for the following reasons:* No role mapping for the `principalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderGetRequestConfiguration)(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.RoleMappingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateRoleMappingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.RoleMappingable), nil
}
// Put updates a single role mapping for one user/principal.This operation can fail for the following reasons:* No role mapping for the principalId exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) Put(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.UpdateRoleable, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
        "500": i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation deletes a single role mapping, effectively denying access to a user/principal.This operation can fail for the following reasons:* No role mapping for the principalId exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets the details of a single role mapping (by `principalId`).This operation can fail for the following reasons:* No role mapping for the `principalId` exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates a single role mapping for one user/principal.This operation can fail for the following reasons:* No role mapping for the principalId exists (HTTP error `404`)* A server error occurred (HTTP error `500`)
func (m *RoleMappingsWithPrincipalItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body i061057543182c83985ba910ffdfb7dc0d3545e4ecde182f67803d779bc358586.UpdateRoleable, requestConfiguration *RoleMappingsWithPrincipalItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
