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

// JellyfinImageSavingConvention the model 'JellyfinImageSavingConvention'
type JellyfinImageSavingConvention string

// List of ImageSavingConvention
const (
	LEGACY JellyfinImageSavingConvention = "Legacy"
	COMPATIBLE JellyfinImageSavingConvention = "Compatible"
)

// All allowed values of JellyfinImageSavingConvention enum
var AllowedJellyfinImageSavingConventionEnumValues = []JellyfinImageSavingConvention{
	"Legacy",
	"Compatible",
}

func (v *JellyfinImageSavingConvention) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinImageSavingConvention(value)
	for _, existing := range AllowedJellyfinImageSavingConventionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinImageSavingConvention", value)
}

// NewJellyfinImageSavingConventionFromValue returns a pointer to a valid JellyfinImageSavingConvention
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinImageSavingConventionFromValue(v string) (*JellyfinImageSavingConvention, error) {
	ev := JellyfinImageSavingConvention(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinImageSavingConvention: valid values are %v", v, AllowedJellyfinImageSavingConventionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinImageSavingConvention) IsValid() bool {
	for _, existing := range AllowedJellyfinImageSavingConventionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageSavingConvention value
func (v JellyfinImageSavingConvention) Ptr() *JellyfinImageSavingConvention {
	return &v
}

type NullableJellyfinImageSavingConvention struct {
	value *JellyfinImageSavingConvention
	isSet bool
}

func (v NullableJellyfinImageSavingConvention) Get() *JellyfinImageSavingConvention {
	return v.value
}

func (v *NullableJellyfinImageSavingConvention) Set(val *JellyfinImageSavingConvention) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinImageSavingConvention) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinImageSavingConvention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinImageSavingConvention(val *JellyfinImageSavingConvention) *NullableJellyfinImageSavingConvention {
	return &NullableJellyfinImageSavingConvention{value: val, isSet: true}
}

func (v NullableJellyfinImageSavingConvention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinImageSavingConvention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

