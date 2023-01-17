/*
Service Registry Management API

Service Registry Management API is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.

API version: 0.0.6
Contact: rhosak-eval-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
	"time"
)

// RootTypeForRegistry Service Registry instance in a multi-tenant deployment.
type RootTypeForRegistry struct {
	Id string `json:"id"`
	Status RegistryStatusValue `json:"status"`
	RegistryUrl *string `json:"registryUrl,omitempty"`
	BrowserUrl *string `json:"browserUrl,omitempty"`
	// User-defined Registry instance name. Does not have to be unique.
	Name string `json:"name"`
	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`
	// Registry instance owner.
	Owner string `json:"owner"`
	// Description of the Registry instance.
	Description *string `json:"description,omitempty"`
	// ISO 8601 UTC timestamp.
	CreatedAt time.Time `json:"created_at"`
	// ISO 8601 UTC timestamp.
	UpdatedAt time.Time `json:"updated_at"`
	InstanceType RegistryInstanceTypeValue `json:"instance_type"`
}

// NewRootTypeForRegistry instantiates a new RootTypeForRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootTypeForRegistry(id string, status RegistryStatusValue, name string, owner string, createdAt time.Time, updatedAt time.Time, instanceType RegistryInstanceTypeValue) *RootTypeForRegistry {
	this := RootTypeForRegistry{}
	this.Id = id
	this.Status = status
	this.Name = name
	this.Owner = owner
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.InstanceType = instanceType
	return &this
}

// NewRootTypeForRegistryWithDefaults instantiates a new RootTypeForRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTypeForRegistryWithDefaults() *RootTypeForRegistry {
	this := RootTypeForRegistry{}
	return &this
}

// GetId returns the Id field value
func (o *RootTypeForRegistry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RootTypeForRegistry) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *RootTypeForRegistry) GetStatus() RegistryStatusValue {
	if o == nil {
		var ret RegistryStatusValue
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetStatusOk() (*RegistryStatusValue, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RootTypeForRegistry) SetStatus(v RegistryStatusValue) {
	o.Status = v
}

// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetRegistryUrl() string {
	if o == nil || isNil(o.RegistryUrl) {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetRegistryUrlOk() (*string, bool) {
	if o == nil || isNil(o.RegistryUrl) {
    return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasRegistryUrl() bool {
	if o != nil && !isNil(o.RegistryUrl) {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *RootTypeForRegistry) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}

// GetBrowserUrl returns the BrowserUrl field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetBrowserUrl() string {
	if o == nil || isNil(o.BrowserUrl) {
		var ret string
		return ret
	}
	return *o.BrowserUrl
}

// GetBrowserUrlOk returns a tuple with the BrowserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetBrowserUrlOk() (*string, bool) {
	if o == nil || isNil(o.BrowserUrl) {
    return nil, false
	}
	return o.BrowserUrl, true
}

// HasBrowserUrl returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasBrowserUrl() bool {
	if o != nil && !isNil(o.BrowserUrl) {
		return true
	}

	return false
}

// SetBrowserUrl gets a reference to the given string and assigns it to the BrowserUrl field.
func (o *RootTypeForRegistry) SetBrowserUrl(v string) {
	o.BrowserUrl = &v
}

// GetName returns the Name field value
func (o *RootTypeForRegistry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RootTypeForRegistry) SetName(v string) {
	o.Name = v
}

// GetRegistryDeploymentId returns the RegistryDeploymentId field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetRegistryDeploymentId() int32 {
	if o == nil || isNil(o.RegistryDeploymentId) {
		var ret int32
		return ret
	}
	return *o.RegistryDeploymentId
}

// GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetRegistryDeploymentIdOk() (*int32, bool) {
	if o == nil || isNil(o.RegistryDeploymentId) {
    return nil, false
	}
	return o.RegistryDeploymentId, true
}

// HasRegistryDeploymentId returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasRegistryDeploymentId() bool {
	if o != nil && !isNil(o.RegistryDeploymentId) {
		return true
	}

	return false
}

// SetRegistryDeploymentId gets a reference to the given int32 and assigns it to the RegistryDeploymentId field.
func (o *RootTypeForRegistry) SetRegistryDeploymentId(v int32) {
	o.RegistryDeploymentId = &v
}

// GetOwner returns the Owner field value
func (o *RootTypeForRegistry) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetOwnerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *RootTypeForRegistry) SetOwner(v string) {
	o.Owner = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RootTypeForRegistry) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RootTypeForRegistry) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RootTypeForRegistry) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RootTypeForRegistry) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RootTypeForRegistry) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetInstanceType returns the InstanceType field value
func (o *RootTypeForRegistry) GetInstanceType() RegistryInstanceTypeValue {
	if o == nil {
		var ret RegistryInstanceTypeValue
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetInstanceTypeOk() (*RegistryInstanceTypeValue, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *RootTypeForRegistry) SetInstanceType(v RegistryInstanceTypeValue) {
	o.InstanceType = v
}

func (o RootTypeForRegistry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.RegistryUrl) {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
	if !isNil(o.BrowserUrl) {
		toSerialize["browserUrl"] = o.BrowserUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.RegistryDeploymentId) {
		toSerialize["registryDeploymentId"] = o.RegistryDeploymentId
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["instance_type"] = o.InstanceType
	}
	return json.Marshal(toSerialize)
}

type NullableRootTypeForRegistry struct {
	value *RootTypeForRegistry
	isSet bool
}

func (v NullableRootTypeForRegistry) Get() *RootTypeForRegistry {
	return v.value
}

func (v *NullableRootTypeForRegistry) Set(val *RootTypeForRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableRootTypeForRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableRootTypeForRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootTypeForRegistry(val *RootTypeForRegistry) *NullableRootTypeForRegistry {
	return &NullableRootTypeForRegistry{value: val, isSet: true}
}

func (v NullableRootTypeForRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootTypeForRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

