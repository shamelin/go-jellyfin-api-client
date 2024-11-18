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

// JellyfinGroupRepeatMode Enum GroupRepeatMode.
type JellyfinGroupRepeatMode string

// List of GroupRepeatMode
const (
	REPEAT_ONE JellyfinGroupRepeatMode = "RepeatOne"
	REPEAT_ALL JellyfinGroupRepeatMode = "RepeatAll"
	REPEAT_NONE JellyfinGroupRepeatMode = "RepeatNone"
)

// All allowed values of JellyfinGroupRepeatMode enum
var AllowedJellyfinGroupRepeatModeEnumValues = []JellyfinGroupRepeatMode{
	"RepeatOne",
	"RepeatAll",
	"RepeatNone",
}

func (v *JellyfinGroupRepeatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinGroupRepeatMode(value)
	for _, existing := range AllowedJellyfinGroupRepeatModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinGroupRepeatMode", value)
}

// NewJellyfinGroupRepeatModeFromValue returns a pointer to a valid JellyfinGroupRepeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinGroupRepeatModeFromValue(v string) (*JellyfinGroupRepeatMode, error) {
	ev := JellyfinGroupRepeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinGroupRepeatMode: valid values are %v", v, AllowedJellyfinGroupRepeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinGroupRepeatMode) IsValid() bool {
	for _, existing := range AllowedJellyfinGroupRepeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupRepeatMode value
func (v JellyfinGroupRepeatMode) Ptr() *JellyfinGroupRepeatMode {
	return &v
}

type NullableJellyfinGroupRepeatMode struct {
	value *JellyfinGroupRepeatMode
	isSet bool
}

func (v NullableJellyfinGroupRepeatMode) Get() *JellyfinGroupRepeatMode {
	return v.value
}

func (v *NullableJellyfinGroupRepeatMode) Set(val *JellyfinGroupRepeatMode) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinGroupRepeatMode) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinGroupRepeatMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinGroupRepeatMode(val *JellyfinGroupRepeatMode) *NullableJellyfinGroupRepeatMode {
	return &NullableJellyfinGroupRepeatMode{value: val, isSet: true}
}

func (v NullableJellyfinGroupRepeatMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinGroupRepeatMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

