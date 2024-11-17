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

// TranscodeSeekInfo the model 'TranscodeSeekInfo'
type TranscodeSeekInfo string

// List of TranscodeSeekInfo
const (
	AUTO TranscodeSeekInfo = "Auto"
	BYTES TranscodeSeekInfo = "Bytes"
)

// All allowed values of TranscodeSeekInfo enum
var AllowedTranscodeSeekInfoEnumValues = []TranscodeSeekInfo{
	"Auto",
	"Bytes",
}

func (v *TranscodeSeekInfo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TranscodeSeekInfo(value)
	for _, existing := range AllowedTranscodeSeekInfoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TranscodeSeekInfo", value)
}

// NewTranscodeSeekInfoFromValue returns a pointer to a valid TranscodeSeekInfo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTranscodeSeekInfoFromValue(v string) (*TranscodeSeekInfo, error) {
	ev := TranscodeSeekInfo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TranscodeSeekInfo: valid values are %v", v, AllowedTranscodeSeekInfoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TranscodeSeekInfo) IsValid() bool {
	for _, existing := range AllowedTranscodeSeekInfoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TranscodeSeekInfo value
func (v TranscodeSeekInfo) Ptr() *TranscodeSeekInfo {
	return &v
}

type NullableTranscodeSeekInfo struct {
	value *TranscodeSeekInfo
	isSet bool
}

func (v NullableTranscodeSeekInfo) Get() *TranscodeSeekInfo {
	return v.value
}

func (v *NullableTranscodeSeekInfo) Set(val *TranscodeSeekInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscodeSeekInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscodeSeekInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscodeSeekInfo(val *TranscodeSeekInfo) *NullableTranscodeSeekInfo {
	return &NullableTranscodeSeekInfo{value: val, isSet: true}
}

func (v NullableTranscodeSeekInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscodeSeekInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

