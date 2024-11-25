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

// JellyfinSyncPlayUserAccessType Enum SyncPlayUserAccessType.
type JellyfinSyncPlayUserAccessType string

// List of SyncPlayUserAccessType
const (
	JELLYFINSYNCPLAYUSERACCESSTYPE_CREATE_AND_JOIN_GROUPS JellyfinSyncPlayUserAccessType = "CreateAndJoinGroups"
	JELLYFINSYNCPLAYUSERACCESSTYPE_JOIN_GROUPS JellyfinSyncPlayUserAccessType = "JoinGroups"
	JELLYFINSYNCPLAYUSERACCESSTYPE_NONE JellyfinSyncPlayUserAccessType = "None"
)

// All allowed values of JellyfinSyncPlayUserAccessType enum
var AllowedJellyfinSyncPlayUserAccessTypeEnumValues = []JellyfinSyncPlayUserAccessType{
	"CreateAndJoinGroups",
	"JoinGroups",
	"None",
}

func (v *JellyfinSyncPlayUserAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JellyfinSyncPlayUserAccessType(value)
	for _, existing := range AllowedJellyfinSyncPlayUserAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JellyfinSyncPlayUserAccessType", value)
}

// NewJellyfinSyncPlayUserAccessTypeFromValue returns a pointer to a valid JellyfinSyncPlayUserAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJellyfinSyncPlayUserAccessTypeFromValue(v string) (*JellyfinSyncPlayUserAccessType, error) {
	ev := JellyfinSyncPlayUserAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JellyfinSyncPlayUserAccessType: valid values are %v", v, AllowedJellyfinSyncPlayUserAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JellyfinSyncPlayUserAccessType) IsValid() bool {
	for _, existing := range AllowedJellyfinSyncPlayUserAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyncPlayUserAccessType value
func (v JellyfinSyncPlayUserAccessType) Ptr() *JellyfinSyncPlayUserAccessType {
	return &v
}

type NullableJellyfinSyncPlayUserAccessType struct {
	value *JellyfinSyncPlayUserAccessType
	isSet bool
}

func (v NullableJellyfinSyncPlayUserAccessType) Get() *JellyfinSyncPlayUserAccessType {
	return v.value
}

func (v *NullableJellyfinSyncPlayUserAccessType) Set(val *JellyfinSyncPlayUserAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSyncPlayUserAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSyncPlayUserAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSyncPlayUserAccessType(val *JellyfinSyncPlayUserAccessType) *NullableJellyfinSyncPlayUserAccessType {
	return &NullableJellyfinSyncPlayUserAccessType{value: val, isSet: true}
}

func (v NullableJellyfinSyncPlayUserAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSyncPlayUserAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

