/*
Account Management Service API

Manage user subscriptions and clusters

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// CapabilityReview struct for CapabilityReview
type CapabilityReview struct {
	Result string `json:"result"`
}

// NewCapabilityReview instantiates a new CapabilityReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityReview(result string) *CapabilityReview {
	this := CapabilityReview{}
	this.Result = result
	return &this
}

// NewCapabilityReviewWithDefaults instantiates a new CapabilityReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityReviewWithDefaults() *CapabilityReview {
	this := CapabilityReview{}
	return &this
}

// GetResult returns the Result field value
func (o *CapabilityReview) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CapabilityReview) GetResultOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CapabilityReview) SetResult(v string) {
	o.Result = v
}

func (o CapabilityReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityReview struct {
	value *CapabilityReview
	isSet bool
}

func (v NullableCapabilityReview) Get() *CapabilityReview {
	return v.value
}

func (v *NullableCapabilityReview) Set(val *CapabilityReview) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityReview) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityReview(val *CapabilityReview) *NullableCapabilityReview {
	return &NullableCapabilityReview{value: val, isSet: true}
}

func (v NullableCapabilityReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


