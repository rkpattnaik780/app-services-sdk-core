package client

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i56a034b1b6f11c31d6b71f7d83c50e13d9fc521fff5c35ca18f13e05aadbff6d "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/users"
    i91224c0bf4f8d34f768616bfb0459b647f269f264a30cbd2c8c7283b486c75fc "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/ids"
    ibadcb7ab6aa0b35bc45eb0f8c002c013cbba61e28ac679132f62e590c8e847c3 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/admin"
    ic9fbc1bd9c816b36e6f39796872238fd84a72e2cea282e6b0da91a563f070261 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/groups"
    idc9491da031399ea7f8443d50fb7ccf2118851151b581e0abaf52dba04fd2b64 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/system"
    if2764e35c7c10e53c3ed49c79e3d6d69ef0e4ce6531921de45a36cac7b64fa1c "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/search"
)

// RegistryInstance the main entry point of the SDK, exposes the configuration and the fluent API.
type RegistryInstance struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Admin the admin property
func (m *RegistryInstance) Admin()(*ibadcb7ab6aa0b35bc45eb0f8c002c013cbba61e28ac679132f62e590c8e847c3.AdminRequestBuilder) {
    return ibadcb7ab6aa0b35bc45eb0f8c002c013cbba61e28ac679132f62e590c8e847c3.NewAdminRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRegistryInstance instantiates a new RegistryInstance and sets the default values.
func NewRegistryInstance(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegistryInstance) {
    m := &RegistryInstance{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    return m
}
// Groups collection of the groups in the registry.
func (m *RegistryInstance) Groups()(*ic9fbc1bd9c816b36e6f39796872238fd84a72e2cea282e6b0da91a563f070261.GroupsRequestBuilder) {
    return ic9fbc1bd9c816b36e6f39796872238fd84a72e2cea282e6b0da91a563f070261.NewGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ids the ids property
func (m *RegistryInstance) Ids()(*i91224c0bf4f8d34f768616bfb0459b647f269f264a30cbd2c8c7283b486c75fc.IdsRequestBuilder) {
    return i91224c0bf4f8d34f768616bfb0459b647f269f264a30cbd2c8c7283b486c75fc.NewIdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
func (m *RegistryInstance) Search()(*if2764e35c7c10e53c3ed49c79e3d6d69ef0e4ce6531921de45a36cac7b64fa1c.SearchRequestBuilder) {
    return if2764e35c7c10e53c3ed49c79e3d6d69ef0e4ce6531921de45a36cac7b64fa1c.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// System the system property
func (m *RegistryInstance) System()(*idc9491da031399ea7f8443d50fb7ccf2118851151b581e0abaf52dba04fd2b64.SystemRequestBuilder) {
    return idc9491da031399ea7f8443d50fb7ccf2118851151b581e0abaf52dba04fd2b64.NewSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
func (m *RegistryInstance) Users()(*i56a034b1b6f11c31d6b71f7d83c50e13d9fc521fff5c35ca18f13e05aadbff6d.UsersRequestBuilder) {
    return i56a034b1b6f11c31d6b71f7d83c50e13d9fc521fff5c35ca18f13e05aadbff6d.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
