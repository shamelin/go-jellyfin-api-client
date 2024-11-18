# JellyfinPlayerStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionTicks** | Pointer to **NullableInt64** | Gets or sets the now playing position ticks. | [optional] 
**CanSeek** | Pointer to **bool** | Gets or sets a value indicating whether this instance can seek. | [optional] 
**IsPaused** | Pointer to **bool** | Gets or sets a value indicating whether this instance is paused. | [optional] 
**IsMuted** | Pointer to **bool** | Gets or sets a value indicating whether this instance is muted. | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** | Gets or sets the volume level. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the now playing audio stream. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the now playing subtitle stream. | [optional] 
**MediaSourceId** | Pointer to **NullableString** | Gets or sets the now playing media version identifier. | [optional] 
**PlayMethod** | Pointer to [**NullableJellyfinJellyfinPlayMethod**](JellyfinPlayMethod.md) | Gets or sets the play method. | [optional] 
**RepeatMode** | Pointer to [**JellyfinJellyfinRepeatMode**](JellyfinRepeatMode.md) | Gets or sets the repeat mode. | [optional] 
**PlaybackOrder** | Pointer to [**JellyfinJellyfinPlaybackOrder**](JellyfinPlaybackOrder.md) | Gets or sets the playback order. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the now playing live stream identifier. | [optional] 

## Methods

### NewJellyfinPlayerStateInfo

`func NewJellyfinPlayerStateInfo() *JellyfinPlayerStateInfo`

NewJellyfinPlayerStateInfo instantiates a new JellyfinPlayerStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlayerStateInfoWithDefaults

`func NewJellyfinPlayerStateInfoWithDefaults() *JellyfinPlayerStateInfo`

NewJellyfinPlayerStateInfoWithDefaults instantiates a new JellyfinPlayerStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionTicks

`func (o *JellyfinPlayerStateInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *JellyfinPlayerStateInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *JellyfinPlayerStateInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *JellyfinPlayerStateInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *JellyfinPlayerStateInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *JellyfinPlayerStateInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetCanSeek

`func (o *JellyfinPlayerStateInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *JellyfinPlayerStateInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *JellyfinPlayerStateInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *JellyfinPlayerStateInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetIsPaused

`func (o *JellyfinPlayerStateInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *JellyfinPlayerStateInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *JellyfinPlayerStateInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *JellyfinPlayerStateInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *JellyfinPlayerStateInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *JellyfinPlayerStateInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *JellyfinPlayerStateInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *JellyfinPlayerStateInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetVolumeLevel

`func (o *JellyfinPlayerStateInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *JellyfinPlayerStateInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *JellyfinPlayerStateInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *JellyfinPlayerStateInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *JellyfinPlayerStateInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *JellyfinPlayerStateInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetAudioStreamIndex

`func (o *JellyfinPlayerStateInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *JellyfinPlayerStateInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *JellyfinPlayerStateInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *JellyfinPlayerStateInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *JellyfinPlayerStateInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *JellyfinPlayerStateInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *JellyfinPlayerStateInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *JellyfinPlayerStateInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *JellyfinPlayerStateInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *JellyfinPlayerStateInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *JellyfinPlayerStateInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *JellyfinPlayerStateInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMediaSourceId

`func (o *JellyfinPlayerStateInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *JellyfinPlayerStateInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *JellyfinPlayerStateInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *JellyfinPlayerStateInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *JellyfinPlayerStateInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *JellyfinPlayerStateInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetPlayMethod

`func (o *JellyfinPlayerStateInfo) GetPlayMethod() JellyfinJellyfinPlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *JellyfinPlayerStateInfo) GetPlayMethodOk() (*JellyfinJellyfinPlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *JellyfinPlayerStateInfo) SetPlayMethod(v JellyfinJellyfinPlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *JellyfinPlayerStateInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### SetPlayMethodNil

`func (o *JellyfinPlayerStateInfo) SetPlayMethodNil(b bool)`

 SetPlayMethodNil sets the value for PlayMethod to be an explicit nil

### UnsetPlayMethod
`func (o *JellyfinPlayerStateInfo) UnsetPlayMethod()`

UnsetPlayMethod ensures that no value is present for PlayMethod, not even an explicit nil
### GetRepeatMode

`func (o *JellyfinPlayerStateInfo) GetRepeatMode() JellyfinJellyfinRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *JellyfinPlayerStateInfo) GetRepeatModeOk() (*JellyfinJellyfinRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *JellyfinPlayerStateInfo) SetRepeatMode(v JellyfinJellyfinRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *JellyfinPlayerStateInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetPlaybackOrder

`func (o *JellyfinPlayerStateInfo) GetPlaybackOrder() JellyfinJellyfinPlaybackOrder`

GetPlaybackOrder returns the PlaybackOrder field if non-nil, zero value otherwise.

### GetPlaybackOrderOk

`func (o *JellyfinPlayerStateInfo) GetPlaybackOrderOk() (*JellyfinJellyfinPlaybackOrder, bool)`

GetPlaybackOrderOk returns a tuple with the PlaybackOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackOrder

`func (o *JellyfinPlayerStateInfo) SetPlaybackOrder(v JellyfinJellyfinPlaybackOrder)`

SetPlaybackOrder sets PlaybackOrder field to given value.

### HasPlaybackOrder

`func (o *JellyfinPlayerStateInfo) HasPlaybackOrder() bool`

HasPlaybackOrder returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *JellyfinPlayerStateInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *JellyfinPlayerStateInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *JellyfinPlayerStateInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *JellyfinPlayerStateInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *JellyfinPlayerStateInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *JellyfinPlayerStateInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


