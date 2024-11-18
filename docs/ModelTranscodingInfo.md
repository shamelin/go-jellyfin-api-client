# ModelTranscodingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioCodec** | Pointer to **NullableString** | Gets or sets the thread count used for encoding. | [optional] 
**VideoCodec** | Pointer to **NullableString** | Gets or sets the thread count used for encoding. | [optional] 
**Container** | Pointer to **NullableString** | Gets or sets the thread count used for encoding. | [optional] 
**IsVideoDirect** | Pointer to **bool** | Gets or sets a value indicating whether the video is passed through. | [optional] 
**IsAudioDirect** | Pointer to **bool** | Gets or sets a value indicating whether the audio is passed through. | [optional] 
**Bitrate** | Pointer to **NullableInt32** | Gets or sets the bitrate. | [optional] 
**Framerate** | Pointer to **NullableFloat32** | Gets or sets the framerate. | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** | Gets or sets the completion percentage. | [optional] 
**Width** | Pointer to **NullableInt32** | Gets or sets the video width. | [optional] 
**Height** | Pointer to **NullableInt32** | Gets or sets the video height. | [optional] 
**AudioChannels** | Pointer to **NullableInt32** | Gets or sets the audio channels. | [optional] 
**HardwareAccelerationType** | Pointer to [**NullableModelModelHardwareAccelerationType**](ModelHardwareAccelerationType.md) | Gets or sets the hardware acceleration type. | [optional] 
**TranscodeReasons** | Pointer to [**[]ModelModelTranscodeReason**](ModelModelTranscodeReason.md) | Gets or sets the transcode reasons. | [optional] 

## Methods

### NewModelTranscodingInfo

`func NewModelTranscodingInfo() *ModelTranscodingInfo`

NewModelTranscodingInfo instantiates a new ModelTranscodingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTranscodingInfoWithDefaults

`func NewModelTranscodingInfoWithDefaults() *ModelTranscodingInfo`

NewModelTranscodingInfoWithDefaults instantiates a new ModelTranscodingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioCodec

`func (o *ModelTranscodingInfo) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelTranscodingInfo) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelTranscodingInfo) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelTranscodingInfo) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *ModelTranscodingInfo) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *ModelTranscodingInfo) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetVideoCodec

`func (o *ModelTranscodingInfo) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelTranscodingInfo) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelTranscodingInfo) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelTranscodingInfo) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *ModelTranscodingInfo) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *ModelTranscodingInfo) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetContainer

`func (o *ModelTranscodingInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelTranscodingInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelTranscodingInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelTranscodingInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ModelTranscodingInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ModelTranscodingInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetIsVideoDirect

`func (o *ModelTranscodingInfo) GetIsVideoDirect() bool`

GetIsVideoDirect returns the IsVideoDirect field if non-nil, zero value otherwise.

### GetIsVideoDirectOk

`func (o *ModelTranscodingInfo) GetIsVideoDirectOk() (*bool, bool)`

GetIsVideoDirectOk returns a tuple with the IsVideoDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideoDirect

`func (o *ModelTranscodingInfo) SetIsVideoDirect(v bool)`

SetIsVideoDirect sets IsVideoDirect field to given value.

### HasIsVideoDirect

`func (o *ModelTranscodingInfo) HasIsVideoDirect() bool`

HasIsVideoDirect returns a boolean if a field has been set.

### GetIsAudioDirect

`func (o *ModelTranscodingInfo) GetIsAudioDirect() bool`

GetIsAudioDirect returns the IsAudioDirect field if non-nil, zero value otherwise.

### GetIsAudioDirectOk

`func (o *ModelTranscodingInfo) GetIsAudioDirectOk() (*bool, bool)`

GetIsAudioDirectOk returns a tuple with the IsAudioDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAudioDirect

`func (o *ModelTranscodingInfo) SetIsAudioDirect(v bool)`

SetIsAudioDirect sets IsAudioDirect field to given value.

### HasIsAudioDirect

`func (o *ModelTranscodingInfo) HasIsAudioDirect() bool`

HasIsAudioDirect returns a boolean if a field has been set.

### GetBitrate

`func (o *ModelTranscodingInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelTranscodingInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelTranscodingInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelTranscodingInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelTranscodingInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelTranscodingInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetFramerate

`func (o *ModelTranscodingInfo) GetFramerate() float32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *ModelTranscodingInfo) GetFramerateOk() (*float32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *ModelTranscodingInfo) SetFramerate(v float32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *ModelTranscodingInfo) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### SetFramerateNil

`func (o *ModelTranscodingInfo) SetFramerateNil(b bool)`

 SetFramerateNil sets the value for Framerate to be an explicit nil

### UnsetFramerate
`func (o *ModelTranscodingInfo) UnsetFramerate()`

UnsetFramerate ensures that no value is present for Framerate, not even an explicit nil
### GetCompletionPercentage

`func (o *ModelTranscodingInfo) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *ModelTranscodingInfo) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *ModelTranscodingInfo) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *ModelTranscodingInfo) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *ModelTranscodingInfo) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *ModelTranscodingInfo) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetWidth

`func (o *ModelTranscodingInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelTranscodingInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelTranscodingInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelTranscodingInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelTranscodingInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelTranscodingInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *ModelTranscodingInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelTranscodingInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelTranscodingInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelTranscodingInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelTranscodingInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelTranscodingInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetAudioChannels

`func (o *ModelTranscodingInfo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *ModelTranscodingInfo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *ModelTranscodingInfo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *ModelTranscodingInfo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannelsNil

`func (o *ModelTranscodingInfo) SetAudioChannelsNil(b bool)`

 SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil

### UnsetAudioChannels
`func (o *ModelTranscodingInfo) UnsetAudioChannels()`

UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
### GetHardwareAccelerationType

`func (o *ModelTranscodingInfo) GetHardwareAccelerationType() ModelModelHardwareAccelerationType`

GetHardwareAccelerationType returns the HardwareAccelerationType field if non-nil, zero value otherwise.

### GetHardwareAccelerationTypeOk

`func (o *ModelTranscodingInfo) GetHardwareAccelerationTypeOk() (*ModelModelHardwareAccelerationType, bool)`

GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationType

`func (o *ModelTranscodingInfo) SetHardwareAccelerationType(v ModelModelHardwareAccelerationType)`

SetHardwareAccelerationType sets HardwareAccelerationType field to given value.

### HasHardwareAccelerationType

`func (o *ModelTranscodingInfo) HasHardwareAccelerationType() bool`

HasHardwareAccelerationType returns a boolean if a field has been set.

### SetHardwareAccelerationTypeNil

`func (o *ModelTranscodingInfo) SetHardwareAccelerationTypeNil(b bool)`

 SetHardwareAccelerationTypeNil sets the value for HardwareAccelerationType to be an explicit nil

### UnsetHardwareAccelerationType
`func (o *ModelTranscodingInfo) UnsetHardwareAccelerationType()`

UnsetHardwareAccelerationType ensures that no value is present for HardwareAccelerationType, not even an explicit nil
### GetTranscodeReasons

`func (o *ModelTranscodingInfo) GetTranscodeReasons() []ModelModelTranscodeReason`

GetTranscodeReasons returns the TranscodeReasons field if non-nil, zero value otherwise.

### GetTranscodeReasonsOk

`func (o *ModelTranscodingInfo) GetTranscodeReasonsOk() (*[]ModelModelTranscodeReason, bool)`

GetTranscodeReasonsOk returns a tuple with the TranscodeReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeReasons

`func (o *ModelTranscodingInfo) SetTranscodeReasons(v []ModelModelTranscodeReason)`

SetTranscodeReasons sets TranscodeReasons field to given value.

### HasTranscodeReasons

`func (o *ModelTranscodingInfo) HasTranscodeReasons() bool`

HasTranscodeReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


