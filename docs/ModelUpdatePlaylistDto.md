# ModelUpdatePlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the new playlist. | [optional] 
**Ids** | Pointer to **[]string** | Gets or sets item ids of the playlist. | [optional] 
**Users** | Pointer to [**[]ModelModelPlaylistUserPermissions**](ModelModelPlaylistUserPermissions.md) | Gets or sets the playlist users. | [optional] 
**IsPublic** | Pointer to **NullableBool** | Gets or sets a value indicating whether the playlist is public. | [optional] 

## Methods

### NewModelUpdatePlaylistDto

`func NewModelUpdatePlaylistDto() *ModelUpdatePlaylistDto`

NewModelUpdatePlaylistDto instantiates a new ModelUpdatePlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpdatePlaylistDtoWithDefaults

`func NewModelUpdatePlaylistDtoWithDefaults() *ModelUpdatePlaylistDto`

NewModelUpdatePlaylistDtoWithDefaults instantiates a new ModelUpdatePlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelUpdatePlaylistDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelUpdatePlaylistDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelUpdatePlaylistDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelUpdatePlaylistDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelUpdatePlaylistDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelUpdatePlaylistDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *ModelUpdatePlaylistDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ModelUpdatePlaylistDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ModelUpdatePlaylistDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ModelUpdatePlaylistDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *ModelUpdatePlaylistDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *ModelUpdatePlaylistDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetUsers

`func (o *ModelUpdatePlaylistDto) GetUsers() []ModelModelPlaylistUserPermissions`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModelUpdatePlaylistDto) GetUsersOk() (*[]ModelModelPlaylistUserPermissions, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModelUpdatePlaylistDto) SetUsers(v []ModelModelPlaylistUserPermissions)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ModelUpdatePlaylistDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ModelUpdatePlaylistDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ModelUpdatePlaylistDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetIsPublic

`func (o *ModelUpdatePlaylistDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ModelUpdatePlaylistDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ModelUpdatePlaylistDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ModelUpdatePlaylistDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *ModelUpdatePlaylistDto) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *ModelUpdatePlaylistDto) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


