package admin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AdminRequestBuilder builds and executes requests for operations under \admin
type AdminRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ArtifactTypes the list of artifact types supported by this instance of Registry.
func (m *AdminRequestBuilder) ArtifactTypes()(*ArtifactTypesRequestBuilder) {
    return NewArtifactTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Config the config property
func (m *AdminRequestBuilder) Config()(*ConfigRequestBuilder) {
    return NewConfigRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewAdminRequestBuilderInternal instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    m := &AdminRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAdminRequestBuilder instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminRequestBuilderInternal(urlParams, requestAdapter)
}
// Export provides a way to export registry data.
func (m *AdminRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ImportEscaped provides a way to import data into the registry.
func (m *AdminRequestBuilder) ImportEscaped()(*ImportRequestBuilder) {
    return NewImportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Loggers manage logger settings/configurations.
func (m *AdminRequestBuilder) Loggers()(*LoggersRequestBuilder) {
    return NewLoggersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LoggersById manage logger settings/configurations.
func (m *AdminRequestBuilder) LoggersById(id string)(*LoggersWithLoggerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["logger"] = id
    }
    return NewLoggersWithLoggerItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// RoleMappings collection to manage role mappings for authenticated principals
func (m *AdminRequestBuilder) RoleMappings()(*RoleMappingsRequestBuilder) {
    return NewRoleMappingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleMappingsById manage the configuration of a single role mapping.
func (m *AdminRequestBuilder) RoleMappingsById(id string)(*RoleMappingsWithPrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["principalId"] = id
    }
    return NewRoleMappingsWithPrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Rules manage the global rules that apply to all artifacts if not otherwise configured.
func (m *AdminRequestBuilder) Rules()(*RulesRequestBuilder) {
    return NewRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RulesById manage the configuration of a single global artifact rule.
func (m *AdminRequestBuilder) RulesById(id string)(*RulesWithRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["rule"] = id
    }
    return NewRulesWithRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
