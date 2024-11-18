# JellyfinPlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenAccess** | Pointer to **bool** | Gets or sets a value indicating whether the playlist is publicly readable. | [optional] 
**Shares** | Pointer to [**[]JellyfinJellyfinPlaylistUserPermissions**](JellyfinJellyfinPlaylistUserPermissions.md) | Gets or sets the share permissions. | [optional] 
**ItemIds** | Pointer to **[]string** | Gets or sets the item ids. | [optional] 

## Methods

### NewJellyfinPlaylistDto

`func NewJellyfinPlaylistDto() *JellyfinPlaylistDto`

NewJellyfinPlaylistDto instantiates a new JellyfinPlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlaylistDtoWithDefaults

`func NewJellyfinPlaylistDtoWithDefaults() *JellyfinPlaylistDto`

NewJellyfinPlaylistDtoWithDefaults instantiates a new JellyfinPlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenAccess

`func (o *JellyfinPlaylistDto) GetOpenAccess() bool`

GetOpenAccess returns the OpenAccess field if non-nil, zero value otherwise.

### GetOpenAccessOk

`func (o *JellyfinPlaylistDto) GetOpenAccessOk() (*bool, bool)`

GetOpenAccessOk returns a tuple with the OpenAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccess

`func (o *JellyfinPlaylistDto) SetOpenAccess(v bool)`

SetOpenAccess sets OpenAccess field to given value.

### HasOpenAccess

`func (o *JellyfinPlaylistDto) HasOpenAccess() bool`

HasOpenAccess returns a boolean if a field has been set.

### GetShares

`func (o *JellyfinPlaylistDto) GetShares() []JellyfinJellyfinPlaylistUserPermissions`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *JellyfinPlaylistDto) GetSharesOk() (*[]JellyfinJellyfinPlaylistUserPermissions, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *JellyfinPlaylistDto) SetShares(v []JellyfinJellyfinPlaylistUserPermissions)`

SetShares sets Shares field to given value.

### HasShares

`func (o *JellyfinPlaylistDto) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetItemIds

`func (o *JellyfinPlaylistDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *JellyfinPlaylistDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *JellyfinPlaylistDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *JellyfinPlaylistDto) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


