# JellyfinUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioLanguagePreference** | Pointer to **NullableString** | Gets or sets the audio language preference. | [optional] 
**PlayDefaultAudioTrack** | Pointer to **bool** | Gets or sets a value indicating whether [play default audio track]. | [optional] 
**SubtitleLanguagePreference** | Pointer to **NullableString** | Gets or sets the subtitle language preference. | [optional] 
**DisplayMissingEpisodes** | Pointer to **bool** |  | [optional] 
**GroupedFolders** | Pointer to **[]string** |  | [optional] 
**SubtitleMode** | Pointer to [**JellyfinJellyfinSubtitlePlaybackMode**](JellyfinSubtitlePlaybackMode.md) | An enum representing a subtitle playback mode. | [optional] 
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

### NewJellyfinUserConfiguration

`func NewJellyfinUserConfiguration() *JellyfinUserConfiguration`

NewJellyfinUserConfiguration instantiates a new JellyfinUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUserConfigurationWithDefaults

`func NewJellyfinUserConfigurationWithDefaults() *JellyfinUserConfiguration`

NewJellyfinUserConfigurationWithDefaults instantiates a new JellyfinUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioLanguagePreference

`func (o *JellyfinUserConfiguration) GetAudioLanguagePreference() string`

GetAudioLanguagePreference returns the AudioLanguagePreference field if non-nil, zero value otherwise.

### GetAudioLanguagePreferenceOk

`func (o *JellyfinUserConfiguration) GetAudioLanguagePreferenceOk() (*string, bool)`

GetAudioLanguagePreferenceOk returns a tuple with the AudioLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguagePreference

`func (o *JellyfinUserConfiguration) SetAudioLanguagePreference(v string)`

SetAudioLanguagePreference sets AudioLanguagePreference field to given value.

### HasAudioLanguagePreference

`func (o *JellyfinUserConfiguration) HasAudioLanguagePreference() bool`

HasAudioLanguagePreference returns a boolean if a field has been set.

### SetAudioLanguagePreferenceNil

`func (o *JellyfinUserConfiguration) SetAudioLanguagePreferenceNil(b bool)`

 SetAudioLanguagePreferenceNil sets the value for AudioLanguagePreference to be an explicit nil

### UnsetAudioLanguagePreference
`func (o *JellyfinUserConfiguration) UnsetAudioLanguagePreference()`

UnsetAudioLanguagePreference ensures that no value is present for AudioLanguagePreference, not even an explicit nil
### GetPlayDefaultAudioTrack

`func (o *JellyfinUserConfiguration) GetPlayDefaultAudioTrack() bool`

GetPlayDefaultAudioTrack returns the PlayDefaultAudioTrack field if non-nil, zero value otherwise.

### GetPlayDefaultAudioTrackOk

`func (o *JellyfinUserConfiguration) GetPlayDefaultAudioTrackOk() (*bool, bool)`

GetPlayDefaultAudioTrackOk returns a tuple with the PlayDefaultAudioTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayDefaultAudioTrack

`func (o *JellyfinUserConfiguration) SetPlayDefaultAudioTrack(v bool)`

SetPlayDefaultAudioTrack sets PlayDefaultAudioTrack field to given value.

### HasPlayDefaultAudioTrack

`func (o *JellyfinUserConfiguration) HasPlayDefaultAudioTrack() bool`

HasPlayDefaultAudioTrack returns a boolean if a field has been set.

### GetSubtitleLanguagePreference

`func (o *JellyfinUserConfiguration) GetSubtitleLanguagePreference() string`

GetSubtitleLanguagePreference returns the SubtitleLanguagePreference field if non-nil, zero value otherwise.

### GetSubtitleLanguagePreferenceOk

`func (o *JellyfinUserConfiguration) GetSubtitleLanguagePreferenceOk() (*string, bool)`

GetSubtitleLanguagePreferenceOk returns a tuple with the SubtitleLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguagePreference

`func (o *JellyfinUserConfiguration) SetSubtitleLanguagePreference(v string)`

SetSubtitleLanguagePreference sets SubtitleLanguagePreference field to given value.

### HasSubtitleLanguagePreference

`func (o *JellyfinUserConfiguration) HasSubtitleLanguagePreference() bool`

HasSubtitleLanguagePreference returns a boolean if a field has been set.

### SetSubtitleLanguagePreferenceNil

`func (o *JellyfinUserConfiguration) SetSubtitleLanguagePreferenceNil(b bool)`

 SetSubtitleLanguagePreferenceNil sets the value for SubtitleLanguagePreference to be an explicit nil

### UnsetSubtitleLanguagePreference
`func (o *JellyfinUserConfiguration) UnsetSubtitleLanguagePreference()`

UnsetSubtitleLanguagePreference ensures that no value is present for SubtitleLanguagePreference, not even an explicit nil
### GetDisplayMissingEpisodes

`func (o *JellyfinUserConfiguration) GetDisplayMissingEpisodes() bool`

GetDisplayMissingEpisodes returns the DisplayMissingEpisodes field if non-nil, zero value otherwise.

### GetDisplayMissingEpisodesOk

`func (o *JellyfinUserConfiguration) GetDisplayMissingEpisodesOk() (*bool, bool)`

GetDisplayMissingEpisodesOk returns a tuple with the DisplayMissingEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMissingEpisodes

`func (o *JellyfinUserConfiguration) SetDisplayMissingEpisodes(v bool)`

SetDisplayMissingEpisodes sets DisplayMissingEpisodes field to given value.

### HasDisplayMissingEpisodes

`func (o *JellyfinUserConfiguration) HasDisplayMissingEpisodes() bool`

HasDisplayMissingEpisodes returns a boolean if a field has been set.

### GetGroupedFolders

`func (o *JellyfinUserConfiguration) GetGroupedFolders() []string`

GetGroupedFolders returns the GroupedFolders field if non-nil, zero value otherwise.

### GetGroupedFoldersOk

`func (o *JellyfinUserConfiguration) GetGroupedFoldersOk() (*[]string, bool)`

GetGroupedFoldersOk returns a tuple with the GroupedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedFolders

`func (o *JellyfinUserConfiguration) SetGroupedFolders(v []string)`

SetGroupedFolders sets GroupedFolders field to given value.

### HasGroupedFolders

`func (o *JellyfinUserConfiguration) HasGroupedFolders() bool`

HasGroupedFolders returns a boolean if a field has been set.

### GetSubtitleMode

`func (o *JellyfinUserConfiguration) GetSubtitleMode() JellyfinJellyfinSubtitlePlaybackMode`

GetSubtitleMode returns the SubtitleMode field if non-nil, zero value otherwise.

### GetSubtitleModeOk

`func (o *JellyfinUserConfiguration) GetSubtitleModeOk() (*JellyfinJellyfinSubtitlePlaybackMode, bool)`

GetSubtitleModeOk returns a tuple with the SubtitleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleMode

`func (o *JellyfinUserConfiguration) SetSubtitleMode(v JellyfinJellyfinSubtitlePlaybackMode)`

SetSubtitleMode sets SubtitleMode field to given value.

### HasSubtitleMode

`func (o *JellyfinUserConfiguration) HasSubtitleMode() bool`

HasSubtitleMode returns a boolean if a field has been set.

### GetDisplayCollectionsView

`func (o *JellyfinUserConfiguration) GetDisplayCollectionsView() bool`

GetDisplayCollectionsView returns the DisplayCollectionsView field if non-nil, zero value otherwise.

### GetDisplayCollectionsViewOk

`func (o *JellyfinUserConfiguration) GetDisplayCollectionsViewOk() (*bool, bool)`

GetDisplayCollectionsViewOk returns a tuple with the DisplayCollectionsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCollectionsView

`func (o *JellyfinUserConfiguration) SetDisplayCollectionsView(v bool)`

SetDisplayCollectionsView sets DisplayCollectionsView field to given value.

### HasDisplayCollectionsView

`func (o *JellyfinUserConfiguration) HasDisplayCollectionsView() bool`

HasDisplayCollectionsView returns a boolean if a field has been set.

### GetEnableLocalPassword

`func (o *JellyfinUserConfiguration) GetEnableLocalPassword() bool`

GetEnableLocalPassword returns the EnableLocalPassword field if non-nil, zero value otherwise.

### GetEnableLocalPasswordOk

`func (o *JellyfinUserConfiguration) GetEnableLocalPasswordOk() (*bool, bool)`

GetEnableLocalPasswordOk returns a tuple with the EnableLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalPassword

`func (o *JellyfinUserConfiguration) SetEnableLocalPassword(v bool)`

SetEnableLocalPassword sets EnableLocalPassword field to given value.

### HasEnableLocalPassword

`func (o *JellyfinUserConfiguration) HasEnableLocalPassword() bool`

HasEnableLocalPassword returns a boolean if a field has been set.

### GetOrderedViews

`func (o *JellyfinUserConfiguration) GetOrderedViews() []string`

GetOrderedViews returns the OrderedViews field if non-nil, zero value otherwise.

### GetOrderedViewsOk

`func (o *JellyfinUserConfiguration) GetOrderedViewsOk() (*[]string, bool)`

GetOrderedViewsOk returns a tuple with the OrderedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedViews

`func (o *JellyfinUserConfiguration) SetOrderedViews(v []string)`

SetOrderedViews sets OrderedViews field to given value.

### HasOrderedViews

`func (o *JellyfinUserConfiguration) HasOrderedViews() bool`

HasOrderedViews returns a boolean if a field has been set.

### GetLatestItemsExcludes

`func (o *JellyfinUserConfiguration) GetLatestItemsExcludes() []string`

GetLatestItemsExcludes returns the LatestItemsExcludes field if non-nil, zero value otherwise.

### GetLatestItemsExcludesOk

`func (o *JellyfinUserConfiguration) GetLatestItemsExcludesOk() (*[]string, bool)`

GetLatestItemsExcludesOk returns a tuple with the LatestItemsExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestItemsExcludes

`func (o *JellyfinUserConfiguration) SetLatestItemsExcludes(v []string)`

SetLatestItemsExcludes sets LatestItemsExcludes field to given value.

### HasLatestItemsExcludes

`func (o *JellyfinUserConfiguration) HasLatestItemsExcludes() bool`

HasLatestItemsExcludes returns a boolean if a field has been set.

### GetMyMediaExcludes

`func (o *JellyfinUserConfiguration) GetMyMediaExcludes() []string`

GetMyMediaExcludes returns the MyMediaExcludes field if non-nil, zero value otherwise.

### GetMyMediaExcludesOk

`func (o *JellyfinUserConfiguration) GetMyMediaExcludesOk() (*[]string, bool)`

GetMyMediaExcludesOk returns a tuple with the MyMediaExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyMediaExcludes

`func (o *JellyfinUserConfiguration) SetMyMediaExcludes(v []string)`

SetMyMediaExcludes sets MyMediaExcludes field to given value.

### HasMyMediaExcludes

`func (o *JellyfinUserConfiguration) HasMyMediaExcludes() bool`

HasMyMediaExcludes returns a boolean if a field has been set.

### GetHidePlayedInLatest

`func (o *JellyfinUserConfiguration) GetHidePlayedInLatest() bool`

GetHidePlayedInLatest returns the HidePlayedInLatest field if non-nil, zero value otherwise.

### GetHidePlayedInLatestOk

`func (o *JellyfinUserConfiguration) GetHidePlayedInLatestOk() (*bool, bool)`

GetHidePlayedInLatestOk returns a tuple with the HidePlayedInLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInLatest

`func (o *JellyfinUserConfiguration) SetHidePlayedInLatest(v bool)`

SetHidePlayedInLatest sets HidePlayedInLatest field to given value.

### HasHidePlayedInLatest

`func (o *JellyfinUserConfiguration) HasHidePlayedInLatest() bool`

HasHidePlayedInLatest returns a boolean if a field has been set.

### GetRememberAudioSelections

`func (o *JellyfinUserConfiguration) GetRememberAudioSelections() bool`

GetRememberAudioSelections returns the RememberAudioSelections field if non-nil, zero value otherwise.

### GetRememberAudioSelectionsOk

`func (o *JellyfinUserConfiguration) GetRememberAudioSelectionsOk() (*bool, bool)`

GetRememberAudioSelectionsOk returns a tuple with the RememberAudioSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberAudioSelections

`func (o *JellyfinUserConfiguration) SetRememberAudioSelections(v bool)`

SetRememberAudioSelections sets RememberAudioSelections field to given value.

### HasRememberAudioSelections

`func (o *JellyfinUserConfiguration) HasRememberAudioSelections() bool`

HasRememberAudioSelections returns a boolean if a field has been set.

### GetRememberSubtitleSelections

`func (o *JellyfinUserConfiguration) GetRememberSubtitleSelections() bool`

GetRememberSubtitleSelections returns the RememberSubtitleSelections field if non-nil, zero value otherwise.

### GetRememberSubtitleSelectionsOk

`func (o *JellyfinUserConfiguration) GetRememberSubtitleSelectionsOk() (*bool, bool)`

GetRememberSubtitleSelectionsOk returns a tuple with the RememberSubtitleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSubtitleSelections

`func (o *JellyfinUserConfiguration) SetRememberSubtitleSelections(v bool)`

SetRememberSubtitleSelections sets RememberSubtitleSelections field to given value.

### HasRememberSubtitleSelections

`func (o *JellyfinUserConfiguration) HasRememberSubtitleSelections() bool`

HasRememberSubtitleSelections returns a boolean if a field has been set.

### GetEnableNextEpisodeAutoPlay

`func (o *JellyfinUserConfiguration) GetEnableNextEpisodeAutoPlay() bool`

GetEnableNextEpisodeAutoPlay returns the EnableNextEpisodeAutoPlay field if non-nil, zero value otherwise.

### GetEnableNextEpisodeAutoPlayOk

`func (o *JellyfinUserConfiguration) GetEnableNextEpisodeAutoPlayOk() (*bool, bool)`

GetEnableNextEpisodeAutoPlayOk returns a tuple with the EnableNextEpisodeAutoPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNextEpisodeAutoPlay

`func (o *JellyfinUserConfiguration) SetEnableNextEpisodeAutoPlay(v bool)`

SetEnableNextEpisodeAutoPlay sets EnableNextEpisodeAutoPlay field to given value.

### HasEnableNextEpisodeAutoPlay

`func (o *JellyfinUserConfiguration) HasEnableNextEpisodeAutoPlay() bool`

HasEnableNextEpisodeAutoPlay returns a boolean if a field has been set.

### GetCastReceiverId

`func (o *JellyfinUserConfiguration) GetCastReceiverId() string`

GetCastReceiverId returns the CastReceiverId field if non-nil, zero value otherwise.

### GetCastReceiverIdOk

`func (o *JellyfinUserConfiguration) GetCastReceiverIdOk() (*string, bool)`

GetCastReceiverIdOk returns a tuple with the CastReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReceiverId

`func (o *JellyfinUserConfiguration) SetCastReceiverId(v string)`

SetCastReceiverId sets CastReceiverId field to given value.

### HasCastReceiverId

`func (o *JellyfinUserConfiguration) HasCastReceiverId() bool`

HasCastReceiverId returns a boolean if a field has been set.

### SetCastReceiverIdNil

`func (o *JellyfinUserConfiguration) SetCastReceiverIdNil(b bool)`

 SetCastReceiverIdNil sets the value for CastReceiverId to be an explicit nil

### UnsetCastReceiverId
`func (o *JellyfinUserConfiguration) UnsetCastReceiverId()`

UnsetCastReceiverId ensures that no value is present for CastReceiverId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


