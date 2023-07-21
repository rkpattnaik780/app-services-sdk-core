package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// Service_accountsV1V1ItemRequestBuilder builds and executes requests for operations under \apis\service_accounts\v1\{id}
type Service_accountsV1V1ItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Service_accountsV1V1ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1V1ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Service_accountsV1V1ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1V1ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Service_accountsV1V1ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1V1ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewService_accountsV1V1ItemRequestBuilderInternal instantiates a new V1ItemRequestBuilder and sets the default values.
func NewService_accountsV1V1ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1V1ItemRequestBuilder) {
    m := &Service_accountsV1V1ItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/service_accounts/v1/{id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewService_accountsV1V1ItemRequestBuilder instantiates a new V1ItemRequestBuilder and sets the default values.
func NewService_accountsV1V1ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1V1ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewService_accountsV1V1ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete service account by id. Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1V1ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *Service_accountsV1V1ItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "404": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get returns service account by id. Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1V1ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *Service_accountsV1V1ItemRequestBuilderGetRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "404": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateServiceAccountDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable), nil
}
// Patch update a service account by id.
func (m *Service_accountsV1V1ItemRequestBuilder) Patch(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountRequestDataable, requestConfiguration *Service_accountsV1V1ItemRequestBuilderPatchRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateValidationExceptionDataFromDiscriminatorValue,
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "404": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateServiceAccountDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable), nil
}
// ResetSecret the resetSecret property
func (m *Service_accountsV1V1ItemRequestBuilder) ResetSecret()(*Service_accountsV1ItemResetSecretRequestBuilder) {
    return NewService_accountsV1ItemResetSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete service account by id. Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1V1ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *Service_accountsV1V1ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns service account by id. Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1V1ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Service_accountsV1V1ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a service account by id.
func (m *Service_accountsV1V1ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountRequestDataable, requestConfiguration *Service_accountsV1V1ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
