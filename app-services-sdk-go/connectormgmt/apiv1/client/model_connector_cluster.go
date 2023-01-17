/*
Connector Management API

Connector Management API is a REST API to manage connectors.

API version: 0.1.0
Contact: rhosak-support@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorCluster struct for ConnectorCluster
type ConnectorCluster struct {
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
	Owner *string `json:"owner,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name *string `json:"name,omitempty"`
	// Name-value string annotations for resource
	Annotations *map[string]string `json:"annotations,omitempty"`
	Status *ConnectorClusterStatusStatus `json:"status,omitempty"`
}

// NewConnectorCluster instantiates a new ConnectorCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorCluster() *ConnectorCluster {
	this := ConnectorCluster{}
	return &this
}

// NewConnectorClusterWithDefaults instantiates a new ConnectorCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterWithDefaults() *ConnectorCluster {
	this := ConnectorCluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectorCluster) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectorCluster) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectorCluster) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConnectorCluster) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConnectorCluster) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConnectorCluster) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConnectorCluster) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConnectorCluster) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConnectorCluster) SetHref(v string) {
	o.Href = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConnectorCluster) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConnectorCluster) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ConnectorCluster) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectorCluster) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectorCluster) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectorCluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ConnectorCluster) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ConnectorCluster) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ConnectorCluster) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorCluster) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorCluster) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorCluster) SetName(v string) {
	o.Name = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorCluster) GetAnnotations() map[string]string {
	if o == nil || isNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorCluster) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorCluster) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectorCluster) GetStatus() ConnectorClusterStatusStatus {
	if o == nil || isNil(o.Status) {
		var ret ConnectorClusterStatusStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorCluster) GetStatusOk() (*ConnectorClusterStatusStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectorCluster) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ConnectorClusterStatusStatus and assigns it to the Status field.
func (o *ConnectorCluster) SetStatus(v ConnectorClusterStatusStatus) {
	o.Status = &v
}

func (o ConnectorCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorCluster struct {
	value *ConnectorCluster
	isSet bool
}

func (v NullableConnectorCluster) Get() *ConnectorCluster {
	return v.value
}

func (v *NullableConnectorCluster) Set(val *ConnectorCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorCluster(val *ConnectorCluster) *NullableConnectorCluster {
	return &NullableConnectorCluster{value: val, isSet: true}
}

func (v NullableConnectorCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

