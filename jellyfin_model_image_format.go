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

// JellyfinImageFormat Enum ImageOutputFormat.
type JellyfinImageFormat string

// List of ImageFormat
const (
	JELLYFINIMAGEFORMAT_BMP JellyfinImageFormat = "Bmp"
	JELLYFINIMAGEFORMAT_GIF JellyfinImageFormat = "Gif"
	JELLYFINIMAGEFORMAT_JPG JellyfinImageFormat = "Jpg"
	JELLYFINIMAGEFORMAT_PNG JellyfinImageFormat = "Png"
	JELLYFINIMAGEFORMAT_WEBP JellyfinImageFormat = "Webp"
	JELLYFINIMAGEFORMAT_SVG JellyfinImageFormat = "Svg"
)

// All allowed values of JellyfinImageFormat enum
var AllowedJellyfinImageFormatEnumValues = []JellyfinImageFormat{
	"Bmp",
	"Gif",
	"Jpg",
	"Png",
	"Webp",
	"Svg",
}

func (v *JellyfinImageFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinImageFormat(value)
	for _, existing := range AllowedJellyfinImageFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinImageFormat", value)
}

// NewJellyfinImageFormatFromValue returns a pointer to a valid JellyfinImageFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinImageFormatFromValue(v string) (*JellyfinImageFormat, error) {
	ev := JellyfinImageFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinImageFormat: valid values are %v", v, AllowedJellyfinImageFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinImageFormat) IsValid() bool {
	for _, existing := range AllowedJellyfinImageFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageFormat value
func (v JellyfinImageFormat) Ptr() *JellyfinImageFormat {
	return &v
}

type NullableJellyfinImageFormat struct {
	value *JellyfinImageFormat
	isSet bool
}

func (v NullableJellyfinImageFormat) Get() *JellyfinImageFormat {
	return v.value
}

func (v *NullableJellyfinImageFormat) Set(val *JellyfinImageFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinImageFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinImageFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinImageFormat(val *JellyfinImageFormat) *NullableJellyfinImageFormat {
	return &NullableJellyfinImageFormat{value: val, isSet: true}
}

func (v NullableJellyfinImageFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinImageFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

