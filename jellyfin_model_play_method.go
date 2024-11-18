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

// JellyfinPlayMethod the model 'JellyfinPlayMethod'
type JellyfinPlayMethod string

// List of PlayMethod
const (
	JELLYFINPLAYMETHOD_TRANSCODE JellyfinPlayMethod = "Transcode"
	JELLYFINPLAYMETHOD_DIRECT_STREAM JellyfinPlayMethod = "DirectStream"
	JELLYFINPLAYMETHOD_DIRECT_PLAY JellyfinPlayMethod = "DirectPlay"
)

// All allowed values of JellyfinPlayMethod enum
var AllowedJellyfinPlayMethodEnumValues = []JellyfinPlayMethod{
	"Transcode",
	"DirectStream",
	"DirectPlay",
}

func (v *JellyfinPlayMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinPlayMethod(value)
	for _, existing := range AllowedJellyfinPlayMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinPlayMethod", value)
}

// NewJellyfinPlayMethodFromValue returns a pointer to a valid JellyfinPlayMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinPlayMethodFromValue(v string) (*JellyfinPlayMethod, error) {
	ev := JellyfinPlayMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinPlayMethod: valid values are %v", v, AllowedJellyfinPlayMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinPlayMethod) IsValid() bool {
	for _, existing := range AllowedJellyfinPlayMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayMethod value
func (v JellyfinPlayMethod) Ptr() *JellyfinPlayMethod {
	return &v
}

type NullableJellyfinPlayMethod struct {
	value *JellyfinPlayMethod
	isSet bool
}

func (v NullableJellyfinPlayMethod) Get() *JellyfinPlayMethod {
	return v.value
}

func (v *NullableJellyfinPlayMethod) Set(val *JellyfinPlayMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPlayMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPlayMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPlayMethod(val *JellyfinPlayMethod) *NullableJellyfinPlayMethod {
	return &NullableJellyfinPlayMethod{value: val, isSet: true}
}

func (v NullableJellyfinPlayMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPlayMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

