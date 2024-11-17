/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"encoding/json"
	"fmt"
)

// MetadataRefreshMode the model 'MetadataRefreshMode'
type MetadataRefreshMode string

// List of MetadataRefreshMode
const (
	NONE MetadataRefreshMode = "None"
	VALIDATION_ONLY MetadataRefreshMode = "ValidationOnly"
	DEFAULT MetadataRefreshMode = "Default"
	FULL_REFRESH MetadataRefreshMode = "FullRefresh"
)

// All allowed values of MetadataRefreshMode enum
var AllowedMetadataRefreshModeEnumValues = []MetadataRefreshMode{
	"None",
	"ValidationOnly",
	"Default",
	"FullRefresh",
}

func (v *MetadataRefreshMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataRefreshMode(value)
	for _, existing := range AllowedMetadataRefreshModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataRefreshMode", value)
}

// NewMetadataRefreshModeFromValue returns a pointer to a valid MetadataRefreshMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataRefreshModeFromValue(v string) (*MetadataRefreshMode, error) {
	ev := MetadataRefreshMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataRefreshMode: valid values are %v", v, AllowedMetadataRefreshModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataRefreshMode) IsValid() bool {
	for _, existing := range AllowedMetadataRefreshModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataRefreshMode value
func (v MetadataRefreshMode) Ptr() *MetadataRefreshMode {
	return &v
}

type NullableMetadataRefreshMode struct {
	value *MetadataRefreshMode
	isSet bool
}

func (v NullableMetadataRefreshMode) Get() *MetadataRefreshMode {
	return v.value
}

func (v *NullableMetadataRefreshMode) Set(val *MetadataRefreshMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataRefreshMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataRefreshMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataRefreshMode(val *MetadataRefreshMode) *NullableMetadataRefreshMode {
	return &NullableMetadataRefreshMode{value: val, isSet: true}
}

func (v NullableMetadataRefreshMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataRefreshMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

