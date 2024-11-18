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

// JellyfinTonemappingMode Enum containing tonemapping modes.
type JellyfinTonemappingMode string

// List of TonemappingMode
const (
	AUTO JellyfinTonemappingMode = "auto"
	MAX JellyfinTonemappingMode = "max"
	RGB JellyfinTonemappingMode = "rgb"
	LUM JellyfinTonemappingMode = "lum"
	ITP JellyfinTonemappingMode = "itp"
)

// All allowed values of JellyfinTonemappingMode enum
var AllowedJellyfinTonemappingModeEnumValues = []JellyfinTonemappingMode{
	"auto",
	"max",
	"rgb",
	"lum",
	"itp",
}

func (v *JellyfinTonemappingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinTonemappingMode(value)
	for _, existing := range AllowedJellyfinTonemappingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinTonemappingMode", value)
}

// NewJellyfinTonemappingModeFromValue returns a pointer to a valid JellyfinTonemappingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinTonemappingModeFromValue(v string) (*JellyfinTonemappingMode, error) {
	ev := JellyfinTonemappingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinTonemappingMode: valid values are %v", v, AllowedJellyfinTonemappingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinTonemappingMode) IsValid() bool {
	for _, existing := range AllowedJellyfinTonemappingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TonemappingMode value
func (v JellyfinTonemappingMode) Ptr() *JellyfinTonemappingMode {
	return &v
}

type NullableJellyfinTonemappingMode struct {
	value *JellyfinTonemappingMode
	isSet bool
}

func (v NullableJellyfinTonemappingMode) Get() *JellyfinTonemappingMode {
	return v.value
}

func (v *NullableJellyfinTonemappingMode) Set(val *JellyfinTonemappingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTonemappingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTonemappingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTonemappingMode(val *JellyfinTonemappingMode) *NullableJellyfinTonemappingMode {
	return &NullableJellyfinTonemappingMode{value: val, isSet: true}
}

func (v NullableJellyfinTonemappingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTonemappingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
