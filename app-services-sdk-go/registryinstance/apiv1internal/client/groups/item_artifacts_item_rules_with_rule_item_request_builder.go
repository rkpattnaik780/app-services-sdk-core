package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/models"
)

// ItemArtifactsItemRulesWithRuleItemRequestBuilder manage the configuration of a single artifact rule.
type ItemArtifactsItemRulesWithRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemArtifactsItemRulesWithRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemRulesWithRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemArtifactsItemRulesWithRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemRulesWithRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemArtifactsItemRulesWithRuleItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemArtifactsItemRulesWithRuleItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemArtifactsItemRulesWithRuleItemRequestBuilderInternal instantiates a new WithRuleItemRequestBuilder and sets the default values.
func NewItemArtifactsItemRulesWithRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemRulesWithRuleItemRequestBuilder) {
    m := &ItemArtifactsItemRulesWithRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{groupId}/artifacts/{artifactId}/rules/{rule}", pathParameters),
    }
    return m
}
// NewItemArtifactsItemRulesWithRuleItemRequestBuilder instantiates a new WithRuleItemRequestBuilder and sets the default values.
func NewItemArtifactsItemRulesWithRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemArtifactsItemRulesWithRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemArtifactsItemRulesWithRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a rule from the artifact.  This results in the rule no longer applying forthis artifact.  If this is the only rule configured for the artifact, this is the same as deleting **all** rules, and the globally configured rules now apply tothis artifact.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns information about a single rule configured for an artifact.  This is usefulwhen you want to know what the current configuration settings are for a specific rule.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderGetRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable), nil
}
// Put updates the configuration of a single rule for the artifact.  The configuration datais specific to each rule type, so the configuration of the `COMPATIBILITY` rule is in a different format from the configuration of the `VALIDITY` rule.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) Put(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderPutRequestConfiguration)(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
        "500": ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.CreateRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable), nil
}
// ToDeleteRequestInformation deletes a rule from the artifact.  This results in the rule no longer applying forthis artifact.  If this is the only rule configured for the artifact, this is the same as deleting **all** rules, and the globally configured rules now apply tothis artifact.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns information about a single rule configured for an artifact.  This is usefulwhen you want to know what the current configuration settings are for a specific rule.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation updates the configuration of a single rule for the artifact.  The configuration datais specific to each rule type, so the configuration of the `COMPATIBILITY` rule is in a different format from the configuration of the `VALIDITY` rule.This operation can fail for the following reasons:* No artifact with this `artifactId` exists (HTTP error `404`)* No rule with this name/type is configured for this artifact (HTTP error `404`)* Invalid rule type (HTTP error `400`)* A server error occurred (HTTP error `500`)
func (m *ItemArtifactsItemRulesWithRuleItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie6243ad9ab62ed9cb687ec1fd7d609a23aa5c7a7654973c2beba0d848a532945.Ruleable, requestConfiguration *ItemArtifactsItemRulesWithRuleItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.Add("Accept", "application/json")
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
