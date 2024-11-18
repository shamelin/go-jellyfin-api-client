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

// checks if the JellyfinLyricDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinLyricDto{}

// JellyfinLyricDto LyricResponse model.
type JellyfinLyricDto struct {
	// Gets or sets Metadata for the lyrics.
	Metadata *JellyfinLyricMetadata `json:"Metadata,omitempty"`
	// Gets or sets a collection of individual lyric lines.
	Lyrics []JellyfinLyricLine `json:"Lyrics,omitempty"`
}

// NewJellyfinLyricDto instantiates a new JellyfinLyricDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinLyricDto() *JellyfinLyricDto {
	this := JellyfinLyricDto{}
	return &this
}

// NewJellyfinLyricDtoWithDefaults instantiates a new JellyfinLyricDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinLyricDtoWithDefaults() *JellyfinLyricDto {
	this := JellyfinLyricDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *JellyfinLyricDto) GetMetadata() JellyfinLyricMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret JellyfinLyricMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLyricDto) GetMetadataOk() (*JellyfinLyricMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *JellyfinLyricDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given JellyfinLyricMetadata and assigns it to the Metadata field.
func (o *JellyfinLyricDto) SetMetadata(v JellyfinLyricMetadata) {
	o.Metadata = &v
}

// GetLyrics returns the Lyrics field value if set, zero value otherwise.
func (o *JellyfinLyricDto) GetLyrics() []JellyfinLyricLine {
	if o == nil || IsNil(o.Lyrics) {
		var ret []JellyfinLyricLine
		return ret
	}
	return o.Lyrics
}

// GetLyricsOk returns a tuple with the Lyrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLyricDto) GetLyricsOk() ([]JellyfinLyricLine, bool) {
	if o == nil || IsNil(o.Lyrics) {
		return nil, false
	}
	return o.Lyrics, true
}

// HasLyrics returns a boolean if a field has been set.
func (o *JellyfinLyricDto) HasLyrics() bool {
	if o != nil && !IsNil(o.Lyrics) {
		return true
	}

	return false
}

// SetLyrics gets a reference to the given []JellyfinLyricLine and assigns it to the Lyrics field.
func (o *JellyfinLyricDto) SetLyrics(v []JellyfinLyricLine) {
	o.Lyrics = v
}

func (o JellyfinLyricDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinLyricDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Lyrics) {
		toSerialize["Lyrics"] = o.Lyrics
	}
	return toSerialize, nil
}

type NullableJellyfinLyricDto struct {
	value *JellyfinLyricDto
	isSet bool
}

func (v NullableJellyfinLyricDto) Get() *JellyfinLyricDto {
	return v.value
}

func (v *NullableJellyfinLyricDto) Set(val *JellyfinLyricDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinLyricDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinLyricDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinLyricDto(val *JellyfinLyricDto) *NullableJellyfinLyricDto {
	return &NullableJellyfinLyricDto{value: val, isSet: true}
}

func (v NullableJellyfinLyricDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinLyricDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

