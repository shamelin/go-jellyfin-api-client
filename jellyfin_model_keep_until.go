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

// JellyfinKeepUntil the model 'JellyfinKeepUntil'
type JellyfinKeepUntil string

// List of KeepUntil
const (
	JELLYFINKEEPUNTIL_UNTIL_DELETED JellyfinKeepUntil = "UntilDeleted"
	JELLYFINKEEPUNTIL_UNTIL_SPACE_NEEDED JellyfinKeepUntil = "UntilSpaceNeeded"
	JELLYFINKEEPUNTIL_UNTIL_WATCHED JellyfinKeepUntil = "UntilWatched"
	JELLYFINKEEPUNTIL_UNTIL_DATE JellyfinKeepUntil = "UntilDate"
)

// All allowed values of JellyfinKeepUntil enum
var AllowedJellyfinKeepUntilEnumValues = []JellyfinKeepUntil{
	"UntilDeleted",
	"UntilSpaceNeeded",
	"UntilWatched",
	"UntilDate",
}

func (v *JellyfinKeepUntil) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinKeepUntil(value)
	for _, existing := range AllowedJellyfinKeepUntilEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinKeepUntil", value)
}

// NewJellyfinKeepUntilFromValue returns a pointer to a valid JellyfinKeepUntil
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinKeepUntilFromValue(v string) (*JellyfinKeepUntil, error) {
	ev := JellyfinKeepUntil(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinKeepUntil: valid values are %v", v, AllowedJellyfinKeepUntilEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinKeepUntil) IsValid() bool {
	for _, existing := range AllowedJellyfinKeepUntilEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeepUntil value
func (v JellyfinKeepUntil) Ptr() *JellyfinKeepUntil {
	return &v
}

type NullableJellyfinKeepUntil struct {
	value *JellyfinKeepUntil
	isSet bool
}

func (v NullableJellyfinKeepUntil) Get() *JellyfinKeepUntil {
	return v.value
}

func (v *NullableJellyfinKeepUntil) Set(val *JellyfinKeepUntil) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinKeepUntil) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinKeepUntil) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinKeepUntil(val *JellyfinKeepUntil) *NullableJellyfinKeepUntil {
	return &NullableJellyfinKeepUntil{value: val, isSet: true}
}

func (v NullableJellyfinKeepUntil) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinKeepUntil) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

