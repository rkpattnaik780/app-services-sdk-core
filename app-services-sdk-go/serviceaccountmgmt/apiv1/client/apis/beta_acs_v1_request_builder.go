package apis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811 "app-services-sdk-go/serviceaccountmgmt/apiv1/client/models"
)

// BetaAcsV1RequestBuilder builds and executes requests for operations under \apis\beta\acs\v1
type BetaAcsV1RequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BetaAcsV1RequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BetaAcsV1RequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBetaAcsV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewBetaAcsV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsV1RequestBuilder) {
    m := &BetaAcsV1RequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/apis/beta/acs/v1";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewBetaAcsV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewBetaAcsV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BetaAcsV1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBetaAcsV1RequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an ACS managed central client. Created ACS managed central clients are associated with the supplied organization id.
func (m *BetaAcsV1RequestBuilder) Post(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AcsClientRequestDataable, requestConfiguration *BetaAcsV1RequestBuilderPostRequestConfiguration)(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AcsClientResponseDataable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateValidationExceptionDataFromDiscriminatorValue,
        "401": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateErrorFromDiscriminatorValue,
        "403": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
        "405": idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateRedHatErrorRepresentationFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.CreateAcsClientResponseDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AcsClientResponseDataable), nil
}
// ToPostRequestInformation create an ACS managed central client. Created ACS managed central clients are associated with the supplied organization id.
func (m *BetaAcsV1RequestBuilder) ToPostRequestInformation(ctx context.Context, body idb896b5d7454acc9be6945db672b5fc828d0a7148dfb758c75ebc5a85d81e811.AcsClientRequestDataable, requestConfiguration *BetaAcsV1RequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
