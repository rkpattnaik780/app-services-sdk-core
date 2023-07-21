package client

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i47f701f5814736c5508f6075be50b4baf8a55044cfa8d6325c212c9e0eb605aa "app-services-sdk-go/registryinstance/apiv1internal/client/ids"
    i598cf0ce26505fedc0e8ccaca4654652780c7c460460e9365521402b612ef92b "app-services-sdk-go/registryinstance/apiv1internal/client/search"
    i7d74922feeb007e3968b92e7d36f12726b27f255b71403723049b614eac94cfe "app-services-sdk-go/registryinstance/apiv1internal/client/users"
    i93a114846ccf438923e4e3572cb98fed71c6e388d867cfb34e2926141f82d972 "app-services-sdk-go/registryinstance/apiv1internal/client/admin"
    id6048425b7bfbfe68519cd90411af6b726e67b4d9334ad040ad59e485f07d596 "app-services-sdk-go/registryinstance/apiv1internal/client/system"
    if2001ac53aaa5190ebfe229b0162e37af1cd33cc434bea2d0854c026b9ba61b1 "app-services-sdk-go/registryinstance/apiv1internal/client/groups"
)

// RegistryInstance the main entry point of the SDK, exposes the configuration and the fluent API.
type RegistryInstance struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Admin the admin property
func (m *RegistryInstance) Admin()(*i93a114846ccf438923e4e3572cb98fed71c6e388d867cfb34e2926141f82d972.AdminRequestBuilder) {
    return i93a114846ccf438923e4e3572cb98fed71c6e388d867cfb34e2926141f82d972.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewRegistryInstance instantiates a new RegistryInstance and sets the default values.
func NewRegistryInstance(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegistryInstance) {
    m := &RegistryInstance{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    return m
}
// Groups collection of the groups in the registry.
func (m *RegistryInstance) Groups()(*if2001ac53aaa5190ebfe229b0162e37af1cd33cc434bea2d0854c026b9ba61b1.GroupsRequestBuilder) {
    return if2001ac53aaa5190ebfe229b0162e37af1cd33cc434bea2d0854c026b9ba61b1.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupsById collection to manage a single group in the registry.
func (m *RegistryInstance) GroupsById(id string)(*if2001ac53aaa5190ebfe229b0162e37af1cd33cc434bea2d0854c026b9ba61b1.WithGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupId"] = id
    }
    return if2001ac53aaa5190ebfe229b0162e37af1cd33cc434bea2d0854c026b9ba61b1.NewWithGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Ids the ids property
func (m *RegistryInstance) Ids()(*i47f701f5814736c5508f6075be50b4baf8a55044cfa8d6325c212c9e0eb605aa.IdsRequestBuilder) {
    return i47f701f5814736c5508f6075be50b4baf8a55044cfa8d6325c212c9e0eb605aa.NewIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Search the search property
func (m *RegistryInstance) Search()(*i598cf0ce26505fedc0e8ccaca4654652780c7c460460e9365521402b612ef92b.SearchRequestBuilder) {
    return i598cf0ce26505fedc0e8ccaca4654652780c7c460460e9365521402b612ef92b.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// System the system property
func (m *RegistryInstance) System()(*id6048425b7bfbfe68519cd90411af6b726e67b4d9334ad040ad59e485f07d596.SystemRequestBuilder) {
    return id6048425b7bfbfe68519cd90411af6b726e67b4d9334ad040ad59e485f07d596.NewSystemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Users the users property
func (m *RegistryInstance) Users()(*i7d74922feeb007e3968b92e7d36f12726b27f255b71403723049b614eac94cfe.UsersRequestBuilder) {
    return i7d74922feeb007e3968b92e7d36f12726b27f255b71403723049b614eac94cfe.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
