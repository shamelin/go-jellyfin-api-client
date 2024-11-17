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

// checks if the ImageProviderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageProviderInfo{}

// ImageProviderInfo Class ImageProviderInfo.
type ImageProviderInfo struct {
	// Gets the name.
	Name *string `json:"Name,omitempty"`
	// Gets the supported image types.
	SupportedImages []ImageType `json:"SupportedImages,omitempty"`
}

// NewImageProviderInfo instantiates a new ImageProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageProviderInfo() *ImageProviderInfo {
	this := ImageProviderInfo{}
	return &this
}

// NewImageProviderInfoWithDefaults instantiates a new ImageProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageProviderInfoWithDefaults() *ImageProviderInfo {
	this := ImageProviderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageProviderInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageProviderInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageProviderInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageProviderInfo) SetName(v string) {
	o.Name = &v
}

// GetSupportedImages returns the SupportedImages field value if set, zero value otherwise.
func (o *ImageProviderInfo) GetSupportedImages() []ImageType {
	if o == nil || IsNil(o.SupportedImages) {
		var ret []ImageType
		return ret
	}
	return o.SupportedImages
}

// GetSupportedImagesOk returns a tuple with the SupportedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageProviderInfo) GetSupportedImagesOk() ([]ImageType, bool) {
	if o == nil || IsNil(o.SupportedImages) {
		return nil, false
	}
	return o.SupportedImages, true
}

// HasSupportedImages returns a boolean if a field has been set.
func (o *ImageProviderInfo) HasSupportedImages() bool {
	if o != nil && !IsNil(o.SupportedImages) {
		return true
	}

	return false
}

// SetSupportedImages gets a reference to the given []ImageType and assigns it to the SupportedImages field.
func (o *ImageProviderInfo) SetSupportedImages(v []ImageType) {
	o.SupportedImages = v
}

func (o ImageProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageProviderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SupportedImages) {
		toSerialize["SupportedImages"] = o.SupportedImages
	}
	return toSerialize, nil
}

type NullableImageProviderInfo struct {
	value *ImageProviderInfo
	isSet bool
}

func (v NullableImageProviderInfo) Get() *ImageProviderInfo {
	return v.value
}

func (v *NullableImageProviderInfo) Set(val *ImageProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImageProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImageProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageProviderInfo(val *ImageProviderInfo) *NullableImageProviderInfo {
	return &NullableImageProviderInfo{value: val, isSet: true}
}

func (v NullableImageProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


