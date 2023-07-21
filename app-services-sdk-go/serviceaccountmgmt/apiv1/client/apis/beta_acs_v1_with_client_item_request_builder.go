package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// BetaAcsV1WithClientItemRequestBuilder builds and executes requests for operations under \apis\beta\acs\v1\{clientId}
type BetaAcsV1WithClientItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BetaAcsV1WithClientItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BetaAcsV1WithClientItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBetaAcsV1WithClientItemRequestBuilderInternal instantiates a new WithClientItemRequestBuilder and sets the default values.
func NewBetaAcsV1WithClientItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsV1WithClientItemRequestBuilder) {
    m := &BetaAcsV1WithClientItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/beta/acs/v1/{clientId}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBetaAcsV1WithClientItemRequestBuilder instantiates a new WithClientItemRequestBuilder and sets the default values.
func NewBetaAcsV1WithClientItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsV1WithClientItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBetaAcsV1WithClientItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ACS managed central client by clientId. Throws not found exception if the client is not found
func (m *BetaAcsV1WithClientItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BetaAcsV1WithClientItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateValidationExceptionDataFromDiscriminatorValue,
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "404": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
        "405": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete ACS managed central client by clientId. Throws not found exception if the client is not found
func (m *BetaAcsV1WithClientItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BetaAcsV1WithClientItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
