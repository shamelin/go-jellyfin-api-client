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

// VideoRangeType An enum representing types of video ranges.
type VideoRangeType string

// List of VideoRangeType
const (
	VIDEORANGETYPE_UNKNOWN VideoRangeType = "Unknown"
	VIDEORANGETYPE_SDR VideoRangeType = "SDR"
	VIDEORANGETYPE_HDR10 VideoRangeType = "HDR10"
	VIDEORANGETYPE_HLG VideoRangeType = "HLG"
	VIDEORANGETYPE_DOVI VideoRangeType = "DOVI"
	VIDEORANGETYPE_DOVI_WITH_HDR10 VideoRangeType = "DOVIWithHDR10"
	VIDEORANGETYPE_DOVI_WITH_HLG VideoRangeType = "DOVIWithHLG"
	VIDEORANGETYPE_DOVI_WITH_SDR VideoRangeType = "DOVIWithSDR"
	VIDEORANGETYPE_HDR10_PLUS VideoRangeType = "HDR10Plus"
)

// All allowed values of VideoRangeType enum
var AllowedVideoRangeTypeEnumValues = []VideoRangeType{
	"Unknown",
	"SDR",
	"HDR10",
	"HLG",
	"DOVI",
	"DOVIWithHDR10",
	"DOVIWithHLG",
	"DOVIWithSDR",
	"HDR10Plus",
}

func (v *VideoRangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoRangeType(value)
	for _, existing := range AllowedVideoRangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoRangeType", value)
}

// NewVideoRangeTypeFromValue returns a pointer to a valid VideoRangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoRangeTypeFromValue(v string) (*VideoRangeType, error) {
	ev := VideoRangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoRangeType: valid values are %v", v, AllowedVideoRangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoRangeType) IsValid() bool {
	for _, existing := range AllowedVideoRangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoRangeType value
func (v VideoRangeType) Ptr() *VideoRangeType {
	return &v
}

type NullableVideoRangeType struct {
	value *VideoRangeType
	isSet bool
}

func (v NullableVideoRangeType) Get() *VideoRangeType {
	return v.value
}

func (v *NullableVideoRangeType) Set(val *VideoRangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoRangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoRangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoRangeType(val *VideoRangeType) *NullableVideoRangeType {
	return &NullableVideoRangeType{value: val, isSet: true}
}

func (v NullableVideoRangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoRangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

