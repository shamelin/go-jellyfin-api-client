# JellyfinLiveTvOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuideDays** | Pointer to **NullableInt32** |  | [optional] 
**RecordingPath** | Pointer to **NullableString** |  | [optional] 
**MovieRecordingPath** | Pointer to **NullableString** |  | [optional] 
**SeriesRecordingPath** | Pointer to **NullableString** |  | [optional] 
**EnableRecordingSubfolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalAudioWithEncodedRecordings** | Pointer to **bool** |  | [optional] 
**TunerHosts** | Pointer to [**[]JellyfinJellyfinTunerHostInfo**](JellyfinJellyfinTunerHostInfo.md) |  | [optional] 
**ListingProviders** | Pointer to [**[]JellyfinJellyfinListingsProviderInfo**](JellyfinJellyfinListingsProviderInfo.md) |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**MediaLocationsCreated** | Pointer to **[]string** |  | [optional] 
**RecordingPostProcessor** | Pointer to **NullableString** |  | [optional] 
**RecordingPostProcessorArguments** | Pointer to **NullableString** |  | [optional] 
**SaveRecordingNFO** | Pointer to **bool** |  | [optional] 
**SaveRecordingImages** | Pointer to **bool** |  | [optional] 

## Methods

### NewJellyfinLiveTvOptions

`func NewJellyfinLiveTvOptions() *JellyfinLiveTvOptions`

NewJellyfinLiveTvOptions instantiates a new JellyfinLiveTvOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLiveTvOptionsWithDefaults

`func NewJellyfinLiveTvOptionsWithDefaults() *JellyfinLiveTvOptions`

NewJellyfinLiveTvOptionsWithDefaults instantiates a new JellyfinLiveTvOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuideDays

`func (o *JellyfinLiveTvOptions) GetGuideDays() int32`

GetGuideDays returns the GuideDays field if non-nil, zero value otherwise.

### GetGuideDaysOk

`func (o *JellyfinLiveTvOptions) GetGuideDaysOk() (*int32, bool)`

GetGuideDaysOk returns a tuple with the GuideDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuideDays

`func (o *JellyfinLiveTvOptions) SetGuideDays(v int32)`

SetGuideDays sets GuideDays field to given value.

### HasGuideDays

`func (o *JellyfinLiveTvOptions) HasGuideDays() bool`

HasGuideDays returns a boolean if a field has been set.

### SetGuideDaysNil

`func (o *JellyfinLiveTvOptions) SetGuideDaysNil(b bool)`

 SetGuideDaysNil sets the value for GuideDays to be an explicit nil

### UnsetGuideDays
`func (o *JellyfinLiveTvOptions) UnsetGuideDays()`

UnsetGuideDays ensures that no value is present for GuideDays, not even an explicit nil
### GetRecordingPath

`func (o *JellyfinLiveTvOptions) GetRecordingPath() string`

GetRecordingPath returns the RecordingPath field if non-nil, zero value otherwise.

### GetRecordingPathOk

`func (o *JellyfinLiveTvOptions) GetRecordingPathOk() (*string, bool)`

GetRecordingPathOk returns a tuple with the RecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPath

`func (o *JellyfinLiveTvOptions) SetRecordingPath(v string)`

SetRecordingPath sets RecordingPath field to given value.

### HasRecordingPath

`func (o *JellyfinLiveTvOptions) HasRecordingPath() bool`

HasRecordingPath returns a boolean if a field has been set.

### SetRecordingPathNil

`func (o *JellyfinLiveTvOptions) SetRecordingPathNil(b bool)`

 SetRecordingPathNil sets the value for RecordingPath to be an explicit nil

### UnsetRecordingPath
`func (o *JellyfinLiveTvOptions) UnsetRecordingPath()`

UnsetRecordingPath ensures that no value is present for RecordingPath, not even an explicit nil
### GetMovieRecordingPath

`func (o *JellyfinLiveTvOptions) GetMovieRecordingPath() string`

GetMovieRecordingPath returns the MovieRecordingPath field if non-nil, zero value otherwise.

### GetMovieRecordingPathOk

`func (o *JellyfinLiveTvOptions) GetMovieRecordingPathOk() (*string, bool)`

GetMovieRecordingPathOk returns a tuple with the MovieRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieRecordingPath

`func (o *JellyfinLiveTvOptions) SetMovieRecordingPath(v string)`

SetMovieRecordingPath sets MovieRecordingPath field to given value.

### HasMovieRecordingPath

`func (o *JellyfinLiveTvOptions) HasMovieRecordingPath() bool`

HasMovieRecordingPath returns a boolean if a field has been set.

### SetMovieRecordingPathNil

`func (o *JellyfinLiveTvOptions) SetMovieRecordingPathNil(b bool)`

 SetMovieRecordingPathNil sets the value for MovieRecordingPath to be an explicit nil

### UnsetMovieRecordingPath
`func (o *JellyfinLiveTvOptions) UnsetMovieRecordingPath()`

UnsetMovieRecordingPath ensures that no value is present for MovieRecordingPath, not even an explicit nil
### GetSeriesRecordingPath

`func (o *JellyfinLiveTvOptions) GetSeriesRecordingPath() string`

GetSeriesRecordingPath returns the SeriesRecordingPath field if non-nil, zero value otherwise.

### GetSeriesRecordingPathOk

`func (o *JellyfinLiveTvOptions) GetSeriesRecordingPathOk() (*string, bool)`

GetSeriesRecordingPathOk returns a tuple with the SeriesRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesRecordingPath

`func (o *JellyfinLiveTvOptions) SetSeriesRecordingPath(v string)`

SetSeriesRecordingPath sets SeriesRecordingPath field to given value.

### HasSeriesRecordingPath

`func (o *JellyfinLiveTvOptions) HasSeriesRecordingPath() bool`

HasSeriesRecordingPath returns a boolean if a field has been set.

### SetSeriesRecordingPathNil

`func (o *JellyfinLiveTvOptions) SetSeriesRecordingPathNil(b bool)`

 SetSeriesRecordingPathNil sets the value for SeriesRecordingPath to be an explicit nil

### UnsetSeriesRecordingPath
`func (o *JellyfinLiveTvOptions) UnsetSeriesRecordingPath()`

UnsetSeriesRecordingPath ensures that no value is present for SeriesRecordingPath, not even an explicit nil
### GetEnableRecordingSubfolders

`func (o *JellyfinLiveTvOptions) GetEnableRecordingSubfolders() bool`

GetEnableRecordingSubfolders returns the EnableRecordingSubfolders field if non-nil, zero value otherwise.

### GetEnableRecordingSubfoldersOk

`func (o *JellyfinLiveTvOptions) GetEnableRecordingSubfoldersOk() (*bool, bool)`

GetEnableRecordingSubfoldersOk returns a tuple with the EnableRecordingSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecordingSubfolders

`func (o *JellyfinLiveTvOptions) SetEnableRecordingSubfolders(v bool)`

SetEnableRecordingSubfolders sets EnableRecordingSubfolders field to given value.

### HasEnableRecordingSubfolders

`func (o *JellyfinLiveTvOptions) HasEnableRecordingSubfolders() bool`

HasEnableRecordingSubfolders returns a boolean if a field has been set.

### GetEnableOriginalAudioWithEncodedRecordings

`func (o *JellyfinLiveTvOptions) GetEnableOriginalAudioWithEncodedRecordings() bool`

GetEnableOriginalAudioWithEncodedRecordings returns the EnableOriginalAudioWithEncodedRecordings field if non-nil, zero value otherwise.

### GetEnableOriginalAudioWithEncodedRecordingsOk

`func (o *JellyfinLiveTvOptions) GetEnableOriginalAudioWithEncodedRecordingsOk() (*bool, bool)`

GetEnableOriginalAudioWithEncodedRecordingsOk returns a tuple with the EnableOriginalAudioWithEncodedRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalAudioWithEncodedRecordings

`func (o *JellyfinLiveTvOptions) SetEnableOriginalAudioWithEncodedRecordings(v bool)`

SetEnableOriginalAudioWithEncodedRecordings sets EnableOriginalAudioWithEncodedRecordings field to given value.

### HasEnableOriginalAudioWithEncodedRecordings

`func (o *JellyfinLiveTvOptions) HasEnableOriginalAudioWithEncodedRecordings() bool`

HasEnableOriginalAudioWithEncodedRecordings returns a boolean if a field has been set.

### GetTunerHosts

`func (o *JellyfinLiveTvOptions) GetTunerHosts() []JellyfinJellyfinTunerHostInfo`

GetTunerHosts returns the TunerHosts field if non-nil, zero value otherwise.

### GetTunerHostsOk

`func (o *JellyfinLiveTvOptions) GetTunerHostsOk() (*[]JellyfinJellyfinTunerHostInfo, bool)`

GetTunerHostsOk returns a tuple with the TunerHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerHosts

`func (o *JellyfinLiveTvOptions) SetTunerHosts(v []JellyfinJellyfinTunerHostInfo)`

SetTunerHosts sets TunerHosts field to given value.

### HasTunerHosts

`func (o *JellyfinLiveTvOptions) HasTunerHosts() bool`

HasTunerHosts returns a boolean if a field has been set.

### SetTunerHostsNil

`func (o *JellyfinLiveTvOptions) SetTunerHostsNil(b bool)`

 SetTunerHostsNil sets the value for TunerHosts to be an explicit nil

### UnsetTunerHosts
`func (o *JellyfinLiveTvOptions) UnsetTunerHosts()`

UnsetTunerHosts ensures that no value is present for TunerHosts, not even an explicit nil
### GetListingProviders

`func (o *JellyfinLiveTvOptions) GetListingProviders() []JellyfinJellyfinListingsProviderInfo`

GetListingProviders returns the ListingProviders field if non-nil, zero value otherwise.

### GetListingProvidersOk

`func (o *JellyfinLiveTvOptions) GetListingProvidersOk() (*[]JellyfinJellyfinListingsProviderInfo, bool)`

GetListingProvidersOk returns a tuple with the ListingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingProviders

`func (o *JellyfinLiveTvOptions) SetListingProviders(v []JellyfinJellyfinListingsProviderInfo)`

SetListingProviders sets ListingProviders field to given value.

### HasListingProviders

`func (o *JellyfinLiveTvOptions) HasListingProviders() bool`

HasListingProviders returns a boolean if a field has been set.

### SetListingProvidersNil

`func (o *JellyfinLiveTvOptions) SetListingProvidersNil(b bool)`

 SetListingProvidersNil sets the value for ListingProviders to be an explicit nil

### UnsetListingProviders
`func (o *JellyfinLiveTvOptions) UnsetListingProviders()`

UnsetListingProviders ensures that no value is present for ListingProviders, not even an explicit nil
### GetPrePaddingSeconds

`func (o *JellyfinLiveTvOptions) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *JellyfinLiveTvOptions) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *JellyfinLiveTvOptions) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *JellyfinLiveTvOptions) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *JellyfinLiveTvOptions) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *JellyfinLiveTvOptions) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *JellyfinLiveTvOptions) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *JellyfinLiveTvOptions) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetMediaLocationsCreated

`func (o *JellyfinLiveTvOptions) GetMediaLocationsCreated() []string`

GetMediaLocationsCreated returns the MediaLocationsCreated field if non-nil, zero value otherwise.

### GetMediaLocationsCreatedOk

`func (o *JellyfinLiveTvOptions) GetMediaLocationsCreatedOk() (*[]string, bool)`

GetMediaLocationsCreatedOk returns a tuple with the MediaLocationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaLocationsCreated

`func (o *JellyfinLiveTvOptions) SetMediaLocationsCreated(v []string)`

SetMediaLocationsCreated sets MediaLocationsCreated field to given value.

### HasMediaLocationsCreated

`func (o *JellyfinLiveTvOptions) HasMediaLocationsCreated() bool`

HasMediaLocationsCreated returns a boolean if a field has been set.

### SetMediaLocationsCreatedNil

`func (o *JellyfinLiveTvOptions) SetMediaLocationsCreatedNil(b bool)`

 SetMediaLocationsCreatedNil sets the value for MediaLocationsCreated to be an explicit nil

### UnsetMediaLocationsCreated
`func (o *JellyfinLiveTvOptions) UnsetMediaLocationsCreated()`

UnsetMediaLocationsCreated ensures that no value is present for MediaLocationsCreated, not even an explicit nil
### GetRecordingPostProcessor

`func (o *JellyfinLiveTvOptions) GetRecordingPostProcessor() string`

GetRecordingPostProcessor returns the RecordingPostProcessor field if non-nil, zero value otherwise.

### GetRecordingPostProcessorOk

`func (o *JellyfinLiveTvOptions) GetRecordingPostProcessorOk() (*string, bool)`

GetRecordingPostProcessorOk returns a tuple with the RecordingPostProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessor

`func (o *JellyfinLiveTvOptions) SetRecordingPostProcessor(v string)`

SetRecordingPostProcessor sets RecordingPostProcessor field to given value.

### HasRecordingPostProcessor

`func (o *JellyfinLiveTvOptions) HasRecordingPostProcessor() bool`

HasRecordingPostProcessor returns a boolean if a field has been set.

### SetRecordingPostProcessorNil

`func (o *JellyfinLiveTvOptions) SetRecordingPostProcessorNil(b bool)`

 SetRecordingPostProcessorNil sets the value for RecordingPostProcessor to be an explicit nil

### UnsetRecordingPostProcessor
`func (o *JellyfinLiveTvOptions) UnsetRecordingPostProcessor()`

UnsetRecordingPostProcessor ensures that no value is present for RecordingPostProcessor, not even an explicit nil
### GetRecordingPostProcessorArguments

`func (o *JellyfinLiveTvOptions) GetRecordingPostProcessorArguments() string`

GetRecordingPostProcessorArguments returns the RecordingPostProcessorArguments field if non-nil, zero value otherwise.

### GetRecordingPostProcessorArgumentsOk

`func (o *JellyfinLiveTvOptions) GetRecordingPostProcessorArgumentsOk() (*string, bool)`

GetRecordingPostProcessorArgumentsOk returns a tuple with the RecordingPostProcessorArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessorArguments

`func (o *JellyfinLiveTvOptions) SetRecordingPostProcessorArguments(v string)`

SetRecordingPostProcessorArguments sets RecordingPostProcessorArguments field to given value.

### HasRecordingPostProcessorArguments

`func (o *JellyfinLiveTvOptions) HasRecordingPostProcessorArguments() bool`

HasRecordingPostProcessorArguments returns a boolean if a field has been set.

### SetRecordingPostProcessorArgumentsNil

`func (o *JellyfinLiveTvOptions) SetRecordingPostProcessorArgumentsNil(b bool)`

 SetRecordingPostProcessorArgumentsNil sets the value for RecordingPostProcessorArguments to be an explicit nil

### UnsetRecordingPostProcessorArguments
`func (o *JellyfinLiveTvOptions) UnsetRecordingPostProcessorArguments()`

UnsetRecordingPostProcessorArguments ensures that no value is present for RecordingPostProcessorArguments, not even an explicit nil
### GetSaveRecordingNFO

`func (o *JellyfinLiveTvOptions) GetSaveRecordingNFO() bool`

GetSaveRecordingNFO returns the SaveRecordingNFO field if non-nil, zero value otherwise.

### GetSaveRecordingNFOOk

`func (o *JellyfinLiveTvOptions) GetSaveRecordingNFOOk() (*bool, bool)`

GetSaveRecordingNFOOk returns a tuple with the SaveRecordingNFO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveRecordingNFO

`func (o *JellyfinLiveTvOptions) SetSaveRecordingNFO(v bool)`

SetSaveRecordingNFO sets SaveRecordingNFO field to given value.

### HasSaveRecordingNFO

`func (o *JellyfinLiveTvOptions) HasSaveRecordingNFO() bool`

HasSaveRecordingNFO returns a boolean if a field has been set.

### GetSaveRecordingImages

`func (o *JellyfinLiveTvOptions) GetSaveRecordingImages() bool`

GetSaveRecordingImages returns the SaveRecordingImages field if non-nil, zero value otherwise.

### GetSaveRecordingImagesOk

`func (o *JellyfinLiveTvOptions) GetSaveRecordingImagesOk() (*bool, bool)`

GetSaveRecordingImagesOk returns a tuple with the SaveRecordingImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveRecordingImages

`func (o *JellyfinLiveTvOptions) SetSaveRecordingImages(v bool)`

SetSaveRecordingImages sets SaveRecordingImages field to given value.

### HasSaveRecordingImages

`func (o *JellyfinLiveTvOptions) HasSaveRecordingImages() bool`

HasSaveRecordingImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


