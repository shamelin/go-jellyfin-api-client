# JellyfinPlaybackStartInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSeek** | Pointer to **bool** | Gets or sets a value indicating whether this instance can seek. | [optional] 
**Item** | Pointer to [**NullableJellyfinJellyfinBaseItemDto**](JellyfinBaseItemDto.md) | Gets or sets the item. | [optional] 
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
**PlayMethod** | Pointer to [**JellyfinJellyfinPlayMethod**](JellyfinPlayMethod.md) | Gets or sets the play method. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the live stream identifier. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**RepeatMode** | Pointer to [**JellyfinJellyfinRepeatMode**](JellyfinRepeatMode.md) | Gets or sets the repeat mode. | [optional] 
**PlaybackOrder** | Pointer to [**JellyfinJellyfinPlaybackOrder**](JellyfinPlaybackOrder.md) | Gets or sets the playback order. | [optional] 
**NowPlayingQueue** | Pointer to [**[]JellyfinJellyfinQueueItem**](JellyfinJellyfinQueueItem.md) |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJellyfinPlaybackStartInfo

`func NewJellyfinPlaybackStartInfo() *JellyfinPlaybackStartInfo`

NewJellyfinPlaybackStartInfo instantiates a new JellyfinPlaybackStartInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlaybackStartInfoWithDefaults

`func NewJellyfinPlaybackStartInfoWithDefaults() *JellyfinPlaybackStartInfo`

NewJellyfinPlaybackStartInfoWithDefaults instantiates a new JellyfinPlaybackStartInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSeek

`func (o *JellyfinPlaybackStartInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *JellyfinPlaybackStartInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *JellyfinPlaybackStartInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *JellyfinPlaybackStartInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetItem

`func (o *JellyfinPlaybackStartInfo) GetItem() JellyfinJellyfinBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *JellyfinPlaybackStartInfo) GetItemOk() (*JellyfinJellyfinBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *JellyfinPlaybackStartInfo) SetItem(v JellyfinJellyfinBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *JellyfinPlaybackStartInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *JellyfinPlaybackStartInfo) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *JellyfinPlaybackStartInfo) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetItemId

`func (o *JellyfinPlaybackStartInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinPlaybackStartInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinPlaybackStartInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinPlaybackStartInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *JellyfinPlaybackStartInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JellyfinPlaybackStartInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JellyfinPlaybackStartInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JellyfinPlaybackStartInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *JellyfinPlaybackStartInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *JellyfinPlaybackStartInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMediaSourceId

`func (o *JellyfinPlaybackStartInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *JellyfinPlaybackStartInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *JellyfinPlaybackStartInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *JellyfinPlaybackStartInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *JellyfinPlaybackStartInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *JellyfinPlaybackStartInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetAudioStreamIndex

`func (o *JellyfinPlaybackStartInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *JellyfinPlaybackStartInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *JellyfinPlaybackStartInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *JellyfinPlaybackStartInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *JellyfinPlaybackStartInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *JellyfinPlaybackStartInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *JellyfinPlaybackStartInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *JellyfinPlaybackStartInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *JellyfinPlaybackStartInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *JellyfinPlaybackStartInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *JellyfinPlaybackStartInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *JellyfinPlaybackStartInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetIsPaused

`func (o *JellyfinPlaybackStartInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *JellyfinPlaybackStartInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *JellyfinPlaybackStartInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *JellyfinPlaybackStartInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *JellyfinPlaybackStartInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *JellyfinPlaybackStartInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *JellyfinPlaybackStartInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *JellyfinPlaybackStartInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetPositionTicks

`func (o *JellyfinPlaybackStartInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *JellyfinPlaybackStartInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *JellyfinPlaybackStartInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *JellyfinPlaybackStartInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *JellyfinPlaybackStartInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *JellyfinPlaybackStartInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlaybackStartTimeTicks

`func (o *JellyfinPlaybackStartInfo) GetPlaybackStartTimeTicks() int64`

GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field if non-nil, zero value otherwise.

### GetPlaybackStartTimeTicksOk

`func (o *JellyfinPlaybackStartInfo) GetPlaybackStartTimeTicksOk() (*int64, bool)`

GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackStartTimeTicks

`func (o *JellyfinPlaybackStartInfo) SetPlaybackStartTimeTicks(v int64)`

SetPlaybackStartTimeTicks sets PlaybackStartTimeTicks field to given value.

### HasPlaybackStartTimeTicks

`func (o *JellyfinPlaybackStartInfo) HasPlaybackStartTimeTicks() bool`

HasPlaybackStartTimeTicks returns a boolean if a field has been set.

### SetPlaybackStartTimeTicksNil

`func (o *JellyfinPlaybackStartInfo) SetPlaybackStartTimeTicksNil(b bool)`

 SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil

### UnsetPlaybackStartTimeTicks
`func (o *JellyfinPlaybackStartInfo) UnsetPlaybackStartTimeTicks()`

UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
### GetVolumeLevel

`func (o *JellyfinPlaybackStartInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *JellyfinPlaybackStartInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *JellyfinPlaybackStartInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *JellyfinPlaybackStartInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *JellyfinPlaybackStartInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *JellyfinPlaybackStartInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetBrightness

`func (o *JellyfinPlaybackStartInfo) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *JellyfinPlaybackStartInfo) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *JellyfinPlaybackStartInfo) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *JellyfinPlaybackStartInfo) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### SetBrightnessNil

`func (o *JellyfinPlaybackStartInfo) SetBrightnessNil(b bool)`

 SetBrightnessNil sets the value for Brightness to be an explicit nil

### UnsetBrightness
`func (o *JellyfinPlaybackStartInfo) UnsetBrightness()`

UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
### GetAspectRatio

`func (o *JellyfinPlaybackStartInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *JellyfinPlaybackStartInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *JellyfinPlaybackStartInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *JellyfinPlaybackStartInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *JellyfinPlaybackStartInfo) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *JellyfinPlaybackStartInfo) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetPlayMethod

`func (o *JellyfinPlaybackStartInfo) GetPlayMethod() JellyfinJellyfinPlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *JellyfinPlaybackStartInfo) GetPlayMethodOk() (*JellyfinJellyfinPlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *JellyfinPlaybackStartInfo) SetPlayMethod(v JellyfinJellyfinPlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *JellyfinPlaybackStartInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *JellyfinPlaybackStartInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *JellyfinPlaybackStartInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *JellyfinPlaybackStartInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *JellyfinPlaybackStartInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *JellyfinPlaybackStartInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *JellyfinPlaybackStartInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetPlaySessionId

`func (o *JellyfinPlaybackStartInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *JellyfinPlaybackStartInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *JellyfinPlaybackStartInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *JellyfinPlaybackStartInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *JellyfinPlaybackStartInfo) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *JellyfinPlaybackStartInfo) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetRepeatMode

`func (o *JellyfinPlaybackStartInfo) GetRepeatMode() JellyfinJellyfinRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *JellyfinPlaybackStartInfo) GetRepeatModeOk() (*JellyfinJellyfinRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *JellyfinPlaybackStartInfo) SetRepeatMode(v JellyfinJellyfinRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *JellyfinPlaybackStartInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetPlaybackOrder

`func (o *JellyfinPlaybackStartInfo) GetPlaybackOrder() JellyfinJellyfinPlaybackOrder`

GetPlaybackOrder returns the PlaybackOrder field if non-nil, zero value otherwise.

### GetPlaybackOrderOk

`func (o *JellyfinPlaybackStartInfo) GetPlaybackOrderOk() (*JellyfinJellyfinPlaybackOrder, bool)`

GetPlaybackOrderOk returns a tuple with the PlaybackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackOrder

`func (o *JellyfinPlaybackStartInfo) SetPlaybackOrder(v JellyfinJellyfinPlaybackOrder)`

SetPlaybackOrder sets PlaybackOrder field to given value.

### HasPlaybackOrder

`func (o *JellyfinPlaybackStartInfo) HasPlaybackOrder() bool`

HasPlaybackOrder returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *JellyfinPlaybackStartInfo) GetNowPlayingQueue() []JellyfinJellyfinQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *JellyfinPlaybackStartInfo) GetNowPlayingQueueOk() (*[]JellyfinJellyfinQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *JellyfinPlaybackStartInfo) SetNowPlayingQueue(v []JellyfinJellyfinQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *JellyfinPlaybackStartInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *JellyfinPlaybackStartInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *JellyfinPlaybackStartInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetPlaylistItemId

`func (o *JellyfinPlaybackStartInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *JellyfinPlaybackStartInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *JellyfinPlaybackStartInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *JellyfinPlaybackStartInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *JellyfinPlaybackStartInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *JellyfinPlaybackStartInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


