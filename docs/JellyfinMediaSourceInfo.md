# JellyfinMediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**JellyfinJellyfinMediaProtocol**](JellyfinMediaProtocol.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**EncoderPath** | Pointer to **NullableString** |  | [optional] 
**EncoderProtocol** | Pointer to [**NullableJellyfinJellyfinMediaProtocol**](JellyfinMediaProtocol.md) |  | [optional] 
**Type** | Pointer to [**JellyfinJellyfinMediaSourceType**](JellyfinMediaSourceType.md) |  | [optional] 
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
**VideoType** | Pointer to [**NullableJellyfinJellyfinVideoType**](JellyfinVideoType.md) |  | [optional] 
**IsoType** | Pointer to [**NullableJellyfinJellyfinIsoType**](JellyfinIsoType.md) |  | [optional] 
**Video3DFormat** | Pointer to [**NullableJellyfinJellyfinVideo3DFormat**](JellyfinVideo3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]JellyfinJellyfinMediaStream**](JellyfinJellyfinMediaStream.md) |  | [optional] 
**MediaAttachments** | Pointer to [**[]JellyfinJellyfinMediaAttachment**](JellyfinJellyfinMediaAttachment.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**FallbackMaxStreamingBitrate** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**NullableJellyfinJellyfinTransportStreamTimestamp**](JellyfinTransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**TranscodingUrl** | Pointer to **NullableString** |  | [optional] 
**TranscodingSubProtocol** | Pointer to [**JellyfinJellyfinMediaStreamProtocol**](JellyfinMediaStreamProtocol.md) | Media streaming protocol.  Lowercase for backwards compatibility. | [optional] 
**TranscodingContainer** | Pointer to **NullableString** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **NullableInt32** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**HasSegments** | Pointer to **bool** |  | [optional] 

## Methods

### NewJellyfinMediaSourceInfo

`func NewJellyfinMediaSourceInfo() *JellyfinMediaSourceInfo`

NewJellyfinMediaSourceInfo instantiates a new JellyfinMediaSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinMediaSourceInfoWithDefaults

`func NewJellyfinMediaSourceInfoWithDefaults() *JellyfinMediaSourceInfo`

NewJellyfinMediaSourceInfoWithDefaults instantiates a new JellyfinMediaSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *JellyfinMediaSourceInfo) GetProtocol() JellyfinJellyfinMediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *JellyfinMediaSourceInfo) GetProtocolOk() (*JellyfinJellyfinMediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *JellyfinMediaSourceInfo) SetProtocol(v JellyfinJellyfinMediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *JellyfinMediaSourceInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetId

`func (o *JellyfinMediaSourceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinMediaSourceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinMediaSourceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinMediaSourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinMediaSourceInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinMediaSourceInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *JellyfinMediaSourceInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JellyfinMediaSourceInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JellyfinMediaSourceInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *JellyfinMediaSourceInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *JellyfinMediaSourceInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *JellyfinMediaSourceInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEncoderPath

`func (o *JellyfinMediaSourceInfo) GetEncoderPath() string`

GetEncoderPath returns the EncoderPath field if non-nil, zero value otherwise.

### GetEncoderPathOk

`func (o *JellyfinMediaSourceInfo) GetEncoderPathOk() (*string, bool)`

GetEncoderPathOk returns a tuple with the EncoderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPath

`func (o *JellyfinMediaSourceInfo) SetEncoderPath(v string)`

SetEncoderPath sets EncoderPath field to given value.

### HasEncoderPath

`func (o *JellyfinMediaSourceInfo) HasEncoderPath() bool`

HasEncoderPath returns a boolean if a field has been set.

### SetEncoderPathNil

`func (o *JellyfinMediaSourceInfo) SetEncoderPathNil(b bool)`

 SetEncoderPathNil sets the value for EncoderPath to be an explicit nil

### UnsetEncoderPath
`func (o *JellyfinMediaSourceInfo) UnsetEncoderPath()`

UnsetEncoderPath ensures that no value is present for EncoderPath, not even an explicit nil
### GetEncoderProtocol

`func (o *JellyfinMediaSourceInfo) GetEncoderProtocol() JellyfinJellyfinMediaProtocol`

GetEncoderProtocol returns the EncoderProtocol field if non-nil, zero value otherwise.

### GetEncoderProtocolOk

`func (o *JellyfinMediaSourceInfo) GetEncoderProtocolOk() (*JellyfinJellyfinMediaProtocol, bool)`

GetEncoderProtocolOk returns a tuple with the EncoderProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderProtocol

`func (o *JellyfinMediaSourceInfo) SetEncoderProtocol(v JellyfinJellyfinMediaProtocol)`

SetEncoderProtocol sets EncoderProtocol field to given value.

### HasEncoderProtocol

`func (o *JellyfinMediaSourceInfo) HasEncoderProtocol() bool`

HasEncoderProtocol returns a boolean if a field has been set.

### SetEncoderProtocolNil

`func (o *JellyfinMediaSourceInfo) SetEncoderProtocolNil(b bool)`

 SetEncoderProtocolNil sets the value for EncoderProtocol to be an explicit nil

### UnsetEncoderProtocol
`func (o *JellyfinMediaSourceInfo) UnsetEncoderProtocol()`

UnsetEncoderProtocol ensures that no value is present for EncoderProtocol, not even an explicit nil
### GetType

`func (o *JellyfinMediaSourceInfo) GetType() JellyfinJellyfinMediaSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinMediaSourceInfo) GetTypeOk() (*JellyfinJellyfinMediaSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinMediaSourceInfo) SetType(v JellyfinJellyfinMediaSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinMediaSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContainer

`func (o *JellyfinMediaSourceInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *JellyfinMediaSourceInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *JellyfinMediaSourceInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *JellyfinMediaSourceInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *JellyfinMediaSourceInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *JellyfinMediaSourceInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSize

`func (o *JellyfinMediaSourceInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *JellyfinMediaSourceInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *JellyfinMediaSourceInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *JellyfinMediaSourceInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *JellyfinMediaSourceInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *JellyfinMediaSourceInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *JellyfinMediaSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinMediaSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinMediaSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinMediaSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinMediaSourceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinMediaSourceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsRemote

`func (o *JellyfinMediaSourceInfo) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *JellyfinMediaSourceInfo) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *JellyfinMediaSourceInfo) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *JellyfinMediaSourceInfo) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetETag

`func (o *JellyfinMediaSourceInfo) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *JellyfinMediaSourceInfo) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *JellyfinMediaSourceInfo) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *JellyfinMediaSourceInfo) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *JellyfinMediaSourceInfo) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *JellyfinMediaSourceInfo) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetRunTimeTicks

`func (o *JellyfinMediaSourceInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *JellyfinMediaSourceInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *JellyfinMediaSourceInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *JellyfinMediaSourceInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *JellyfinMediaSourceInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *JellyfinMediaSourceInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetReadAtNativeFramerate

`func (o *JellyfinMediaSourceInfo) GetReadAtNativeFramerate() bool`

GetReadAtNativeFramerate returns the ReadAtNativeFramerate field if non-nil, zero value otherwise.

### GetReadAtNativeFramerateOk

`func (o *JellyfinMediaSourceInfo) GetReadAtNativeFramerateOk() (*bool, bool)`

GetReadAtNativeFramerateOk returns a tuple with the ReadAtNativeFramerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAtNativeFramerate

`func (o *JellyfinMediaSourceInfo) SetReadAtNativeFramerate(v bool)`

SetReadAtNativeFramerate sets ReadAtNativeFramerate field to given value.

### HasReadAtNativeFramerate

`func (o *JellyfinMediaSourceInfo) HasReadAtNativeFramerate() bool`

HasReadAtNativeFramerate returns a boolean if a field has been set.

### GetIgnoreDts

`func (o *JellyfinMediaSourceInfo) GetIgnoreDts() bool`

GetIgnoreDts returns the IgnoreDts field if non-nil, zero value otherwise.

### GetIgnoreDtsOk

`func (o *JellyfinMediaSourceInfo) GetIgnoreDtsOk() (*bool, bool)`

GetIgnoreDtsOk returns a tuple with the IgnoreDts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDts

`func (o *JellyfinMediaSourceInfo) SetIgnoreDts(v bool)`

SetIgnoreDts sets IgnoreDts field to given value.

### HasIgnoreDts

`func (o *JellyfinMediaSourceInfo) HasIgnoreDts() bool`

HasIgnoreDts returns a boolean if a field has been set.

### GetIgnoreIndex

`func (o *JellyfinMediaSourceInfo) GetIgnoreIndex() bool`

GetIgnoreIndex returns the IgnoreIndex field if non-nil, zero value otherwise.

### GetIgnoreIndexOk

`func (o *JellyfinMediaSourceInfo) GetIgnoreIndexOk() (*bool, bool)`

GetIgnoreIndexOk returns a tuple with the IgnoreIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIndex

`func (o *JellyfinMediaSourceInfo) SetIgnoreIndex(v bool)`

SetIgnoreIndex sets IgnoreIndex field to given value.

### HasIgnoreIndex

`func (o *JellyfinMediaSourceInfo) HasIgnoreIndex() bool`

HasIgnoreIndex returns a boolean if a field has been set.

### GetGenPtsInput

`func (o *JellyfinMediaSourceInfo) GetGenPtsInput() bool`

GetGenPtsInput returns the GenPtsInput field if non-nil, zero value otherwise.

### GetGenPtsInputOk

`func (o *JellyfinMediaSourceInfo) GetGenPtsInputOk() (*bool, bool)`

GetGenPtsInputOk returns a tuple with the GenPtsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenPtsInput

`func (o *JellyfinMediaSourceInfo) SetGenPtsInput(v bool)`

SetGenPtsInput sets GenPtsInput field to given value.

### HasGenPtsInput

`func (o *JellyfinMediaSourceInfo) HasGenPtsInput() bool`

HasGenPtsInput returns a boolean if a field has been set.

### GetSupportsTranscoding

`func (o *JellyfinMediaSourceInfo) GetSupportsTranscoding() bool`

GetSupportsTranscoding returns the SupportsTranscoding field if non-nil, zero value otherwise.

### GetSupportsTranscodingOk

`func (o *JellyfinMediaSourceInfo) GetSupportsTranscodingOk() (*bool, bool)`

GetSupportsTranscodingOk returns a tuple with the SupportsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTranscoding

`func (o *JellyfinMediaSourceInfo) SetSupportsTranscoding(v bool)`

SetSupportsTranscoding sets SupportsTranscoding field to given value.

### HasSupportsTranscoding

`func (o *JellyfinMediaSourceInfo) HasSupportsTranscoding() bool`

HasSupportsTranscoding returns a boolean if a field has been set.

### GetSupportsDirectStream

`func (o *JellyfinMediaSourceInfo) GetSupportsDirectStream() bool`

GetSupportsDirectStream returns the SupportsDirectStream field if non-nil, zero value otherwise.

### GetSupportsDirectStreamOk

`func (o *JellyfinMediaSourceInfo) GetSupportsDirectStreamOk() (*bool, bool)`

GetSupportsDirectStreamOk returns a tuple with the SupportsDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectStream

`func (o *JellyfinMediaSourceInfo) SetSupportsDirectStream(v bool)`

SetSupportsDirectStream sets SupportsDirectStream field to given value.

### HasSupportsDirectStream

`func (o *JellyfinMediaSourceInfo) HasSupportsDirectStream() bool`

HasSupportsDirectStream returns a boolean if a field has been set.

### GetSupportsDirectPlay

`func (o *JellyfinMediaSourceInfo) GetSupportsDirectPlay() bool`

GetSupportsDirectPlay returns the SupportsDirectPlay field if non-nil, zero value otherwise.

### GetSupportsDirectPlayOk

`func (o *JellyfinMediaSourceInfo) GetSupportsDirectPlayOk() (*bool, bool)`

GetSupportsDirectPlayOk returns a tuple with the SupportsDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectPlay

`func (o *JellyfinMediaSourceInfo) SetSupportsDirectPlay(v bool)`

SetSupportsDirectPlay sets SupportsDirectPlay field to given value.

### HasSupportsDirectPlay

`func (o *JellyfinMediaSourceInfo) HasSupportsDirectPlay() bool`

HasSupportsDirectPlay returns a boolean if a field has been set.

### GetIsInfiniteStream

`func (o *JellyfinMediaSourceInfo) GetIsInfiniteStream() bool`

GetIsInfiniteStream returns the IsInfiniteStream field if non-nil, zero value otherwise.

### GetIsInfiniteStreamOk

`func (o *JellyfinMediaSourceInfo) GetIsInfiniteStreamOk() (*bool, bool)`

GetIsInfiniteStreamOk returns a tuple with the IsInfiniteStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfiniteStream

`func (o *JellyfinMediaSourceInfo) SetIsInfiniteStream(v bool)`

SetIsInfiniteStream sets IsInfiniteStream field to given value.

### HasIsInfiniteStream

`func (o *JellyfinMediaSourceInfo) HasIsInfiniteStream() bool`

HasIsInfiniteStream returns a boolean if a field has been set.

### GetUseMostCompatibleTranscodingProfile

`func (o *JellyfinMediaSourceInfo) GetUseMostCompatibleTranscodingProfile() bool`

GetUseMostCompatibleTranscodingProfile returns the UseMostCompatibleTranscodingProfile field if non-nil, zero value otherwise.

### GetUseMostCompatibleTranscodingProfileOk

`func (o *JellyfinMediaSourceInfo) GetUseMostCompatibleTranscodingProfileOk() (*bool, bool)`

GetUseMostCompatibleTranscodingProfileOk returns a tuple with the UseMostCompatibleTranscodingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMostCompatibleTranscodingProfile

`func (o *JellyfinMediaSourceInfo) SetUseMostCompatibleTranscodingProfile(v bool)`

SetUseMostCompatibleTranscodingProfile sets UseMostCompatibleTranscodingProfile field to given value.

### HasUseMostCompatibleTranscodingProfile

`func (o *JellyfinMediaSourceInfo) HasUseMostCompatibleTranscodingProfile() bool`

HasUseMostCompatibleTranscodingProfile returns a boolean if a field has been set.

### GetRequiresOpening

`func (o *JellyfinMediaSourceInfo) GetRequiresOpening() bool`

GetRequiresOpening returns the RequiresOpening field if non-nil, zero value otherwise.

### GetRequiresOpeningOk

`func (o *JellyfinMediaSourceInfo) GetRequiresOpeningOk() (*bool, bool)`

GetRequiresOpeningOk returns a tuple with the RequiresOpening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOpening

`func (o *JellyfinMediaSourceInfo) SetRequiresOpening(v bool)`

SetRequiresOpening sets RequiresOpening field to given value.

### HasRequiresOpening

`func (o *JellyfinMediaSourceInfo) HasRequiresOpening() bool`

HasRequiresOpening returns a boolean if a field has been set.

### GetOpenToken

`func (o *JellyfinMediaSourceInfo) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *JellyfinMediaSourceInfo) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *JellyfinMediaSourceInfo) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *JellyfinMediaSourceInfo) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *JellyfinMediaSourceInfo) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *JellyfinMediaSourceInfo) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetRequiresClosing

`func (o *JellyfinMediaSourceInfo) GetRequiresClosing() bool`

GetRequiresClosing returns the RequiresClosing field if non-nil, zero value otherwise.

### GetRequiresClosingOk

`func (o *JellyfinMediaSourceInfo) GetRequiresClosingOk() (*bool, bool)`

GetRequiresClosingOk returns a tuple with the RequiresClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClosing

`func (o *JellyfinMediaSourceInfo) SetRequiresClosing(v bool)`

SetRequiresClosing sets RequiresClosing field to given value.

### HasRequiresClosing

`func (o *JellyfinMediaSourceInfo) HasRequiresClosing() bool`

HasRequiresClosing returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *JellyfinMediaSourceInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *JellyfinMediaSourceInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *JellyfinMediaSourceInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *JellyfinMediaSourceInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *JellyfinMediaSourceInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *JellyfinMediaSourceInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetBufferMs

`func (o *JellyfinMediaSourceInfo) GetBufferMs() int32`

GetBufferMs returns the BufferMs field if non-nil, zero value otherwise.

### GetBufferMsOk

`func (o *JellyfinMediaSourceInfo) GetBufferMsOk() (*int32, bool)`

GetBufferMsOk returns a tuple with the BufferMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMs

`func (o *JellyfinMediaSourceInfo) SetBufferMs(v int32)`

SetBufferMs sets BufferMs field to given value.

### HasBufferMs

`func (o *JellyfinMediaSourceInfo) HasBufferMs() bool`

HasBufferMs returns a boolean if a field has been set.

### SetBufferMsNil

`func (o *JellyfinMediaSourceInfo) SetBufferMsNil(b bool)`

 SetBufferMsNil sets the value for BufferMs to be an explicit nil

### UnsetBufferMs
`func (o *JellyfinMediaSourceInfo) UnsetBufferMs()`

UnsetBufferMs ensures that no value is present for BufferMs, not even an explicit nil
### GetRequiresLooping

`func (o *JellyfinMediaSourceInfo) GetRequiresLooping() bool`

GetRequiresLooping returns the RequiresLooping field if non-nil, zero value otherwise.

### GetRequiresLoopingOk

`func (o *JellyfinMediaSourceInfo) GetRequiresLoopingOk() (*bool, bool)`

GetRequiresLoopingOk returns a tuple with the RequiresLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLooping

`func (o *JellyfinMediaSourceInfo) SetRequiresLooping(v bool)`

SetRequiresLooping sets RequiresLooping field to given value.

### HasRequiresLooping

`func (o *JellyfinMediaSourceInfo) HasRequiresLooping() bool`

HasRequiresLooping returns a boolean if a field has been set.

### GetSupportsProbing

`func (o *JellyfinMediaSourceInfo) GetSupportsProbing() bool`

GetSupportsProbing returns the SupportsProbing field if non-nil, zero value otherwise.

### GetSupportsProbingOk

`func (o *JellyfinMediaSourceInfo) GetSupportsProbingOk() (*bool, bool)`

GetSupportsProbingOk returns a tuple with the SupportsProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsProbing

`func (o *JellyfinMediaSourceInfo) SetSupportsProbing(v bool)`

SetSupportsProbing sets SupportsProbing field to given value.

### HasSupportsProbing

`func (o *JellyfinMediaSourceInfo) HasSupportsProbing() bool`

HasSupportsProbing returns a boolean if a field has been set.

### GetVideoType

`func (o *JellyfinMediaSourceInfo) GetVideoType() JellyfinJellyfinVideoType`

GetVideoType returns the VideoType field if non-nil, zero value otherwise.

### GetVideoTypeOk

`func (o *JellyfinMediaSourceInfo) GetVideoTypeOk() (*JellyfinJellyfinVideoType, bool)`

GetVideoTypeOk returns a tuple with the VideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoType

`func (o *JellyfinMediaSourceInfo) SetVideoType(v JellyfinJellyfinVideoType)`

SetVideoType sets VideoType field to given value.

### HasVideoType

`func (o *JellyfinMediaSourceInfo) HasVideoType() bool`

HasVideoType returns a boolean if a field has been set.

### SetVideoTypeNil

`func (o *JellyfinMediaSourceInfo) SetVideoTypeNil(b bool)`

 SetVideoTypeNil sets the value for VideoType to be an explicit nil

### UnsetVideoType
`func (o *JellyfinMediaSourceInfo) UnsetVideoType()`

UnsetVideoType ensures that no value is present for VideoType, not even an explicit nil
### GetIsoType

`func (o *JellyfinMediaSourceInfo) GetIsoType() JellyfinJellyfinIsoType`

GetIsoType returns the IsoType field if non-nil, zero value otherwise.

### GetIsoTypeOk

`func (o *JellyfinMediaSourceInfo) GetIsoTypeOk() (*JellyfinJellyfinIsoType, bool)`

GetIsoTypeOk returns a tuple with the IsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoType

`func (o *JellyfinMediaSourceInfo) SetIsoType(v JellyfinJellyfinIsoType)`

SetIsoType sets IsoType field to given value.

### HasIsoType

`func (o *JellyfinMediaSourceInfo) HasIsoType() bool`

HasIsoType returns a boolean if a field has been set.

### SetIsoTypeNil

`func (o *JellyfinMediaSourceInfo) SetIsoTypeNil(b bool)`

 SetIsoTypeNil sets the value for IsoType to be an explicit nil

### UnsetIsoType
`func (o *JellyfinMediaSourceInfo) UnsetIsoType()`

UnsetIsoType ensures that no value is present for IsoType, not even an explicit nil
### GetVideo3DFormat

`func (o *JellyfinMediaSourceInfo) GetVideo3DFormat() JellyfinJellyfinVideo3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *JellyfinMediaSourceInfo) GetVideo3DFormatOk() (*JellyfinJellyfinVideo3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *JellyfinMediaSourceInfo) SetVideo3DFormat(v JellyfinJellyfinVideo3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *JellyfinMediaSourceInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### SetVideo3DFormatNil

`func (o *JellyfinMediaSourceInfo) SetVideo3DFormatNil(b bool)`

 SetVideo3DFormatNil sets the value for Video3DFormat to be an explicit nil

### UnsetVideo3DFormat
`func (o *JellyfinMediaSourceInfo) UnsetVideo3DFormat()`

UnsetVideo3DFormat ensures that no value is present for Video3DFormat, not even an explicit nil
### GetMediaStreams

`func (o *JellyfinMediaSourceInfo) GetMediaStreams() []JellyfinJellyfinMediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *JellyfinMediaSourceInfo) GetMediaStreamsOk() (*[]JellyfinJellyfinMediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *JellyfinMediaSourceInfo) SetMediaStreams(v []JellyfinJellyfinMediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *JellyfinMediaSourceInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### SetMediaStreamsNil

`func (o *JellyfinMediaSourceInfo) SetMediaStreamsNil(b bool)`

 SetMediaStreamsNil sets the value for MediaStreams to be an explicit nil

### UnsetMediaStreams
`func (o *JellyfinMediaSourceInfo) UnsetMediaStreams()`

UnsetMediaStreams ensures that no value is present for MediaStreams, not even an explicit nil
### GetMediaAttachments

`func (o *JellyfinMediaSourceInfo) GetMediaAttachments() []JellyfinJellyfinMediaAttachment`

GetMediaAttachments returns the MediaAttachments field if non-nil, zero value otherwise.

### GetMediaAttachmentsOk

`func (o *JellyfinMediaSourceInfo) GetMediaAttachmentsOk() (*[]JellyfinJellyfinMediaAttachment, bool)`

GetMediaAttachmentsOk returns a tuple with the MediaAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAttachments

`func (o *JellyfinMediaSourceInfo) SetMediaAttachments(v []JellyfinJellyfinMediaAttachment)`

SetMediaAttachments sets MediaAttachments field to given value.

### HasMediaAttachments

`func (o *JellyfinMediaSourceInfo) HasMediaAttachments() bool`

HasMediaAttachments returns a boolean if a field has been set.

### SetMediaAttachmentsNil

`func (o *JellyfinMediaSourceInfo) SetMediaAttachmentsNil(b bool)`

 SetMediaAttachmentsNil sets the value for MediaAttachments to be an explicit nil

### UnsetMediaAttachments
`func (o *JellyfinMediaSourceInfo) UnsetMediaAttachments()`

UnsetMediaAttachments ensures that no value is present for MediaAttachments, not even an explicit nil
### GetFormats

`func (o *JellyfinMediaSourceInfo) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *JellyfinMediaSourceInfo) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *JellyfinMediaSourceInfo) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *JellyfinMediaSourceInfo) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### SetFormatsNil

`func (o *JellyfinMediaSourceInfo) SetFormatsNil(b bool)`

 SetFormatsNil sets the value for Formats to be an explicit nil

### UnsetFormats
`func (o *JellyfinMediaSourceInfo) UnsetFormats()`

UnsetFormats ensures that no value is present for Formats, not even an explicit nil
### GetBitrate

`func (o *JellyfinMediaSourceInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *JellyfinMediaSourceInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *JellyfinMediaSourceInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *JellyfinMediaSourceInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *JellyfinMediaSourceInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *JellyfinMediaSourceInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetFallbackMaxStreamingBitrate

`func (o *JellyfinMediaSourceInfo) GetFallbackMaxStreamingBitrate() int32`

GetFallbackMaxStreamingBitrate returns the FallbackMaxStreamingBitrate field if non-nil, zero value otherwise.

### GetFallbackMaxStreamingBitrateOk

`func (o *JellyfinMediaSourceInfo) GetFallbackMaxStreamingBitrateOk() (*int32, bool)`

GetFallbackMaxStreamingBitrateOk returns a tuple with the FallbackMaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMaxStreamingBitrate

`func (o *JellyfinMediaSourceInfo) SetFallbackMaxStreamingBitrate(v int32)`

SetFallbackMaxStreamingBitrate sets FallbackMaxStreamingBitrate field to given value.

### HasFallbackMaxStreamingBitrate

`func (o *JellyfinMediaSourceInfo) HasFallbackMaxStreamingBitrate() bool`

HasFallbackMaxStreamingBitrate returns a boolean if a field has been set.

### SetFallbackMaxStreamingBitrateNil

`func (o *JellyfinMediaSourceInfo) SetFallbackMaxStreamingBitrateNil(b bool)`

 SetFallbackMaxStreamingBitrateNil sets the value for FallbackMaxStreamingBitrate to be an explicit nil

### UnsetFallbackMaxStreamingBitrate
`func (o *JellyfinMediaSourceInfo) UnsetFallbackMaxStreamingBitrate()`

UnsetFallbackMaxStreamingBitrate ensures that no value is present for FallbackMaxStreamingBitrate, not even an explicit nil
### GetTimestamp

`func (o *JellyfinMediaSourceInfo) GetTimestamp() JellyfinJellyfinTransportStreamTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JellyfinMediaSourceInfo) GetTimestampOk() (*JellyfinJellyfinTransportStreamTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JellyfinMediaSourceInfo) SetTimestamp(v JellyfinJellyfinTransportStreamTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JellyfinMediaSourceInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JellyfinMediaSourceInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JellyfinMediaSourceInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequiredHttpHeaders

`func (o *JellyfinMediaSourceInfo) GetRequiredHttpHeaders() map[string]string`

GetRequiredHttpHeaders returns the RequiredHttpHeaders field if non-nil, zero value otherwise.

### GetRequiredHttpHeadersOk

`func (o *JellyfinMediaSourceInfo) GetRequiredHttpHeadersOk() (*map[string]string, bool)`

GetRequiredHttpHeadersOk returns a tuple with the RequiredHttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHttpHeaders

`func (o *JellyfinMediaSourceInfo) SetRequiredHttpHeaders(v map[string]string)`

SetRequiredHttpHeaders sets RequiredHttpHeaders field to given value.

### HasRequiredHttpHeaders

`func (o *JellyfinMediaSourceInfo) HasRequiredHttpHeaders() bool`

HasRequiredHttpHeaders returns a boolean if a field has been set.

### SetRequiredHttpHeadersNil

`func (o *JellyfinMediaSourceInfo) SetRequiredHttpHeadersNil(b bool)`

 SetRequiredHttpHeadersNil sets the value for RequiredHttpHeaders to be an explicit nil

### UnsetRequiredHttpHeaders
`func (o *JellyfinMediaSourceInfo) UnsetRequiredHttpHeaders()`

UnsetRequiredHttpHeaders ensures that no value is present for RequiredHttpHeaders, not even an explicit nil
### GetTranscodingUrl

`func (o *JellyfinMediaSourceInfo) GetTranscodingUrl() string`

GetTranscodingUrl returns the TranscodingUrl field if non-nil, zero value otherwise.

### GetTranscodingUrlOk

`func (o *JellyfinMediaSourceInfo) GetTranscodingUrlOk() (*string, bool)`

GetTranscodingUrlOk returns a tuple with the TranscodingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUrl

`func (o *JellyfinMediaSourceInfo) SetTranscodingUrl(v string)`

SetTranscodingUrl sets TranscodingUrl field to given value.

### HasTranscodingUrl

`func (o *JellyfinMediaSourceInfo) HasTranscodingUrl() bool`

HasTranscodingUrl returns a boolean if a field has been set.

### SetTranscodingUrlNil

`func (o *JellyfinMediaSourceInfo) SetTranscodingUrlNil(b bool)`

 SetTranscodingUrlNil sets the value for TranscodingUrl to be an explicit nil

### UnsetTranscodingUrl
`func (o *JellyfinMediaSourceInfo) UnsetTranscodingUrl()`

UnsetTranscodingUrl ensures that no value is present for TranscodingUrl, not even an explicit nil
### GetTranscodingSubProtocol

`func (o *JellyfinMediaSourceInfo) GetTranscodingSubProtocol() JellyfinJellyfinMediaStreamProtocol`

GetTranscodingSubProtocol returns the TranscodingSubProtocol field if non-nil, zero value otherwise.

### GetTranscodingSubProtocolOk

`func (o *JellyfinMediaSourceInfo) GetTranscodingSubProtocolOk() (*JellyfinJellyfinMediaStreamProtocol, bool)`

GetTranscodingSubProtocolOk returns a tuple with the TranscodingSubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingSubProtocol

`func (o *JellyfinMediaSourceInfo) SetTranscodingSubProtocol(v JellyfinJellyfinMediaStreamProtocol)`

SetTranscodingSubProtocol sets TranscodingSubProtocol field to given value.

### HasTranscodingSubProtocol

`func (o *JellyfinMediaSourceInfo) HasTranscodingSubProtocol() bool`

HasTranscodingSubProtocol returns a boolean if a field has been set.

### GetTranscodingContainer

`func (o *JellyfinMediaSourceInfo) GetTranscodingContainer() string`

GetTranscodingContainer returns the TranscodingContainer field if non-nil, zero value otherwise.

### GetTranscodingContainerOk

`func (o *JellyfinMediaSourceInfo) GetTranscodingContainerOk() (*string, bool)`

GetTranscodingContainerOk returns a tuple with the TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingContainer

`func (o *JellyfinMediaSourceInfo) SetTranscodingContainer(v string)`

SetTranscodingContainer sets TranscodingContainer field to given value.

### HasTranscodingContainer

`func (o *JellyfinMediaSourceInfo) HasTranscodingContainer() bool`

HasTranscodingContainer returns a boolean if a field has been set.

### SetTranscodingContainerNil

`func (o *JellyfinMediaSourceInfo) SetTranscodingContainerNil(b bool)`

 SetTranscodingContainerNil sets the value for TranscodingContainer to be an explicit nil

### UnsetTranscodingContainer
`func (o *JellyfinMediaSourceInfo) UnsetTranscodingContainer()`

UnsetTranscodingContainer ensures that no value is present for TranscodingContainer, not even an explicit nil
### GetAnalyzeDurationMs

`func (o *JellyfinMediaSourceInfo) GetAnalyzeDurationMs() int32`

GetAnalyzeDurationMs returns the AnalyzeDurationMs field if non-nil, zero value otherwise.

### GetAnalyzeDurationMsOk

`func (o *JellyfinMediaSourceInfo) GetAnalyzeDurationMsOk() (*int32, bool)`

GetAnalyzeDurationMsOk returns a tuple with the AnalyzeDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeDurationMs

`func (o *JellyfinMediaSourceInfo) SetAnalyzeDurationMs(v int32)`

SetAnalyzeDurationMs sets AnalyzeDurationMs field to given value.

### HasAnalyzeDurationMs

`func (o *JellyfinMediaSourceInfo) HasAnalyzeDurationMs() bool`

HasAnalyzeDurationMs returns a boolean if a field has been set.

### SetAnalyzeDurationMsNil

`func (o *JellyfinMediaSourceInfo) SetAnalyzeDurationMsNil(b bool)`

 SetAnalyzeDurationMsNil sets the value for AnalyzeDurationMs to be an explicit nil

### UnsetAnalyzeDurationMs
`func (o *JellyfinMediaSourceInfo) UnsetAnalyzeDurationMs()`

UnsetAnalyzeDurationMs ensures that no value is present for AnalyzeDurationMs, not even an explicit nil
### GetDefaultAudioStreamIndex

`func (o *JellyfinMediaSourceInfo) GetDefaultAudioStreamIndex() int32`

GetDefaultAudioStreamIndex returns the DefaultAudioStreamIndex field if non-nil, zero value otherwise.

### GetDefaultAudioStreamIndexOk

`func (o *JellyfinMediaSourceInfo) GetDefaultAudioStreamIndexOk() (*int32, bool)`

GetDefaultAudioStreamIndexOk returns a tuple with the DefaultAudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioStreamIndex

`func (o *JellyfinMediaSourceInfo) SetDefaultAudioStreamIndex(v int32)`

SetDefaultAudioStreamIndex sets DefaultAudioStreamIndex field to given value.

### HasDefaultAudioStreamIndex

`func (o *JellyfinMediaSourceInfo) HasDefaultAudioStreamIndex() bool`

HasDefaultAudioStreamIndex returns a boolean if a field has been set.

### SetDefaultAudioStreamIndexNil

`func (o *JellyfinMediaSourceInfo) SetDefaultAudioStreamIndexNil(b bool)`

 SetDefaultAudioStreamIndexNil sets the value for DefaultAudioStreamIndex to be an explicit nil

### UnsetDefaultAudioStreamIndex
`func (o *JellyfinMediaSourceInfo) UnsetDefaultAudioStreamIndex()`

UnsetDefaultAudioStreamIndex ensures that no value is present for DefaultAudioStreamIndex, not even an explicit nil
### GetDefaultSubtitleStreamIndex

`func (o *JellyfinMediaSourceInfo) GetDefaultSubtitleStreamIndex() int32`

GetDefaultSubtitleStreamIndex returns the DefaultSubtitleStreamIndex field if non-nil, zero value otherwise.

### GetDefaultSubtitleStreamIndexOk

`func (o *JellyfinMediaSourceInfo) GetDefaultSubtitleStreamIndexOk() (*int32, bool)`

GetDefaultSubtitleStreamIndexOk returns a tuple with the DefaultSubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubtitleStreamIndex

`func (o *JellyfinMediaSourceInfo) SetDefaultSubtitleStreamIndex(v int32)`

SetDefaultSubtitleStreamIndex sets DefaultSubtitleStreamIndex field to given value.

### HasDefaultSubtitleStreamIndex

`func (o *JellyfinMediaSourceInfo) HasDefaultSubtitleStreamIndex() bool`

HasDefaultSubtitleStreamIndex returns a boolean if a field has been set.

### SetDefaultSubtitleStreamIndexNil

`func (o *JellyfinMediaSourceInfo) SetDefaultSubtitleStreamIndexNil(b bool)`

 SetDefaultSubtitleStreamIndexNil sets the value for DefaultSubtitleStreamIndex to be an explicit nil

### UnsetDefaultSubtitleStreamIndex
`func (o *JellyfinMediaSourceInfo) UnsetDefaultSubtitleStreamIndex()`

UnsetDefaultSubtitleStreamIndex ensures that no value is present for DefaultSubtitleStreamIndex, not even an explicit nil
### GetHasSegments

`func (o *JellyfinMediaSourceInfo) GetHasSegments() bool`

GetHasSegments returns the HasSegments field if non-nil, zero value otherwise.

### GetHasSegmentsOk

`func (o *JellyfinMediaSourceInfo) GetHasSegmentsOk() (*bool, bool)`

GetHasSegmentsOk returns a tuple with the HasSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSegments

`func (o *JellyfinMediaSourceInfo) SetHasSegments(v bool)`

SetHasSegments sets HasSegments field to given value.

### HasHasSegments

`func (o *JellyfinMediaSourceInfo) HasHasSegments() bool`

HasHasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


