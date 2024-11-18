# JellyfinBaseItemPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Id** | Pointer to **string** | Gets or sets the identifier. | [optional] 
**Role** | Pointer to **NullableString** | Gets or sets the role. | [optional] 
**Type** | Pointer to [**JellyfinJellyfinPersonKind**](JellyfinPersonKind.md) | The person kind. | [optional] 
**PrimaryImageTag** | Pointer to **NullableString** | Gets or sets the primary image tag. | [optional] 
**ImageBlurHashes** | Pointer to [**NullableJellyfinJellyfinBaseItemPersonImageBlurHashes**](JellyfinBaseItemPersonImageBlurHashes.md) |  | [optional] 

## Methods

### NewJellyfinBaseItemPerson

`func NewJellyfinBaseItemPerson() *JellyfinBaseItemPerson`

NewJellyfinBaseItemPerson instantiates a new JellyfinBaseItemPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinBaseItemPersonWithDefaults

`func NewJellyfinBaseItemPersonWithDefaults() *JellyfinBaseItemPerson`

NewJellyfinBaseItemPersonWithDefaults instantiates a new JellyfinBaseItemPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinBaseItemPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinBaseItemPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinBaseItemPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinBaseItemPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinBaseItemPerson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinBaseItemPerson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *JellyfinBaseItemPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinBaseItemPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinBaseItemPerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinBaseItemPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *JellyfinBaseItemPerson) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *JellyfinBaseItemPerson) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *JellyfinBaseItemPerson) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *JellyfinBaseItemPerson) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *JellyfinBaseItemPerson) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *JellyfinBaseItemPerson) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetType

`func (o *JellyfinBaseItemPerson) GetType() JellyfinJellyfinPersonKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinBaseItemPerson) GetTypeOk() (*JellyfinJellyfinPersonKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinBaseItemPerson) SetType(v JellyfinJellyfinPersonKind)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinBaseItemPerson) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *JellyfinBaseItemPerson) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *JellyfinBaseItemPerson) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *JellyfinBaseItemPerson) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *JellyfinBaseItemPerson) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *JellyfinBaseItemPerson) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *JellyfinBaseItemPerson) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetImageBlurHashes

`func (o *JellyfinBaseItemPerson) GetImageBlurHashes() JellyfinJellyfinBaseItemPersonImageBlurHashes`

GetImageBlurHashes returns the ImageBlurHashes field if non-nil, zero value otherwise.

### GetImageBlurHashesOk

`func (o *JellyfinBaseItemPerson) GetImageBlurHashesOk() (*JellyfinJellyfinBaseItemPersonImageBlurHashes, bool)`

GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlurHashes

`func (o *JellyfinBaseItemPerson) SetImageBlurHashes(v JellyfinJellyfinBaseItemPersonImageBlurHashes)`

SetImageBlurHashes sets ImageBlurHashes field to given value.

### HasImageBlurHashes

`func (o *JellyfinBaseItemPerson) HasImageBlurHashes() bool`

HasImageBlurHashes returns a boolean if a field has been set.

### SetImageBlurHashesNil

`func (o *JellyfinBaseItemPerson) SetImageBlurHashesNil(b bool)`

 SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil

### UnsetImageBlurHashes
`func (o *JellyfinBaseItemPerson) UnsetImageBlurHashes()`

UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


