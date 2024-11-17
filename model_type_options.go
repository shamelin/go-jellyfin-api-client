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

// checks if the TypeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypeOptions{}

// TypeOptions struct for TypeOptions
type TypeOptions struct {
	Type NullableString `json:"Type,omitempty"`
	MetadataFetchers []string `json:"MetadataFetchers,omitempty"`
	MetadataFetcherOrder []string `json:"MetadataFetcherOrder,omitempty"`
	ImageFetchers []string `json:"ImageFetchers,omitempty"`
	ImageFetcherOrder []string `json:"ImageFetcherOrder,omitempty"`
	ImageOptions []ImageOption `json:"ImageOptions,omitempty"`
}

// NewTypeOptions instantiates a new TypeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeOptions() *TypeOptions {
	this := TypeOptions{}
	return &this
}

// NewTypeOptionsWithDefaults instantiates a new TypeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeOptionsWithDefaults() *TypeOptions {
	this := TypeOptions{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TypeOptions) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TypeOptions) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TypeOptions) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TypeOptions) UnsetType() {
	o.Type.Unset()
}

// GetMetadataFetchers returns the MetadataFetchers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetMetadataFetchers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MetadataFetchers
}

// GetMetadataFetchersOk returns a tuple with the MetadataFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetMetadataFetchersOk() ([]string, bool) {
	if o == nil || IsNil(o.MetadataFetchers) {
		return nil, false
	}
	return o.MetadataFetchers, true
}

// HasMetadataFetchers returns a boolean if a field has been set.
func (o *TypeOptions) HasMetadataFetchers() bool {
	if o != nil && !IsNil(o.MetadataFetchers) {
		return true
	}

	return false
}

// SetMetadataFetchers gets a reference to the given []string and assigns it to the MetadataFetchers field.
func (o *TypeOptions) SetMetadataFetchers(v []string) {
	o.MetadataFetchers = v
}

// GetMetadataFetcherOrder returns the MetadataFetcherOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetMetadataFetcherOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MetadataFetcherOrder
}

// GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetMetadataFetcherOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.MetadataFetcherOrder) {
		return nil, false
	}
	return o.MetadataFetcherOrder, true
}

// HasMetadataFetcherOrder returns a boolean if a field has been set.
func (o *TypeOptions) HasMetadataFetcherOrder() bool {
	if o != nil && !IsNil(o.MetadataFetcherOrder) {
		return true
	}

	return false
}

// SetMetadataFetcherOrder gets a reference to the given []string and assigns it to the MetadataFetcherOrder field.
func (o *TypeOptions) SetMetadataFetcherOrder(v []string) {
	o.MetadataFetcherOrder = v
}

// GetImageFetchers returns the ImageFetchers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetImageFetchers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ImageFetchers
}

// GetImageFetchersOk returns a tuple with the ImageFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetImageFetchersOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageFetchers) {
		return nil, false
	}
	return o.ImageFetchers, true
}

// HasImageFetchers returns a boolean if a field has been set.
func (o *TypeOptions) HasImageFetchers() bool {
	if o != nil && !IsNil(o.ImageFetchers) {
		return true
	}

	return false
}

// SetImageFetchers gets a reference to the given []string and assigns it to the ImageFetchers field.
func (o *TypeOptions) SetImageFetchers(v []string) {
	o.ImageFetchers = v
}

// GetImageFetcherOrder returns the ImageFetcherOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetImageFetcherOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ImageFetcherOrder
}

// GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetImageFetcherOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageFetcherOrder) {
		return nil, false
	}
	return o.ImageFetcherOrder, true
}

// HasImageFetcherOrder returns a boolean if a field has been set.
func (o *TypeOptions) HasImageFetcherOrder() bool {
	if o != nil && !IsNil(o.ImageFetcherOrder) {
		return true
	}

	return false
}

// SetImageFetcherOrder gets a reference to the given []string and assigns it to the ImageFetcherOrder field.
func (o *TypeOptions) SetImageFetcherOrder(v []string) {
	o.ImageFetcherOrder = v
}

// GetImageOptions returns the ImageOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeOptions) GetImageOptions() []ImageOption {
	if o == nil {
		var ret []ImageOption
		return ret
	}
	return o.ImageOptions
}

// GetImageOptionsOk returns a tuple with the ImageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeOptions) GetImageOptionsOk() ([]ImageOption, bool) {
	if o == nil || IsNil(o.ImageOptions) {
		return nil, false
	}
	return o.ImageOptions, true
}

// HasImageOptions returns a boolean if a field has been set.
func (o *TypeOptions) HasImageOptions() bool {
	if o != nil && !IsNil(o.ImageOptions) {
		return true
	}

	return false
}

// SetImageOptions gets a reference to the given []ImageOption and assigns it to the ImageOptions field.
func (o *TypeOptions) SetImageOptions(v []ImageOption) {
	o.ImageOptions = v
}

func (o TypeOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypeOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if o.MetadataFetchers != nil {
		toSerialize["MetadataFetchers"] = o.MetadataFetchers
	}
	if o.MetadataFetcherOrder != nil {
		toSerialize["MetadataFetcherOrder"] = o.MetadataFetcherOrder
	}
	if o.ImageFetchers != nil {
		toSerialize["ImageFetchers"] = o.ImageFetchers
	}
	if o.ImageFetcherOrder != nil {
		toSerialize["ImageFetcherOrder"] = o.ImageFetcherOrder
	}
	if o.ImageOptions != nil {
		toSerialize["ImageOptions"] = o.ImageOptions
	}
	return toSerialize, nil
}

type NullableTypeOptions struct {
	value *TypeOptions
	isSet bool
}

func (v NullableTypeOptions) Get() *TypeOptions {
	return v.value
}

func (v *NullableTypeOptions) Set(val *TypeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeOptions(val *TypeOptions) *NullableTypeOptions {
	return &NullableTypeOptions{value: val, isSet: true}
}

func (v NullableTypeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


