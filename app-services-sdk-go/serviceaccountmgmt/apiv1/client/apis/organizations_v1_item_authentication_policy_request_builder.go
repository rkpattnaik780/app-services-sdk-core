package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// OrganizationsV1ItemAuthenticationPolicyRequestBuilder builds and executes requests for operations under \apis\organizations\v1\{id}\authentication-policy
type OrganizationsV1ItemAuthenticationPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationsV1ItemAuthenticationPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationsV1ItemAuthenticationPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationsV1ItemAuthenticationPolicyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationsV1ItemAuthenticationPolicyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrganizationsV1ItemAuthenticationPolicyRequestBuilderInternal instantiates a new AuthenticationPolicyRequestBuilder and sets the default values.
func NewOrganizationsV1ItemAuthenticationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1ItemAuthenticationPolicyRequestBuilder) {
    m := &OrganizationsV1ItemAuthenticationPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/organizations/v1/{id}/authentication-policy";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrganizationsV1ItemAuthenticationPolicyRequestBuilder instantiates a new AuthenticationPolicyRequestBuilder and sets the default values.
func NewOrganizationsV1ItemAuthenticationPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationsV1ItemAuthenticationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationsV1ItemAuthenticationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get current authentication policy information
func (m *OrganizationsV1ItemAuthenticationPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationsV1ItemAuthenticationPolicyRequestBuilderGetRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateAuthenticationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable), nil
}
// Post update current authentication policy information
func (m *OrganizationsV1ItemAuthenticationPolicyRequestBuilder) Post(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable, requestConfiguration *OrganizationsV1ItemAuthenticationPolicyRequestBuilderPostRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateAuthenticationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable), nil
}
// ToGetRequestInformation get current authentication policy information
func (m *OrganizationsV1ItemAuthenticationPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationsV1ItemAuthenticationPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation update current authentication policy information
func (m *OrganizationsV1ItemAuthenticationPolicyRequestBuilder) ToPostRequestInformation(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AuthenticationPolicyable, requestConfiguration *OrganizationsV1ItemAuthenticationPolicyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
