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

// JellyfinChannelMediaType the model 'JellyfinChannelMediaType'
type JellyfinChannelMediaType string

// List of ChannelMediaType
const (
	AUDIO JellyfinChannelMediaType = "Audio"
	VIDEO JellyfinChannelMediaType = "Video"
	PHOTO JellyfinChannelMediaType = "Photo"
)

// All allowed values of JellyfinChannelMediaType enum
var AllowedJellyfinChannelMediaTypeEnumValues = []JellyfinChannelMediaType{
	"Audio",
	"Video",
	"Photo",
}

func (v *JellyfinChannelMediaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinChannelMediaType(value)
	for _, existing := range AllowedJellyfinChannelMediaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinChannelMediaType", value)
}

// NewJellyfinChannelMediaTypeFromValue returns a pointer to a valid JellyfinChannelMediaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinChannelMediaTypeFromValue(v string) (*JellyfinChannelMediaType, error) {
	ev := JellyfinChannelMediaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinChannelMediaType: valid values are %v", v, AllowedJellyfinChannelMediaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinChannelMediaType) IsValid() bool {
	for _, existing := range AllowedJellyfinChannelMediaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelMediaType value
func (v JellyfinChannelMediaType) Ptr() *JellyfinChannelMediaType {
	return &v
}

type NullableJellyfinChannelMediaType struct {
	value *JellyfinChannelMediaType
	isSet bool
}

func (v NullableJellyfinChannelMediaType) Get() *JellyfinChannelMediaType {
	return v.value
}

func (v *NullableJellyfinChannelMediaType) Set(val *JellyfinChannelMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinChannelMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinChannelMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinChannelMediaType(val *JellyfinChannelMediaType) *NullableJellyfinChannelMediaType {
	return &NullableJellyfinChannelMediaType{value: val, isSet: true}
}

func (v NullableJellyfinChannelMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinChannelMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

