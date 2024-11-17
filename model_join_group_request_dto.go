/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the JoinGroupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoinGroupRequestDto{}

// JoinGroupRequestDto Class JoinGroupRequestDto.
type JoinGroupRequestDto struct {
	// Gets or sets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
}

// NewJoinGroupRequestDto instantiates a new JoinGroupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinGroupRequestDto() *JoinGroupRequestDto {
	this := JoinGroupRequestDto{}
	return &this
}

// NewJoinGroupRequestDtoWithDefaults instantiates a new JoinGroupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinGroupRequestDtoWithDefaults() *JoinGroupRequestDto {
	this := JoinGroupRequestDto{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *JoinGroupRequestDto) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinGroupRequestDto) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *JoinGroupRequestDto) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *JoinGroupRequestDto) SetGroupId(v string) {
	o.GroupId = &v
}

func (o JoinGroupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoinGroupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	return toSerialize, nil
}

type NullableJoinGroupRequestDto struct {
	value *JoinGroupRequestDto
	isSet bool
}

func (v NullableJoinGroupRequestDto) Get() *JoinGroupRequestDto {
	return v.value
}

func (v *NullableJoinGroupRequestDto) Set(val *JoinGroupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinGroupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinGroupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinGroupRequestDto(val *JoinGroupRequestDto) *NullableJoinGroupRequestDto {
	return &NullableJoinGroupRequestDto{value: val, isSet: true}
}

func (v NullableJoinGroupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinGroupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


