/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"time"
)

// checks if the JellyfinUserItemDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinUserItemDataDto{}

// JellyfinUserItemDataDto Class UserItemDataDto.
type JellyfinUserItemDataDto struct {
	// Gets or sets the rating.
	Rating NullableFloat64 `json:"Rating,omitempty"`
	// Gets or sets the played percentage.
	PlayedPercentage NullableFloat64 `json:"PlayedPercentage,omitempty"`
	// Gets or sets the unplayed item count.
	UnplayedItemCount NullableInt32 `json:"UnplayedItemCount,omitempty"`
	// Gets or sets the playback position ticks.
	PlaybackPositionTicks *int64 `json:"PlaybackPositionTicks,omitempty"`
	// Gets or sets the play count.
	PlayCount *int32 `json:"PlayCount,omitempty"`
	// Gets or sets a value indicating whether this instance is favorite.
	IsFavorite *bool `json:"IsFavorite,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is likes.
	Likes NullableBool `json:"Likes,omitempty"`
	// Gets or sets the last played date.
	LastPlayedDate NullableTime `json:"LastPlayedDate,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played.
	Played *bool `json:"Played,omitempty"`
	// Gets or sets the key.
	Key *string `json:"Key,omitempty"`
	// Gets or sets the item identifier.
	ItemId *string `json:"ItemId,omitempty"`
}

// NewJellyfinUserItemDataDto instantiates a new JellyfinUserItemDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinUserItemDataDto() *JellyfinUserItemDataDto {
	this := JellyfinUserItemDataDto{}
	return &this
}

// NewJellyfinUserItemDataDtoWithDefaults instantiates a new JellyfinUserItemDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinUserItemDataDtoWithDefaults() *JellyfinUserItemDataDto {
	this := JellyfinUserItemDataDto{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserItemDataDto) GetRating() float64 {
	if o == nil || IsNil(o.Rating.Get()) {
		var ret float64
		return ret
	}
	return *o.Rating.Get()
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserItemDataDto) GetRatingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating.Get(), o.Rating.IsSet()
}

// HasRating returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasRating() bool {
	if o != nil && o.Rating.IsSet() {
		return true
	}

	return false
}

// SetRating gets a reference to the given NullableFloat64 and assigns it to the Rating field.
func (o *JellyfinUserItemDataDto) SetRating(v float64) {
	o.Rating.Set(&v)
}
// SetRatingNil sets the value for Rating to be an explicit nil
func (o *JellyfinUserItemDataDto) SetRatingNil() {
	o.Rating.Set(nil)
}

// UnsetRating ensures that no value is present for Rating, not even an explicit nil
func (o *JellyfinUserItemDataDto) UnsetRating() {
	o.Rating.Unset()
}

// GetPlayedPercentage returns the PlayedPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserItemDataDto) GetPlayedPercentage() float64 {
	if o == nil || IsNil(o.PlayedPercentage.Get()) {
		var ret float64
		return ret
	}
	return *o.PlayedPercentage.Get()
}

// GetPlayedPercentageOk returns a tuple with the PlayedPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserItemDataDto) GetPlayedPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayedPercentage.Get(), o.PlayedPercentage.IsSet()
}

// HasPlayedPercentage returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasPlayedPercentage() bool {
	if o != nil && o.PlayedPercentage.IsSet() {
		return true
	}

	return false
}

// SetPlayedPercentage gets a reference to the given NullableFloat64 and assigns it to the PlayedPercentage field.
func (o *JellyfinUserItemDataDto) SetPlayedPercentage(v float64) {
	o.PlayedPercentage.Set(&v)
}
// SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil
func (o *JellyfinUserItemDataDto) SetPlayedPercentageNil() {
	o.PlayedPercentage.Set(nil)
}

// UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
func (o *JellyfinUserItemDataDto) UnsetPlayedPercentage() {
	o.PlayedPercentage.Unset()
}

// GetUnplayedItemCount returns the UnplayedItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserItemDataDto) GetUnplayedItemCount() int32 {
	if o == nil || IsNil(o.UnplayedItemCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnplayedItemCount.Get()
}

// GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserItemDataDto) GetUnplayedItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnplayedItemCount.Get(), o.UnplayedItemCount.IsSet()
}

// HasUnplayedItemCount returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasUnplayedItemCount() bool {
	if o != nil && o.UnplayedItemCount.IsSet() {
		return true
	}

	return false
}

// SetUnplayedItemCount gets a reference to the given NullableInt32 and assigns it to the UnplayedItemCount field.
func (o *JellyfinUserItemDataDto) SetUnplayedItemCount(v int32) {
	o.UnplayedItemCount.Set(&v)
}
// SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil
func (o *JellyfinUserItemDataDto) SetUnplayedItemCountNil() {
	o.UnplayedItemCount.Set(nil)
}

// UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
func (o *JellyfinUserItemDataDto) UnsetUnplayedItemCount() {
	o.UnplayedItemCount.Unset()
}

// GetPlaybackPositionTicks returns the PlaybackPositionTicks field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetPlaybackPositionTicks() int64 {
	if o == nil || IsNil(o.PlaybackPositionTicks) {
		var ret int64
		return ret
	}
	return *o.PlaybackPositionTicks
}

// GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.PlaybackPositionTicks) {
		return nil, false
	}
	return o.PlaybackPositionTicks, true
}

// HasPlaybackPositionTicks returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasPlaybackPositionTicks() bool {
	if o != nil && !IsNil(o.PlaybackPositionTicks) {
		return true
	}

	return false
}

// SetPlaybackPositionTicks gets a reference to the given int64 and assigns it to the PlaybackPositionTicks field.
func (o *JellyfinUserItemDataDto) SetPlaybackPositionTicks(v int64) {
	o.PlaybackPositionTicks = &v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetPlayCount() int32 {
	if o == nil || IsNil(o.PlayCount) {
		var ret int32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetPlayCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayCount) {
		return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasPlayCount() bool {
	if o != nil && !IsNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given int32 and assigns it to the PlayCount field.
func (o *JellyfinUserItemDataDto) SetPlayCount(v int32) {
	o.PlayCount = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasIsFavorite() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *JellyfinUserItemDataDto) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserItemDataDto) GetLikes() bool {
	if o == nil || IsNil(o.Likes.Get()) {
		var ret bool
		return ret
	}
	return *o.Likes.Get()
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserItemDataDto) GetLikesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Likes.Get(), o.Likes.IsSet()
}

// HasLikes returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasLikes() bool {
	if o != nil && o.Likes.IsSet() {
		return true
	}

	return false
}

// SetLikes gets a reference to the given NullableBool and assigns it to the Likes field.
func (o *JellyfinUserItemDataDto) SetLikes(v bool) {
	o.Likes.Set(&v)
}
// SetLikesNil sets the value for Likes to be an explicit nil
func (o *JellyfinUserItemDataDto) SetLikesNil() {
	o.Likes.Set(nil)
}

// UnsetLikes ensures that no value is present for Likes, not even an explicit nil
func (o *JellyfinUserItemDataDto) UnsetLikes() {
	o.Likes.Unset()
}

// GetLastPlayedDate returns the LastPlayedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserItemDataDto) GetLastPlayedDate() time.Time {
	if o == nil || IsNil(o.LastPlayedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPlayedDate.Get()
}

// GetLastPlayedDateOk returns a tuple with the LastPlayedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPlayedDate.Get(), o.LastPlayedDate.IsSet()
}

// HasLastPlayedDate returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasLastPlayedDate() bool {
	if o != nil && o.LastPlayedDate.IsSet() {
		return true
	}

	return false
}

// SetLastPlayedDate gets a reference to the given NullableTime and assigns it to the LastPlayedDate field.
func (o *JellyfinUserItemDataDto) SetLastPlayedDate(v time.Time) {
	o.LastPlayedDate.Set(&v)
}
// SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil
func (o *JellyfinUserItemDataDto) SetLastPlayedDateNil() {
	o.LastPlayedDate.Set(nil)
}

// UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
func (o *JellyfinUserItemDataDto) UnsetLastPlayedDate() {
	o.LastPlayedDate.Unset()
}

// GetPlayed returns the Played field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetPlayed() bool {
	if o == nil || IsNil(o.Played) {
		var ret bool
		return ret
	}
	return *o.Played
}

// GetPlayedOk returns a tuple with the Played field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetPlayedOk() (*bool, bool) {
	if o == nil || IsNil(o.Played) {
		return nil, false
	}
	return o.Played, true
}

// HasPlayed returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasPlayed() bool {
	if o != nil && !IsNil(o.Played) {
		return true
	}

	return false
}

// SetPlayed gets a reference to the given bool and assigns it to the Played field.
func (o *JellyfinUserItemDataDto) SetPlayed(v bool) {
	o.Played = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *JellyfinUserItemDataDto) SetKey(v string) {
	o.Key = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *JellyfinUserItemDataDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserItemDataDto) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *JellyfinUserItemDataDto) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *JellyfinUserItemDataDto) SetItemId(v string) {
	o.ItemId = &v
}

func (o JellyfinUserItemDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinUserItemDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rating.IsSet() {
		toSerialize["Rating"] = o.Rating.Get()
	}
	if o.PlayedPercentage.IsSet() {
		toSerialize["PlayedPercentage"] = o.PlayedPercentage.Get()
	}
	if o.UnplayedItemCount.IsSet() {
		toSerialize["UnplayedItemCount"] = o.UnplayedItemCount.Get()
	}
	if !IsNil(o.PlaybackPositionTicks) {
		toSerialize["PlaybackPositionTicks"] = o.PlaybackPositionTicks
	}
	if !IsNil(o.PlayCount) {
		toSerialize["PlayCount"] = o.PlayCount
	}
	if !IsNil(o.IsFavorite) {
		toSerialize["IsFavorite"] = o.IsFavorite
	}
	if o.Likes.IsSet() {
		toSerialize["Likes"] = o.Likes.Get()
	}
	if o.LastPlayedDate.IsSet() {
		toSerialize["LastPlayedDate"] = o.LastPlayedDate.Get()
	}
	if !IsNil(o.Played) {
		toSerialize["Played"] = o.Played
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	return toSerialize, nil
}

type NullableJellyfinUserItemDataDto struct {
	value *JellyfinUserItemDataDto
	isSet bool
}

func (v NullableJellyfinUserItemDataDto) Get() *JellyfinUserItemDataDto {
	return v.value
}

func (v *NullableJellyfinUserItemDataDto) Set(val *JellyfinUserItemDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinUserItemDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinUserItemDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinUserItemDataDto(val *JellyfinUserItemDataDto) *NullableJellyfinUserItemDataDto {
	return &NullableJellyfinUserItemDataDto{value: val, isSet: true}
}

func (v NullableJellyfinUserItemDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinUserItemDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


