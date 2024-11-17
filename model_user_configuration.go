/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserConfiguration{}

// UserConfiguration Class UserConfiguration.
type UserConfiguration struct {
	// Gets or sets the audio language preference.
	AudioLanguagePreference NullableString `json:"AudioLanguagePreference,omitempty"`
	// Gets or sets a value indicating whether [play default audio track].
	PlayDefaultAudioTrack *bool `json:"PlayDefaultAudioTrack,omitempty"`
	// Gets or sets the subtitle language preference.
	SubtitleLanguagePreference NullableString `json:"SubtitleLanguagePreference,omitempty"`
	DisplayMissingEpisodes *bool `json:"DisplayMissingEpisodes,omitempty"`
	GroupedFolders []string `json:"GroupedFolders,omitempty"`
	// An enum representing a subtitle playback mode.
	SubtitleMode *SubtitlePlaybackMode `json:"SubtitleMode,omitempty"`
	DisplayCollectionsView *bool `json:"DisplayCollectionsView,omitempty"`
	EnableLocalPassword *bool `json:"EnableLocalPassword,omitempty"`
	OrderedViews []string `json:"OrderedViews,omitempty"`
	LatestItemsExcludes []string `json:"LatestItemsExcludes,omitempty"`
	MyMediaExcludes []string `json:"MyMediaExcludes,omitempty"`
	HidePlayedInLatest *bool `json:"HidePlayedInLatest,omitempty"`
	RememberAudioSelections *bool `json:"RememberAudioSelections,omitempty"`
	RememberSubtitleSelections *bool `json:"RememberSubtitleSelections,omitempty"`
	EnableNextEpisodeAutoPlay *bool `json:"EnableNextEpisodeAutoPlay,omitempty"`
	// Gets or sets the id of the selected cast receiver.
	CastReceiverId NullableString `json:"CastReceiverId,omitempty"`
}

// NewUserConfiguration instantiates a new UserConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConfiguration() *UserConfiguration {
	this := UserConfiguration{}
	return &this
}

// NewUserConfigurationWithDefaults instantiates a new UserConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConfigurationWithDefaults() *UserConfiguration {
	this := UserConfiguration{}
	return &this
}

// GetAudioLanguagePreference returns the AudioLanguagePreference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserConfiguration) GetAudioLanguagePreference() string {
	if o == nil || IsNil(o.AudioLanguagePreference.Get()) {
		var ret string
		return ret
	}
	return *o.AudioLanguagePreference.Get()
}

// GetAudioLanguagePreferenceOk returns a tuple with the AudioLanguagePreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserConfiguration) GetAudioLanguagePreferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioLanguagePreference.Get(), o.AudioLanguagePreference.IsSet()
}

// HasAudioLanguagePreference returns a boolean if a field has been set.
func (o *UserConfiguration) HasAudioLanguagePreference() bool {
	if o != nil && o.AudioLanguagePreference.IsSet() {
		return true
	}

	return false
}

// SetAudioLanguagePreference gets a reference to the given NullableString and assigns it to the AudioLanguagePreference field.
func (o *UserConfiguration) SetAudioLanguagePreference(v string) {
	o.AudioLanguagePreference.Set(&v)
}
// SetAudioLanguagePreferenceNil sets the value for AudioLanguagePreference to be an explicit nil
func (o *UserConfiguration) SetAudioLanguagePreferenceNil() {
	o.AudioLanguagePreference.Set(nil)
}

// UnsetAudioLanguagePreference ensures that no value is present for AudioLanguagePreference, not even an explicit nil
func (o *UserConfiguration) UnsetAudioLanguagePreference() {
	o.AudioLanguagePreference.Unset()
}

// GetPlayDefaultAudioTrack returns the PlayDefaultAudioTrack field value if set, zero value otherwise.
func (o *UserConfiguration) GetPlayDefaultAudioTrack() bool {
	if o == nil || IsNil(o.PlayDefaultAudioTrack) {
		var ret bool
		return ret
	}
	return *o.PlayDefaultAudioTrack
}

// GetPlayDefaultAudioTrackOk returns a tuple with the PlayDefaultAudioTrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetPlayDefaultAudioTrackOk() (*bool, bool) {
	if o == nil || IsNil(o.PlayDefaultAudioTrack) {
		return nil, false
	}
	return o.PlayDefaultAudioTrack, true
}

// HasPlayDefaultAudioTrack returns a boolean if a field has been set.
func (o *UserConfiguration) HasPlayDefaultAudioTrack() bool {
	if o != nil && !IsNil(o.PlayDefaultAudioTrack) {
		return true
	}

	return false
}

// SetPlayDefaultAudioTrack gets a reference to the given bool and assigns it to the PlayDefaultAudioTrack field.
func (o *UserConfiguration) SetPlayDefaultAudioTrack(v bool) {
	o.PlayDefaultAudioTrack = &v
}

// GetSubtitleLanguagePreference returns the SubtitleLanguagePreference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserConfiguration) GetSubtitleLanguagePreference() string {
	if o == nil || IsNil(o.SubtitleLanguagePreference.Get()) {
		var ret string
		return ret
	}
	return *o.SubtitleLanguagePreference.Get()
}

// GetSubtitleLanguagePreferenceOk returns a tuple with the SubtitleLanguagePreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserConfiguration) GetSubtitleLanguagePreferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleLanguagePreference.Get(), o.SubtitleLanguagePreference.IsSet()
}

// HasSubtitleLanguagePreference returns a boolean if a field has been set.
func (o *UserConfiguration) HasSubtitleLanguagePreference() bool {
	if o != nil && o.SubtitleLanguagePreference.IsSet() {
		return true
	}

	return false
}

// SetSubtitleLanguagePreference gets a reference to the given NullableString and assigns it to the SubtitleLanguagePreference field.
func (o *UserConfiguration) SetSubtitleLanguagePreference(v string) {
	o.SubtitleLanguagePreference.Set(&v)
}
// SetSubtitleLanguagePreferenceNil sets the value for SubtitleLanguagePreference to be an explicit nil
func (o *UserConfiguration) SetSubtitleLanguagePreferenceNil() {
	o.SubtitleLanguagePreference.Set(nil)
}

// UnsetSubtitleLanguagePreference ensures that no value is present for SubtitleLanguagePreference, not even an explicit nil
func (o *UserConfiguration) UnsetSubtitleLanguagePreference() {
	o.SubtitleLanguagePreference.Unset()
}

// GetDisplayMissingEpisodes returns the DisplayMissingEpisodes field value if set, zero value otherwise.
func (o *UserConfiguration) GetDisplayMissingEpisodes() bool {
	if o == nil || IsNil(o.DisplayMissingEpisodes) {
		var ret bool
		return ret
	}
	return *o.DisplayMissingEpisodes
}

// GetDisplayMissingEpisodesOk returns a tuple with the DisplayMissingEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetDisplayMissingEpisodesOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayMissingEpisodes) {
		return nil, false
	}
	return o.DisplayMissingEpisodes, true
}

// HasDisplayMissingEpisodes returns a boolean if a field has been set.
func (o *UserConfiguration) HasDisplayMissingEpisodes() bool {
	if o != nil && !IsNil(o.DisplayMissingEpisodes) {
		return true
	}

	return false
}

// SetDisplayMissingEpisodes gets a reference to the given bool and assigns it to the DisplayMissingEpisodes field.
func (o *UserConfiguration) SetDisplayMissingEpisodes(v bool) {
	o.DisplayMissingEpisodes = &v
}

// GetGroupedFolders returns the GroupedFolders field value if set, zero value otherwise.
func (o *UserConfiguration) GetGroupedFolders() []string {
	if o == nil || IsNil(o.GroupedFolders) {
		var ret []string
		return ret
	}
	return o.GroupedFolders
}

// GetGroupedFoldersOk returns a tuple with the GroupedFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetGroupedFoldersOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupedFolders) {
		return nil, false
	}
	return o.GroupedFolders, true
}

// HasGroupedFolders returns a boolean if a field has been set.
func (o *UserConfiguration) HasGroupedFolders() bool {
	if o != nil && !IsNil(o.GroupedFolders) {
		return true
	}

	return false
}

// SetGroupedFolders gets a reference to the given []string and assigns it to the GroupedFolders field.
func (o *UserConfiguration) SetGroupedFolders(v []string) {
	o.GroupedFolders = v
}

// GetSubtitleMode returns the SubtitleMode field value if set, zero value otherwise.
func (o *UserConfiguration) GetSubtitleMode() SubtitlePlaybackMode {
	if o == nil || IsNil(o.SubtitleMode) {
		var ret SubtitlePlaybackMode
		return ret
	}
	return *o.SubtitleMode
}

// GetSubtitleModeOk returns a tuple with the SubtitleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetSubtitleModeOk() (*SubtitlePlaybackMode, bool) {
	if o == nil || IsNil(o.SubtitleMode) {
		return nil, false
	}
	return o.SubtitleMode, true
}

// HasSubtitleMode returns a boolean if a field has been set.
func (o *UserConfiguration) HasSubtitleMode() bool {
	if o != nil && !IsNil(o.SubtitleMode) {
		return true
	}

	return false
}

// SetSubtitleMode gets a reference to the given SubtitlePlaybackMode and assigns it to the SubtitleMode field.
func (o *UserConfiguration) SetSubtitleMode(v SubtitlePlaybackMode) {
	o.SubtitleMode = &v
}

// GetDisplayCollectionsView returns the DisplayCollectionsView field value if set, zero value otherwise.
func (o *UserConfiguration) GetDisplayCollectionsView() bool {
	if o == nil || IsNil(o.DisplayCollectionsView) {
		var ret bool
		return ret
	}
	return *o.DisplayCollectionsView
}

// GetDisplayCollectionsViewOk returns a tuple with the DisplayCollectionsView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetDisplayCollectionsViewOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayCollectionsView) {
		return nil, false
	}
	return o.DisplayCollectionsView, true
}

// HasDisplayCollectionsView returns a boolean if a field has been set.
func (o *UserConfiguration) HasDisplayCollectionsView() bool {
	if o != nil && !IsNil(o.DisplayCollectionsView) {
		return true
	}

	return false
}

// SetDisplayCollectionsView gets a reference to the given bool and assigns it to the DisplayCollectionsView field.
func (o *UserConfiguration) SetDisplayCollectionsView(v bool) {
	o.DisplayCollectionsView = &v
}

// GetEnableLocalPassword returns the EnableLocalPassword field value if set, zero value otherwise.
func (o *UserConfiguration) GetEnableLocalPassword() bool {
	if o == nil || IsNil(o.EnableLocalPassword) {
		var ret bool
		return ret
	}
	return *o.EnableLocalPassword
}

// GetEnableLocalPasswordOk returns a tuple with the EnableLocalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetEnableLocalPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLocalPassword) {
		return nil, false
	}
	return o.EnableLocalPassword, true
}

// HasEnableLocalPassword returns a boolean if a field has been set.
func (o *UserConfiguration) HasEnableLocalPassword() bool {
	if o != nil && !IsNil(o.EnableLocalPassword) {
		return true
	}

	return false
}

// SetEnableLocalPassword gets a reference to the given bool and assigns it to the EnableLocalPassword field.
func (o *UserConfiguration) SetEnableLocalPassword(v bool) {
	o.EnableLocalPassword = &v
}

// GetOrderedViews returns the OrderedViews field value if set, zero value otherwise.
func (o *UserConfiguration) GetOrderedViews() []string {
	if o == nil || IsNil(o.OrderedViews) {
		var ret []string
		return ret
	}
	return o.OrderedViews
}

// GetOrderedViewsOk returns a tuple with the OrderedViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetOrderedViewsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderedViews) {
		return nil, false
	}
	return o.OrderedViews, true
}

// HasOrderedViews returns a boolean if a field has been set.
func (o *UserConfiguration) HasOrderedViews() bool {
	if o != nil && !IsNil(o.OrderedViews) {
		return true
	}

	return false
}

// SetOrderedViews gets a reference to the given []string and assigns it to the OrderedViews field.
func (o *UserConfiguration) SetOrderedViews(v []string) {
	o.OrderedViews = v
}

// GetLatestItemsExcludes returns the LatestItemsExcludes field value if set, zero value otherwise.
func (o *UserConfiguration) GetLatestItemsExcludes() []string {
	if o == nil || IsNil(o.LatestItemsExcludes) {
		var ret []string
		return ret
	}
	return o.LatestItemsExcludes
}

// GetLatestItemsExcludesOk returns a tuple with the LatestItemsExcludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetLatestItemsExcludesOk() ([]string, bool) {
	if o == nil || IsNil(o.LatestItemsExcludes) {
		return nil, false
	}
	return o.LatestItemsExcludes, true
}

// HasLatestItemsExcludes returns a boolean if a field has been set.
func (o *UserConfiguration) HasLatestItemsExcludes() bool {
	if o != nil && !IsNil(o.LatestItemsExcludes) {
		return true
	}

	return false
}

// SetLatestItemsExcludes gets a reference to the given []string and assigns it to the LatestItemsExcludes field.
func (o *UserConfiguration) SetLatestItemsExcludes(v []string) {
	o.LatestItemsExcludes = v
}

// GetMyMediaExcludes returns the MyMediaExcludes field value if set, zero value otherwise.
func (o *UserConfiguration) GetMyMediaExcludes() []string {
	if o == nil || IsNil(o.MyMediaExcludes) {
		var ret []string
		return ret
	}
	return o.MyMediaExcludes
}

// GetMyMediaExcludesOk returns a tuple with the MyMediaExcludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetMyMediaExcludesOk() ([]string, bool) {
	if o == nil || IsNil(o.MyMediaExcludes) {
		return nil, false
	}
	return o.MyMediaExcludes, true
}

// HasMyMediaExcludes returns a boolean if a field has been set.
func (o *UserConfiguration) HasMyMediaExcludes() bool {
	if o != nil && !IsNil(o.MyMediaExcludes) {
		return true
	}

	return false
}

// SetMyMediaExcludes gets a reference to the given []string and assigns it to the MyMediaExcludes field.
func (o *UserConfiguration) SetMyMediaExcludes(v []string) {
	o.MyMediaExcludes = v
}

// GetHidePlayedInLatest returns the HidePlayedInLatest field value if set, zero value otherwise.
func (o *UserConfiguration) GetHidePlayedInLatest() bool {
	if o == nil || IsNil(o.HidePlayedInLatest) {
		var ret bool
		return ret
	}
	return *o.HidePlayedInLatest
}

// GetHidePlayedInLatestOk returns a tuple with the HidePlayedInLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetHidePlayedInLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.HidePlayedInLatest) {
		return nil, false
	}
	return o.HidePlayedInLatest, true
}

// HasHidePlayedInLatest returns a boolean if a field has been set.
func (o *UserConfiguration) HasHidePlayedInLatest() bool {
	if o != nil && !IsNil(o.HidePlayedInLatest) {
		return true
	}

	return false
}

// SetHidePlayedInLatest gets a reference to the given bool and assigns it to the HidePlayedInLatest field.
func (o *UserConfiguration) SetHidePlayedInLatest(v bool) {
	o.HidePlayedInLatest = &v
}

// GetRememberAudioSelections returns the RememberAudioSelections field value if set, zero value otherwise.
func (o *UserConfiguration) GetRememberAudioSelections() bool {
	if o == nil || IsNil(o.RememberAudioSelections) {
		var ret bool
		return ret
	}
	return *o.RememberAudioSelections
}

// GetRememberAudioSelectionsOk returns a tuple with the RememberAudioSelections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetRememberAudioSelectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberAudioSelections) {
		return nil, false
	}
	return o.RememberAudioSelections, true
}

// HasRememberAudioSelections returns a boolean if a field has been set.
func (o *UserConfiguration) HasRememberAudioSelections() bool {
	if o != nil && !IsNil(o.RememberAudioSelections) {
		return true
	}

	return false
}

// SetRememberAudioSelections gets a reference to the given bool and assigns it to the RememberAudioSelections field.
func (o *UserConfiguration) SetRememberAudioSelections(v bool) {
	o.RememberAudioSelections = &v
}

// GetRememberSubtitleSelections returns the RememberSubtitleSelections field value if set, zero value otherwise.
func (o *UserConfiguration) GetRememberSubtitleSelections() bool {
	if o == nil || IsNil(o.RememberSubtitleSelections) {
		var ret bool
		return ret
	}
	return *o.RememberSubtitleSelections
}

// GetRememberSubtitleSelectionsOk returns a tuple with the RememberSubtitleSelections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetRememberSubtitleSelectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberSubtitleSelections) {
		return nil, false
	}
	return o.RememberSubtitleSelections, true
}

// HasRememberSubtitleSelections returns a boolean if a field has been set.
func (o *UserConfiguration) HasRememberSubtitleSelections() bool {
	if o != nil && !IsNil(o.RememberSubtitleSelections) {
		return true
	}

	return false
}

// SetRememberSubtitleSelections gets a reference to the given bool and assigns it to the RememberSubtitleSelections field.
func (o *UserConfiguration) SetRememberSubtitleSelections(v bool) {
	o.RememberSubtitleSelections = &v
}

// GetEnableNextEpisodeAutoPlay returns the EnableNextEpisodeAutoPlay field value if set, zero value otherwise.
func (o *UserConfiguration) GetEnableNextEpisodeAutoPlay() bool {
	if o == nil || IsNil(o.EnableNextEpisodeAutoPlay) {
		var ret bool
		return ret
	}
	return *o.EnableNextEpisodeAutoPlay
}

// GetEnableNextEpisodeAutoPlayOk returns a tuple with the EnableNextEpisodeAutoPlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConfiguration) GetEnableNextEpisodeAutoPlayOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableNextEpisodeAutoPlay) {
		return nil, false
	}
	return o.EnableNextEpisodeAutoPlay, true
}

// HasEnableNextEpisodeAutoPlay returns a boolean if a field has been set.
func (o *UserConfiguration) HasEnableNextEpisodeAutoPlay() bool {
	if o != nil && !IsNil(o.EnableNextEpisodeAutoPlay) {
		return true
	}

	return false
}

// SetEnableNextEpisodeAutoPlay gets a reference to the given bool and assigns it to the EnableNextEpisodeAutoPlay field.
func (o *UserConfiguration) SetEnableNextEpisodeAutoPlay(v bool) {
	o.EnableNextEpisodeAutoPlay = &v
}

// GetCastReceiverId returns the CastReceiverId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserConfiguration) GetCastReceiverId() string {
	if o == nil || IsNil(o.CastReceiverId.Get()) {
		var ret string
		return ret
	}
	return *o.CastReceiverId.Get()
}

// GetCastReceiverIdOk returns a tuple with the CastReceiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserConfiguration) GetCastReceiverIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CastReceiverId.Get(), o.CastReceiverId.IsSet()
}

// HasCastReceiverId returns a boolean if a field has been set.
func (o *UserConfiguration) HasCastReceiverId() bool {
	if o != nil && o.CastReceiverId.IsSet() {
		return true
	}

	return false
}

// SetCastReceiverId gets a reference to the given NullableString and assigns it to the CastReceiverId field.
func (o *UserConfiguration) SetCastReceiverId(v string) {
	o.CastReceiverId.Set(&v)
}
// SetCastReceiverIdNil sets the value for CastReceiverId to be an explicit nil
func (o *UserConfiguration) SetCastReceiverIdNil() {
	o.CastReceiverId.Set(nil)
}

// UnsetCastReceiverId ensures that no value is present for CastReceiverId, not even an explicit nil
func (o *UserConfiguration) UnsetCastReceiverId() {
	o.CastReceiverId.Unset()
}

func (o UserConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioLanguagePreference.IsSet() {
		toSerialize["AudioLanguagePreference"] = o.AudioLanguagePreference.Get()
	}
	if !IsNil(o.PlayDefaultAudioTrack) {
		toSerialize["PlayDefaultAudioTrack"] = o.PlayDefaultAudioTrack
	}
	if o.SubtitleLanguagePreference.IsSet() {
		toSerialize["SubtitleLanguagePreference"] = o.SubtitleLanguagePreference.Get()
	}
	if !IsNil(o.DisplayMissingEpisodes) {
		toSerialize["DisplayMissingEpisodes"] = o.DisplayMissingEpisodes
	}
	if !IsNil(o.GroupedFolders) {
		toSerialize["GroupedFolders"] = o.GroupedFolders
	}
	if !IsNil(o.SubtitleMode) {
		toSerialize["SubtitleMode"] = o.SubtitleMode
	}
	if !IsNil(o.DisplayCollectionsView) {
		toSerialize["DisplayCollectionsView"] = o.DisplayCollectionsView
	}
	if !IsNil(o.EnableLocalPassword) {
		toSerialize["EnableLocalPassword"] = o.EnableLocalPassword
	}
	if !IsNil(o.OrderedViews) {
		toSerialize["OrderedViews"] = o.OrderedViews
	}
	if !IsNil(o.LatestItemsExcludes) {
		toSerialize["LatestItemsExcludes"] = o.LatestItemsExcludes
	}
	if !IsNil(o.MyMediaExcludes) {
		toSerialize["MyMediaExcludes"] = o.MyMediaExcludes
	}
	if !IsNil(o.HidePlayedInLatest) {
		toSerialize["HidePlayedInLatest"] = o.HidePlayedInLatest
	}
	if !IsNil(o.RememberAudioSelections) {
		toSerialize["RememberAudioSelections"] = o.RememberAudioSelections
	}
	if !IsNil(o.RememberSubtitleSelections) {
		toSerialize["RememberSubtitleSelections"] = o.RememberSubtitleSelections
	}
	if !IsNil(o.EnableNextEpisodeAutoPlay) {
		toSerialize["EnableNextEpisodeAutoPlay"] = o.EnableNextEpisodeAutoPlay
	}
	if o.CastReceiverId.IsSet() {
		toSerialize["CastReceiverId"] = o.CastReceiverId.Get()
	}
	return toSerialize, nil
}

type NullableUserConfiguration struct {
	value *UserConfiguration
	isSet bool
}

func (v NullableUserConfiguration) Get() *UserConfiguration {
	return v.value
}

func (v *NullableUserConfiguration) Set(val *UserConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConfiguration(val *UserConfiguration) *NullableUserConfiguration {
	return &NullableUserConfiguration{value: val, isSet: true}
}

func (v NullableUserConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


