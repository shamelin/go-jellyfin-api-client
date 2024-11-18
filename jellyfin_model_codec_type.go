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

// JellyfinCodecType the model 'JellyfinCodecType'
type JellyfinCodecType string

// List of CodecType
const (
	VIDEO JellyfinCodecType = "Video"
	VIDEO_AUDIO JellyfinCodecType = "VideoAudio"
	AUDIO JellyfinCodecType = "Audio"
)

// All allowed values of JellyfinCodecType enum
var AllowedJellyfinCodecTypeEnumValues = []JellyfinCodecType{
	"Video",
	"VideoAudio",
	"Audio",
}

func (v *JellyfinCodecType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinCodecType(value)
	for _, existing := range AllowedJellyfinCodecTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinCodecType", value)
}

// NewJellyfinCodecTypeFromValue returns a pointer to a valid JellyfinCodecType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinCodecTypeFromValue(v string) (*JellyfinCodecType, error) {
	ev := JellyfinCodecType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinCodecType: valid values are %v", v, AllowedJellyfinCodecTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinCodecType) IsValid() bool {
	for _, existing := range AllowedJellyfinCodecTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CodecType value
func (v JellyfinCodecType) Ptr() *JellyfinCodecType {
	return &v
}

type NullableJellyfinCodecType struct {
	value *JellyfinCodecType
	isSet bool
}

func (v NullableJellyfinCodecType) Get() *JellyfinCodecType {
	return v.value
}

func (v *NullableJellyfinCodecType) Set(val *JellyfinCodecType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinCodecType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinCodecType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinCodecType(val *JellyfinCodecType) *NullableJellyfinCodecType {
	return &NullableJellyfinCodecType{value: val, isSet: true}
}

func (v NullableJellyfinCodecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinCodecType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
