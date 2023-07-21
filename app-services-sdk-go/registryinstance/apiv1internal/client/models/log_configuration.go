package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LogConfiguration 
type LogConfiguration struct {
    NamedLogConfiguration
    // The level property
    level *LogLevel
}
// NewLogConfiguration instantiates a new LogConfiguration and sets the default values.
func NewLogConfiguration()(*LogConfiguration) {
    m := &LogConfiguration{
        NamedLogConfiguration: *NewNamedLogConfiguration(),
    }
    return m
}
// CreateLogConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLogConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NamedLogConfiguration.GetFieldDeserializers()
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLogLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val.(*LogLevel))
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The level property
func (m *LogConfiguration) GetLevel()(*LogLevel) {
    return m.level
}
// Serialize serializes information the current object
func (m *LogConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NamedLogConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLevel() != nil {
        cast := (*m.GetLevel()).String()
        err = writer.WriteStringValue("level", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLevel sets the level property value. The level property
func (m *LogConfiguration) SetLevel(value *LogLevel)() {
    m.level = value
}
// LogConfigurationable 
type LogConfigurationable interface {
    NamedLogConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLevel()(*LogLevel)
    SetLevel(value *LogLevel)()
}
