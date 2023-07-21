package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationPolicy 
type AuthenticationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authenticationFactors property
    authenticationFactors AuthenticationFactorsable
}
// NewAuthenticationPolicy instantiates a new AuthenticationPolicy and sets the default values.
func NewAuthenticationPolicy()(*AuthenticationPolicy) {
    m := &AuthenticationPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthenticationFactors gets the authenticationFactors property value. The authenticationFactors property
func (m *AuthenticationPolicy) GetAuthenticationFactors()(AuthenticationFactorsable) {
    return m.authenticationFactors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationFactors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationFactorsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationFactors(val.(AuthenticationFactorsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationFactors", m.GetAuthenticationFactors())
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
func (m *AuthenticationPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthenticationFactors sets the authenticationFactors property value. The authenticationFactors property
func (m *AuthenticationPolicy) SetAuthenticationFactors(value AuthenticationFactorsable)() {
    m.authenticationFactors = value
}
// AuthenticationPolicyable 
type AuthenticationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationFactors()(AuthenticationFactorsable)
    SetAuthenticationFactors(value AuthenticationFactorsable)()
}
