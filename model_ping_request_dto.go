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

// checks if the PingRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingRequestDto{}

// PingRequestDto Class PingRequestDto.
type PingRequestDto struct {
	// Gets or sets the ping time.
	Ping *int64 `json:"Ping,omitempty"`
}

// NewPingRequestDto instantiates a new PingRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingRequestDto() *PingRequestDto {
	this := PingRequestDto{}
	return &this
}

// NewPingRequestDtoWithDefaults instantiates a new PingRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingRequestDtoWithDefaults() *PingRequestDto {
	this := PingRequestDto{}
	return &this
}

// GetPing returns the Ping field value if set, zero value otherwise.
func (o *PingRequestDto) GetPing() int64 {
	if o == nil || IsNil(o.Ping) {
		var ret int64
		return ret
	}
	return *o.Ping
}

// GetPingOk returns a tuple with the Ping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingRequestDto) GetPingOk() (*int64, bool) {
	if o == nil || IsNil(o.Ping) {
		return nil, false
	}
	return o.Ping, true
}

// HasPing returns a boolean if a field has been set.
func (o *PingRequestDto) HasPing() bool {
	if o != nil && !IsNil(o.Ping) {
		return true
	}

	return false
}

// SetPing gets a reference to the given int64 and assigns it to the Ping field.
func (o *PingRequestDto) SetPing(v int64) {
	o.Ping = &v
}

func (o PingRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ping) {
		toSerialize["Ping"] = o.Ping
	}
	return toSerialize, nil
}

type NullablePingRequestDto struct {
	value *PingRequestDto
	isSet bool
}

func (v NullablePingRequestDto) Get() *PingRequestDto {
	return v.value
}

func (v *NullablePingRequestDto) Set(val *PingRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePingRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePingRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingRequestDto(val *PingRequestDto) *NullablePingRequestDto {
	return &NullablePingRequestDto{value: val, isSet: true}
}

func (v NullablePingRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


