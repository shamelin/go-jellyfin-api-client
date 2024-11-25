/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
)

// checks if the JellyfinSubtitleOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinSubtitleOptions{}

// JellyfinSubtitleOptions struct for JellyfinSubtitleOptions
type JellyfinSubtitleOptions struct {
	SkipIfEmbeddedSubtitlesPresent *bool `json:"SkipIfEmbeddedSubtitlesPresent,omitempty"`
	SkipIfAudioTrackMatches *bool `json:"SkipIfAudioTrackMatches,omitempty"`
	DownloadLanguages []string `json:"DownloadLanguages,omitempty"`
	DownloadMovieSubtitles *bool `json:"DownloadMovieSubtitles,omitempty"`
	DownloadEpisodeSubtitles *bool `json:"DownloadEpisodeSubtitles,omitempty"`
	OpenSubtitlesUsername NullableString `json:"OpenSubtitlesUsername,omitempty"`
	OpenSubtitlesPasswordHash NullableString `json:"OpenSubtitlesPasswordHash,omitempty"`
	IsOpenSubtitleVipAccount *bool `json:"IsOpenSubtitleVipAccount,omitempty"`
	RequirePerfectMatch *bool `json:"RequirePerfectMatch,omitempty"`
}

// NewJellyfinSubtitleOptions instantiates a new JellyfinSubtitleOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinSubtitleOptions() *JellyfinSubtitleOptions {
	this := JellyfinSubtitleOptions{}
	return &this
}

// NewJellyfinSubtitleOptionsWithDefaults instantiates a new JellyfinSubtitleOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinSubtitleOptionsWithDefaults() *JellyfinSubtitleOptions {
	this := JellyfinSubtitleOptions{}
	return &this
}

// GetSkipIfEmbeddedSubtitlesPresent returns the SkipIfEmbeddedSubtitlesPresent field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetSkipIfEmbeddedSubtitlesPresent() bool {
	if o == nil || IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		var ret bool
		return ret
	}
	return *o.SkipIfEmbeddedSubtitlesPresent
}

// GetSkipIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipIfEmbeddedSubtitlesPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetSkipIfEmbeddedSubtitlesPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		return nil, false
	}
	return o.SkipIfEmbeddedSubtitlesPresent, true
}

// HasSkipIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasSkipIfEmbeddedSubtitlesPresent() bool {
	if o != nil && !IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		return true
	}

	return false
}

// SetSkipIfEmbeddedSubtitlesPresent gets a reference to the given bool and assigns it to the SkipIfEmbeddedSubtitlesPresent field.
func (o *JellyfinSubtitleOptions) SetSkipIfEmbeddedSubtitlesPresent(v bool) {
	o.SkipIfEmbeddedSubtitlesPresent = &v
}

// GetSkipIfAudioTrackMatches returns the SkipIfAudioTrackMatches field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetSkipIfAudioTrackMatches() bool {
	if o == nil || IsNil(o.SkipIfAudioTrackMatches) {
		var ret bool
		return ret
	}
	return *o.SkipIfAudioTrackMatches
}

// GetSkipIfAudioTrackMatchesOk returns a tuple with the SkipIfAudioTrackMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetSkipIfAudioTrackMatchesOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipIfAudioTrackMatches) {
		return nil, false
	}
	return o.SkipIfAudioTrackMatches, true
}

// HasSkipIfAudioTrackMatches returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasSkipIfAudioTrackMatches() bool {
	if o != nil && !IsNil(o.SkipIfAudioTrackMatches) {
		return true
	}

	return false
}

// SetSkipIfAudioTrackMatches gets a reference to the given bool and assigns it to the SkipIfAudioTrackMatches field.
func (o *JellyfinSubtitleOptions) SetSkipIfAudioTrackMatches(v bool) {
	o.SkipIfAudioTrackMatches = &v
}

// GetDownloadLanguages returns the DownloadLanguages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinSubtitleOptions) GetDownloadLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DownloadLanguages
}

// GetDownloadLanguagesOk returns a tuple with the DownloadLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinSubtitleOptions) GetDownloadLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.DownloadLanguages) {
		return nil, false
	}
	return o.DownloadLanguages, true
}

// HasDownloadLanguages returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasDownloadLanguages() bool {
	if o != nil && !IsNil(o.DownloadLanguages) {
		return true
	}

	return false
}

// SetDownloadLanguages gets a reference to the given []string and assigns it to the DownloadLanguages field.
func (o *JellyfinSubtitleOptions) SetDownloadLanguages(v []string) {
	o.DownloadLanguages = v
}

// GetDownloadMovieSubtitles returns the DownloadMovieSubtitles field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetDownloadMovieSubtitles() bool {
	if o == nil || IsNil(o.DownloadMovieSubtitles) {
		var ret bool
		return ret
	}
	return *o.DownloadMovieSubtitles
}

// GetDownloadMovieSubtitlesOk returns a tuple with the DownloadMovieSubtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetDownloadMovieSubtitlesOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadMovieSubtitles) {
		return nil, false
	}
	return o.DownloadMovieSubtitles, true
}

// HasDownloadMovieSubtitles returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasDownloadMovieSubtitles() bool {
	if o != nil && !IsNil(o.DownloadMovieSubtitles) {
		return true
	}

	return false
}

// SetDownloadMovieSubtitles gets a reference to the given bool and assigns it to the DownloadMovieSubtitles field.
func (o *JellyfinSubtitleOptions) SetDownloadMovieSubtitles(v bool) {
	o.DownloadMovieSubtitles = &v
}

// GetDownloadEpisodeSubtitles returns the DownloadEpisodeSubtitles field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetDownloadEpisodeSubtitles() bool {
	if o == nil || IsNil(o.DownloadEpisodeSubtitles) {
		var ret bool
		return ret
	}
	return *o.DownloadEpisodeSubtitles
}

// GetDownloadEpisodeSubtitlesOk returns a tuple with the DownloadEpisodeSubtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetDownloadEpisodeSubtitlesOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadEpisodeSubtitles) {
		return nil, false
	}
	return o.DownloadEpisodeSubtitles, true
}

// HasDownloadEpisodeSubtitles returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasDownloadEpisodeSubtitles() bool {
	if o != nil && !IsNil(o.DownloadEpisodeSubtitles) {
		return true
	}

	return false
}

// SetDownloadEpisodeSubtitles gets a reference to the given bool and assigns it to the DownloadEpisodeSubtitles field.
func (o *JellyfinSubtitleOptions) SetDownloadEpisodeSubtitles(v bool) {
	o.DownloadEpisodeSubtitles = &v
}

// GetOpenSubtitlesUsername returns the OpenSubtitlesUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinSubtitleOptions) GetOpenSubtitlesUsername() string {
	if o == nil || IsNil(o.OpenSubtitlesUsername.Get()) {
		var ret string
		return ret
	}
	return *o.OpenSubtitlesUsername.Get()
}

// GetOpenSubtitlesUsernameOk returns a tuple with the OpenSubtitlesUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinSubtitleOptions) GetOpenSubtitlesUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenSubtitlesUsername.Get(), o.OpenSubtitlesUsername.IsSet()
}

// HasOpenSubtitlesUsername returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasOpenSubtitlesUsername() bool {
	if o != nil && o.OpenSubtitlesUsername.IsSet() {
		return true
	}

	return false
}

// SetOpenSubtitlesUsername gets a reference to the given NullableString and assigns it to the OpenSubtitlesUsername field.
func (o *JellyfinSubtitleOptions) SetOpenSubtitlesUsername(v string) {
	o.OpenSubtitlesUsername.Set(&v)
}
// SetOpenSubtitlesUsernameNil sets the value for OpenSubtitlesUsername to be an explicit nil
func (o *JellyfinSubtitleOptions) SetOpenSubtitlesUsernameNil() {
	o.OpenSubtitlesUsername.Set(nil)
}

// UnsetOpenSubtitlesUsername ensures that no value is present for OpenSubtitlesUsername, not even an explicit nil
func (o *JellyfinSubtitleOptions) UnsetOpenSubtitlesUsername() {
	o.OpenSubtitlesUsername.Unset()
}

// GetOpenSubtitlesPasswordHash returns the OpenSubtitlesPasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinSubtitleOptions) GetOpenSubtitlesPasswordHash() string {
	if o == nil || IsNil(o.OpenSubtitlesPasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.OpenSubtitlesPasswordHash.Get()
}

// GetOpenSubtitlesPasswordHashOk returns a tuple with the OpenSubtitlesPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinSubtitleOptions) GetOpenSubtitlesPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenSubtitlesPasswordHash.Get(), o.OpenSubtitlesPasswordHash.IsSet()
}

// HasOpenSubtitlesPasswordHash returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasOpenSubtitlesPasswordHash() bool {
	if o != nil && o.OpenSubtitlesPasswordHash.IsSet() {
		return true
	}

	return false
}

// SetOpenSubtitlesPasswordHash gets a reference to the given NullableString and assigns it to the OpenSubtitlesPasswordHash field.
func (o *JellyfinSubtitleOptions) SetOpenSubtitlesPasswordHash(v string) {
	o.OpenSubtitlesPasswordHash.Set(&v)
}
// SetOpenSubtitlesPasswordHashNil sets the value for OpenSubtitlesPasswordHash to be an explicit nil
func (o *JellyfinSubtitleOptions) SetOpenSubtitlesPasswordHashNil() {
	o.OpenSubtitlesPasswordHash.Set(nil)
}

// UnsetOpenSubtitlesPasswordHash ensures that no value is present for OpenSubtitlesPasswordHash, not even an explicit nil
func (o *JellyfinSubtitleOptions) UnsetOpenSubtitlesPasswordHash() {
	o.OpenSubtitlesPasswordHash.Unset()
}

// GetIsOpenSubtitleVipAccount returns the IsOpenSubtitleVipAccount field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetIsOpenSubtitleVipAccount() bool {
	if o == nil || IsNil(o.IsOpenSubtitleVipAccount) {
		var ret bool
		return ret
	}
	return *o.IsOpenSubtitleVipAccount
}

// GetIsOpenSubtitleVipAccountOk returns a tuple with the IsOpenSubtitleVipAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetIsOpenSubtitleVipAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpenSubtitleVipAccount) {
		return nil, false
	}
	return o.IsOpenSubtitleVipAccount, true
}

// HasIsOpenSubtitleVipAccount returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasIsOpenSubtitleVipAccount() bool {
	if o != nil && !IsNil(o.IsOpenSubtitleVipAccount) {
		return true
	}

	return false
}

// SetIsOpenSubtitleVipAccount gets a reference to the given bool and assigns it to the IsOpenSubtitleVipAccount field.
func (o *JellyfinSubtitleOptions) SetIsOpenSubtitleVipAccount(v bool) {
	o.IsOpenSubtitleVipAccount = &v
}

// GetRequirePerfectMatch returns the RequirePerfectMatch field value if set, zero value otherwise.
func (o *JellyfinSubtitleOptions) GetRequirePerfectMatch() bool {
	if o == nil || IsNil(o.RequirePerfectMatch) {
		var ret bool
		return ret
	}
	return *o.RequirePerfectMatch
}

// GetRequirePerfectMatchOk returns a tuple with the RequirePerfectMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSubtitleOptions) GetRequirePerfectMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePerfectMatch) {
		return nil, false
	}
	return o.RequirePerfectMatch, true
}

// HasRequirePerfectMatch returns a boolean if a field has been set.
func (o *JellyfinSubtitleOptions) HasRequirePerfectMatch() bool {
	if o != nil && !IsNil(o.RequirePerfectMatch) {
		return true
	}

	return false
}

// SetRequirePerfectMatch gets a reference to the given bool and assigns it to the RequirePerfectMatch field.
func (o *JellyfinSubtitleOptions) SetRequirePerfectMatch(v bool) {
	o.RequirePerfectMatch = &v
}

func (o JellyfinSubtitleOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinSubtitleOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		toSerialize["SkipIfEmbeddedSubtitlesPresent"] = o.SkipIfEmbeddedSubtitlesPresent
	}
	if !IsNil(o.SkipIfAudioTrackMatches) {
		toSerialize["SkipIfAudioTrackMatches"] = o.SkipIfAudioTrackMatches
	}
	if o.DownloadLanguages != nil {
		toSerialize["DownloadLanguages"] = o.DownloadLanguages
	}
	if !IsNil(o.DownloadMovieSubtitles) {
		toSerialize["DownloadMovieSubtitles"] = o.DownloadMovieSubtitles
	}
	if !IsNil(o.DownloadEpisodeSubtitles) {
		toSerialize["DownloadEpisodeSubtitles"] = o.DownloadEpisodeSubtitles
	}
	if o.OpenSubtitlesUsername.IsSet() {
		toSerialize["OpenSubtitlesUsername"] = o.OpenSubtitlesUsername.Get()
	}
	if o.OpenSubtitlesPasswordHash.IsSet() {
		toSerialize["OpenSubtitlesPasswordHash"] = o.OpenSubtitlesPasswordHash.Get()
	}
	if !IsNil(o.IsOpenSubtitleVipAccount) {
		toSerialize["IsOpenSubtitleVipAccount"] = o.IsOpenSubtitleVipAccount
	}
	if !IsNil(o.RequirePerfectMatch) {
		toSerialize["RequirePerfectMatch"] = o.RequirePerfectMatch
	}
	return toSerialize, nil
}

type NullableJellyfinSubtitleOptions struct {
	value *JellyfinSubtitleOptions
	isSet bool
}

func (v NullableJellyfinSubtitleOptions) Get() *JellyfinSubtitleOptions {
	return v.value
}

func (v *NullableJellyfinSubtitleOptions) Set(val *JellyfinSubtitleOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSubtitleOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSubtitleOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSubtitleOptions(val *JellyfinSubtitleOptions) *NullableJellyfinSubtitleOptions {
	return &NullableJellyfinSubtitleOptions{value: val, isSet: true}
}

func (v NullableJellyfinSubtitleOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSubtitleOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


