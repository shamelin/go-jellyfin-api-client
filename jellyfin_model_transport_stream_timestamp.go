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

// JellyfinTransportStreamTimestamp the model 'JellyfinTransportStreamTimestamp'
type JellyfinTransportStreamTimestamp string

// List of TransportStreamTimestamp
const (
	NONE JellyfinTransportStreamTimestamp = "None"
	ZERO JellyfinTransportStreamTimestamp = "Zero"
	VALID JellyfinTransportStreamTimestamp = "Valid"
)

// All allowed values of JellyfinTransportStreamTimestamp enum
var AllowedJellyfinTransportStreamTimestampEnumValues = []JellyfinTransportStreamTimestamp{
	"None",
	"Zero",
	"Valid",
}

func (v *JellyfinTransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinTransportStreamTimestamp(value)
	for _, existing := range AllowedJellyfinTransportStreamTimestampEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinTransportStreamTimestamp", value)
}

// NewJellyfinTransportStreamTimestampFromValue returns a pointer to a valid JellyfinTransportStreamTimestamp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinTransportStreamTimestampFromValue(v string) (*JellyfinTransportStreamTimestamp, error) {
	ev := JellyfinTransportStreamTimestamp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinTransportStreamTimestamp: valid values are %v", v, AllowedJellyfinTransportStreamTimestampEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinTransportStreamTimestamp) IsValid() bool {
	for _, existing := range AllowedJellyfinTransportStreamTimestampEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransportStreamTimestamp value
func (v JellyfinTransportStreamTimestamp) Ptr() *JellyfinTransportStreamTimestamp {
	return &v
}

type NullableJellyfinTransportStreamTimestamp struct {
	value *JellyfinTransportStreamTimestamp
	isSet bool
}

func (v NullableJellyfinTransportStreamTimestamp) Get() *JellyfinTransportStreamTimestamp {
	return v.value
}

func (v *NullableJellyfinTransportStreamTimestamp) Set(val *JellyfinTransportStreamTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTransportStreamTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTransportStreamTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTransportStreamTimestamp(val *JellyfinTransportStreamTimestamp) *NullableJellyfinTransportStreamTimestamp {
	return &NullableJellyfinTransportStreamTimestamp{value: val, isSet: true}
}

func (v NullableJellyfinTransportStreamTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

