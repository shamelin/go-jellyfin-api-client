# JellyfinBufferRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | Pointer to **time.Time** | Gets or sets when the request has been made by the client. | [optional] 
**PositionTicks** | Pointer to **int64** | Gets or sets the position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets or sets a value indicating whether the client playback is unpaused. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist item identifier of the playing item. | [optional] 

## Methods

### NewJellyfinBufferRequestDto

`func NewJellyfinBufferRequestDto() *JellyfinBufferRequestDto`

NewJellyfinBufferRequestDto instantiates a new JellyfinBufferRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinBufferRequestDtoWithDefaults

`func NewJellyfinBufferRequestDtoWithDefaults() *JellyfinBufferRequestDto`

NewJellyfinBufferRequestDtoWithDefaults instantiates a new JellyfinBufferRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *JellyfinBufferRequestDto) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *JellyfinBufferRequestDto) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *JellyfinBufferRequestDto) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *JellyfinBufferRequestDto) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPositionTicks

`func (o *JellyfinBufferRequestDto) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *JellyfinBufferRequestDto) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *JellyfinBufferRequestDto) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *JellyfinBufferRequestDto) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *JellyfinBufferRequestDto) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *JellyfinBufferRequestDto) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *JellyfinBufferRequestDto) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *JellyfinBufferRequestDto) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *JellyfinBufferRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *JellyfinBufferRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *JellyfinBufferRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *JellyfinBufferRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


