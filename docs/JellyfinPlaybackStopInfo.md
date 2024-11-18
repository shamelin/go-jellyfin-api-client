# JellyfinPlaybackStopInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**NullableJellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md) | Gets or sets the item. | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the item identifier. | [optional] 
**SessionId** | Pointer to **NullableString** | Gets or sets the session id. | [optional] 
**MediaSourceId** | Pointer to **NullableString** | Gets or sets the media version identifier. | [optional] 
**PositionTicks** | Pointer to **NullableInt64** | Gets or sets the position ticks. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the live stream identifier. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**Failed** | Pointer to **bool** | Gets or sets a value indicating whether this MediaBrowser.Model.Session.PlaybackStopInfo is failed. | [optional] 
**NextMediaType** | Pointer to **NullableString** |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 
**NowPlayingQueue** | Pointer to [**[]JellyfinJellyfinQueueItem**](JellyfinJellyfinQueueItem.md) |  | [optional] 

## Methods

### NewJellyfinPlaybackStopInfo

`func NewJellyfinPlaybackStopInfo() *JellyfinPlaybackStopInfo`

NewJellyfinPlaybackStopInfo instantiates a new JellyfinPlaybackStopInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlaybackStopInfoWithDefaults

`func NewJellyfinPlaybackStopInfoWithDefaults() *JellyfinPlaybackStopInfo`

NewJellyfinPlaybackStopInfoWithDefaults instantiates a new JellyfinPlaybackStopInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *JellyfinPlaybackStopInfo) GetItem() JellyfinJellyfinBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *JellyfinPlaybackStopInfo) GetItemOk() (*JellyfinJellyfinBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *JellyfinPlaybackStopInfo) SetItem(v JellyfinJellyfinBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *JellyfinPlaybackStopInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *JellyfinPlaybackStopInfo) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *JellyfinPlaybackStopInfo) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetItemId

`func (o *JellyfinPlaybackStopInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinPlaybackStopInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinPlaybackStopInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinPlaybackStopInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *JellyfinPlaybackStopInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JellyfinPlaybackStopInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JellyfinPlaybackStopInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JellyfinPlaybackStopInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *JellyfinPlaybackStopInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *JellyfinPlaybackStopInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMediaSourceId

`func (o *JellyfinPlaybackStopInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *JellyfinPlaybackStopInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *JellyfinPlaybackStopInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *JellyfinPlaybackStopInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *JellyfinPlaybackStopInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *JellyfinPlaybackStopInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetPositionTicks

`func (o *JellyfinPlaybackStopInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *JellyfinPlaybackStopInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *JellyfinPlaybackStopInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *JellyfinPlaybackStopInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *JellyfinPlaybackStopInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *JellyfinPlaybackStopInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetLiveStreamId

`func (o *JellyfinPlaybackStopInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *JellyfinPlaybackStopInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *JellyfinPlaybackStopInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *JellyfinPlaybackStopInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *JellyfinPlaybackStopInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *JellyfinPlaybackStopInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetPlaySessionId

`func (o *JellyfinPlaybackStopInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *JellyfinPlaybackStopInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *JellyfinPlaybackStopInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *JellyfinPlaybackStopInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *JellyfinPlaybackStopInfo) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *JellyfinPlaybackStopInfo) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetFailed

`func (o *JellyfinPlaybackStopInfo) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *JellyfinPlaybackStopInfo) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *JellyfinPlaybackStopInfo) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *JellyfinPlaybackStopInfo) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetNextMediaType

`func (o *JellyfinPlaybackStopInfo) GetNextMediaType() string`

GetNextMediaType returns the NextMediaType field if non-nil, zero value otherwise.

### GetNextMediaTypeOk

`func (o *JellyfinPlaybackStopInfo) GetNextMediaTypeOk() (*string, bool)`

GetNextMediaTypeOk returns a tuple with the NextMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMediaType

`func (o *JellyfinPlaybackStopInfo) SetNextMediaType(v string)`

SetNextMediaType sets NextMediaType field to given value.

### HasNextMediaType

`func (o *JellyfinPlaybackStopInfo) HasNextMediaType() bool`

HasNextMediaType returns a boolean if a field has been set.

### SetNextMediaTypeNil

`func (o *JellyfinPlaybackStopInfo) SetNextMediaTypeNil(b bool)`

 SetNextMediaTypeNil sets the value for NextMediaType to be an explicit nil

### UnsetNextMediaType
`func (o *JellyfinPlaybackStopInfo) UnsetNextMediaType()`

UnsetNextMediaType ensures that no value is present for NextMediaType, not even an explicit nil
### GetPlaylistItemId

`func (o *JellyfinPlaybackStopInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *JellyfinPlaybackStopInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *JellyfinPlaybackStopInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *JellyfinPlaybackStopInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *JellyfinPlaybackStopInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *JellyfinPlaybackStopInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetNowPlayingQueue

`func (o *JellyfinPlaybackStopInfo) GetNowPlayingQueue() []JellyfinJellyfinQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *JellyfinPlaybackStopInfo) GetNowPlayingQueueOk() (*[]JellyfinJellyfinQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *JellyfinPlaybackStopInfo) SetNowPlayingQueue(v []JellyfinJellyfinQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *JellyfinPlaybackStopInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *JellyfinPlaybackStopInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *JellyfinPlaybackStopInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


