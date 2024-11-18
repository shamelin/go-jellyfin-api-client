# ModelOpenLiveStreamDto

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
**DeviceProfile** | Pointer to [**NullableModelModelDeviceProfile**](ModelDeviceProfile.md) | Gets or sets the device profile. | [optional] 
**DirectPlayProtocols** | Pointer to [**[]ModelModelMediaProtocol**](ModelModelMediaProtocol.md) | Gets or sets the device play protocols. | [optional] 

## Methods

### NewModelOpenLiveStreamDto

`func NewModelOpenLiveStreamDto() *ModelOpenLiveStreamDto`

NewModelOpenLiveStreamDto instantiates a new ModelOpenLiveStreamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOpenLiveStreamDtoWithDefaults

`func NewModelOpenLiveStreamDtoWithDefaults() *ModelOpenLiveStreamDto`

NewModelOpenLiveStreamDtoWithDefaults instantiates a new ModelOpenLiveStreamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenToken

`func (o *ModelOpenLiveStreamDto) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *ModelOpenLiveStreamDto) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *ModelOpenLiveStreamDto) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *ModelOpenLiveStreamDto) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *ModelOpenLiveStreamDto) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *ModelOpenLiveStreamDto) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetUserId

`func (o *ModelOpenLiveStreamDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelOpenLiveStreamDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelOpenLiveStreamDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelOpenLiveStreamDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ModelOpenLiveStreamDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ModelOpenLiveStreamDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetPlaySessionId

`func (o *ModelOpenLiveStreamDto) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *ModelOpenLiveStreamDto) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *ModelOpenLiveStreamDto) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *ModelOpenLiveStreamDto) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *ModelOpenLiveStreamDto) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *ModelOpenLiveStreamDto) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *ModelOpenLiveStreamDto) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ModelOpenLiveStreamDto) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ModelOpenLiveStreamDto) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ModelOpenLiveStreamDto) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ModelOpenLiveStreamDto) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ModelOpenLiveStreamDto) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *ModelOpenLiveStreamDto) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *ModelOpenLiveStreamDto) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *ModelOpenLiveStreamDto) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *ModelOpenLiveStreamDto) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *ModelOpenLiveStreamDto) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *ModelOpenLiveStreamDto) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *ModelOpenLiveStreamDto) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelOpenLiveStreamDto) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelOpenLiveStreamDto) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelOpenLiveStreamDto) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelOpenLiveStreamDto) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelOpenLiveStreamDto) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *ModelOpenLiveStreamDto) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelOpenLiveStreamDto) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelOpenLiveStreamDto) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelOpenLiveStreamDto) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelOpenLiveStreamDto) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelOpenLiveStreamDto) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *ModelOpenLiveStreamDto) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *ModelOpenLiveStreamDto) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *ModelOpenLiveStreamDto) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *ModelOpenLiveStreamDto) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *ModelOpenLiveStreamDto) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *ModelOpenLiveStreamDto) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetItemId

`func (o *ModelOpenLiveStreamDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelOpenLiveStreamDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelOpenLiveStreamDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelOpenLiveStreamDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ModelOpenLiveStreamDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ModelOpenLiveStreamDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetEnableDirectPlay

`func (o *ModelOpenLiveStreamDto) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *ModelOpenLiveStreamDto) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *ModelOpenLiveStreamDto) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *ModelOpenLiveStreamDto) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### SetEnableDirectPlayNil

`func (o *ModelOpenLiveStreamDto) SetEnableDirectPlayNil(b bool)`

 SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil

### UnsetEnableDirectPlay
`func (o *ModelOpenLiveStreamDto) UnsetEnableDirectPlay()`

UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
### GetEnableDirectStream

`func (o *ModelOpenLiveStreamDto) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *ModelOpenLiveStreamDto) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *ModelOpenLiveStreamDto) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *ModelOpenLiveStreamDto) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### SetEnableDirectStreamNil

`func (o *ModelOpenLiveStreamDto) SetEnableDirectStreamNil(b bool)`

 SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil

### UnsetEnableDirectStream
`func (o *ModelOpenLiveStreamDto) UnsetEnableDirectStream()`

UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
### GetAlwaysBurnInSubtitleWhenTranscoding

`func (o *ModelOpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscoding() bool`

GetAlwaysBurnInSubtitleWhenTranscoding returns the AlwaysBurnInSubtitleWhenTranscoding field if non-nil, zero value otherwise.

### GetAlwaysBurnInSubtitleWhenTranscodingOk

`func (o *ModelOpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscodingOk() (*bool, bool)`

GetAlwaysBurnInSubtitleWhenTranscodingOk returns a tuple with the AlwaysBurnInSubtitleWhenTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysBurnInSubtitleWhenTranscoding

`func (o *ModelOpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscoding(v bool)`

SetAlwaysBurnInSubtitleWhenTranscoding sets AlwaysBurnInSubtitleWhenTranscoding field to given value.

### HasAlwaysBurnInSubtitleWhenTranscoding

`func (o *ModelOpenLiveStreamDto) HasAlwaysBurnInSubtitleWhenTranscoding() bool`

HasAlwaysBurnInSubtitleWhenTranscoding returns a boolean if a field has been set.

### SetAlwaysBurnInSubtitleWhenTranscodingNil

`func (o *ModelOpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscodingNil(b bool)`

 SetAlwaysBurnInSubtitleWhenTranscodingNil sets the value for AlwaysBurnInSubtitleWhenTranscoding to be an explicit nil

### UnsetAlwaysBurnInSubtitleWhenTranscoding
`func (o *ModelOpenLiveStreamDto) UnsetAlwaysBurnInSubtitleWhenTranscoding()`

UnsetAlwaysBurnInSubtitleWhenTranscoding ensures that no value is present for AlwaysBurnInSubtitleWhenTranscoding, not even an explicit nil
### GetDeviceProfile

`func (o *ModelOpenLiveStreamDto) GetDeviceProfile() ModelModelDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ModelOpenLiveStreamDto) GetDeviceProfileOk() (*ModelModelDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ModelOpenLiveStreamDto) SetDeviceProfile(v ModelModelDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ModelOpenLiveStreamDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *ModelOpenLiveStreamDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *ModelOpenLiveStreamDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetDirectPlayProtocols

`func (o *ModelOpenLiveStreamDto) GetDirectPlayProtocols() []ModelModelMediaProtocol`

GetDirectPlayProtocols returns the DirectPlayProtocols field if non-nil, zero value otherwise.

### GetDirectPlayProtocolsOk

`func (o *ModelOpenLiveStreamDto) GetDirectPlayProtocolsOk() (*[]ModelModelMediaProtocol, bool)`

GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProtocols

`func (o *ModelOpenLiveStreamDto) SetDirectPlayProtocols(v []ModelModelMediaProtocol)`

SetDirectPlayProtocols sets DirectPlayProtocols field to given value.

### HasDirectPlayProtocols

`func (o *ModelOpenLiveStreamDto) HasDirectPlayProtocols() bool`

HasDirectPlayProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


