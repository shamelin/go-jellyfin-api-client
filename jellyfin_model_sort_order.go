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

// JellyfinSortOrder An enum representing the sorting order.
type JellyfinSortOrder string

// List of SortOrder
const (
	ASCENDING JellyfinSortOrder = "Ascending"
	DESCENDING JellyfinSortOrder = "Descending"
)

// All allowed values of JellyfinSortOrder enum
var AllowedJellyfinSortOrderEnumValues = []JellyfinSortOrder{
	"Ascending",
	"Descending",
}

func (v *JellyfinSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinSortOrder(value)
	for _, existing := range AllowedJellyfinSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinSortOrder", value)
}

// NewJellyfinSortOrderFromValue returns a pointer to a valid JellyfinSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinSortOrderFromValue(v string) (*JellyfinSortOrder, error) {
	ev := JellyfinSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinSortOrder: valid values are %v", v, AllowedJellyfinSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinSortOrder) IsValid() bool {
	for _, existing := range AllowedJellyfinSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrder value
func (v JellyfinSortOrder) Ptr() *JellyfinSortOrder {
	return &v
}

type NullableJellyfinSortOrder struct {
	value *JellyfinSortOrder
	isSet bool
}

func (v NullableJellyfinSortOrder) Get() *JellyfinSortOrder {
	return v.value
}

func (v *NullableJellyfinSortOrder) Set(val *JellyfinSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSortOrder(val *JellyfinSortOrder) *NullableJellyfinSortOrder {
	return &NullableJellyfinSortOrder{value: val, isSet: true}
}

func (v NullableJellyfinSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

