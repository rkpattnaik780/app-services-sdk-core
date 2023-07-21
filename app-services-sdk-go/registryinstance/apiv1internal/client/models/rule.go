package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Rule 
type Rule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The config property
    config *string
    // The type property
    typeEscaped *RuleType
}
// NewRule instantiates a new Rule and sets the default values.
func NewRule()(*Rule) {
    m := &Rule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Rule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. The config property
func (m *Rule) GetConfig()(*string) {
    return m.config
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Rule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*RuleType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type property
func (m *Rule) GetType()(*RuleType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Rule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Rule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. The config property
func (m *Rule) SetConfig(value *string)() {
    m.config = value
}
// SetType sets the type property value. The type property
func (m *Rule) SetType(value *RuleType)() {
    m.typeEscaped = value
}
// Ruleable 
type Ruleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfig()(*string)
    GetType()(*RuleType)
    SetConfig(value *string)()
    SetType(value *RuleType)()
}
