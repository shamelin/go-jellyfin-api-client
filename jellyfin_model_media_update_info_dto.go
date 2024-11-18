/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
)

// checks if the JellyfinMediaUpdateInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinMediaUpdateInfoDto{}

// JellyfinMediaUpdateInfoDto Media Update Info Dto.
type JellyfinMediaUpdateInfoDto struct {
	// Gets or sets the list of updates.
	Updates []JellyfinMediaUpdateInfoPathDto `json:"Updates,omitempty"`
}

// NewJellyfinMediaUpdateInfoDto instantiates a new JellyfinMediaUpdateInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinMediaUpdateInfoDto() *JellyfinMediaUpdateInfoDto {
	this := JellyfinMediaUpdateInfoDto{}
	return &this
}

// NewJellyfinMediaUpdateInfoDtoWithDefaults instantiates a new JellyfinMediaUpdateInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinMediaUpdateInfoDtoWithDefaults() *JellyfinMediaUpdateInfoDto {
	this := JellyfinMediaUpdateInfoDto{}
	return &this
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *JellyfinMediaUpdateInfoDto) GetUpdates() []JellyfinMediaUpdateInfoPathDto {
	if o == nil || IsNil(o.Updates) {
		var ret []JellyfinMediaUpdateInfoPathDto
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinMediaUpdateInfoDto) GetUpdatesOk() ([]JellyfinMediaUpdateInfoPathDto, bool) {
	if o == nil || IsNil(o.Updates) {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *JellyfinMediaUpdateInfoDto) HasUpdates() bool {
	if o != nil && !IsNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given []JellyfinMediaUpdateInfoPathDto and assigns it to the Updates field.
func (o *JellyfinMediaUpdateInfoDto) SetUpdates(v []JellyfinMediaUpdateInfoPathDto) {
	o.Updates = v
}

func (o JellyfinMediaUpdateInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinMediaUpdateInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Updates) {
		toSerialize["Updates"] = o.Updates
	}
	return toSerialize, nil
}

type NullableJellyfinMediaUpdateInfoDto struct {
	value *JellyfinMediaUpdateInfoDto
	isSet bool
}

func (v NullableJellyfinMediaUpdateInfoDto) Get() *JellyfinMediaUpdateInfoDto {
	return v.value
}

func (v *NullableJellyfinMediaUpdateInfoDto) Set(val *JellyfinMediaUpdateInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMediaUpdateInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMediaUpdateInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMediaUpdateInfoDto(val *JellyfinMediaUpdateInfoDto) *NullableJellyfinMediaUpdateInfoDto {
	return &NullableJellyfinMediaUpdateInfoDto{value: val, isSet: true}
}

func (v NullableJellyfinMediaUpdateInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMediaUpdateInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

