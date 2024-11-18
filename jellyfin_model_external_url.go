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

// checks if the JellyfinExternalUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinExternalUrl{}

// JellyfinExternalUrl struct for JellyfinExternalUrl
type JellyfinExternalUrl struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the type of the item.
	Url NullableString `json:"Url,omitempty"`
}

// NewJellyfinExternalUrl instantiates a new JellyfinExternalUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinExternalUrl() *JellyfinExternalUrl {
	this := JellyfinExternalUrl{}
	return &this
}

// NewJellyfinExternalUrlWithDefaults instantiates a new JellyfinExternalUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinExternalUrlWithDefaults() *JellyfinExternalUrl {
	this := JellyfinExternalUrl{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinExternalUrl) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinExternalUrl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinExternalUrl) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JellyfinExternalUrl) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JellyfinExternalUrl) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JellyfinExternalUrl) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinExternalUrl) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinExternalUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *JellyfinExternalUrl) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *JellyfinExternalUrl) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *JellyfinExternalUrl) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *JellyfinExternalUrl) UnsetUrl() {
	o.Url.Unset()
}

func (o JellyfinExternalUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinExternalUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["Url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinExternalUrl struct {
	value *JellyfinExternalUrl
	isSet bool
}

func (v NullableJellyfinExternalUrl) Get() *JellyfinExternalUrl {
	return v.value
}

func (v *NullableJellyfinExternalUrl) Set(val *JellyfinExternalUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinExternalUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinExternalUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinExternalUrl(val *JellyfinExternalUrl) *NullableJellyfinExternalUrl {
	return &NullableJellyfinExternalUrl{value: val, isSet: true}
}

func (v NullableJellyfinExternalUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinExternalUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

