# ModelEncodingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingThreadCount** | Pointer to **int32** | Gets or sets the thread count used for encoding. | [optional] 
**TranscodingTempPath** | Pointer to **NullableString** | Gets or sets the temporary transcoding path. | [optional] 
**FallbackFontPath** | Pointer to **NullableString** | Gets or sets the path to the fallback font. | [optional] 
**EnableFallbackFont** | Pointer to **bool** | Gets or sets a value indicating whether to use the fallback font. | [optional] 
**EnableAudioVbr** | Pointer to **bool** | Gets or sets a value indicating whether audio VBR is enabled. | [optional] 
**DownMixAudioBoost** | Pointer to **float64** | Gets or sets the audio boost applied when downmixing audio. | [optional] 
**DownMixStereoAlgorithm** | Pointer to [**ModelModelDownMixStereoAlgorithms**](ModelDownMixStereoAlgorithms.md) | Gets or sets the algorithm used for downmixing audio to stereo. | [optional] 
**MaxMuxingQueueSize** | Pointer to **int32** | Gets or sets the maximum size of the muxing queue. | [optional] 
**EnableThrottling** | Pointer to **bool** | Gets or sets a value indicating whether throttling is enabled. | [optional] 
**ThrottleDelaySeconds** | Pointer to **int32** | Gets or sets the delay after which throttling happens. | [optional] 
**EnableSegmentDeletion** | Pointer to **bool** | Gets or sets a value indicating whether segment deletion is enabled. | [optional] 
**SegmentKeepSeconds** | Pointer to **int32** | Gets or sets seconds for which segments should be kept before being deleted. | [optional] 
**HardwareAccelerationType** | Pointer to [**ModelModelHardwareAccelerationType**](ModelHardwareAccelerationType.md) | Gets or sets the hardware acceleration type. | [optional] 
**EncoderAppPath** | Pointer to **NullableString** | Gets or sets the FFmpeg path as set by the user via the UI. | [optional] 
**EncoderAppPathDisplay** | Pointer to **NullableString** | Gets or sets the current FFmpeg path being used by the system and displayed on the transcode page. | [optional] 
**VaapiDevice** | Pointer to **NullableString** | Gets or sets the VA-API device. | [optional] 
**QsvDevice** | Pointer to **NullableString** | Gets or sets the QSV device. | [optional] 
**EnableTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether tonemapping is enabled. | [optional] 
**EnableVppTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether VPP tonemapping is enabled. | [optional] 
**EnableVideoToolboxTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether videotoolbox tonemapping is enabled. | [optional] 
**TonemappingAlgorithm** | Pointer to [**ModelModelTonemappingAlgorithm**](ModelTonemappingAlgorithm.md) | Gets or sets the tone-mapping algorithm. | [optional] 
**TonemappingMode** | Pointer to [**ModelModelTonemappingMode**](ModelTonemappingMode.md) | Gets or sets the tone-mapping mode. | [optional] 
**TonemappingRange** | Pointer to [**ModelModelTonemappingRange**](ModelTonemappingRange.md) | Gets or sets the tone-mapping range. | [optional] 
**TonemappingDesat** | Pointer to **float64** | Gets or sets the tone-mapping desaturation. | [optional] 
**TonemappingPeak** | Pointer to **float64** | Gets or sets the tone-mapping peak. | [optional] 
**TonemappingParam** | Pointer to **float64** | Gets or sets the tone-mapping parameters. | [optional] 
**VppTonemappingBrightness** | Pointer to **float64** | Gets or sets the VPP tone-mapping brightness. | [optional] 
**VppTonemappingContrast** | Pointer to **float64** | Gets or sets the VPP tone-mapping contrast. | [optional] 
**H264Crf** | Pointer to **int32** | Gets or sets the H264 CRF. | [optional] 
**H265Crf** | Pointer to **int32** | Gets or sets the H265 CRF. | [optional] 
**EncoderPreset** | Pointer to [**NullableModelModelEncoderPreset**](ModelEncoderPreset.md) | Gets or sets the encoder preset. | [optional] 
**DeinterlaceDoubleRate** | Pointer to **bool** | Gets or sets a value indicating whether the framerate is doubled when deinterlacing. | [optional] 
**DeinterlaceMethod** | Pointer to [**ModelModelDeinterlaceMethod**](ModelDeinterlaceMethod.md) | Gets or sets the deinterlace method. | [optional] 
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

### NewModelEncodingOptions

`func NewModelEncodingOptions() *ModelEncodingOptions`

NewModelEncodingOptions instantiates a new ModelEncodingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEncodingOptionsWithDefaults

`func NewModelEncodingOptionsWithDefaults() *ModelEncodingOptions`

NewModelEncodingOptionsWithDefaults instantiates a new ModelEncodingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingThreadCount

`func (o *ModelEncodingOptions) GetEncodingThreadCount() int32`

GetEncodingThreadCount returns the EncodingThreadCount field if non-nil, zero value otherwise.

### GetEncodingThreadCountOk

`func (o *ModelEncodingOptions) GetEncodingThreadCountOk() (*int32, bool)`

GetEncodingThreadCountOk returns a tuple with the EncodingThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingThreadCount

`func (o *ModelEncodingOptions) SetEncodingThreadCount(v int32)`

SetEncodingThreadCount sets EncodingThreadCount field to given value.

### HasEncodingThreadCount

`func (o *ModelEncodingOptions) HasEncodingThreadCount() bool`

HasEncodingThreadCount returns a boolean if a field has been set.

### GetTranscodingTempPath

`func (o *ModelEncodingOptions) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *ModelEncodingOptions) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *ModelEncodingOptions) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *ModelEncodingOptions) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### SetTranscodingTempPathNil

`func (o *ModelEncodingOptions) SetTranscodingTempPathNil(b bool)`

 SetTranscodingTempPathNil sets the value for TranscodingTempPath to be an explicit nil

### UnsetTranscodingTempPath
`func (o *ModelEncodingOptions) UnsetTranscodingTempPath()`

UnsetTranscodingTempPath ensures that no value is present for TranscodingTempPath, not even an explicit nil
### GetFallbackFontPath

`func (o *ModelEncodingOptions) GetFallbackFontPath() string`

GetFallbackFontPath returns the FallbackFontPath field if non-nil, zero value otherwise.

### GetFallbackFontPathOk

`func (o *ModelEncodingOptions) GetFallbackFontPathOk() (*string, bool)`

GetFallbackFontPathOk returns a tuple with the FallbackFontPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackFontPath

`func (o *ModelEncodingOptions) SetFallbackFontPath(v string)`

SetFallbackFontPath sets FallbackFontPath field to given value.

### HasFallbackFontPath

`func (o *ModelEncodingOptions) HasFallbackFontPath() bool`

HasFallbackFontPath returns a boolean if a field has been set.

### SetFallbackFontPathNil

`func (o *ModelEncodingOptions) SetFallbackFontPathNil(b bool)`

 SetFallbackFontPathNil sets the value for FallbackFontPath to be an explicit nil

### UnsetFallbackFontPath
`func (o *ModelEncodingOptions) UnsetFallbackFontPath()`

UnsetFallbackFontPath ensures that no value is present for FallbackFontPath, not even an explicit nil
### GetEnableFallbackFont

`func (o *ModelEncodingOptions) GetEnableFallbackFont() bool`

GetEnableFallbackFont returns the EnableFallbackFont field if non-nil, zero value otherwise.

### GetEnableFallbackFontOk

`func (o *ModelEncodingOptions) GetEnableFallbackFontOk() (*bool, bool)`

GetEnableFallbackFontOk returns a tuple with the EnableFallbackFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFallbackFont

`func (o *ModelEncodingOptions) SetEnableFallbackFont(v bool)`

SetEnableFallbackFont sets EnableFallbackFont field to given value.

### HasEnableFallbackFont

`func (o *ModelEncodingOptions) HasEnableFallbackFont() bool`

HasEnableFallbackFont returns a boolean if a field has been set.

### GetEnableAudioVbr

`func (o *ModelEncodingOptions) GetEnableAudioVbr() bool`

GetEnableAudioVbr returns the EnableAudioVbr field if non-nil, zero value otherwise.

### GetEnableAudioVbrOk

`func (o *ModelEncodingOptions) GetEnableAudioVbrOk() (*bool, bool)`

GetEnableAudioVbrOk returns a tuple with the EnableAudioVbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioVbr

`func (o *ModelEncodingOptions) SetEnableAudioVbr(v bool)`

SetEnableAudioVbr sets EnableAudioVbr field to given value.

### HasEnableAudioVbr

`func (o *ModelEncodingOptions) HasEnableAudioVbr() bool`

HasEnableAudioVbr returns a boolean if a field has been set.

### GetDownMixAudioBoost

`func (o *ModelEncodingOptions) GetDownMixAudioBoost() float64`

GetDownMixAudioBoost returns the DownMixAudioBoost field if non-nil, zero value otherwise.

### GetDownMixAudioBoostOk

`func (o *ModelEncodingOptions) GetDownMixAudioBoostOk() (*float64, bool)`

GetDownMixAudioBoostOk returns a tuple with the DownMixAudioBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixAudioBoost

`func (o *ModelEncodingOptions) SetDownMixAudioBoost(v float64)`

SetDownMixAudioBoost sets DownMixAudioBoost field to given value.

### HasDownMixAudioBoost

`func (o *ModelEncodingOptions) HasDownMixAudioBoost() bool`

HasDownMixAudioBoost returns a boolean if a field has been set.

### GetDownMixStereoAlgorithm

`func (o *ModelEncodingOptions) GetDownMixStereoAlgorithm() ModelModelDownMixStereoAlgorithms`

GetDownMixStereoAlgorithm returns the DownMixStereoAlgorithm field if non-nil, zero value otherwise.

### GetDownMixStereoAlgorithmOk

`func (o *ModelEncodingOptions) GetDownMixStereoAlgorithmOk() (*ModelModelDownMixStereoAlgorithms, bool)`

GetDownMixStereoAlgorithmOk returns a tuple with the DownMixStereoAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixStereoAlgorithm

`func (o *ModelEncodingOptions) SetDownMixStereoAlgorithm(v ModelModelDownMixStereoAlgorithms)`

SetDownMixStereoAlgorithm sets DownMixStereoAlgorithm field to given value.

### HasDownMixStereoAlgorithm

`func (o *ModelEncodingOptions) HasDownMixStereoAlgorithm() bool`

HasDownMixStereoAlgorithm returns a boolean if a field has been set.

### GetMaxMuxingQueueSize

`func (o *ModelEncodingOptions) GetMaxMuxingQueueSize() int32`

GetMaxMuxingQueueSize returns the MaxMuxingQueueSize field if non-nil, zero value otherwise.

### GetMaxMuxingQueueSizeOk

`func (o *ModelEncodingOptions) GetMaxMuxingQueueSizeOk() (*int32, bool)`

GetMaxMuxingQueueSizeOk returns a tuple with the MaxMuxingQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMuxingQueueSize

`func (o *ModelEncodingOptions) SetMaxMuxingQueueSize(v int32)`

SetMaxMuxingQueueSize sets MaxMuxingQueueSize field to given value.

### HasMaxMuxingQueueSize

`func (o *ModelEncodingOptions) HasMaxMuxingQueueSize() bool`

HasMaxMuxingQueueSize returns a boolean if a field has been set.

### GetEnableThrottling

`func (o *ModelEncodingOptions) GetEnableThrottling() bool`

GetEnableThrottling returns the EnableThrottling field if non-nil, zero value otherwise.

### GetEnableThrottlingOk

`func (o *ModelEncodingOptions) GetEnableThrottlingOk() (*bool, bool)`

GetEnableThrottlingOk returns a tuple with the EnableThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThrottling

`func (o *ModelEncodingOptions) SetEnableThrottling(v bool)`

SetEnableThrottling sets EnableThrottling field to given value.

### HasEnableThrottling

`func (o *ModelEncodingOptions) HasEnableThrottling() bool`

HasEnableThrottling returns a boolean if a field has been set.

### GetThrottleDelaySeconds

`func (o *ModelEncodingOptions) GetThrottleDelaySeconds() int32`

GetThrottleDelaySeconds returns the ThrottleDelaySeconds field if non-nil, zero value otherwise.

### GetThrottleDelaySecondsOk

`func (o *ModelEncodingOptions) GetThrottleDelaySecondsOk() (*int32, bool)`

GetThrottleDelaySecondsOk returns a tuple with the ThrottleDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDelaySeconds

`func (o *ModelEncodingOptions) SetThrottleDelaySeconds(v int32)`

SetThrottleDelaySeconds sets ThrottleDelaySeconds field to given value.

### HasThrottleDelaySeconds

`func (o *ModelEncodingOptions) HasThrottleDelaySeconds() bool`

HasThrottleDelaySeconds returns a boolean if a field has been set.

### GetEnableSegmentDeletion

`func (o *ModelEncodingOptions) GetEnableSegmentDeletion() bool`

GetEnableSegmentDeletion returns the EnableSegmentDeletion field if non-nil, zero value otherwise.

### GetEnableSegmentDeletionOk

`func (o *ModelEncodingOptions) GetEnableSegmentDeletionOk() (*bool, bool)`

GetEnableSegmentDeletionOk returns a tuple with the EnableSegmentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSegmentDeletion

`func (o *ModelEncodingOptions) SetEnableSegmentDeletion(v bool)`

SetEnableSegmentDeletion sets EnableSegmentDeletion field to given value.

### HasEnableSegmentDeletion

`func (o *ModelEncodingOptions) HasEnableSegmentDeletion() bool`

HasEnableSegmentDeletion returns a boolean if a field has been set.

### GetSegmentKeepSeconds

`func (o *ModelEncodingOptions) GetSegmentKeepSeconds() int32`

GetSegmentKeepSeconds returns the SegmentKeepSeconds field if non-nil, zero value otherwise.

### GetSegmentKeepSecondsOk

`func (o *ModelEncodingOptions) GetSegmentKeepSecondsOk() (*int32, bool)`

GetSegmentKeepSecondsOk returns a tuple with the SegmentKeepSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeepSeconds

`func (o *ModelEncodingOptions) SetSegmentKeepSeconds(v int32)`

SetSegmentKeepSeconds sets SegmentKeepSeconds field to given value.

### HasSegmentKeepSeconds

`func (o *ModelEncodingOptions) HasSegmentKeepSeconds() bool`

HasSegmentKeepSeconds returns a boolean if a field has been set.

### GetHardwareAccelerationType

`func (o *ModelEncodingOptions) GetHardwareAccelerationType() ModelModelHardwareAccelerationType`

GetHardwareAccelerationType returns the HardwareAccelerationType field if non-nil, zero value otherwise.

### GetHardwareAccelerationTypeOk

`func (o *ModelEncodingOptions) GetHardwareAccelerationTypeOk() (*ModelModelHardwareAccelerationType, bool)`

GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationType

`func (o *ModelEncodingOptions) SetHardwareAccelerationType(v ModelModelHardwareAccelerationType)`

SetHardwareAccelerationType sets HardwareAccelerationType field to given value.

### HasHardwareAccelerationType

`func (o *ModelEncodingOptions) HasHardwareAccelerationType() bool`

HasHardwareAccelerationType returns a boolean if a field has been set.

### GetEncoderAppPath

`func (o *ModelEncodingOptions) GetEncoderAppPath() string`

GetEncoderAppPath returns the EncoderAppPath field if non-nil, zero value otherwise.

### GetEncoderAppPathOk

`func (o *ModelEncodingOptions) GetEncoderAppPathOk() (*string, bool)`

GetEncoderAppPathOk returns a tuple with the EncoderAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPath

`func (o *ModelEncodingOptions) SetEncoderAppPath(v string)`

SetEncoderAppPath sets EncoderAppPath field to given value.

### HasEncoderAppPath

`func (o *ModelEncodingOptions) HasEncoderAppPath() bool`

HasEncoderAppPath returns a boolean if a field has been set.

### SetEncoderAppPathNil

`func (o *ModelEncodingOptions) SetEncoderAppPathNil(b bool)`

 SetEncoderAppPathNil sets the value for EncoderAppPath to be an explicit nil

### UnsetEncoderAppPath
`func (o *ModelEncodingOptions) UnsetEncoderAppPath()`

UnsetEncoderAppPath ensures that no value is present for EncoderAppPath, not even an explicit nil
### GetEncoderAppPathDisplay

`func (o *ModelEncodingOptions) GetEncoderAppPathDisplay() string`

GetEncoderAppPathDisplay returns the EncoderAppPathDisplay field if non-nil, zero value otherwise.

### GetEncoderAppPathDisplayOk

`func (o *ModelEncodingOptions) GetEncoderAppPathDisplayOk() (*string, bool)`

GetEncoderAppPathDisplayOk returns a tuple with the EncoderAppPathDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPathDisplay

`func (o *ModelEncodingOptions) SetEncoderAppPathDisplay(v string)`

SetEncoderAppPathDisplay sets EncoderAppPathDisplay field to given value.

### HasEncoderAppPathDisplay

`func (o *ModelEncodingOptions) HasEncoderAppPathDisplay() bool`

HasEncoderAppPathDisplay returns a boolean if a field has been set.

### SetEncoderAppPathDisplayNil

`func (o *ModelEncodingOptions) SetEncoderAppPathDisplayNil(b bool)`

 SetEncoderAppPathDisplayNil sets the value for EncoderAppPathDisplay to be an explicit nil

### UnsetEncoderAppPathDisplay
`func (o *ModelEncodingOptions) UnsetEncoderAppPathDisplay()`

UnsetEncoderAppPathDisplay ensures that no value is present for EncoderAppPathDisplay, not even an explicit nil
### GetVaapiDevice

`func (o *ModelEncodingOptions) GetVaapiDevice() string`

GetVaapiDevice returns the VaapiDevice field if non-nil, zero value otherwise.

### GetVaapiDeviceOk

`func (o *ModelEncodingOptions) GetVaapiDeviceOk() (*string, bool)`

GetVaapiDeviceOk returns a tuple with the VaapiDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaapiDevice

`func (o *ModelEncodingOptions) SetVaapiDevice(v string)`

SetVaapiDevice sets VaapiDevice field to given value.

### HasVaapiDevice

`func (o *ModelEncodingOptions) HasVaapiDevice() bool`

HasVaapiDevice returns a boolean if a field has been set.

### SetVaapiDeviceNil

`func (o *ModelEncodingOptions) SetVaapiDeviceNil(b bool)`

 SetVaapiDeviceNil sets the value for VaapiDevice to be an explicit nil

### UnsetVaapiDevice
`func (o *ModelEncodingOptions) UnsetVaapiDevice()`

UnsetVaapiDevice ensures that no value is present for VaapiDevice, not even an explicit nil
### GetQsvDevice

`func (o *ModelEncodingOptions) GetQsvDevice() string`

GetQsvDevice returns the QsvDevice field if non-nil, zero value otherwise.

### GetQsvDeviceOk

`func (o *ModelEncodingOptions) GetQsvDeviceOk() (*string, bool)`

GetQsvDeviceOk returns a tuple with the QsvDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQsvDevice

`func (o *ModelEncodingOptions) SetQsvDevice(v string)`

SetQsvDevice sets QsvDevice field to given value.

### HasQsvDevice

`func (o *ModelEncodingOptions) HasQsvDevice() bool`

HasQsvDevice returns a boolean if a field has been set.

### SetQsvDeviceNil

`func (o *ModelEncodingOptions) SetQsvDeviceNil(b bool)`

 SetQsvDeviceNil sets the value for QsvDevice to be an explicit nil

### UnsetQsvDevice
`func (o *ModelEncodingOptions) UnsetQsvDevice()`

UnsetQsvDevice ensures that no value is present for QsvDevice, not even an explicit nil
### GetEnableTonemapping

`func (o *ModelEncodingOptions) GetEnableTonemapping() bool`

GetEnableTonemapping returns the EnableTonemapping field if non-nil, zero value otherwise.

### GetEnableTonemappingOk

`func (o *ModelEncodingOptions) GetEnableTonemappingOk() (*bool, bool)`

GetEnableTonemappingOk returns a tuple with the EnableTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTonemapping

`func (o *ModelEncodingOptions) SetEnableTonemapping(v bool)`

SetEnableTonemapping sets EnableTonemapping field to given value.

### HasEnableTonemapping

`func (o *ModelEncodingOptions) HasEnableTonemapping() bool`

HasEnableTonemapping returns a boolean if a field has been set.

### GetEnableVppTonemapping

`func (o *ModelEncodingOptions) GetEnableVppTonemapping() bool`

GetEnableVppTonemapping returns the EnableVppTonemapping field if non-nil, zero value otherwise.

### GetEnableVppTonemappingOk

`func (o *ModelEncodingOptions) GetEnableVppTonemappingOk() (*bool, bool)`

GetEnableVppTonemappingOk returns a tuple with the EnableVppTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVppTonemapping

`func (o *ModelEncodingOptions) SetEnableVppTonemapping(v bool)`

SetEnableVppTonemapping sets EnableVppTonemapping field to given value.

### HasEnableVppTonemapping

`func (o *ModelEncodingOptions) HasEnableVppTonemapping() bool`

HasEnableVppTonemapping returns a boolean if a field has been set.

### GetEnableVideoToolboxTonemapping

`func (o *ModelEncodingOptions) GetEnableVideoToolboxTonemapping() bool`

GetEnableVideoToolboxTonemapping returns the EnableVideoToolboxTonemapping field if non-nil, zero value otherwise.

### GetEnableVideoToolboxTonemappingOk

`func (o *ModelEncodingOptions) GetEnableVideoToolboxTonemappingOk() (*bool, bool)`

GetEnableVideoToolboxTonemappingOk returns a tuple with the EnableVideoToolboxTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoToolboxTonemapping

`func (o *ModelEncodingOptions) SetEnableVideoToolboxTonemapping(v bool)`

SetEnableVideoToolboxTonemapping sets EnableVideoToolboxTonemapping field to given value.

### HasEnableVideoToolboxTonemapping

`func (o *ModelEncodingOptions) HasEnableVideoToolboxTonemapping() bool`

HasEnableVideoToolboxTonemapping returns a boolean if a field has been set.

### GetTonemappingAlgorithm

`func (o *ModelEncodingOptions) GetTonemappingAlgorithm() ModelModelTonemappingAlgorithm`

GetTonemappingAlgorithm returns the TonemappingAlgorithm field if non-nil, zero value otherwise.

### GetTonemappingAlgorithmOk

`func (o *ModelEncodingOptions) GetTonemappingAlgorithmOk() (*ModelModelTonemappingAlgorithm, bool)`

GetTonemappingAlgorithmOk returns a tuple with the TonemappingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingAlgorithm

`func (o *ModelEncodingOptions) SetTonemappingAlgorithm(v ModelModelTonemappingAlgorithm)`

SetTonemappingAlgorithm sets TonemappingAlgorithm field to given value.

### HasTonemappingAlgorithm

`func (o *ModelEncodingOptions) HasTonemappingAlgorithm() bool`

HasTonemappingAlgorithm returns a boolean if a field has been set.

### GetTonemappingMode

`func (o *ModelEncodingOptions) GetTonemappingMode() ModelModelTonemappingMode`

GetTonemappingMode returns the TonemappingMode field if non-nil, zero value otherwise.

### GetTonemappingModeOk

`func (o *ModelEncodingOptions) GetTonemappingModeOk() (*ModelModelTonemappingMode, bool)`

GetTonemappingModeOk returns a tuple with the TonemappingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingMode

`func (o *ModelEncodingOptions) SetTonemappingMode(v ModelModelTonemappingMode)`

SetTonemappingMode sets TonemappingMode field to given value.

### HasTonemappingMode

`func (o *ModelEncodingOptions) HasTonemappingMode() bool`

HasTonemappingMode returns a boolean if a field has been set.

### GetTonemappingRange

`func (o *ModelEncodingOptions) GetTonemappingRange() ModelModelTonemappingRange`

GetTonemappingRange returns the TonemappingRange field if non-nil, zero value otherwise.

### GetTonemappingRangeOk

`func (o *ModelEncodingOptions) GetTonemappingRangeOk() (*ModelModelTonemappingRange, bool)`

GetTonemappingRangeOk returns a tuple with the TonemappingRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingRange

`func (o *ModelEncodingOptions) SetTonemappingRange(v ModelModelTonemappingRange)`

SetTonemappingRange sets TonemappingRange field to given value.

### HasTonemappingRange

`func (o *ModelEncodingOptions) HasTonemappingRange() bool`

HasTonemappingRange returns a boolean if a field has been set.

### GetTonemappingDesat

`func (o *ModelEncodingOptions) GetTonemappingDesat() float64`

GetTonemappingDesat returns the TonemappingDesat field if non-nil, zero value otherwise.

### GetTonemappingDesatOk

`func (o *ModelEncodingOptions) GetTonemappingDesatOk() (*float64, bool)`

GetTonemappingDesatOk returns a tuple with the TonemappingDesat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingDesat

`func (o *ModelEncodingOptions) SetTonemappingDesat(v float64)`

SetTonemappingDesat sets TonemappingDesat field to given value.

### HasTonemappingDesat

`func (o *ModelEncodingOptions) HasTonemappingDesat() bool`

HasTonemappingDesat returns a boolean if a field has been set.

### GetTonemappingPeak

`func (o *ModelEncodingOptions) GetTonemappingPeak() float64`

GetTonemappingPeak returns the TonemappingPeak field if non-nil, zero value otherwise.

### GetTonemappingPeakOk

`func (o *ModelEncodingOptions) GetTonemappingPeakOk() (*float64, bool)`

GetTonemappingPeakOk returns a tuple with the TonemappingPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingPeak

`func (o *ModelEncodingOptions) SetTonemappingPeak(v float64)`

SetTonemappingPeak sets TonemappingPeak field to given value.

### HasTonemappingPeak

`func (o *ModelEncodingOptions) HasTonemappingPeak() bool`

HasTonemappingPeak returns a boolean if a field has been set.

### GetTonemappingParam

`func (o *ModelEncodingOptions) GetTonemappingParam() float64`

GetTonemappingParam returns the TonemappingParam field if non-nil, zero value otherwise.

### GetTonemappingParamOk

`func (o *ModelEncodingOptions) GetTonemappingParamOk() (*float64, bool)`

GetTonemappingParamOk returns a tuple with the TonemappingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingParam

`func (o *ModelEncodingOptions) SetTonemappingParam(v float64)`

SetTonemappingParam sets TonemappingParam field to given value.

### HasTonemappingParam

`func (o *ModelEncodingOptions) HasTonemappingParam() bool`

HasTonemappingParam returns a boolean if a field has been set.

### GetVppTonemappingBrightness

`func (o *ModelEncodingOptions) GetVppTonemappingBrightness() float64`

GetVppTonemappingBrightness returns the VppTonemappingBrightness field if non-nil, zero value otherwise.

### GetVppTonemappingBrightnessOk

`func (o *ModelEncodingOptions) GetVppTonemappingBrightnessOk() (*float64, bool)`

GetVppTonemappingBrightnessOk returns a tuple with the VppTonemappingBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingBrightness

`func (o *ModelEncodingOptions) SetVppTonemappingBrightness(v float64)`

SetVppTonemappingBrightness sets VppTonemappingBrightness field to given value.

### HasVppTonemappingBrightness

`func (o *ModelEncodingOptions) HasVppTonemappingBrightness() bool`

HasVppTonemappingBrightness returns a boolean if a field has been set.

### GetVppTonemappingContrast

`func (o *ModelEncodingOptions) GetVppTonemappingContrast() float64`

GetVppTonemappingContrast returns the VppTonemappingContrast field if non-nil, zero value otherwise.

### GetVppTonemappingContrastOk

`func (o *ModelEncodingOptions) GetVppTonemappingContrastOk() (*float64, bool)`

GetVppTonemappingContrastOk returns a tuple with the VppTonemappingContrast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingContrast

`func (o *ModelEncodingOptions) SetVppTonemappingContrast(v float64)`

SetVppTonemappingContrast sets VppTonemappingContrast field to given value.

### HasVppTonemappingContrast

`func (o *ModelEncodingOptions) HasVppTonemappingContrast() bool`

HasVppTonemappingContrast returns a boolean if a field has been set.

### GetH264Crf

`func (o *ModelEncodingOptions) GetH264Crf() int32`

GetH264Crf returns the H264Crf field if non-nil, zero value otherwise.

### GetH264CrfOk

`func (o *ModelEncodingOptions) GetH264CrfOk() (*int32, bool)`

GetH264CrfOk returns a tuple with the H264Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH264Crf

`func (o *ModelEncodingOptions) SetH264Crf(v int32)`

SetH264Crf sets H264Crf field to given value.

### HasH264Crf

`func (o *ModelEncodingOptions) HasH264Crf() bool`

HasH264Crf returns a boolean if a field has been set.

### GetH265Crf

`func (o *ModelEncodingOptions) GetH265Crf() int32`

GetH265Crf returns the H265Crf field if non-nil, zero value otherwise.

### GetH265CrfOk

`func (o *ModelEncodingOptions) GetH265CrfOk() (*int32, bool)`

GetH265CrfOk returns a tuple with the H265Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH265Crf

`func (o *ModelEncodingOptions) SetH265Crf(v int32)`

SetH265Crf sets H265Crf field to given value.

### HasH265Crf

`func (o *ModelEncodingOptions) HasH265Crf() bool`

HasH265Crf returns a boolean if a field has been set.

### GetEncoderPreset

`func (o *ModelEncodingOptions) GetEncoderPreset() ModelModelEncoderPreset`

GetEncoderPreset returns the EncoderPreset field if non-nil, zero value otherwise.

### GetEncoderPresetOk

`func (o *ModelEncodingOptions) GetEncoderPresetOk() (*ModelModelEncoderPreset, bool)`

GetEncoderPresetOk returns a tuple with the EncoderPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPreset

`func (o *ModelEncodingOptions) SetEncoderPreset(v ModelModelEncoderPreset)`

SetEncoderPreset sets EncoderPreset field to given value.

### HasEncoderPreset

`func (o *ModelEncodingOptions) HasEncoderPreset() bool`

HasEncoderPreset returns a boolean if a field has been set.

### SetEncoderPresetNil

`func (o *ModelEncodingOptions) SetEncoderPresetNil(b bool)`

 SetEncoderPresetNil sets the value for EncoderPreset to be an explicit nil

### UnsetEncoderPreset
`func (o *ModelEncodingOptions) UnsetEncoderPreset()`

UnsetEncoderPreset ensures that no value is present for EncoderPreset, not even an explicit nil
### GetDeinterlaceDoubleRate

`func (o *ModelEncodingOptions) GetDeinterlaceDoubleRate() bool`

GetDeinterlaceDoubleRate returns the DeinterlaceDoubleRate field if non-nil, zero value otherwise.

### GetDeinterlaceDoubleRateOk

`func (o *ModelEncodingOptions) GetDeinterlaceDoubleRateOk() (*bool, bool)`

GetDeinterlaceDoubleRateOk returns a tuple with the DeinterlaceDoubleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceDoubleRate

`func (o *ModelEncodingOptions) SetDeinterlaceDoubleRate(v bool)`

SetDeinterlaceDoubleRate sets DeinterlaceDoubleRate field to given value.

### HasDeinterlaceDoubleRate

`func (o *ModelEncodingOptions) HasDeinterlaceDoubleRate() bool`

HasDeinterlaceDoubleRate returns a boolean if a field has been set.

### GetDeinterlaceMethod

`func (o *ModelEncodingOptions) GetDeinterlaceMethod() ModelModelDeinterlaceMethod`

GetDeinterlaceMethod returns the DeinterlaceMethod field if non-nil, zero value otherwise.

### GetDeinterlaceMethodOk

`func (o *ModelEncodingOptions) GetDeinterlaceMethodOk() (*ModelModelDeinterlaceMethod, bool)`

GetDeinterlaceMethodOk returns a tuple with the DeinterlaceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceMethod

`func (o *ModelEncodingOptions) SetDeinterlaceMethod(v ModelModelDeinterlaceMethod)`

SetDeinterlaceMethod sets DeinterlaceMethod field to given value.

### HasDeinterlaceMethod

`func (o *ModelEncodingOptions) HasDeinterlaceMethod() bool`

HasDeinterlaceMethod returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10Hevc

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10Hevc() bool`

GetEnableDecodingColorDepth10Hevc returns the EnableDecodingColorDepth10Hevc field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10HevcOk

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10HevcOk() (*bool, bool)`

GetEnableDecodingColorDepth10HevcOk returns a tuple with the EnableDecodingColorDepth10Hevc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Hevc

`func (o *ModelEncodingOptions) SetEnableDecodingColorDepth10Hevc(v bool)`

SetEnableDecodingColorDepth10Hevc sets EnableDecodingColorDepth10Hevc field to given value.

### HasEnableDecodingColorDepth10Hevc

`func (o *ModelEncodingOptions) HasEnableDecodingColorDepth10Hevc() bool`

HasEnableDecodingColorDepth10Hevc returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10Vp9

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10Vp9() bool`

GetEnableDecodingColorDepth10Vp9 returns the EnableDecodingColorDepth10Vp9 field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10Vp9Ok

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10Vp9Ok() (*bool, bool)`

GetEnableDecodingColorDepth10Vp9Ok returns a tuple with the EnableDecodingColorDepth10Vp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Vp9

`func (o *ModelEncodingOptions) SetEnableDecodingColorDepth10Vp9(v bool)`

SetEnableDecodingColorDepth10Vp9 sets EnableDecodingColorDepth10Vp9 field to given value.

### HasEnableDecodingColorDepth10Vp9

`func (o *ModelEncodingOptions) HasEnableDecodingColorDepth10Vp9() bool`

HasEnableDecodingColorDepth10Vp9 returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10HevcRext

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10HevcRext() bool`

GetEnableDecodingColorDepth10HevcRext returns the EnableDecodingColorDepth10HevcRext field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10HevcRextOk

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth10HevcRextOk() (*bool, bool)`

GetEnableDecodingColorDepth10HevcRextOk returns a tuple with the EnableDecodingColorDepth10HevcRext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10HevcRext

`func (o *ModelEncodingOptions) SetEnableDecodingColorDepth10HevcRext(v bool)`

SetEnableDecodingColorDepth10HevcRext sets EnableDecodingColorDepth10HevcRext field to given value.

### HasEnableDecodingColorDepth10HevcRext

`func (o *ModelEncodingOptions) HasEnableDecodingColorDepth10HevcRext() bool`

HasEnableDecodingColorDepth10HevcRext returns a boolean if a field has been set.

### GetEnableDecodingColorDepth12HevcRext

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth12HevcRext() bool`

GetEnableDecodingColorDepth12HevcRext returns the EnableDecodingColorDepth12HevcRext field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth12HevcRextOk

`func (o *ModelEncodingOptions) GetEnableDecodingColorDepth12HevcRextOk() (*bool, bool)`

GetEnableDecodingColorDepth12HevcRextOk returns a tuple with the EnableDecodingColorDepth12HevcRext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth12HevcRext

`func (o *ModelEncodingOptions) SetEnableDecodingColorDepth12HevcRext(v bool)`

SetEnableDecodingColorDepth12HevcRext sets EnableDecodingColorDepth12HevcRext field to given value.

### HasEnableDecodingColorDepth12HevcRext

`func (o *ModelEncodingOptions) HasEnableDecodingColorDepth12HevcRext() bool`

HasEnableDecodingColorDepth12HevcRext returns a boolean if a field has been set.

### GetEnableEnhancedNvdecDecoder

`func (o *ModelEncodingOptions) GetEnableEnhancedNvdecDecoder() bool`

GetEnableEnhancedNvdecDecoder returns the EnableEnhancedNvdecDecoder field if non-nil, zero value otherwise.

### GetEnableEnhancedNvdecDecoderOk

`func (o *ModelEncodingOptions) GetEnableEnhancedNvdecDecoderOk() (*bool, bool)`

GetEnableEnhancedNvdecDecoderOk returns a tuple with the EnableEnhancedNvdecDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEnhancedNvdecDecoder

`func (o *ModelEncodingOptions) SetEnableEnhancedNvdecDecoder(v bool)`

SetEnableEnhancedNvdecDecoder sets EnableEnhancedNvdecDecoder field to given value.

### HasEnableEnhancedNvdecDecoder

`func (o *ModelEncodingOptions) HasEnableEnhancedNvdecDecoder() bool`

HasEnableEnhancedNvdecDecoder returns a boolean if a field has been set.

### GetPreferSystemNativeHwDecoder

`func (o *ModelEncodingOptions) GetPreferSystemNativeHwDecoder() bool`

GetPreferSystemNativeHwDecoder returns the PreferSystemNativeHwDecoder field if non-nil, zero value otherwise.

### GetPreferSystemNativeHwDecoderOk

`func (o *ModelEncodingOptions) GetPreferSystemNativeHwDecoderOk() (*bool, bool)`

GetPreferSystemNativeHwDecoderOk returns a tuple with the PreferSystemNativeHwDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferSystemNativeHwDecoder

`func (o *ModelEncodingOptions) SetPreferSystemNativeHwDecoder(v bool)`

SetPreferSystemNativeHwDecoder sets PreferSystemNativeHwDecoder field to given value.

### HasPreferSystemNativeHwDecoder

`func (o *ModelEncodingOptions) HasPreferSystemNativeHwDecoder() bool`

HasPreferSystemNativeHwDecoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerH264HwEncoder

`func (o *ModelEncodingOptions) GetEnableIntelLowPowerH264HwEncoder() bool`

GetEnableIntelLowPowerH264HwEncoder returns the EnableIntelLowPowerH264HwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerH264HwEncoderOk

`func (o *ModelEncodingOptions) GetEnableIntelLowPowerH264HwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerH264HwEncoderOk returns a tuple with the EnableIntelLowPowerH264HwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerH264HwEncoder

`func (o *ModelEncodingOptions) SetEnableIntelLowPowerH264HwEncoder(v bool)`

SetEnableIntelLowPowerH264HwEncoder sets EnableIntelLowPowerH264HwEncoder field to given value.

### HasEnableIntelLowPowerH264HwEncoder

`func (o *ModelEncodingOptions) HasEnableIntelLowPowerH264HwEncoder() bool`

HasEnableIntelLowPowerH264HwEncoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerHevcHwEncoder

`func (o *ModelEncodingOptions) GetEnableIntelLowPowerHevcHwEncoder() bool`

GetEnableIntelLowPowerHevcHwEncoder returns the EnableIntelLowPowerHevcHwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerHevcHwEncoderOk

`func (o *ModelEncodingOptions) GetEnableIntelLowPowerHevcHwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerHevcHwEncoderOk returns a tuple with the EnableIntelLowPowerHevcHwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerHevcHwEncoder

`func (o *ModelEncodingOptions) SetEnableIntelLowPowerHevcHwEncoder(v bool)`

SetEnableIntelLowPowerHevcHwEncoder sets EnableIntelLowPowerHevcHwEncoder field to given value.

### HasEnableIntelLowPowerHevcHwEncoder

`func (o *ModelEncodingOptions) HasEnableIntelLowPowerHevcHwEncoder() bool`

HasEnableIntelLowPowerHevcHwEncoder returns a boolean if a field has been set.

### GetEnableHardwareEncoding

`func (o *ModelEncodingOptions) GetEnableHardwareEncoding() bool`

GetEnableHardwareEncoding returns the EnableHardwareEncoding field if non-nil, zero value otherwise.

### GetEnableHardwareEncodingOk

`func (o *ModelEncodingOptions) GetEnableHardwareEncodingOk() (*bool, bool)`

GetEnableHardwareEncodingOk returns a tuple with the EnableHardwareEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHardwareEncoding

`func (o *ModelEncodingOptions) SetEnableHardwareEncoding(v bool)`

SetEnableHardwareEncoding sets EnableHardwareEncoding field to given value.

### HasEnableHardwareEncoding

`func (o *ModelEncodingOptions) HasEnableHardwareEncoding() bool`

HasEnableHardwareEncoding returns a boolean if a field has been set.

### GetAllowHevcEncoding

`func (o *ModelEncodingOptions) GetAllowHevcEncoding() bool`

GetAllowHevcEncoding returns the AllowHevcEncoding field if non-nil, zero value otherwise.

### GetAllowHevcEncodingOk

`func (o *ModelEncodingOptions) GetAllowHevcEncodingOk() (*bool, bool)`

GetAllowHevcEncodingOk returns a tuple with the AllowHevcEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHevcEncoding

`func (o *ModelEncodingOptions) SetAllowHevcEncoding(v bool)`

SetAllowHevcEncoding sets AllowHevcEncoding field to given value.

### HasAllowHevcEncoding

`func (o *ModelEncodingOptions) HasAllowHevcEncoding() bool`

HasAllowHevcEncoding returns a boolean if a field has been set.

### GetAllowAv1Encoding

`func (o *ModelEncodingOptions) GetAllowAv1Encoding() bool`

GetAllowAv1Encoding returns the AllowAv1Encoding field if non-nil, zero value otherwise.

### GetAllowAv1EncodingOk

`func (o *ModelEncodingOptions) GetAllowAv1EncodingOk() (*bool, bool)`

GetAllowAv1EncodingOk returns a tuple with the AllowAv1Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAv1Encoding

`func (o *ModelEncodingOptions) SetAllowAv1Encoding(v bool)`

SetAllowAv1Encoding sets AllowAv1Encoding field to given value.

### HasAllowAv1Encoding

`func (o *ModelEncodingOptions) HasAllowAv1Encoding() bool`

HasAllowAv1Encoding returns a boolean if a field has been set.

### GetEnableSubtitleExtraction

`func (o *ModelEncodingOptions) GetEnableSubtitleExtraction() bool`

GetEnableSubtitleExtraction returns the EnableSubtitleExtraction field if non-nil, zero value otherwise.

### GetEnableSubtitleExtractionOk

`func (o *ModelEncodingOptions) GetEnableSubtitleExtractionOk() (*bool, bool)`

GetEnableSubtitleExtractionOk returns a tuple with the EnableSubtitleExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleExtraction

`func (o *ModelEncodingOptions) SetEnableSubtitleExtraction(v bool)`

SetEnableSubtitleExtraction sets EnableSubtitleExtraction field to given value.

### HasEnableSubtitleExtraction

`func (o *ModelEncodingOptions) HasEnableSubtitleExtraction() bool`

HasEnableSubtitleExtraction returns a boolean if a field has been set.

### GetHardwareDecodingCodecs

`func (o *ModelEncodingOptions) GetHardwareDecodingCodecs() []string`

GetHardwareDecodingCodecs returns the HardwareDecodingCodecs field if non-nil, zero value otherwise.

### GetHardwareDecodingCodecsOk

`func (o *ModelEncodingOptions) GetHardwareDecodingCodecsOk() (*[]string, bool)`

GetHardwareDecodingCodecsOk returns a tuple with the HardwareDecodingCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareDecodingCodecs

`func (o *ModelEncodingOptions) SetHardwareDecodingCodecs(v []string)`

SetHardwareDecodingCodecs sets HardwareDecodingCodecs field to given value.

### HasHardwareDecodingCodecs

`func (o *ModelEncodingOptions) HasHardwareDecodingCodecs() bool`

HasHardwareDecodingCodecs returns a boolean if a field has been set.

### SetHardwareDecodingCodecsNil

`func (o *ModelEncodingOptions) SetHardwareDecodingCodecsNil(b bool)`

 SetHardwareDecodingCodecsNil sets the value for HardwareDecodingCodecs to be an explicit nil

### UnsetHardwareDecodingCodecs
`func (o *ModelEncodingOptions) UnsetHardwareDecodingCodecs()`

UnsetHardwareDecodingCodecs ensures that no value is present for HardwareDecodingCodecs, not even an explicit nil
### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *ModelEncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() []string`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if non-nil, zero value otherwise.

### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk

`func (o *ModelEncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk() (*[]string, bool)`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk returns a tuple with the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *ModelEncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions(v []string)`

SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions sets AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field to given value.

### HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *ModelEncodingOptions) HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() bool`

HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns a boolean if a field has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil

`func (o *ModelEncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil(b bool)`

 SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil sets the value for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions to be an explicit nil

### UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions
`func (o *ModelEncodingOptions) UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions()`

UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions ensures that no value is present for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


