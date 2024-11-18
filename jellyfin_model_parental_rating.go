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

// checks if the JellyfinParentalRating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinParentalRating{}

// JellyfinParentalRating Class ParentalRating.
type JellyfinParentalRating struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the value.
	Value NullableInt32 `json:"Value,omitempty"`
}

// NewJellyfinParentalRating instantiates a new JellyfinParentalRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinParentalRating() *JellyfinParentalRating {
	this := JellyfinParentalRating{}
	return &this
}

// NewJellyfinParentalRatingWithDefaults instantiates a new JellyfinParentalRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinParentalRatingWithDefaults() *JellyfinParentalRating {
	this := JellyfinParentalRating{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinParentalRating) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinParentalRating) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinParentalRating) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JellyfinParentalRating) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JellyfinParentalRating) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JellyfinParentalRating) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinParentalRating) GetValue() int32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinParentalRating) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *JellyfinParentalRating) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt32 and assigns it to the Value field.
func (o *JellyfinParentalRating) SetValue(v int32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *JellyfinParentalRating) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *JellyfinParentalRating) UnsetValue() {
	o.Value.Unset()
}

func (o JellyfinParentalRating) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinParentalRating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinParentalRating struct {
	value *JellyfinParentalRating
	isSet bool
}

func (v NullableJellyfinParentalRating) Get() *JellyfinParentalRating {
	return v.value
}

func (v *NullableJellyfinParentalRating) Set(val *JellyfinParentalRating) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinParentalRating) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinParentalRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinParentalRating(val *JellyfinParentalRating) *NullableJellyfinParentalRating {
	return &NullableJellyfinParentalRating{value: val, isSet: true}
}

func (v NullableJellyfinParentalRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinParentalRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

