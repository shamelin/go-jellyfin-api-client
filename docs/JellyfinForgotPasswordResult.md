# JellyfinForgotPasswordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**JellyfinJellyfinForgotPasswordAction**](JellyfinForgotPasswordAction.md) | Gets or sets the action. | [optional] 
**PinFile** | Pointer to **NullableString** | Gets or sets the pin file. | [optional] 
**PinExpirationDate** | Pointer to **NullableTime** | Gets or sets the pin expiration date. | [optional] 

## Methods

### NewJellyfinForgotPasswordResult

`func NewJellyfinForgotPasswordResult() *JellyfinForgotPasswordResult`

NewJellyfinForgotPasswordResult instantiates a new JellyfinForgotPasswordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinForgotPasswordResultWithDefaults

`func NewJellyfinForgotPasswordResultWithDefaults() *JellyfinForgotPasswordResult`

NewJellyfinForgotPasswordResultWithDefaults instantiates a new JellyfinForgotPasswordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *JellyfinForgotPasswordResult) GetAction() JellyfinJellyfinForgotPasswordAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *JellyfinForgotPasswordResult) GetActionOk() (*JellyfinJellyfinForgotPasswordAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *JellyfinForgotPasswordResult) SetAction(v JellyfinJellyfinForgotPasswordAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *JellyfinForgotPasswordResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPinFile

`func (o *JellyfinForgotPasswordResult) GetPinFile() string`

GetPinFile returns the PinFile field if non-nil, zero value otherwise.

### GetPinFileOk

`func (o *JellyfinForgotPasswordResult) GetPinFileOk() (*string, bool)`

GetPinFileOk returns a tuple with the PinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFile

`func (o *JellyfinForgotPasswordResult) SetPinFile(v string)`

SetPinFile sets PinFile field to given value.

### HasPinFile

`func (o *JellyfinForgotPasswordResult) HasPinFile() bool`

HasPinFile returns a boolean if a field has been set.

### SetPinFileNil

`func (o *JellyfinForgotPasswordResult) SetPinFileNil(b bool)`

 SetPinFileNil sets the value for PinFile to be an explicit nil

### UnsetPinFile
`func (o *JellyfinForgotPasswordResult) UnsetPinFile()`

UnsetPinFile ensures that no value is present for PinFile, not even an explicit nil
### GetPinExpirationDate

`func (o *JellyfinForgotPasswordResult) GetPinExpirationDate() time.Time`

GetPinExpirationDate returns the PinExpirationDate field if non-nil, zero value otherwise.

### GetPinExpirationDateOk

`func (o *JellyfinForgotPasswordResult) GetPinExpirationDateOk() (*time.Time, bool)`

GetPinExpirationDateOk returns a tuple with the PinExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinExpirationDate

`func (o *JellyfinForgotPasswordResult) SetPinExpirationDate(v time.Time)`

SetPinExpirationDate sets PinExpirationDate field to given value.

### HasPinExpirationDate

`func (o *JellyfinForgotPasswordResult) HasPinExpirationDate() bool`

HasPinExpirationDate returns a boolean if a field has been set.

### SetPinExpirationDateNil

`func (o *JellyfinForgotPasswordResult) SetPinExpirationDateNil(b bool)`

 SetPinExpirationDateNil sets the value for PinExpirationDate to be an explicit nil

### UnsetPinExpirationDate
`func (o *JellyfinForgotPasswordResult) UnsetPinExpirationDate()`

UnsetPinExpirationDate ensures that no value is present for PinExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


