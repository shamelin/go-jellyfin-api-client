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

// PlayAccess the model 'PlayAccess'
type PlayAccess string

// List of PlayAccess
const (
	FULL PlayAccess = "Full"
	NONE PlayAccess = "None"
)

// All allowed values of PlayAccess enum
var AllowedPlayAccessEnumValues = []PlayAccess{
	"Full",
	"None",
}

func (v *PlayAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlayAccess(value)
	for _, existing := range AllowedPlayAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlayAccess", value)
}

// NewPlayAccessFromValue returns a pointer to a valid PlayAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlayAccessFromValue(v string) (*PlayAccess, error) {
	ev := PlayAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlayAccess: valid values are %v", v, AllowedPlayAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlayAccess) IsValid() bool {
	for _, existing := range AllowedPlayAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayAccess value
func (v PlayAccess) Ptr() *PlayAccess {
	return &v
}

type NullablePlayAccess struct {
	value *PlayAccess
	isSet bool
}

func (v NullablePlayAccess) Get() *PlayAccess {
	return v.value
}

func (v *NullablePlayAccess) Set(val *PlayAccess) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayAccess) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayAccess(val *PlayAccess) *NullablePlayAccess {
	return &NullablePlayAccess{value: val, isSet: true}
}

func (v NullablePlayAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

