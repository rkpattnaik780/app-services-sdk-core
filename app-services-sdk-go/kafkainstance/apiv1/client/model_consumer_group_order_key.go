/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.0-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// ConsumerGroupOrderKey the model 'ConsumerGroupOrderKey'
type ConsumerGroupOrderKey string

// List of ConsumerGroupOrderKey
const (
	CONSUMERGROUPORDERKEY_NAME ConsumerGroupOrderKey = "name"
)

// All allowed values of ConsumerGroupOrderKey enum
var AllowedConsumerGroupOrderKeyEnumValues = []ConsumerGroupOrderKey{
	"name",
}

func (v *ConsumerGroupOrderKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsumerGroupOrderKey(value)
	for _, existing := range AllowedConsumerGroupOrderKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsumerGroupOrderKey", value)
}

// NewConsumerGroupOrderKeyFromValue returns a pointer to a valid ConsumerGroupOrderKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsumerGroupOrderKeyFromValue(v string) (*ConsumerGroupOrderKey, error) {
	ev := ConsumerGroupOrderKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsumerGroupOrderKey: valid values are %v", v, AllowedConsumerGroupOrderKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsumerGroupOrderKey) IsValid() bool {
	for _, existing := range AllowedConsumerGroupOrderKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsumerGroupOrderKey value
func (v ConsumerGroupOrderKey) Ptr() *ConsumerGroupOrderKey {
	return &v
}

type NullableConsumerGroupOrderKey struct {
	value *ConsumerGroupOrderKey
	isSet bool
}

func (v NullableConsumerGroupOrderKey) Get() *ConsumerGroupOrderKey {
	return v.value
}

func (v *NullableConsumerGroupOrderKey) Set(val *ConsumerGroupOrderKey) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupOrderKey) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupOrderKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupOrderKey(val *ConsumerGroupOrderKey) *NullableConsumerGroupOrderKey {
	return &NullableConsumerGroupOrderKey{value: val, isSet: true}
}

func (v NullableConsumerGroupOrderKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupOrderKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
