package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// Service_accountsV1RequestBuilder builds and executes requests for operations under \apis\service_accounts\v1
type Service_accountsV1RequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Service_accountsV1RequestBuilderGetQueryParameters returns a list of service accounts created by a user. User information is obtained from the bearer token. The list is paginated with starting index as 'first' and page size as 'max'.
type Service_accountsV1RequestBuilderGetQueryParameters struct {
    // 
    ClientId []string
    // 
    First *int32
    // 
    Max *int32
}
// Service_accountsV1RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Service_accountsV1RequestBuilderGetQueryParameters
}
// Service_accountsV1RequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Service_accountsV1RequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewService_accountsV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewService_accountsV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1RequestBuilder) {
    m := &Service_accountsV1RequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/service_accounts/v1{?first*,max*,clientId*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewService_accountsV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewService_accountsV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Service_accountsV1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewService_accountsV1RequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of service accounts created by a user. User information is obtained from the bearer token. The list is paginated with starting index as 'first' and page size as 'max'.
func (m *Service_accountsV1RequestBuilder) Get(ctx context.Context, requestConfiguration *Service_accountsV1RequestBuilderGetRequestConfiguration)([]idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateValidationExceptionDataFromDiscriminatorValue,
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateServiceAccountDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, len(res))
    for i, v := range res {
        val[i] = v.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable)
    }
    return val, nil
}
// Post create a service account. Created service account is associated with the user defined in the bearer token.
func (m *Service_accountsV1RequestBuilder) Post(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountCreateRequestDataable, requestConfiguration *Service_accountsV1RequestBuilderPostRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateValidationExceptionDataFromDiscriminatorValue,
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "403": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
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
// ToGetRequestInformation returns a list of service accounts created by a user. User information is obtained from the bearer token. The list is paginated with starting index as 'first' and page size as 'max'.
func (m *Service_accountsV1RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Service_accountsV1RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create a service account. Created service account is associated with the user defined in the bearer token.
func (m *Service_accountsV1RequestBuilder) ToPostRequestInformation(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.ServiceAccountCreateRequestDataable, requestConfiguration *Service_accountsV1RequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
