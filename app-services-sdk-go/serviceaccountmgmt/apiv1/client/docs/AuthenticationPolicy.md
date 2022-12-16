# AuthenticationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationFactors** | Pointer to [**AuthenticationFactors**](AuthenticationFactors.md) |  | [optional] 

## Methods

### NewAuthenticationPolicy

`func NewAuthenticationPolicy() *AuthenticationPolicy`

NewAuthenticationPolicy instantiates a new AuthenticationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyWithDefaults

`func NewAuthenticationPolicyWithDefaults() *AuthenticationPolicy`

NewAuthenticationPolicyWithDefaults instantiates a new AuthenticationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationFactors

`func (o *AuthenticationPolicy) GetAuthenticationFactors() AuthenticationFactors`

GetAuthenticationFactors returns the AuthenticationFactors field if non-nil, zero value otherwise.

### GetAuthenticationFactorsOk

`func (o *AuthenticationPolicy) GetAuthenticationFactorsOk() (*AuthenticationFactors, bool)`

GetAuthenticationFactorsOk returns a tuple with the AuthenticationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFactors

`func (o *AuthenticationPolicy) SetAuthenticationFactors(v AuthenticationFactors)`

SetAuthenticationFactors sets AuthenticationFactors field to given value.

### HasAuthenticationFactors

`func (o *AuthenticationPolicy) HasAuthenticationFactors() bool`

HasAuthenticationFactors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


