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

// checks if the LyricLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LyricLine{}

// LyricLine Lyric model.
type LyricLine struct {
	// Gets the text of this lyric line.
	Text *string `json:"Text,omitempty"`
	// Gets the start time in ticks.
	Start NullableInt64 `json:"Start,omitempty"`
}

// NewLyricLine instantiates a new LyricLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLyricLine() *LyricLine {
	this := LyricLine{}
	return &this
}

// NewLyricLineWithDefaults instantiates a new LyricLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLyricLineWithDefaults() *LyricLine {
	this := LyricLine{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *LyricLine) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LyricLine) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *LyricLine) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *LyricLine) SetText(v string) {
	o.Text = &v
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricLine) GetStart() int64 {
	if o == nil || IsNil(o.Start.Get()) {
		var ret int64
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricLine) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *LyricLine) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableInt64 and assigns it to the Start field.
func (o *LyricLine) SetStart(v int64) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *LyricLine) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *LyricLine) UnsetStart() {
	o.Start.Unset()
}

func (o LyricLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LyricLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["Text"] = o.Text
	}
	if o.Start.IsSet() {
		toSerialize["Start"] = o.Start.Get()
	}
	return toSerialize, nil
}

type NullableLyricLine struct {
	value *LyricLine
	isSet bool
}

func (v NullableLyricLine) Get() *LyricLine {
	return v.value
}

func (v *NullableLyricLine) Set(val *LyricLine) {
	v.value = val
	v.isSet = true
}

func (v NullableLyricLine) IsSet() bool {
	return v.isSet
}

func (v *NullableLyricLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLyricLine(val *LyricLine) *NullableLyricLine {
	return &NullableLyricLine{value: val, isSet: true}
}

func (v NullableLyricLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLyricLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


