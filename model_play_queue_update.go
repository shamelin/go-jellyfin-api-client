/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PlayQueueUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayQueueUpdate{}

// PlayQueueUpdate Class PlayQueueUpdate.
type PlayQueueUpdate struct {
	// Gets the request type that originated this update.
	Reason *PlayQueueUpdateReason `json:"Reason,omitempty"`
	// Gets the UTC time of the last change to the playing queue.
	LastUpdate *time.Time `json:"LastUpdate,omitempty"`
	// Gets the playlist.
	Playlist []SyncPlayQueueItem `json:"Playlist,omitempty"`
	// Gets the playing item index in the playlist.
	PlayingItemIndex *int32 `json:"PlayingItemIndex,omitempty"`
	// Gets the start position ticks.
	StartPositionTicks *int64 `json:"StartPositionTicks,omitempty"`
	// Gets a value indicating whether the current item is playing.
	IsPlaying *bool `json:"IsPlaying,omitempty"`
	// Gets the shuffle mode.
	ShuffleMode *GroupShuffleMode `json:"ShuffleMode,omitempty"`
	// Gets the repeat mode.
	RepeatMode *GroupRepeatMode `json:"RepeatMode,omitempty"`
}

// NewPlayQueueUpdate instantiates a new PlayQueueUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayQueueUpdate() *PlayQueueUpdate {
	this := PlayQueueUpdate{}
	return &this
}

// NewPlayQueueUpdateWithDefaults instantiates a new PlayQueueUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayQueueUpdateWithDefaults() *PlayQueueUpdate {
	this := PlayQueueUpdate{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetReason() PlayQueueUpdateReason {
	if o == nil || IsNil(o.Reason) {
		var ret PlayQueueUpdateReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetReasonOk() (*PlayQueueUpdateReason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given PlayQueueUpdateReason and assigns it to the Reason field.
func (o *PlayQueueUpdate) SetReason(v PlayQueueUpdateReason) {
	o.Reason = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *PlayQueueUpdate) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetPlaylist returns the Playlist field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetPlaylist() []SyncPlayQueueItem {
	if o == nil || IsNil(o.Playlist) {
		var ret []SyncPlayQueueItem
		return ret
	}
	return o.Playlist
}

// GetPlaylistOk returns a tuple with the Playlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetPlaylistOk() ([]SyncPlayQueueItem, bool) {
	if o == nil || IsNil(o.Playlist) {
		return nil, false
	}
	return o.Playlist, true
}

// HasPlaylist returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasPlaylist() bool {
	if o != nil && !IsNil(o.Playlist) {
		return true
	}

	return false
}

// SetPlaylist gets a reference to the given []SyncPlayQueueItem and assigns it to the Playlist field.
func (o *PlayQueueUpdate) SetPlaylist(v []SyncPlayQueueItem) {
	o.Playlist = v
}

// GetPlayingItemIndex returns the PlayingItemIndex field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetPlayingItemIndex() int32 {
	if o == nil || IsNil(o.PlayingItemIndex) {
		var ret int32
		return ret
	}
	return *o.PlayingItemIndex
}

// GetPlayingItemIndexOk returns a tuple with the PlayingItemIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetPlayingItemIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayingItemIndex) {
		return nil, false
	}
	return o.PlayingItemIndex, true
}

// HasPlayingItemIndex returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasPlayingItemIndex() bool {
	if o != nil && !IsNil(o.PlayingItemIndex) {
		return true
	}

	return false
}

// SetPlayingItemIndex gets a reference to the given int32 and assigns it to the PlayingItemIndex field.
func (o *PlayQueueUpdate) SetPlayingItemIndex(v int32) {
	o.PlayingItemIndex = &v
}

// GetStartPositionTicks returns the StartPositionTicks field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetStartPositionTicks() int64 {
	if o == nil || IsNil(o.StartPositionTicks) {
		var ret int64
		return ret
	}
	return *o.StartPositionTicks
}

// GetStartPositionTicksOk returns a tuple with the StartPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetStartPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPositionTicks) {
		return nil, false
	}
	return o.StartPositionTicks, true
}

// HasStartPositionTicks returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasStartPositionTicks() bool {
	if o != nil && !IsNil(o.StartPositionTicks) {
		return true
	}

	return false
}

// SetStartPositionTicks gets a reference to the given int64 and assigns it to the StartPositionTicks field.
func (o *PlayQueueUpdate) SetStartPositionTicks(v int64) {
	o.StartPositionTicks = &v
}

// GetIsPlaying returns the IsPlaying field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetIsPlaying() bool {
	if o == nil || IsNil(o.IsPlaying) {
		var ret bool
		return ret
	}
	return *o.IsPlaying
}

// GetIsPlayingOk returns a tuple with the IsPlaying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetIsPlayingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlaying) {
		return nil, false
	}
	return o.IsPlaying, true
}

// HasIsPlaying returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasIsPlaying() bool {
	if o != nil && !IsNil(o.IsPlaying) {
		return true
	}

	return false
}

// SetIsPlaying gets a reference to the given bool and assigns it to the IsPlaying field.
func (o *PlayQueueUpdate) SetIsPlaying(v bool) {
	o.IsPlaying = &v
}

// GetShuffleMode returns the ShuffleMode field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetShuffleMode() GroupShuffleMode {
	if o == nil || IsNil(o.ShuffleMode) {
		var ret GroupShuffleMode
		return ret
	}
	return *o.ShuffleMode
}

// GetShuffleModeOk returns a tuple with the ShuffleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetShuffleModeOk() (*GroupShuffleMode, bool) {
	if o == nil || IsNil(o.ShuffleMode) {
		return nil, false
	}
	return o.ShuffleMode, true
}

// HasShuffleMode returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasShuffleMode() bool {
	if o != nil && !IsNil(o.ShuffleMode) {
		return true
	}

	return false
}

// SetShuffleMode gets a reference to the given GroupShuffleMode and assigns it to the ShuffleMode field.
func (o *PlayQueueUpdate) SetShuffleMode(v GroupShuffleMode) {
	o.ShuffleMode = &v
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *PlayQueueUpdate) GetRepeatMode() GroupRepeatMode {
	if o == nil || IsNil(o.RepeatMode) {
		var ret GroupRepeatMode
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdate) GetRepeatModeOk() (*GroupRepeatMode, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *PlayQueueUpdate) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given GroupRepeatMode and assigns it to the RepeatMode field.
func (o *PlayQueueUpdate) SetRepeatMode(v GroupRepeatMode) {
	o.RepeatMode = &v
}

func (o PlayQueueUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayQueueUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["LastUpdate"] = o.LastUpdate
	}
	if !IsNil(o.Playlist) {
		toSerialize["Playlist"] = o.Playlist
	}
	if !IsNil(o.PlayingItemIndex) {
		toSerialize["PlayingItemIndex"] = o.PlayingItemIndex
	}
	if !IsNil(o.StartPositionTicks) {
		toSerialize["StartPositionTicks"] = o.StartPositionTicks
	}
	if !IsNil(o.IsPlaying) {
		toSerialize["IsPlaying"] = o.IsPlaying
	}
	if !IsNil(o.ShuffleMode) {
		toSerialize["ShuffleMode"] = o.ShuffleMode
	}
	if !IsNil(o.RepeatMode) {
		toSerialize["RepeatMode"] = o.RepeatMode
	}
	return toSerialize, nil
}

type NullablePlayQueueUpdate struct {
	value *PlayQueueUpdate
	isSet bool
}

func (v NullablePlayQueueUpdate) Get() *PlayQueueUpdate {
	return v.value
}

func (v *NullablePlayQueueUpdate) Set(val *PlayQueueUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayQueueUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayQueueUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayQueueUpdate(val *PlayQueueUpdate) *NullablePlayQueueUpdate {
	return &NullablePlayQueueUpdate{value: val, isSet: true}
}

func (v NullablePlayQueueUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayQueueUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

