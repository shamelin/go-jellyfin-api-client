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

// JellyfinMetadataRefreshMode the model 'JellyfinMetadataRefreshMode'
type JellyfinMetadataRefreshMode string

// List of MetadataRefreshMode
const (
	NONE JellyfinMetadataRefreshMode = "None"
	VALIDATION_ONLY JellyfinMetadataRefreshMode = "ValidationOnly"
	DEFAULT JellyfinMetadataRefreshMode = "Default"
	FULL_REFRESH JellyfinMetadataRefreshMode = "FullRefresh"
)

// All allowed values of JellyfinMetadataRefreshMode enum
var AllowedJellyfinMetadataRefreshModeEnumValues = []JellyfinMetadataRefreshMode{
	"None",
	"ValidationOnly",
	"Default",
	"FullRefresh",
}

func (v *JellyfinMetadataRefreshMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinMetadataRefreshMode(value)
	for _, existing := range AllowedJellyfinMetadataRefreshModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinMetadataRefreshMode", value)
}

// NewJellyfinMetadataRefreshModeFromValue returns a pointer to a valid JellyfinMetadataRefreshMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinMetadataRefreshModeFromValue(v string) (*JellyfinMetadataRefreshMode, error) {
	ev := JellyfinMetadataRefreshMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinMetadataRefreshMode: valid values are %v", v, AllowedJellyfinMetadataRefreshModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinMetadataRefreshMode) IsValid() bool {
	for _, existing := range AllowedJellyfinMetadataRefreshModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataRefreshMode value
func (v JellyfinMetadataRefreshMode) Ptr() *JellyfinMetadataRefreshMode {
	return &v
}

type NullableJellyfinMetadataRefreshMode struct {
	value *JellyfinMetadataRefreshMode
	isSet bool
}

func (v NullableJellyfinMetadataRefreshMode) Get() *JellyfinMetadataRefreshMode {
	return v.value
}

func (v *NullableJellyfinMetadataRefreshMode) Set(val *JellyfinMetadataRefreshMode) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMetadataRefreshMode) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMetadataRefreshMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMetadataRefreshMode(val *JellyfinMetadataRefreshMode) *NullableJellyfinMetadataRefreshMode {
	return &NullableJellyfinMetadataRefreshMode{value: val, isSet: true}
}

func (v NullableJellyfinMetadataRefreshMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMetadataRefreshMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
