# JellyfinPlaybackInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | Gets or sets the playback userId. | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt32** | Gets or sets the max streaming bitrate. | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** | Gets or sets the start time in ticks. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the audio stream index. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the subtitle stream index. | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** | Gets or sets the max audio channels. | [optional] 
**MediaSourceId** | Pointer to **NullableString** | Gets or sets the media source id. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the live stream id. | [optional] 
**DeviceProfile** | Pointer to [**NullableJellyfinJellyfinDeviceProfile**](JellyfinDeviceProfile.md) | Gets or sets the device profile. | [optional] 
**EnableDirectPlay** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable direct play. | [optional] 
**EnableDirectStream** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable direct stream. | [optional] 
**EnableTranscoding** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable transcoding. | [optional] 
**AllowVideoStreamCopy** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable video stream copy. | [optional] 
**AllowAudioStreamCopy** | Pointer to **NullableBool** | Gets or sets a value indicating whether to allow audio stream copy. | [optional] 
**AutoOpenLiveStream** | Pointer to **NullableBool** | Gets or sets a value indicating whether to auto open the live stream. | [optional] 
**AlwaysBurnInSubtitleWhenTranscoding** | Pointer to **NullableBool** | Gets or sets a value indicating whether always burn in subtitles when transcoding. | [optional] 

## Methods

### NewJellyfinPlaybackInfoDto

`func NewJellyfinPlaybackInfoDto() *JellyfinPlaybackInfoDto`

NewJellyfinPlaybackInfoDto instantiates a new JellyfinPlaybackInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlaybackInfoDtoWithDefaults

`func NewJellyfinPlaybackInfoDtoWithDefaults() *JellyfinPlaybackInfoDto`

NewJellyfinPlaybackInfoDtoWithDefaults instantiates a new JellyfinPlaybackInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *JellyfinPlaybackInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinPlaybackInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinPlaybackInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinPlaybackInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *JellyfinPlaybackInfoDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *JellyfinPlaybackInfoDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *JellyfinPlaybackInfoDto) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *JellyfinPlaybackInfoDto) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *JellyfinPlaybackInfoDto) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *JellyfinPlaybackInfoDto) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *JellyfinPlaybackInfoDto) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *JellyfinPlaybackInfoDto) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *JellyfinPlaybackInfoDto) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *JellyfinPlaybackInfoDto) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *JellyfinPlaybackInfoDto) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *JellyfinPlaybackInfoDto) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *JellyfinPlaybackInfoDto) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *JellyfinPlaybackInfoDto) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *JellyfinPlaybackInfoDto) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *JellyfinPlaybackInfoDto) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *JellyfinPlaybackInfoDto) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *JellyfinPlaybackInfoDto) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *JellyfinPlaybackInfoDto) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *JellyfinPlaybackInfoDto) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *JellyfinPlaybackInfoDto) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *JellyfinPlaybackInfoDto) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *JellyfinPlaybackInfoDto) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *JellyfinPlaybackInfoDto) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *JellyfinPlaybackInfoDto) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *JellyfinPlaybackInfoDto) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *JellyfinPlaybackInfoDto) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *JellyfinPlaybackInfoDto) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *JellyfinPlaybackInfoDto) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *JellyfinPlaybackInfoDto) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *JellyfinPlaybackInfoDto) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *JellyfinPlaybackInfoDto) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetMediaSourceId

`func (o *JellyfinPlaybackInfoDto) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *JellyfinPlaybackInfoDto) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *JellyfinPlaybackInfoDto) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *JellyfinPlaybackInfoDto) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *JellyfinPlaybackInfoDto) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *JellyfinPlaybackInfoDto) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetLiveStreamId

`func (o *JellyfinPlaybackInfoDto) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *JellyfinPlaybackInfoDto) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *JellyfinPlaybackInfoDto) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *JellyfinPlaybackInfoDto) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *JellyfinPlaybackInfoDto) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *JellyfinPlaybackInfoDto) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetDeviceProfile

`func (o *JellyfinPlaybackInfoDto) GetDeviceProfile() JellyfinJellyfinDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *JellyfinPlaybackInfoDto) GetDeviceProfileOk() (*JellyfinJellyfinDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *JellyfinPlaybackInfoDto) SetDeviceProfile(v JellyfinJellyfinDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *JellyfinPlaybackInfoDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *JellyfinPlaybackInfoDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *JellyfinPlaybackInfoDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetEnableDirectPlay

`func (o *JellyfinPlaybackInfoDto) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *JellyfinPlaybackInfoDto) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *JellyfinPlaybackInfoDto) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *JellyfinPlaybackInfoDto) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### SetEnableDirectPlayNil

`func (o *JellyfinPlaybackInfoDto) SetEnableDirectPlayNil(b bool)`

 SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil

### UnsetEnableDirectPlay
`func (o *JellyfinPlaybackInfoDto) UnsetEnableDirectPlay()`

UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
### GetEnableDirectStream

`func (o *JellyfinPlaybackInfoDto) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *JellyfinPlaybackInfoDto) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *JellyfinPlaybackInfoDto) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *JellyfinPlaybackInfoDto) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### SetEnableDirectStreamNil

`func (o *JellyfinPlaybackInfoDto) SetEnableDirectStreamNil(b bool)`

 SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil

### UnsetEnableDirectStream
`func (o *JellyfinPlaybackInfoDto) UnsetEnableDirectStream()`

UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
### GetEnableTranscoding

`func (o *JellyfinPlaybackInfoDto) GetEnableTranscoding() bool`

GetEnableTranscoding returns the EnableTranscoding field if non-nil, zero value otherwise.

### GetEnableTranscodingOk

`func (o *JellyfinPlaybackInfoDto) GetEnableTranscodingOk() (*bool, bool)`

GetEnableTranscodingOk returns a tuple with the EnableTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTranscoding

`func (o *JellyfinPlaybackInfoDto) SetEnableTranscoding(v bool)`

SetEnableTranscoding sets EnableTranscoding field to given value.

### HasEnableTranscoding

`func (o *JellyfinPlaybackInfoDto) HasEnableTranscoding() bool`

HasEnableTranscoding returns a boolean if a field has been set.

### SetEnableTranscodingNil

`func (o *JellyfinPlaybackInfoDto) SetEnableTranscodingNil(b bool)`

 SetEnableTranscodingNil sets the value for EnableTranscoding to be an explicit nil

### UnsetEnableTranscoding
`func (o *JellyfinPlaybackInfoDto) UnsetEnableTranscoding()`

UnsetEnableTranscoding ensures that no value is present for EnableTranscoding, not even an explicit nil
### GetAllowVideoStreamCopy

`func (o *JellyfinPlaybackInfoDto) GetAllowVideoStreamCopy() bool`

GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowVideoStreamCopyOk

`func (o *JellyfinPlaybackInfoDto) GetAllowVideoStreamCopyOk() (*bool, bool)`

GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideoStreamCopy

`func (o *JellyfinPlaybackInfoDto) SetAllowVideoStreamCopy(v bool)`

SetAllowVideoStreamCopy sets AllowVideoStreamCopy field to given value.

### HasAllowVideoStreamCopy

`func (o *JellyfinPlaybackInfoDto) HasAllowVideoStreamCopy() bool`

HasAllowVideoStreamCopy returns a boolean if a field has been set.

### SetAllowVideoStreamCopyNil

`func (o *JellyfinPlaybackInfoDto) SetAllowVideoStreamCopyNil(b bool)`

 SetAllowVideoStreamCopyNil sets the value for AllowVideoStreamCopy to be an explicit nil

### UnsetAllowVideoStreamCopy
`func (o *JellyfinPlaybackInfoDto) UnsetAllowVideoStreamCopy()`

UnsetAllowVideoStreamCopy ensures that no value is present for AllowVideoStreamCopy, not even an explicit nil
### GetAllowAudioStreamCopy

`func (o *JellyfinPlaybackInfoDto) GetAllowAudioStreamCopy() bool`

GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field if non-nil, zero value otherwise.

### GetAllowAudioStreamCopyOk

`func (o *JellyfinPlaybackInfoDto) GetAllowAudioStreamCopyOk() (*bool, bool)`

GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAudioStreamCopy

`func (o *JellyfinPlaybackInfoDto) SetAllowAudioStreamCopy(v bool)`

SetAllowAudioStreamCopy sets AllowAudioStreamCopy field to given value.

### HasAllowAudioStreamCopy

`func (o *JellyfinPlaybackInfoDto) HasAllowAudioStreamCopy() bool`

HasAllowAudioStreamCopy returns a boolean if a field has been set.

### SetAllowAudioStreamCopyNil

`func (o *JellyfinPlaybackInfoDto) SetAllowAudioStreamCopyNil(b bool)`

 SetAllowAudioStreamCopyNil sets the value for AllowAudioStreamCopy to be an explicit nil

### UnsetAllowAudioStreamCopy
`func (o *JellyfinPlaybackInfoDto) UnsetAllowAudioStreamCopy()`

UnsetAllowAudioStreamCopy ensures that no value is present for AllowAudioStreamCopy, not even an explicit nil
### GetAutoOpenLiveStream

`func (o *JellyfinPlaybackInfoDto) GetAutoOpenLiveStream() bool`

GetAutoOpenLiveStream returns the AutoOpenLiveStream field if non-nil, zero value otherwise.

### GetAutoOpenLiveStreamOk

`func (o *JellyfinPlaybackInfoDto) GetAutoOpenLiveStreamOk() (*bool, bool)`

GetAutoOpenLiveStreamOk returns a tuple with the AutoOpenLiveStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOpenLiveStream

`func (o *JellyfinPlaybackInfoDto) SetAutoOpenLiveStream(v bool)`

SetAutoOpenLiveStream sets AutoOpenLiveStream field to given value.

### HasAutoOpenLiveStream

`func (o *JellyfinPlaybackInfoDto) HasAutoOpenLiveStream() bool`

HasAutoOpenLiveStream returns a boolean if a field has been set.

### SetAutoOpenLiveStreamNil

`func (o *JellyfinPlaybackInfoDto) SetAutoOpenLiveStreamNil(b bool)`

 SetAutoOpenLiveStreamNil sets the value for AutoOpenLiveStream to be an explicit nil

### UnsetAutoOpenLiveStream
`func (o *JellyfinPlaybackInfoDto) UnsetAutoOpenLiveStream()`

UnsetAutoOpenLiveStream ensures that no value is present for AutoOpenLiveStream, not even an explicit nil
### GetAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinPlaybackInfoDto) GetAlwaysBurnInSubtitleWhenTranscoding() bool`

GetAlwaysBurnInSubtitleWhenTranscoding returns the AlwaysBurnInSubtitleWhenTranscoding field if non-nil, zero value otherwise.

### GetAlwaysBurnInSubtitleWhenTranscodingOk

`func (o *JellyfinPlaybackInfoDto) GetAlwaysBurnInSubtitleWhenTranscodingOk() (*bool, bool)`

GetAlwaysBurnInSubtitleWhenTranscodingOk returns a tuple with the AlwaysBurnInSubtitleWhenTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinPlaybackInfoDto) SetAlwaysBurnInSubtitleWhenTranscoding(v bool)`

SetAlwaysBurnInSubtitleWhenTranscoding sets AlwaysBurnInSubtitleWhenTranscoding field to given value.

### HasAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinPlaybackInfoDto) HasAlwaysBurnInSubtitleWhenTranscoding() bool`

HasAlwaysBurnInSubtitleWhenTranscoding returns a boolean if a field has been set.

### SetAlwaysBurnInSubtitleWhenTranscodingNil

`func (o *JellyfinPlaybackInfoDto) SetAlwaysBurnInSubtitleWhenTranscodingNil(b bool)`

 SetAlwaysBurnInSubtitleWhenTranscodingNil sets the value for AlwaysBurnInSubtitleWhenTranscoding to be an explicit nil

### UnsetAlwaysBurnInSubtitleWhenTranscoding
`func (o *JellyfinPlaybackInfoDto) UnsetAlwaysBurnInSubtitleWhenTranscoding()`

UnsetAlwaysBurnInSubtitleWhenTranscoding ensures that no value is present for AlwaysBurnInSubtitleWhenTranscoding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


