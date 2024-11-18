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

// JellyfinFileSystemEntryType Enum FileSystemEntryType.
type JellyfinFileSystemEntryType string

// List of FileSystemEntryType
const (
	JELLYFINFILESYSTEMENTRYTYPE_FILE JellyfinFileSystemEntryType = "File"
	JELLYFINFILESYSTEMENTRYTYPE_DIRECTORY JellyfinFileSystemEntryType = "Directory"
	JELLYFINFILESYSTEMENTRYTYPE_NETWORK_COMPUTER JellyfinFileSystemEntryType = "NetworkComputer"
	JELLYFINFILESYSTEMENTRYTYPE_NETWORK_SHARE JellyfinFileSystemEntryType = "NetworkShare"
)

// All allowed values of JellyfinFileSystemEntryType enum
var AllowedJellyfinFileSystemEntryTypeEnumValues = []JellyfinFileSystemEntryType{
	"File",
	"Directory",
	"NetworkComputer",
	"NetworkShare",
}

func (v *JellyfinFileSystemEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinFileSystemEntryType(value)
	for _, existing := range AllowedJellyfinFileSystemEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinFileSystemEntryType", value)
}

// NewJellyfinFileSystemEntryTypeFromValue returns a pointer to a valid JellyfinFileSystemEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinFileSystemEntryTypeFromValue(v string) (*JellyfinFileSystemEntryType, error) {
	ev := JellyfinFileSystemEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinFileSystemEntryType: valid values are %v", v, AllowedJellyfinFileSystemEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinFileSystemEntryType) IsValid() bool {
	for _, existing := range AllowedJellyfinFileSystemEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileSystemEntryType value
func (v JellyfinFileSystemEntryType) Ptr() *JellyfinFileSystemEntryType {
	return &v
}

type NullableJellyfinFileSystemEntryType struct {
	value *JellyfinFileSystemEntryType
	isSet bool
}

func (v NullableJellyfinFileSystemEntryType) Get() *JellyfinFileSystemEntryType {
	return v.value
}

func (v *NullableJellyfinFileSystemEntryType) Set(val *JellyfinFileSystemEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinFileSystemEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinFileSystemEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinFileSystemEntryType(val *JellyfinFileSystemEntryType) *NullableJellyfinFileSystemEntryType {
	return &NullableJellyfinFileSystemEntryType{value: val, isSet: true}
}

func (v NullableJellyfinFileSystemEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinFileSystemEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

