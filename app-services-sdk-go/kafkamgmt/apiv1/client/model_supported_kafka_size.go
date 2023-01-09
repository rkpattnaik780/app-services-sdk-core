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

// SupportedKafkaSize Supported Kafka Size
type SupportedKafkaSize struct {
	// Unique identifier of this Kafka instance size.
	Id *string `json:"id,omitempty"`
	// Display name of this Kafka instance size.
	DisplayName *string `json:"display_name,omitempty"`
	IngressThroughputPerSec SupportedKafkaSizeBytesValueItem `json:"ingress_throughput_per_sec,omitempty"`
	EgressThroughputPerSec SupportedKafkaSizeBytesValueItem `json:"egress_throughput_per_sec,omitempty"`
	// Maximum amount of total connections available to this Kafka instance size.
	TotalMaxConnections *int32 `json:"total_max_connections,omitempty"`
	MaxDataRetentionSize SupportedKafkaSizeBytesValueItem `json:"max_data_retention_size,omitempty"`
	// Maximum amount of total partitions available to this Kafka instance size.
	MaxPartitions *int32 `json:"max_partitions,omitempty"`
	// Maximum data retention period available to this Kafka instance size.
	MaxDataRetentionPeriod *string `json:"max_data_retention_period,omitempty"`
	// Maximium connection attempts per second available to this Kafka instance size.
	MaxConnectionAttemptsPerSec *int32 `json:"max_connection_attempts_per_sec,omitempty"`
	MaxMessageSize SupportedKafkaSizeBytesValueItem `json:"max_message_size,omitempty"`
	// Minimum number of in-sync replicas.
	MinInSyncReplicas *int32 `json:"min_in_sync_replicas,omitempty"`
	// Replication factor available to this Kafka instance size.
	ReplicationFactor *int32 `json:"replication_factor,omitempty"`
	// List of Availability Zone modes that this Kafka instance size supports. The possible values are \"single\", \"multi\".
	SupportedAzModes []string `json:"supported_az_modes,omitempty"`
	// The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires.
	LifespanSeconds NullableInt32 `json:"lifespan_seconds,omitempty"`
	// Quota consumed by this Kafka instance size.
	QuotaConsumed *int32 `json:"quota_consumed,omitempty"`
	// Quota type used by this Kafka instance size. This is now deprecated, please refer to supported_billing_models at instance-type level instead.
	// Deprecated
	QuotaType *string `json:"quota_type,omitempty"`
	// Data plane cluster capacity consumed by this Kafka instance size.
	CapacityConsumed *int32 `json:"capacity_consumed,omitempty"`
	// Maturity level of the size. Can be \"stable\" or \"preview\".
	MaturityStatus *string `json:"maturity_status,omitempty"`
}

// NewSupportedKafkaSize instantiates a new SupportedKafkaSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedKafkaSize() *SupportedKafkaSize {
	this := SupportedKafkaSize{}
	return &this
}

// NewSupportedKafkaSizeWithDefaults instantiates a new SupportedKafkaSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedKafkaSizeWithDefaults() *SupportedKafkaSize {
	this := SupportedKafkaSize{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportedKafkaSize) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SupportedKafkaSize) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIngressThroughputPerSec returns the IngressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetIngressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || isNil(o.IngressThroughputPerSec) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.IngressThroughputPerSec
}

// GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetIngressThroughputPerSecOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || isNil(o.IngressThroughputPerSec) {
    return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.IngressThroughputPerSec, true
}

// HasIngressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasIngressThroughputPerSec() bool {
	if o != nil && !isNil(o.IngressThroughputPerSec) {
		return true
	}

	return false
}

// SetIngressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the IngressThroughputPerSec field.
func (o *SupportedKafkaSize) SetIngressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.IngressThroughputPerSec = v
}

// GetEgressThroughputPerSec returns the EgressThroughputPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetEgressThroughputPerSec() SupportedKafkaSizeBytesValueItem {
	if o == nil || isNil(o.EgressThroughputPerSec) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.EgressThroughputPerSec
}

// GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetEgressThroughputPerSecOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || isNil(o.EgressThroughputPerSec) {
    return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.EgressThroughputPerSec, true
}

// HasEgressThroughputPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasEgressThroughputPerSec() bool {
	if o != nil && !isNil(o.EgressThroughputPerSec) {
		return true
	}

	return false
}

// SetEgressThroughputPerSec gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the EgressThroughputPerSec field.
func (o *SupportedKafkaSize) SetEgressThroughputPerSec(v SupportedKafkaSizeBytesValueItem) {
	o.EgressThroughputPerSec = v
}

// GetTotalMaxConnections returns the TotalMaxConnections field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetTotalMaxConnections() int32 {
	if o == nil || isNil(o.TotalMaxConnections) {
		var ret int32
		return ret
	}
	return *o.TotalMaxConnections
}

// GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetTotalMaxConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalMaxConnections) {
    return nil, false
	}
	return o.TotalMaxConnections, true
}

// HasTotalMaxConnections returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasTotalMaxConnections() bool {
	if o != nil && !isNil(o.TotalMaxConnections) {
		return true
	}

	return false
}

// SetTotalMaxConnections gets a reference to the given int32 and assigns it to the TotalMaxConnections field.
func (o *SupportedKafkaSize) SetTotalMaxConnections(v int32) {
	o.TotalMaxConnections = &v
}

// GetMaxDataRetentionSize returns the MaxDataRetentionSize field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem {
	if o == nil || isNil(o.MaxDataRetentionSize) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.MaxDataRetentionSize
}

// GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxDataRetentionSizeOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || isNil(o.MaxDataRetentionSize) {
    return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.MaxDataRetentionSize, true
}

// HasMaxDataRetentionSize returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxDataRetentionSize() bool {
	if o != nil && !isNil(o.MaxDataRetentionSize) {
		return true
	}

	return false
}

// SetMaxDataRetentionSize gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the MaxDataRetentionSize field.
func (o *SupportedKafkaSize) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem) {
	o.MaxDataRetentionSize = v
}

// GetMaxPartitions returns the MaxPartitions field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxPartitions() int32 {
	if o == nil || isNil(o.MaxPartitions) {
		var ret int32
		return ret
	}
	return *o.MaxPartitions
}

// GetMaxPartitionsOk returns a tuple with the MaxPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxPartitionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPartitions) {
    return nil, false
	}
	return o.MaxPartitions, true
}

// HasMaxPartitions returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxPartitions() bool {
	if o != nil && !isNil(o.MaxPartitions) {
		return true
	}

	return false
}

// SetMaxPartitions gets a reference to the given int32 and assigns it to the MaxPartitions field.
func (o *SupportedKafkaSize) SetMaxPartitions(v int32) {
	o.MaxPartitions = &v
}

// GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxDataRetentionPeriod() string {
	if o == nil || isNil(o.MaxDataRetentionPeriod) {
		var ret string
		return ret
	}
	return *o.MaxDataRetentionPeriod
}

// GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxDataRetentionPeriodOk() (*string, bool) {
	if o == nil || isNil(o.MaxDataRetentionPeriod) {
    return nil, false
	}
	return o.MaxDataRetentionPeriod, true
}

// HasMaxDataRetentionPeriod returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxDataRetentionPeriod() bool {
	if o != nil && !isNil(o.MaxDataRetentionPeriod) {
		return true
	}

	return false
}

// SetMaxDataRetentionPeriod gets a reference to the given string and assigns it to the MaxDataRetentionPeriod field.
func (o *SupportedKafkaSize) SetMaxDataRetentionPeriod(v string) {
	o.MaxDataRetentionPeriod = &v
}

// GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSec() int32 {
	if o == nil || isNil(o.MaxConnectionAttemptsPerSec) {
		var ret int32
		return ret
	}
	return *o.MaxConnectionAttemptsPerSec
}

// GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSecOk() (*int32, bool) {
	if o == nil || isNil(o.MaxConnectionAttemptsPerSec) {
    return nil, false
	}
	return o.MaxConnectionAttemptsPerSec, true
}

// HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxConnectionAttemptsPerSec() bool {
	if o != nil && !isNil(o.MaxConnectionAttemptsPerSec) {
		return true
	}

	return false
}

// SetMaxConnectionAttemptsPerSec gets a reference to the given int32 and assigns it to the MaxConnectionAttemptsPerSec field.
func (o *SupportedKafkaSize) SetMaxConnectionAttemptsPerSec(v int32) {
	o.MaxConnectionAttemptsPerSec = &v
}

// GetMaxMessageSize returns the MaxMessageSize field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaxMessageSize() SupportedKafkaSizeBytesValueItem {
	if o == nil || isNil(o.MaxMessageSize) {
		var ret SupportedKafkaSizeBytesValueItem
		return ret
	}
	return o.MaxMessageSize
}

// GetMaxMessageSizeOk returns a tuple with the MaxMessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaxMessageSizeOk() (SupportedKafkaSizeBytesValueItem, bool) {
	if o == nil || isNil(o.MaxMessageSize) {
    return SupportedKafkaSizeBytesValueItem{}, false
	}
	return o.MaxMessageSize, true
}

// HasMaxMessageSize returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaxMessageSize() bool {
	if o != nil && !isNil(o.MaxMessageSize) {
		return true
	}

	return false
}

// SetMaxMessageSize gets a reference to the given SupportedKafkaSizeBytesValueItem and assigns it to the MaxMessageSize field.
func (o *SupportedKafkaSize) SetMaxMessageSize(v SupportedKafkaSizeBytesValueItem) {
	o.MaxMessageSize = v
}

// GetMinInSyncReplicas returns the MinInSyncReplicas field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMinInSyncReplicas() int32 {
	if o == nil || isNil(o.MinInSyncReplicas) {
		var ret int32
		return ret
	}
	return *o.MinInSyncReplicas
}

// GetMinInSyncReplicasOk returns a tuple with the MinInSyncReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMinInSyncReplicasOk() (*int32, bool) {
	if o == nil || isNil(o.MinInSyncReplicas) {
    return nil, false
	}
	return o.MinInSyncReplicas, true
}

// HasMinInSyncReplicas returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMinInSyncReplicas() bool {
	if o != nil && !isNil(o.MinInSyncReplicas) {
		return true
	}

	return false
}

// SetMinInSyncReplicas gets a reference to the given int32 and assigns it to the MinInSyncReplicas field.
func (o *SupportedKafkaSize) SetMinInSyncReplicas(v int32) {
	o.MinInSyncReplicas = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetReplicationFactor() int32 {
	if o == nil || isNil(o.ReplicationFactor) {
		var ret int32
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetReplicationFactorOk() (*int32, bool) {
	if o == nil || isNil(o.ReplicationFactor) {
    return nil, false
	}
	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasReplicationFactor() bool {
	if o != nil && !isNil(o.ReplicationFactor) {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int32 and assigns it to the ReplicationFactor field.
func (o *SupportedKafkaSize) SetReplicationFactor(v int32) {
	o.ReplicationFactor = &v
}

// GetSupportedAzModes returns the SupportedAzModes field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetSupportedAzModes() []string {
	if o == nil || isNil(o.SupportedAzModes) {
		var ret []string
		return ret
	}
	return o.SupportedAzModes
}

// GetSupportedAzModesOk returns a tuple with the SupportedAzModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetSupportedAzModesOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedAzModes) {
    return nil, false
	}
	return o.SupportedAzModes, true
}

// HasSupportedAzModes returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasSupportedAzModes() bool {
	if o != nil && !isNil(o.SupportedAzModes) {
		return true
	}

	return false
}

// SetSupportedAzModes gets a reference to the given []string and assigns it to the SupportedAzModes field.
func (o *SupportedKafkaSize) SetSupportedAzModes(v []string) {
	o.SupportedAzModes = v
}

// GetLifespanSeconds returns the LifespanSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SupportedKafkaSize) GetLifespanSeconds() int32 {
	if o == nil || isNil(o.LifespanSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.LifespanSeconds.Get()
}

// GetLifespanSecondsOk returns a tuple with the LifespanSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SupportedKafkaSize) GetLifespanSecondsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.LifespanSeconds.Get(), o.LifespanSeconds.IsSet()
}

// HasLifespanSeconds returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasLifespanSeconds() bool {
	if o != nil && o.LifespanSeconds.IsSet() {
		return true
	}

	return false
}

// SetLifespanSeconds gets a reference to the given NullableInt32 and assigns it to the LifespanSeconds field.
func (o *SupportedKafkaSize) SetLifespanSeconds(v int32) {
	o.LifespanSeconds.Set(&v)
}
// SetLifespanSecondsNil sets the value for LifespanSeconds to be an explicit nil
func (o *SupportedKafkaSize) SetLifespanSecondsNil() {
	o.LifespanSeconds.Set(nil)
}

// UnsetLifespanSeconds ensures that no value is present for LifespanSeconds, not even an explicit nil
func (o *SupportedKafkaSize) UnsetLifespanSeconds() {
	o.LifespanSeconds.Unset()
}

// GetQuotaConsumed returns the QuotaConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetQuotaConsumed() int32 {
	if o == nil || isNil(o.QuotaConsumed) {
		var ret int32
		return ret
	}
	return *o.QuotaConsumed
}

// GetQuotaConsumedOk returns a tuple with the QuotaConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetQuotaConsumedOk() (*int32, bool) {
	if o == nil || isNil(o.QuotaConsumed) {
    return nil, false
	}
	return o.QuotaConsumed, true
}

// HasQuotaConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasQuotaConsumed() bool {
	if o != nil && !isNil(o.QuotaConsumed) {
		return true
	}

	return false
}

// SetQuotaConsumed gets a reference to the given int32 and assigns it to the QuotaConsumed field.
func (o *SupportedKafkaSize) SetQuotaConsumed(v int32) {
	o.QuotaConsumed = &v
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
// Deprecated
func (o *SupportedKafkaSize) GetQuotaType() string {
	if o == nil || isNil(o.QuotaType) {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SupportedKafkaSize) GetQuotaTypeOk() (*string, bool) {
	if o == nil || isNil(o.QuotaType) {
    return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasQuotaType() bool {
	if o != nil && !isNil(o.QuotaType) {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
// Deprecated
func (o *SupportedKafkaSize) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetCapacityConsumed returns the CapacityConsumed field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetCapacityConsumed() int32 {
	if o == nil || isNil(o.CapacityConsumed) {
		var ret int32
		return ret
	}
	return *o.CapacityConsumed
}

// GetCapacityConsumedOk returns a tuple with the CapacityConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetCapacityConsumedOk() (*int32, bool) {
	if o == nil || isNil(o.CapacityConsumed) {
    return nil, false
	}
	return o.CapacityConsumed, true
}

// HasCapacityConsumed returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasCapacityConsumed() bool {
	if o != nil && !isNil(o.CapacityConsumed) {
		return true
	}

	return false
}

// SetCapacityConsumed gets a reference to the given int32 and assigns it to the CapacityConsumed field.
func (o *SupportedKafkaSize) SetCapacityConsumed(v int32) {
	o.CapacityConsumed = &v
}

// GetMaturityStatus returns the MaturityStatus field value if set, zero value otherwise.
func (o *SupportedKafkaSize) GetMaturityStatus() string {
	if o == nil || isNil(o.MaturityStatus) {
		var ret string
		return ret
	}
	return *o.MaturityStatus
}

// GetMaturityStatusOk returns a tuple with the MaturityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaSize) GetMaturityStatusOk() (*string, bool) {
	if o == nil || isNil(o.MaturityStatus) {
    return nil, false
	}
	return o.MaturityStatus, true
}

// HasMaturityStatus returns a boolean if a field has been set.
func (o *SupportedKafkaSize) HasMaturityStatus() bool {
	if o != nil && !isNil(o.MaturityStatus) {
		return true
	}

	return false
}

// SetMaturityStatus gets a reference to the given string and assigns it to the MaturityStatus field.
func (o *SupportedKafkaSize) SetMaturityStatus(v string) {
	o.MaturityStatus = &v
}

func (o SupportedKafkaSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.IngressThroughputPerSec) {
		toSerialize["ingress_throughput_per_sec"] = o.IngressThroughputPerSec
	}
	if !isNil(o.EgressThroughputPerSec) {
		toSerialize["egress_throughput_per_sec"] = o.EgressThroughputPerSec
	}
	if !isNil(o.TotalMaxConnections) {
		toSerialize["total_max_connections"] = o.TotalMaxConnections
	}
	if !isNil(o.MaxDataRetentionSize) {
		toSerialize["max_data_retention_size"] = o.MaxDataRetentionSize
	}
	if !isNil(o.MaxPartitions) {
		toSerialize["max_partitions"] = o.MaxPartitions
	}
	if !isNil(o.MaxDataRetentionPeriod) {
		toSerialize["max_data_retention_period"] = o.MaxDataRetentionPeriod
	}
	if !isNil(o.MaxConnectionAttemptsPerSec) {
		toSerialize["max_connection_attempts_per_sec"] = o.MaxConnectionAttemptsPerSec
	}
	if !isNil(o.MaxMessageSize) {
		toSerialize["max_message_size"] = o.MaxMessageSize
	}
	if !isNil(o.MinInSyncReplicas) {
		toSerialize["min_in_sync_replicas"] = o.MinInSyncReplicas
	}
	if !isNil(o.ReplicationFactor) {
		toSerialize["replication_factor"] = o.ReplicationFactor
	}
	if !isNil(o.SupportedAzModes) {
		toSerialize["supported_az_modes"] = o.SupportedAzModes
	}
	if o.LifespanSeconds.IsSet() {
		toSerialize["lifespan_seconds"] = o.LifespanSeconds.Get()
	}
	if !isNil(o.QuotaConsumed) {
		toSerialize["quota_consumed"] = o.QuotaConsumed
	}
	if !isNil(o.QuotaType) {
		toSerialize["quota_type"] = o.QuotaType
	}
	if !isNil(o.CapacityConsumed) {
		toSerialize["capacity_consumed"] = o.CapacityConsumed
	}
	if !isNil(o.MaturityStatus) {
		toSerialize["maturity_status"] = o.MaturityStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSupportedKafkaSize struct {
	value *SupportedKafkaSize
	isSet bool
}

func (v NullableSupportedKafkaSize) Get() *SupportedKafkaSize {
	return v.value
}

func (v *NullableSupportedKafkaSize) Set(val *SupportedKafkaSize) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedKafkaSize) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedKafkaSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedKafkaSize(val *SupportedKafkaSize) *NullableSupportedKafkaSize {
	return &NullableSupportedKafkaSize{value: val, isSet: true}
}

func (v NullableSupportedKafkaSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedKafkaSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


