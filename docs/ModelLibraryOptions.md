# ModelLibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableLUFSScan** | Pointer to **bool** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**EnableTrickplayImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractTrickplayImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]ModelModelMediaPathInfo**](ModelModelMediaPathInfo.md) |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**EnableInternetProviders** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedExtrasTitles** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedEpisodeInfos** | Pointer to **bool** |  | [optional] 
**AutomaticRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **NullableString** | Gets or sets the preferred metadata language. | [optional] 
**MetadataCountryCode** | Pointer to **NullableString** | Gets or sets the metadata country code. | [optional] 
**SeasonZeroDisplayName** | Pointer to **string** |  | [optional] 
**MetadataSavers** | Pointer to **[]string** |  | [optional] 
**DisabledLocalMetadataReaders** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledSubtitleFetchers** | Pointer to **[]string** |  | [optional] 
**SubtitleFetcherOrder** | Pointer to **[]string** |  | [optional] 
**DisabledMediaSegmentProviders** | Pointer to **[]string** |  | [optional] 
**MediaSegmentProvideOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**SaveLyricsWithMedia** | Pointer to **bool** |  | [optional] [default to false]
**SaveTrickplayWithMedia** | Pointer to **bool** |  | [optional] [default to false]
**DisabledLyricFetchers** | Pointer to **[]string** |  | [optional] 
**LyricFetcherOrder** | Pointer to **[]string** |  | [optional] 
**PreferNonstandardArtistsTag** | Pointer to **bool** |  | [optional] [default to false]
**UseCustomTagDelimiters** | Pointer to **bool** |  | [optional] [default to false]
**CustomTagDelimiters** | Pointer to **[]string** |  | [optional] 
**DelimiterWhitelist** | Pointer to **[]string** |  | [optional] 
**AutomaticallyAddToCollection** | Pointer to **bool** |  | [optional] 
**AllowEmbeddedSubtitles** | Pointer to [**ModelModelEmbeddedSubtitleOptions**](ModelEmbeddedSubtitleOptions.md) | An enum representing the options to disable embedded subs. | [optional] 
**TypeOptions** | Pointer to [**[]ModelModelTypeOptions**](ModelModelTypeOptions.md) |  | [optional] 

## Methods

### NewModelLibraryOptions

`func NewModelLibraryOptions() *ModelLibraryOptions`

NewModelLibraryOptions instantiates a new ModelLibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLibraryOptionsWithDefaults

`func NewModelLibraryOptionsWithDefaults() *ModelLibraryOptions`

NewModelLibraryOptionsWithDefaults instantiates a new ModelLibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ModelLibraryOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelLibraryOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelLibraryOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelLibraryOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *ModelLibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *ModelLibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *ModelLibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *ModelLibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *ModelLibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *ModelLibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *ModelLibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *ModelLibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableLUFSScan

`func (o *ModelLibraryOptions) GetEnableLUFSScan() bool`

GetEnableLUFSScan returns the EnableLUFSScan field if non-nil, zero value otherwise.

### GetEnableLUFSScanOk

`func (o *ModelLibraryOptions) GetEnableLUFSScanOk() (*bool, bool)`

GetEnableLUFSScanOk returns a tuple with the EnableLUFSScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLUFSScan

`func (o *ModelLibraryOptions) SetEnableLUFSScan(v bool)`

SetEnableLUFSScan sets EnableLUFSScan field to given value.

### HasEnableLUFSScan

`func (o *ModelLibraryOptions) HasEnableLUFSScan() bool`

HasEnableLUFSScan returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *ModelLibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *ModelLibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *ModelLibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *ModelLibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *ModelLibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *ModelLibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetEnableTrickplayImageExtraction

`func (o *ModelLibraryOptions) GetEnableTrickplayImageExtraction() bool`

GetEnableTrickplayImageExtraction returns the EnableTrickplayImageExtraction field if non-nil, zero value otherwise.

### GetEnableTrickplayImageExtractionOk

`func (o *ModelLibraryOptions) GetEnableTrickplayImageExtractionOk() (*bool, bool)`

GetEnableTrickplayImageExtractionOk returns a tuple with the EnableTrickplayImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrickplayImageExtraction

`func (o *ModelLibraryOptions) SetEnableTrickplayImageExtraction(v bool)`

SetEnableTrickplayImageExtraction sets EnableTrickplayImageExtraction field to given value.

### HasEnableTrickplayImageExtraction

`func (o *ModelLibraryOptions) HasEnableTrickplayImageExtraction() bool`

HasEnableTrickplayImageExtraction returns a boolean if a field has been set.

### GetExtractTrickplayImagesDuringLibraryScan

`func (o *ModelLibraryOptions) GetExtractTrickplayImagesDuringLibraryScan() bool`

GetExtractTrickplayImagesDuringLibraryScan returns the ExtractTrickplayImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractTrickplayImagesDuringLibraryScanOk

`func (o *ModelLibraryOptions) GetExtractTrickplayImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractTrickplayImagesDuringLibraryScanOk returns a tuple with the ExtractTrickplayImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractTrickplayImagesDuringLibraryScan

`func (o *ModelLibraryOptions) SetExtractTrickplayImagesDuringLibraryScan(v bool)`

SetExtractTrickplayImagesDuringLibraryScan sets ExtractTrickplayImagesDuringLibraryScan field to given value.

### HasExtractTrickplayImagesDuringLibraryScan

`func (o *ModelLibraryOptions) HasExtractTrickplayImagesDuringLibraryScan() bool`

HasExtractTrickplayImagesDuringLibraryScan returns a boolean if a field has been set.

### GetPathInfos

`func (o *ModelLibraryOptions) GetPathInfos() []ModelModelMediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *ModelLibraryOptions) GetPathInfosOk() (*[]ModelModelMediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *ModelLibraryOptions) SetPathInfos(v []ModelModelMediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *ModelLibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *ModelLibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *ModelLibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *ModelLibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *ModelLibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetEnableInternetProviders

`func (o *ModelLibraryOptions) GetEnableInternetProviders() bool`

GetEnableInternetProviders returns the EnableInternetProviders field if non-nil, zero value otherwise.

### GetEnableInternetProvidersOk

`func (o *ModelLibraryOptions) GetEnableInternetProvidersOk() (*bool, bool)`

GetEnableInternetProvidersOk returns a tuple with the EnableInternetProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetProviders

`func (o *ModelLibraryOptions) SetEnableInternetProviders(v bool)`

SetEnableInternetProviders sets EnableInternetProviders field to given value.

### HasEnableInternetProviders

`func (o *ModelLibraryOptions) HasEnableInternetProviders() bool`

HasEnableInternetProviders returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *ModelLibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *ModelLibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *ModelLibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *ModelLibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *ModelLibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *ModelLibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableEmbeddedExtrasTitles

`func (o *ModelLibraryOptions) GetEnableEmbeddedExtrasTitles() bool`

GetEnableEmbeddedExtrasTitles returns the EnableEmbeddedExtrasTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedExtrasTitlesOk

`func (o *ModelLibraryOptions) GetEnableEmbeddedExtrasTitlesOk() (*bool, bool)`

GetEnableEmbeddedExtrasTitlesOk returns a tuple with the EnableEmbeddedExtrasTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedExtrasTitles

`func (o *ModelLibraryOptions) SetEnableEmbeddedExtrasTitles(v bool)`

SetEnableEmbeddedExtrasTitles sets EnableEmbeddedExtrasTitles field to given value.

### HasEnableEmbeddedExtrasTitles

`func (o *ModelLibraryOptions) HasEnableEmbeddedExtrasTitles() bool`

HasEnableEmbeddedExtrasTitles returns a boolean if a field has been set.

### GetEnableEmbeddedEpisodeInfos

`func (o *ModelLibraryOptions) GetEnableEmbeddedEpisodeInfos() bool`

GetEnableEmbeddedEpisodeInfos returns the EnableEmbeddedEpisodeInfos field if non-nil, zero value otherwise.

### GetEnableEmbeddedEpisodeInfosOk

`func (o *ModelLibraryOptions) GetEnableEmbeddedEpisodeInfosOk() (*bool, bool)`

GetEnableEmbeddedEpisodeInfosOk returns a tuple with the EnableEmbeddedEpisodeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedEpisodeInfos

`func (o *ModelLibraryOptions) SetEnableEmbeddedEpisodeInfos(v bool)`

SetEnableEmbeddedEpisodeInfos sets EnableEmbeddedEpisodeInfos field to given value.

### HasEnableEmbeddedEpisodeInfos

`func (o *ModelLibraryOptions) HasEnableEmbeddedEpisodeInfos() bool`

HasEnableEmbeddedEpisodeInfos returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *ModelLibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *ModelLibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ModelLibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ModelLibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ModelLibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ModelLibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *ModelLibraryOptions) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *ModelLibraryOptions) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *ModelLibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelLibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelLibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelLibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *ModelLibraryOptions) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *ModelLibraryOptions) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetSeasonZeroDisplayName

`func (o *ModelLibraryOptions) GetSeasonZeroDisplayName() string`

GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field if non-nil, zero value otherwise.

### GetSeasonZeroDisplayNameOk

`func (o *ModelLibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool)`

GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonZeroDisplayName

`func (o *ModelLibraryOptions) SetSeasonZeroDisplayName(v string)`

SetSeasonZeroDisplayName sets SeasonZeroDisplayName field to given value.

### HasSeasonZeroDisplayName

`func (o *ModelLibraryOptions) HasSeasonZeroDisplayName() bool`

HasSeasonZeroDisplayName returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *ModelLibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *ModelLibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *ModelLibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *ModelLibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### SetMetadataSaversNil

`func (o *ModelLibraryOptions) SetMetadataSaversNil(b bool)`

 SetMetadataSaversNil sets the value for MetadataSavers to be an explicit nil

### UnsetMetadataSavers
`func (o *ModelLibraryOptions) UnsetMetadataSavers()`

UnsetMetadataSavers ensures that no value is present for MetadataSavers, not even an explicit nil
### GetDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *ModelLibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *ModelLibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *ModelLibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *ModelLibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *ModelLibraryOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *ModelLibraryOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *ModelLibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *ModelLibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *ModelLibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *ModelLibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *ModelLibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *ModelLibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetDisabledMediaSegmentProviders

`func (o *ModelLibraryOptions) GetDisabledMediaSegmentProviders() []string`

GetDisabledMediaSegmentProviders returns the DisabledMediaSegmentProviders field if non-nil, zero value otherwise.

### GetDisabledMediaSegmentProvidersOk

`func (o *ModelLibraryOptions) GetDisabledMediaSegmentProvidersOk() (*[]string, bool)`

GetDisabledMediaSegmentProvidersOk returns a tuple with the DisabledMediaSegmentProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMediaSegmentProviders

`func (o *ModelLibraryOptions) SetDisabledMediaSegmentProviders(v []string)`

SetDisabledMediaSegmentProviders sets DisabledMediaSegmentProviders field to given value.

### HasDisabledMediaSegmentProviders

`func (o *ModelLibraryOptions) HasDisabledMediaSegmentProviders() bool`

HasDisabledMediaSegmentProviders returns a boolean if a field has been set.

### GetMediaSegmentProvideOrder

`func (o *ModelLibraryOptions) GetMediaSegmentProvideOrder() []string`

GetMediaSegmentProvideOrder returns the MediaSegmentProvideOrder field if non-nil, zero value otherwise.

### GetMediaSegmentProvideOrderOk

`func (o *ModelLibraryOptions) GetMediaSegmentProvideOrderOk() (*[]string, bool)`

GetMediaSegmentProvideOrderOk returns a tuple with the MediaSegmentProvideOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSegmentProvideOrder

`func (o *ModelLibraryOptions) SetMediaSegmentProvideOrder(v []string)`

SetMediaSegmentProvideOrder sets MediaSegmentProvideOrder field to given value.

### HasMediaSegmentProvideOrder

`func (o *ModelLibraryOptions) HasMediaSegmentProvideOrder() bool`

HasMediaSegmentProvideOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *ModelLibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *ModelLibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *ModelLibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *ModelLibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *ModelLibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### SetSubtitleDownloadLanguagesNil

`func (o *ModelLibraryOptions) SetSubtitleDownloadLanguagesNil(b bool)`

 SetSubtitleDownloadLanguagesNil sets the value for SubtitleDownloadLanguages to be an explicit nil

### UnsetSubtitleDownloadLanguages
`func (o *ModelLibraryOptions) UnsetSubtitleDownloadLanguages()`

UnsetSubtitleDownloadLanguages ensures that no value is present for SubtitleDownloadLanguages, not even an explicit nil
### GetRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *ModelLibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *ModelLibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *ModelLibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *ModelLibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetSaveLyricsWithMedia

`func (o *ModelLibraryOptions) GetSaveLyricsWithMedia() bool`

GetSaveLyricsWithMedia returns the SaveLyricsWithMedia field if non-nil, zero value otherwise.

### GetSaveLyricsWithMediaOk

`func (o *ModelLibraryOptions) GetSaveLyricsWithMediaOk() (*bool, bool)`

GetSaveLyricsWithMediaOk returns a tuple with the SaveLyricsWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLyricsWithMedia

`func (o *ModelLibraryOptions) SetSaveLyricsWithMedia(v bool)`

SetSaveLyricsWithMedia sets SaveLyricsWithMedia field to given value.

### HasSaveLyricsWithMedia

`func (o *ModelLibraryOptions) HasSaveLyricsWithMedia() bool`

HasSaveLyricsWithMedia returns a boolean if a field has been set.

### GetSaveTrickplayWithMedia

`func (o *ModelLibraryOptions) GetSaveTrickplayWithMedia() bool`

GetSaveTrickplayWithMedia returns the SaveTrickplayWithMedia field if non-nil, zero value otherwise.

### GetSaveTrickplayWithMediaOk

`func (o *ModelLibraryOptions) GetSaveTrickplayWithMediaOk() (*bool, bool)`

GetSaveTrickplayWithMediaOk returns a tuple with the SaveTrickplayWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveTrickplayWithMedia

`func (o *ModelLibraryOptions) SetSaveTrickplayWithMedia(v bool)`

SetSaveTrickplayWithMedia sets SaveTrickplayWithMedia field to given value.

### HasSaveTrickplayWithMedia

`func (o *ModelLibraryOptions) HasSaveTrickplayWithMedia() bool`

HasSaveTrickplayWithMedia returns a boolean if a field has been set.

### GetDisabledLyricFetchers

`func (o *ModelLibraryOptions) GetDisabledLyricFetchers() []string`

GetDisabledLyricFetchers returns the DisabledLyricFetchers field if non-nil, zero value otherwise.

### GetDisabledLyricFetchersOk

`func (o *ModelLibraryOptions) GetDisabledLyricFetchersOk() (*[]string, bool)`

GetDisabledLyricFetchersOk returns a tuple with the DisabledLyricFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLyricFetchers

`func (o *ModelLibraryOptions) SetDisabledLyricFetchers(v []string)`

SetDisabledLyricFetchers sets DisabledLyricFetchers field to given value.

### HasDisabledLyricFetchers

`func (o *ModelLibraryOptions) HasDisabledLyricFetchers() bool`

HasDisabledLyricFetchers returns a boolean if a field has been set.

### GetLyricFetcherOrder

`func (o *ModelLibraryOptions) GetLyricFetcherOrder() []string`

GetLyricFetcherOrder returns the LyricFetcherOrder field if non-nil, zero value otherwise.

### GetLyricFetcherOrderOk

`func (o *ModelLibraryOptions) GetLyricFetcherOrderOk() (*[]string, bool)`

GetLyricFetcherOrderOk returns a tuple with the LyricFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricFetcherOrder

`func (o *ModelLibraryOptions) SetLyricFetcherOrder(v []string)`

SetLyricFetcherOrder sets LyricFetcherOrder field to given value.

### HasLyricFetcherOrder

`func (o *ModelLibraryOptions) HasLyricFetcherOrder() bool`

HasLyricFetcherOrder returns a boolean if a field has been set.

### GetPreferNonstandardArtistsTag

`func (o *ModelLibraryOptions) GetPreferNonstandardArtistsTag() bool`

GetPreferNonstandardArtistsTag returns the PreferNonstandardArtistsTag field if non-nil, zero value otherwise.

### GetPreferNonstandardArtistsTagOk

`func (o *ModelLibraryOptions) GetPreferNonstandardArtistsTagOk() (*bool, bool)`

GetPreferNonstandardArtistsTagOk returns a tuple with the PreferNonstandardArtistsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferNonstandardArtistsTag

`func (o *ModelLibraryOptions) SetPreferNonstandardArtistsTag(v bool)`

SetPreferNonstandardArtistsTag sets PreferNonstandardArtistsTag field to given value.

### HasPreferNonstandardArtistsTag

`func (o *ModelLibraryOptions) HasPreferNonstandardArtistsTag() bool`

HasPreferNonstandardArtistsTag returns a boolean if a field has been set.

### GetUseCustomTagDelimiters

`func (o *ModelLibraryOptions) GetUseCustomTagDelimiters() bool`

GetUseCustomTagDelimiters returns the UseCustomTagDelimiters field if non-nil, zero value otherwise.

### GetUseCustomTagDelimitersOk

`func (o *ModelLibraryOptions) GetUseCustomTagDelimitersOk() (*bool, bool)`

GetUseCustomTagDelimitersOk returns a tuple with the UseCustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomTagDelimiters

`func (o *ModelLibraryOptions) SetUseCustomTagDelimiters(v bool)`

SetUseCustomTagDelimiters sets UseCustomTagDelimiters field to given value.

### HasUseCustomTagDelimiters

`func (o *ModelLibraryOptions) HasUseCustomTagDelimiters() bool`

HasUseCustomTagDelimiters returns a boolean if a field has been set.

### GetCustomTagDelimiters

`func (o *ModelLibraryOptions) GetCustomTagDelimiters() []string`

GetCustomTagDelimiters returns the CustomTagDelimiters field if non-nil, zero value otherwise.

### GetCustomTagDelimitersOk

`func (o *ModelLibraryOptions) GetCustomTagDelimitersOk() (*[]string, bool)`

GetCustomTagDelimitersOk returns a tuple with the CustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTagDelimiters

`func (o *ModelLibraryOptions) SetCustomTagDelimiters(v []string)`

SetCustomTagDelimiters sets CustomTagDelimiters field to given value.

### HasCustomTagDelimiters

`func (o *ModelLibraryOptions) HasCustomTagDelimiters() bool`

HasCustomTagDelimiters returns a boolean if a field has been set.

### GetDelimiterWhitelist

`func (o *ModelLibraryOptions) GetDelimiterWhitelist() []string`

GetDelimiterWhitelist returns the DelimiterWhitelist field if non-nil, zero value otherwise.

### GetDelimiterWhitelistOk

`func (o *ModelLibraryOptions) GetDelimiterWhitelistOk() (*[]string, bool)`

GetDelimiterWhitelistOk returns a tuple with the DelimiterWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiterWhitelist

`func (o *ModelLibraryOptions) SetDelimiterWhitelist(v []string)`

SetDelimiterWhitelist sets DelimiterWhitelist field to given value.

### HasDelimiterWhitelist

`func (o *ModelLibraryOptions) HasDelimiterWhitelist() bool`

HasDelimiterWhitelist returns a boolean if a field has been set.

### GetAutomaticallyAddToCollection

`func (o *ModelLibraryOptions) GetAutomaticallyAddToCollection() bool`

GetAutomaticallyAddToCollection returns the AutomaticallyAddToCollection field if non-nil, zero value otherwise.

### GetAutomaticallyAddToCollectionOk

`func (o *ModelLibraryOptions) GetAutomaticallyAddToCollectionOk() (*bool, bool)`

GetAutomaticallyAddToCollectionOk returns a tuple with the AutomaticallyAddToCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAddToCollection

`func (o *ModelLibraryOptions) SetAutomaticallyAddToCollection(v bool)`

SetAutomaticallyAddToCollection sets AutomaticallyAddToCollection field to given value.

### HasAutomaticallyAddToCollection

`func (o *ModelLibraryOptions) HasAutomaticallyAddToCollection() bool`

HasAutomaticallyAddToCollection returns a boolean if a field has been set.

### GetAllowEmbeddedSubtitles

`func (o *ModelLibraryOptions) GetAllowEmbeddedSubtitles() ModelModelEmbeddedSubtitleOptions`

GetAllowEmbeddedSubtitles returns the AllowEmbeddedSubtitles field if non-nil, zero value otherwise.

### GetAllowEmbeddedSubtitlesOk

`func (o *ModelLibraryOptions) GetAllowEmbeddedSubtitlesOk() (*ModelModelEmbeddedSubtitleOptions, bool)`

GetAllowEmbeddedSubtitlesOk returns a tuple with the AllowEmbeddedSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmbeddedSubtitles

`func (o *ModelLibraryOptions) SetAllowEmbeddedSubtitles(v ModelModelEmbeddedSubtitleOptions)`

SetAllowEmbeddedSubtitles sets AllowEmbeddedSubtitles field to given value.

### HasAllowEmbeddedSubtitles

`func (o *ModelLibraryOptions) HasAllowEmbeddedSubtitles() bool`

HasAllowEmbeddedSubtitles returns a boolean if a field has been set.

### GetTypeOptions

`func (o *ModelLibraryOptions) GetTypeOptions() []ModelModelTypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *ModelLibraryOptions) GetTypeOptionsOk() (*[]ModelModelTypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *ModelLibraryOptions) SetTypeOptions(v []ModelModelTypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *ModelLibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


