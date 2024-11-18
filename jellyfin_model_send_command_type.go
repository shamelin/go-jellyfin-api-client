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

// JellyfinSendCommandType Enum SendCommandType.
type JellyfinSendCommandType string

// List of SendCommandType
const (
	UNPAUSE JellyfinSendCommandType = "Unpause"
	PAUSE JellyfinSendCommandType = "Pause"
	STOP JellyfinSendCommandType = "Stop"
	SEEK JellyfinSendCommandType = "Seek"
)

// All allowed values of JellyfinSendCommandType enum
var AllowedJellyfinSendCommandTypeEnumValues = []JellyfinSendCommandType{
	"Unpause",
	"Pause",
	"Stop",
	"Seek",
}

func (v *JellyfinSendCommandType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinSendCommandType(value)
	for _, existing := range AllowedJellyfinSendCommandTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinSendCommandType", value)
}

// NewJellyfinSendCommandTypeFromValue returns a pointer to a valid JellyfinSendCommandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinSendCommandTypeFromValue(v string) (*JellyfinSendCommandType, error) {
	ev := JellyfinSendCommandType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinSendCommandType: valid values are %v", v, AllowedJellyfinSendCommandTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinSendCommandType) IsValid() bool {
	for _, existing := range AllowedJellyfinSendCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SendCommandType value
func (v JellyfinSendCommandType) Ptr() *JellyfinSendCommandType {
	return &v
}

type NullableJellyfinSendCommandType struct {
	value *JellyfinSendCommandType
	isSet bool
}

func (v NullableJellyfinSendCommandType) Get() *JellyfinSendCommandType {
	return v.value
}

func (v *NullableJellyfinSendCommandType) Set(val *JellyfinSendCommandType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSendCommandType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSendCommandType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSendCommandType(val *JellyfinSendCommandType) *NullableJellyfinSendCommandType {
	return &NullableJellyfinSendCommandType{value: val, isSet: true}
}

func (v NullableJellyfinSendCommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSendCommandType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
