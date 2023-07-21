package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedHatErrorRepresentation 
type RedHatErrorRepresentation struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error_description property
    error_description *string
    // The error property
    errorEscaped *RedHatErrorRepresentation_error
}
// NewRedHatErrorRepresentation instantiates a new RedHatErrorRepresentation and sets the default values.
func NewRedHatErrorRepresentation()(*RedHatErrorRepresentation) {
    m := &RedHatErrorRepresentation{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRedHatErrorRepresentationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedHatErrorRepresentationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedHatErrorRepresentation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedHatErrorRepresentation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error property
func (m *RedHatErrorRepresentation) GetError()(*RedHatErrorRepresentation_error) {
    return m.errorEscaped
}
// GetErrorDescription gets the error_description property value. The error_description property
func (m *RedHatErrorRepresentation) GetErrorDescription()(*string) {
    return m.error_description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedHatErrorRepresentation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error_description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRedHatErrorRepresentation_error)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*RedHatErrorRepresentation_error))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RedHatErrorRepresentation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetError() != nil {
        cast := (*m.GetError()).String()
        err := writer.WriteStringValue("error", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error_description", m.GetErrorDescription())
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
func (m *RedHatErrorRepresentation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error property
func (m *RedHatErrorRepresentation) SetError(value *RedHatErrorRepresentation_error)() {
    m.errorEscaped = value
}
// SetErrorDescription sets the error_description property value. The error_description property
func (m *RedHatErrorRepresentation) SetErrorDescription(value *string)() {
    m.error_description = value
}
// RedHatErrorRepresentationable 
type RedHatErrorRepresentationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*RedHatErrorRepresentation_error)
    GetErrorDescription()(*string)
    SetError(value *RedHatErrorRepresentation_error)()
    SetErrorDescription(value *string)()
}
