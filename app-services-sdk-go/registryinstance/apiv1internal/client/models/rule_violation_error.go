package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RuleViolationError all error responses, whether `4xx` or `5xx` will include one of these as the responsebody.
type RuleViolationError struct {
    Error
}
// NewRuleViolationError instantiates a new RuleViolationError and sets the default values.
func NewRuleViolationError()(*RuleViolationError) {
    m := &RuleViolationError{
        Error: *NewError(),
    }
    return m
}
// CreateRuleViolationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRuleViolationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleViolationError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RuleViolationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Error.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RuleViolationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Error.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RuleViolationErrorable 
type RuleViolationErrorable interface {
    Errorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
