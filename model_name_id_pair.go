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

// checks if the NameIdPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameIdPair{}

// NameIdPair struct for NameIdPair
type NameIdPair struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the identifier.
	Id NullableString `json:"Id,omitempty"`
}

// NewNameIdPair instantiates a new NameIdPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameIdPair() *NameIdPair {
	this := NameIdPair{}
	return &this
}

// NewNameIdPairWithDefaults instantiates a new NameIdPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameIdPairWithDefaults() *NameIdPair {
	this := NameIdPair{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameIdPair) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameIdPair) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NameIdPair) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NameIdPair) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NameIdPair) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NameIdPair) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameIdPair) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameIdPair) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *NameIdPair) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *NameIdPair) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *NameIdPair) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *NameIdPair) UnsetId() {
	o.Id.Unset()
}

func (o NameIdPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameIdPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableNameIdPair struct {
	value *NameIdPair
	isSet bool
}

func (v NullableNameIdPair) Get() *NameIdPair {
	return v.value
}

func (v *NullableNameIdPair) Set(val *NameIdPair) {
	v.value = val
	v.isSet = true
}

func (v NullableNameIdPair) IsSet() bool {
	return v.isSet
}

func (v *NullableNameIdPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameIdPair(val *NameIdPair) *NullableNameIdPair {
	return &NullableNameIdPair{value: val, isSet: true}
}

func (v NullableNameIdPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameIdPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


