/*
Kafka Management API

Kafka Management API is a REST API to manage Kafka instances

API version: 1.14.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// VersionMetadata struct for VersionMetadata
type VersionMetadata struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	ServerVersion *string `json:"server_version,omitempty"`
	Collections []ObjectReference `json:"collections,omitempty"`
}

// NewVersionMetadata instantiates a new VersionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionMetadata(id string, kind string, href string) *VersionMetadata {
	this := VersionMetadata{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	return &this
}

// NewVersionMetadataWithDefaults instantiates a new VersionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionMetadataWithDefaults() *VersionMetadata {
	this := VersionMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *VersionMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VersionMetadata) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *VersionMetadata) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetKindOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *VersionMetadata) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *VersionMetadata) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *VersionMetadata) SetHref(v string) {
	o.Href = v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *VersionMetadata) GetServerVersion() string {
	if o == nil || isNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetServerVersionOk() (*string, bool) {
	if o == nil || isNil(o.ServerVersion) {
    return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *VersionMetadata) HasServerVersion() bool {
	if o != nil && !isNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *VersionMetadata) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *VersionMetadata) GetCollections() []ObjectReference {
	if o == nil || isNil(o.Collections) {
		var ret []ObjectReference
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetCollectionsOk() ([]ObjectReference, bool) {
	if o == nil || isNil(o.Collections) {
    return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *VersionMetadata) HasCollections() bool {
	if o != nil && !isNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []ObjectReference and assigns it to the Collections field.
func (o *VersionMetadata) SetCollections(v []ObjectReference) {
	o.Collections = v
}

func (o VersionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.ServerVersion) {
		toSerialize["server_version"] = o.ServerVersion
	}
	if !isNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	return json.Marshal(toSerialize)
}

type NullableVersionMetadata struct {
	value *VersionMetadata
	isSet bool
}

func (v NullableVersionMetadata) Get() *VersionMetadata {
	return v.value
}

func (v *NullableVersionMetadata) Set(val *VersionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionMetadata(val *VersionMetadata) *NullableVersionMetadata {
	return &NullableVersionMetadata{value: val, isSet: true}
}

func (v NullableVersionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


