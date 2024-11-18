# JellyfinUpdateUserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **NullableString** | Gets or sets the current sha1-hashed password. | [optional] 
**CurrentPw** | Pointer to **NullableString** | Gets or sets the current plain text password. | [optional] 
**NewPw** | Pointer to **NullableString** | Gets or sets the new plain text password. | [optional] 
**ResetPassword** | Pointer to **bool** | Gets or sets a value indicating whether to reset the password. | [optional] 

## Methods

### NewJellyfinUpdateUserPassword

`func NewJellyfinUpdateUserPassword() *JellyfinUpdateUserPassword`

NewJellyfinUpdateUserPassword instantiates a new JellyfinUpdateUserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUpdateUserPasswordWithDefaults

`func NewJellyfinUpdateUserPasswordWithDefaults() *JellyfinUpdateUserPassword`

NewJellyfinUpdateUserPasswordWithDefaults instantiates a new JellyfinUpdateUserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *JellyfinUpdateUserPassword) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *JellyfinUpdateUserPassword) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *JellyfinUpdateUserPassword) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *JellyfinUpdateUserPassword) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *JellyfinUpdateUserPassword) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *JellyfinUpdateUserPassword) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetCurrentPw

`func (o *JellyfinUpdateUserPassword) GetCurrentPw() string`

GetCurrentPw returns the CurrentPw field if non-nil, zero value otherwise.

### GetCurrentPwOk

`func (o *JellyfinUpdateUserPassword) GetCurrentPwOk() (*string, bool)`

GetCurrentPwOk returns a tuple with the CurrentPw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPw

`func (o *JellyfinUpdateUserPassword) SetCurrentPw(v string)`

SetCurrentPw sets CurrentPw field to given value.

### HasCurrentPw

`func (o *JellyfinUpdateUserPassword) HasCurrentPw() bool`

HasCurrentPw returns a boolean if a field has been set.

### SetCurrentPwNil

`func (o *JellyfinUpdateUserPassword) SetCurrentPwNil(b bool)`

 SetCurrentPwNil sets the value for CurrentPw to be an explicit nil

### UnsetCurrentPw
`func (o *JellyfinUpdateUserPassword) UnsetCurrentPw()`

UnsetCurrentPw ensures that no value is present for CurrentPw, not even an explicit nil
### GetNewPw

`func (o *JellyfinUpdateUserPassword) GetNewPw() string`

GetNewPw returns the NewPw field if non-nil, zero value otherwise.

### GetNewPwOk

`func (o *JellyfinUpdateUserPassword) GetNewPwOk() (*string, bool)`

GetNewPwOk returns a tuple with the NewPw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPw

`func (o *JellyfinUpdateUserPassword) SetNewPw(v string)`

SetNewPw sets NewPw field to given value.

### HasNewPw

`func (o *JellyfinUpdateUserPassword) HasNewPw() bool`

HasNewPw returns a boolean if a field has been set.

### SetNewPwNil

`func (o *JellyfinUpdateUserPassword) SetNewPwNil(b bool)`

 SetNewPwNil sets the value for NewPw to be an explicit nil

### UnsetNewPw
`func (o *JellyfinUpdateUserPassword) UnsetNewPw()`

UnsetNewPw ensures that no value is present for NewPw, not even an explicit nil
### GetResetPassword

`func (o *JellyfinUpdateUserPassword) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *JellyfinUpdateUserPassword) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *JellyfinUpdateUserPassword) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *JellyfinUpdateUserPassword) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


