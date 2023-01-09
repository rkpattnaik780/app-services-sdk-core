# SSOHealthResultError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackTrace** | Pointer to [**[]SSOHealthResultErrorStackTraceInner**](SSOHealthResultErrorStackTraceInner.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to [**[]SSOHealthResultErrorSuppressedInner**](SSOHealthResultErrorSuppressedInner.md) |  | [optional] 
**LocalizedMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewSSOHealthResultError

`func NewSSOHealthResultError() *SSOHealthResultError`

NewSSOHealthResultError instantiates a new SSOHealthResultError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOHealthResultErrorWithDefaults

`func NewSSOHealthResultErrorWithDefaults() *SSOHealthResultError`

NewSSOHealthResultErrorWithDefaults instantiates a new SSOHealthResultError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackTrace

`func (o *SSOHealthResultError) GetStackTrace() []SSOHealthResultErrorStackTraceInner`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *SSOHealthResultError) GetStackTraceOk() (*[]SSOHealthResultErrorStackTraceInner, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *SSOHealthResultError) SetStackTrace(v []SSOHealthResultErrorStackTraceInner)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *SSOHealthResultError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetMessage

`func (o *SSOHealthResultError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SSOHealthResultError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SSOHealthResultError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SSOHealthResultError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuppressed

`func (o *SSOHealthResultError) GetSuppressed() []SSOHealthResultErrorSuppressedInner`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *SSOHealthResultError) GetSuppressedOk() (*[]SSOHealthResultErrorSuppressedInner, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *SSOHealthResultError) SetSuppressed(v []SSOHealthResultErrorSuppressedInner)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *SSOHealthResultError) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetLocalizedMessage

`func (o *SSOHealthResultError) GetLocalizedMessage() string`

GetLocalizedMessage returns the LocalizedMessage field if non-nil, zero value otherwise.

### GetLocalizedMessageOk

`func (o *SSOHealthResultError) GetLocalizedMessageOk() (*string, bool)`

GetLocalizedMessageOk returns a tuple with the LocalizedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMessage

`func (o *SSOHealthResultError) SetLocalizedMessage(v string)`

SetLocalizedMessage sets LocalizedMessage field to given value.

### HasLocalizedMessage

`func (o *SSOHealthResultError) HasLocalizedMessage() bool`

HasLocalizedMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


