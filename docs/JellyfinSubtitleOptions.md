# JellyfinSubtitleOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**DownloadLanguages** | Pointer to **[]string** |  | [optional] 
**DownloadMovieSubtitles** | Pointer to **bool** |  | [optional] 
**DownloadEpisodeSubtitles** | Pointer to **bool** |  | [optional] 
**OpenSubtitlesUsername** | Pointer to **NullableString** |  | [optional] 
**OpenSubtitlesPasswordHash** | Pointer to **NullableString** |  | [optional] 
**IsOpenSubtitleVipAccount** | Pointer to **bool** |  | [optional] 
**RequirePerfectMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewJellyfinSubtitleOptions

`func NewJellyfinSubtitleOptions() *JellyfinSubtitleOptions`

NewJellyfinSubtitleOptions instantiates a new JellyfinSubtitleOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSubtitleOptionsWithDefaults

`func NewJellyfinSubtitleOptionsWithDefaults() *JellyfinSubtitleOptions`

NewJellyfinSubtitleOptionsWithDefaults instantiates a new JellyfinSubtitleOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipIfEmbeddedSubtitlesPresent

`func (o *JellyfinSubtitleOptions) GetSkipIfEmbeddedSubtitlesPresent() bool`

GetSkipIfEmbeddedSubtitlesPresent returns the SkipIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipIfEmbeddedSubtitlesPresentOk

`func (o *JellyfinSubtitleOptions) GetSkipIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfEmbeddedSubtitlesPresent

`func (o *JellyfinSubtitleOptions) SetSkipIfEmbeddedSubtitlesPresent(v bool)`

SetSkipIfEmbeddedSubtitlesPresent sets SkipIfEmbeddedSubtitlesPresent field to given value.

### HasSkipIfEmbeddedSubtitlesPresent

`func (o *JellyfinSubtitleOptions) HasSkipIfEmbeddedSubtitlesPresent() bool`

HasSkipIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipIfAudioTrackMatches

`func (o *JellyfinSubtitleOptions) GetSkipIfAudioTrackMatches() bool`

GetSkipIfAudioTrackMatches returns the SkipIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipIfAudioTrackMatchesOk

`func (o *JellyfinSubtitleOptions) GetSkipIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipIfAudioTrackMatchesOk returns a tuple with the SkipIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfAudioTrackMatches

`func (o *JellyfinSubtitleOptions) SetSkipIfAudioTrackMatches(v bool)`

SetSkipIfAudioTrackMatches sets SkipIfAudioTrackMatches field to given value.

### HasSkipIfAudioTrackMatches

`func (o *JellyfinSubtitleOptions) HasSkipIfAudioTrackMatches() bool`

HasSkipIfAudioTrackMatches returns a boolean if a field has been set.

### GetDownloadLanguages

`func (o *JellyfinSubtitleOptions) GetDownloadLanguages() []string`

GetDownloadLanguages returns the DownloadLanguages field if non-nil, zero value otherwise.

### GetDownloadLanguagesOk

`func (o *JellyfinSubtitleOptions) GetDownloadLanguagesOk() (*[]string, bool)`

GetDownloadLanguagesOk returns a tuple with the DownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLanguages

`func (o *JellyfinSubtitleOptions) SetDownloadLanguages(v []string)`

SetDownloadLanguages sets DownloadLanguages field to given value.

### HasDownloadLanguages

`func (o *JellyfinSubtitleOptions) HasDownloadLanguages() bool`

HasDownloadLanguages returns a boolean if a field has been set.

### SetDownloadLanguagesNil

`func (o *JellyfinSubtitleOptions) SetDownloadLanguagesNil(b bool)`

 SetDownloadLanguagesNil sets the value for DownloadLanguages to be an explicit nil

### UnsetDownloadLanguages
`func (o *JellyfinSubtitleOptions) UnsetDownloadLanguages()`

UnsetDownloadLanguages ensures that no value is present for DownloadLanguages, not even an explicit nil
### GetDownloadMovieSubtitles

`func (o *JellyfinSubtitleOptions) GetDownloadMovieSubtitles() bool`

GetDownloadMovieSubtitles returns the DownloadMovieSubtitles field if non-nil, zero value otherwise.

### GetDownloadMovieSubtitlesOk

`func (o *JellyfinSubtitleOptions) GetDownloadMovieSubtitlesOk() (*bool, bool)`

GetDownloadMovieSubtitlesOk returns a tuple with the DownloadMovieSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMovieSubtitles

`func (o *JellyfinSubtitleOptions) SetDownloadMovieSubtitles(v bool)`

SetDownloadMovieSubtitles sets DownloadMovieSubtitles field to given value.

### HasDownloadMovieSubtitles

`func (o *JellyfinSubtitleOptions) HasDownloadMovieSubtitles() bool`

HasDownloadMovieSubtitles returns a boolean if a field has been set.

### GetDownloadEpisodeSubtitles

`func (o *JellyfinSubtitleOptions) GetDownloadEpisodeSubtitles() bool`

GetDownloadEpisodeSubtitles returns the DownloadEpisodeSubtitles field if non-nil, zero value otherwise.

### GetDownloadEpisodeSubtitlesOk

`func (o *JellyfinSubtitleOptions) GetDownloadEpisodeSubtitlesOk() (*bool, bool)`

GetDownloadEpisodeSubtitlesOk returns a tuple with the DownloadEpisodeSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEpisodeSubtitles

`func (o *JellyfinSubtitleOptions) SetDownloadEpisodeSubtitles(v bool)`

SetDownloadEpisodeSubtitles sets DownloadEpisodeSubtitles field to given value.

### HasDownloadEpisodeSubtitles

`func (o *JellyfinSubtitleOptions) HasDownloadEpisodeSubtitles() bool`

HasDownloadEpisodeSubtitles returns a boolean if a field has been set.

### GetOpenSubtitlesUsername

`func (o *JellyfinSubtitleOptions) GetOpenSubtitlesUsername() string`

GetOpenSubtitlesUsername returns the OpenSubtitlesUsername field if non-nil, zero value otherwise.

### GetOpenSubtitlesUsernameOk

`func (o *JellyfinSubtitleOptions) GetOpenSubtitlesUsernameOk() (*string, bool)`

GetOpenSubtitlesUsernameOk returns a tuple with the OpenSubtitlesUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSubtitlesUsername

`func (o *JellyfinSubtitleOptions) SetOpenSubtitlesUsername(v string)`

SetOpenSubtitlesUsername sets OpenSubtitlesUsername field to given value.

### HasOpenSubtitlesUsername

`func (o *JellyfinSubtitleOptions) HasOpenSubtitlesUsername() bool`

HasOpenSubtitlesUsername returns a boolean if a field has been set.

### SetOpenSubtitlesUsernameNil

`func (o *JellyfinSubtitleOptions) SetOpenSubtitlesUsernameNil(b bool)`

 SetOpenSubtitlesUsernameNil sets the value for OpenSubtitlesUsername to be an explicit nil

### UnsetOpenSubtitlesUsername
`func (o *JellyfinSubtitleOptions) UnsetOpenSubtitlesUsername()`

UnsetOpenSubtitlesUsername ensures that no value is present for OpenSubtitlesUsername, not even an explicit nil
### GetOpenSubtitlesPasswordHash

`func (o *JellyfinSubtitleOptions) GetOpenSubtitlesPasswordHash() string`

GetOpenSubtitlesPasswordHash returns the OpenSubtitlesPasswordHash field if non-nil, zero value otherwise.

### GetOpenSubtitlesPasswordHashOk

`func (o *JellyfinSubtitleOptions) GetOpenSubtitlesPasswordHashOk() (*string, bool)`

GetOpenSubtitlesPasswordHashOk returns a tuple with the OpenSubtitlesPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSubtitlesPasswordHash

`func (o *JellyfinSubtitleOptions) SetOpenSubtitlesPasswordHash(v string)`

SetOpenSubtitlesPasswordHash sets OpenSubtitlesPasswordHash field to given value.

### HasOpenSubtitlesPasswordHash

`func (o *JellyfinSubtitleOptions) HasOpenSubtitlesPasswordHash() bool`

HasOpenSubtitlesPasswordHash returns a boolean if a field has been set.

### SetOpenSubtitlesPasswordHashNil

`func (o *JellyfinSubtitleOptions) SetOpenSubtitlesPasswordHashNil(b bool)`

 SetOpenSubtitlesPasswordHashNil sets the value for OpenSubtitlesPasswordHash to be an explicit nil

### UnsetOpenSubtitlesPasswordHash
`func (o *JellyfinSubtitleOptions) UnsetOpenSubtitlesPasswordHash()`

UnsetOpenSubtitlesPasswordHash ensures that no value is present for OpenSubtitlesPasswordHash, not even an explicit nil
### GetIsOpenSubtitleVipAccount

`func (o *JellyfinSubtitleOptions) GetIsOpenSubtitleVipAccount() bool`

GetIsOpenSubtitleVipAccount returns the IsOpenSubtitleVipAccount field if non-nil, zero value otherwise.

### GetIsOpenSubtitleVipAccountOk

`func (o *JellyfinSubtitleOptions) GetIsOpenSubtitleVipAccountOk() (*bool, bool)`

GetIsOpenSubtitleVipAccountOk returns a tuple with the IsOpenSubtitleVipAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenSubtitleVipAccount

`func (o *JellyfinSubtitleOptions) SetIsOpenSubtitleVipAccount(v bool)`

SetIsOpenSubtitleVipAccount sets IsOpenSubtitleVipAccount field to given value.

### HasIsOpenSubtitleVipAccount

`func (o *JellyfinSubtitleOptions) HasIsOpenSubtitleVipAccount() bool`

HasIsOpenSubtitleVipAccount returns a boolean if a field has been set.

### GetRequirePerfectMatch

`func (o *JellyfinSubtitleOptions) GetRequirePerfectMatch() bool`

GetRequirePerfectMatch returns the RequirePerfectMatch field if non-nil, zero value otherwise.

### GetRequirePerfectMatchOk

`func (o *JellyfinSubtitleOptions) GetRequirePerfectMatchOk() (*bool, bool)`

GetRequirePerfectMatchOk returns a tuple with the RequirePerfectMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectMatch

`func (o *JellyfinSubtitleOptions) SetRequirePerfectMatch(v bool)`

SetRequirePerfectMatch sets RequirePerfectMatch field to given value.

### HasRequirePerfectMatch

`func (o *JellyfinSubtitleOptions) HasRequirePerfectMatch() bool`

HasRequirePerfectMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


