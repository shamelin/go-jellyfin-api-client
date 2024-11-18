# ModelMediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**ModelModelMediaProtocol**](ModelMediaProtocol.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**EncoderPath** | Pointer to **NullableString** |  | [optional] 
**EncoderProtocol** | Pointer to [**NullableModelModelMediaProtocol**](ModelMediaProtocol.md) |  | [optional] 
**Type** | Pointer to [**ModelModelMediaSourceType**](ModelMediaSourceType.md) |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsRemote** | Pointer to **bool** | Gets or sets a value indicating whether the media is remote.  Differentiate internet url vs local network. | [optional] 
**ETag** | Pointer to **NullableString** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**ReadAtNativeFramerate** | Pointer to **bool** |  | [optional] 
**IgnoreDts** | Pointer to **bool** |  | [optional] 
**IgnoreIndex** | Pointer to **bool** |  | [optional] 
**GenPtsInput** | Pointer to **bool** |  | [optional] 
**SupportsTranscoding** | Pointer to **bool** |  | [optional] 
**SupportsDirectStream** | Pointer to **bool** |  | [optional] 
**SupportsDirectPlay** | Pointer to **bool** |  | [optional] 
**IsInfiniteStream** | Pointer to **bool** |  | [optional] 
**UseMostCompatibleTranscodingProfile** | Pointer to **bool** |  | [optional] [default to false]
**RequiresOpening** | Pointer to **bool** |  | [optional] 
**OpenToken** | Pointer to **NullableString** |  | [optional] 
**RequiresClosing** | Pointer to **bool** |  | [optional] 
**LiveStreamId** | Pointer to **NullableString** |  | [optional] 
**BufferMs** | Pointer to **NullableInt32** |  | [optional] 
**RequiresLooping** | Pointer to **bool** |  | [optional] 
**SupportsProbing** | Pointer to **bool** |  | [optional] 
**VideoType** | Pointer to [**NullableModelModelVideoType**](ModelVideoType.md) |  | [optional] 
**IsoType** | Pointer to [**NullableModelModelIsoType**](ModelIsoType.md) |  | [optional] 
**Video3DFormat** | Pointer to [**NullableModelModelVideo3DFormat**](ModelVideo3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]ModelModelMediaStream**](ModelModelMediaStream.md) |  | [optional] 
**MediaAttachments** | Pointer to [**[]ModelModelMediaAttachment**](ModelModelMediaAttachment.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**FallbackMaxStreamingBitrate** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**NullableModelModelTransportStreamTimestamp**](ModelTransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**TranscodingUrl** | Pointer to **NullableString** |  | [optional] 
**TranscodingSubProtocol** | Pointer to [**ModelModelMediaStreamProtocol**](ModelMediaStreamProtocol.md) | Media streaming protocol.  Lowercase for backwards compatibility. | [optional] 
**TranscodingContainer** | Pointer to **NullableString** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **NullableInt32** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**HasSegments** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelMediaSourceInfo

`func NewModelMediaSourceInfo() *ModelMediaSourceInfo`

NewModelMediaSourceInfo instantiates a new ModelMediaSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMediaSourceInfoWithDefaults

`func NewModelMediaSourceInfoWithDefaults() *ModelMediaSourceInfo`

NewModelMediaSourceInfoWithDefaults instantiates a new ModelMediaSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ModelMediaSourceInfo) GetProtocol() ModelModelMediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelMediaSourceInfo) GetProtocolOk() (*ModelModelMediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelMediaSourceInfo) SetProtocol(v ModelModelMediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ModelMediaSourceInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetId

`func (o *ModelMediaSourceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelMediaSourceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelMediaSourceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelMediaSourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelMediaSourceInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelMediaSourceInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *ModelMediaSourceInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelMediaSourceInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelMediaSourceInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelMediaSourceInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ModelMediaSourceInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ModelMediaSourceInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEncoderPath

`func (o *ModelMediaSourceInfo) GetEncoderPath() string`

GetEncoderPath returns the EncoderPath field if non-nil, zero value otherwise.

### GetEncoderPathOk

`func (o *ModelMediaSourceInfo) GetEncoderPathOk() (*string, bool)`

GetEncoderPathOk returns a tuple with the EncoderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPath

`func (o *ModelMediaSourceInfo) SetEncoderPath(v string)`

SetEncoderPath sets EncoderPath field to given value.

### HasEncoderPath

`func (o *ModelMediaSourceInfo) HasEncoderPath() bool`

HasEncoderPath returns a boolean if a field has been set.

### SetEncoderPathNil

`func (o *ModelMediaSourceInfo) SetEncoderPathNil(b bool)`

 SetEncoderPathNil sets the value for EncoderPath to be an explicit nil

### UnsetEncoderPath
`func (o *ModelMediaSourceInfo) UnsetEncoderPath()`

UnsetEncoderPath ensures that no value is present for EncoderPath, not even an explicit nil
### GetEncoderProtocol

`func (o *ModelMediaSourceInfo) GetEncoderProtocol() ModelModelMediaProtocol`

GetEncoderProtocol returns the EncoderProtocol field if non-nil, zero value otherwise.

### GetEncoderProtocolOk

`func (o *ModelMediaSourceInfo) GetEncoderProtocolOk() (*ModelModelMediaProtocol, bool)`

GetEncoderProtocolOk returns a tuple with the EncoderProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderProtocol

`func (o *ModelMediaSourceInfo) SetEncoderProtocol(v ModelModelMediaProtocol)`

SetEncoderProtocol sets EncoderProtocol field to given value.

### HasEncoderProtocol

`func (o *ModelMediaSourceInfo) HasEncoderProtocol() bool`

HasEncoderProtocol returns a boolean if a field has been set.

### SetEncoderProtocolNil

`func (o *ModelMediaSourceInfo) SetEncoderProtocolNil(b bool)`

 SetEncoderProtocolNil sets the value for EncoderProtocol to be an explicit nil

### UnsetEncoderProtocol
`func (o *ModelMediaSourceInfo) UnsetEncoderProtocol()`

UnsetEncoderProtocol ensures that no value is present for EncoderProtocol, not even an explicit nil
### GetType

`func (o *ModelMediaSourceInfo) GetType() ModelModelMediaSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelMediaSourceInfo) GetTypeOk() (*ModelModelMediaSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelMediaSourceInfo) SetType(v ModelModelMediaSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelMediaSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContainer

`func (o *ModelMediaSourceInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelMediaSourceInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelMediaSourceInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelMediaSourceInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ModelMediaSourceInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ModelMediaSourceInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSize

`func (o *ModelMediaSourceInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ModelMediaSourceInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ModelMediaSourceInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ModelMediaSourceInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ModelMediaSourceInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ModelMediaSourceInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *ModelMediaSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMediaSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMediaSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelMediaSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelMediaSourceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelMediaSourceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsRemote

`func (o *ModelMediaSourceInfo) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *ModelMediaSourceInfo) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *ModelMediaSourceInfo) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *ModelMediaSourceInfo) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetETag

`func (o *ModelMediaSourceInfo) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *ModelMediaSourceInfo) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *ModelMediaSourceInfo) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *ModelMediaSourceInfo) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *ModelMediaSourceInfo) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *ModelMediaSourceInfo) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetRunTimeTicks

`func (o *ModelMediaSourceInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelMediaSourceInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelMediaSourceInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelMediaSourceInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelMediaSourceInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelMediaSourceInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) GetReadAtNativeFramerate() bool`

GetReadAtNativeFramerate returns the ReadAtNativeFramerate field if non-nil, zero value otherwise.

### GetReadAtNativeFramerateOk

`func (o *ModelMediaSourceInfo) GetReadAtNativeFramerateOk() (*bool, bool)`

GetReadAtNativeFramerateOk returns a tuple with the ReadAtNativeFramerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) SetReadAtNativeFramerate(v bool)`

SetReadAtNativeFramerate sets ReadAtNativeFramerate field to given value.

### HasReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) HasReadAtNativeFramerate() bool`

HasReadAtNativeFramerate returns a boolean if a field has been set.

### GetIgnoreDts

`func (o *ModelMediaSourceInfo) GetIgnoreDts() bool`

GetIgnoreDts returns the IgnoreDts field if non-nil, zero value otherwise.

### GetIgnoreDtsOk

`func (o *ModelMediaSourceInfo) GetIgnoreDtsOk() (*bool, bool)`

GetIgnoreDtsOk returns a tuple with the IgnoreDts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDts

`func (o *ModelMediaSourceInfo) SetIgnoreDts(v bool)`

SetIgnoreDts sets IgnoreDts field to given value.

### HasIgnoreDts

`func (o *ModelMediaSourceInfo) HasIgnoreDts() bool`

HasIgnoreDts returns a boolean if a field has been set.

### GetIgnoreIndex

`func (o *ModelMediaSourceInfo) GetIgnoreIndex() bool`

GetIgnoreIndex returns the IgnoreIndex field if non-nil, zero value otherwise.

### GetIgnoreIndexOk

`func (o *ModelMediaSourceInfo) GetIgnoreIndexOk() (*bool, bool)`

GetIgnoreIndexOk returns a tuple with the IgnoreIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIndex

`func (o *ModelMediaSourceInfo) SetIgnoreIndex(v bool)`

SetIgnoreIndex sets IgnoreIndex field to given value.

### HasIgnoreIndex

`func (o *ModelMediaSourceInfo) HasIgnoreIndex() bool`

HasIgnoreIndex returns a boolean if a field has been set.

### GetGenPtsInput

`func (o *ModelMediaSourceInfo) GetGenPtsInput() bool`

GetGenPtsInput returns the GenPtsInput field if non-nil, zero value otherwise.

### GetGenPtsInputOk

`func (o *ModelMediaSourceInfo) GetGenPtsInputOk() (*bool, bool)`

GetGenPtsInputOk returns a tuple with the GenPtsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenPtsInput

`func (o *ModelMediaSourceInfo) SetGenPtsInput(v bool)`

SetGenPtsInput sets GenPtsInput field to given value.

### HasGenPtsInput

`func (o *ModelMediaSourceInfo) HasGenPtsInput() bool`

HasGenPtsInput returns a boolean if a field has been set.

### GetSupportsTranscoding

`func (o *ModelMediaSourceInfo) GetSupportsTranscoding() bool`

GetSupportsTranscoding returns the SupportsTranscoding field if non-nil, zero value otherwise.

### GetSupportsTranscodingOk

`func (o *ModelMediaSourceInfo) GetSupportsTranscodingOk() (*bool, bool)`

GetSupportsTranscodingOk returns a tuple with the SupportsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTranscoding

`func (o *ModelMediaSourceInfo) SetSupportsTranscoding(v bool)`

SetSupportsTranscoding sets SupportsTranscoding field to given value.

### HasSupportsTranscoding

`func (o *ModelMediaSourceInfo) HasSupportsTranscoding() bool`

HasSupportsTranscoding returns a boolean if a field has been set.

### GetSupportsDirectStream

`func (o *ModelMediaSourceInfo) GetSupportsDirectStream() bool`

GetSupportsDirectStream returns the SupportsDirectStream field if non-nil, zero value otherwise.

### GetSupportsDirectStreamOk

`func (o *ModelMediaSourceInfo) GetSupportsDirectStreamOk() (*bool, bool)`

GetSupportsDirectStreamOk returns a tuple with the SupportsDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectStream

`func (o *ModelMediaSourceInfo) SetSupportsDirectStream(v bool)`

SetSupportsDirectStream sets SupportsDirectStream field to given value.

### HasSupportsDirectStream

`func (o *ModelMediaSourceInfo) HasSupportsDirectStream() bool`

HasSupportsDirectStream returns a boolean if a field has been set.

### GetSupportsDirectPlay

`func (o *ModelMediaSourceInfo) GetSupportsDirectPlay() bool`

GetSupportsDirectPlay returns the SupportsDirectPlay field if non-nil, zero value otherwise.

### GetSupportsDirectPlayOk

`func (o *ModelMediaSourceInfo) GetSupportsDirectPlayOk() (*bool, bool)`

GetSupportsDirectPlayOk returns a tuple with the SupportsDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectPlay

`func (o *ModelMediaSourceInfo) SetSupportsDirectPlay(v bool)`

SetSupportsDirectPlay sets SupportsDirectPlay field to given value.

### HasSupportsDirectPlay

`func (o *ModelMediaSourceInfo) HasSupportsDirectPlay() bool`

HasSupportsDirectPlay returns a boolean if a field has been set.

### GetIsInfiniteStream

`func (o *ModelMediaSourceInfo) GetIsInfiniteStream() bool`

GetIsInfiniteStream returns the IsInfiniteStream field if non-nil, zero value otherwise.

### GetIsInfiniteStreamOk

`func (o *ModelMediaSourceInfo) GetIsInfiniteStreamOk() (*bool, bool)`

GetIsInfiniteStreamOk returns a tuple with the IsInfiniteStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfiniteStream

`func (o *ModelMediaSourceInfo) SetIsInfiniteStream(v bool)`

SetIsInfiniteStream sets IsInfiniteStream field to given value.

### HasIsInfiniteStream

`func (o *ModelMediaSourceInfo) HasIsInfiniteStream() bool`

HasIsInfiniteStream returns a boolean if a field has been set.

### GetUseMostCompatibleTranscodingProfile

`func (o *ModelMediaSourceInfo) GetUseMostCompatibleTranscodingProfile() bool`

GetUseMostCompatibleTranscodingProfile returns the UseMostCompatibleTranscodingProfile field if non-nil, zero value otherwise.

### GetUseMostCompatibleTranscodingProfileOk

`func (o *ModelMediaSourceInfo) GetUseMostCompatibleTranscodingProfileOk() (*bool, bool)`

GetUseMostCompatibleTranscodingProfileOk returns a tuple with the UseMostCompatibleTranscodingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMostCompatibleTranscodingProfile

`func (o *ModelMediaSourceInfo) SetUseMostCompatibleTranscodingProfile(v bool)`

SetUseMostCompatibleTranscodingProfile sets UseMostCompatibleTranscodingProfile field to given value.

### HasUseMostCompatibleTranscodingProfile

`func (o *ModelMediaSourceInfo) HasUseMostCompatibleTranscodingProfile() bool`

HasUseMostCompatibleTranscodingProfile returns a boolean if a field has been set.

### GetRequiresOpening

`func (o *ModelMediaSourceInfo) GetRequiresOpening() bool`

GetRequiresOpening returns the RequiresOpening field if non-nil, zero value otherwise.

### GetRequiresOpeningOk

`func (o *ModelMediaSourceInfo) GetRequiresOpeningOk() (*bool, bool)`

GetRequiresOpeningOk returns a tuple with the RequiresOpening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOpening

`func (o *ModelMediaSourceInfo) SetRequiresOpening(v bool)`

SetRequiresOpening sets RequiresOpening field to given value.

### HasRequiresOpening

`func (o *ModelMediaSourceInfo) HasRequiresOpening() bool`

HasRequiresOpening returns a boolean if a field has been set.

### GetOpenToken

`func (o *ModelMediaSourceInfo) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *ModelMediaSourceInfo) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *ModelMediaSourceInfo) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *ModelMediaSourceInfo) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *ModelMediaSourceInfo) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *ModelMediaSourceInfo) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetRequiresClosing

`func (o *ModelMediaSourceInfo) GetRequiresClosing() bool`

GetRequiresClosing returns the RequiresClosing field if non-nil, zero value otherwise.

### GetRequiresClosingOk

`func (o *ModelMediaSourceInfo) GetRequiresClosingOk() (*bool, bool)`

GetRequiresClosingOk returns a tuple with the RequiresClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClosing

`func (o *ModelMediaSourceInfo) SetRequiresClosing(v bool)`

SetRequiresClosing sets RequiresClosing field to given value.

### HasRequiresClosing

`func (o *ModelMediaSourceInfo) HasRequiresClosing() bool`

HasRequiresClosing returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *ModelMediaSourceInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelMediaSourceInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelMediaSourceInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelMediaSourceInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *ModelMediaSourceInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *ModelMediaSourceInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetBufferMs

`func (o *ModelMediaSourceInfo) GetBufferMs() int32`

GetBufferMs returns the BufferMs field if non-nil, zero value otherwise.

### GetBufferMsOk

`func (o *ModelMediaSourceInfo) GetBufferMsOk() (*int32, bool)`

GetBufferMsOk returns a tuple with the BufferMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMs

`func (o *ModelMediaSourceInfo) SetBufferMs(v int32)`

SetBufferMs sets BufferMs field to given value.

### HasBufferMs

`func (o *ModelMediaSourceInfo) HasBufferMs() bool`

HasBufferMs returns a boolean if a field has been set.

### SetBufferMsNil

`func (o *ModelMediaSourceInfo) SetBufferMsNil(b bool)`

 SetBufferMsNil sets the value for BufferMs to be an explicit nil

### UnsetBufferMs
`func (o *ModelMediaSourceInfo) UnsetBufferMs()`

UnsetBufferMs ensures that no value is present for BufferMs, not even an explicit nil
### GetRequiresLooping

`func (o *ModelMediaSourceInfo) GetRequiresLooping() bool`

GetRequiresLooping returns the RequiresLooping field if non-nil, zero value otherwise.

### GetRequiresLoopingOk

`func (o *ModelMediaSourceInfo) GetRequiresLoopingOk() (*bool, bool)`

GetRequiresLoopingOk returns a tuple with the RequiresLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLooping

`func (o *ModelMediaSourceInfo) SetRequiresLooping(v bool)`

SetRequiresLooping sets RequiresLooping field to given value.

### HasRequiresLooping

`func (o *ModelMediaSourceInfo) HasRequiresLooping() bool`

HasRequiresLooping returns a boolean if a field has been set.

### GetSupportsProbing

`func (o *ModelMediaSourceInfo) GetSupportsProbing() bool`

GetSupportsProbing returns the SupportsProbing field if non-nil, zero value otherwise.

### GetSupportsProbingOk

`func (o *ModelMediaSourceInfo) GetSupportsProbingOk() (*bool, bool)`

GetSupportsProbingOk returns a tuple with the SupportsProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsProbing

`func (o *ModelMediaSourceInfo) SetSupportsProbing(v bool)`

SetSupportsProbing sets SupportsProbing field to given value.

### HasSupportsProbing

`func (o *ModelMediaSourceInfo) HasSupportsProbing() bool`

HasSupportsProbing returns a boolean if a field has been set.

### GetVideoType

`func (o *ModelMediaSourceInfo) GetVideoType() ModelModelVideoType`

GetVideoType returns the VideoType field if non-nil, zero value otherwise.

### GetVideoTypeOk

`func (o *ModelMediaSourceInfo) GetVideoTypeOk() (*ModelModelVideoType, bool)`

GetVideoTypeOk returns a tuple with the VideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoType

`func (o *ModelMediaSourceInfo) SetVideoType(v ModelModelVideoType)`

SetVideoType sets VideoType field to given value.

### HasVideoType

`func (o *ModelMediaSourceInfo) HasVideoType() bool`

HasVideoType returns a boolean if a field has been set.

### SetVideoTypeNil

`func (o *ModelMediaSourceInfo) SetVideoTypeNil(b bool)`

 SetVideoTypeNil sets the value for VideoType to be an explicit nil

### UnsetVideoType
`func (o *ModelMediaSourceInfo) UnsetVideoType()`

UnsetVideoType ensures that no value is present for VideoType, not even an explicit nil
### GetIsoType

`func (o *ModelMediaSourceInfo) GetIsoType() ModelModelIsoType`

GetIsoType returns the IsoType field if non-nil, zero value otherwise.

### GetIsoTypeOk

`func (o *ModelMediaSourceInfo) GetIsoTypeOk() (*ModelModelIsoType, bool)`

GetIsoTypeOk returns a tuple with the IsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoType

`func (o *ModelMediaSourceInfo) SetIsoType(v ModelModelIsoType)`

SetIsoType sets IsoType field to given value.

### HasIsoType

`func (o *ModelMediaSourceInfo) HasIsoType() bool`

HasIsoType returns a boolean if a field has been set.

### SetIsoTypeNil

`func (o *ModelMediaSourceInfo) SetIsoTypeNil(b bool)`

 SetIsoTypeNil sets the value for IsoType to be an explicit nil

### UnsetIsoType
`func (o *ModelMediaSourceInfo) UnsetIsoType()`

UnsetIsoType ensures that no value is present for IsoType, not even an explicit nil
### GetVideo3DFormat

`func (o *ModelMediaSourceInfo) GetVideo3DFormat() ModelModelVideo3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *ModelMediaSourceInfo) GetVideo3DFormatOk() (*ModelModelVideo3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *ModelMediaSourceInfo) SetVideo3DFormat(v ModelModelVideo3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *ModelMediaSourceInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### SetVideo3DFormatNil

`func (o *ModelMediaSourceInfo) SetVideo3DFormatNil(b bool)`

 SetVideo3DFormatNil sets the value for Video3DFormat to be an explicit nil

### UnsetVideo3DFormat
`func (o *ModelMediaSourceInfo) UnsetVideo3DFormat()`

UnsetVideo3DFormat ensures that no value is present for Video3DFormat, not even an explicit nil
### GetMediaStreams

`func (o *ModelMediaSourceInfo) GetMediaStreams() []ModelModelMediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *ModelMediaSourceInfo) GetMediaStreamsOk() (*[]ModelModelMediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *ModelMediaSourceInfo) SetMediaStreams(v []ModelModelMediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *ModelMediaSourceInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### SetMediaStreamsNil

`func (o *ModelMediaSourceInfo) SetMediaStreamsNil(b bool)`

 SetMediaStreamsNil sets the value for MediaStreams to be an explicit nil

### UnsetMediaStreams
`func (o *ModelMediaSourceInfo) UnsetMediaStreams()`

UnsetMediaStreams ensures that no value is present for MediaStreams, not even an explicit nil
### GetMediaAttachments

`func (o *ModelMediaSourceInfo) GetMediaAttachments() []ModelModelMediaAttachment`

GetMediaAttachments returns the MediaAttachments field if non-nil, zero value otherwise.

### GetMediaAttachmentsOk

`func (o *ModelMediaSourceInfo) GetMediaAttachmentsOk() (*[]ModelModelMediaAttachment, bool)`

GetMediaAttachmentsOk returns a tuple with the MediaAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAttachments

`func (o *ModelMediaSourceInfo) SetMediaAttachments(v []ModelModelMediaAttachment)`

SetMediaAttachments sets MediaAttachments field to given value.

### HasMediaAttachments

`func (o *ModelMediaSourceInfo) HasMediaAttachments() bool`

HasMediaAttachments returns a boolean if a field has been set.

### SetMediaAttachmentsNil

`func (o *ModelMediaSourceInfo) SetMediaAttachmentsNil(b bool)`

 SetMediaAttachmentsNil sets the value for MediaAttachments to be an explicit nil

### UnsetMediaAttachments
`func (o *ModelMediaSourceInfo) UnsetMediaAttachments()`

UnsetMediaAttachments ensures that no value is present for MediaAttachments, not even an explicit nil
### GetFormats

`func (o *ModelMediaSourceInfo) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *ModelMediaSourceInfo) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *ModelMediaSourceInfo) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *ModelMediaSourceInfo) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### SetFormatsNil

`func (o *ModelMediaSourceInfo) SetFormatsNil(b bool)`

 SetFormatsNil sets the value for Formats to be an explicit nil

### UnsetFormats
`func (o *ModelMediaSourceInfo) UnsetFormats()`

UnsetFormats ensures that no value is present for Formats, not even an explicit nil
### GetBitrate

`func (o *ModelMediaSourceInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelMediaSourceInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelMediaSourceInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelMediaSourceInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelMediaSourceInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelMediaSourceInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetFallbackMaxStreamingBitrate

`func (o *ModelMediaSourceInfo) GetFallbackMaxStreamingBitrate() int32`

GetFallbackMaxStreamingBitrate returns the FallbackMaxStreamingBitrate field if non-nil, zero value otherwise.

### GetFallbackMaxStreamingBitrateOk

`func (o *ModelMediaSourceInfo) GetFallbackMaxStreamingBitrateOk() (*int32, bool)`

GetFallbackMaxStreamingBitrateOk returns a tuple with the FallbackMaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMaxStreamingBitrate

`func (o *ModelMediaSourceInfo) SetFallbackMaxStreamingBitrate(v int32)`

SetFallbackMaxStreamingBitrate sets FallbackMaxStreamingBitrate field to given value.

### HasFallbackMaxStreamingBitrate

`func (o *ModelMediaSourceInfo) HasFallbackMaxStreamingBitrate() bool`

HasFallbackMaxStreamingBitrate returns a boolean if a field has been set.

### SetFallbackMaxStreamingBitrateNil

`func (o *ModelMediaSourceInfo) SetFallbackMaxStreamingBitrateNil(b bool)`

 SetFallbackMaxStreamingBitrateNil sets the value for FallbackMaxStreamingBitrate to be an explicit nil

### UnsetFallbackMaxStreamingBitrate
`func (o *ModelMediaSourceInfo) UnsetFallbackMaxStreamingBitrate()`

UnsetFallbackMaxStreamingBitrate ensures that no value is present for FallbackMaxStreamingBitrate, not even an explicit nil
### GetTimestamp

`func (o *ModelMediaSourceInfo) GetTimestamp() ModelModelTransportStreamTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelMediaSourceInfo) GetTimestampOk() (*ModelModelTransportStreamTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelMediaSourceInfo) SetTimestamp(v ModelModelTransportStreamTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelMediaSourceInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ModelMediaSourceInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ModelMediaSourceInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) GetRequiredHttpHeaders() map[string]string`

GetRequiredHttpHeaders returns the RequiredHttpHeaders field if non-nil, zero value otherwise.

### GetRequiredHttpHeadersOk

`func (o *ModelMediaSourceInfo) GetRequiredHttpHeadersOk() (*map[string]string, bool)`

GetRequiredHttpHeadersOk returns a tuple with the RequiredHttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) SetRequiredHttpHeaders(v map[string]string)`

SetRequiredHttpHeaders sets RequiredHttpHeaders field to given value.

### HasRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) HasRequiredHttpHeaders() bool`

HasRequiredHttpHeaders returns a boolean if a field has been set.

### SetRequiredHttpHeadersNil

`func (o *ModelMediaSourceInfo) SetRequiredHttpHeadersNil(b bool)`

 SetRequiredHttpHeadersNil sets the value for RequiredHttpHeaders to be an explicit nil

### UnsetRequiredHttpHeaders
`func (o *ModelMediaSourceInfo) UnsetRequiredHttpHeaders()`

UnsetRequiredHttpHeaders ensures that no value is present for RequiredHttpHeaders, not even an explicit nil
### GetTranscodingUrl

`func (o *ModelMediaSourceInfo) GetTranscodingUrl() string`

GetTranscodingUrl returns the TranscodingUrl field if non-nil, zero value otherwise.

### GetTranscodingUrlOk

`func (o *ModelMediaSourceInfo) GetTranscodingUrlOk() (*string, bool)`

GetTranscodingUrlOk returns a tuple with the TranscodingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUrl

`func (o *ModelMediaSourceInfo) SetTranscodingUrl(v string)`

SetTranscodingUrl sets TranscodingUrl field to given value.

### HasTranscodingUrl

`func (o *ModelMediaSourceInfo) HasTranscodingUrl() bool`

HasTranscodingUrl returns a boolean if a field has been set.

### SetTranscodingUrlNil

`func (o *ModelMediaSourceInfo) SetTranscodingUrlNil(b bool)`

 SetTranscodingUrlNil sets the value for TranscodingUrl to be an explicit nil

### UnsetTranscodingUrl
`func (o *ModelMediaSourceInfo) UnsetTranscodingUrl()`

UnsetTranscodingUrl ensures that no value is present for TranscodingUrl, not even an explicit nil
### GetTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) GetTranscodingSubProtocol() ModelModelMediaStreamProtocol`

GetTranscodingSubProtocol returns the TranscodingSubProtocol field if non-nil, zero value otherwise.

### GetTranscodingSubProtocolOk

`func (o *ModelMediaSourceInfo) GetTranscodingSubProtocolOk() (*ModelModelMediaStreamProtocol, bool)`

GetTranscodingSubProtocolOk returns a tuple with the TranscodingSubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) SetTranscodingSubProtocol(v ModelModelMediaStreamProtocol)`

SetTranscodingSubProtocol sets TranscodingSubProtocol field to given value.

### HasTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) HasTranscodingSubProtocol() bool`

HasTranscodingSubProtocol returns a boolean if a field has been set.

### GetTranscodingContainer

`func (o *ModelMediaSourceInfo) GetTranscodingContainer() string`

GetTranscodingContainer returns the TranscodingContainer field if non-nil, zero value otherwise.

### GetTranscodingContainerOk

`func (o *ModelMediaSourceInfo) GetTranscodingContainerOk() (*string, bool)`

GetTranscodingContainerOk returns a tuple with the TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingContainer

`func (o *ModelMediaSourceInfo) SetTranscodingContainer(v string)`

SetTranscodingContainer sets TranscodingContainer field to given value.

### HasTranscodingContainer

`func (o *ModelMediaSourceInfo) HasTranscodingContainer() bool`

HasTranscodingContainer returns a boolean if a field has been set.

### SetTranscodingContainerNil

`func (o *ModelMediaSourceInfo) SetTranscodingContainerNil(b bool)`

 SetTranscodingContainerNil sets the value for TranscodingContainer to be an explicit nil

### UnsetTranscodingContainer
`func (o *ModelMediaSourceInfo) UnsetTranscodingContainer()`

UnsetTranscodingContainer ensures that no value is present for TranscodingContainer, not even an explicit nil
### GetAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) GetAnalyzeDurationMs() int32`

GetAnalyzeDurationMs returns the AnalyzeDurationMs field if non-nil, zero value otherwise.

### GetAnalyzeDurationMsOk

`func (o *ModelMediaSourceInfo) GetAnalyzeDurationMsOk() (*int32, bool)`

GetAnalyzeDurationMsOk returns a tuple with the AnalyzeDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) SetAnalyzeDurationMs(v int32)`

SetAnalyzeDurationMs sets AnalyzeDurationMs field to given value.

### HasAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) HasAnalyzeDurationMs() bool`

HasAnalyzeDurationMs returns a boolean if a field has been set.

### SetAnalyzeDurationMsNil

`func (o *ModelMediaSourceInfo) SetAnalyzeDurationMsNil(b bool)`

 SetAnalyzeDurationMsNil sets the value for AnalyzeDurationMs to be an explicit nil

### UnsetAnalyzeDurationMs
`func (o *ModelMediaSourceInfo) UnsetAnalyzeDurationMs()`

UnsetAnalyzeDurationMs ensures that no value is present for AnalyzeDurationMs, not even an explicit nil
### GetDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) GetDefaultAudioStreamIndex() int32`

GetDefaultAudioStreamIndex returns the DefaultAudioStreamIndex field if non-nil, zero value otherwise.

### GetDefaultAudioStreamIndexOk

`func (o *ModelMediaSourceInfo) GetDefaultAudioStreamIndexOk() (*int32, bool)`

GetDefaultAudioStreamIndexOk returns a tuple with the DefaultAudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) SetDefaultAudioStreamIndex(v int32)`

SetDefaultAudioStreamIndex sets DefaultAudioStreamIndex field to given value.

### HasDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) HasDefaultAudioStreamIndex() bool`

HasDefaultAudioStreamIndex returns a boolean if a field has been set.

### SetDefaultAudioStreamIndexNil

`func (o *ModelMediaSourceInfo) SetDefaultAudioStreamIndexNil(b bool)`

 SetDefaultAudioStreamIndexNil sets the value for DefaultAudioStreamIndex to be an explicit nil

### UnsetDefaultAudioStreamIndex
`func (o *ModelMediaSourceInfo) UnsetDefaultAudioStreamIndex()`

UnsetDefaultAudioStreamIndex ensures that no value is present for DefaultAudioStreamIndex, not even an explicit nil
### GetDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) GetDefaultSubtitleStreamIndex() int32`

GetDefaultSubtitleStreamIndex returns the DefaultSubtitleStreamIndex field if non-nil, zero value otherwise.

### GetDefaultSubtitleStreamIndexOk

`func (o *ModelMediaSourceInfo) GetDefaultSubtitleStreamIndexOk() (*int32, bool)`

GetDefaultSubtitleStreamIndexOk returns a tuple with the DefaultSubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) SetDefaultSubtitleStreamIndex(v int32)`

SetDefaultSubtitleStreamIndex sets DefaultSubtitleStreamIndex field to given value.

### HasDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) HasDefaultSubtitleStreamIndex() bool`

HasDefaultSubtitleStreamIndex returns a boolean if a field has been set.

### SetDefaultSubtitleStreamIndexNil

`func (o *ModelMediaSourceInfo) SetDefaultSubtitleStreamIndexNil(b bool)`

 SetDefaultSubtitleStreamIndexNil sets the value for DefaultSubtitleStreamIndex to be an explicit nil

### UnsetDefaultSubtitleStreamIndex
`func (o *ModelMediaSourceInfo) UnsetDefaultSubtitleStreamIndex()`

UnsetDefaultSubtitleStreamIndex ensures that no value is present for DefaultSubtitleStreamIndex, not even an explicit nil
### GetHasSegments

`func (o *ModelMediaSourceInfo) GetHasSegments() bool`

GetHasSegments returns the HasSegments field if non-nil, zero value otherwise.

### GetHasSegmentsOk

`func (o *ModelMediaSourceInfo) GetHasSegmentsOk() (*bool, bool)`

GetHasSegmentsOk returns a tuple with the HasSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSegments

`func (o *ModelMediaSourceInfo) SetHasSegments(v bool)`

SetHasSegments sets HasSegments field to given value.

### HasHasSegments

`func (o *ModelMediaSourceInfo) HasHasSegments() bool`

HasHasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


