# JellyfinServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogFileRetentionDays** | Pointer to **int32** | Gets or sets the number of days we should retain log files. | [optional] 
**IsStartupWizardCompleted** | Pointer to **bool** | Gets or sets a value indicating whether this instance is first run. | [optional] 
**CachePath** | Pointer to **NullableString** | Gets or sets the cache path. | [optional] 
**PreviousVersion** | Pointer to **NullableString** | Gets or sets the last known version that was ran using the configuration. | [optional] 
**PreviousVersionStr** | Pointer to **NullableString** | Gets or sets the stringified PreviousVersion to be stored/loaded,  because System.Version itself isn&#39;t xml-serializable. | [optional] 
**EnableMetrics** | Pointer to **bool** | Gets or sets a value indicating whether to enable prometheus metrics exporting. | [optional] 
**EnableNormalizedItemByNameIds** | Pointer to **bool** |  | [optional] 
**IsPortAuthorized** | Pointer to **bool** | Gets or sets a value indicating whether this instance is port authorized. | [optional] 
**QuickConnectAvailable** | Pointer to **bool** | Gets or sets a value indicating whether quick connect is available for use on this server. | [optional] 
**EnableCaseSensitiveItemIds** | Pointer to **bool** | Gets or sets a value indicating whether [enable case sensitive item ids]. | [optional] 
**DisableLiveTvChannelUserDataName** | Pointer to **bool** |  | [optional] 
**MetadataPath** | Pointer to **string** | Gets or sets the metadata path. | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** | Gets or sets the preferred metadata language. | [optional] 
**MetadataCountryCode** | Pointer to **string** | Gets or sets the metadata country code. | [optional] 
**SortReplaceCharacters** | Pointer to **[]string** | Gets or sets characters to be replaced with a &#39; &#39; in strings to create a sort name. | [optional] 
**SortRemoveCharacters** | Pointer to **[]string** | Gets or sets characters to be removed from strings to create a sort name. | [optional] 
**SortRemoveWords** | Pointer to **[]string** | Gets or sets words to be removed from strings to create a sort name. | [optional] 
**MinResumePct** | Pointer to **int32** | Gets or sets the minimum percentage of an item that must be played in order for playstate to be updated. | [optional] 
**MaxResumePct** | Pointer to **int32** | Gets or sets the maximum percentage of an item that can be played while still saving playstate. If this percentage is crossed playstate will be reset to the beginning and the item will be marked watched. | [optional] 
**MinResumeDurationSeconds** | Pointer to **int32** | Gets or sets the minimum duration that an item must have in order to be eligible for playstate updates.. | [optional] 
**MinAudiobookResume** | Pointer to **int32** | Gets or sets the minimum minutes of a book that must be played in order for playstate to be updated. | [optional] 
**MaxAudiobookResume** | Pointer to **int32** | Gets or sets the remaining minutes of a book that can be played while still saving playstate. If this percentage is crossed playstate will be reset to the beginning and the item will be marked watched. | [optional] 
**InactiveSessionThreshold** | Pointer to **int32** | Gets or sets the threshold in minutes after a inactive session gets closed automatically.  If set to 0 the check for inactive sessions gets disabled. | [optional] 
**LibraryMonitorDelay** | Pointer to **int32** | Gets or sets the delay in seconds that we will wait after a file system change to try and discover what has been added/removed  Some delay is necessary with some items because their creation is not atomic.  It involves the creation of several  different directories and files. | [optional] 
**LibraryUpdateDuration** | Pointer to **int32** | Gets or sets the duration in seconds that we will wait after a library updated event before executing the library changed notification. | [optional] 
**ImageSavingConvention** | Pointer to [**JellyfinJellyfinImageSavingConvention**](JellyfinImageSavingConvention.md) | Gets or sets the image saving convention. | [optional] 
**MetadataOptions** | Pointer to [**[]JellyfinJellyfinMetadataOptions**](JellyfinJellyfinMetadataOptions.md) |  | [optional] 
**SkipDeserializationForBasicTypes** | Pointer to **bool** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**UICulture** | Pointer to **string** |  | [optional] 
**SaveMetadataHidden** | Pointer to **bool** |  | [optional] 
**ContentTypes** | Pointer to [**[]JellyfinJellyfinNameValuePair**](JellyfinJellyfinNameValuePair.md) |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**EnableFolderView** | Pointer to **bool** |  | [optional] 
**EnableGroupingIntoCollections** | Pointer to **bool** |  | [optional] 
**DisplaySpecialsWithinSeasons** | Pointer to **bool** |  | [optional] 
**CodecsUsed** | Pointer to **[]string** |  | [optional] 
**PluginRepositories** | Pointer to [**[]JellyfinJellyfinRepositoryInfo**](JellyfinJellyfinRepositoryInfo.md) |  | [optional] 
**EnableExternalContentInSuggestions** | Pointer to **bool** |  | [optional] 
**ImageExtractionTimeoutMs** | Pointer to **int32** |  | [optional] 
**PathSubstitutions** | Pointer to [**[]JellyfinJellyfinPathSubstitution**](JellyfinJellyfinPathSubstitution.md) |  | [optional] 
**EnableSlowResponseWarning** | Pointer to **bool** | Gets or sets a value indicating whether slow server responses should be logged as a warning. | [optional] 
**SlowResponseThresholdMs** | Pointer to **int64** | Gets or sets the threshold for the slow response time warning in ms. | [optional] 
**CorsHosts** | Pointer to **[]string** | Gets or sets the cors hosts. | [optional] 
**ActivityLogRetentionDays** | Pointer to **NullableInt32** | Gets or sets the number of days we should retain activity logs. | [optional] 
**LibraryScanFanoutConcurrency** | Pointer to **int32** | Gets or sets the how the library scan fans out. | [optional] 
**LibraryMetadataRefreshConcurrency** | Pointer to **int32** | Gets or sets the how many metadata refreshes can run concurrently. | [optional] 
**RemoveOldPlugins** | Pointer to **bool** | Gets or sets a value indicating whether older plugins should automatically be deleted from the plugin folder. | [optional] 
**AllowClientLogUpload** | Pointer to **bool** | Gets or sets a value indicating whether clients should be allowed to upload logs. | [optional] 
**DummyChapterDuration** | Pointer to **int32** | Gets or sets the dummy chapter duration in seconds, use 0 (zero) or less to disable generation alltogether. | [optional] 
**ChapterImageResolution** | Pointer to [**JellyfinJellyfinImageResolution**](JellyfinImageResolution.md) | Gets or sets the chapter image resolution. | [optional] 
**ParallelImageEncodingLimit** | Pointer to **int32** | Gets or sets the limit for parallel image encoding. | [optional] 
**CastReceiverApplications** | Pointer to [**[]JellyfinJellyfinCastReceiverApplication**](JellyfinJellyfinCastReceiverApplication.md) | Gets or sets the list of cast receiver applications. | [optional] 
**TrickplayOptions** | Pointer to [**JellyfinJellyfinTrickplayOptions**](JellyfinTrickplayOptions.md) | Gets or sets the trickplay options. | [optional] 

## Methods

### NewJellyfinServerConfiguration

`func NewJellyfinServerConfiguration() *JellyfinServerConfiguration`

NewJellyfinServerConfiguration instantiates a new JellyfinServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinServerConfigurationWithDefaults

`func NewJellyfinServerConfigurationWithDefaults() *JellyfinServerConfiguration`

NewJellyfinServerConfigurationWithDefaults instantiates a new JellyfinServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogFileRetentionDays

`func (o *JellyfinServerConfiguration) GetLogFileRetentionDays() int32`

GetLogFileRetentionDays returns the LogFileRetentionDays field if non-nil, zero value otherwise.

### GetLogFileRetentionDaysOk

`func (o *JellyfinServerConfiguration) GetLogFileRetentionDaysOk() (*int32, bool)`

GetLogFileRetentionDaysOk returns a tuple with the LogFileRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileRetentionDays

`func (o *JellyfinServerConfiguration) SetLogFileRetentionDays(v int32)`

SetLogFileRetentionDays sets LogFileRetentionDays field to given value.

### HasLogFileRetentionDays

`func (o *JellyfinServerConfiguration) HasLogFileRetentionDays() bool`

HasLogFileRetentionDays returns a boolean if a field has been set.

### GetIsStartupWizardCompleted

`func (o *JellyfinServerConfiguration) GetIsStartupWizardCompleted() bool`

GetIsStartupWizardCompleted returns the IsStartupWizardCompleted field if non-nil, zero value otherwise.

### GetIsStartupWizardCompletedOk

`func (o *JellyfinServerConfiguration) GetIsStartupWizardCompletedOk() (*bool, bool)`

GetIsStartupWizardCompletedOk returns a tuple with the IsStartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartupWizardCompleted

`func (o *JellyfinServerConfiguration) SetIsStartupWizardCompleted(v bool)`

SetIsStartupWizardCompleted sets IsStartupWizardCompleted field to given value.

### HasIsStartupWizardCompleted

`func (o *JellyfinServerConfiguration) HasIsStartupWizardCompleted() bool`

HasIsStartupWizardCompleted returns a boolean if a field has been set.

### GetCachePath

`func (o *JellyfinServerConfiguration) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *JellyfinServerConfiguration) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *JellyfinServerConfiguration) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *JellyfinServerConfiguration) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### SetCachePathNil

`func (o *JellyfinServerConfiguration) SetCachePathNil(b bool)`

 SetCachePathNil sets the value for CachePath to be an explicit nil

### UnsetCachePath
`func (o *JellyfinServerConfiguration) UnsetCachePath()`

UnsetCachePath ensures that no value is present for CachePath, not even an explicit nil
### GetPreviousVersion

`func (o *JellyfinServerConfiguration) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *JellyfinServerConfiguration) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *JellyfinServerConfiguration) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *JellyfinServerConfiguration) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *JellyfinServerConfiguration) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *JellyfinServerConfiguration) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetPreviousVersionStr

`func (o *JellyfinServerConfiguration) GetPreviousVersionStr() string`

GetPreviousVersionStr returns the PreviousVersionStr field if non-nil, zero value otherwise.

### GetPreviousVersionStrOk

`func (o *JellyfinServerConfiguration) GetPreviousVersionStrOk() (*string, bool)`

GetPreviousVersionStrOk returns a tuple with the PreviousVersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersionStr

`func (o *JellyfinServerConfiguration) SetPreviousVersionStr(v string)`

SetPreviousVersionStr sets PreviousVersionStr field to given value.

### HasPreviousVersionStr

`func (o *JellyfinServerConfiguration) HasPreviousVersionStr() bool`

HasPreviousVersionStr returns a boolean if a field has been set.

### SetPreviousVersionStrNil

`func (o *JellyfinServerConfiguration) SetPreviousVersionStrNil(b bool)`

 SetPreviousVersionStrNil sets the value for PreviousVersionStr to be an explicit nil

### UnsetPreviousVersionStr
`func (o *JellyfinServerConfiguration) UnsetPreviousVersionStr()`

UnsetPreviousVersionStr ensures that no value is present for PreviousVersionStr, not even an explicit nil
### GetEnableMetrics

`func (o *JellyfinServerConfiguration) GetEnableMetrics() bool`

GetEnableMetrics returns the EnableMetrics field if non-nil, zero value otherwise.

### GetEnableMetricsOk

`func (o *JellyfinServerConfiguration) GetEnableMetricsOk() (*bool, bool)`

GetEnableMetricsOk returns a tuple with the EnableMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetrics

`func (o *JellyfinServerConfiguration) SetEnableMetrics(v bool)`

SetEnableMetrics sets EnableMetrics field to given value.

### HasEnableMetrics

`func (o *JellyfinServerConfiguration) HasEnableMetrics() bool`

HasEnableMetrics returns a boolean if a field has been set.

### GetEnableNormalizedItemByNameIds

`func (o *JellyfinServerConfiguration) GetEnableNormalizedItemByNameIds() bool`

GetEnableNormalizedItemByNameIds returns the EnableNormalizedItemByNameIds field if non-nil, zero value otherwise.

### GetEnableNormalizedItemByNameIdsOk

`func (o *JellyfinServerConfiguration) GetEnableNormalizedItemByNameIdsOk() (*bool, bool)`

GetEnableNormalizedItemByNameIdsOk returns a tuple with the EnableNormalizedItemByNameIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNormalizedItemByNameIds

`func (o *JellyfinServerConfiguration) SetEnableNormalizedItemByNameIds(v bool)`

SetEnableNormalizedItemByNameIds sets EnableNormalizedItemByNameIds field to given value.

### HasEnableNormalizedItemByNameIds

`func (o *JellyfinServerConfiguration) HasEnableNormalizedItemByNameIds() bool`

HasEnableNormalizedItemByNameIds returns a boolean if a field has been set.

### GetIsPortAuthorized

`func (o *JellyfinServerConfiguration) GetIsPortAuthorized() bool`

GetIsPortAuthorized returns the IsPortAuthorized field if non-nil, zero value otherwise.

### GetIsPortAuthorizedOk

`func (o *JellyfinServerConfiguration) GetIsPortAuthorizedOk() (*bool, bool)`

GetIsPortAuthorizedOk returns a tuple with the IsPortAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPortAuthorized

`func (o *JellyfinServerConfiguration) SetIsPortAuthorized(v bool)`

SetIsPortAuthorized sets IsPortAuthorized field to given value.

### HasIsPortAuthorized

`func (o *JellyfinServerConfiguration) HasIsPortAuthorized() bool`

HasIsPortAuthorized returns a boolean if a field has been set.

### GetQuickConnectAvailable

`func (o *JellyfinServerConfiguration) GetQuickConnectAvailable() bool`

GetQuickConnectAvailable returns the QuickConnectAvailable field if non-nil, zero value otherwise.

### GetQuickConnectAvailableOk

`func (o *JellyfinServerConfiguration) GetQuickConnectAvailableOk() (*bool, bool)`

GetQuickConnectAvailableOk returns a tuple with the QuickConnectAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickConnectAvailable

`func (o *JellyfinServerConfiguration) SetQuickConnectAvailable(v bool)`

SetQuickConnectAvailable sets QuickConnectAvailable field to given value.

### HasQuickConnectAvailable

`func (o *JellyfinServerConfiguration) HasQuickConnectAvailable() bool`

HasQuickConnectAvailable returns a boolean if a field has been set.

### GetEnableCaseSensitiveItemIds

`func (o *JellyfinServerConfiguration) GetEnableCaseSensitiveItemIds() bool`

GetEnableCaseSensitiveItemIds returns the EnableCaseSensitiveItemIds field if non-nil, zero value otherwise.

### GetEnableCaseSensitiveItemIdsOk

`func (o *JellyfinServerConfiguration) GetEnableCaseSensitiveItemIdsOk() (*bool, bool)`

GetEnableCaseSensitiveItemIdsOk returns a tuple with the EnableCaseSensitiveItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaseSensitiveItemIds

`func (o *JellyfinServerConfiguration) SetEnableCaseSensitiveItemIds(v bool)`

SetEnableCaseSensitiveItemIds sets EnableCaseSensitiveItemIds field to given value.

### HasEnableCaseSensitiveItemIds

`func (o *JellyfinServerConfiguration) HasEnableCaseSensitiveItemIds() bool`

HasEnableCaseSensitiveItemIds returns a boolean if a field has been set.

### GetDisableLiveTvChannelUserDataName

`func (o *JellyfinServerConfiguration) GetDisableLiveTvChannelUserDataName() bool`

GetDisableLiveTvChannelUserDataName returns the DisableLiveTvChannelUserDataName field if non-nil, zero value otherwise.

### GetDisableLiveTvChannelUserDataNameOk

`func (o *JellyfinServerConfiguration) GetDisableLiveTvChannelUserDataNameOk() (*bool, bool)`

GetDisableLiveTvChannelUserDataNameOk returns a tuple with the DisableLiveTvChannelUserDataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLiveTvChannelUserDataName

`func (o *JellyfinServerConfiguration) SetDisableLiveTvChannelUserDataName(v bool)`

SetDisableLiveTvChannelUserDataName sets DisableLiveTvChannelUserDataName field to given value.

### HasDisableLiveTvChannelUserDataName

`func (o *JellyfinServerConfiguration) HasDisableLiveTvChannelUserDataName() bool`

HasDisableLiveTvChannelUserDataName returns a boolean if a field has been set.

### GetMetadataPath

`func (o *JellyfinServerConfiguration) GetMetadataPath() string`

GetMetadataPath returns the MetadataPath field if non-nil, zero value otherwise.

### GetMetadataPathOk

`func (o *JellyfinServerConfiguration) GetMetadataPathOk() (*string, bool)`

GetMetadataPathOk returns a tuple with the MetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPath

`func (o *JellyfinServerConfiguration) SetMetadataPath(v string)`

SetMetadataPath sets MetadataPath field to given value.

### HasMetadataPath

`func (o *JellyfinServerConfiguration) HasMetadataPath() bool`

HasMetadataPath returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *JellyfinServerConfiguration) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *JellyfinServerConfiguration) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *JellyfinServerConfiguration) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *JellyfinServerConfiguration) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *JellyfinServerConfiguration) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *JellyfinServerConfiguration) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *JellyfinServerConfiguration) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *JellyfinServerConfiguration) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSortReplaceCharacters

`func (o *JellyfinServerConfiguration) GetSortReplaceCharacters() []string`

GetSortReplaceCharacters returns the SortReplaceCharacters field if non-nil, zero value otherwise.

### GetSortReplaceCharactersOk

`func (o *JellyfinServerConfiguration) GetSortReplaceCharactersOk() (*[]string, bool)`

GetSortReplaceCharactersOk returns a tuple with the SortReplaceCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortReplaceCharacters

`func (o *JellyfinServerConfiguration) SetSortReplaceCharacters(v []string)`

SetSortReplaceCharacters sets SortReplaceCharacters field to given value.

### HasSortReplaceCharacters

`func (o *JellyfinServerConfiguration) HasSortReplaceCharacters() bool`

HasSortReplaceCharacters returns a boolean if a field has been set.

### GetSortRemoveCharacters

`func (o *JellyfinServerConfiguration) GetSortRemoveCharacters() []string`

GetSortRemoveCharacters returns the SortRemoveCharacters field if non-nil, zero value otherwise.

### GetSortRemoveCharactersOk

`func (o *JellyfinServerConfiguration) GetSortRemoveCharactersOk() (*[]string, bool)`

GetSortRemoveCharactersOk returns a tuple with the SortRemoveCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveCharacters

`func (o *JellyfinServerConfiguration) SetSortRemoveCharacters(v []string)`

SetSortRemoveCharacters sets SortRemoveCharacters field to given value.

### HasSortRemoveCharacters

`func (o *JellyfinServerConfiguration) HasSortRemoveCharacters() bool`

HasSortRemoveCharacters returns a boolean if a field has been set.

### GetSortRemoveWords

`func (o *JellyfinServerConfiguration) GetSortRemoveWords() []string`

GetSortRemoveWords returns the SortRemoveWords field if non-nil, zero value otherwise.

### GetSortRemoveWordsOk

`func (o *JellyfinServerConfiguration) GetSortRemoveWordsOk() (*[]string, bool)`

GetSortRemoveWordsOk returns a tuple with the SortRemoveWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveWords

`func (o *JellyfinServerConfiguration) SetSortRemoveWords(v []string)`

SetSortRemoveWords sets SortRemoveWords field to given value.

### HasSortRemoveWords

`func (o *JellyfinServerConfiguration) HasSortRemoveWords() bool`

HasSortRemoveWords returns a boolean if a field has been set.

### GetMinResumePct

`func (o *JellyfinServerConfiguration) GetMinResumePct() int32`

GetMinResumePct returns the MinResumePct field if non-nil, zero value otherwise.

### GetMinResumePctOk

`func (o *JellyfinServerConfiguration) GetMinResumePctOk() (*int32, bool)`

GetMinResumePctOk returns a tuple with the MinResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumePct

`func (o *JellyfinServerConfiguration) SetMinResumePct(v int32)`

SetMinResumePct sets MinResumePct field to given value.

### HasMinResumePct

`func (o *JellyfinServerConfiguration) HasMinResumePct() bool`

HasMinResumePct returns a boolean if a field has been set.

### GetMaxResumePct

`func (o *JellyfinServerConfiguration) GetMaxResumePct() int32`

GetMaxResumePct returns the MaxResumePct field if non-nil, zero value otherwise.

### GetMaxResumePctOk

`func (o *JellyfinServerConfiguration) GetMaxResumePctOk() (*int32, bool)`

GetMaxResumePctOk returns a tuple with the MaxResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResumePct

`func (o *JellyfinServerConfiguration) SetMaxResumePct(v int32)`

SetMaxResumePct sets MaxResumePct field to given value.

### HasMaxResumePct

`func (o *JellyfinServerConfiguration) HasMaxResumePct() bool`

HasMaxResumePct returns a boolean if a field has been set.

### GetMinResumeDurationSeconds

`func (o *JellyfinServerConfiguration) GetMinResumeDurationSeconds() int32`

GetMinResumeDurationSeconds returns the MinResumeDurationSeconds field if non-nil, zero value otherwise.

### GetMinResumeDurationSecondsOk

`func (o *JellyfinServerConfiguration) GetMinResumeDurationSecondsOk() (*int32, bool)`

GetMinResumeDurationSecondsOk returns a tuple with the MinResumeDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumeDurationSeconds

`func (o *JellyfinServerConfiguration) SetMinResumeDurationSeconds(v int32)`

SetMinResumeDurationSeconds sets MinResumeDurationSeconds field to given value.

### HasMinResumeDurationSeconds

`func (o *JellyfinServerConfiguration) HasMinResumeDurationSeconds() bool`

HasMinResumeDurationSeconds returns a boolean if a field has been set.

### GetMinAudiobookResume

`func (o *JellyfinServerConfiguration) GetMinAudiobookResume() int32`

GetMinAudiobookResume returns the MinAudiobookResume field if non-nil, zero value otherwise.

### GetMinAudiobookResumeOk

`func (o *JellyfinServerConfiguration) GetMinAudiobookResumeOk() (*int32, bool)`

GetMinAudiobookResumeOk returns a tuple with the MinAudiobookResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAudiobookResume

`func (o *JellyfinServerConfiguration) SetMinAudiobookResume(v int32)`

SetMinAudiobookResume sets MinAudiobookResume field to given value.

### HasMinAudiobookResume

`func (o *JellyfinServerConfiguration) HasMinAudiobookResume() bool`

HasMinAudiobookResume returns a boolean if a field has been set.

### GetMaxAudiobookResume

`func (o *JellyfinServerConfiguration) GetMaxAudiobookResume() int32`

GetMaxAudiobookResume returns the MaxAudiobookResume field if non-nil, zero value otherwise.

### GetMaxAudiobookResumeOk

`func (o *JellyfinServerConfiguration) GetMaxAudiobookResumeOk() (*int32, bool)`

GetMaxAudiobookResumeOk returns a tuple with the MaxAudiobookResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudiobookResume

`func (o *JellyfinServerConfiguration) SetMaxAudiobookResume(v int32)`

SetMaxAudiobookResume sets MaxAudiobookResume field to given value.

### HasMaxAudiobookResume

`func (o *JellyfinServerConfiguration) HasMaxAudiobookResume() bool`

HasMaxAudiobookResume returns a boolean if a field has been set.

### GetInactiveSessionThreshold

`func (o *JellyfinServerConfiguration) GetInactiveSessionThreshold() int32`

GetInactiveSessionThreshold returns the InactiveSessionThreshold field if non-nil, zero value otherwise.

### GetInactiveSessionThresholdOk

`func (o *JellyfinServerConfiguration) GetInactiveSessionThresholdOk() (*int32, bool)`

GetInactiveSessionThresholdOk returns a tuple with the InactiveSessionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSessionThreshold

`func (o *JellyfinServerConfiguration) SetInactiveSessionThreshold(v int32)`

SetInactiveSessionThreshold sets InactiveSessionThreshold field to given value.

### HasInactiveSessionThreshold

`func (o *JellyfinServerConfiguration) HasInactiveSessionThreshold() bool`

HasInactiveSessionThreshold returns a boolean if a field has been set.

### GetLibraryMonitorDelay

`func (o *JellyfinServerConfiguration) GetLibraryMonitorDelay() int32`

GetLibraryMonitorDelay returns the LibraryMonitorDelay field if non-nil, zero value otherwise.

### GetLibraryMonitorDelayOk

`func (o *JellyfinServerConfiguration) GetLibraryMonitorDelayOk() (*int32, bool)`

GetLibraryMonitorDelayOk returns a tuple with the LibraryMonitorDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMonitorDelay

`func (o *JellyfinServerConfiguration) SetLibraryMonitorDelay(v int32)`

SetLibraryMonitorDelay sets LibraryMonitorDelay field to given value.

### HasLibraryMonitorDelay

`func (o *JellyfinServerConfiguration) HasLibraryMonitorDelay() bool`

HasLibraryMonitorDelay returns a boolean if a field has been set.

### GetLibraryUpdateDuration

`func (o *JellyfinServerConfiguration) GetLibraryUpdateDuration() int32`

GetLibraryUpdateDuration returns the LibraryUpdateDuration field if non-nil, zero value otherwise.

### GetLibraryUpdateDurationOk

`func (o *JellyfinServerConfiguration) GetLibraryUpdateDurationOk() (*int32, bool)`

GetLibraryUpdateDurationOk returns a tuple with the LibraryUpdateDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryUpdateDuration

`func (o *JellyfinServerConfiguration) SetLibraryUpdateDuration(v int32)`

SetLibraryUpdateDuration sets LibraryUpdateDuration field to given value.

### HasLibraryUpdateDuration

`func (o *JellyfinServerConfiguration) HasLibraryUpdateDuration() bool`

HasLibraryUpdateDuration returns a boolean if a field has been set.

### GetImageSavingConvention

`func (o *JellyfinServerConfiguration) GetImageSavingConvention() JellyfinJellyfinImageSavingConvention`

GetImageSavingConvention returns the ImageSavingConvention field if non-nil, zero value otherwise.

### GetImageSavingConventionOk

`func (o *JellyfinServerConfiguration) GetImageSavingConventionOk() (*JellyfinJellyfinImageSavingConvention, bool)`

GetImageSavingConventionOk returns a tuple with the ImageSavingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSavingConvention

`func (o *JellyfinServerConfiguration) SetImageSavingConvention(v JellyfinJellyfinImageSavingConvention)`

SetImageSavingConvention sets ImageSavingConvention field to given value.

### HasImageSavingConvention

`func (o *JellyfinServerConfiguration) HasImageSavingConvention() bool`

HasImageSavingConvention returns a boolean if a field has been set.

### GetMetadataOptions

`func (o *JellyfinServerConfiguration) GetMetadataOptions() []JellyfinJellyfinMetadataOptions`

GetMetadataOptions returns the MetadataOptions field if non-nil, zero value otherwise.

### GetMetadataOptionsOk

`func (o *JellyfinServerConfiguration) GetMetadataOptionsOk() (*[]JellyfinJellyfinMetadataOptions, bool)`

GetMetadataOptionsOk returns a tuple with the MetadataOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataOptions

`func (o *JellyfinServerConfiguration) SetMetadataOptions(v []JellyfinJellyfinMetadataOptions)`

SetMetadataOptions sets MetadataOptions field to given value.

### HasMetadataOptions

`func (o *JellyfinServerConfiguration) HasMetadataOptions() bool`

HasMetadataOptions returns a boolean if a field has been set.

### GetSkipDeserializationForBasicTypes

`func (o *JellyfinServerConfiguration) GetSkipDeserializationForBasicTypes() bool`

GetSkipDeserializationForBasicTypes returns the SkipDeserializationForBasicTypes field if non-nil, zero value otherwise.

### GetSkipDeserializationForBasicTypesOk

`func (o *JellyfinServerConfiguration) GetSkipDeserializationForBasicTypesOk() (*bool, bool)`

GetSkipDeserializationForBasicTypesOk returns a tuple with the SkipDeserializationForBasicTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDeserializationForBasicTypes

`func (o *JellyfinServerConfiguration) SetSkipDeserializationForBasicTypes(v bool)`

SetSkipDeserializationForBasicTypes sets SkipDeserializationForBasicTypes field to given value.

### HasSkipDeserializationForBasicTypes

`func (o *JellyfinServerConfiguration) HasSkipDeserializationForBasicTypes() bool`

HasSkipDeserializationForBasicTypes returns a boolean if a field has been set.

### GetServerName

`func (o *JellyfinServerConfiguration) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *JellyfinServerConfiguration) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *JellyfinServerConfiguration) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *JellyfinServerConfiguration) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetUICulture

`func (o *JellyfinServerConfiguration) GetUICulture() string`

GetUICulture returns the UICulture field if non-nil, zero value otherwise.

### GetUICultureOk

`func (o *JellyfinServerConfiguration) GetUICultureOk() (*string, bool)`

GetUICultureOk returns a tuple with the UICulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUICulture

`func (o *JellyfinServerConfiguration) SetUICulture(v string)`

SetUICulture sets UICulture field to given value.

### HasUICulture

`func (o *JellyfinServerConfiguration) HasUICulture() bool`

HasUICulture returns a boolean if a field has been set.

### GetSaveMetadataHidden

`func (o *JellyfinServerConfiguration) GetSaveMetadataHidden() bool`

GetSaveMetadataHidden returns the SaveMetadataHidden field if non-nil, zero value otherwise.

### GetSaveMetadataHiddenOk

`func (o *JellyfinServerConfiguration) GetSaveMetadataHiddenOk() (*bool, bool)`

GetSaveMetadataHiddenOk returns a tuple with the SaveMetadataHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMetadataHidden

`func (o *JellyfinServerConfiguration) SetSaveMetadataHidden(v bool)`

SetSaveMetadataHidden sets SaveMetadataHidden field to given value.

### HasSaveMetadataHidden

`func (o *JellyfinServerConfiguration) HasSaveMetadataHidden() bool`

HasSaveMetadataHidden returns a boolean if a field has been set.

### GetContentTypes

`func (o *JellyfinServerConfiguration) GetContentTypes() []JellyfinJellyfinNameValuePair`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *JellyfinServerConfiguration) GetContentTypesOk() (*[]JellyfinJellyfinNameValuePair, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *JellyfinServerConfiguration) SetContentTypes(v []JellyfinJellyfinNameValuePair)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *JellyfinServerConfiguration) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *JellyfinServerConfiguration) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *JellyfinServerConfiguration) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *JellyfinServerConfiguration) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *JellyfinServerConfiguration) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetEnableFolderView

`func (o *JellyfinServerConfiguration) GetEnableFolderView() bool`

GetEnableFolderView returns the EnableFolderView field if non-nil, zero value otherwise.

### GetEnableFolderViewOk

`func (o *JellyfinServerConfiguration) GetEnableFolderViewOk() (*bool, bool)`

GetEnableFolderViewOk returns a tuple with the EnableFolderView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFolderView

`func (o *JellyfinServerConfiguration) SetEnableFolderView(v bool)`

SetEnableFolderView sets EnableFolderView field to given value.

### HasEnableFolderView

`func (o *JellyfinServerConfiguration) HasEnableFolderView() bool`

HasEnableFolderView returns a boolean if a field has been set.

### GetEnableGroupingIntoCollections

`func (o *JellyfinServerConfiguration) GetEnableGroupingIntoCollections() bool`

GetEnableGroupingIntoCollections returns the EnableGroupingIntoCollections field if non-nil, zero value otherwise.

### GetEnableGroupingIntoCollectionsOk

`func (o *JellyfinServerConfiguration) GetEnableGroupingIntoCollectionsOk() (*bool, bool)`

GetEnableGroupingIntoCollectionsOk returns a tuple with the EnableGroupingIntoCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupingIntoCollections

`func (o *JellyfinServerConfiguration) SetEnableGroupingIntoCollections(v bool)`

SetEnableGroupingIntoCollections sets EnableGroupingIntoCollections field to given value.

### HasEnableGroupingIntoCollections

`func (o *JellyfinServerConfiguration) HasEnableGroupingIntoCollections() bool`

HasEnableGroupingIntoCollections returns a boolean if a field has been set.

### GetDisplaySpecialsWithinSeasons

`func (o *JellyfinServerConfiguration) GetDisplaySpecialsWithinSeasons() bool`

GetDisplaySpecialsWithinSeasons returns the DisplaySpecialsWithinSeasons field if non-nil, zero value otherwise.

### GetDisplaySpecialsWithinSeasonsOk

`func (o *JellyfinServerConfiguration) GetDisplaySpecialsWithinSeasonsOk() (*bool, bool)`

GetDisplaySpecialsWithinSeasonsOk returns a tuple with the DisplaySpecialsWithinSeasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySpecialsWithinSeasons

`func (o *JellyfinServerConfiguration) SetDisplaySpecialsWithinSeasons(v bool)`

SetDisplaySpecialsWithinSeasons sets DisplaySpecialsWithinSeasons field to given value.

### HasDisplaySpecialsWithinSeasons

`func (o *JellyfinServerConfiguration) HasDisplaySpecialsWithinSeasons() bool`

HasDisplaySpecialsWithinSeasons returns a boolean if a field has been set.

### GetCodecsUsed

`func (o *JellyfinServerConfiguration) GetCodecsUsed() []string`

GetCodecsUsed returns the CodecsUsed field if non-nil, zero value otherwise.

### GetCodecsUsedOk

`func (o *JellyfinServerConfiguration) GetCodecsUsedOk() (*[]string, bool)`

GetCodecsUsedOk returns a tuple with the CodecsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecsUsed

`func (o *JellyfinServerConfiguration) SetCodecsUsed(v []string)`

SetCodecsUsed sets CodecsUsed field to given value.

### HasCodecsUsed

`func (o *JellyfinServerConfiguration) HasCodecsUsed() bool`

HasCodecsUsed returns a boolean if a field has been set.

### GetPluginRepositories

`func (o *JellyfinServerConfiguration) GetPluginRepositories() []JellyfinJellyfinRepositoryInfo`

GetPluginRepositories returns the PluginRepositories field if non-nil, zero value otherwise.

### GetPluginRepositoriesOk

`func (o *JellyfinServerConfiguration) GetPluginRepositoriesOk() (*[]JellyfinJellyfinRepositoryInfo, bool)`

GetPluginRepositoriesOk returns a tuple with the PluginRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRepositories

`func (o *JellyfinServerConfiguration) SetPluginRepositories(v []JellyfinJellyfinRepositoryInfo)`

SetPluginRepositories sets PluginRepositories field to given value.

### HasPluginRepositories

`func (o *JellyfinServerConfiguration) HasPluginRepositories() bool`

HasPluginRepositories returns a boolean if a field has been set.

### GetEnableExternalContentInSuggestions

`func (o *JellyfinServerConfiguration) GetEnableExternalContentInSuggestions() bool`

GetEnableExternalContentInSuggestions returns the EnableExternalContentInSuggestions field if non-nil, zero value otherwise.

### GetEnableExternalContentInSuggestionsOk

`func (o *JellyfinServerConfiguration) GetEnableExternalContentInSuggestionsOk() (*bool, bool)`

GetEnableExternalContentInSuggestionsOk returns a tuple with the EnableExternalContentInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalContentInSuggestions

`func (o *JellyfinServerConfiguration) SetEnableExternalContentInSuggestions(v bool)`

SetEnableExternalContentInSuggestions sets EnableExternalContentInSuggestions field to given value.

### HasEnableExternalContentInSuggestions

`func (o *JellyfinServerConfiguration) HasEnableExternalContentInSuggestions() bool`

HasEnableExternalContentInSuggestions returns a boolean if a field has been set.

### GetImageExtractionTimeoutMs

`func (o *JellyfinServerConfiguration) GetImageExtractionTimeoutMs() int32`

GetImageExtractionTimeoutMs returns the ImageExtractionTimeoutMs field if non-nil, zero value otherwise.

### GetImageExtractionTimeoutMsOk

`func (o *JellyfinServerConfiguration) GetImageExtractionTimeoutMsOk() (*int32, bool)`

GetImageExtractionTimeoutMsOk returns a tuple with the ImageExtractionTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractionTimeoutMs

`func (o *JellyfinServerConfiguration) SetImageExtractionTimeoutMs(v int32)`

SetImageExtractionTimeoutMs sets ImageExtractionTimeoutMs field to given value.

### HasImageExtractionTimeoutMs

`func (o *JellyfinServerConfiguration) HasImageExtractionTimeoutMs() bool`

HasImageExtractionTimeoutMs returns a boolean if a field has been set.

### GetPathSubstitutions

`func (o *JellyfinServerConfiguration) GetPathSubstitutions() []JellyfinJellyfinPathSubstitution`

GetPathSubstitutions returns the PathSubstitutions field if non-nil, zero value otherwise.

### GetPathSubstitutionsOk

`func (o *JellyfinServerConfiguration) GetPathSubstitutionsOk() (*[]JellyfinJellyfinPathSubstitution, bool)`

GetPathSubstitutionsOk returns a tuple with the PathSubstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSubstitutions

`func (o *JellyfinServerConfiguration) SetPathSubstitutions(v []JellyfinJellyfinPathSubstitution)`

SetPathSubstitutions sets PathSubstitutions field to given value.

### HasPathSubstitutions

`func (o *JellyfinServerConfiguration) HasPathSubstitutions() bool`

HasPathSubstitutions returns a boolean if a field has been set.

### GetEnableSlowResponseWarning

`func (o *JellyfinServerConfiguration) GetEnableSlowResponseWarning() bool`

GetEnableSlowResponseWarning returns the EnableSlowResponseWarning field if non-nil, zero value otherwise.

### GetEnableSlowResponseWarningOk

`func (o *JellyfinServerConfiguration) GetEnableSlowResponseWarningOk() (*bool, bool)`

GetEnableSlowResponseWarningOk returns a tuple with the EnableSlowResponseWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSlowResponseWarning

`func (o *JellyfinServerConfiguration) SetEnableSlowResponseWarning(v bool)`

SetEnableSlowResponseWarning sets EnableSlowResponseWarning field to given value.

### HasEnableSlowResponseWarning

`func (o *JellyfinServerConfiguration) HasEnableSlowResponseWarning() bool`

HasEnableSlowResponseWarning returns a boolean if a field has been set.

### GetSlowResponseThresholdMs

`func (o *JellyfinServerConfiguration) GetSlowResponseThresholdMs() int64`

GetSlowResponseThresholdMs returns the SlowResponseThresholdMs field if non-nil, zero value otherwise.

### GetSlowResponseThresholdMsOk

`func (o *JellyfinServerConfiguration) GetSlowResponseThresholdMsOk() (*int64, bool)`

GetSlowResponseThresholdMsOk returns a tuple with the SlowResponseThresholdMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowResponseThresholdMs

`func (o *JellyfinServerConfiguration) SetSlowResponseThresholdMs(v int64)`

SetSlowResponseThresholdMs sets SlowResponseThresholdMs field to given value.

### HasSlowResponseThresholdMs

`func (o *JellyfinServerConfiguration) HasSlowResponseThresholdMs() bool`

HasSlowResponseThresholdMs returns a boolean if a field has been set.

### GetCorsHosts

`func (o *JellyfinServerConfiguration) GetCorsHosts() []string`

GetCorsHosts returns the CorsHosts field if non-nil, zero value otherwise.

### GetCorsHostsOk

`func (o *JellyfinServerConfiguration) GetCorsHostsOk() (*[]string, bool)`

GetCorsHostsOk returns a tuple with the CorsHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHosts

`func (o *JellyfinServerConfiguration) SetCorsHosts(v []string)`

SetCorsHosts sets CorsHosts field to given value.

### HasCorsHosts

`func (o *JellyfinServerConfiguration) HasCorsHosts() bool`

HasCorsHosts returns a boolean if a field has been set.

### GetActivityLogRetentionDays

`func (o *JellyfinServerConfiguration) GetActivityLogRetentionDays() int32`

GetActivityLogRetentionDays returns the ActivityLogRetentionDays field if non-nil, zero value otherwise.

### GetActivityLogRetentionDaysOk

`func (o *JellyfinServerConfiguration) GetActivityLogRetentionDaysOk() (*int32, bool)`

GetActivityLogRetentionDaysOk returns a tuple with the ActivityLogRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLogRetentionDays

`func (o *JellyfinServerConfiguration) SetActivityLogRetentionDays(v int32)`

SetActivityLogRetentionDays sets ActivityLogRetentionDays field to given value.

### HasActivityLogRetentionDays

`func (o *JellyfinServerConfiguration) HasActivityLogRetentionDays() bool`

HasActivityLogRetentionDays returns a boolean if a field has been set.

### SetActivityLogRetentionDaysNil

`func (o *JellyfinServerConfiguration) SetActivityLogRetentionDaysNil(b bool)`

 SetActivityLogRetentionDaysNil sets the value for ActivityLogRetentionDays to be an explicit nil

### UnsetActivityLogRetentionDays
`func (o *JellyfinServerConfiguration) UnsetActivityLogRetentionDays()`

UnsetActivityLogRetentionDays ensures that no value is present for ActivityLogRetentionDays, not even an explicit nil
### GetLibraryScanFanoutConcurrency

`func (o *JellyfinServerConfiguration) GetLibraryScanFanoutConcurrency() int32`

GetLibraryScanFanoutConcurrency returns the LibraryScanFanoutConcurrency field if non-nil, zero value otherwise.

### GetLibraryScanFanoutConcurrencyOk

`func (o *JellyfinServerConfiguration) GetLibraryScanFanoutConcurrencyOk() (*int32, bool)`

GetLibraryScanFanoutConcurrencyOk returns a tuple with the LibraryScanFanoutConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryScanFanoutConcurrency

`func (o *JellyfinServerConfiguration) SetLibraryScanFanoutConcurrency(v int32)`

SetLibraryScanFanoutConcurrency sets LibraryScanFanoutConcurrency field to given value.

### HasLibraryScanFanoutConcurrency

`func (o *JellyfinServerConfiguration) HasLibraryScanFanoutConcurrency() bool`

HasLibraryScanFanoutConcurrency returns a boolean if a field has been set.

### GetLibraryMetadataRefreshConcurrency

`func (o *JellyfinServerConfiguration) GetLibraryMetadataRefreshConcurrency() int32`

GetLibraryMetadataRefreshConcurrency returns the LibraryMetadataRefreshConcurrency field if non-nil, zero value otherwise.

### GetLibraryMetadataRefreshConcurrencyOk

`func (o *JellyfinServerConfiguration) GetLibraryMetadataRefreshConcurrencyOk() (*int32, bool)`

GetLibraryMetadataRefreshConcurrencyOk returns a tuple with the LibraryMetadataRefreshConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMetadataRefreshConcurrency

`func (o *JellyfinServerConfiguration) SetLibraryMetadataRefreshConcurrency(v int32)`

SetLibraryMetadataRefreshConcurrency sets LibraryMetadataRefreshConcurrency field to given value.

### HasLibraryMetadataRefreshConcurrency

`func (o *JellyfinServerConfiguration) HasLibraryMetadataRefreshConcurrency() bool`

HasLibraryMetadataRefreshConcurrency returns a boolean if a field has been set.

### GetRemoveOldPlugins

`func (o *JellyfinServerConfiguration) GetRemoveOldPlugins() bool`

GetRemoveOldPlugins returns the RemoveOldPlugins field if non-nil, zero value otherwise.

### GetRemoveOldPluginsOk

`func (o *JellyfinServerConfiguration) GetRemoveOldPluginsOk() (*bool, bool)`

GetRemoveOldPluginsOk returns a tuple with the RemoveOldPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOldPlugins

`func (o *JellyfinServerConfiguration) SetRemoveOldPlugins(v bool)`

SetRemoveOldPlugins sets RemoveOldPlugins field to given value.

### HasRemoveOldPlugins

`func (o *JellyfinServerConfiguration) HasRemoveOldPlugins() bool`

HasRemoveOldPlugins returns a boolean if a field has been set.

### GetAllowClientLogUpload

`func (o *JellyfinServerConfiguration) GetAllowClientLogUpload() bool`

GetAllowClientLogUpload returns the AllowClientLogUpload field if non-nil, zero value otherwise.

### GetAllowClientLogUploadOk

`func (o *JellyfinServerConfiguration) GetAllowClientLogUploadOk() (*bool, bool)`

GetAllowClientLogUploadOk returns a tuple with the AllowClientLogUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClientLogUpload

`func (o *JellyfinServerConfiguration) SetAllowClientLogUpload(v bool)`

SetAllowClientLogUpload sets AllowClientLogUpload field to given value.

### HasAllowClientLogUpload

`func (o *JellyfinServerConfiguration) HasAllowClientLogUpload() bool`

HasAllowClientLogUpload returns a boolean if a field has been set.

### GetDummyChapterDuration

`func (o *JellyfinServerConfiguration) GetDummyChapterDuration() int32`

GetDummyChapterDuration returns the DummyChapterDuration field if non-nil, zero value otherwise.

### GetDummyChapterDurationOk

`func (o *JellyfinServerConfiguration) GetDummyChapterDurationOk() (*int32, bool)`

GetDummyChapterDurationOk returns a tuple with the DummyChapterDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDummyChapterDuration

`func (o *JellyfinServerConfiguration) SetDummyChapterDuration(v int32)`

SetDummyChapterDuration sets DummyChapterDuration field to given value.

### HasDummyChapterDuration

`func (o *JellyfinServerConfiguration) HasDummyChapterDuration() bool`

HasDummyChapterDuration returns a boolean if a field has been set.

### GetChapterImageResolution

`func (o *JellyfinServerConfiguration) GetChapterImageResolution() JellyfinJellyfinImageResolution`

GetChapterImageResolution returns the ChapterImageResolution field if non-nil, zero value otherwise.

### GetChapterImageResolutionOk

`func (o *JellyfinServerConfiguration) GetChapterImageResolutionOk() (*JellyfinJellyfinImageResolution, bool)`

GetChapterImageResolutionOk returns a tuple with the ChapterImageResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterImageResolution

`func (o *JellyfinServerConfiguration) SetChapterImageResolution(v JellyfinJellyfinImageResolution)`

SetChapterImageResolution sets ChapterImageResolution field to given value.

### HasChapterImageResolution

`func (o *JellyfinServerConfiguration) HasChapterImageResolution() bool`

HasChapterImageResolution returns a boolean if a field has been set.

### GetParallelImageEncodingLimit

`func (o *JellyfinServerConfiguration) GetParallelImageEncodingLimit() int32`

GetParallelImageEncodingLimit returns the ParallelImageEncodingLimit field if non-nil, zero value otherwise.

### GetParallelImageEncodingLimitOk

`func (o *JellyfinServerConfiguration) GetParallelImageEncodingLimitOk() (*int32, bool)`

GetParallelImageEncodingLimitOk returns a tuple with the ParallelImageEncodingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelImageEncodingLimit

`func (o *JellyfinServerConfiguration) SetParallelImageEncodingLimit(v int32)`

SetParallelImageEncodingLimit sets ParallelImageEncodingLimit field to given value.

### HasParallelImageEncodingLimit

`func (o *JellyfinServerConfiguration) HasParallelImageEncodingLimit() bool`

HasParallelImageEncodingLimit returns a boolean if a field has been set.

### GetCastReceiverApplications

`func (o *JellyfinServerConfiguration) GetCastReceiverApplications() []JellyfinJellyfinCastReceiverApplication`

GetCastReceiverApplications returns the CastReceiverApplications field if non-nil, zero value otherwise.

### GetCastReceiverApplicationsOk

`func (o *JellyfinServerConfiguration) GetCastReceiverApplicationsOk() (*[]JellyfinJellyfinCastReceiverApplication, bool)`

GetCastReceiverApplicationsOk returns a tuple with the CastReceiverApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReceiverApplications

`func (o *JellyfinServerConfiguration) SetCastReceiverApplications(v []JellyfinJellyfinCastReceiverApplication)`

SetCastReceiverApplications sets CastReceiverApplications field to given value.

### HasCastReceiverApplications

`func (o *JellyfinServerConfiguration) HasCastReceiverApplications() bool`

HasCastReceiverApplications returns a boolean if a field has been set.

### GetTrickplayOptions

`func (o *JellyfinServerConfiguration) GetTrickplayOptions() JellyfinJellyfinTrickplayOptions`

GetTrickplayOptions returns the TrickplayOptions field if non-nil, zero value otherwise.

### GetTrickplayOptionsOk

`func (o *JellyfinServerConfiguration) GetTrickplayOptionsOk() (*JellyfinJellyfinTrickplayOptions, bool)`

GetTrickplayOptionsOk returns a tuple with the TrickplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrickplayOptions

`func (o *JellyfinServerConfiguration) SetTrickplayOptions(v JellyfinJellyfinTrickplayOptions)`

SetTrickplayOptions sets TrickplayOptions field to given value.

### HasTrickplayOptions

`func (o *JellyfinServerConfiguration) HasTrickplayOptions() bool`

HasTrickplayOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


