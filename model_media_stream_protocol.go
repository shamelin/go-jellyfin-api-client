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

// MediaStreamProtocol Media streaming protocol.  Lowercase for backwards compatibility.
type MediaStreamProtocol string

// List of MediaStreamProtocol
const (
	HTTP MediaStreamProtocol = "http"
	HLS MediaStreamProtocol = "hls"
)

// All allowed values of MediaStreamProtocol enum
var AllowedMediaStreamProtocolEnumValues = []MediaStreamProtocol{
	"http",
	"hls",
}

func (v *MediaStreamProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaStreamProtocol(value)
	for _, existing := range AllowedMediaStreamProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaStreamProtocol", value)
}

// NewMediaStreamProtocolFromValue returns a pointer to a valid MediaStreamProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaStreamProtocolFromValue(v string) (*MediaStreamProtocol, error) {
	ev := MediaStreamProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaStreamProtocol: valid values are %v", v, AllowedMediaStreamProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaStreamProtocol) IsValid() bool {
	for _, existing := range AllowedMediaStreamProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaStreamProtocol value
func (v MediaStreamProtocol) Ptr() *MediaStreamProtocol {
	return &v
}

type NullableMediaStreamProtocol struct {
	value *MediaStreamProtocol
	isSet bool
}

func (v NullableMediaStreamProtocol) Get() *MediaStreamProtocol {
	return v.value
}

func (v *NullableMediaStreamProtocol) Set(val *MediaStreamProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamProtocol(val *MediaStreamProtocol) *NullableMediaStreamProtocol {
	return &NullableMediaStreamProtocol{value: val, isSet: true}
}

func (v NullableMediaStreamProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

