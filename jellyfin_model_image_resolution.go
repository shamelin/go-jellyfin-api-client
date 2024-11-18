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

// JellyfinImageResolution Enum ImageResolution.
type JellyfinImageResolution string

// List of ImageResolution
const (
	MATCH_SOURCE JellyfinImageResolution = "MatchSource"
	P144 JellyfinImageResolution = "P144"
	P240 JellyfinImageResolution = "P240"
	P360 JellyfinImageResolution = "P360"
	P480 JellyfinImageResolution = "P480"
	P720 JellyfinImageResolution = "P720"
	P1080 JellyfinImageResolution = "P1080"
	P1440 JellyfinImageResolution = "P1440"
	P2160 JellyfinImageResolution = "P2160"
)

// All allowed values of JellyfinImageResolution enum
var AllowedJellyfinImageResolutionEnumValues = []JellyfinImageResolution{
	"MatchSource",
	"P144",
	"P240",
	"P360",
	"P480",
	"P720",
	"P1080",
	"P1440",
	"P2160",
}

func (v *JellyfinImageResolution) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinImageResolution(value)
	for _, existing := range AllowedJellyfinImageResolutionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinImageResolution", value)
}

// NewJellyfinImageResolutionFromValue returns a pointer to a valid JellyfinImageResolution
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinImageResolutionFromValue(v string) (*JellyfinImageResolution, error) {
	ev := JellyfinImageResolution(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinImageResolution: valid values are %v", v, AllowedJellyfinImageResolutionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinImageResolution) IsValid() bool {
	for _, existing := range AllowedJellyfinImageResolutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageResolution value
func (v JellyfinImageResolution) Ptr() *JellyfinImageResolution {
	return &v
}

type NullableJellyfinImageResolution struct {
	value *JellyfinImageResolution
	isSet bool
}

func (v NullableJellyfinImageResolution) Get() *JellyfinImageResolution {
	return v.value
}

func (v *NullableJellyfinImageResolution) Set(val *JellyfinImageResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinImageResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinImageResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinImageResolution(val *JellyfinImageResolution) *NullableJellyfinImageResolution {
	return &NullableJellyfinImageResolution{value: val, isSet: true}
}

func (v NullableJellyfinImageResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinImageResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
