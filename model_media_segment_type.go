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

// MediaSegmentType Defines the types of content an individual Jellyfin.Data.Entities.MediaSegment represents.
type MediaSegmentType string

// List of MediaSegmentType
const (
	UNKNOWN MediaSegmentType = "Unknown"
	COMMERCIAL MediaSegmentType = "Commercial"
	PREVIEW MediaSegmentType = "Preview"
	RECAP MediaSegmentType = "Recap"
	OUTRO MediaSegmentType = "Outro"
	INTRO MediaSegmentType = "Intro"
)

// All allowed values of MediaSegmentType enum
var AllowedMediaSegmentTypeEnumValues = []MediaSegmentType{
	"Unknown",
	"Commercial",
	"Preview",
	"Recap",
	"Outro",
	"Intro",
}

func (v *MediaSegmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaSegmentType(value)
	for _, existing := range AllowedMediaSegmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaSegmentType", value)
}

// NewMediaSegmentTypeFromValue returns a pointer to a valid MediaSegmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaSegmentTypeFromValue(v string) (*MediaSegmentType, error) {
	ev := MediaSegmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaSegmentType: valid values are %v", v, AllowedMediaSegmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaSegmentType) IsValid() bool {
	for _, existing := range AllowedMediaSegmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaSegmentType value
func (v MediaSegmentType) Ptr() *MediaSegmentType {
	return &v
}

type NullableMediaSegmentType struct {
	value *MediaSegmentType
	isSet bool
}

func (v NullableMediaSegmentType) Get() *MediaSegmentType {
	return v.value
}

func (v *NullableMediaSegmentType) Set(val *MediaSegmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSegmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSegmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSegmentType(val *MediaSegmentType) *NullableMediaSegmentType {
	return &NullableMediaSegmentType{value: val, isSet: true}
}

func (v NullableMediaSegmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSegmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

