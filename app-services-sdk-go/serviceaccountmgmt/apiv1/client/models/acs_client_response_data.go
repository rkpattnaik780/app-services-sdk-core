package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AcsClientResponseData 
type AcsClientResponseData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientId property
    clientId *string
    // The createdAt property
    createdAt *int64
    // The name property
    name *string
    // The secret property
    secret *string
}
// NewAcsClientResponseData instantiates a new AcsClientResponseData and sets the default values.
func NewAcsClientResponseData()(*AcsClientResponseData) {
    m := &AcsClientResponseData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAcsClientResponseDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcsClientResponseDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcsClientResponseData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AcsClientResponseData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientId gets the clientId property value. The clientId property
func (m *AcsClientResponseData) GetClientId()(*string) {
    return m.clientId
}
// GetCreatedAt gets the createdAt property value. The createdAt property
func (m *AcsClientResponseData) GetCreatedAt()(*int64) {
    return m.createdAt
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AcsClientResponseData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetName gets the name property value. The name property
func (m *AcsClientResponseData) GetName()(*string) {
    return m.name
}
// GetSecret gets the secret property value. The secret property
func (m *AcsClientResponseData) GetSecret()(*string) {
    return m.secret
}
// Serialize serializes information the current object
func (m *AcsClientResponseData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AcsClientResponseData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientId sets the clientId property value. The clientId property
func (m *AcsClientResponseData) SetClientId(value *string)() {
    m.clientId = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *AcsClientResponseData) SetCreatedAt(value *int64)() {
    m.createdAt = value
}
// SetName sets the name property value. The name property
func (m *AcsClientResponseData) SetName(value *string)() {
    m.name = value
}
// SetSecret sets the secret property value. The secret property
func (m *AcsClientResponseData) SetSecret(value *string)() {
    m.secret = value
}
// AcsClientResponseDataable 
type AcsClientResponseDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetCreatedAt()(*int64)
    GetName()(*string)
    GetSecret()(*string)
    SetClientId(value *string)()
    SetCreatedAt(value *int64)()
    SetName(value *string)()
    SetSecret(value *string)()
}
