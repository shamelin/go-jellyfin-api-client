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

// checks if the JellyfinQueryFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinQueryFilters{}

// JellyfinQueryFilters struct for JellyfinQueryFilters
type JellyfinQueryFilters struct {
	Genres []JellyfinNameGuidPair `json:"Genres,omitempty"`
	Tags []string `json:"Tags,omitempty"`
}

// NewJellyfinQueryFilters instantiates a new JellyfinQueryFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinQueryFilters() *JellyfinQueryFilters {
	this := JellyfinQueryFilters{}
	return &this
}

// NewJellyfinQueryFiltersWithDefaults instantiates a new JellyfinQueryFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinQueryFiltersWithDefaults() *JellyfinQueryFilters {
	this := JellyfinQueryFilters{}
	return &this
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinQueryFilters) GetGenres() []JellyfinNameGuidPair {
	if o == nil {
		var ret []JellyfinNameGuidPair
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinQueryFilters) GetGenresOk() ([]JellyfinNameGuidPair, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *JellyfinQueryFilters) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []JellyfinNameGuidPair and assigns it to the Genres field.
func (o *JellyfinQueryFilters) SetGenres(v []JellyfinNameGuidPair) {
	o.Genres = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinQueryFilters) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinQueryFilters) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *JellyfinQueryFilters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *JellyfinQueryFilters) SetTags(v []string) {
	o.Tags = v
}

func (o JellyfinQueryFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinQueryFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Genres != nil {
		toSerialize["Genres"] = o.Genres
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableJellyfinQueryFilters struct {
	value *JellyfinQueryFilters
	isSet bool
}

func (v NullableJellyfinQueryFilters) Get() *JellyfinQueryFilters {
	return v.value
}

func (v *NullableJellyfinQueryFilters) Set(val *JellyfinQueryFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinQueryFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinQueryFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinQueryFilters(val *JellyfinQueryFilters) *NullableJellyfinQueryFilters {
	return &NullableJellyfinQueryFilters{value: val, isSet: true}
}

func (v NullableJellyfinQueryFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinQueryFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


