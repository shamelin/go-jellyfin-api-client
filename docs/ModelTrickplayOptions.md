# ModelTrickplayOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableHwAcceleration** | Pointer to **bool** | Gets or sets a value indicating whether or not to use HW acceleration. | [optional] 
**EnableHwEncoding** | Pointer to **bool** | Gets or sets a value indicating whether or not to use HW accelerated MJPEG encoding. | [optional] 
**EnableKeyFrameOnlyExtraction** | Pointer to **bool** | Gets or sets a value indicating whether to only extract key frames.  Significantly faster, but is not compatible with all decoders and/or video files. | [optional] 
**ScanBehavior** | Pointer to [**ModelModelTrickplayScanBehavior**](ModelTrickplayScanBehavior.md) | Gets or sets the behavior used by trickplay provider on library scan/update. | [optional] 
**ProcessPriority** | Pointer to [**ModelModelProcessPriorityClass**](ModelProcessPriorityClass.md) | Gets or sets the process priority for the ffmpeg process. | [optional] 
**Interval** | Pointer to **int32** | Gets or sets the interval, in ms, between each new trickplay image. | [optional] 
**WidthResolutions** | Pointer to **[]int32** | Gets or sets the target width resolutions, in px, to generates preview images for. | [optional] 
**TileWidth** | Pointer to **int32** | Gets or sets number of tile images to allow in X dimension. | [optional] 
**TileHeight** | Pointer to **int32** | Gets or sets number of tile images to allow in Y dimension. | [optional] 
**Qscale** | Pointer to **int32** | Gets or sets the ffmpeg output quality level. | [optional] 
**JpegQuality** | Pointer to **int32** | Gets or sets the jpeg quality to use for image tiles. | [optional] 
**ProcessThreads** | Pointer to **int32** | Gets or sets the number of threads to be used by ffmpeg. | [optional] 

## Methods

### NewModelTrickplayOptions

`func NewModelTrickplayOptions() *ModelTrickplayOptions`

NewModelTrickplayOptions instantiates a new ModelTrickplayOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTrickplayOptionsWithDefaults

`func NewModelTrickplayOptionsWithDefaults() *ModelTrickplayOptions`

NewModelTrickplayOptionsWithDefaults instantiates a new ModelTrickplayOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableHwAcceleration

`func (o *ModelTrickplayOptions) GetEnableHwAcceleration() bool`

GetEnableHwAcceleration returns the EnableHwAcceleration field if non-nil, zero value otherwise.

### GetEnableHwAccelerationOk

`func (o *ModelTrickplayOptions) GetEnableHwAccelerationOk() (*bool, bool)`

GetEnableHwAccelerationOk returns a tuple with the EnableHwAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHwAcceleration

`func (o *ModelTrickplayOptions) SetEnableHwAcceleration(v bool)`

SetEnableHwAcceleration sets EnableHwAcceleration field to given value.

### HasEnableHwAcceleration

`func (o *ModelTrickplayOptions) HasEnableHwAcceleration() bool`

HasEnableHwAcceleration returns a boolean if a field has been set.

### GetEnableHwEncoding

`func (o *ModelTrickplayOptions) GetEnableHwEncoding() bool`

GetEnableHwEncoding returns the EnableHwEncoding field if non-nil, zero value otherwise.

### GetEnableHwEncodingOk

`func (o *ModelTrickplayOptions) GetEnableHwEncodingOk() (*bool, bool)`

GetEnableHwEncodingOk returns a tuple with the EnableHwEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHwEncoding

`func (o *ModelTrickplayOptions) SetEnableHwEncoding(v bool)`

SetEnableHwEncoding sets EnableHwEncoding field to given value.

### HasEnableHwEncoding

`func (o *ModelTrickplayOptions) HasEnableHwEncoding() bool`

HasEnableHwEncoding returns a boolean if a field has been set.

### GetEnableKeyFrameOnlyExtraction

`func (o *ModelTrickplayOptions) GetEnableKeyFrameOnlyExtraction() bool`

GetEnableKeyFrameOnlyExtraction returns the EnableKeyFrameOnlyExtraction field if non-nil, zero value otherwise.

### GetEnableKeyFrameOnlyExtractionOk

`func (o *ModelTrickplayOptions) GetEnableKeyFrameOnlyExtractionOk() (*bool, bool)`

GetEnableKeyFrameOnlyExtractionOk returns a tuple with the EnableKeyFrameOnlyExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKeyFrameOnlyExtraction

`func (o *ModelTrickplayOptions) SetEnableKeyFrameOnlyExtraction(v bool)`

SetEnableKeyFrameOnlyExtraction sets EnableKeyFrameOnlyExtraction field to given value.

### HasEnableKeyFrameOnlyExtraction

`func (o *ModelTrickplayOptions) HasEnableKeyFrameOnlyExtraction() bool`

HasEnableKeyFrameOnlyExtraction returns a boolean if a field has been set.

### GetScanBehavior

`func (o *ModelTrickplayOptions) GetScanBehavior() ModelModelTrickplayScanBehavior`

GetScanBehavior returns the ScanBehavior field if non-nil, zero value otherwise.

### GetScanBehaviorOk

`func (o *ModelTrickplayOptions) GetScanBehaviorOk() (*ModelModelTrickplayScanBehavior, bool)`

GetScanBehaviorOk returns a tuple with the ScanBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanBehavior

`func (o *ModelTrickplayOptions) SetScanBehavior(v ModelModelTrickplayScanBehavior)`

SetScanBehavior sets ScanBehavior field to given value.

### HasScanBehavior

`func (o *ModelTrickplayOptions) HasScanBehavior() bool`

HasScanBehavior returns a boolean if a field has been set.

### GetProcessPriority

`func (o *ModelTrickplayOptions) GetProcessPriority() ModelModelProcessPriorityClass`

GetProcessPriority returns the ProcessPriority field if non-nil, zero value otherwise.

### GetProcessPriorityOk

`func (o *ModelTrickplayOptions) GetProcessPriorityOk() (*ModelModelProcessPriorityClass, bool)`

GetProcessPriorityOk returns a tuple with the ProcessPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessPriority

`func (o *ModelTrickplayOptions) SetProcessPriority(v ModelModelProcessPriorityClass)`

SetProcessPriority sets ProcessPriority field to given value.

### HasProcessPriority

`func (o *ModelTrickplayOptions) HasProcessPriority() bool`

HasProcessPriority returns a boolean if a field has been set.

### GetInterval

`func (o *ModelTrickplayOptions) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ModelTrickplayOptions) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ModelTrickplayOptions) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ModelTrickplayOptions) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetWidthResolutions

`func (o *ModelTrickplayOptions) GetWidthResolutions() []int32`

GetWidthResolutions returns the WidthResolutions field if non-nil, zero value otherwise.

### GetWidthResolutionsOk

`func (o *ModelTrickplayOptions) GetWidthResolutionsOk() (*[]int32, bool)`

GetWidthResolutionsOk returns a tuple with the WidthResolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthResolutions

`func (o *ModelTrickplayOptions) SetWidthResolutions(v []int32)`

SetWidthResolutions sets WidthResolutions field to given value.

### HasWidthResolutions

`func (o *ModelTrickplayOptions) HasWidthResolutions() bool`

HasWidthResolutions returns a boolean if a field has been set.

### GetTileWidth

`func (o *ModelTrickplayOptions) GetTileWidth() int32`

GetTileWidth returns the TileWidth field if non-nil, zero value otherwise.

### GetTileWidthOk

`func (o *ModelTrickplayOptions) GetTileWidthOk() (*int32, bool)`

GetTileWidthOk returns a tuple with the TileWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileWidth

`func (o *ModelTrickplayOptions) SetTileWidth(v int32)`

SetTileWidth sets TileWidth field to given value.

### HasTileWidth

`func (o *ModelTrickplayOptions) HasTileWidth() bool`

HasTileWidth returns a boolean if a field has been set.

### GetTileHeight

`func (o *ModelTrickplayOptions) GetTileHeight() int32`

GetTileHeight returns the TileHeight field if non-nil, zero value otherwise.

### GetTileHeightOk

`func (o *ModelTrickplayOptions) GetTileHeightOk() (*int32, bool)`

GetTileHeightOk returns a tuple with the TileHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileHeight

`func (o *ModelTrickplayOptions) SetTileHeight(v int32)`

SetTileHeight sets TileHeight field to given value.

### HasTileHeight

`func (o *ModelTrickplayOptions) HasTileHeight() bool`

HasTileHeight returns a boolean if a field has been set.

### GetQscale

`func (o *ModelTrickplayOptions) GetQscale() int32`

GetQscale returns the Qscale field if non-nil, zero value otherwise.

### GetQscaleOk

`func (o *ModelTrickplayOptions) GetQscaleOk() (*int32, bool)`

GetQscaleOk returns a tuple with the Qscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQscale

`func (o *ModelTrickplayOptions) SetQscale(v int32)`

SetQscale sets Qscale field to given value.

### HasQscale

`func (o *ModelTrickplayOptions) HasQscale() bool`

HasQscale returns a boolean if a field has been set.

### GetJpegQuality

`func (o *ModelTrickplayOptions) GetJpegQuality() int32`

GetJpegQuality returns the JpegQuality field if non-nil, zero value otherwise.

### GetJpegQualityOk

`func (o *ModelTrickplayOptions) GetJpegQualityOk() (*int32, bool)`

GetJpegQualityOk returns a tuple with the JpegQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJpegQuality

`func (o *ModelTrickplayOptions) SetJpegQuality(v int32)`

SetJpegQuality sets JpegQuality field to given value.

### HasJpegQuality

`func (o *ModelTrickplayOptions) HasJpegQuality() bool`

HasJpegQuality returns a boolean if a field has been set.

### GetProcessThreads

`func (o *ModelTrickplayOptions) GetProcessThreads() int32`

GetProcessThreads returns the ProcessThreads field if non-nil, zero value otherwise.

### GetProcessThreadsOk

`func (o *ModelTrickplayOptions) GetProcessThreadsOk() (*int32, bool)`

GetProcessThreadsOk returns a tuple with the ProcessThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessThreads

`func (o *ModelTrickplayOptions) SetProcessThreads(v int32)`

SetProcessThreads sets ProcessThreads field to given value.

### HasProcessThreads

`func (o *ModelTrickplayOptions) HasProcessThreads() bool`

HasProcessThreads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


