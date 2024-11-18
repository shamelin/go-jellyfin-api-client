# JellyfinOpenLiveStreamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenToken** | Pointer to **NullableString** | Gets or sets the open token. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session id. | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt32** | Gets or sets the max streaming bitrate. | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** | Gets or sets the start time in ticks. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the audio stream index. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the subtitle stream index. | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** | Gets or sets the max audio channels. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item id. | [optional] 
**EnableDirectPlay** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable direct play. | [optional] 
**EnableDirectStream** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enale direct stream. | [optional] 
**AlwaysBurnInSubtitleWhenTranscoding** | Pointer to **NullableBool** | Gets or sets a value indicating whether always burn in subtitles when transcoding. | [optional] 
**DeviceProfile** | Pointer to [**NullableJellyfinJellyfinDeviceProfile**](JellyfinDeviceProfile.md) | Gets or sets the device profile. | [optional] 
**DirectPlayProtocols** | Pointer to [**[]JellyfinJellyfinMediaProtocol**](JellyfinJellyfinMediaProtocol.md) | Gets or sets the device play protocols. | [optional] 

## Methods

### NewJellyfinOpenLiveStreamDto

`func NewJellyfinOpenLiveStreamDto() *JellyfinOpenLiveStreamDto`

NewJellyfinOpenLiveStreamDto instantiates a new JellyfinOpenLiveStreamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinOpenLiveStreamDtoWithDefaults

`func NewJellyfinOpenLiveStreamDtoWithDefaults() *JellyfinOpenLiveStreamDto`

NewJellyfinOpenLiveStreamDtoWithDefaults instantiates a new JellyfinOpenLiveStreamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenToken

`func (o *JellyfinOpenLiveStreamDto) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *JellyfinOpenLiveStreamDto) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *JellyfinOpenLiveStreamDto) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *JellyfinOpenLiveStreamDto) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *JellyfinOpenLiveStreamDto) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *JellyfinOpenLiveStreamDto) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetUserId

`func (o *JellyfinOpenLiveStreamDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinOpenLiveStreamDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinOpenLiveStreamDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinOpenLiveStreamDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *JellyfinOpenLiveStreamDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *JellyfinOpenLiveStreamDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetPlaySessionId

`func (o *JellyfinOpenLiveStreamDto) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *JellyfinOpenLiveStreamDto) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *JellyfinOpenLiveStreamDto) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *JellyfinOpenLiveStreamDto) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *JellyfinOpenLiveStreamDto) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *JellyfinOpenLiveStreamDto) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *JellyfinOpenLiveStreamDto) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *JellyfinOpenLiveStreamDto) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *JellyfinOpenLiveStreamDto) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *JellyfinOpenLiveStreamDto) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *JellyfinOpenLiveStreamDto) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *JellyfinOpenLiveStreamDto) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *JellyfinOpenLiveStreamDto) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *JellyfinOpenLiveStreamDto) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *JellyfinOpenLiveStreamDto) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *JellyfinOpenLiveStreamDto) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *JellyfinOpenLiveStreamDto) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *JellyfinOpenLiveStreamDto) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *JellyfinOpenLiveStreamDto) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *JellyfinOpenLiveStreamDto) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *JellyfinOpenLiveStreamDto) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *JellyfinOpenLiveStreamDto) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *JellyfinOpenLiveStreamDto) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *JellyfinOpenLiveStreamDto) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *JellyfinOpenLiveStreamDto) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *JellyfinOpenLiveStreamDto) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *JellyfinOpenLiveStreamDto) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *JellyfinOpenLiveStreamDto) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *JellyfinOpenLiveStreamDto) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *JellyfinOpenLiveStreamDto) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *JellyfinOpenLiveStreamDto) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *JellyfinOpenLiveStreamDto) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *JellyfinOpenLiveStreamDto) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *JellyfinOpenLiveStreamDto) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *JellyfinOpenLiveStreamDto) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *JellyfinOpenLiveStreamDto) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetItemId

`func (o *JellyfinOpenLiveStreamDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinOpenLiveStreamDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinOpenLiveStreamDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinOpenLiveStreamDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *JellyfinOpenLiveStreamDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *JellyfinOpenLiveStreamDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetEnableDirectPlay

`func (o *JellyfinOpenLiveStreamDto) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *JellyfinOpenLiveStreamDto) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *JellyfinOpenLiveStreamDto) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *JellyfinOpenLiveStreamDto) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### SetEnableDirectPlayNil

`func (o *JellyfinOpenLiveStreamDto) SetEnableDirectPlayNil(b bool)`

 SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil

### UnsetEnableDirectPlay
`func (o *JellyfinOpenLiveStreamDto) UnsetEnableDirectPlay()`

UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
### GetEnableDirectStream

`func (o *JellyfinOpenLiveStreamDto) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *JellyfinOpenLiveStreamDto) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *JellyfinOpenLiveStreamDto) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *JellyfinOpenLiveStreamDto) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### SetEnableDirectStreamNil

`func (o *JellyfinOpenLiveStreamDto) SetEnableDirectStreamNil(b bool)`

 SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil

### UnsetEnableDirectStream
`func (o *JellyfinOpenLiveStreamDto) UnsetEnableDirectStream()`

UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
### GetAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinOpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscoding() bool`

GetAlwaysBurnInSubtitleWhenTranscoding returns the AlwaysBurnInSubtitleWhenTranscoding field if non-nil, zero value otherwise.

### GetAlwaysBurnInSubtitleWhenTranscodingOk

`func (o *JellyfinOpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscodingOk() (*bool, bool)`

GetAlwaysBurnInSubtitleWhenTranscodingOk returns a tuple with the AlwaysBurnInSubtitleWhenTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinOpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscoding(v bool)`

SetAlwaysBurnInSubtitleWhenTranscoding sets AlwaysBurnInSubtitleWhenTranscoding field to given value.

### HasAlwaysBurnInSubtitleWhenTranscoding

`func (o *JellyfinOpenLiveStreamDto) HasAlwaysBurnInSubtitleWhenTranscoding() bool`

HasAlwaysBurnInSubtitleWhenTranscoding returns a boolean if a field has been set.

### SetAlwaysBurnInSubtitleWhenTranscodingNil

`func (o *JellyfinOpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscodingNil(b bool)`

 SetAlwaysBurnInSubtitleWhenTranscodingNil sets the value for AlwaysBurnInSubtitleWhenTranscoding to be an explicit nil

### UnsetAlwaysBurnInSubtitleWhenTranscoding
`func (o *JellyfinOpenLiveStreamDto) UnsetAlwaysBurnInSubtitleWhenTranscoding()`

UnsetAlwaysBurnInSubtitleWhenTranscoding ensures that no value is present for AlwaysBurnInSubtitleWhenTranscoding, not even an explicit nil
### GetDeviceProfile

`func (o *JellyfinOpenLiveStreamDto) GetDeviceProfile() JellyfinJellyfinDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *JellyfinOpenLiveStreamDto) GetDeviceProfileOk() (*JellyfinJellyfinDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *JellyfinOpenLiveStreamDto) SetDeviceProfile(v JellyfinJellyfinDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *JellyfinOpenLiveStreamDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *JellyfinOpenLiveStreamDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *JellyfinOpenLiveStreamDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetDirectPlayProtocols

`func (o *JellyfinOpenLiveStreamDto) GetDirectPlayProtocols() []JellyfinJellyfinMediaProtocol`

GetDirectPlayProtocols returns the DirectPlayProtocols field if non-nil, zero value otherwise.

### GetDirectPlayProtocolsOk

`func (o *JellyfinOpenLiveStreamDto) GetDirectPlayProtocolsOk() (*[]JellyfinJellyfinMediaProtocol, bool)`

GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProtocols

`func (o *JellyfinOpenLiveStreamDto) SetDirectPlayProtocols(v []JellyfinJellyfinMediaProtocol)`

SetDirectPlayProtocols sets DirectPlayProtocols field to given value.

### HasDirectPlayProtocols

`func (o *JellyfinOpenLiveStreamDto) HasDirectPlayProtocols() bool`

HasDirectPlayProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


