/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.16.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// FleetshardParameter Fleetshard parameter consumed by enterprise cluster
type FleetshardParameter struct {
	Id *string `json:"id,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewFleetshardParameter instantiates a new FleetshardParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFleetshardParameter() *FleetshardParameter {
	this := FleetshardParameter{}
	return &this
}

// NewFleetshardParameterWithDefaults instantiates a new FleetshardParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFleetshardParameterWithDefaults() *FleetshardParameter {
	this := FleetshardParameter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FleetshardParameter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetshardParameter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FleetshardParameter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FleetshardParameter) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FleetshardParameter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetshardParameter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FleetshardParameter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FleetshardParameter) SetValue(v string) {
	o.Value = &v
}

func (o FleetshardParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFleetshardParameter struct {
	value *FleetshardParameter
	isSet bool
}

func (v NullableFleetshardParameter) Get() *FleetshardParameter {
	return v.value
}

func (v *NullableFleetshardParameter) Set(val *FleetshardParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFleetshardParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFleetshardParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFleetshardParameter(val *FleetshardParameter) *NullableFleetshardParameter {
	return &NullableFleetshardParameter{value: val, isSet: true}
}

func (v NullableFleetshardParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFleetshardParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


