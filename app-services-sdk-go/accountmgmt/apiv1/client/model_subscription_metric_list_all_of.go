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

// SubscriptionMetricListAllOf struct for SubscriptionMetricListAllOf
type SubscriptionMetricListAllOf struct {
	Items []SubscriptionMetric `json:"items,omitempty"`
}

// NewSubscriptionMetricListAllOf instantiates a new SubscriptionMetricListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionMetricListAllOf() *SubscriptionMetricListAllOf {
	this := SubscriptionMetricListAllOf{}
	return &this
}

// NewSubscriptionMetricListAllOfWithDefaults instantiates a new SubscriptionMetricListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionMetricListAllOfWithDefaults() *SubscriptionMetricListAllOf {
	this := SubscriptionMetricListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SubscriptionMetricListAllOf) GetItems() []SubscriptionMetric {
	if o == nil || isNil(o.Items) {
		var ret []SubscriptionMetric
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionMetricListAllOf) GetItemsOk() ([]SubscriptionMetric, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SubscriptionMetricListAllOf) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SubscriptionMetric and assigns it to the Items field.
func (o *SubscriptionMetricListAllOf) SetItems(v []SubscriptionMetric) {
	o.Items = v
}

func (o SubscriptionMetricListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionMetricListAllOf struct {
	value *SubscriptionMetricListAllOf
	isSet bool
}

func (v NullableSubscriptionMetricListAllOf) Get() *SubscriptionMetricListAllOf {
	return v.value
}

func (v *NullableSubscriptionMetricListAllOf) Set(val *SubscriptionMetricListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionMetricListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionMetricListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionMetricListAllOf(val *SubscriptionMetricListAllOf) *NullableSubscriptionMetricListAllOf {
	return &NullableSubscriptionMetricListAllOf{value: val, isSet: true}
}

func (v NullableSubscriptionMetricListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionMetricListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

