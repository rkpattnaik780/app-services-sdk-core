# EnterpriseClusterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKafkasViaPrivateNetwork** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | Pointer to **string** | The OCM&#39;s cluster id of the registered Enterprise cluster. | [optional] 
**Status** | Pointer to **string** | The status of Enterprise cluster registration | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider for this cluster. This valus will be used as the Kafka&#39;s cloud provider value when a Kafka is created on this cluster | [optional] 
**Region** | Pointer to **string** | The region of this cluster. This valus will be used as the Kafka&#39;s region value when a Kafka is created on this cluster | [optional] 
**MultiAz** | **bool** | A flag indicating whether this cluster is available on multiple availability zones or not | 
**SupportedInstanceTypes** | Pointer to [**SupportedKafkaInstanceTypesList**](SupportedKafkaInstanceTypesList.md) |  | [optional] 
**CapacityInformation** | Pointer to [**EnterpriseClusterAllOfCapacityInformation**](EnterpriseClusterAllOfCapacityInformation.md) |  | [optional] 

## Methods

### NewEnterpriseClusterAllOf

`func NewEnterpriseClusterAllOf(accessKafkasViaPrivateNetwork bool, multiAz bool, ) *EnterpriseClusterAllOf`

NewEnterpriseClusterAllOf instantiates a new EnterpriseClusterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterAllOfWithDefaults

`func NewEnterpriseClusterAllOfWithDefaults() *EnterpriseClusterAllOf`

NewEnterpriseClusterAllOfWithDefaults instantiates a new EnterpriseClusterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterAllOf) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseClusterAllOf) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterAllOf) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseClusterAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnterpriseClusterAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnterpriseClusterAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnterpriseClusterAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *EnterpriseClusterAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *EnterpriseClusterAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnterpriseClusterAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnterpriseClusterAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnterpriseClusterAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMultiAz

`func (o *EnterpriseClusterAllOf) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *EnterpriseClusterAllOf) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *EnterpriseClusterAllOf) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.


### GetSupportedInstanceTypes

`func (o *EnterpriseClusterAllOf) GetSupportedInstanceTypes() SupportedKafkaInstanceTypesList`

GetSupportedInstanceTypes returns the SupportedInstanceTypes field if non-nil, zero value otherwise.

### GetSupportedInstanceTypesOk

`func (o *EnterpriseClusterAllOf) GetSupportedInstanceTypesOk() (*SupportedKafkaInstanceTypesList, bool)`

GetSupportedInstanceTypesOk returns a tuple with the SupportedInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInstanceTypes

`func (o *EnterpriseClusterAllOf) SetSupportedInstanceTypes(v SupportedKafkaInstanceTypesList)`

SetSupportedInstanceTypes sets SupportedInstanceTypes field to given value.

### HasSupportedInstanceTypes

`func (o *EnterpriseClusterAllOf) HasSupportedInstanceTypes() bool`

HasSupportedInstanceTypes returns a boolean if a field has been set.

### GetCapacityInformation

`func (o *EnterpriseClusterAllOf) GetCapacityInformation() EnterpriseClusterAllOfCapacityInformation`

GetCapacityInformation returns the CapacityInformation field if non-nil, zero value otherwise.

### GetCapacityInformationOk

`func (o *EnterpriseClusterAllOf) GetCapacityInformationOk() (*EnterpriseClusterAllOfCapacityInformation, bool)`

GetCapacityInformationOk returns a tuple with the CapacityInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInformation

`func (o *EnterpriseClusterAllOf) SetCapacityInformation(v EnterpriseClusterAllOfCapacityInformation)`

SetCapacityInformation sets CapacityInformation field to given value.

### HasCapacityInformation

`func (o *EnterpriseClusterAllOf) HasCapacityInformation() bool`

HasCapacityInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


