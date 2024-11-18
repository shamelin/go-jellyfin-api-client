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

// JellyfinSubtitleDeliveryMethod Delivery method to use during playback of a specific subtitle format.
type JellyfinSubtitleDeliveryMethod string

// List of SubtitleDeliveryMethod
const (
	ENCODE JellyfinSubtitleDeliveryMethod = "Encode"
	EMBED JellyfinSubtitleDeliveryMethod = "Embed"
	EXTERNAL JellyfinSubtitleDeliveryMethod = "External"
	HLS JellyfinSubtitleDeliveryMethod = "Hls"
	DROP JellyfinSubtitleDeliveryMethod = "Drop"
)

// All allowed values of JellyfinSubtitleDeliveryMethod enum
var AllowedJellyfinSubtitleDeliveryMethodEnumValues = []JellyfinSubtitleDeliveryMethod{
	"Encode",
	"Embed",
	"External",
	"Hls",
	"Drop",
}

func (v *JellyfinSubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinSubtitleDeliveryMethod(value)
	for _, existing := range AllowedJellyfinSubtitleDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinSubtitleDeliveryMethod", value)
}

// NewJellyfinSubtitleDeliveryMethodFromValue returns a pointer to a valid JellyfinSubtitleDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinSubtitleDeliveryMethodFromValue(v string) (*JellyfinSubtitleDeliveryMethod, error) {
	ev := JellyfinSubtitleDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinSubtitleDeliveryMethod: valid values are %v", v, AllowedJellyfinSubtitleDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinSubtitleDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedJellyfinSubtitleDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubtitleDeliveryMethod value
func (v JellyfinSubtitleDeliveryMethod) Ptr() *JellyfinSubtitleDeliveryMethod {
	return &v
}

type NullableJellyfinSubtitleDeliveryMethod struct {
	value *JellyfinSubtitleDeliveryMethod
	isSet bool
}

func (v NullableJellyfinSubtitleDeliveryMethod) Get() *JellyfinSubtitleDeliveryMethod {
	return v.value
}

func (v *NullableJellyfinSubtitleDeliveryMethod) Set(val *JellyfinSubtitleDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSubtitleDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSubtitleDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSubtitleDeliveryMethod(val *JellyfinSubtitleDeliveryMethod) *NullableJellyfinSubtitleDeliveryMethod {
	return &NullableJellyfinSubtitleDeliveryMethod{value: val, isSet: true}
}

func (v NullableJellyfinSubtitleDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

