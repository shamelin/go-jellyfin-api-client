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

// ProcessPriorityClass the model 'ProcessPriorityClass'
type ProcessPriorityClass string

// List of ProcessPriorityClass
const (
	PROCESSPRIORITYCLASS_NORMAL ProcessPriorityClass = "Normal"
	PROCESSPRIORITYCLASS_IDLE ProcessPriorityClass = "Idle"
	PROCESSPRIORITYCLASS_HIGH ProcessPriorityClass = "High"
	PROCESSPRIORITYCLASS_REAL_TIME ProcessPriorityClass = "RealTime"
	PROCESSPRIORITYCLASS_BELOW_NORMAL ProcessPriorityClass = "BelowNormal"
	PROCESSPRIORITYCLASS_ABOVE_NORMAL ProcessPriorityClass = "AboveNormal"
)

// All allowed values of ProcessPriorityClass enum
var AllowedProcessPriorityClassEnumValues = []ProcessPriorityClass{
	"Normal",
	"Idle",
	"High",
	"RealTime",
	"BelowNormal",
	"AboveNormal",
}

func (v *ProcessPriorityClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProcessPriorityClass(value)
	for _, existing := range AllowedProcessPriorityClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProcessPriorityClass", value)
}

// NewProcessPriorityClassFromValue returns a pointer to a valid ProcessPriorityClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProcessPriorityClassFromValue(v string) (*ProcessPriorityClass, error) {
	ev := ProcessPriorityClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProcessPriorityClass: valid values are %v", v, AllowedProcessPriorityClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProcessPriorityClass) IsValid() bool {
	for _, existing := range AllowedProcessPriorityClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProcessPriorityClass value
func (v ProcessPriorityClass) Ptr() *ProcessPriorityClass {
	return &v
}

type NullableProcessPriorityClass struct {
	value *ProcessPriorityClass
	isSet bool
}

func (v NullableProcessPriorityClass) Get() *ProcessPriorityClass {
	return v.value
}

func (v *NullableProcessPriorityClass) Set(val *ProcessPriorityClass) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessPriorityClass) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessPriorityClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessPriorityClass(val *ProcessPriorityClass) *NullableProcessPriorityClass {
	return &NullableProcessPriorityClass{value: val, isSet: true}
}

func (v NullableProcessPriorityClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessPriorityClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

