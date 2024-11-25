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

// JellyfinLocationType Enum LocationType.
type JellyfinLocationType string

// List of LocationType
const (
	JELLYFINLOCATIONTYPE_FILE_SYSTEM JellyfinLocationType = "FileSystem"
	JELLYFINLOCATIONTYPE_REMOTE JellyfinLocationType = "Remote"
	JELLYFINLOCATIONTYPE_VIRTUAL JellyfinLocationType = "Virtual"
	JELLYFINLOCATIONTYPE_OFFLINE JellyfinLocationType = "Offline"
)

// All allowed values of JellyfinLocationType enum
var AllowedJellyfinLocationTypeEnumValues = []JellyfinLocationType{
	"FileSystem",
	"Remote",
	"Virtual",
	"Offline",
}

func (v *JellyfinLocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinLocationType(value)
	for _, existing := range AllowedJellyfinLocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinLocationType", value)
}

// NewJellyfinLocationTypeFromValue returns a pointer to a valid JellyfinLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinLocationTypeFromValue(v string) (*JellyfinLocationType, error) {
	ev := JellyfinLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinLocationType: valid values are %v", v, AllowedJellyfinLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinLocationType) IsValid() bool {
	for _, existing := range AllowedJellyfinLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationType value
func (v JellyfinLocationType) Ptr() *JellyfinLocationType {
	return &v
}

type NullableJellyfinLocationType struct {
	value *JellyfinLocationType
	isSet bool
}

func (v NullableJellyfinLocationType) Get() *JellyfinLocationType {
	return v.value
}

func (v *NullableJellyfinLocationType) Set(val *JellyfinLocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinLocationType(val *JellyfinLocationType) *NullableJellyfinLocationType {
	return &NullableJellyfinLocationType{value: val, isSet: true}
}

func (v NullableJellyfinLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

