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

// checks if the JellyfinValidatePathDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinValidatePathDto{}

// JellyfinValidatePathDto Validate path object.
type JellyfinValidatePathDto struct {
	// Gets or sets a value indicating whether validate if path is writable.
	ValidateWritable *bool `json:"ValidateWritable,omitempty"`
	// Gets or sets the path.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets is path file.
	IsFile NullableBool `json:"IsFile,omitempty"`
}

// NewJellyfinValidatePathDto instantiates a new JellyfinValidatePathDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinValidatePathDto() *JellyfinValidatePathDto {
	this := JellyfinValidatePathDto{}
	return &this
}

// NewJellyfinValidatePathDtoWithDefaults instantiates a new JellyfinValidatePathDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinValidatePathDtoWithDefaults() *JellyfinValidatePathDto {
	this := JellyfinValidatePathDto{}
	return &this
}

// GetValidateWritable returns the ValidateWritable field value if set, zero value otherwise.
func (o *JellyfinValidatePathDto) GetValidateWritable() bool {
	if o == nil || IsNil(o.ValidateWritable) {
		var ret bool
		return ret
	}
	return *o.ValidateWritable
}

// GetValidateWritableOk returns a tuple with the ValidateWritable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinValidatePathDto) GetValidateWritableOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateWritable) {
		return nil, false
	}
	return o.ValidateWritable, true
}

// HasValidateWritable returns a boolean if a field has been set.
func (o *JellyfinValidatePathDto) HasValidateWritable() bool {
	if o != nil && !IsNil(o.ValidateWritable) {
		return true
	}

	return false
}

// SetValidateWritable gets a reference to the given bool and assigns it to the ValidateWritable field.
func (o *JellyfinValidatePathDto) SetValidateWritable(v bool) {
	o.ValidateWritable = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinValidatePathDto) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinValidatePathDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *JellyfinValidatePathDto) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *JellyfinValidatePathDto) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *JellyfinValidatePathDto) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *JellyfinValidatePathDto) UnsetPath() {
	o.Path.Unset()
}

// GetIsFile returns the IsFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinValidatePathDto) GetIsFile() bool {
	if o == nil || IsNil(o.IsFile.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFile.Get()
}

// GetIsFileOk returns a tuple with the IsFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinValidatePathDto) GetIsFileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFile.Get(), o.IsFile.IsSet()
}

// HasIsFile returns a boolean if a field has been set.
func (o *JellyfinValidatePathDto) HasIsFile() bool {
	if o != nil && o.IsFile.IsSet() {
		return true
	}

	return false
}

// SetIsFile gets a reference to the given NullableBool and assigns it to the IsFile field.
func (o *JellyfinValidatePathDto) SetIsFile(v bool) {
	o.IsFile.Set(&v)
}
// SetIsFileNil sets the value for IsFile to be an explicit nil
func (o *JellyfinValidatePathDto) SetIsFileNil() {
	o.IsFile.Set(nil)
}

// UnsetIsFile ensures that no value is present for IsFile, not even an explicit nil
func (o *JellyfinValidatePathDto) UnsetIsFile() {
	o.IsFile.Unset()
}

func (o JellyfinValidatePathDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinValidatePathDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValidateWritable) {
		toSerialize["ValidateWritable"] = o.ValidateWritable
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.IsFile.IsSet() {
		toSerialize["IsFile"] = o.IsFile.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinValidatePathDto struct {
	value *JellyfinValidatePathDto
	isSet bool
}

func (v NullableJellyfinValidatePathDto) Get() *JellyfinValidatePathDto {
	return v.value
}

func (v *NullableJellyfinValidatePathDto) Set(val *JellyfinValidatePathDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinValidatePathDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinValidatePathDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinValidatePathDto(val *JellyfinValidatePathDto) *NullableJellyfinValidatePathDto {
	return &NullableJellyfinValidatePathDto{value: val, isSet: true}
}

func (v NullableJellyfinValidatePathDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinValidatePathDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


