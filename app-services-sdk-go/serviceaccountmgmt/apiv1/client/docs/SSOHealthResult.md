# SSOHealthResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**SSOHealthResultError**](SSOHealthResultError.md) |  | [optional] 
**Details** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 

## Methods

### NewSSOHealthResult

`func NewSSOHealthResult() *SSOHealthResult`

NewSSOHealthResult instantiates a new SSOHealthResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOHealthResultWithDefaults

`func NewSSOHealthResultWithDefaults() *SSOHealthResult`

NewSSOHealthResultWithDefaults instantiates a new SSOHealthResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SSOHealthResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SSOHealthResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SSOHealthResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SSOHealthResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *SSOHealthResult) GetError() SSOHealthResultError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SSOHealthResult) GetErrorOk() (*SSOHealthResultError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SSOHealthResult) SetError(v SSOHealthResultError)`

SetError sets Error field to given value.

### HasError

`func (o *SSOHealthResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *SSOHealthResult) GetDetails() map[string]map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SSOHealthResult) GetDetailsOk() (*map[string]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SSOHealthResult) SetDetails(v map[string]map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SSOHealthResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTime

`func (o *SSOHealthResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SSOHealthResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SSOHealthResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SSOHealthResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetHealthy

`func (o *SSOHealthResult) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *SSOHealthResult) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *SSOHealthResult) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *SSOHealthResult) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


