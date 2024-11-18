# ModelCreateUserByName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the username. | 
**Password** | Pointer to **NullableString** | Gets or sets the password. | [optional] 

## Methods

### NewModelCreateUserByName

`func NewModelCreateUserByName(name string, ) *ModelCreateUserByName`

NewModelCreateUserByName instantiates a new ModelCreateUserByName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreateUserByNameWithDefaults

`func NewModelCreateUserByNameWithDefaults() *ModelCreateUserByName`

NewModelCreateUserByNameWithDefaults instantiates a new ModelCreateUserByName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelCreateUserByName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCreateUserByName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCreateUserByName) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *ModelCreateUserByName) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelCreateUserByName) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelCreateUserByName) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelCreateUserByName) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ModelCreateUserByName) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ModelCreateUserByName) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


