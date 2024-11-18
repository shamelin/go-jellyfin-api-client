# JellyfinReadyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | Pointer to **time.Time** | Gets or sets when the request has been made by the client. | [optional] 
**PositionTicks** | Pointer to **int64** | Gets or sets the position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets or sets a value indicating whether the client playback is unpaused. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist item identifier of the playing item. | [optional] 

## Methods

### NewJellyfinReadyRequestDto

`func NewJellyfinReadyRequestDto() *JellyfinReadyRequestDto`

NewJellyfinReadyRequestDto instantiates a new JellyfinReadyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinReadyRequestDtoWithDefaults

`func NewJellyfinReadyRequestDtoWithDefaults() *JellyfinReadyRequestDto`

NewJellyfinReadyRequestDtoWithDefaults instantiates a new JellyfinReadyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *JellyfinReadyRequestDto) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *JellyfinReadyRequestDto) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *JellyfinReadyRequestDto) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *JellyfinReadyRequestDto) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPositionTicks

`func (o *JellyfinReadyRequestDto) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *JellyfinReadyRequestDto) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *JellyfinReadyRequestDto) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *JellyfinReadyRequestDto) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *JellyfinReadyRequestDto) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *JellyfinReadyRequestDto) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *JellyfinReadyRequestDto) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *JellyfinReadyRequestDto) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *JellyfinReadyRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *JellyfinReadyRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *JellyfinReadyRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *JellyfinReadyRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


