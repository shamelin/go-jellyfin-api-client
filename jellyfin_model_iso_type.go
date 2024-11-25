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

// JellyfinIsoType Enum IsoType.
type JellyfinIsoType string

// List of IsoType
const (
	JELLYFINISOTYPE_DVD JellyfinIsoType = "Dvd"
	JELLYFINISOTYPE_BLU_RAY JellyfinIsoType = "BluRay"
)

// All allowed values of JellyfinIsoType enum
var AllowedJellyfinIsoTypeEnumValues = []JellyfinIsoType{
	"Dvd",
	"BluRay",
}

func (v *JellyfinIsoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinIsoType(value)
	for _, existing := range AllowedJellyfinIsoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinIsoType", value)
}

// NewJellyfinIsoTypeFromValue returns a pointer to a valid JellyfinIsoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinIsoTypeFromValue(v string) (*JellyfinIsoType, error) {
	ev := JellyfinIsoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinIsoType: valid values are %v", v, AllowedJellyfinIsoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinIsoType) IsValid() bool {
	for _, existing := range AllowedJellyfinIsoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IsoType value
func (v JellyfinIsoType) Ptr() *JellyfinIsoType {
	return &v
}

type NullableJellyfinIsoType struct {
	value *JellyfinIsoType
	isSet bool
}

func (v NullableJellyfinIsoType) Get() *JellyfinIsoType {
	return v.value
}

func (v *NullableJellyfinIsoType) Set(val *JellyfinIsoType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinIsoType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinIsoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinIsoType(val *JellyfinIsoType) *NullableJellyfinIsoType {
	return &NullableJellyfinIsoType{value: val, isSet: true}
}

func (v NullableJellyfinIsoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinIsoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

