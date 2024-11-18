# JellyfinLibraryOptions

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
**PathInfos** | Pointer to [**[]JellyfinJellyfinMediaPathInfo**](JellyfinJellyfinMediaPathInfo.md) |  | [optional] 
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
**AllowEmbeddedSubtitles** | Pointer to [**JellyfinJellyfinEmbeddedSubtitleOptions**](JellyfinEmbeddedSubtitleOptions.md) | An enum representing the options to disable embedded subs. | [optional] 
**TypeOptions** | Pointer to [**[]JellyfinJellyfinTypeOptions**](JellyfinJellyfinTypeOptions.md) |  | [optional] 

## Methods

### NewJellyfinLibraryOptions

`func NewJellyfinLibraryOptions() *JellyfinLibraryOptions`

NewJellyfinLibraryOptions instantiates a new JellyfinLibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLibraryOptionsWithDefaults

`func NewJellyfinLibraryOptionsWithDefaults() *JellyfinLibraryOptions`

NewJellyfinLibraryOptionsWithDefaults instantiates a new JellyfinLibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JellyfinLibraryOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JellyfinLibraryOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JellyfinLibraryOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JellyfinLibraryOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *JellyfinLibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *JellyfinLibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *JellyfinLibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *JellyfinLibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *JellyfinLibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *JellyfinLibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *JellyfinLibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *JellyfinLibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableLUFSScan

`func (o *JellyfinLibraryOptions) GetEnableLUFSScan() bool`

GetEnableLUFSScan returns the EnableLUFSScan field if non-nil, zero value otherwise.

### GetEnableLUFSScanOk

`func (o *JellyfinLibraryOptions) GetEnableLUFSScanOk() (*bool, bool)`

GetEnableLUFSScanOk returns a tuple with the EnableLUFSScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLUFSScan

`func (o *JellyfinLibraryOptions) SetEnableLUFSScan(v bool)`

SetEnableLUFSScan sets EnableLUFSScan field to given value.

### HasEnableLUFSScan

`func (o *JellyfinLibraryOptions) HasEnableLUFSScan() bool`

HasEnableLUFSScan returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *JellyfinLibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *JellyfinLibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *JellyfinLibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *JellyfinLibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *JellyfinLibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetEnableTrickplayImageExtraction

`func (o *JellyfinLibraryOptions) GetEnableTrickplayImageExtraction() bool`

GetEnableTrickplayImageExtraction returns the EnableTrickplayImageExtraction field if non-nil, zero value otherwise.

### GetEnableTrickplayImageExtractionOk

`func (o *JellyfinLibraryOptions) GetEnableTrickplayImageExtractionOk() (*bool, bool)`

GetEnableTrickplayImageExtractionOk returns a tuple with the EnableTrickplayImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrickplayImageExtraction

`func (o *JellyfinLibraryOptions) SetEnableTrickplayImageExtraction(v bool)`

SetEnableTrickplayImageExtraction sets EnableTrickplayImageExtraction field to given value.

### HasEnableTrickplayImageExtraction

`func (o *JellyfinLibraryOptions) HasEnableTrickplayImageExtraction() bool`

HasEnableTrickplayImageExtraction returns a boolean if a field has been set.

### GetExtractTrickplayImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) GetExtractTrickplayImagesDuringLibraryScan() bool`

GetExtractTrickplayImagesDuringLibraryScan returns the ExtractTrickplayImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractTrickplayImagesDuringLibraryScanOk

`func (o *JellyfinLibraryOptions) GetExtractTrickplayImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractTrickplayImagesDuringLibraryScanOk returns a tuple with the ExtractTrickplayImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractTrickplayImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) SetExtractTrickplayImagesDuringLibraryScan(v bool)`

SetExtractTrickplayImagesDuringLibraryScan sets ExtractTrickplayImagesDuringLibraryScan field to given value.

### HasExtractTrickplayImagesDuringLibraryScan

`func (o *JellyfinLibraryOptions) HasExtractTrickplayImagesDuringLibraryScan() bool`

HasExtractTrickplayImagesDuringLibraryScan returns a boolean if a field has been set.

### GetPathInfos

`func (o *JellyfinLibraryOptions) GetPathInfos() []JellyfinJellyfinMediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *JellyfinLibraryOptions) GetPathInfosOk() (*[]JellyfinJellyfinMediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *JellyfinLibraryOptions) SetPathInfos(v []JellyfinJellyfinMediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *JellyfinLibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *JellyfinLibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *JellyfinLibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *JellyfinLibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *JellyfinLibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetEnableInternetProviders

`func (o *JellyfinLibraryOptions) GetEnableInternetProviders() bool`

GetEnableInternetProviders returns the EnableInternetProviders field if non-nil, zero value otherwise.

### GetEnableInternetProvidersOk

`func (o *JellyfinLibraryOptions) GetEnableInternetProvidersOk() (*bool, bool)`

GetEnableInternetProvidersOk returns a tuple with the EnableInternetProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetProviders

`func (o *JellyfinLibraryOptions) SetEnableInternetProviders(v bool)`

SetEnableInternetProviders sets EnableInternetProviders field to given value.

### HasEnableInternetProviders

`func (o *JellyfinLibraryOptions) HasEnableInternetProviders() bool`

HasEnableInternetProviders returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *JellyfinLibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *JellyfinLibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *JellyfinLibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *JellyfinLibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *JellyfinLibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *JellyfinLibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableEmbeddedExtrasTitles

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedExtrasTitles() bool`

GetEnableEmbeddedExtrasTitles returns the EnableEmbeddedExtrasTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedExtrasTitlesOk

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedExtrasTitlesOk() (*bool, bool)`

GetEnableEmbeddedExtrasTitlesOk returns a tuple with the EnableEmbeddedExtrasTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedExtrasTitles

`func (o *JellyfinLibraryOptions) SetEnableEmbeddedExtrasTitles(v bool)`

SetEnableEmbeddedExtrasTitles sets EnableEmbeddedExtrasTitles field to given value.

### HasEnableEmbeddedExtrasTitles

`func (o *JellyfinLibraryOptions) HasEnableEmbeddedExtrasTitles() bool`

HasEnableEmbeddedExtrasTitles returns a boolean if a field has been set.

### GetEnableEmbeddedEpisodeInfos

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedEpisodeInfos() bool`

GetEnableEmbeddedEpisodeInfos returns the EnableEmbeddedEpisodeInfos field if non-nil, zero value otherwise.

### GetEnableEmbeddedEpisodeInfosOk

`func (o *JellyfinLibraryOptions) GetEnableEmbeddedEpisodeInfosOk() (*bool, bool)`

GetEnableEmbeddedEpisodeInfosOk returns a tuple with the EnableEmbeddedEpisodeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedEpisodeInfos

`func (o *JellyfinLibraryOptions) SetEnableEmbeddedEpisodeInfos(v bool)`

SetEnableEmbeddedEpisodeInfos sets EnableEmbeddedEpisodeInfos field to given value.

### HasEnableEmbeddedEpisodeInfos

`func (o *JellyfinLibraryOptions) HasEnableEmbeddedEpisodeInfos() bool`

HasEnableEmbeddedEpisodeInfos returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *JellyfinLibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *JellyfinLibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *JellyfinLibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *JellyfinLibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *JellyfinLibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *JellyfinLibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *JellyfinLibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *JellyfinLibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *JellyfinLibraryOptions) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *JellyfinLibraryOptions) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *JellyfinLibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *JellyfinLibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *JellyfinLibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *JellyfinLibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *JellyfinLibraryOptions) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *JellyfinLibraryOptions) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetSeasonZeroDisplayName

`func (o *JellyfinLibraryOptions) GetSeasonZeroDisplayName() string`

GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field if non-nil, zero value otherwise.

### GetSeasonZeroDisplayNameOk

`func (o *JellyfinLibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool)`

GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonZeroDisplayName

`func (o *JellyfinLibraryOptions) SetSeasonZeroDisplayName(v string)`

SetSeasonZeroDisplayName sets SeasonZeroDisplayName field to given value.

### HasSeasonZeroDisplayName

`func (o *JellyfinLibraryOptions) HasSeasonZeroDisplayName() bool`

HasSeasonZeroDisplayName returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *JellyfinLibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *JellyfinLibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *JellyfinLibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *JellyfinLibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### SetMetadataSaversNil

`func (o *JellyfinLibraryOptions) SetMetadataSaversNil(b bool)`

 SetMetadataSaversNil sets the value for MetadataSavers to be an explicit nil

### UnsetMetadataSavers
`func (o *JellyfinLibraryOptions) UnsetMetadataSavers()`

UnsetMetadataSavers ensures that no value is present for MetadataSavers, not even an explicit nil
### GetDisabledLocalMetadataReaders

`func (o *JellyfinLibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *JellyfinLibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *JellyfinLibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *JellyfinLibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *JellyfinLibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *JellyfinLibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *JellyfinLibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *JellyfinLibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *JellyfinLibraryOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *JellyfinLibraryOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledSubtitleFetchers

`func (o *JellyfinLibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *JellyfinLibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *JellyfinLibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *JellyfinLibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *JellyfinLibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *JellyfinLibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *JellyfinLibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *JellyfinLibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetDisabledMediaSegmentProviders

`func (o *JellyfinLibraryOptions) GetDisabledMediaSegmentProviders() []string`

GetDisabledMediaSegmentProviders returns the DisabledMediaSegmentProviders field if non-nil, zero value otherwise.

### GetDisabledMediaSegmentProvidersOk

`func (o *JellyfinLibraryOptions) GetDisabledMediaSegmentProvidersOk() (*[]string, bool)`

GetDisabledMediaSegmentProvidersOk returns a tuple with the DisabledMediaSegmentProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMediaSegmentProviders

`func (o *JellyfinLibraryOptions) SetDisabledMediaSegmentProviders(v []string)`

SetDisabledMediaSegmentProviders sets DisabledMediaSegmentProviders field to given value.

### HasDisabledMediaSegmentProviders

`func (o *JellyfinLibraryOptions) HasDisabledMediaSegmentProviders() bool`

HasDisabledMediaSegmentProviders returns a boolean if a field has been set.

### GetMediaSegmentProvideOrder

`func (o *JellyfinLibraryOptions) GetMediaSegmentProvideOrder() []string`

GetMediaSegmentProvideOrder returns the MediaSegmentProvideOrder field if non-nil, zero value otherwise.

### GetMediaSegmentProvideOrderOk

`func (o *JellyfinLibraryOptions) GetMediaSegmentProvideOrderOk() (*[]string, bool)`

GetMediaSegmentProvideOrderOk returns a tuple with the MediaSegmentProvideOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSegmentProvideOrder

`func (o *JellyfinLibraryOptions) SetMediaSegmentProvideOrder(v []string)`

SetMediaSegmentProvideOrder sets MediaSegmentProvideOrder field to given value.

### HasMediaSegmentProvideOrder

`func (o *JellyfinLibraryOptions) HasMediaSegmentProvideOrder() bool`

HasMediaSegmentProvideOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *JellyfinLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *JellyfinLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *JellyfinLibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *JellyfinLibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *JellyfinLibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *JellyfinLibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *JellyfinLibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *JellyfinLibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *JellyfinLibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *JellyfinLibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *JellyfinLibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *JellyfinLibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### SetSubtitleDownloadLanguagesNil

`func (o *JellyfinLibraryOptions) SetSubtitleDownloadLanguagesNil(b bool)`

 SetSubtitleDownloadLanguagesNil sets the value for SubtitleDownloadLanguages to be an explicit nil

### UnsetSubtitleDownloadLanguages
`func (o *JellyfinLibraryOptions) UnsetSubtitleDownloadLanguages()`

UnsetSubtitleDownloadLanguages ensures that no value is present for SubtitleDownloadLanguages, not even an explicit nil
### GetRequirePerfectSubtitleMatch

`func (o *JellyfinLibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *JellyfinLibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *JellyfinLibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *JellyfinLibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *JellyfinLibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *JellyfinLibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *JellyfinLibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *JellyfinLibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetSaveLyricsWithMedia

`func (o *JellyfinLibraryOptions) GetSaveLyricsWithMedia() bool`

GetSaveLyricsWithMedia returns the SaveLyricsWithMedia field if non-nil, zero value otherwise.

### GetSaveLyricsWithMediaOk

`func (o *JellyfinLibraryOptions) GetSaveLyricsWithMediaOk() (*bool, bool)`

GetSaveLyricsWithMediaOk returns a tuple with the SaveLyricsWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLyricsWithMedia

`func (o *JellyfinLibraryOptions) SetSaveLyricsWithMedia(v bool)`

SetSaveLyricsWithMedia sets SaveLyricsWithMedia field to given value.

### HasSaveLyricsWithMedia

`func (o *JellyfinLibraryOptions) HasSaveLyricsWithMedia() bool`

HasSaveLyricsWithMedia returns a boolean if a field has been set.

### GetSaveTrickplayWithMedia

`func (o *JellyfinLibraryOptions) GetSaveTrickplayWithMedia() bool`

GetSaveTrickplayWithMedia returns the SaveTrickplayWithMedia field if non-nil, zero value otherwise.

### GetSaveTrickplayWithMediaOk

`func (o *JellyfinLibraryOptions) GetSaveTrickplayWithMediaOk() (*bool, bool)`

GetSaveTrickplayWithMediaOk returns a tuple with the SaveTrickplayWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveTrickplayWithMedia

`func (o *JellyfinLibraryOptions) SetSaveTrickplayWithMedia(v bool)`

SetSaveTrickplayWithMedia sets SaveTrickplayWithMedia field to given value.

### HasSaveTrickplayWithMedia

`func (o *JellyfinLibraryOptions) HasSaveTrickplayWithMedia() bool`

HasSaveTrickplayWithMedia returns a boolean if a field has been set.

### GetDisabledLyricFetchers

`func (o *JellyfinLibraryOptions) GetDisabledLyricFetchers() []string`

GetDisabledLyricFetchers returns the DisabledLyricFetchers field if non-nil, zero value otherwise.

### GetDisabledLyricFetchersOk

`func (o *JellyfinLibraryOptions) GetDisabledLyricFetchersOk() (*[]string, bool)`

GetDisabledLyricFetchersOk returns a tuple with the DisabledLyricFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLyricFetchers

`func (o *JellyfinLibraryOptions) SetDisabledLyricFetchers(v []string)`

SetDisabledLyricFetchers sets DisabledLyricFetchers field to given value.

### HasDisabledLyricFetchers

`func (o *JellyfinLibraryOptions) HasDisabledLyricFetchers() bool`

HasDisabledLyricFetchers returns a boolean if a field has been set.

### GetLyricFetcherOrder

`func (o *JellyfinLibraryOptions) GetLyricFetcherOrder() []string`

GetLyricFetcherOrder returns the LyricFetcherOrder field if non-nil, zero value otherwise.

### GetLyricFetcherOrderOk

`func (o *JellyfinLibraryOptions) GetLyricFetcherOrderOk() (*[]string, bool)`

GetLyricFetcherOrderOk returns a tuple with the LyricFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricFetcherOrder

`func (o *JellyfinLibraryOptions) SetLyricFetcherOrder(v []string)`

SetLyricFetcherOrder sets LyricFetcherOrder field to given value.

### HasLyricFetcherOrder

`func (o *JellyfinLibraryOptions) HasLyricFetcherOrder() bool`

HasLyricFetcherOrder returns a boolean if a field has been set.

### GetPreferNonstandardArtistsTag

`func (o *JellyfinLibraryOptions) GetPreferNonstandardArtistsTag() bool`

GetPreferNonstandardArtistsTag returns the PreferNonstandardArtistsTag field if non-nil, zero value otherwise.

### GetPreferNonstandardArtistsTagOk

`func (o *JellyfinLibraryOptions) GetPreferNonstandardArtistsTagOk() (*bool, bool)`

GetPreferNonstandardArtistsTagOk returns a tuple with the PreferNonstandardArtistsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferNonstandardArtistsTag

`func (o *JellyfinLibraryOptions) SetPreferNonstandardArtistsTag(v bool)`

SetPreferNonstandardArtistsTag sets PreferNonstandardArtistsTag field to given value.

### HasPreferNonstandardArtistsTag

`func (o *JellyfinLibraryOptions) HasPreferNonstandardArtistsTag() bool`

HasPreferNonstandardArtistsTag returns a boolean if a field has been set.

### GetUseCustomTagDelimiters

`func (o *JellyfinLibraryOptions) GetUseCustomTagDelimiters() bool`

GetUseCustomTagDelimiters returns the UseCustomTagDelimiters field if non-nil, zero value otherwise.

### GetUseCustomTagDelimitersOk

`func (o *JellyfinLibraryOptions) GetUseCustomTagDelimitersOk() (*bool, bool)`

GetUseCustomTagDelimitersOk returns a tuple with the UseCustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomTagDelimiters

`func (o *JellyfinLibraryOptions) SetUseCustomTagDelimiters(v bool)`

SetUseCustomTagDelimiters sets UseCustomTagDelimiters field to given value.

### HasUseCustomTagDelimiters

`func (o *JellyfinLibraryOptions) HasUseCustomTagDelimiters() bool`

HasUseCustomTagDelimiters returns a boolean if a field has been set.

### GetCustomTagDelimiters

`func (o *JellyfinLibraryOptions) GetCustomTagDelimiters() []string`

GetCustomTagDelimiters returns the CustomTagDelimiters field if non-nil, zero value otherwise.

### GetCustomTagDelimitersOk

`func (o *JellyfinLibraryOptions) GetCustomTagDelimitersOk() (*[]string, bool)`

GetCustomTagDelimitersOk returns a tuple with the CustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTagDelimiters

`func (o *JellyfinLibraryOptions) SetCustomTagDelimiters(v []string)`

SetCustomTagDelimiters sets CustomTagDelimiters field to given value.

### HasCustomTagDelimiters

`func (o *JellyfinLibraryOptions) HasCustomTagDelimiters() bool`

HasCustomTagDelimiters returns a boolean if a field has been set.

### GetDelimiterWhitelist

`func (o *JellyfinLibraryOptions) GetDelimiterWhitelist() []string`

GetDelimiterWhitelist returns the DelimiterWhitelist field if non-nil, zero value otherwise.

### GetDelimiterWhitelistOk

`func (o *JellyfinLibraryOptions) GetDelimiterWhitelistOk() (*[]string, bool)`

GetDelimiterWhitelistOk returns a tuple with the DelimiterWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiterWhitelist

`func (o *JellyfinLibraryOptions) SetDelimiterWhitelist(v []string)`

SetDelimiterWhitelist sets DelimiterWhitelist field to given value.

### HasDelimiterWhitelist

`func (o *JellyfinLibraryOptions) HasDelimiterWhitelist() bool`

HasDelimiterWhitelist returns a boolean if a field has been set.

### GetAutomaticallyAddToCollection

`func (o *JellyfinLibraryOptions) GetAutomaticallyAddToCollection() bool`

GetAutomaticallyAddToCollection returns the AutomaticallyAddToCollection field if non-nil, zero value otherwise.

### GetAutomaticallyAddToCollectionOk

`func (o *JellyfinLibraryOptions) GetAutomaticallyAddToCollectionOk() (*bool, bool)`

GetAutomaticallyAddToCollectionOk returns a tuple with the AutomaticallyAddToCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAddToCollection

`func (o *JellyfinLibraryOptions) SetAutomaticallyAddToCollection(v bool)`

SetAutomaticallyAddToCollection sets AutomaticallyAddToCollection field to given value.

### HasAutomaticallyAddToCollection

`func (o *JellyfinLibraryOptions) HasAutomaticallyAddToCollection() bool`

HasAutomaticallyAddToCollection returns a boolean if a field has been set.

### GetAllowEmbeddedSubtitles

`func (o *JellyfinLibraryOptions) GetAllowEmbeddedSubtitles() JellyfinJellyfinEmbeddedSubtitleOptions`

GetAllowEmbeddedSubtitles returns the AllowEmbeddedSubtitles field if non-nil, zero value otherwise.

### GetAllowEmbeddedSubtitlesOk

`func (o *JellyfinLibraryOptions) GetAllowEmbeddedSubtitlesOk() (*JellyfinJellyfinEmbeddedSubtitleOptions, bool)`

GetAllowEmbeddedSubtitlesOk returns a tuple with the AllowEmbeddedSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmbeddedSubtitles

`func (o *JellyfinLibraryOptions) SetAllowEmbeddedSubtitles(v JellyfinJellyfinEmbeddedSubtitleOptions)`

SetAllowEmbeddedSubtitles sets AllowEmbeddedSubtitles field to given value.

### HasAllowEmbeddedSubtitles

`func (o *JellyfinLibraryOptions) HasAllowEmbeddedSubtitles() bool`

HasAllowEmbeddedSubtitles returns a boolean if a field has been set.

### GetTypeOptions

`func (o *JellyfinLibraryOptions) GetTypeOptions() []JellyfinJellyfinTypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *JellyfinLibraryOptions) GetTypeOptionsOk() (*[]JellyfinJellyfinTypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *JellyfinLibraryOptions) SetTypeOptions(v []JellyfinJellyfinTypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *JellyfinLibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


