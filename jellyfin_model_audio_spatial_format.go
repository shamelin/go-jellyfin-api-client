/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"fmt"
)

// JellyfinAudioSpatialFormat An enum representing formats of spatial audio.
type JellyfinAudioSpatialFormat string

// List of AudioSpatialFormat
const (
	JELLYFINAUDIOSPATIALFORMAT_NONE JellyfinAudioSpatialFormat = "None"
	JELLYFINAUDIOSPATIALFORMAT_DOLBY_ATMOS JellyfinAudioSpatialFormat = "DolbyAtmos"
	JELLYFINAUDIOSPATIALFORMAT_DTSX JellyfinAudioSpatialFormat = "DTSX"
)

// All allowed values of JellyfinAudioSpatialFormat enum
var AllowedJellyfinAudioSpatialFormatEnumValues = []JellyfinAudioSpatialFormat{
	"None",
	"DolbyAtmos",
	"DTSX",
}

func (v *JellyfinAudioSpatialFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinAudioSpatialFormat(value)
	for _, existing := range AllowedJellyfinAudioSpatialFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinAudioSpatialFormat", value)
}

// NewJellyfinAudioSpatialFormatFromValue returns a pointer to a valid JellyfinAudioSpatialFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinAudioSpatialFormatFromValue(v string) (*JellyfinAudioSpatialFormat, error) {
	ev := JellyfinAudioSpatialFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinAudioSpatialFormat: valid values are %v", v, AllowedJellyfinAudioSpatialFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinAudioSpatialFormat) IsValid() bool {
	for _, existing := range AllowedJellyfinAudioSpatialFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudioSpatialFormat value
func (v JellyfinAudioSpatialFormat) Ptr() *JellyfinAudioSpatialFormat {
	return &v
}

type NullableJellyfinAudioSpatialFormat struct {
	value *JellyfinAudioSpatialFormat
	isSet bool
}

func (v NullableJellyfinAudioSpatialFormat) Get() *JellyfinAudioSpatialFormat {
	return v.value
}

func (v *NullableJellyfinAudioSpatialFormat) Set(val *JellyfinAudioSpatialFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinAudioSpatialFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinAudioSpatialFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinAudioSpatialFormat(val *JellyfinAudioSpatialFormat) *NullableJellyfinAudioSpatialFormat {
	return &NullableJellyfinAudioSpatialFormat{value: val, isSet: true}
}

func (v NullableJellyfinAudioSpatialFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinAudioSpatialFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

