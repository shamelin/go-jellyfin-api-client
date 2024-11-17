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

// checks if the RemoteImageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteImageInfo{}

// RemoteImageInfo Class RemoteImageInfo.
type RemoteImageInfo struct {
	// Gets or sets the name of the provider.
	ProviderName NullableString `json:"ProviderName,omitempty"`
	// Gets or sets the URL.
	Url NullableString `json:"Url,omitempty"`
	// Gets or sets a url used for previewing a smaller version.
	ThumbnailUrl NullableString `json:"ThumbnailUrl,omitempty"`
	// Gets or sets the height.
	Height NullableInt32 `json:"Height,omitempty"`
	// Gets or sets the width.
	Width NullableInt32 `json:"Width,omitempty"`
	// Gets or sets the community rating.
	CommunityRating NullableFloat64 `json:"CommunityRating,omitempty"`
	// Gets or sets the vote count.
	VoteCount NullableInt32 `json:"VoteCount,omitempty"`
	// Gets or sets the language.
	Language NullableString `json:"Language,omitempty"`
	// Gets or sets the type.
	Type *ImageType `json:"Type,omitempty"`
	// Gets or sets the type of the rating.
	RatingType *RatingType `json:"RatingType,omitempty"`
}

// NewRemoteImageInfo instantiates a new RemoteImageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteImageInfo() *RemoteImageInfo {
	this := RemoteImageInfo{}
	return &this
}

// NewRemoteImageInfoWithDefaults instantiates a new RemoteImageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteImageInfoWithDefaults() *RemoteImageInfo {
	this := RemoteImageInfo{}
	return &this
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderName.Get()
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderName.Get(), o.ProviderName.IsSet()
}

// HasProviderName returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasProviderName() bool {
	if o != nil && o.ProviderName.IsSet() {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given NullableString and assigns it to the ProviderName field.
func (o *RemoteImageInfo) SetProviderName(v string) {
	o.ProviderName.Set(&v)
}
// SetProviderNameNil sets the value for ProviderName to be an explicit nil
func (o *RemoteImageInfo) SetProviderNameNil() {
	o.ProviderName.Set(nil)
}

// UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
func (o *RemoteImageInfo) UnsetProviderName() {
	o.ProviderName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RemoteImageInfo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *RemoteImageInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RemoteImageInfo) UnsetUrl() {
	o.Url.Unset()
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *RemoteImageInfo) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *RemoteImageInfo) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *RemoteImageInfo) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *RemoteImageInfo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *RemoteImageInfo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *RemoteImageInfo) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *RemoteImageInfo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *RemoteImageInfo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *RemoteImageInfo) UnsetWidth() {
	o.Width.Unset()
}

// GetCommunityRating returns the CommunityRating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetCommunityRating() float64 {
	if o == nil || IsNil(o.CommunityRating.Get()) {
		var ret float64
		return ret
	}
	return *o.CommunityRating.Get()
}

// GetCommunityRatingOk returns a tuple with the CommunityRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetCommunityRatingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommunityRating.Get(), o.CommunityRating.IsSet()
}

// HasCommunityRating returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasCommunityRating() bool {
	if o != nil && o.CommunityRating.IsSet() {
		return true
	}

	return false
}

// SetCommunityRating gets a reference to the given NullableFloat64 and assigns it to the CommunityRating field.
func (o *RemoteImageInfo) SetCommunityRating(v float64) {
	o.CommunityRating.Set(&v)
}
// SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil
func (o *RemoteImageInfo) SetCommunityRatingNil() {
	o.CommunityRating.Set(nil)
}

// UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
func (o *RemoteImageInfo) UnsetCommunityRating() {
	o.CommunityRating.Unset()
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetVoteCount() int32 {
	if o == nil || IsNil(o.VoteCount.Get()) {
		var ret int32
		return ret
	}
	return *o.VoteCount.Get()
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetVoteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoteCount.Get(), o.VoteCount.IsSet()
}

// HasVoteCount returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasVoteCount() bool {
	if o != nil && o.VoteCount.IsSet() {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given NullableInt32 and assigns it to the VoteCount field.
func (o *RemoteImageInfo) SetVoteCount(v int32) {
	o.VoteCount.Set(&v)
}
// SetVoteCountNil sets the value for VoteCount to be an explicit nil
func (o *RemoteImageInfo) SetVoteCountNil() {
	o.VoteCount.Set(nil)
}

// UnsetVoteCount ensures that no value is present for VoteCount, not even an explicit nil
func (o *RemoteImageInfo) UnsetVoteCount() {
	o.VoteCount.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageInfo) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageInfo) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *RemoteImageInfo) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *RemoteImageInfo) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *RemoteImageInfo) UnsetLanguage() {
	o.Language.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RemoteImageInfo) GetType() ImageType {
	if o == nil || IsNil(o.Type) {
		var ret ImageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteImageInfo) GetTypeOk() (*ImageType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ImageType and assigns it to the Type field.
func (o *RemoteImageInfo) SetType(v ImageType) {
	o.Type = &v
}

// GetRatingType returns the RatingType field value if set, zero value otherwise.
func (o *RemoteImageInfo) GetRatingType() RatingType {
	if o == nil || IsNil(o.RatingType) {
		var ret RatingType
		return ret
	}
	return *o.RatingType
}

// GetRatingTypeOk returns a tuple with the RatingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteImageInfo) GetRatingTypeOk() (*RatingType, bool) {
	if o == nil || IsNil(o.RatingType) {
		return nil, false
	}
	return o.RatingType, true
}

// HasRatingType returns a boolean if a field has been set.
func (o *RemoteImageInfo) HasRatingType() bool {
	if o != nil && !IsNil(o.RatingType) {
		return true
	}

	return false
}

// SetRatingType gets a reference to the given RatingType and assigns it to the RatingType field.
func (o *RemoteImageInfo) SetRatingType(v RatingType) {
	o.RatingType = &v
}

func (o RemoteImageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteImageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderName.IsSet() {
		toSerialize["ProviderName"] = o.ProviderName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["Url"] = o.Url.Get()
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["ThumbnailUrl"] = o.ThumbnailUrl.Get()
	}
	if o.Height.IsSet() {
		toSerialize["Height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["Width"] = o.Width.Get()
	}
	if o.CommunityRating.IsSet() {
		toSerialize["CommunityRating"] = o.CommunityRating.Get()
	}
	if o.VoteCount.IsSet() {
		toSerialize["VoteCount"] = o.VoteCount.Get()
	}
	if o.Language.IsSet() {
		toSerialize["Language"] = o.Language.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.RatingType) {
		toSerialize["RatingType"] = o.RatingType
	}
	return toSerialize, nil
}

type NullableRemoteImageInfo struct {
	value *RemoteImageInfo
	isSet bool
}

func (v NullableRemoteImageInfo) Get() *RemoteImageInfo {
	return v.value
}

func (v *NullableRemoteImageInfo) Set(val *RemoteImageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteImageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteImageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteImageInfo(val *RemoteImageInfo) *NullableRemoteImageInfo {
	return &NullableRemoteImageInfo{value: val, isSet: true}
}

func (v NullableRemoteImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteImageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


