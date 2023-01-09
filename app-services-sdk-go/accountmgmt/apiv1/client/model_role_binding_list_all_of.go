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

// RoleBindingListAllOf struct for RoleBindingListAllOf
type RoleBindingListAllOf struct {
	Items []RoleBinding `json:"items,omitempty"`
}

// NewRoleBindingListAllOf instantiates a new RoleBindingListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBindingListAllOf() *RoleBindingListAllOf {
	this := RoleBindingListAllOf{}
	return &this
}

// NewRoleBindingListAllOfWithDefaults instantiates a new RoleBindingListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingListAllOfWithDefaults() *RoleBindingListAllOf {
	this := RoleBindingListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RoleBindingListAllOf) GetItems() []RoleBinding {
	if o == nil || isNil(o.Items) {
		var ret []RoleBinding
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingListAllOf) GetItemsOk() ([]RoleBinding, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RoleBindingListAllOf) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RoleBinding and assigns it to the Items field.
func (o *RoleBindingListAllOf) SetItems(v []RoleBinding) {
	o.Items = v
}

func (o RoleBindingListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBindingListAllOf struct {
	value *RoleBindingListAllOf
	isSet bool
}

func (v NullableRoleBindingListAllOf) Get() *RoleBindingListAllOf {
	return v.value
}

func (v *NullableRoleBindingListAllOf) Set(val *RoleBindingListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBindingListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBindingListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBindingListAllOf(val *RoleBindingListAllOf) *NullableRoleBindingListAllOf {
	return &NullableRoleBindingListAllOf{value: val, isSet: true}
}

func (v NullableRoleBindingListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBindingListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


