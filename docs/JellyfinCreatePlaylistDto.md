# JellyfinCreatePlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name of the new playlist. | [optional] 
**Ids** | Pointer to **[]string** | Gets or sets item ids to add to the playlist. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**MediaType** | Pointer to [**NullableJellyfinJellyfinMediaType**](JellyfinMediaType.md) | Gets or sets the media type. | [optional] 
**Users** | Pointer to [**[]JellyfinJellyfinPlaylistUserPermissions**](JellyfinJellyfinPlaylistUserPermissions.md) | Gets or sets the playlist users. | [optional] 
**IsPublic** | Pointer to **bool** | Gets or sets a value indicating whether the playlist is public. | [optional] 

## Methods

### NewJellyfinCreatePlaylistDto

`func NewJellyfinCreatePlaylistDto() *JellyfinCreatePlaylistDto`

NewJellyfinCreatePlaylistDto instantiates a new JellyfinCreatePlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinCreatePlaylistDtoWithDefaults

`func NewJellyfinCreatePlaylistDtoWithDefaults() *JellyfinCreatePlaylistDto`

NewJellyfinCreatePlaylistDtoWithDefaults instantiates a new JellyfinCreatePlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinCreatePlaylistDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinCreatePlaylistDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinCreatePlaylistDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinCreatePlaylistDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIds

`func (o *JellyfinCreatePlaylistDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *JellyfinCreatePlaylistDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *JellyfinCreatePlaylistDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *JellyfinCreatePlaylistDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetUserId

`func (o *JellyfinCreatePlaylistDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinCreatePlaylistDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinCreatePlaylistDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinCreatePlaylistDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *JellyfinCreatePlaylistDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *JellyfinCreatePlaylistDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMediaType

`func (o *JellyfinCreatePlaylistDto) GetMediaType() JellyfinJellyfinMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *JellyfinCreatePlaylistDto) GetMediaTypeOk() (*JellyfinJellyfinMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *JellyfinCreatePlaylistDto) SetMediaType(v JellyfinJellyfinMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *JellyfinCreatePlaylistDto) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *JellyfinCreatePlaylistDto) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *JellyfinCreatePlaylistDto) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetUsers

`func (o *JellyfinCreatePlaylistDto) GetUsers() []JellyfinJellyfinPlaylistUserPermissions`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *JellyfinCreatePlaylistDto) GetUsersOk() (*[]JellyfinJellyfinPlaylistUserPermissions, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *JellyfinCreatePlaylistDto) SetUsers(v []JellyfinJellyfinPlaylistUserPermissions)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *JellyfinCreatePlaylistDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetIsPublic

`func (o *JellyfinCreatePlaylistDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *JellyfinCreatePlaylistDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *JellyfinCreatePlaylistDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *JellyfinCreatePlaylistDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


