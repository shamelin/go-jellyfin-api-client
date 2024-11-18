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

// checks if the JellyfinAddVirtualFolderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinAddVirtualFolderDto{}

// JellyfinAddVirtualFolderDto Add virtual folder dto.
type JellyfinAddVirtualFolderDto struct {
	// Gets or sets library options.
	LibraryOptions NullableJellyfinLibraryOptions `json:"LibraryOptions,omitempty"`
}

// NewJellyfinAddVirtualFolderDto instantiates a new JellyfinAddVirtualFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinAddVirtualFolderDto() *JellyfinAddVirtualFolderDto {
	this := JellyfinAddVirtualFolderDto{}
	return &this
}

// NewJellyfinAddVirtualFolderDtoWithDefaults instantiates a new JellyfinAddVirtualFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinAddVirtualFolderDtoWithDefaults() *JellyfinAddVirtualFolderDto {
	this := JellyfinAddVirtualFolderDto{}
	return &this
}

// GetLibraryOptions returns the LibraryOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinAddVirtualFolderDto) GetLibraryOptions() JellyfinLibraryOptions {
	if o == nil || IsNil(o.LibraryOptions.Get()) {
		var ret JellyfinLibraryOptions
		return ret
	}
	return *o.LibraryOptions.Get()
}

// GetLibraryOptionsOk returns a tuple with the LibraryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinAddVirtualFolderDto) GetLibraryOptionsOk() (*JellyfinLibraryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibraryOptions.Get(), o.LibraryOptions.IsSet()
}

// HasLibraryOptions returns a boolean if a field has been set.
func (o *JellyfinAddVirtualFolderDto) HasLibraryOptions() bool {
	if o != nil && o.LibraryOptions.IsSet() {
		return true
	}

	return false
}

// SetLibraryOptions gets a reference to the given NullableJellyfinLibraryOptions and assigns it to the LibraryOptions field.
func (o *JellyfinAddVirtualFolderDto) SetLibraryOptions(v JellyfinLibraryOptions) {
	o.LibraryOptions.Set(&v)
}
// SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil
func (o *JellyfinAddVirtualFolderDto) SetLibraryOptionsNil() {
	o.LibraryOptions.Set(nil)
}

// UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil
func (o *JellyfinAddVirtualFolderDto) UnsetLibraryOptions() {
	o.LibraryOptions.Unset()
}

func (o JellyfinAddVirtualFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinAddVirtualFolderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryOptions.IsSet() {
		toSerialize["LibraryOptions"] = o.LibraryOptions.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinAddVirtualFolderDto struct {
	value *JellyfinAddVirtualFolderDto
	isSet bool
}

func (v NullableJellyfinAddVirtualFolderDto) Get() *JellyfinAddVirtualFolderDto {
	return v.value
}

func (v *NullableJellyfinAddVirtualFolderDto) Set(val *JellyfinAddVirtualFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinAddVirtualFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinAddVirtualFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinAddVirtualFolderDto(val *JellyfinAddVirtualFolderDto) *NullableJellyfinAddVirtualFolderDto {
	return &NullableJellyfinAddVirtualFolderDto{value: val, isSet: true}
}

func (v NullableJellyfinAddVirtualFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinAddVirtualFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


