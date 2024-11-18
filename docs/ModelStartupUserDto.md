# ModelStartupUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Password** | Pointer to **NullableString** | Gets or sets the user&#39;s password. | [optional] 

## Methods

### NewModelStartupUserDto

`func NewModelStartupUserDto() *ModelStartupUserDto`

NewModelStartupUserDto instantiates a new ModelStartupUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelStartupUserDtoWithDefaults

`func NewModelStartupUserDtoWithDefaults() *ModelStartupUserDto`

NewModelStartupUserDtoWithDefaults instantiates a new ModelStartupUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelStartupUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelStartupUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelStartupUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelStartupUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelStartupUserDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelStartupUserDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPassword

`func (o *ModelStartupUserDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelStartupUserDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelStartupUserDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelStartupUserDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ModelStartupUserDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ModelStartupUserDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


