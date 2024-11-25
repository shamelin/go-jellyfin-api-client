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

// JellyfinMediaProtocol the model 'JellyfinMediaProtocol'
type JellyfinMediaProtocol string

// List of MediaProtocol
const (
	JELLYFINMEDIAPROTOCOL_FILE JellyfinMediaProtocol = "File"
	JELLYFINMEDIAPROTOCOL_HTTP JellyfinMediaProtocol = "Http"
	JELLYFINMEDIAPROTOCOL_RTMP JellyfinMediaProtocol = "Rtmp"
	JELLYFINMEDIAPROTOCOL_RTSP JellyfinMediaProtocol = "Rtsp"
	JELLYFINMEDIAPROTOCOL_UDP JellyfinMediaProtocol = "Udp"
	JELLYFINMEDIAPROTOCOL_RTP JellyfinMediaProtocol = "Rtp"
	JELLYFINMEDIAPROTOCOL_FTP JellyfinMediaProtocol = "Ftp"
)

// All allowed values of JellyfinMediaProtocol enum
var AllowedJellyfinMediaProtocolEnumValues = []JellyfinMediaProtocol{
	"File",
	"Http",
	"Rtmp",
	"Rtsp",
	"Udp",
	"Rtp",
	"Ftp",
}

func (v *JellyfinMediaProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinMediaProtocol(value)
	for _, existing := range AllowedJellyfinMediaProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinMediaProtocol", value)
}

// NewJellyfinMediaProtocolFromValue returns a pointer to a valid JellyfinMediaProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinMediaProtocolFromValue(v string) (*JellyfinMediaProtocol, error) {
	ev := JellyfinMediaProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinMediaProtocol: valid values are %v", v, AllowedJellyfinMediaProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinMediaProtocol) IsValid() bool {
	for _, existing := range AllowedJellyfinMediaProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaProtocol value
func (v JellyfinMediaProtocol) Ptr() *JellyfinMediaProtocol {
	return &v
}

type NullableJellyfinMediaProtocol struct {
	value *JellyfinMediaProtocol
	isSet bool
}

func (v NullableJellyfinMediaProtocol) Get() *JellyfinMediaProtocol {
	return v.value
}

func (v *NullableJellyfinMediaProtocol) Set(val *JellyfinMediaProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMediaProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMediaProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMediaProtocol(val *JellyfinMediaProtocol) *NullableJellyfinMediaProtocol {
	return &NullableJellyfinMediaProtocol{value: val, isSet: true}
}

func (v NullableJellyfinMediaProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMediaProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

