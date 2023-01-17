/*
Kafka Instance API

API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs

API version: 0.13.0-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ConsumerGroupMetrics struct for ConsumerGroupMetrics
type ConsumerGroupMetrics struct {
	LaggingPartitions *int32 `json:"laggingPartitions,omitempty"`
	ActiveConsumers *int32 `json:"activeConsumers,omitempty"`
	UnassignedPartitions *int32 `json:"unassignedPartitions,omitempty"`
}

// NewConsumerGroupMetrics instantiates a new ConsumerGroupMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupMetrics() *ConsumerGroupMetrics {
	this := ConsumerGroupMetrics{}
	return &this
}

// NewConsumerGroupMetricsWithDefaults instantiates a new ConsumerGroupMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupMetricsWithDefaults() *ConsumerGroupMetrics {
	this := ConsumerGroupMetrics{}
	return &this
}

// GetLaggingPartitions returns the LaggingPartitions field value if set, zero value otherwise.
func (o *ConsumerGroupMetrics) GetLaggingPartitions() int32 {
	if o == nil || isNil(o.LaggingPartitions) {
		var ret int32
		return ret
	}
	return *o.LaggingPartitions
}

// GetLaggingPartitionsOk returns a tuple with the LaggingPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupMetrics) GetLaggingPartitionsOk() (*int32, bool) {
	if o == nil || isNil(o.LaggingPartitions) {
    return nil, false
	}
	return o.LaggingPartitions, true
}

// HasLaggingPartitions returns a boolean if a field has been set.
func (o *ConsumerGroupMetrics) HasLaggingPartitions() bool {
	if o != nil && !isNil(o.LaggingPartitions) {
		return true
	}

	return false
}

// SetLaggingPartitions gets a reference to the given int32 and assigns it to the LaggingPartitions field.
func (o *ConsumerGroupMetrics) SetLaggingPartitions(v int32) {
	o.LaggingPartitions = &v
}

// GetActiveConsumers returns the ActiveConsumers field value if set, zero value otherwise.
func (o *ConsumerGroupMetrics) GetActiveConsumers() int32 {
	if o == nil || isNil(o.ActiveConsumers) {
		var ret int32
		return ret
	}
	return *o.ActiveConsumers
}

// GetActiveConsumersOk returns a tuple with the ActiveConsumers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupMetrics) GetActiveConsumersOk() (*int32, bool) {
	if o == nil || isNil(o.ActiveConsumers) {
    return nil, false
	}
	return o.ActiveConsumers, true
}

// HasActiveConsumers returns a boolean if a field has been set.
func (o *ConsumerGroupMetrics) HasActiveConsumers() bool {
	if o != nil && !isNil(o.ActiveConsumers) {
		return true
	}

	return false
}

// SetActiveConsumers gets a reference to the given int32 and assigns it to the ActiveConsumers field.
func (o *ConsumerGroupMetrics) SetActiveConsumers(v int32) {
	o.ActiveConsumers = &v
}

// GetUnassignedPartitions returns the UnassignedPartitions field value if set, zero value otherwise.
func (o *ConsumerGroupMetrics) GetUnassignedPartitions() int32 {
	if o == nil || isNil(o.UnassignedPartitions) {
		var ret int32
		return ret
	}
	return *o.UnassignedPartitions
}

// GetUnassignedPartitionsOk returns a tuple with the UnassignedPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupMetrics) GetUnassignedPartitionsOk() (*int32, bool) {
	if o == nil || isNil(o.UnassignedPartitions) {
    return nil, false
	}
	return o.UnassignedPartitions, true
}

// HasUnassignedPartitions returns a boolean if a field has been set.
func (o *ConsumerGroupMetrics) HasUnassignedPartitions() bool {
	if o != nil && !isNil(o.UnassignedPartitions) {
		return true
	}

	return false
}

// SetUnassignedPartitions gets a reference to the given int32 and assigns it to the UnassignedPartitions field.
func (o *ConsumerGroupMetrics) SetUnassignedPartitions(v int32) {
	o.UnassignedPartitions = &v
}

func (o ConsumerGroupMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LaggingPartitions) {
		toSerialize["laggingPartitions"] = o.LaggingPartitions
	}
	if !isNil(o.ActiveConsumers) {
		toSerialize["activeConsumers"] = o.ActiveConsumers
	}
	if !isNil(o.UnassignedPartitions) {
		toSerialize["unassignedPartitions"] = o.UnassignedPartitions
	}
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupMetrics struct {
	value *ConsumerGroupMetrics
	isSet bool
}

func (v NullableConsumerGroupMetrics) Get() *ConsumerGroupMetrics {
	return v.value
}

func (v *NullableConsumerGroupMetrics) Set(val *ConsumerGroupMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupMetrics(val *ConsumerGroupMetrics) *NullableConsumerGroupMetrics {
	return &NullableConsumerGroupMetrics{value: val, isSet: true}
}

func (v NullableConsumerGroupMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

