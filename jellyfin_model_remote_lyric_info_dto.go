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

// checks if the JellyfinRemoteLyricInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinRemoteLyricInfoDto{}

// JellyfinRemoteLyricInfoDto The remote lyric info dto.
type JellyfinRemoteLyricInfoDto struct {
	// Gets or sets the id for the lyric.
	Id *string `json:"Id,omitempty"`
	// Gets the provider name.
	ProviderName *string `json:"ProviderName,omitempty"`
	// Gets the lyrics.
	Lyrics *JellyfinLyricDto `json:"Lyrics,omitempty"`
}

// NewJellyfinRemoteLyricInfoDto instantiates a new JellyfinRemoteLyricInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinRemoteLyricInfoDto() *JellyfinRemoteLyricInfoDto {
	this := JellyfinRemoteLyricInfoDto{}
	return &this
}

// NewJellyfinRemoteLyricInfoDtoWithDefaults instantiates a new JellyfinRemoteLyricInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinRemoteLyricInfoDtoWithDefaults() *JellyfinRemoteLyricInfoDto {
	this := JellyfinRemoteLyricInfoDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JellyfinRemoteLyricInfoDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinRemoteLyricInfoDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JellyfinRemoteLyricInfoDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JellyfinRemoteLyricInfoDto) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *JellyfinRemoteLyricInfoDto) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinRemoteLyricInfoDto) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *JellyfinRemoteLyricInfoDto) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *JellyfinRemoteLyricInfoDto) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetLyrics returns the Lyrics field value if set, zero value otherwise.
func (o *JellyfinRemoteLyricInfoDto) GetLyrics() JellyfinLyricDto {
	if o == nil || IsNil(o.Lyrics) {
		var ret JellyfinLyricDto
		return ret
	}
	return *o.Lyrics
}

// GetLyricsOk returns a tuple with the Lyrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinRemoteLyricInfoDto) GetLyricsOk() (*JellyfinLyricDto, bool) {
	if o == nil || IsNil(o.Lyrics) {
		return nil, false
	}
	return o.Lyrics, true
}

// HasLyrics returns a boolean if a field has been set.
func (o *JellyfinRemoteLyricInfoDto) HasLyrics() bool {
	if o != nil && !IsNil(o.Lyrics) {
		return true
	}

	return false
}

// SetLyrics gets a reference to the given JellyfinLyricDto and assigns it to the Lyrics field.
func (o *JellyfinRemoteLyricInfoDto) SetLyrics(v JellyfinLyricDto) {
	o.Lyrics = &v
}

func (o JellyfinRemoteLyricInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinRemoteLyricInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.ProviderName) {
		toSerialize["ProviderName"] = o.ProviderName
	}
	if !IsNil(o.Lyrics) {
		toSerialize["Lyrics"] = o.Lyrics
	}
	return toSerialize, nil
}

type NullableJellyfinRemoteLyricInfoDto struct {
	value *JellyfinRemoteLyricInfoDto
	isSet bool
}

func (v NullableJellyfinRemoteLyricInfoDto) Get() *JellyfinRemoteLyricInfoDto {
	return v.value
}

func (v *NullableJellyfinRemoteLyricInfoDto) Set(val *JellyfinRemoteLyricInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinRemoteLyricInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinRemoteLyricInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinRemoteLyricInfoDto(val *JellyfinRemoteLyricInfoDto) *NullableJellyfinRemoteLyricInfoDto {
	return &NullableJellyfinRemoteLyricInfoDto{value: val, isSet: true}
}

func (v NullableJellyfinRemoteLyricInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinRemoteLyricInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


