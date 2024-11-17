/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GuideInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuideInfo{}

// GuideInfo struct for GuideInfo
type GuideInfo struct {
	// Gets or sets the start date.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Gets or sets the end date.
	EndDate *time.Time `json:"EndDate,omitempty"`
}

// NewGuideInfo instantiates a new GuideInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuideInfo() *GuideInfo {
	this := GuideInfo{}
	return &this
}

// NewGuideInfoWithDefaults instantiates a new GuideInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuideInfoWithDefaults() *GuideInfo {
	this := GuideInfo{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GuideInfo) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideInfo) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GuideInfo) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *GuideInfo) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GuideInfo) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuideInfo) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GuideInfo) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GuideInfo) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o GuideInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuideInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableGuideInfo struct {
	value *GuideInfo
	isSet bool
}

func (v NullableGuideInfo) Get() *GuideInfo {
	return v.value
}

func (v *NullableGuideInfo) Set(val *GuideInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGuideInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGuideInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuideInfo(val *GuideInfo) *NullableGuideInfo {
	return &NullableGuideInfo{value: val, isSet: true}
}

func (v NullableGuideInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuideInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


