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

// JellyfinMediaStreamType Enum MediaStreamType.
type JellyfinMediaStreamType string

// List of MediaStreamType
const (
	AUDIO JellyfinMediaStreamType = "Audio"
	VIDEO JellyfinMediaStreamType = "Video"
	SUBTITLE JellyfinMediaStreamType = "Subtitle"
	EMBEDDED_IMAGE JellyfinMediaStreamType = "EmbeddedImage"
	DATA JellyfinMediaStreamType = "Data"
	LYRIC JellyfinMediaStreamType = "Lyric"
)

// All allowed values of JellyfinMediaStreamType enum
var AllowedJellyfinMediaStreamTypeEnumValues = []JellyfinMediaStreamType{
	"Audio",
	"Video",
	"Subtitle",
	"EmbeddedImage",
	"Data",
	"Lyric",
}

func (v *JellyfinMediaStreamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinMediaStreamType(value)
	for _, existing := range AllowedJellyfinMediaStreamTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinMediaStreamType", value)
}

// NewJellyfinMediaStreamTypeFromValue returns a pointer to a valid JellyfinMediaStreamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinMediaStreamTypeFromValue(v string) (*JellyfinMediaStreamType, error) {
	ev := JellyfinMediaStreamType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinMediaStreamType: valid values are %v", v, AllowedJellyfinMediaStreamTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinMediaStreamType) IsValid() bool {
	for _, existing := range AllowedJellyfinMediaStreamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaStreamType value
func (v JellyfinMediaStreamType) Ptr() *JellyfinMediaStreamType {
	return &v
}

type NullableJellyfinMediaStreamType struct {
	value *JellyfinMediaStreamType
	isSet bool
}

func (v NullableJellyfinMediaStreamType) Get() *JellyfinMediaStreamType {
	return v.value
}

func (v *NullableJellyfinMediaStreamType) Set(val *JellyfinMediaStreamType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMediaStreamType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMediaStreamType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMediaStreamType(val *JellyfinMediaStreamType) *NullableJellyfinMediaStreamType {
	return &NullableJellyfinMediaStreamType{value: val, isSet: true}
}

func (v NullableJellyfinMediaStreamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMediaStreamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
