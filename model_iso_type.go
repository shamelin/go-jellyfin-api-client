/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// IsoType Enum IsoType.
type IsoType string

// List of IsoType
const (
	DVD IsoType = "Dvd"
	BLU_RAY IsoType = "BluRay"
)

// All allowed values of IsoType enum
var AllowedIsoTypeEnumValues = []IsoType{
	"Dvd",
	"BluRay",
}

func (v *IsoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IsoType(value)
	for _, existing := range AllowedIsoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IsoType", value)
}

// NewIsoTypeFromValue returns a pointer to a valid IsoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIsoTypeFromValue(v string) (*IsoType, error) {
	ev := IsoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IsoType: valid values are %v", v, AllowedIsoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IsoType) IsValid() bool {
	for _, existing := range AllowedIsoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IsoType value
func (v IsoType) Ptr() *IsoType {
	return &v
}

type NullableIsoType struct {
	value *IsoType
	isSet bool
}

func (v NullableIsoType) Get() *IsoType {
	return v.value
}

func (v *NullableIsoType) Set(val *IsoType) {
	v.value = val
	v.isSet = true
}

func (v NullableIsoType) IsSet() bool {
	return v.isSet
}

func (v *NullableIsoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsoType(val *IsoType) *NullableIsoType {
	return &NullableIsoType{value: val, isSet: true}
}

func (v NullableIsoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

