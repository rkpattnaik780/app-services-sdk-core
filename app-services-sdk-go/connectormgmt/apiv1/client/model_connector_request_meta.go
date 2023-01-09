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

// ConnectorRequestMeta struct for ConnectorRequestMeta
type ConnectorRequestMeta struct {
	Name string `json:"name"`
	ConnectorTypeId string `json:"connector_type_id"`
	NamespaceId string `json:"namespace_id"`
	Channel *Channel `json:"channel,omitempty"`
	DesiredState ConnectorDesiredState `json:"desired_state"`
	// Name-value string annotations for resource
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewConnectorRequestMeta instantiates a new ConnectorRequestMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRequestMeta(name string, connectorTypeId string, namespaceId string, desiredState ConnectorDesiredState) *ConnectorRequestMeta {
	this := ConnectorRequestMeta{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	this.NamespaceId = namespaceId
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	this.DesiredState = desiredState
	return &this
}

// NewConnectorRequestMetaWithDefaults instantiates a new ConnectorRequestMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRequestMetaWithDefaults() *ConnectorRequestMeta {
	this := ConnectorRequestMeta{}
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	return &this
}

// GetName returns the Name field value
func (o *ConnectorRequestMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorRequestMeta) SetName(v string) {
	o.Name = v
}

// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *ConnectorRequestMeta) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *ConnectorRequestMeta) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}

// GetNamespaceId returns the NamespaceId field value
func (o *ConnectorRequestMeta) GetNamespaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetNamespaceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NamespaceId, true
}

// SetNamespaceId sets field value
func (o *ConnectorRequestMeta) SetNamespaceId(v string) {
	o.NamespaceId = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ConnectorRequestMeta) GetChannel() Channel {
	if o == nil || isNil(o.Channel) {
		var ret Channel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetChannelOk() (*Channel, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ConnectorRequestMeta) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given Channel and assigns it to the Channel field.
func (o *ConnectorRequestMeta) SetChannel(v Channel) {
	o.Channel = &v
}

// GetDesiredState returns the DesiredState field value
func (o *ConnectorRequestMeta) GetDesiredState() ConnectorDesiredState {
	if o == nil {
		var ret ConnectorDesiredState
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetDesiredStateOk() (*ConnectorDesiredState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *ConnectorRequestMeta) SetDesiredState(v ConnectorDesiredState) {
	o.DesiredState = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorRequestMeta) GetAnnotations() map[string]string {
	if o == nil || isNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorRequestMeta) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorRequestMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o ConnectorRequestMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["connector_type_id"] = o.ConnectorTypeId
	}
	if true {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if true {
		toSerialize["desired_state"] = o.DesiredState
	}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorRequestMeta struct {
	value *ConnectorRequestMeta
	isSet bool
}

func (v NullableConnectorRequestMeta) Get() *ConnectorRequestMeta {
	return v.value
}

func (v *NullableConnectorRequestMeta) Set(val *ConnectorRequestMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRequestMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRequestMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRequestMeta(val *ConnectorRequestMeta) *NullableConnectorRequestMeta {
	return &NullableConnectorRequestMeta{value: val, isSet: true}
}

func (v NullableConnectorRequestMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRequestMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


