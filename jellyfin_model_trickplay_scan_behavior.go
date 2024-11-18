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

// JellyfinTrickplayScanBehavior Enum TrickplayScanBehavior.
type JellyfinTrickplayScanBehavior string

// List of TrickplayScanBehavior
const (
	JELLYFINTRICKPLAYSCANBEHAVIOR_BLOCKING JellyfinTrickplayScanBehavior = "Blocking"
	JELLYFINTRICKPLAYSCANBEHAVIOR_NON_BLOCKING JellyfinTrickplayScanBehavior = "NonBlocking"
)

// All allowed values of JellyfinTrickplayScanBehavior enum
var AllowedJellyfinTrickplayScanBehaviorEnumValues = []JellyfinTrickplayScanBehavior{
	"Blocking",
	"NonBlocking",
}

func (v *JellyfinTrickplayScanBehavior) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinTrickplayScanBehavior(value)
	for _, existing := range AllowedJellyfinTrickplayScanBehaviorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinTrickplayScanBehavior", value)
}

// NewJellyfinTrickplayScanBehaviorFromValue returns a pointer to a valid JellyfinTrickplayScanBehavior
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinTrickplayScanBehaviorFromValue(v string) (*JellyfinTrickplayScanBehavior, error) {
	ev := JellyfinTrickplayScanBehavior(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinTrickplayScanBehavior: valid values are %v", v, AllowedJellyfinTrickplayScanBehaviorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinTrickplayScanBehavior) IsValid() bool {
	for _, existing := range AllowedJellyfinTrickplayScanBehaviorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrickplayScanBehavior value
func (v JellyfinTrickplayScanBehavior) Ptr() *JellyfinTrickplayScanBehavior {
	return &v
}

type NullableJellyfinTrickplayScanBehavior struct {
	value *JellyfinTrickplayScanBehavior
	isSet bool
}

func (v NullableJellyfinTrickplayScanBehavior) Get() *JellyfinTrickplayScanBehavior {
	return v.value
}

func (v *NullableJellyfinTrickplayScanBehavior) Set(val *JellyfinTrickplayScanBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTrickplayScanBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTrickplayScanBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTrickplayScanBehavior(val *JellyfinTrickplayScanBehavior) *NullableJellyfinTrickplayScanBehavior {
	return &NullableJellyfinTrickplayScanBehavior{value: val, isSet: true}
}

func (v NullableJellyfinTrickplayScanBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTrickplayScanBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

