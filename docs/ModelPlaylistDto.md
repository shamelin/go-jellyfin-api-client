# ModelPlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenAccess** | Pointer to **bool** | Gets or sets a value indicating whether the playlist is publicly readable. | [optional] 
**Shares** | Pointer to [**[]ModelModelPlaylistUserPermissions**](ModelModelPlaylistUserPermissions.md) | Gets or sets the share permissions. | [optional] 
**ItemIds** | Pointer to **[]string** | Gets or sets the item ids. | [optional] 

## Methods

### NewModelPlaylistDto

`func NewModelPlaylistDto() *ModelPlaylistDto`

NewModelPlaylistDto instantiates a new ModelPlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaylistDtoWithDefaults

`func NewModelPlaylistDtoWithDefaults() *ModelPlaylistDto`

NewModelPlaylistDtoWithDefaults instantiates a new ModelPlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenAccess

`func (o *ModelPlaylistDto) GetOpenAccess() bool`

GetOpenAccess returns the OpenAccess field if non-nil, zero value otherwise.

### GetOpenAccessOk

`func (o *ModelPlaylistDto) GetOpenAccessOk() (*bool, bool)`

GetOpenAccessOk returns a tuple with the OpenAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccess

`func (o *ModelPlaylistDto) SetOpenAccess(v bool)`

SetOpenAccess sets OpenAccess field to given value.

### HasOpenAccess

`func (o *ModelPlaylistDto) HasOpenAccess() bool`

HasOpenAccess returns a boolean if a field has been set.

### GetShares

`func (o *ModelPlaylistDto) GetShares() []ModelModelPlaylistUserPermissions`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *ModelPlaylistDto) GetSharesOk() (*[]ModelModelPlaylistUserPermissions, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *ModelPlaylistDto) SetShares(v []ModelModelPlaylistUserPermissions)`

SetShares sets Shares field to given value.

### HasShares

`func (o *ModelPlaylistDto) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetItemIds

`func (o *ModelPlaylistDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ModelPlaylistDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ModelPlaylistDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ModelPlaylistDto) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


