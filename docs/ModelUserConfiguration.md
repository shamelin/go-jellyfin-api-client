# ModelUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioLanguagePreference** | Pointer to **NullableString** | Gets or sets the audio language preference. | [optional] 
**PlayDefaultAudioTrack** | Pointer to **bool** | Gets or sets a value indicating whether [play default audio track]. | [optional] 
**SubtitleLanguagePreference** | Pointer to **NullableString** | Gets or sets the subtitle language preference. | [optional] 
**DisplayMissingEpisodes** | Pointer to **bool** |  | [optional] 
**GroupedFolders** | Pointer to **[]string** |  | [optional] 
**SubtitleMode** | Pointer to [**ModelModelSubtitlePlaybackMode**](ModelSubtitlePlaybackMode.md) | An enum representing a subtitle playback mode. | [optional] 
**DisplayCollectionsView** | Pointer to **bool** |  | [optional] 
**EnableLocalPassword** | Pointer to **bool** |  | [optional] 
**OrderedViews** | Pointer to **[]string** |  | [optional] 
**LatestItemsExcludes** | Pointer to **[]string** |  | [optional] 
**MyMediaExcludes** | Pointer to **[]string** |  | [optional] 
**HidePlayedInLatest** | Pointer to **bool** |  | [optional] 
**RememberAudioSelections** | Pointer to **bool** |  | [optional] 
**RememberSubtitleSelections** | Pointer to **bool** |  | [optional] 
**EnableNextEpisodeAutoPlay** | Pointer to **bool** |  | [optional] 
**CastReceiverId** | Pointer to **NullableString** | Gets or sets the id of the selected cast receiver. | [optional] 

## Methods

### NewModelUserConfiguration

`func NewModelUserConfiguration() *ModelUserConfiguration`

NewModelUserConfiguration instantiates a new ModelUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserConfigurationWithDefaults

`func NewModelUserConfigurationWithDefaults() *ModelUserConfiguration`

NewModelUserConfigurationWithDefaults instantiates a new ModelUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioLanguagePreference

`func (o *ModelUserConfiguration) GetAudioLanguagePreference() string`

GetAudioLanguagePreference returns the AudioLanguagePreference field if non-nil, zero value otherwise.

### GetAudioLanguagePreferenceOk

`func (o *ModelUserConfiguration) GetAudioLanguagePreferenceOk() (*string, bool)`

GetAudioLanguagePreferenceOk returns a tuple with the AudioLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguagePreference

`func (o *ModelUserConfiguration) SetAudioLanguagePreference(v string)`

SetAudioLanguagePreference sets AudioLanguagePreference field to given value.

### HasAudioLanguagePreference

`func (o *ModelUserConfiguration) HasAudioLanguagePreference() bool`

HasAudioLanguagePreference returns a boolean if a field has been set.

### SetAudioLanguagePreferenceNil

`func (o *ModelUserConfiguration) SetAudioLanguagePreferenceNil(b bool)`

 SetAudioLanguagePreferenceNil sets the value for AudioLanguagePreference to be an explicit nil

### UnsetAudioLanguagePreference
`func (o *ModelUserConfiguration) UnsetAudioLanguagePreference()`

UnsetAudioLanguagePreference ensures that no value is present for AudioLanguagePreference, not even an explicit nil
### GetPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) GetPlayDefaultAudioTrack() bool`

GetPlayDefaultAudioTrack returns the PlayDefaultAudioTrack field if non-nil, zero value otherwise.

### GetPlayDefaultAudioTrackOk

`func (o *ModelUserConfiguration) GetPlayDefaultAudioTrackOk() (*bool, bool)`

GetPlayDefaultAudioTrackOk returns a tuple with the PlayDefaultAudioTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) SetPlayDefaultAudioTrack(v bool)`

SetPlayDefaultAudioTrack sets PlayDefaultAudioTrack field to given value.

### HasPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) HasPlayDefaultAudioTrack() bool`

HasPlayDefaultAudioTrack returns a boolean if a field has been set.

### GetSubtitleLanguagePreference

`func (o *ModelUserConfiguration) GetSubtitleLanguagePreference() string`

GetSubtitleLanguagePreference returns the SubtitleLanguagePreference field if non-nil, zero value otherwise.

### GetSubtitleLanguagePreferenceOk

`func (o *ModelUserConfiguration) GetSubtitleLanguagePreferenceOk() (*string, bool)`

GetSubtitleLanguagePreferenceOk returns a tuple with the SubtitleLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguagePreference

`func (o *ModelUserConfiguration) SetSubtitleLanguagePreference(v string)`

SetSubtitleLanguagePreference sets SubtitleLanguagePreference field to given value.

### HasSubtitleLanguagePreference

`func (o *ModelUserConfiguration) HasSubtitleLanguagePreference() bool`

HasSubtitleLanguagePreference returns a boolean if a field has been set.

### SetSubtitleLanguagePreferenceNil

`func (o *ModelUserConfiguration) SetSubtitleLanguagePreferenceNil(b bool)`

 SetSubtitleLanguagePreferenceNil sets the value for SubtitleLanguagePreference to be an explicit nil

### UnsetSubtitleLanguagePreference
`func (o *ModelUserConfiguration) UnsetSubtitleLanguagePreference()`

UnsetSubtitleLanguagePreference ensures that no value is present for SubtitleLanguagePreference, not even an explicit nil
### GetDisplayMissingEpisodes

`func (o *ModelUserConfiguration) GetDisplayMissingEpisodes() bool`

GetDisplayMissingEpisodes returns the DisplayMissingEpisodes field if non-nil, zero value otherwise.

### GetDisplayMissingEpisodesOk

`func (o *ModelUserConfiguration) GetDisplayMissingEpisodesOk() (*bool, bool)`

GetDisplayMissingEpisodesOk returns a tuple with the DisplayMissingEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMissingEpisodes

`func (o *ModelUserConfiguration) SetDisplayMissingEpisodes(v bool)`

SetDisplayMissingEpisodes sets DisplayMissingEpisodes field to given value.

### HasDisplayMissingEpisodes

`func (o *ModelUserConfiguration) HasDisplayMissingEpisodes() bool`

HasDisplayMissingEpisodes returns a boolean if a field has been set.

### GetGroupedFolders

`func (o *ModelUserConfiguration) GetGroupedFolders() []string`

GetGroupedFolders returns the GroupedFolders field if non-nil, zero value otherwise.

### GetGroupedFoldersOk

`func (o *ModelUserConfiguration) GetGroupedFoldersOk() (*[]string, bool)`

GetGroupedFoldersOk returns a tuple with the GroupedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedFolders

`func (o *ModelUserConfiguration) SetGroupedFolders(v []string)`

SetGroupedFolders sets GroupedFolders field to given value.

### HasGroupedFolders

`func (o *ModelUserConfiguration) HasGroupedFolders() bool`

HasGroupedFolders returns a boolean if a field has been set.

### GetSubtitleMode

`func (o *ModelUserConfiguration) GetSubtitleMode() ModelModelSubtitlePlaybackMode`

GetSubtitleMode returns the SubtitleMode field if non-nil, zero value otherwise.

### GetSubtitleModeOk

`func (o *ModelUserConfiguration) GetSubtitleModeOk() (*ModelModelSubtitlePlaybackMode, bool)`

GetSubtitleModeOk returns a tuple with the SubtitleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleMode

`func (o *ModelUserConfiguration) SetSubtitleMode(v ModelModelSubtitlePlaybackMode)`

SetSubtitleMode sets SubtitleMode field to given value.

### HasSubtitleMode

`func (o *ModelUserConfiguration) HasSubtitleMode() bool`

HasSubtitleMode returns a boolean if a field has been set.

### GetDisplayCollectionsView

`func (o *ModelUserConfiguration) GetDisplayCollectionsView() bool`

GetDisplayCollectionsView returns the DisplayCollectionsView field if non-nil, zero value otherwise.

### GetDisplayCollectionsViewOk

`func (o *ModelUserConfiguration) GetDisplayCollectionsViewOk() (*bool, bool)`

GetDisplayCollectionsViewOk returns a tuple with the DisplayCollectionsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCollectionsView

`func (o *ModelUserConfiguration) SetDisplayCollectionsView(v bool)`

SetDisplayCollectionsView sets DisplayCollectionsView field to given value.

### HasDisplayCollectionsView

`func (o *ModelUserConfiguration) HasDisplayCollectionsView() bool`

HasDisplayCollectionsView returns a boolean if a field has been set.

### GetEnableLocalPassword

`func (o *ModelUserConfiguration) GetEnableLocalPassword() bool`

GetEnableLocalPassword returns the EnableLocalPassword field if non-nil, zero value otherwise.

### GetEnableLocalPasswordOk

`func (o *ModelUserConfiguration) GetEnableLocalPasswordOk() (*bool, bool)`

GetEnableLocalPasswordOk returns a tuple with the EnableLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalPassword

`func (o *ModelUserConfiguration) SetEnableLocalPassword(v bool)`

SetEnableLocalPassword sets EnableLocalPassword field to given value.

### HasEnableLocalPassword

`func (o *ModelUserConfiguration) HasEnableLocalPassword() bool`

HasEnableLocalPassword returns a boolean if a field has been set.

### GetOrderedViews

`func (o *ModelUserConfiguration) GetOrderedViews() []string`

GetOrderedViews returns the OrderedViews field if non-nil, zero value otherwise.

### GetOrderedViewsOk

`func (o *ModelUserConfiguration) GetOrderedViewsOk() (*[]string, bool)`

GetOrderedViewsOk returns a tuple with the OrderedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedViews

`func (o *ModelUserConfiguration) SetOrderedViews(v []string)`

SetOrderedViews sets OrderedViews field to given value.

### HasOrderedViews

`func (o *ModelUserConfiguration) HasOrderedViews() bool`

HasOrderedViews returns a boolean if a field has been set.

### GetLatestItemsExcludes

`func (o *ModelUserConfiguration) GetLatestItemsExcludes() []string`

GetLatestItemsExcludes returns the LatestItemsExcludes field if non-nil, zero value otherwise.

### GetLatestItemsExcludesOk

`func (o *ModelUserConfiguration) GetLatestItemsExcludesOk() (*[]string, bool)`

GetLatestItemsExcludesOk returns a tuple with the LatestItemsExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestItemsExcludes

`func (o *ModelUserConfiguration) SetLatestItemsExcludes(v []string)`

SetLatestItemsExcludes sets LatestItemsExcludes field to given value.

### HasLatestItemsExcludes

`func (o *ModelUserConfiguration) HasLatestItemsExcludes() bool`

HasLatestItemsExcludes returns a boolean if a field has been set.

### GetMyMediaExcludes

`func (o *ModelUserConfiguration) GetMyMediaExcludes() []string`

GetMyMediaExcludes returns the MyMediaExcludes field if non-nil, zero value otherwise.

### GetMyMediaExcludesOk

`func (o *ModelUserConfiguration) GetMyMediaExcludesOk() (*[]string, bool)`

GetMyMediaExcludesOk returns a tuple with the MyMediaExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyMediaExcludes

`func (o *ModelUserConfiguration) SetMyMediaExcludes(v []string)`

SetMyMediaExcludes sets MyMediaExcludes field to given value.

### HasMyMediaExcludes

`func (o *ModelUserConfiguration) HasMyMediaExcludes() bool`

HasMyMediaExcludes returns a boolean if a field has been set.

### GetHidePlayedInLatest

`func (o *ModelUserConfiguration) GetHidePlayedInLatest() bool`

GetHidePlayedInLatest returns the HidePlayedInLatest field if non-nil, zero value otherwise.

### GetHidePlayedInLatestOk

`func (o *ModelUserConfiguration) GetHidePlayedInLatestOk() (*bool, bool)`

GetHidePlayedInLatestOk returns a tuple with the HidePlayedInLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInLatest

`func (o *ModelUserConfiguration) SetHidePlayedInLatest(v bool)`

SetHidePlayedInLatest sets HidePlayedInLatest field to given value.

### HasHidePlayedInLatest

`func (o *ModelUserConfiguration) HasHidePlayedInLatest() bool`

HasHidePlayedInLatest returns a boolean if a field has been set.

### GetRememberAudioSelections

`func (o *ModelUserConfiguration) GetRememberAudioSelections() bool`

GetRememberAudioSelections returns the RememberAudioSelections field if non-nil, zero value otherwise.

### GetRememberAudioSelectionsOk

`func (o *ModelUserConfiguration) GetRememberAudioSelectionsOk() (*bool, bool)`

GetRememberAudioSelectionsOk returns a tuple with the RememberAudioSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberAudioSelections

`func (o *ModelUserConfiguration) SetRememberAudioSelections(v bool)`

SetRememberAudioSelections sets RememberAudioSelections field to given value.

### HasRememberAudioSelections

`func (o *ModelUserConfiguration) HasRememberAudioSelections() bool`

HasRememberAudioSelections returns a boolean if a field has been set.

### GetRememberSubtitleSelections

`func (o *ModelUserConfiguration) GetRememberSubtitleSelections() bool`

GetRememberSubtitleSelections returns the RememberSubtitleSelections field if non-nil, zero value otherwise.

### GetRememberSubtitleSelectionsOk

`func (o *ModelUserConfiguration) GetRememberSubtitleSelectionsOk() (*bool, bool)`

GetRememberSubtitleSelectionsOk returns a tuple with the RememberSubtitleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSubtitleSelections

`func (o *ModelUserConfiguration) SetRememberSubtitleSelections(v bool)`

SetRememberSubtitleSelections sets RememberSubtitleSelections field to given value.

### HasRememberSubtitleSelections

`func (o *ModelUserConfiguration) HasRememberSubtitleSelections() bool`

HasRememberSubtitleSelections returns a boolean if a field has been set.

### GetEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) GetEnableNextEpisodeAutoPlay() bool`

GetEnableNextEpisodeAutoPlay returns the EnableNextEpisodeAutoPlay field if non-nil, zero value otherwise.

### GetEnableNextEpisodeAutoPlayOk

`func (o *ModelUserConfiguration) GetEnableNextEpisodeAutoPlayOk() (*bool, bool)`

GetEnableNextEpisodeAutoPlayOk returns a tuple with the EnableNextEpisodeAutoPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) SetEnableNextEpisodeAutoPlay(v bool)`

SetEnableNextEpisodeAutoPlay sets EnableNextEpisodeAutoPlay field to given value.

### HasEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) HasEnableNextEpisodeAutoPlay() bool`

HasEnableNextEpisodeAutoPlay returns a boolean if a field has been set.

### GetCastReceiverId

`func (o *ModelUserConfiguration) GetCastReceiverId() string`

GetCastReceiverId returns the CastReceiverId field if non-nil, zero value otherwise.

### GetCastReceiverIdOk

`func (o *ModelUserConfiguration) GetCastReceiverIdOk() (*string, bool)`

GetCastReceiverIdOk returns a tuple with the CastReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReceiverId

`func (o *ModelUserConfiguration) SetCastReceiverId(v string)`

SetCastReceiverId sets CastReceiverId field to given value.

### HasCastReceiverId

`func (o *ModelUserConfiguration) HasCastReceiverId() bool`

HasCastReceiverId returns a boolean if a field has been set.

### SetCastReceiverIdNil

`func (o *ModelUserConfiguration) SetCastReceiverIdNil(b bool)`

 SetCastReceiverIdNil sets the value for CastReceiverId to be an explicit nil

### UnsetCastReceiverId
`func (o *ModelUserConfiguration) UnsetCastReceiverId()`

UnsetCastReceiverId ensures that no value is present for CastReceiverId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


