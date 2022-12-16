# AcsClientRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RedirectUris** | **[]string** |  | 
**OrgId** | **string** |  | 

## Methods

### NewAcsClientRequestData

`func NewAcsClientRequestData(redirectUris []string, orgId string, ) *AcsClientRequestData`

NewAcsClientRequestData instantiates a new AcsClientRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcsClientRequestDataWithDefaults

`func NewAcsClientRequestDataWithDefaults() *AcsClientRequestData`

NewAcsClientRequestDataWithDefaults instantiates a new AcsClientRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AcsClientRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcsClientRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcsClientRequestData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcsClientRequestData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedirectUris

`func (o *AcsClientRequestData) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *AcsClientRequestData) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *AcsClientRequestData) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetOrgId

`func (o *AcsClientRequestData) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AcsClientRequestData) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AcsClientRequestData) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


