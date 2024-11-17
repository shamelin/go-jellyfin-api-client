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

// CollectionType Collection type.
type CollectionType string

// List of CollectionType
const (
	UNKNOWN CollectionType = "unknown"
	MOVIES CollectionType = "movies"
	TVSHOWS CollectionType = "tvshows"
	MUSIC CollectionType = "music"
	MUSICVIDEOS CollectionType = "musicvideos"
	TRAILERS CollectionType = "trailers"
	HOMEVIDEOS CollectionType = "homevideos"
	BOXSETS CollectionType = "boxsets"
	BOOKS CollectionType = "books"
	PHOTOS CollectionType = "photos"
	LIVETV CollectionType = "livetv"
	PLAYLISTS CollectionType = "playlists"
	FOLDERS CollectionType = "folders"
)

// All allowed values of CollectionType enum
var AllowedCollectionTypeEnumValues = []CollectionType{
	"unknown",
	"movies",
	"tvshows",
	"music",
	"musicvideos",
	"trailers",
	"homevideos",
	"boxsets",
	"books",
	"photos",
	"livetv",
	"playlists",
	"folders",
}

func (v *CollectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionType(value)
	for _, existing := range AllowedCollectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionType", value)
}

// NewCollectionTypeFromValue returns a pointer to a valid CollectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionTypeFromValue(v string) (*CollectionType, error) {
	ev := CollectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionType: valid values are %v", v, AllowedCollectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionType) IsValid() bool {
	for _, existing := range AllowedCollectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectionType value
func (v CollectionType) Ptr() *CollectionType {
	return &v
}

type NullableCollectionType struct {
	value *CollectionType
	isSet bool
}

func (v NullableCollectionType) Get() *CollectionType {
	return v.value
}

func (v *NullableCollectionType) Set(val *CollectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionType(val *CollectionType) *NullableCollectionType {
	return &NullableCollectionType{value: val, isSet: true}
}

func (v NullableCollectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
