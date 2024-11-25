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

// checks if the JellyfinQueueItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinQueueItem{}

// JellyfinQueueItem struct for JellyfinQueueItem
type JellyfinQueueItem struct {
	Id *string `json:"Id,omitempty"`
	PlaylistItemId NullableString `json:"PlaylistItemId,omitempty"`
}

// NewJellyfinQueueItem instantiates a new JellyfinQueueItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinQueueItem() *JellyfinQueueItem {
	this := JellyfinQueueItem{}
	return &this
}

// NewJellyfinQueueItemWithDefaults instantiates a new JellyfinQueueItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinQueueItemWithDefaults() *JellyfinQueueItem {
	this := JellyfinQueueItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JellyfinQueueItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinQueueItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JellyfinQueueItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JellyfinQueueItem) SetId(v string) {
	o.Id = &v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinQueueItem) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId.Get()
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinQueueItem) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaylistItemId.Get(), o.PlaylistItemId.IsSet()
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *JellyfinQueueItem) HasPlaylistItemId() bool {
	if o != nil && o.PlaylistItemId.IsSet() {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given NullableString and assigns it to the PlaylistItemId field.
func (o *JellyfinQueueItem) SetPlaylistItemId(v string) {
	o.PlaylistItemId.Set(&v)
}
// SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil
func (o *JellyfinQueueItem) SetPlaylistItemIdNil() {
	o.PlaylistItemId.Set(nil)
}

// UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
func (o *JellyfinQueueItem) UnsetPlaylistItemId() {
	o.PlaylistItemId.Unset()
}

func (o JellyfinQueueItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinQueueItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.PlaylistItemId.IsSet() {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinQueueItem struct {
	value *JellyfinQueueItem
	isSet bool
}

func (v NullableJellyfinQueueItem) Get() *JellyfinQueueItem {
	return v.value
}

func (v *NullableJellyfinQueueItem) Set(val *JellyfinQueueItem) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinQueueItem) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinQueueItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinQueueItem(val *JellyfinQueueItem) *NullableJellyfinQueueItem {
	return &NullableJellyfinQueueItem{value: val, isSet: true}
}

func (v NullableJellyfinQueueItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinQueueItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


