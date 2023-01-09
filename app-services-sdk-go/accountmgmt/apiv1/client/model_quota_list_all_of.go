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

// QuotaListAllOf struct for QuotaListAllOf
type QuotaListAllOf struct {
	Items []Quota `json:"items,omitempty"`
}

// NewQuotaListAllOf instantiates a new QuotaListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListAllOf() *QuotaListAllOf {
	this := QuotaListAllOf{}
	return &this
}

// NewQuotaListAllOfWithDefaults instantiates a new QuotaListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListAllOfWithDefaults() *QuotaListAllOf {
	this := QuotaListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QuotaListAllOf) GetItems() []Quota {
	if o == nil || isNil(o.Items) {
		var ret []Quota
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaListAllOf) GetItemsOk() ([]Quota, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QuotaListAllOf) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Quota and assigns it to the Items field.
func (o *QuotaListAllOf) SetItems(v []Quota) {
	o.Items = v
}

func (o QuotaListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaListAllOf struct {
	value *QuotaListAllOf
	isSet bool
}

func (v NullableQuotaListAllOf) Get() *QuotaListAllOf {
	return v.value
}

func (v *NullableQuotaListAllOf) Set(val *QuotaListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListAllOf(val *QuotaListAllOf) *NullableQuotaListAllOf {
	return &NullableQuotaListAllOf{value: val, isSet: true}
}

func (v NullableQuotaListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


