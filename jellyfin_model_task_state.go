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

// JellyfinTaskState Enum TaskState.
type JellyfinTaskState string

// List of TaskState
const (
	IDLE JellyfinTaskState = "Idle"
	CANCELLING JellyfinTaskState = "Cancelling"
	RUNNING JellyfinTaskState = "Running"
)

// All allowed values of JellyfinTaskState enum
var AllowedJellyfinTaskStateEnumValues = []JellyfinTaskState{
	"Idle",
	"Cancelling",
	"Running",
}

func (v *JellyfinTaskState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinTaskState(value)
	for _, existing := range AllowedJellyfinTaskStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinTaskState", value)
}

// NewJellyfinTaskStateFromValue returns a pointer to a valid JellyfinTaskState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinTaskStateFromValue(v string) (*JellyfinTaskState, error) {
	ev := JellyfinTaskState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinTaskState: valid values are %v", v, AllowedJellyfinTaskStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinTaskState) IsValid() bool {
	for _, existing := range AllowedJellyfinTaskStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskState value
func (v JellyfinTaskState) Ptr() *JellyfinTaskState {
	return &v
}

type NullableJellyfinTaskState struct {
	value *JellyfinTaskState
	isSet bool
}

func (v NullableJellyfinTaskState) Get() *JellyfinTaskState {
	return v.value
}

func (v *NullableJellyfinTaskState) Set(val *JellyfinTaskState) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTaskState) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTaskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTaskState(val *JellyfinTaskState) *NullableJellyfinTaskState {
	return &NullableJellyfinTaskState{value: val, isSet: true}
}

func (v NullableJellyfinTaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTaskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
