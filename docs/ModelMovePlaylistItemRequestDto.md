# ModelMovePlaylistItemRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist identifier of the item. | [optional] 
**NewIndex** | Pointer to **int32** | Gets or sets the new position. | [optional] 

## Methods

### NewModelMovePlaylistItemRequestDto

`func NewModelMovePlaylistItemRequestDto() *ModelMovePlaylistItemRequestDto`

NewModelMovePlaylistItemRequestDto instantiates a new ModelMovePlaylistItemRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMovePlaylistItemRequestDtoWithDefaults

`func NewModelMovePlaylistItemRequestDtoWithDefaults() *ModelMovePlaylistItemRequestDto`

NewModelMovePlaylistItemRequestDtoWithDefaults instantiates a new ModelMovePlaylistItemRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistItemId

`func (o *ModelMovePlaylistItemRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelMovePlaylistItemRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelMovePlaylistItemRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelMovePlaylistItemRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetNewIndex

`func (o *ModelMovePlaylistItemRequestDto) GetNewIndex() int32`

GetNewIndex returns the NewIndex field if non-nil, zero value otherwise.

### GetNewIndexOk

`func (o *ModelMovePlaylistItemRequestDto) GetNewIndexOk() (*int32, bool)`

GetNewIndexOk returns a tuple with the NewIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIndex

`func (o *ModelMovePlaylistItemRequestDto) SetNewIndex(v int32)`

SetNewIndex sets NewIndex field to given value.

### HasNewIndex

`func (o *ModelMovePlaylistItemRequestDto) HasNewIndex() bool`

HasNewIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


