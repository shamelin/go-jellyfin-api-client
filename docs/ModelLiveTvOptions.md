# ModelLiveTvOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuideDays** | Pointer to **NullableInt32** |  | [optional] 
**RecordingPath** | Pointer to **NullableString** |  | [optional] 
**MovieRecordingPath** | Pointer to **NullableString** |  | [optional] 
**SeriesRecordingPath** | Pointer to **NullableString** |  | [optional] 
**EnableRecordingSubfolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalAudioWithEncodedRecordings** | Pointer to **bool** |  | [optional] 
**TunerHosts** | Pointer to [**[]ModelModelTunerHostInfo**](ModelModelTunerHostInfo.md) |  | [optional] 
**ListingProviders** | Pointer to [**[]ModelModelListingsProviderInfo**](ModelModelListingsProviderInfo.md) |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**MediaLocationsCreated** | Pointer to **[]string** |  | [optional] 
**RecordingPostProcessor** | Pointer to **NullableString** |  | [optional] 
**RecordingPostProcessorArguments** | Pointer to **NullableString** |  | [optional] 
**SaveRecordingNFO** | Pointer to **bool** |  | [optional] 
**SaveRecordingImages** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelLiveTvOptions

`func NewModelLiveTvOptions() *ModelLiveTvOptions`

NewModelLiveTvOptions instantiates a new ModelLiveTvOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvOptionsWithDefaults

`func NewModelLiveTvOptionsWithDefaults() *ModelLiveTvOptions`

NewModelLiveTvOptionsWithDefaults instantiates a new ModelLiveTvOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuideDays

`func (o *ModelLiveTvOptions) GetGuideDays() int32`

GetGuideDays returns the GuideDays field if non-nil, zero value otherwise.

### GetGuideDaysOk

`func (o *ModelLiveTvOptions) GetGuideDaysOk() (*int32, bool)`

GetGuideDaysOk returns a tuple with the GuideDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuideDays

`func (o *ModelLiveTvOptions) SetGuideDays(v int32)`

SetGuideDays sets GuideDays field to given value.

### HasGuideDays

`func (o *ModelLiveTvOptions) HasGuideDays() bool`

HasGuideDays returns a boolean if a field has been set.

### SetGuideDaysNil

`func (o *ModelLiveTvOptions) SetGuideDaysNil(b bool)`

 SetGuideDaysNil sets the value for GuideDays to be an explicit nil

### UnsetGuideDays
`func (o *ModelLiveTvOptions) UnsetGuideDays()`

UnsetGuideDays ensures that no value is present for GuideDays, not even an explicit nil
### GetRecordingPath

`func (o *ModelLiveTvOptions) GetRecordingPath() string`

GetRecordingPath returns the RecordingPath field if non-nil, zero value otherwise.

### GetRecordingPathOk

`func (o *ModelLiveTvOptions) GetRecordingPathOk() (*string, bool)`

GetRecordingPathOk returns a tuple with the RecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPath

`func (o *ModelLiveTvOptions) SetRecordingPath(v string)`

SetRecordingPath sets RecordingPath field to given value.

### HasRecordingPath

`func (o *ModelLiveTvOptions) HasRecordingPath() bool`

HasRecordingPath returns a boolean if a field has been set.

### SetRecordingPathNil

`func (o *ModelLiveTvOptions) SetRecordingPathNil(b bool)`

 SetRecordingPathNil sets the value for RecordingPath to be an explicit nil

### UnsetRecordingPath
`func (o *ModelLiveTvOptions) UnsetRecordingPath()`

UnsetRecordingPath ensures that no value is present for RecordingPath, not even an explicit nil
### GetMovieRecordingPath

`func (o *ModelLiveTvOptions) GetMovieRecordingPath() string`

GetMovieRecordingPath returns the MovieRecordingPath field if non-nil, zero value otherwise.

### GetMovieRecordingPathOk

`func (o *ModelLiveTvOptions) GetMovieRecordingPathOk() (*string, bool)`

GetMovieRecordingPathOk returns a tuple with the MovieRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieRecordingPath

`func (o *ModelLiveTvOptions) SetMovieRecordingPath(v string)`

SetMovieRecordingPath sets MovieRecordingPath field to given value.

### HasMovieRecordingPath

`func (o *ModelLiveTvOptions) HasMovieRecordingPath() bool`

HasMovieRecordingPath returns a boolean if a field has been set.

### SetMovieRecordingPathNil

`func (o *ModelLiveTvOptions) SetMovieRecordingPathNil(b bool)`

 SetMovieRecordingPathNil sets the value for MovieRecordingPath to be an explicit nil

### UnsetMovieRecordingPath
`func (o *ModelLiveTvOptions) UnsetMovieRecordingPath()`

UnsetMovieRecordingPath ensures that no value is present for MovieRecordingPath, not even an explicit nil
### GetSeriesRecordingPath

`func (o *ModelLiveTvOptions) GetSeriesRecordingPath() string`

GetSeriesRecordingPath returns the SeriesRecordingPath field if non-nil, zero value otherwise.

### GetSeriesRecordingPathOk

`func (o *ModelLiveTvOptions) GetSeriesRecordingPathOk() (*string, bool)`

GetSeriesRecordingPathOk returns a tuple with the SeriesRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesRecordingPath

`func (o *ModelLiveTvOptions) SetSeriesRecordingPath(v string)`

SetSeriesRecordingPath sets SeriesRecordingPath field to given value.

### HasSeriesRecordingPath

`func (o *ModelLiveTvOptions) HasSeriesRecordingPath() bool`

HasSeriesRecordingPath returns a boolean if a field has been set.

### SetSeriesRecordingPathNil

`func (o *ModelLiveTvOptions) SetSeriesRecordingPathNil(b bool)`

 SetSeriesRecordingPathNil sets the value for SeriesRecordingPath to be an explicit nil

### UnsetSeriesRecordingPath
`func (o *ModelLiveTvOptions) UnsetSeriesRecordingPath()`

UnsetSeriesRecordingPath ensures that no value is present for SeriesRecordingPath, not even an explicit nil
### GetEnableRecordingSubfolders

`func (o *ModelLiveTvOptions) GetEnableRecordingSubfolders() bool`

GetEnableRecordingSubfolders returns the EnableRecordingSubfolders field if non-nil, zero value otherwise.

### GetEnableRecordingSubfoldersOk

`func (o *ModelLiveTvOptions) GetEnableRecordingSubfoldersOk() (*bool, bool)`

GetEnableRecordingSubfoldersOk returns a tuple with the EnableRecordingSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecordingSubfolders

`func (o *ModelLiveTvOptions) SetEnableRecordingSubfolders(v bool)`

SetEnableRecordingSubfolders sets EnableRecordingSubfolders field to given value.

### HasEnableRecordingSubfolders

`func (o *ModelLiveTvOptions) HasEnableRecordingSubfolders() bool`

HasEnableRecordingSubfolders returns a boolean if a field has been set.

### GetEnableOriginalAudioWithEncodedRecordings

`func (o *ModelLiveTvOptions) GetEnableOriginalAudioWithEncodedRecordings() bool`

GetEnableOriginalAudioWithEncodedRecordings returns the EnableOriginalAudioWithEncodedRecordings field if non-nil, zero value otherwise.

### GetEnableOriginalAudioWithEncodedRecordingsOk

`func (o *ModelLiveTvOptions) GetEnableOriginalAudioWithEncodedRecordingsOk() (*bool, bool)`

GetEnableOriginalAudioWithEncodedRecordingsOk returns a tuple with the EnableOriginalAudioWithEncodedRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalAudioWithEncodedRecordings

`func (o *ModelLiveTvOptions) SetEnableOriginalAudioWithEncodedRecordings(v bool)`

SetEnableOriginalAudioWithEncodedRecordings sets EnableOriginalAudioWithEncodedRecordings field to given value.

### HasEnableOriginalAudioWithEncodedRecordings

`func (o *ModelLiveTvOptions) HasEnableOriginalAudioWithEncodedRecordings() bool`

HasEnableOriginalAudioWithEncodedRecordings returns a boolean if a field has been set.

### GetTunerHosts

`func (o *ModelLiveTvOptions) GetTunerHosts() []ModelModelTunerHostInfo`

GetTunerHosts returns the TunerHosts field if non-nil, zero value otherwise.

### GetTunerHostsOk

`func (o *ModelLiveTvOptions) GetTunerHostsOk() (*[]ModelModelTunerHostInfo, bool)`

GetTunerHostsOk returns a tuple with the TunerHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerHosts

`func (o *ModelLiveTvOptions) SetTunerHosts(v []ModelModelTunerHostInfo)`

SetTunerHosts sets TunerHosts field to given value.

### HasTunerHosts

`func (o *ModelLiveTvOptions) HasTunerHosts() bool`

HasTunerHosts returns a boolean if a field has been set.

### SetTunerHostsNil

`func (o *ModelLiveTvOptions) SetTunerHostsNil(b bool)`

 SetTunerHostsNil sets the value for TunerHosts to be an explicit nil

### UnsetTunerHosts
`func (o *ModelLiveTvOptions) UnsetTunerHosts()`

UnsetTunerHosts ensures that no value is present for TunerHosts, not even an explicit nil
### GetListingProviders

`func (o *ModelLiveTvOptions) GetListingProviders() []ModelModelListingsProviderInfo`

GetListingProviders returns the ListingProviders field if non-nil, zero value otherwise.

### GetListingProvidersOk

`func (o *ModelLiveTvOptions) GetListingProvidersOk() (*[]ModelModelListingsProviderInfo, bool)`

GetListingProvidersOk returns a tuple with the ListingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingProviders

`func (o *ModelLiveTvOptions) SetListingProviders(v []ModelModelListingsProviderInfo)`

SetListingProviders sets ListingProviders field to given value.

### HasListingProviders

`func (o *ModelLiveTvOptions) HasListingProviders() bool`

HasListingProviders returns a boolean if a field has been set.

### SetListingProvidersNil

`func (o *ModelLiveTvOptions) SetListingProvidersNil(b bool)`

 SetListingProvidersNil sets the value for ListingProviders to be an explicit nil

### UnsetListingProviders
`func (o *ModelLiveTvOptions) UnsetListingProviders()`

UnsetListingProviders ensures that no value is present for ListingProviders, not even an explicit nil
### GetPrePaddingSeconds

`func (o *ModelLiveTvOptions) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelLiveTvOptions) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelLiveTvOptions) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelLiveTvOptions) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelLiveTvOptions) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelLiveTvOptions) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelLiveTvOptions) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelLiveTvOptions) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetMediaLocationsCreated

`func (o *ModelLiveTvOptions) GetMediaLocationsCreated() []string`

GetMediaLocationsCreated returns the MediaLocationsCreated field if non-nil, zero value otherwise.

### GetMediaLocationsCreatedOk

`func (o *ModelLiveTvOptions) GetMediaLocationsCreatedOk() (*[]string, bool)`

GetMediaLocationsCreatedOk returns a tuple with the MediaLocationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaLocationsCreated

`func (o *ModelLiveTvOptions) SetMediaLocationsCreated(v []string)`

SetMediaLocationsCreated sets MediaLocationsCreated field to given value.

### HasMediaLocationsCreated

`func (o *ModelLiveTvOptions) HasMediaLocationsCreated() bool`

HasMediaLocationsCreated returns a boolean if a field has been set.

### SetMediaLocationsCreatedNil

`func (o *ModelLiveTvOptions) SetMediaLocationsCreatedNil(b bool)`

 SetMediaLocationsCreatedNil sets the value for MediaLocationsCreated to be an explicit nil

### UnsetMediaLocationsCreated
`func (o *ModelLiveTvOptions) UnsetMediaLocationsCreated()`

UnsetMediaLocationsCreated ensures that no value is present for MediaLocationsCreated, not even an explicit nil
### GetRecordingPostProcessor

`func (o *ModelLiveTvOptions) GetRecordingPostProcessor() string`

GetRecordingPostProcessor returns the RecordingPostProcessor field if non-nil, zero value otherwise.

### GetRecordingPostProcessorOk

`func (o *ModelLiveTvOptions) GetRecordingPostProcessorOk() (*string, bool)`

GetRecordingPostProcessorOk returns a tuple with the RecordingPostProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessor

`func (o *ModelLiveTvOptions) SetRecordingPostProcessor(v string)`

SetRecordingPostProcessor sets RecordingPostProcessor field to given value.

### HasRecordingPostProcessor

`func (o *ModelLiveTvOptions) HasRecordingPostProcessor() bool`

HasRecordingPostProcessor returns a boolean if a field has been set.

### SetRecordingPostProcessorNil

`func (o *ModelLiveTvOptions) SetRecordingPostProcessorNil(b bool)`

 SetRecordingPostProcessorNil sets the value for RecordingPostProcessor to be an explicit nil

### UnsetRecordingPostProcessor
`func (o *ModelLiveTvOptions) UnsetRecordingPostProcessor()`

UnsetRecordingPostProcessor ensures that no value is present for RecordingPostProcessor, not even an explicit nil
### GetRecordingPostProcessorArguments

`func (o *ModelLiveTvOptions) GetRecordingPostProcessorArguments() string`

GetRecordingPostProcessorArguments returns the RecordingPostProcessorArguments field if non-nil, zero value otherwise.

### GetRecordingPostProcessorArgumentsOk

`func (o *ModelLiveTvOptions) GetRecordingPostProcessorArgumentsOk() (*string, bool)`

GetRecordingPostProcessorArgumentsOk returns a tuple with the RecordingPostProcessorArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessorArguments

`func (o *ModelLiveTvOptions) SetRecordingPostProcessorArguments(v string)`

SetRecordingPostProcessorArguments sets RecordingPostProcessorArguments field to given value.

### HasRecordingPostProcessorArguments

`func (o *ModelLiveTvOptions) HasRecordingPostProcessorArguments() bool`

HasRecordingPostProcessorArguments returns a boolean if a field has been set.

### SetRecordingPostProcessorArgumentsNil

`func (o *ModelLiveTvOptions) SetRecordingPostProcessorArgumentsNil(b bool)`

 SetRecordingPostProcessorArgumentsNil sets the value for RecordingPostProcessorArguments to be an explicit nil

### UnsetRecordingPostProcessorArguments
`func (o *ModelLiveTvOptions) UnsetRecordingPostProcessorArguments()`

UnsetRecordingPostProcessorArguments ensures that no value is present for RecordingPostProcessorArguments, not even an explicit nil
### GetSaveRecordingNFO

`func (o *ModelLiveTvOptions) GetSaveRecordingNFO() bool`

GetSaveRecordingNFO returns the SaveRecordingNFO field if non-nil, zero value otherwise.

### GetSaveRecordingNFOOk

`func (o *ModelLiveTvOptions) GetSaveRecordingNFOOk() (*bool, bool)`

GetSaveRecordingNFOOk returns a tuple with the SaveRecordingNFO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveRecordingNFO

`func (o *ModelLiveTvOptions) SetSaveRecordingNFO(v bool)`

SetSaveRecordingNFO sets SaveRecordingNFO field to given value.

### HasSaveRecordingNFO

`func (o *ModelLiveTvOptions) HasSaveRecordingNFO() bool`

HasSaveRecordingNFO returns a boolean if a field has been set.

### GetSaveRecordingImages

`func (o *ModelLiveTvOptions) GetSaveRecordingImages() bool`

GetSaveRecordingImages returns the SaveRecordingImages field if non-nil, zero value otherwise.

### GetSaveRecordingImagesOk

`func (o *ModelLiveTvOptions) GetSaveRecordingImagesOk() (*bool, bool)`

GetSaveRecordingImagesOk returns a tuple with the SaveRecordingImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveRecordingImages

`func (o *ModelLiveTvOptions) SetSaveRecordingImages(v bool)`

SetSaveRecordingImages sets SaveRecordingImages field to given value.

### HasSaveRecordingImages

`func (o *ModelLiveTvOptions) HasSaveRecordingImages() bool`

HasSaveRecordingImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


