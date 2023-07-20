/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.4.x
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// NamedLogConfiguration struct for NamedLogConfiguration
type NamedLogConfiguration struct {
	Name string `json:"name"`
	Level LogLevel `json:"level"`
}

// NewNamedLogConfiguration instantiates a new NamedLogConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedLogConfiguration(name string, level LogLevel) *NamedLogConfiguration {
	this := NamedLogConfiguration{}
	this.Name = name
	this.Level = level
	return &this
}

// NewNamedLogConfigurationWithDefaults instantiates a new NamedLogConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedLogConfigurationWithDefaults() *NamedLogConfiguration {
	this := NamedLogConfiguration{}
	return &this
}

// GetName returns the Name field value
func (o *NamedLogConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NamedLogConfiguration) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NamedLogConfiguration) SetName(v string) {
	o.Name = v
}

// GetLevel returns the Level field value
func (o *NamedLogConfiguration) GetLevel() LogLevel {
	if o == nil {
		var ret LogLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *NamedLogConfiguration) GetLevelOk() (*LogLevel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *NamedLogConfiguration) SetLevel(v LogLevel) {
	o.Level = v
}

func (o NamedLogConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableNamedLogConfiguration struct {
	value *NamedLogConfiguration
	isSet bool
}

func (v NullableNamedLogConfiguration) Get() *NamedLogConfiguration {
	return v.value
}

func (v *NullableNamedLogConfiguration) Set(val *NamedLogConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedLogConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedLogConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedLogConfiguration(val *NamedLogConfiguration) *NullableNamedLogConfiguration {
	return &NullableNamedLogConfiguration{value: val, isSet: true}
}

func (v NullableNamedLogConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedLogConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


