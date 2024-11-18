# ModelPlaybackProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSeek** | Pointer to **bool** | Gets or sets a value indicating whether this instance can seek. | [optional] 
**Item** | Pointer to [**NullableModelModelBaseItemDto**](ModelBaseItemDto.md) | Gets or sets the item. | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the item identifier. | [optional] 
**SessionId** | Pointer to **NullableString** | Gets or sets the session id. | [optional] 
**MediaSourceId** | Pointer to **NullableString** | Gets or sets the media version identifier. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the audio stream. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the subtitle stream. | [optional] 
**IsPaused** | Pointer to **bool** | Gets or sets a value indicating whether this instance is paused. | [optional] 
**IsMuted** | Pointer to **bool** | Gets or sets a value indicating whether this instance is muted. | [optional] 
**PositionTicks** | Pointer to **NullableInt64** | Gets or sets the position ticks. | [optional] 
**PlaybackStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** | Gets or sets the volume level. | [optional] 
**Brightness** | Pointer to **NullableInt32** |  | [optional] 
**AspectRatio** | Pointer to **NullableString** |  | [optional] 
**PlayMethod** | Pointer to [**ModelModelPlayMethod**](ModelPlayMethod.md) | Gets or sets the play method. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the live stream identifier. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**RepeatMode** | Pointer to [**ModelModelRepeatMode**](ModelRepeatMode.md) | Gets or sets the repeat mode. | [optional] 
**PlaybackOrder** | Pointer to [**ModelModelPlaybackOrder**](ModelPlaybackOrder.md) | Gets or sets the playback order. | [optional] 
**NowPlayingQueue** | Pointer to [**[]ModelModelQueueItem**](ModelModelQueueItem.md) |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelPlaybackProgressInfo

`func NewModelPlaybackProgressInfo() *ModelPlaybackProgressInfo`

NewModelPlaybackProgressInfo instantiates a new ModelPlaybackProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaybackProgressInfoWithDefaults

`func NewModelPlaybackProgressInfoWithDefaults() *ModelPlaybackProgressInfo`

NewModelPlaybackProgressInfoWithDefaults instantiates a new ModelPlaybackProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSeek

`func (o *ModelPlaybackProgressInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *ModelPlaybackProgressInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *ModelPlaybackProgressInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *ModelPlaybackProgressInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetItem

`func (o *ModelPlaybackProgressInfo) GetItem() ModelModelBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ModelPlaybackProgressInfo) GetItemOk() (*ModelModelBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ModelPlaybackProgressInfo) SetItem(v ModelModelBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *ModelPlaybackProgressInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *ModelPlaybackProgressInfo) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *ModelPlaybackProgressInfo) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetItemId

`func (o *ModelPlaybackProgressInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelPlaybackProgressInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelPlaybackProgressInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelPlaybackProgressInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *ModelPlaybackProgressInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ModelPlaybackProgressInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ModelPlaybackProgressInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ModelPlaybackProgressInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *ModelPlaybackProgressInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *ModelPlaybackProgressInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMediaSourceId

`func (o *ModelPlaybackProgressInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlaybackProgressInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlaybackProgressInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlaybackProgressInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *ModelPlaybackProgressInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *ModelPlaybackProgressInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetAudioStreamIndex

`func (o *ModelPlaybackProgressInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelPlaybackProgressInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelPlaybackProgressInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelPlaybackProgressInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelPlaybackProgressInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelPlaybackProgressInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *ModelPlaybackProgressInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelPlaybackProgressInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelPlaybackProgressInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelPlaybackProgressInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelPlaybackProgressInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelPlaybackProgressInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetIsPaused

`func (o *ModelPlaybackProgressInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ModelPlaybackProgressInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ModelPlaybackProgressInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ModelPlaybackProgressInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *ModelPlaybackProgressInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *ModelPlaybackProgressInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *ModelPlaybackProgressInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *ModelPlaybackProgressInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetPositionTicks

`func (o *ModelPlaybackProgressInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelPlaybackProgressInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelPlaybackProgressInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelPlaybackProgressInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelPlaybackProgressInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelPlaybackProgressInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlaybackStartTimeTicks

`func (o *ModelPlaybackProgressInfo) GetPlaybackStartTimeTicks() int64`

GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field if non-nil, zero value otherwise.

### GetPlaybackStartTimeTicksOk

`func (o *ModelPlaybackProgressInfo) GetPlaybackStartTimeTicksOk() (*int64, bool)`

GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackStartTimeTicks

`func (o *ModelPlaybackProgressInfo) SetPlaybackStartTimeTicks(v int64)`

SetPlaybackStartTimeTicks sets PlaybackStartTimeTicks field to given value.

### HasPlaybackStartTimeTicks

`func (o *ModelPlaybackProgressInfo) HasPlaybackStartTimeTicks() bool`

HasPlaybackStartTimeTicks returns a boolean if a field has been set.

### SetPlaybackStartTimeTicksNil

`func (o *ModelPlaybackProgressInfo) SetPlaybackStartTimeTicksNil(b bool)`

 SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil

### UnsetPlaybackStartTimeTicks
`func (o *ModelPlaybackProgressInfo) UnsetPlaybackStartTimeTicks()`

UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
### GetVolumeLevel

`func (o *ModelPlaybackProgressInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *ModelPlaybackProgressInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *ModelPlaybackProgressInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *ModelPlaybackProgressInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *ModelPlaybackProgressInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *ModelPlaybackProgressInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetBrightness

`func (o *ModelPlaybackProgressInfo) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *ModelPlaybackProgressInfo) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *ModelPlaybackProgressInfo) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *ModelPlaybackProgressInfo) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### SetBrightnessNil

`func (o *ModelPlaybackProgressInfo) SetBrightnessNil(b bool)`

 SetBrightnessNil sets the value for Brightness to be an explicit nil

### UnsetBrightness
`func (o *ModelPlaybackProgressInfo) UnsetBrightness()`

UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
### GetAspectRatio

`func (o *ModelPlaybackProgressInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ModelPlaybackProgressInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ModelPlaybackProgressInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *ModelPlaybackProgressInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *ModelPlaybackProgressInfo) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *ModelPlaybackProgressInfo) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetPlayMethod

`func (o *ModelPlaybackProgressInfo) GetPlayMethod() ModelModelPlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *ModelPlaybackProgressInfo) GetPlayMethodOk() (*ModelModelPlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *ModelPlaybackProgressInfo) SetPlayMethod(v ModelModelPlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *ModelPlaybackProgressInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *ModelPlaybackProgressInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelPlaybackProgressInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelPlaybackProgressInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelPlaybackProgressInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *ModelPlaybackProgressInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *ModelPlaybackProgressInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetPlaySessionId

`func (o *ModelPlaybackProgressInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *ModelPlaybackProgressInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *ModelPlaybackProgressInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *ModelPlaybackProgressInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *ModelPlaybackProgressInfo) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *ModelPlaybackProgressInfo) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetRepeatMode

`func (o *ModelPlaybackProgressInfo) GetRepeatMode() ModelModelRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *ModelPlaybackProgressInfo) GetRepeatModeOk() (*ModelModelRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *ModelPlaybackProgressInfo) SetRepeatMode(v ModelModelRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *ModelPlaybackProgressInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetPlaybackOrder

`func (o *ModelPlaybackProgressInfo) GetPlaybackOrder() ModelModelPlaybackOrder`

GetPlaybackOrder returns the PlaybackOrder field if non-nil, zero value otherwise.

### GetPlaybackOrderOk

`func (o *ModelPlaybackProgressInfo) GetPlaybackOrderOk() (*ModelModelPlaybackOrder, bool)`

GetPlaybackOrderOk returns a tuple with the PlaybackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackOrder

`func (o *ModelPlaybackProgressInfo) SetPlaybackOrder(v ModelModelPlaybackOrder)`

SetPlaybackOrder sets PlaybackOrder field to given value.

### HasPlaybackOrder

`func (o *ModelPlaybackProgressInfo) HasPlaybackOrder() bool`

HasPlaybackOrder returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *ModelPlaybackProgressInfo) GetNowPlayingQueue() []ModelModelQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *ModelPlaybackProgressInfo) GetNowPlayingQueueOk() (*[]ModelModelQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *ModelPlaybackProgressInfo) SetNowPlayingQueue(v []ModelModelQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *ModelPlaybackProgressInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *ModelPlaybackProgressInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *ModelPlaybackProgressInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetPlaylistItemId

`func (o *ModelPlaybackProgressInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelPlaybackProgressInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelPlaybackProgressInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelPlaybackProgressInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *ModelPlaybackProgressInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *ModelPlaybackProgressInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


