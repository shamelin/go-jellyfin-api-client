# JellyfinEncodingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingThreadCount** | Pointer to **int32** | Gets or sets the thread count used for encoding. | [optional] 
**TranscodingTempPath** | Pointer to **NullableString** | Gets or sets the temporary transcoding path. | [optional] 
**FallbackFontPath** | Pointer to **NullableString** | Gets or sets the path to the fallback font. | [optional] 
**EnableFallbackFont** | Pointer to **bool** | Gets or sets a value indicating whether to use the fallback font. | [optional] 
**EnableAudioVbr** | Pointer to **bool** | Gets or sets a value indicating whether audio VBR is enabled. | [optional] 
**DownMixAudioBoost** | Pointer to **float64** | Gets or sets the audio boost applied when downmixing audio. | [optional] 
**DownMixStereoAlgorithm** | Pointer to [**JellyfinJellyfinDownMixStereoAlgorithms**](JellyfinDownMixStereoAlgorithms.md) | Gets or sets the algorithm used for downmixing audio to stereo. | [optional] 
**MaxMuxingQueueSize** | Pointer to **int32** | Gets or sets the maximum size of the muxing queue. | [optional] 
**EnableThrottling** | Pointer to **bool** | Gets or sets a value indicating whether throttling is enabled. | [optional] 
**ThrottleDelaySeconds** | Pointer to **int32** | Gets or sets the delay after which throttling happens. | [optional] 
**EnableSegmentDeletion** | Pointer to **bool** | Gets or sets a value indicating whether segment deletion is enabled. | [optional] 
**SegmentKeepSeconds** | Pointer to **int32** | Gets or sets seconds for which segments should be kept before being deleted. | [optional] 
**HardwareAccelerationType** | Pointer to [**JellyfinJellyfinHardwareAccelerationType**](JellyfinHardwareAccelerationType.md) | Gets or sets the hardware acceleration type. | [optional] 
**EncoderAppPath** | Pointer to **NullableString** | Gets or sets the FFmpeg path as set by the user via the UI. | [optional] 
**EncoderAppPathDisplay** | Pointer to **NullableString** | Gets or sets the current FFmpeg path being used by the system and displayed on the transcode page. | [optional] 
**VaapiDevice** | Pointer to **NullableString** | Gets or sets the VA-API device. | [optional] 
**QsvDevice** | Pointer to **NullableString** | Gets or sets the QSV device. | [optional] 
**EnableTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether tonemapping is enabled. | [optional] 
**EnableVppTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether VPP tonemapping is enabled. | [optional] 
**EnableVideoToolboxTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether videotoolbox tonemapping is enabled. | [optional] 
**TonemappingAlgorithm** | Pointer to [**JellyfinJellyfinTonemappingAlgorithm**](JellyfinTonemappingAlgorithm.md) | Gets or sets the tone-mapping algorithm. | [optional] 
**TonemappingMode** | Pointer to [**JellyfinJellyfinTonemappingMode**](JellyfinTonemappingMode.md) | Gets or sets the tone-mapping mode. | [optional] 
**TonemappingRange** | Pointer to [**JellyfinJellyfinTonemappingRange**](JellyfinTonemappingRange.md) | Gets or sets the tone-mapping range. | [optional] 
**TonemappingDesat** | Pointer to **float64** | Gets or sets the tone-mapping desaturation. | [optional] 
**TonemappingPeak** | Pointer to **float64** | Gets or sets the tone-mapping peak. | [optional] 
**TonemappingParam** | Pointer to **float64** | Gets or sets the tone-mapping parameters. | [optional] 
**VppTonemappingBrightness** | Pointer to **float64** | Gets or sets the VPP tone-mapping brightness. | [optional] 
**VppTonemappingContrast** | Pointer to **float64** | Gets or sets the VPP tone-mapping contrast. | [optional] 
**H264Crf** | Pointer to **int32** | Gets or sets the H264 CRF. | [optional] 
**H265Crf** | Pointer to **int32** | Gets or sets the H265 CRF. | [optional] 
**EncoderPreset** | Pointer to [**NullableJellyfinJellyfinEncoderPreset**](JellyfinEncoderPreset.md) | Gets or sets the encoder preset. | [optional] 
**DeinterlaceDoubleRate** | Pointer to **bool** | Gets or sets a value indicating whether the framerate is doubled when deinterlacing. | [optional] 
**DeinterlaceMethod** | Pointer to [**JellyfinJellyfinDeinterlaceMethod**](JellyfinDeinterlaceMethod.md) | Gets or sets the deinterlace method. | [optional] 
**EnableDecodingColorDepth10Hevc** | Pointer to **bool** | Gets or sets a value indicating whether 10bit HEVC decoding is enabled. | [optional] 
**EnableDecodingColorDepth10Vp9** | Pointer to **bool** | Gets or sets a value indicating whether 10bit VP9 decoding is enabled. | [optional] 
**EnableDecodingColorDepth10HevcRext** | Pointer to **bool** | Gets or sets a value indicating whether 8/10bit HEVC RExt decoding is enabled. | [optional] 
**EnableDecodingColorDepth12HevcRext** | Pointer to **bool** | Gets or sets a value indicating whether 12bit HEVC RExt decoding is enabled. | [optional] 
**EnableEnhancedNvdecDecoder** | Pointer to **bool** | Gets or sets a value indicating whether the enhanced NVDEC is enabled. | [optional] 
**PreferSystemNativeHwDecoder** | Pointer to **bool** | Gets or sets a value indicating whether the system native hardware decoder should be used. | [optional] 
**EnableIntelLowPowerH264HwEncoder** | Pointer to **bool** | Gets or sets a value indicating whether the Intel H264 low-power hardware encoder should be used. | [optional] 
**EnableIntelLowPowerHevcHwEncoder** | Pointer to **bool** | Gets or sets a value indicating whether the Intel HEVC low-power hardware encoder should be used. | [optional] 
**EnableHardwareEncoding** | Pointer to **bool** | Gets or sets a value indicating whether hardware encoding is enabled. | [optional] 
**AllowHevcEncoding** | Pointer to **bool** | Gets or sets a value indicating whether HEVC encoding is enabled. | [optional] 
**AllowAv1Encoding** | Pointer to **bool** | Gets or sets a value indicating whether AV1 encoding is enabled. | [optional] 
**EnableSubtitleExtraction** | Pointer to **bool** | Gets or sets a value indicating whether subtitle extraction is enabled. | [optional] 
**HardwareDecodingCodecs** | Pointer to **[]string** | Gets or sets the codecs hardware encoding is used for. | [optional] 
**AllowOnDemandMetadataBasedKeyframeExtractionForExtensions** | Pointer to **[]string** | Gets or sets the file extensions on-demand metadata based keyframe extraction is enabled for. | [optional] 

## Methods

### NewJellyfinEncodingOptions

`func NewJellyfinEncodingOptions() *JellyfinEncodingOptions`

NewJellyfinEncodingOptions instantiates a new JellyfinEncodingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinEncodingOptionsWithDefaults

`func NewJellyfinEncodingOptionsWithDefaults() *JellyfinEncodingOptions`

NewJellyfinEncodingOptionsWithDefaults instantiates a new JellyfinEncodingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingThreadCount

`func (o *JellyfinEncodingOptions) GetEncodingThreadCount() int32`

GetEncodingThreadCount returns the EncodingThreadCount field if non-nil, zero value otherwise.

### GetEncodingThreadCountOk

`func (o *JellyfinEncodingOptions) GetEncodingThreadCountOk() (*int32, bool)`

GetEncodingThreadCountOk returns a tuple with the EncodingThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingThreadCount

`func (o *JellyfinEncodingOptions) SetEncodingThreadCount(v int32)`

SetEncodingThreadCount sets EncodingThreadCount field to given value.

### HasEncodingThreadCount

`func (o *JellyfinEncodingOptions) HasEncodingThreadCount() bool`

HasEncodingThreadCount returns a boolean if a field has been set.

### GetTranscodingTempPath

`func (o *JellyfinEncodingOptions) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *JellyfinEncodingOptions) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *JellyfinEncodingOptions) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *JellyfinEncodingOptions) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### SetTranscodingTempPathNil

`func (o *JellyfinEncodingOptions) SetTranscodingTempPathNil(b bool)`

 SetTranscodingTempPathNil sets the value for TranscodingTempPath to be an explicit nil

### UnsetTranscodingTempPath
`func (o *JellyfinEncodingOptions) UnsetTranscodingTempPath()`

UnsetTranscodingTempPath ensures that no value is present for TranscodingTempPath, not even an explicit nil
### GetFallbackFontPath

`func (o *JellyfinEncodingOptions) GetFallbackFontPath() string`

GetFallbackFontPath returns the FallbackFontPath field if non-nil, zero value otherwise.

### GetFallbackFontPathOk

`func (o *JellyfinEncodingOptions) GetFallbackFontPathOk() (*string, bool)`

GetFallbackFontPathOk returns a tuple with the FallbackFontPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackFontPath

`func (o *JellyfinEncodingOptions) SetFallbackFontPath(v string)`

SetFallbackFontPath sets FallbackFontPath field to given value.

### HasFallbackFontPath

`func (o *JellyfinEncodingOptions) HasFallbackFontPath() bool`

HasFallbackFontPath returns a boolean if a field has been set.

### SetFallbackFontPathNil

`func (o *JellyfinEncodingOptions) SetFallbackFontPathNil(b bool)`

 SetFallbackFontPathNil sets the value for FallbackFontPath to be an explicit nil

### UnsetFallbackFontPath
`func (o *JellyfinEncodingOptions) UnsetFallbackFontPath()`

UnsetFallbackFontPath ensures that no value is present for FallbackFontPath, not even an explicit nil
### GetEnableFallbackFont

`func (o *JellyfinEncodingOptions) GetEnableFallbackFont() bool`

GetEnableFallbackFont returns the EnableFallbackFont field if non-nil, zero value otherwise.

### GetEnableFallbackFontOk

`func (o *JellyfinEncodingOptions) GetEnableFallbackFontOk() (*bool, bool)`

GetEnableFallbackFontOk returns a tuple with the EnableFallbackFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFallbackFont

`func (o *JellyfinEncodingOptions) SetEnableFallbackFont(v bool)`

SetEnableFallbackFont sets EnableFallbackFont field to given value.

### HasEnableFallbackFont

`func (o *JellyfinEncodingOptions) HasEnableFallbackFont() bool`

HasEnableFallbackFont returns a boolean if a field has been set.

### GetEnableAudioVbr

`func (o *JellyfinEncodingOptions) GetEnableAudioVbr() bool`

GetEnableAudioVbr returns the EnableAudioVbr field if non-nil, zero value otherwise.

### GetEnableAudioVbrOk

`func (o *JellyfinEncodingOptions) GetEnableAudioVbrOk() (*bool, bool)`

GetEnableAudioVbrOk returns a tuple with the EnableAudioVbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioVbr

`func (o *JellyfinEncodingOptions) SetEnableAudioVbr(v bool)`

SetEnableAudioVbr sets EnableAudioVbr field to given value.

### HasEnableAudioVbr

`func (o *JellyfinEncodingOptions) HasEnableAudioVbr() bool`

HasEnableAudioVbr returns a boolean if a field has been set.

### GetDownMixAudioBoost

`func (o *JellyfinEncodingOptions) GetDownMixAudioBoost() float64`

GetDownMixAudioBoost returns the DownMixAudioBoost field if non-nil, zero value otherwise.

### GetDownMixAudioBoostOk

`func (o *JellyfinEncodingOptions) GetDownMixAudioBoostOk() (*float64, bool)`

GetDownMixAudioBoostOk returns a tuple with the DownMixAudioBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixAudioBoost

`func (o *JellyfinEncodingOptions) SetDownMixAudioBoost(v float64)`

SetDownMixAudioBoost sets DownMixAudioBoost field to given value.

### HasDownMixAudioBoost

`func (o *JellyfinEncodingOptions) HasDownMixAudioBoost() bool`

HasDownMixAudioBoost returns a boolean if a field has been set.

### GetDownMixStereoAlgorithm

`func (o *JellyfinEncodingOptions) GetDownMixStereoAlgorithm() JellyfinJellyfinDownMixStereoAlgorithms`

GetDownMixStereoAlgorithm returns the DownMixStereoAlgorithm field if non-nil, zero value otherwise.

### GetDownMixStereoAlgorithmOk

`func (o *JellyfinEncodingOptions) GetDownMixStereoAlgorithmOk() (*JellyfinJellyfinDownMixStereoAlgorithms, bool)`

GetDownMixStereoAlgorithmOk returns a tuple with the DownMixStereoAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixStereoAlgorithm

`func (o *JellyfinEncodingOptions) SetDownMixStereoAlgorithm(v JellyfinJellyfinDownMixStereoAlgorithms)`

SetDownMixStereoAlgorithm sets DownMixStereoAlgorithm field to given value.

### HasDownMixStereoAlgorithm

`func (o *JellyfinEncodingOptions) HasDownMixStereoAlgorithm() bool`

HasDownMixStereoAlgorithm returns a boolean if a field has been set.

### GetMaxMuxingQueueSize

`func (o *JellyfinEncodingOptions) GetMaxMuxingQueueSize() int32`

GetMaxMuxingQueueSize returns the MaxMuxingQueueSize field if non-nil, zero value otherwise.

### GetMaxMuxingQueueSizeOk

`func (o *JellyfinEncodingOptions) GetMaxMuxingQueueSizeOk() (*int32, bool)`

GetMaxMuxingQueueSizeOk returns a tuple with the MaxMuxingQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMuxingQueueSize

`func (o *JellyfinEncodingOptions) SetMaxMuxingQueueSize(v int32)`

SetMaxMuxingQueueSize sets MaxMuxingQueueSize field to given value.

### HasMaxMuxingQueueSize

`func (o *JellyfinEncodingOptions) HasMaxMuxingQueueSize() bool`

HasMaxMuxingQueueSize returns a boolean if a field has been set.

### GetEnableThrottling

`func (o *JellyfinEncodingOptions) GetEnableThrottling() bool`

GetEnableThrottling returns the EnableThrottling field if non-nil, zero value otherwise.

### GetEnableThrottlingOk

`func (o *JellyfinEncodingOptions) GetEnableThrottlingOk() (*bool, bool)`

GetEnableThrottlingOk returns a tuple with the EnableThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThrottling

`func (o *JellyfinEncodingOptions) SetEnableThrottling(v bool)`

SetEnableThrottling sets EnableThrottling field to given value.

### HasEnableThrottling

`func (o *JellyfinEncodingOptions) HasEnableThrottling() bool`

HasEnableThrottling returns a boolean if a field has been set.

### GetThrottleDelaySeconds

`func (o *JellyfinEncodingOptions) GetThrottleDelaySeconds() int32`

GetThrottleDelaySeconds returns the ThrottleDelaySeconds field if non-nil, zero value otherwise.

### GetThrottleDelaySecondsOk

`func (o *JellyfinEncodingOptions) GetThrottleDelaySecondsOk() (*int32, bool)`

GetThrottleDelaySecondsOk returns a tuple with the ThrottleDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDelaySeconds

`func (o *JellyfinEncodingOptions) SetThrottleDelaySeconds(v int32)`

SetThrottleDelaySeconds sets ThrottleDelaySeconds field to given value.

### HasThrottleDelaySeconds

`func (o *JellyfinEncodingOptions) HasThrottleDelaySeconds() bool`

HasThrottleDelaySeconds returns a boolean if a field has been set.

### GetEnableSegmentDeletion

`func (o *JellyfinEncodingOptions) GetEnableSegmentDeletion() bool`

GetEnableSegmentDeletion returns the EnableSegmentDeletion field if non-nil, zero value otherwise.

### GetEnableSegmentDeletionOk

`func (o *JellyfinEncodingOptions) GetEnableSegmentDeletionOk() (*bool, bool)`

GetEnableSegmentDeletionOk returns a tuple with the EnableSegmentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSegmentDeletion

`func (o *JellyfinEncodingOptions) SetEnableSegmentDeletion(v bool)`

SetEnableSegmentDeletion sets EnableSegmentDeletion field to given value.

### HasEnableSegmentDeletion

`func (o *JellyfinEncodingOptions) HasEnableSegmentDeletion() bool`

HasEnableSegmentDeletion returns a boolean if a field has been set.

### GetSegmentKeepSeconds

`func (o *JellyfinEncodingOptions) GetSegmentKeepSeconds() int32`

GetSegmentKeepSeconds returns the SegmentKeepSeconds field if non-nil, zero value otherwise.

### GetSegmentKeepSecondsOk

`func (o *JellyfinEncodingOptions) GetSegmentKeepSecondsOk() (*int32, bool)`

GetSegmentKeepSecondsOk returns a tuple with the SegmentKeepSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeepSeconds

`func (o *JellyfinEncodingOptions) SetSegmentKeepSeconds(v int32)`

SetSegmentKeepSeconds sets SegmentKeepSeconds field to given value.

### HasSegmentKeepSeconds

`func (o *JellyfinEncodingOptions) HasSegmentKeepSeconds() bool`

HasSegmentKeepSeconds returns a boolean if a field has been set.

### GetHardwareAccelerationType

`func (o *JellyfinEncodingOptions) GetHardwareAccelerationType() JellyfinJellyfinHardwareAccelerationType`

GetHardwareAccelerationType returns the HardwareAccelerationType field if non-nil, zero value otherwise.

### GetHardwareAccelerationTypeOk

`func (o *JellyfinEncodingOptions) GetHardwareAccelerationTypeOk() (*JellyfinJellyfinHardwareAccelerationType, bool)`

GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationType

`func (o *JellyfinEncodingOptions) SetHardwareAccelerationType(v JellyfinJellyfinHardwareAccelerationType)`

SetHardwareAccelerationType sets HardwareAccelerationType field to given value.

### HasHardwareAccelerationType

`func (o *JellyfinEncodingOptions) HasHardwareAccelerationType() bool`

HasHardwareAccelerationType returns a boolean if a field has been set.

### GetEncoderAppPath

`func (o *JellyfinEncodingOptions) GetEncoderAppPath() string`

GetEncoderAppPath returns the EncoderAppPath field if non-nil, zero value otherwise.

### GetEncoderAppPathOk

`func (o *JellyfinEncodingOptions) GetEncoderAppPathOk() (*string, bool)`

GetEncoderAppPathOk returns a tuple with the EncoderAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPath

`func (o *JellyfinEncodingOptions) SetEncoderAppPath(v string)`

SetEncoderAppPath sets EncoderAppPath field to given value.

### HasEncoderAppPath

`func (o *JellyfinEncodingOptions) HasEncoderAppPath() bool`

HasEncoderAppPath returns a boolean if a field has been set.

### SetEncoderAppPathNil

`func (o *JellyfinEncodingOptions) SetEncoderAppPathNil(b bool)`

 SetEncoderAppPathNil sets the value for EncoderAppPath to be an explicit nil

### UnsetEncoderAppPath
`func (o *JellyfinEncodingOptions) UnsetEncoderAppPath()`

UnsetEncoderAppPath ensures that no value is present for EncoderAppPath, not even an explicit nil
### GetEncoderAppPathDisplay

`func (o *JellyfinEncodingOptions) GetEncoderAppPathDisplay() string`

GetEncoderAppPathDisplay returns the EncoderAppPathDisplay field if non-nil, zero value otherwise.

### GetEncoderAppPathDisplayOk

`func (o *JellyfinEncodingOptions) GetEncoderAppPathDisplayOk() (*string, bool)`

GetEncoderAppPathDisplayOk returns a tuple with the EncoderAppPathDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPathDisplay

`func (o *JellyfinEncodingOptions) SetEncoderAppPathDisplay(v string)`

SetEncoderAppPathDisplay sets EncoderAppPathDisplay field to given value.

### HasEncoderAppPathDisplay

`func (o *JellyfinEncodingOptions) HasEncoderAppPathDisplay() bool`

HasEncoderAppPathDisplay returns a boolean if a field has been set.

### SetEncoderAppPathDisplayNil

`func (o *JellyfinEncodingOptions) SetEncoderAppPathDisplayNil(b bool)`

 SetEncoderAppPathDisplayNil sets the value for EncoderAppPathDisplay to be an explicit nil

### UnsetEncoderAppPathDisplay
`func (o *JellyfinEncodingOptions) UnsetEncoderAppPathDisplay()`

UnsetEncoderAppPathDisplay ensures that no value is present for EncoderAppPathDisplay, not even an explicit nil
### GetVaapiDevice

`func (o *JellyfinEncodingOptions) GetVaapiDevice() string`

GetVaapiDevice returns the VaapiDevice field if non-nil, zero value otherwise.

### GetVaapiDeviceOk

`func (o *JellyfinEncodingOptions) GetVaapiDeviceOk() (*string, bool)`

GetVaapiDeviceOk returns a tuple with the VaapiDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaapiDevice

`func (o *JellyfinEncodingOptions) SetVaapiDevice(v string)`

SetVaapiDevice sets VaapiDevice field to given value.

### HasVaapiDevice

`func (o *JellyfinEncodingOptions) HasVaapiDevice() bool`

HasVaapiDevice returns a boolean if a field has been set.

### SetVaapiDeviceNil

`func (o *JellyfinEncodingOptions) SetVaapiDeviceNil(b bool)`

 SetVaapiDeviceNil sets the value for VaapiDevice to be an explicit nil

### UnsetVaapiDevice
`func (o *JellyfinEncodingOptions) UnsetVaapiDevice()`

UnsetVaapiDevice ensures that no value is present for VaapiDevice, not even an explicit nil
### GetQsvDevice

`func (o *JellyfinEncodingOptions) GetQsvDevice() string`

GetQsvDevice returns the QsvDevice field if non-nil, zero value otherwise.

### GetQsvDeviceOk

`func (o *JellyfinEncodingOptions) GetQsvDeviceOk() (*string, bool)`

GetQsvDeviceOk returns a tuple with the QsvDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQsvDevice

`func (o *JellyfinEncodingOptions) SetQsvDevice(v string)`

SetQsvDevice sets QsvDevice field to given value.

### HasQsvDevice

`func (o *JellyfinEncodingOptions) HasQsvDevice() bool`

HasQsvDevice returns a boolean if a field has been set.

### SetQsvDeviceNil

`func (o *JellyfinEncodingOptions) SetQsvDeviceNil(b bool)`

 SetQsvDeviceNil sets the value for QsvDevice to be an explicit nil

### UnsetQsvDevice
`func (o *JellyfinEncodingOptions) UnsetQsvDevice()`

UnsetQsvDevice ensures that no value is present for QsvDevice, not even an explicit nil
### GetEnableTonemapping

`func (o *JellyfinEncodingOptions) GetEnableTonemapping() bool`

GetEnableTonemapping returns the EnableTonemapping field if non-nil, zero value otherwise.

### GetEnableTonemappingOk

`func (o *JellyfinEncodingOptions) GetEnableTonemappingOk() (*bool, bool)`

GetEnableTonemappingOk returns a tuple with the EnableTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTonemapping

`func (o *JellyfinEncodingOptions) SetEnableTonemapping(v bool)`

SetEnableTonemapping sets EnableTonemapping field to given value.

### HasEnableTonemapping

`func (o *JellyfinEncodingOptions) HasEnableTonemapping() bool`

HasEnableTonemapping returns a boolean if a field has been set.

### GetEnableVppTonemapping

`func (o *JellyfinEncodingOptions) GetEnableVppTonemapping() bool`

GetEnableVppTonemapping returns the EnableVppTonemapping field if non-nil, zero value otherwise.

### GetEnableVppTonemappingOk

`func (o *JellyfinEncodingOptions) GetEnableVppTonemappingOk() (*bool, bool)`

GetEnableVppTonemappingOk returns a tuple with the EnableVppTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVppTonemapping

`func (o *JellyfinEncodingOptions) SetEnableVppTonemapping(v bool)`

SetEnableVppTonemapping sets EnableVppTonemapping field to given value.

### HasEnableVppTonemapping

`func (o *JellyfinEncodingOptions) HasEnableVppTonemapping() bool`

HasEnableVppTonemapping returns a boolean if a field has been set.

### GetEnableVideoToolboxTonemapping

`func (o *JellyfinEncodingOptions) GetEnableVideoToolboxTonemapping() bool`

GetEnableVideoToolboxTonemapping returns the EnableVideoToolboxTonemapping field if non-nil, zero value otherwise.

### GetEnableVideoToolboxTonemappingOk

`func (o *JellyfinEncodingOptions) GetEnableVideoToolboxTonemappingOk() (*bool, bool)`

GetEnableVideoToolboxTonemappingOk returns a tuple with the EnableVideoToolboxTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoToolboxTonemapping

`func (o *JellyfinEncodingOptions) SetEnableVideoToolboxTonemapping(v bool)`

SetEnableVideoToolboxTonemapping sets EnableVideoToolboxTonemapping field to given value.

### HasEnableVideoToolboxTonemapping

`func (o *JellyfinEncodingOptions) HasEnableVideoToolboxTonemapping() bool`

HasEnableVideoToolboxTonemapping returns a boolean if a field has been set.

### GetTonemappingAlgorithm

`func (o *JellyfinEncodingOptions) GetTonemappingAlgorithm() JellyfinJellyfinTonemappingAlgorithm`

GetTonemappingAlgorithm returns the TonemappingAlgorithm field if non-nil, zero value otherwise.

### GetTonemappingAlgorithmOk

`func (o *JellyfinEncodingOptions) GetTonemappingAlgorithmOk() (*JellyfinJellyfinTonemappingAlgorithm, bool)`

GetTonemappingAlgorithmOk returns a tuple with the TonemappingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingAlgorithm

`func (o *JellyfinEncodingOptions) SetTonemappingAlgorithm(v JellyfinJellyfinTonemappingAlgorithm)`

SetTonemappingAlgorithm sets TonemappingAlgorithm field to given value.

### HasTonemappingAlgorithm

`func (o *JellyfinEncodingOptions) HasTonemappingAlgorithm() bool`

HasTonemappingAlgorithm returns a boolean if a field has been set.

### GetTonemappingMode

`func (o *JellyfinEncodingOptions) GetTonemappingMode() JellyfinJellyfinTonemappingMode`

GetTonemappingMode returns the TonemappingMode field if non-nil, zero value otherwise.

### GetTonemappingModeOk

`func (o *JellyfinEncodingOptions) GetTonemappingModeOk() (*JellyfinJellyfinTonemappingMode, bool)`

GetTonemappingModeOk returns a tuple with the TonemappingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingMode

`func (o *JellyfinEncodingOptions) SetTonemappingMode(v JellyfinJellyfinTonemappingMode)`

SetTonemappingMode sets TonemappingMode field to given value.

### HasTonemappingMode

`func (o *JellyfinEncodingOptions) HasTonemappingMode() bool`

HasTonemappingMode returns a boolean if a field has been set.

### GetTonemappingRange

`func (o *JellyfinEncodingOptions) GetTonemappingRange() JellyfinJellyfinTonemappingRange`

GetTonemappingRange returns the TonemappingRange field if non-nil, zero value otherwise.

### GetTonemappingRangeOk

`func (o *JellyfinEncodingOptions) GetTonemappingRangeOk() (*JellyfinJellyfinTonemappingRange, bool)`

GetTonemappingRangeOk returns a tuple with the TonemappingRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingRange

`func (o *JellyfinEncodingOptions) SetTonemappingRange(v JellyfinJellyfinTonemappingRange)`

SetTonemappingRange sets TonemappingRange field to given value.

### HasTonemappingRange

`func (o *JellyfinEncodingOptions) HasTonemappingRange() bool`

HasTonemappingRange returns a boolean if a field has been set.

### GetTonemappingDesat

`func (o *JellyfinEncodingOptions) GetTonemappingDesat() float64`

GetTonemappingDesat returns the TonemappingDesat field if non-nil, zero value otherwise.

### GetTonemappingDesatOk

`func (o *JellyfinEncodingOptions) GetTonemappingDesatOk() (*float64, bool)`

GetTonemappingDesatOk returns a tuple with the TonemappingDesat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingDesat

`func (o *JellyfinEncodingOptions) SetTonemappingDesat(v float64)`

SetTonemappingDesat sets TonemappingDesat field to given value.

### HasTonemappingDesat

`func (o *JellyfinEncodingOptions) HasTonemappingDesat() bool`

HasTonemappingDesat returns a boolean if a field has been set.

### GetTonemappingPeak

`func (o *JellyfinEncodingOptions) GetTonemappingPeak() float64`

GetTonemappingPeak returns the TonemappingPeak field if non-nil, zero value otherwise.

### GetTonemappingPeakOk

`func (o *JellyfinEncodingOptions) GetTonemappingPeakOk() (*float64, bool)`

GetTonemappingPeakOk returns a tuple with the TonemappingPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingPeak

`func (o *JellyfinEncodingOptions) SetTonemappingPeak(v float64)`

SetTonemappingPeak sets TonemappingPeak field to given value.

### HasTonemappingPeak

`func (o *JellyfinEncodingOptions) HasTonemappingPeak() bool`

HasTonemappingPeak returns a boolean if a field has been set.

### GetTonemappingParam

`func (o *JellyfinEncodingOptions) GetTonemappingParam() float64`

GetTonemappingParam returns the TonemappingParam field if non-nil, zero value otherwise.

### GetTonemappingParamOk

`func (o *JellyfinEncodingOptions) GetTonemappingParamOk() (*float64, bool)`

GetTonemappingParamOk returns a tuple with the TonemappingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingParam

`func (o *JellyfinEncodingOptions) SetTonemappingParam(v float64)`

SetTonemappingParam sets TonemappingParam field to given value.

### HasTonemappingParam

`func (o *JellyfinEncodingOptions) HasTonemappingParam() bool`

HasTonemappingParam returns a boolean if a field has been set.

### GetVppTonemappingBrightness

`func (o *JellyfinEncodingOptions) GetVppTonemappingBrightness() float64`

GetVppTonemappingBrightness returns the VppTonemappingBrightness field if non-nil, zero value otherwise.

### GetVppTonemappingBrightnessOk

`func (o *JellyfinEncodingOptions) GetVppTonemappingBrightnessOk() (*float64, bool)`

GetVppTonemappingBrightnessOk returns a tuple with the VppTonemappingBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingBrightness

`func (o *JellyfinEncodingOptions) SetVppTonemappingBrightness(v float64)`

SetVppTonemappingBrightness sets VppTonemappingBrightness field to given value.

### HasVppTonemappingBrightness

`func (o *JellyfinEncodingOptions) HasVppTonemappingBrightness() bool`

HasVppTonemappingBrightness returns a boolean if a field has been set.

### GetVppTonemappingContrast

`func (o *JellyfinEncodingOptions) GetVppTonemappingContrast() float64`

GetVppTonemappingContrast returns the VppTonemappingContrast field if non-nil, zero value otherwise.

### GetVppTonemappingContrastOk

`func (o *JellyfinEncodingOptions) GetVppTonemappingContrastOk() (*float64, bool)`

GetVppTonemappingContrastOk returns a tuple with the VppTonemappingContrast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingContrast

`func (o *JellyfinEncodingOptions) SetVppTonemappingContrast(v float64)`

SetVppTonemappingContrast sets VppTonemappingContrast field to given value.

### HasVppTonemappingContrast

`func (o *JellyfinEncodingOptions) HasVppTonemappingContrast() bool`

HasVppTonemappingContrast returns a boolean if a field has been set.

### GetH264Crf

`func (o *JellyfinEncodingOptions) GetH264Crf() int32`

GetH264Crf returns the H264Crf field if non-nil, zero value otherwise.

### GetH264CrfOk

`func (o *JellyfinEncodingOptions) GetH264CrfOk() (*int32, bool)`

GetH264CrfOk returns a tuple with the H264Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH264Crf

`func (o *JellyfinEncodingOptions) SetH264Crf(v int32)`

SetH264Crf sets H264Crf field to given value.

### HasH264Crf

`func (o *JellyfinEncodingOptions) HasH264Crf() bool`

HasH264Crf returns a boolean if a field has been set.

### GetH265Crf

`func (o *JellyfinEncodingOptions) GetH265Crf() int32`

GetH265Crf returns the H265Crf field if non-nil, zero value otherwise.

### GetH265CrfOk

`func (o *JellyfinEncodingOptions) GetH265CrfOk() (*int32, bool)`

GetH265CrfOk returns a tuple with the H265Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH265Crf

`func (o *JellyfinEncodingOptions) SetH265Crf(v int32)`

SetH265Crf sets H265Crf field to given value.

### HasH265Crf

`func (o *JellyfinEncodingOptions) HasH265Crf() bool`

HasH265Crf returns a boolean if a field has been set.

### GetEncoderPreset

`func (o *JellyfinEncodingOptions) GetEncoderPreset() JellyfinJellyfinEncoderPreset`

GetEncoderPreset returns the EncoderPreset field if non-nil, zero value otherwise.

### GetEncoderPresetOk

`func (o *JellyfinEncodingOptions) GetEncoderPresetOk() (*JellyfinJellyfinEncoderPreset, bool)`

GetEncoderPresetOk returns a tuple with the EncoderPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPreset

`func (o *JellyfinEncodingOptions) SetEncoderPreset(v JellyfinJellyfinEncoderPreset)`

SetEncoderPreset sets EncoderPreset field to given value.

### HasEncoderPreset

`func (o *JellyfinEncodingOptions) HasEncoderPreset() bool`

HasEncoderPreset returns a boolean if a field has been set.

### SetEncoderPresetNil

`func (o *JellyfinEncodingOptions) SetEncoderPresetNil(b bool)`

 SetEncoderPresetNil sets the value for EncoderPreset to be an explicit nil

### UnsetEncoderPreset
`func (o *JellyfinEncodingOptions) UnsetEncoderPreset()`

UnsetEncoderPreset ensures that no value is present for EncoderPreset, not even an explicit nil
### GetDeinterlaceDoubleRate

`func (o *JellyfinEncodingOptions) GetDeinterlaceDoubleRate() bool`

GetDeinterlaceDoubleRate returns the DeinterlaceDoubleRate field if non-nil, zero value otherwise.

### GetDeinterlaceDoubleRateOk

`func (o *JellyfinEncodingOptions) GetDeinterlaceDoubleRateOk() (*bool, bool)`

GetDeinterlaceDoubleRateOk returns a tuple with the DeinterlaceDoubleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceDoubleRate

`func (o *JellyfinEncodingOptions) SetDeinterlaceDoubleRate(v bool)`

SetDeinterlaceDoubleRate sets DeinterlaceDoubleRate field to given value.

### HasDeinterlaceDoubleRate

`func (o *JellyfinEncodingOptions) HasDeinterlaceDoubleRate() bool`

HasDeinterlaceDoubleRate returns a boolean if a field has been set.

### GetDeinterlaceMethod

`func (o *JellyfinEncodingOptions) GetDeinterlaceMethod() JellyfinJellyfinDeinterlaceMethod`

GetDeinterlaceMethod returns the DeinterlaceMethod field if non-nil, zero value otherwise.

### GetDeinterlaceMethodOk

`func (o *JellyfinEncodingOptions) GetDeinterlaceMethodOk() (*JellyfinJellyfinDeinterlaceMethod, bool)`

GetDeinterlaceMethodOk returns a tuple with the DeinterlaceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceMethod

`func (o *JellyfinEncodingOptions) SetDeinterlaceMethod(v JellyfinJellyfinDeinterlaceMethod)`

SetDeinterlaceMethod sets DeinterlaceMethod field to given value.

### HasDeinterlaceMethod

`func (o *JellyfinEncodingOptions) HasDeinterlaceMethod() bool`

HasDeinterlaceMethod returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10Hevc

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10Hevc() bool`

GetEnableDecodingColorDepth10Hevc returns the EnableDecodingColorDepth10Hevc field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10HevcOk

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10HevcOk() (*bool, bool)`

GetEnableDecodingColorDepth10HevcOk returns a tuple with the EnableDecodingColorDepth10Hevc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Hevc

`func (o *JellyfinEncodingOptions) SetEnableDecodingColorDepth10Hevc(v bool)`

SetEnableDecodingColorDepth10Hevc sets EnableDecodingColorDepth10Hevc field to given value.

### HasEnableDecodingColorDepth10Hevc

`func (o *JellyfinEncodingOptions) HasEnableDecodingColorDepth10Hevc() bool`

HasEnableDecodingColorDepth10Hevc returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10Vp9

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10Vp9() bool`

GetEnableDecodingColorDepth10Vp9 returns the EnableDecodingColorDepth10Vp9 field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10Vp9Ok

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10Vp9Ok() (*bool, bool)`

GetEnableDecodingColorDepth10Vp9Ok returns a tuple with the EnableDecodingColorDepth10Vp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Vp9

`func (o *JellyfinEncodingOptions) SetEnableDecodingColorDepth10Vp9(v bool)`

SetEnableDecodingColorDepth10Vp9 sets EnableDecodingColorDepth10Vp9 field to given value.

### HasEnableDecodingColorDepth10Vp9

`func (o *JellyfinEncodingOptions) HasEnableDecodingColorDepth10Vp9() bool`

HasEnableDecodingColorDepth10Vp9 returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10HevcRext

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10HevcRext() bool`

GetEnableDecodingColorDepth10HevcRext returns the EnableDecodingColorDepth10HevcRext field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10HevcRextOk

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth10HevcRextOk() (*bool, bool)`

GetEnableDecodingColorDepth10HevcRextOk returns a tuple with the EnableDecodingColorDepth10HevcRext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10HevcRext

`func (o *JellyfinEncodingOptions) SetEnableDecodingColorDepth10HevcRext(v bool)`

SetEnableDecodingColorDepth10HevcRext sets EnableDecodingColorDepth10HevcRext field to given value.

### HasEnableDecodingColorDepth10HevcRext

`func (o *JellyfinEncodingOptions) HasEnableDecodingColorDepth10HevcRext() bool`

HasEnableDecodingColorDepth10HevcRext returns a boolean if a field has been set.

### GetEnableDecodingColorDepth12HevcRext

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth12HevcRext() bool`

GetEnableDecodingColorDepth12HevcRext returns the EnableDecodingColorDepth12HevcRext field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth12HevcRextOk

`func (o *JellyfinEncodingOptions) GetEnableDecodingColorDepth12HevcRextOk() (*bool, bool)`

GetEnableDecodingColorDepth12HevcRextOk returns a tuple with the EnableDecodingColorDepth12HevcRext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth12HevcRext

`func (o *JellyfinEncodingOptions) SetEnableDecodingColorDepth12HevcRext(v bool)`

SetEnableDecodingColorDepth12HevcRext sets EnableDecodingColorDepth12HevcRext field to given value.

### HasEnableDecodingColorDepth12HevcRext

`func (o *JellyfinEncodingOptions) HasEnableDecodingColorDepth12HevcRext() bool`

HasEnableDecodingColorDepth12HevcRext returns a boolean if a field has been set.

### GetEnableEnhancedNvdecDecoder

`func (o *JellyfinEncodingOptions) GetEnableEnhancedNvdecDecoder() bool`

GetEnableEnhancedNvdecDecoder returns the EnableEnhancedNvdecDecoder field if non-nil, zero value otherwise.

### GetEnableEnhancedNvdecDecoderOk

`func (o *JellyfinEncodingOptions) GetEnableEnhancedNvdecDecoderOk() (*bool, bool)`

GetEnableEnhancedNvdecDecoderOk returns a tuple with the EnableEnhancedNvdecDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEnhancedNvdecDecoder

`func (o *JellyfinEncodingOptions) SetEnableEnhancedNvdecDecoder(v bool)`

SetEnableEnhancedNvdecDecoder sets EnableEnhancedNvdecDecoder field to given value.

### HasEnableEnhancedNvdecDecoder

`func (o *JellyfinEncodingOptions) HasEnableEnhancedNvdecDecoder() bool`

HasEnableEnhancedNvdecDecoder returns a boolean if a field has been set.

### GetPreferSystemNativeHwDecoder

`func (o *JellyfinEncodingOptions) GetPreferSystemNativeHwDecoder() bool`

GetPreferSystemNativeHwDecoder returns the PreferSystemNativeHwDecoder field if non-nil, zero value otherwise.

### GetPreferSystemNativeHwDecoderOk

`func (o *JellyfinEncodingOptions) GetPreferSystemNativeHwDecoderOk() (*bool, bool)`

GetPreferSystemNativeHwDecoderOk returns a tuple with the PreferSystemNativeHwDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferSystemNativeHwDecoder

`func (o *JellyfinEncodingOptions) SetPreferSystemNativeHwDecoder(v bool)`

SetPreferSystemNativeHwDecoder sets PreferSystemNativeHwDecoder field to given value.

### HasPreferSystemNativeHwDecoder

`func (o *JellyfinEncodingOptions) HasPreferSystemNativeHwDecoder() bool`

HasPreferSystemNativeHwDecoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerH264HwEncoder

`func (o *JellyfinEncodingOptions) GetEnableIntelLowPowerH264HwEncoder() bool`

GetEnableIntelLowPowerH264HwEncoder returns the EnableIntelLowPowerH264HwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerH264HwEncoderOk

`func (o *JellyfinEncodingOptions) GetEnableIntelLowPowerH264HwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerH264HwEncoderOk returns a tuple with the EnableIntelLowPowerH264HwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerH264HwEncoder

`func (o *JellyfinEncodingOptions) SetEnableIntelLowPowerH264HwEncoder(v bool)`

SetEnableIntelLowPowerH264HwEncoder sets EnableIntelLowPowerH264HwEncoder field to given value.

### HasEnableIntelLowPowerH264HwEncoder

`func (o *JellyfinEncodingOptions) HasEnableIntelLowPowerH264HwEncoder() bool`

HasEnableIntelLowPowerH264HwEncoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerHevcHwEncoder

`func (o *JellyfinEncodingOptions) GetEnableIntelLowPowerHevcHwEncoder() bool`

GetEnableIntelLowPowerHevcHwEncoder returns the EnableIntelLowPowerHevcHwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerHevcHwEncoderOk

`func (o *JellyfinEncodingOptions) GetEnableIntelLowPowerHevcHwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerHevcHwEncoderOk returns a tuple with the EnableIntelLowPowerHevcHwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerHevcHwEncoder

`func (o *JellyfinEncodingOptions) SetEnableIntelLowPowerHevcHwEncoder(v bool)`

SetEnableIntelLowPowerHevcHwEncoder sets EnableIntelLowPowerHevcHwEncoder field to given value.

### HasEnableIntelLowPowerHevcHwEncoder

`func (o *JellyfinEncodingOptions) HasEnableIntelLowPowerHevcHwEncoder() bool`

HasEnableIntelLowPowerHevcHwEncoder returns a boolean if a field has been set.

### GetEnableHardwareEncoding

`func (o *JellyfinEncodingOptions) GetEnableHardwareEncoding() bool`

GetEnableHardwareEncoding returns the EnableHardwareEncoding field if non-nil, zero value otherwise.

### GetEnableHardwareEncodingOk

`func (o *JellyfinEncodingOptions) GetEnableHardwareEncodingOk() (*bool, bool)`

GetEnableHardwareEncodingOk returns a tuple with the EnableHardwareEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHardwareEncoding

`func (o *JellyfinEncodingOptions) SetEnableHardwareEncoding(v bool)`

SetEnableHardwareEncoding sets EnableHardwareEncoding field to given value.

### HasEnableHardwareEncoding

`func (o *JellyfinEncodingOptions) HasEnableHardwareEncoding() bool`

HasEnableHardwareEncoding returns a boolean if a field has been set.

### GetAllowHevcEncoding

`func (o *JellyfinEncodingOptions) GetAllowHevcEncoding() bool`

GetAllowHevcEncoding returns the AllowHevcEncoding field if non-nil, zero value otherwise.

### GetAllowHevcEncodingOk

`func (o *JellyfinEncodingOptions) GetAllowHevcEncodingOk() (*bool, bool)`

GetAllowHevcEncodingOk returns a tuple with the AllowHevcEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHevcEncoding

`func (o *JellyfinEncodingOptions) SetAllowHevcEncoding(v bool)`

SetAllowHevcEncoding sets AllowHevcEncoding field to given value.

### HasAllowHevcEncoding

`func (o *JellyfinEncodingOptions) HasAllowHevcEncoding() bool`

HasAllowHevcEncoding returns a boolean if a field has been set.

### GetAllowAv1Encoding

`func (o *JellyfinEncodingOptions) GetAllowAv1Encoding() bool`

GetAllowAv1Encoding returns the AllowAv1Encoding field if non-nil, zero value otherwise.

### GetAllowAv1EncodingOk

`func (o *JellyfinEncodingOptions) GetAllowAv1EncodingOk() (*bool, bool)`

GetAllowAv1EncodingOk returns a tuple with the AllowAv1Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAv1Encoding

`func (o *JellyfinEncodingOptions) SetAllowAv1Encoding(v bool)`

SetAllowAv1Encoding sets AllowAv1Encoding field to given value.

### HasAllowAv1Encoding

`func (o *JellyfinEncodingOptions) HasAllowAv1Encoding() bool`

HasAllowAv1Encoding returns a boolean if a field has been set.

### GetEnableSubtitleExtraction

`func (o *JellyfinEncodingOptions) GetEnableSubtitleExtraction() bool`

GetEnableSubtitleExtraction returns the EnableSubtitleExtraction field if non-nil, zero value otherwise.

### GetEnableSubtitleExtractionOk

`func (o *JellyfinEncodingOptions) GetEnableSubtitleExtractionOk() (*bool, bool)`

GetEnableSubtitleExtractionOk returns a tuple with the EnableSubtitleExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleExtraction

`func (o *JellyfinEncodingOptions) SetEnableSubtitleExtraction(v bool)`

SetEnableSubtitleExtraction sets EnableSubtitleExtraction field to given value.

### HasEnableSubtitleExtraction

`func (o *JellyfinEncodingOptions) HasEnableSubtitleExtraction() bool`

HasEnableSubtitleExtraction returns a boolean if a field has been set.

### GetHardwareDecodingCodecs

`func (o *JellyfinEncodingOptions) GetHardwareDecodingCodecs() []string`

GetHardwareDecodingCodecs returns the HardwareDecodingCodecs field if non-nil, zero value otherwise.

### GetHardwareDecodingCodecsOk

`func (o *JellyfinEncodingOptions) GetHardwareDecodingCodecsOk() (*[]string, bool)`

GetHardwareDecodingCodecsOk returns a tuple with the HardwareDecodingCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareDecodingCodecs

`func (o *JellyfinEncodingOptions) SetHardwareDecodingCodecs(v []string)`

SetHardwareDecodingCodecs sets HardwareDecodingCodecs field to given value.

### HasHardwareDecodingCodecs

`func (o *JellyfinEncodingOptions) HasHardwareDecodingCodecs() bool`

HasHardwareDecodingCodecs returns a boolean if a field has been set.

### SetHardwareDecodingCodecsNil

`func (o *JellyfinEncodingOptions) SetHardwareDecodingCodecsNil(b bool)`

 SetHardwareDecodingCodecsNil sets the value for HardwareDecodingCodecs to be an explicit nil

### UnsetHardwareDecodingCodecs
`func (o *JellyfinEncodingOptions) UnsetHardwareDecodingCodecs()`

UnsetHardwareDecodingCodecs ensures that no value is present for HardwareDecodingCodecs, not even an explicit nil
### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *JellyfinEncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() []string`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if non-nil, zero value otherwise.

### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk

`func (o *JellyfinEncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk() (*[]string, bool)`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk returns a tuple with the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *JellyfinEncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions(v []string)`

SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions sets AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field to given value.

### HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *JellyfinEncodingOptions) HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() bool`

HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns a boolean if a field has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil

`func (o *JellyfinEncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil(b bool)`

 SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil sets the value for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions to be an explicit nil

### UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions
`func (o *JellyfinEncodingOptions) UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions()`

UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions ensures that no value is present for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


