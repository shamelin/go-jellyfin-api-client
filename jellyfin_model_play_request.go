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

// checks if the JellyfinPlayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinPlayRequest{}

// JellyfinPlayRequest Class PlayRequest.
type JellyfinPlayRequest struct {
	// Gets or sets the item ids.
	ItemIds []string `json:"ItemIds,omitempty"`
	// Gets or sets the start position ticks that the first item should be played at.
	StartPositionTicks NullableInt64 `json:"StartPositionTicks,omitempty"`
	// Gets or sets the play command.
	PlayCommand *JellyfinPlayCommand `json:"PlayCommand,omitempty"`
	// Gets or sets the controlling user identifier.
	ControllingUserId *string `json:"ControllingUserId,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	MediaSourceId NullableString `json:"MediaSourceId,omitempty"`
	StartIndex NullableInt32 `json:"StartIndex,omitempty"`
}

// NewJellyfinPlayRequest instantiates a new JellyfinPlayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinPlayRequest() *JellyfinPlayRequest {
	this := JellyfinPlayRequest{}
	return &this
}

// NewJellyfinPlayRequestWithDefaults instantiates a new JellyfinPlayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinPlayRequestWithDefaults() *JellyfinPlayRequest {
	this := JellyfinPlayRequest{}
	return &this
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetItemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemIds) {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasItemIds() bool {
	if o != nil && !IsNil(o.ItemIds) {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *JellyfinPlayRequest) SetItemIds(v []string) {
	o.ItemIds = v
}

// GetStartPositionTicks returns the StartPositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetStartPositionTicks() int64 {
	if o == nil || IsNil(o.StartPositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.StartPositionTicks.Get()
}

// GetStartPositionTicksOk returns a tuple with the StartPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetStartPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartPositionTicks.Get(), o.StartPositionTicks.IsSet()
}

// HasStartPositionTicks returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasStartPositionTicks() bool {
	if o != nil && o.StartPositionTicks.IsSet() {
		return true
	}

	return false
}

// SetStartPositionTicks gets a reference to the given NullableInt64 and assigns it to the StartPositionTicks field.
func (o *JellyfinPlayRequest) SetStartPositionTicks(v int64) {
	o.StartPositionTicks.Set(&v)
}
// SetStartPositionTicksNil sets the value for StartPositionTicks to be an explicit nil
func (o *JellyfinPlayRequest) SetStartPositionTicksNil() {
	o.StartPositionTicks.Set(nil)
}

// UnsetStartPositionTicks ensures that no value is present for StartPositionTicks, not even an explicit nil
func (o *JellyfinPlayRequest) UnsetStartPositionTicks() {
	o.StartPositionTicks.Unset()
}

// GetPlayCommand returns the PlayCommand field value if set, zero value otherwise.
func (o *JellyfinPlayRequest) GetPlayCommand() JellyfinPlayCommand {
	if o == nil || IsNil(o.PlayCommand) {
		var ret JellyfinPlayCommand
		return ret
	}
	return *o.PlayCommand
}

// GetPlayCommandOk returns a tuple with the PlayCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlayRequest) GetPlayCommandOk() (*JellyfinPlayCommand, bool) {
	if o == nil || IsNil(o.PlayCommand) {
		return nil, false
	}
	return o.PlayCommand, true
}

// HasPlayCommand returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasPlayCommand() bool {
	if o != nil && !IsNil(o.PlayCommand) {
		return true
	}

	return false
}

// SetPlayCommand gets a reference to the given JellyfinPlayCommand and assigns it to the PlayCommand field.
func (o *JellyfinPlayRequest) SetPlayCommand(v JellyfinPlayCommand) {
	o.PlayCommand = &v
}

// GetControllingUserId returns the ControllingUserId field value if set, zero value otherwise.
func (o *JellyfinPlayRequest) GetControllingUserId() string {
	if o == nil || IsNil(o.ControllingUserId) {
		var ret string
		return ret
	}
	return *o.ControllingUserId
}

// GetControllingUserIdOk returns a tuple with the ControllingUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPlayRequest) GetControllingUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControllingUserId) {
		return nil, false
	}
	return o.ControllingUserId, true
}

// HasControllingUserId returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasControllingUserId() bool {
	if o != nil && !IsNil(o.ControllingUserId) {
		return true
	}

	return false
}

// SetControllingUserId gets a reference to the given string and assigns it to the ControllingUserId field.
func (o *JellyfinPlayRequest) SetControllingUserId(v string) {
	o.ControllingUserId = &v
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *JellyfinPlayRequest) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *JellyfinPlayRequest) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *JellyfinPlayRequest) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *JellyfinPlayRequest) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *JellyfinPlayRequest) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *JellyfinPlayRequest) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.MediaSourceId.Get()
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetMediaSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaSourceId.Get(), o.MediaSourceId.IsSet()
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasMediaSourceId() bool {
	if o != nil && o.MediaSourceId.IsSet() {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given NullableString and assigns it to the MediaSourceId field.
func (o *JellyfinPlayRequest) SetMediaSourceId(v string) {
	o.MediaSourceId.Set(&v)
}
// SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil
func (o *JellyfinPlayRequest) SetMediaSourceIdNil() {
	o.MediaSourceId.Set(nil)
}

// UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
func (o *JellyfinPlayRequest) UnsetMediaSourceId() {
	o.MediaSourceId.Unset()
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPlayRequest) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.StartIndex.Get()
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPlayRequest) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartIndex.Get(), o.StartIndex.IsSet()
}

// HasStartIndex returns a boolean if a field has been set.
func (o *JellyfinPlayRequest) HasStartIndex() bool {
	if o != nil && o.StartIndex.IsSet() {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given NullableInt32 and assigns it to the StartIndex field.
func (o *JellyfinPlayRequest) SetStartIndex(v int32) {
	o.StartIndex.Set(&v)
}
// SetStartIndexNil sets the value for StartIndex to be an explicit nil
func (o *JellyfinPlayRequest) SetStartIndexNil() {
	o.StartIndex.Set(nil)
}

// UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
func (o *JellyfinPlayRequest) UnsetStartIndex() {
	o.StartIndex.Unset()
}

func (o JellyfinPlayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinPlayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemIds != nil {
		toSerialize["ItemIds"] = o.ItemIds
	}
	if o.StartPositionTicks.IsSet() {
		toSerialize["StartPositionTicks"] = o.StartPositionTicks.Get()
	}
	if !IsNil(o.PlayCommand) {
		toSerialize["PlayCommand"] = o.PlayCommand
	}
	if !IsNil(o.ControllingUserId) {
		toSerialize["ControllingUserId"] = o.ControllingUserId
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.MediaSourceId.IsSet() {
		toSerialize["MediaSourceId"] = o.MediaSourceId.Get()
	}
	if o.StartIndex.IsSet() {
		toSerialize["StartIndex"] = o.StartIndex.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinPlayRequest struct {
	value *JellyfinPlayRequest
	isSet bool
}

func (v NullableJellyfinPlayRequest) Get() *JellyfinPlayRequest {
	return v.value
}

func (v *NullableJellyfinPlayRequest) Set(val *JellyfinPlayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPlayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPlayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPlayRequest(val *JellyfinPlayRequest) *NullableJellyfinPlayRequest {
	return &NullableJellyfinPlayRequest{value: val, isSet: true}
}

func (v NullableJellyfinPlayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPlayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


