/*
Service Registry API

Service Registry Instance API  NOTE: This API cannot be called directly from the portal.

API version: 2.2.5.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// ArtifactReference A reference to a different artifact. Typically used with artifact types that can have dependencies like Protobuf.
type ArtifactReference struct {
	GroupId string `json:"groupId"`
	ArtifactId string `json:"artifactId"`
	Version *string `json:"version,omitempty"`
	Name string `json:"name"`
}

// NewArtifactReference instantiates a new ArtifactReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactReference(groupId string, artifactId string, name string) *ArtifactReference {
	this := ArtifactReference{}
	this.GroupId = groupId
	this.ArtifactId = artifactId
	this.Name = name
	return &this
}

// NewArtifactReferenceWithDefaults instantiates a new ArtifactReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactReferenceWithDefaults() *ArtifactReference {
	this := ArtifactReference{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ArtifactReference) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ArtifactReference) SetGroupId(v string) {
	o.GroupId = v
}

// GetArtifactId returns the ArtifactId field value
func (o *ArtifactReference) GetArtifactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetArtifactIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ArtifactId, true
}

// SetArtifactId sets field value
func (o *ArtifactReference) SetArtifactId(v string) {
	o.ArtifactId = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ArtifactReference) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ArtifactReference) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ArtifactReference) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value
func (o *ArtifactReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArtifactReference) SetName(v string) {
	o.Name = v
}

func (o ArtifactReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactReference struct {
	value *ArtifactReference
	isSet bool
}

func (v NullableArtifactReference) Get() *ArtifactReference {
	return v.value
}

func (v *NullableArtifactReference) Set(val *ArtifactReference) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactReference) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactReference(val *ArtifactReference) *NullableArtifactReference {
	return &NullableArtifactReference{value: val, isSet: true}
}

func (v NullableArtifactReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

