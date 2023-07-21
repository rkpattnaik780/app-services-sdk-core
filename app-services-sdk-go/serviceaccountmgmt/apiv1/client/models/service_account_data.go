package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAccountData 
type ServiceAccountData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientId property
    clientId *string
    // The createdAt property
    createdAt *int64
    // The createdBy property
    createdBy *string
    // The description property
    description *string
    // The id property
    id *string
    // The name property
    name *string
    // Provided during creation and resetting of service account credentials.
    secret *string
}
// NewServiceAccountData instantiates a new ServiceAccountData and sets the default values.
func NewServiceAccountData()(*ServiceAccountData) {
    m := &ServiceAccountData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAccountDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAccountDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAccountData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceAccountData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the clientId property value. The clientId property
func (m *ServiceAccountData) GetClientId()(*string) {
    return m.clientId
}
// GetCreatedAt gets the createdAt property value. The createdAt property
func (m *ServiceAccountData) GetCreatedAt()(*int64) {
    return m.createdAt
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *ServiceAccountData) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetDescription gets the description property value. The description property
func (m *ServiceAccountData) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAccountData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ServiceAccountData) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name property
func (m *ServiceAccountData) GetName()(*string) {
    return m.name
}
// GetSecret gets the secret property value. Provided during creation and resetting of service account credentials.
func (m *ServiceAccountData) GetSecret()(*string) {
    return m.secret
}
// Serialize serializes information the current object
func (m *ServiceAccountData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceAccountData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the clientId property value. The clientId property
func (m *ServiceAccountData) SetClientId(value *string)() {
    m.clientId = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *ServiceAccountData) SetCreatedAt(value *int64)() {
    m.createdAt = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *ServiceAccountData) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetDescription sets the description property value. The description property
func (m *ServiceAccountData) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The id property
func (m *ServiceAccountData) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *ServiceAccountData) SetName(value *string)() {
    m.name = value
}
// SetSecret sets the secret property value. Provided during creation and resetting of service account credentials.
func (m *ServiceAccountData) SetSecret(value *string)() {
    m.secret = value
}
// ServiceAccountDataable 
type ServiceAccountDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetCreatedAt()(*int64)
    GetCreatedBy()(*string)
    GetDescription()(*string)
    GetId()(*string)
    GetName()(*string)
    GetSecret()(*string)
    SetClientId(value *string)()
    SetCreatedAt(value *int64)()
    SetCreatedBy(value *string)()
    SetDescription(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetSecret(value *string)()
}
