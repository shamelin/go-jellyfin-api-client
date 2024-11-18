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

// TonemappingAlgorithm Enum containing tonemapping algorithms.
type TonemappingAlgorithm string

// List of TonemappingAlgorithm
const (
	TONEMAPPINGALGORITHM_NONE TonemappingAlgorithm = "none"
	TONEMAPPINGALGORITHM_CLIP TonemappingAlgorithm = "clip"
	TONEMAPPINGALGORITHM_LINEAR TonemappingAlgorithm = "linear"
	TONEMAPPINGALGORITHM_GAMMA TonemappingAlgorithm = "gamma"
	TONEMAPPINGALGORITHM_REINHARD TonemappingAlgorithm = "reinhard"
	TONEMAPPINGALGORITHM_HABLE TonemappingAlgorithm = "hable"
	TONEMAPPINGALGORITHM_MOBIUS TonemappingAlgorithm = "mobius"
	TONEMAPPINGALGORITHM_BT2390 TonemappingAlgorithm = "bt2390"
)

// All allowed values of TonemappingAlgorithm enum
var AllowedTonemappingAlgorithmEnumValues = []TonemappingAlgorithm{
	"none",
	"clip",
	"linear",
	"gamma",
	"reinhard",
	"hable",
	"mobius",
	"bt2390",
}

func (v *TonemappingAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TonemappingAlgorithm(value)
	for _, existing := range AllowedTonemappingAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TonemappingAlgorithm", value)
}

// NewTonemappingAlgorithmFromValue returns a pointer to a valid TonemappingAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTonemappingAlgorithmFromValue(v string) (*TonemappingAlgorithm, error) {
	ev := TonemappingAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TonemappingAlgorithm: valid values are %v", v, AllowedTonemappingAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TonemappingAlgorithm) IsValid() bool {
	for _, existing := range AllowedTonemappingAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TonemappingAlgorithm value
func (v TonemappingAlgorithm) Ptr() *TonemappingAlgorithm {
	return &v
}

type NullableTonemappingAlgorithm struct {
	value *TonemappingAlgorithm
	isSet bool
}

func (v NullableTonemappingAlgorithm) Get() *TonemappingAlgorithm {
	return v.value
}

func (v *NullableTonemappingAlgorithm) Set(val *TonemappingAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableTonemappingAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableTonemappingAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTonemappingAlgorithm(val *TonemappingAlgorithm) *NullableTonemappingAlgorithm {
	return &NullableTonemappingAlgorithm{value: val, isSet: true}
}

func (v NullableTonemappingAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTonemappingAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

