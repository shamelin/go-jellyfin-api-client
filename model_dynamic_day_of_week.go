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

// DynamicDayOfWeek An enum that represents a day of the week, weekdays, weekends, or all days.
type DynamicDayOfWeek string

// List of DynamicDayOfWeek
const (
	SUNDAY DynamicDayOfWeek = "Sunday"
	MONDAY DynamicDayOfWeek = "Monday"
	TUESDAY DynamicDayOfWeek = "Tuesday"
	WEDNESDAY DynamicDayOfWeek = "Wednesday"
	THURSDAY DynamicDayOfWeek = "Thursday"
	FRIDAY DynamicDayOfWeek = "Friday"
	SATURDAY DynamicDayOfWeek = "Saturday"
	EVERYDAY DynamicDayOfWeek = "Everyday"
	WEEKDAY DynamicDayOfWeek = "Weekday"
	WEEKEND DynamicDayOfWeek = "Weekend"
)

// All allowed values of DynamicDayOfWeek enum
var AllowedDynamicDayOfWeekEnumValues = []DynamicDayOfWeek{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Everyday",
	"Weekday",
	"Weekend",
}

func (v *DynamicDayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DynamicDayOfWeek(value)
	for _, existing := range AllowedDynamicDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DynamicDayOfWeek", value)
}

// NewDynamicDayOfWeekFromValue returns a pointer to a valid DynamicDayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDynamicDayOfWeekFromValue(v string) (*DynamicDayOfWeek, error) {
	ev := DynamicDayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DynamicDayOfWeek: valid values are %v", v, AllowedDynamicDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DynamicDayOfWeek) IsValid() bool {
	for _, existing := range AllowedDynamicDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DynamicDayOfWeek value
func (v DynamicDayOfWeek) Ptr() *DynamicDayOfWeek {
	return &v
}

type NullableDynamicDayOfWeek struct {
	value *DynamicDayOfWeek
	isSet bool
}

func (v NullableDynamicDayOfWeek) Get() *DynamicDayOfWeek {
	return v.value
}

func (v *NullableDynamicDayOfWeek) Set(val *DynamicDayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicDayOfWeek(val *DynamicDayOfWeek) *NullableDynamicDayOfWeek {
	return &NullableDynamicDayOfWeek{value: val, isSet: true}
}

func (v NullableDynamicDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

