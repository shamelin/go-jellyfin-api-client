/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MetadataField Enum MetadataFields.
type MetadataField string

// List of MetadataField
const (
	CAST MetadataField = "Cast"
	GENRES MetadataField = "Genres"
	PRODUCTION_LOCATIONS MetadataField = "ProductionLocations"
	STUDIOS MetadataField = "Studios"
	TAGS MetadataField = "Tags"
	NAME MetadataField = "Name"
	OVERVIEW MetadataField = "Overview"
	RUNTIME MetadataField = "Runtime"
	OFFICIAL_RATING MetadataField = "OfficialRating"
)

// All allowed values of MetadataField enum
var AllowedMetadataFieldEnumValues = []MetadataField{
	"Cast",
	"Genres",
	"ProductionLocations",
	"Studios",
	"Tags",
	"Name",
	"Overview",
	"Runtime",
	"OfficialRating",
}

func (v *MetadataField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataField(value)
	for _, existing := range AllowedMetadataFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataField", value)
}

// NewMetadataFieldFromValue returns a pointer to a valid MetadataField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataFieldFromValue(v string) (*MetadataField, error) {
	ev := MetadataField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataField: valid values are %v", v, AllowedMetadataFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataField) IsValid() bool {
	for _, existing := range AllowedMetadataFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataField value
func (v MetadataField) Ptr() *MetadataField {
	return &v
}

type NullableMetadataField struct {
	value *MetadataField
	isSet bool
}

func (v NullableMetadataField) Get() *MetadataField {
	return v.value
}

func (v *NullableMetadataField) Set(val *MetadataField) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataField) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataField(val *MetadataField) *NullableMetadataField {
	return &NullableMetadataField{value: val, isSet: true}
}

func (v NullableMetadataField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

