# JellyfinTranscodingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | Gets or sets the container. | [optional] 
**Type** | Pointer to [**JellyfinJellyfinDlnaProfileType**](JellyfinDlnaProfileType.md) | Gets or sets the DLNA profile type. | [optional] 
**VideoCodec** | Pointer to **string** | Gets or sets the video codec. | [optional] 
**AudioCodec** | Pointer to **string** | Gets or sets the audio codec. | [optional] 
**Protocol** | Pointer to [**JellyfinJellyfinMediaStreamProtocol**](JellyfinMediaStreamProtocol.md) | Media streaming protocol.  Lowercase for backwards compatibility. | [optional] 
**EstimateContentLength** | Pointer to **bool** | Gets or sets a value indicating whether the content length should be estimated. | [optional] [default to false]
**EnableMpegtsM2TsMode** | Pointer to **bool** | Gets or sets a value indicating whether M2TS mode is enabled. | [optional] [default to false]
**TranscodeSeekInfo** | Pointer to [**JellyfinJellyfinTranscodeSeekInfo**](JellyfinTranscodeSeekInfo.md) | Gets or sets the transcoding seek info mode. | [optional] [default to JELLYFINJELLYFINTRANSCODESEEKINFO_AUTO]
**CopyTimestamps** | Pointer to **bool** | Gets or sets a value indicating whether timestamps should be copied. | [optional] [default to false]
**Context** | Pointer to [**JellyfinJellyfinEncodingContext**](JellyfinEncodingContext.md) | Gets or sets the encoding context. | [optional] [default to JELLYFINJELLYFINENCODINGCONTEXT_STREAMING]
**EnableSubtitlesInManifest** | Pointer to **bool** | Gets or sets a value indicating whether subtitles are allowed in the manifest. | [optional] [default to false]
**MaxAudioChannels** | Pointer to **NullableString** | Gets or sets the maximum audio channels. | [optional] 
**MinSegments** | Pointer to **int32** | Gets or sets the minimum amount of segments. | [optional] [default to 0]
**SegmentLength** | Pointer to **int32** | Gets or sets the segment length. | [optional] [default to 0]
**BreakOnNonKeyFrames** | Pointer to **bool** | Gets or sets a value indicating whether breaking the video stream on non-keyframes is supported. | [optional] [default to false]
**Conditions** | Pointer to [**[]JellyfinJellyfinProfileCondition**](JellyfinJellyfinProfileCondition.md) | Gets or sets the profile conditions. | [optional] 
**EnableAudioVbrEncoding** | Pointer to **bool** | Gets or sets a value indicating whether variable bitrate encoding is supported. | [optional] [default to true]

## Methods

### NewJellyfinTranscodingProfile

`func NewJellyfinTranscodingProfile() *JellyfinTranscodingProfile`

NewJellyfinTranscodingProfile instantiates a new JellyfinTranscodingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinTranscodingProfileWithDefaults

`func NewJellyfinTranscodingProfileWithDefaults() *JellyfinTranscodingProfile`

NewJellyfinTranscodingProfileWithDefaults instantiates a new JellyfinTranscodingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *JellyfinTranscodingProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *JellyfinTranscodingProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *JellyfinTranscodingProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *JellyfinTranscodingProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetType

`func (o *JellyfinTranscodingProfile) GetType() JellyfinJellyfinDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinTranscodingProfile) GetTypeOk() (*JellyfinJellyfinDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinTranscodingProfile) SetType(v JellyfinJellyfinDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinTranscodingProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVideoCodec

`func (o *JellyfinTranscodingProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *JellyfinTranscodingProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *JellyfinTranscodingProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *JellyfinTranscodingProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *JellyfinTranscodingProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *JellyfinTranscodingProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *JellyfinTranscodingProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *JellyfinTranscodingProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetProtocol

`func (o *JellyfinTranscodingProfile) GetProtocol() JellyfinJellyfinMediaStreamProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *JellyfinTranscodingProfile) GetProtocolOk() (*JellyfinJellyfinMediaStreamProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *JellyfinTranscodingProfile) SetProtocol(v JellyfinJellyfinMediaStreamProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *JellyfinTranscodingProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEstimateContentLength

`func (o *JellyfinTranscodingProfile) GetEstimateContentLength() bool`

GetEstimateContentLength returns the EstimateContentLength field if non-nil, zero value otherwise.

### GetEstimateContentLengthOk

`func (o *JellyfinTranscodingProfile) GetEstimateContentLengthOk() (*bool, bool)`

GetEstimateContentLengthOk returns a tuple with the EstimateContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateContentLength

`func (o *JellyfinTranscodingProfile) SetEstimateContentLength(v bool)`

SetEstimateContentLength sets EstimateContentLength field to given value.

### HasEstimateContentLength

`func (o *JellyfinTranscodingProfile) HasEstimateContentLength() bool`

HasEstimateContentLength returns a boolean if a field has been set.

### GetEnableMpegtsM2TsMode

`func (o *JellyfinTranscodingProfile) GetEnableMpegtsM2TsMode() bool`

GetEnableMpegtsM2TsMode returns the EnableMpegtsM2TsMode field if non-nil, zero value otherwise.

### GetEnableMpegtsM2TsModeOk

`func (o *JellyfinTranscodingProfile) GetEnableMpegtsM2TsModeOk() (*bool, bool)`

GetEnableMpegtsM2TsModeOk returns a tuple with the EnableMpegtsM2TsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMpegtsM2TsMode

`func (o *JellyfinTranscodingProfile) SetEnableMpegtsM2TsMode(v bool)`

SetEnableMpegtsM2TsMode sets EnableMpegtsM2TsMode field to given value.

### HasEnableMpegtsM2TsMode

`func (o *JellyfinTranscodingProfile) HasEnableMpegtsM2TsMode() bool`

HasEnableMpegtsM2TsMode returns a boolean if a field has been set.

### GetTranscodeSeekInfo

`func (o *JellyfinTranscodingProfile) GetTranscodeSeekInfo() JellyfinJellyfinTranscodeSeekInfo`

GetTranscodeSeekInfo returns the TranscodeSeekInfo field if non-nil, zero value otherwise.

### GetTranscodeSeekInfoOk

`func (o *JellyfinTranscodingProfile) GetTranscodeSeekInfoOk() (*JellyfinJellyfinTranscodeSeekInfo, bool)`

GetTranscodeSeekInfoOk returns a tuple with the TranscodeSeekInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeSeekInfo

`func (o *JellyfinTranscodingProfile) SetTranscodeSeekInfo(v JellyfinJellyfinTranscodeSeekInfo)`

SetTranscodeSeekInfo sets TranscodeSeekInfo field to given value.

### HasTranscodeSeekInfo

`func (o *JellyfinTranscodingProfile) HasTranscodeSeekInfo() bool`

HasTranscodeSeekInfo returns a boolean if a field has been set.

### GetCopyTimestamps

`func (o *JellyfinTranscodingProfile) GetCopyTimestamps() bool`

GetCopyTimestamps returns the CopyTimestamps field if non-nil, zero value otherwise.

### GetCopyTimestampsOk

`func (o *JellyfinTranscodingProfile) GetCopyTimestampsOk() (*bool, bool)`

GetCopyTimestampsOk returns a tuple with the CopyTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTimestamps

`func (o *JellyfinTranscodingProfile) SetCopyTimestamps(v bool)`

SetCopyTimestamps sets CopyTimestamps field to given value.

### HasCopyTimestamps

`func (o *JellyfinTranscodingProfile) HasCopyTimestamps() bool`

HasCopyTimestamps returns a boolean if a field has been set.

### GetContext

`func (o *JellyfinTranscodingProfile) GetContext() JellyfinJellyfinEncodingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *JellyfinTranscodingProfile) GetContextOk() (*JellyfinJellyfinEncodingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *JellyfinTranscodingProfile) SetContext(v JellyfinJellyfinEncodingContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *JellyfinTranscodingProfile) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEnableSubtitlesInManifest

`func (o *JellyfinTranscodingProfile) GetEnableSubtitlesInManifest() bool`

GetEnableSubtitlesInManifest returns the EnableSubtitlesInManifest field if non-nil, zero value otherwise.

### GetEnableSubtitlesInManifestOk

`func (o *JellyfinTranscodingProfile) GetEnableSubtitlesInManifestOk() (*bool, bool)`

GetEnableSubtitlesInManifestOk returns a tuple with the EnableSubtitlesInManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitlesInManifest

`func (o *JellyfinTranscodingProfile) SetEnableSubtitlesInManifest(v bool)`

SetEnableSubtitlesInManifest sets EnableSubtitlesInManifest field to given value.

### HasEnableSubtitlesInManifest

`func (o *JellyfinTranscodingProfile) HasEnableSubtitlesInManifest() bool`

HasEnableSubtitlesInManifest returns a boolean if a field has been set.

### GetMaxAudioChannels

`func (o *JellyfinTranscodingProfile) GetMaxAudioChannels() string`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *JellyfinTranscodingProfile) GetMaxAudioChannelsOk() (*string, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *JellyfinTranscodingProfile) SetMaxAudioChannels(v string)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *JellyfinTranscodingProfile) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *JellyfinTranscodingProfile) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *JellyfinTranscodingProfile) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetMinSegments

`func (o *JellyfinTranscodingProfile) GetMinSegments() int32`

GetMinSegments returns the MinSegments field if non-nil, zero value otherwise.

### GetMinSegmentsOk

`func (o *JellyfinTranscodingProfile) GetMinSegmentsOk() (*int32, bool)`

GetMinSegmentsOk returns a tuple with the MinSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSegments

`func (o *JellyfinTranscodingProfile) SetMinSegments(v int32)`

SetMinSegments sets MinSegments field to given value.

### HasMinSegments

`func (o *JellyfinTranscodingProfile) HasMinSegments() bool`

HasMinSegments returns a boolean if a field has been set.

### GetSegmentLength

`func (o *JellyfinTranscodingProfile) GetSegmentLength() int32`

GetSegmentLength returns the SegmentLength field if non-nil, zero value otherwise.

### GetSegmentLengthOk

`func (o *JellyfinTranscodingProfile) GetSegmentLengthOk() (*int32, bool)`

GetSegmentLengthOk returns a tuple with the SegmentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentLength

`func (o *JellyfinTranscodingProfile) SetSegmentLength(v int32)`

SetSegmentLength sets SegmentLength field to given value.

### HasSegmentLength

`func (o *JellyfinTranscodingProfile) HasSegmentLength() bool`

HasSegmentLength returns a boolean if a field has been set.

### GetBreakOnNonKeyFrames

`func (o *JellyfinTranscodingProfile) GetBreakOnNonKeyFrames() bool`

GetBreakOnNonKeyFrames returns the BreakOnNonKeyFrames field if non-nil, zero value otherwise.

### GetBreakOnNonKeyFramesOk

`func (o *JellyfinTranscodingProfile) GetBreakOnNonKeyFramesOk() (*bool, bool)`

GetBreakOnNonKeyFramesOk returns a tuple with the BreakOnNonKeyFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakOnNonKeyFrames

`func (o *JellyfinTranscodingProfile) SetBreakOnNonKeyFrames(v bool)`

SetBreakOnNonKeyFrames sets BreakOnNonKeyFrames field to given value.

### HasBreakOnNonKeyFrames

`func (o *JellyfinTranscodingProfile) HasBreakOnNonKeyFrames() bool`

HasBreakOnNonKeyFrames returns a boolean if a field has been set.

### GetConditions

`func (o *JellyfinTranscodingProfile) GetConditions() []JellyfinJellyfinProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *JellyfinTranscodingProfile) GetConditionsOk() (*[]JellyfinJellyfinProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *JellyfinTranscodingProfile) SetConditions(v []JellyfinJellyfinProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *JellyfinTranscodingProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetEnableAudioVbrEncoding

`func (o *JellyfinTranscodingProfile) GetEnableAudioVbrEncoding() bool`

GetEnableAudioVbrEncoding returns the EnableAudioVbrEncoding field if non-nil, zero value otherwise.

### GetEnableAudioVbrEncodingOk

`func (o *JellyfinTranscodingProfile) GetEnableAudioVbrEncodingOk() (*bool, bool)`

GetEnableAudioVbrEncodingOk returns a tuple with the EnableAudioVbrEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioVbrEncoding

`func (o *JellyfinTranscodingProfile) SetEnableAudioVbrEncoding(v bool)`

SetEnableAudioVbrEncoding sets EnableAudioVbrEncoding field to given value.

### HasEnableAudioVbrEncoding

`func (o *JellyfinTranscodingProfile) HasEnableAudioVbrEncoding() bool`

HasEnableAudioVbrEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


