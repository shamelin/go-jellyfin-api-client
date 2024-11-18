# ModelBaseItemPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Id** | Pointer to **string** | Gets or sets the identifier. | [optional] 
**Role** | Pointer to **NullableString** | Gets or sets the role. | [optional] 
**Type** | Pointer to [**ModelModelPersonKind**](ModelPersonKind.md) | The person kind. | [optional] 
**PrimaryImageTag** | Pointer to **NullableString** | Gets or sets the primary image tag. | [optional] 
**ImageBlurHashes** | Pointer to [**NullableModelModelBaseItemPersonImageBlurHashes**](ModelBaseItemPersonImageBlurHashes.md) |  | [optional] 

## Methods

### NewModelBaseItemPerson

`func NewModelBaseItemPerson() *ModelBaseItemPerson`

NewModelBaseItemPerson instantiates a new ModelBaseItemPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBaseItemPersonWithDefaults

`func NewModelBaseItemPersonWithDefaults() *ModelBaseItemPerson`

NewModelBaseItemPersonWithDefaults instantiates a new ModelBaseItemPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelBaseItemPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelBaseItemPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelBaseItemPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelBaseItemPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelBaseItemPerson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelBaseItemPerson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ModelBaseItemPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelBaseItemPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelBaseItemPerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelBaseItemPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *ModelBaseItemPerson) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelBaseItemPerson) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelBaseItemPerson) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModelBaseItemPerson) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ModelBaseItemPerson) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ModelBaseItemPerson) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetType

`func (o *ModelBaseItemPerson) GetType() ModelModelPersonKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelBaseItemPerson) GetTypeOk() (*ModelModelPersonKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelBaseItemPerson) SetType(v ModelModelPersonKind)`

SetType sets Type field to given value.

### HasType

`func (o *ModelBaseItemPerson) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *ModelBaseItemPerson) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *ModelBaseItemPerson) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *ModelBaseItemPerson) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *ModelBaseItemPerson) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *ModelBaseItemPerson) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *ModelBaseItemPerson) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetImageBlurHashes

`func (o *ModelBaseItemPerson) GetImageBlurHashes() ModelModelBaseItemPersonImageBlurHashes`

GetImageBlurHashes returns the ImageBlurHashes field if non-nil, zero value otherwise.

### GetImageBlurHashesOk

`func (o *ModelBaseItemPerson) GetImageBlurHashesOk() (*ModelModelBaseItemPersonImageBlurHashes, bool)`

GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlurHashes

`func (o *ModelBaseItemPerson) SetImageBlurHashes(v ModelModelBaseItemPersonImageBlurHashes)`

SetImageBlurHashes sets ImageBlurHashes field to given value.

### HasImageBlurHashes

`func (o *ModelBaseItemPerson) HasImageBlurHashes() bool`

HasImageBlurHashes returns a boolean if a field has been set.

### SetImageBlurHashesNil

`func (o *ModelBaseItemPerson) SetImageBlurHashesNil(b bool)`

 SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil

### UnsetImageBlurHashes
`func (o *ModelBaseItemPerson) UnsetImageBlurHashes()`

UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


