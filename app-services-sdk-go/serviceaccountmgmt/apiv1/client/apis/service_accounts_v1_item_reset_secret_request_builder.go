package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// Service_accountsV1ItemResetSecretRequestBuilder builds and executes requests for operations under \apis\service_accounts\v1\{id}\resetSecret
type Service_accountsV1ItemResetSecretRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Service_accountsV1ItemResetSecretRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1ItemResetSecretRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewService_accountsV1ItemResetSecretRequestBuilderInternal instantiates a new ResetSecretRequestBuilder and sets the default values.
func NewService_accountsV1ItemResetSecretRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1ItemResetSecretRequestBuilder) {
    m := &Service_accountsV1ItemResetSecretRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/service_accounts/v1/{id}/resetSecret";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewService_accountsV1ItemResetSecretRequestBuilder instantiates a new ResetSecretRequestBuilder and sets the default values.
func NewService_accountsV1ItemResetSecretRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1ItemResetSecretRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewService_accountsV1ItemResetSecretRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset service account secret by id . Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1ItemResetSecretRequestBuilder) Post(ctx context.Context, requestConfiguration *Service_accountsV1ItemResetSecretRequestBuilderPostRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reset service account secret by id . Throws not found exception if the service account is not found or the user does not have access to this service account
func (m *Service_accountsV1ItemResetSecretRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *Service_accountsV1ItemResetSecretRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
