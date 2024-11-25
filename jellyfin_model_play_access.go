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

// JellyfinPlayAccess the model 'JellyfinPlayAccess'
type JellyfinPlayAccess string

// List of PlayAccess
const (
	JELLYFINPLAYACCESS_FULL JellyfinPlayAccess = "Full"
	JELLYFINPLAYACCESS_NONE JellyfinPlayAccess = "None"
)

// All allowed values of JellyfinPlayAccess enum
var AllowedJellyfinPlayAccessEnumValues = []JellyfinPlayAccess{
	"Full",
	"None",
}

func (v *JellyfinPlayAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinPlayAccess(value)
	for _, existing := range AllowedJellyfinPlayAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinPlayAccess", value)
}

// NewJellyfinPlayAccessFromValue returns a pointer to a valid JellyfinPlayAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinPlayAccessFromValue(v string) (*JellyfinPlayAccess, error) {
	ev := JellyfinPlayAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinPlayAccess: valid values are %v", v, AllowedJellyfinPlayAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinPlayAccess) IsValid() bool {
	for _, existing := range AllowedJellyfinPlayAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayAccess value
func (v JellyfinPlayAccess) Ptr() *JellyfinPlayAccess {
	return &v
}

type NullableJellyfinPlayAccess struct {
	value *JellyfinPlayAccess
	isSet bool
}

func (v NullableJellyfinPlayAccess) Get() *JellyfinPlayAccess {
	return v.value
}

func (v *NullableJellyfinPlayAccess) Set(val *JellyfinPlayAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPlayAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPlayAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPlayAccess(val *JellyfinPlayAccess) *NullableJellyfinPlayAccess {
	return &NullableJellyfinPlayAccess{value: val, isSet: true}
}

func (v NullableJellyfinPlayAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPlayAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

