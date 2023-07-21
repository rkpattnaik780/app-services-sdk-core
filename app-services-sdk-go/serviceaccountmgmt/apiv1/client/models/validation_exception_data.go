package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidationExceptionData 
type ValidationExceptionData struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error_description property
    error_description *string
    // The error property
    errorEscaped *string
    // The fields property
    fields ValidationExceptionData_fieldsable
}
// NewValidationExceptionData instantiates a new ValidationExceptionData and sets the default values.
func NewValidationExceptionData()(*ValidationExceptionData) {
    m := &ValidationExceptionData{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateValidationExceptionDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidationExceptionDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationExceptionData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationExceptionData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetError gets the error property value. The error property
func (m *ValidationExceptionData) GetError()(*string) {
    return m.errorEscaped
}
// GetErrorDescription gets the error_description property value. The error_description property
func (m *ValidationExceptionData) GetErrorDescription()(*string) {
    return m.error_description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidationExceptionData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateValidationExceptionData_fieldsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(ValidationExceptionData_fieldsable))
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The fields property
func (m *ValidationExceptionData) GetFields()(ValidationExceptionData_fieldsable) {
    return m.fields
}
// Serialize serializes information the current object
func (m *ValidationExceptionData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("error", m.GetError())
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
        err := writer.WriteObjectValue("fields", m.GetFields())
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
func (m *ValidationExceptionData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetError sets the error property value. The error property
func (m *ValidationExceptionData) SetError(value *string)() {
    m.errorEscaped = value
}
// SetErrorDescription sets the error_description property value. The error_description property
func (m *ValidationExceptionData) SetErrorDescription(value *string)() {
    m.error_description = value
}
// SetFields sets the fields property value. The fields property
func (m *ValidationExceptionData) SetFields(value ValidationExceptionData_fieldsable)() {
    m.fields = value
}
// ValidationExceptionDataable 
type ValidationExceptionDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*string)
    GetErrorDescription()(*string)
    GetFields()(ValidationExceptionData_fieldsable)
    SetError(value *string)()
    SetErrorDescription(value *string)()
    SetFields(value ValidationExceptionData_fieldsable)()
}
