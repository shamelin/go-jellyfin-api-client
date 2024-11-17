/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"encoding/json"
)

// checks if the SubtitleProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubtitleProfile{}

// SubtitleProfile A class for subtitle profile information.
type SubtitleProfile struct {
	// Gets or sets the format.
	Format NullableString `json:"Format,omitempty"`
	// Gets or sets the delivery method.
	Method *SubtitleDeliveryMethod `json:"Method,omitempty"`
	// Gets or sets the DIDL mode.
	DidlMode NullableString `json:"DidlMode,omitempty"`
	// Gets or sets the language.
	Language NullableString `json:"Language,omitempty"`
	// Gets or sets the container.
	Container NullableString `json:"Container,omitempty"`
}

// NewSubtitleProfile instantiates a new SubtitleProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubtitleProfile() *SubtitleProfile {
	this := SubtitleProfile{}
	return &this
}

// NewSubtitleProfileWithDefaults instantiates a new SubtitleProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubtitleProfileWithDefaults() *SubtitleProfile {
	this := SubtitleProfile{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleProfile) GetFormat() string {
	if o == nil || IsNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleProfile) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *SubtitleProfile) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *SubtitleProfile) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *SubtitleProfile) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *SubtitleProfile) UnsetFormat() {
	o.Format.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SubtitleProfile) GetMethod() SubtitleDeliveryMethod {
	if o == nil || IsNil(o.Method) {
		var ret SubtitleDeliveryMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleProfile) GetMethodOk() (*SubtitleDeliveryMethod, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SubtitleProfile) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given SubtitleDeliveryMethod and assigns it to the Method field.
func (o *SubtitleProfile) SetMethod(v SubtitleDeliveryMethod) {
	o.Method = &v
}

// GetDidlMode returns the DidlMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleProfile) GetDidlMode() string {
	if o == nil || IsNil(o.DidlMode.Get()) {
		var ret string
		return ret
	}
	return *o.DidlMode.Get()
}

// GetDidlModeOk returns a tuple with the DidlMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleProfile) GetDidlModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DidlMode.Get(), o.DidlMode.IsSet()
}

// HasDidlMode returns a boolean if a field has been set.
func (o *SubtitleProfile) HasDidlMode() bool {
	if o != nil && o.DidlMode.IsSet() {
		return true
	}

	return false
}

// SetDidlMode gets a reference to the given NullableString and assigns it to the DidlMode field.
func (o *SubtitleProfile) SetDidlMode(v string) {
	o.DidlMode.Set(&v)
}
// SetDidlModeNil sets the value for DidlMode to be an explicit nil
func (o *SubtitleProfile) SetDidlModeNil() {
	o.DidlMode.Set(nil)
}

// UnsetDidlMode ensures that no value is present for DidlMode, not even an explicit nil
func (o *SubtitleProfile) UnsetDidlMode() {
	o.DidlMode.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleProfile) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleProfile) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *SubtitleProfile) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *SubtitleProfile) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *SubtitleProfile) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *SubtitleProfile) UnsetLanguage() {
	o.Language.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleProfile) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleProfile) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *SubtitleProfile) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *SubtitleProfile) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *SubtitleProfile) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *SubtitleProfile) UnsetContainer() {
	o.Container.Unset()
}

func (o SubtitleProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubtitleProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Format.IsSet() {
		toSerialize["Format"] = o.Format.Get()
	}
	if !IsNil(o.Method) {
		toSerialize["Method"] = o.Method
	}
	if o.DidlMode.IsSet() {
		toSerialize["DidlMode"] = o.DidlMode.Get()
	}
	if o.Language.IsSet() {
		toSerialize["Language"] = o.Language.Get()
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	return toSerialize, nil
}

type NullableSubtitleProfile struct {
	value *SubtitleProfile
	isSet bool
}

func (v NullableSubtitleProfile) Get() *SubtitleProfile {
	return v.value
}

func (v *NullableSubtitleProfile) Set(val *SubtitleProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableSubtitleProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableSubtitleProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubtitleProfile(val *SubtitleProfile) *NullableSubtitleProfile {
	return &NullableSubtitleProfile{value: val, isSet: true}
}

func (v NullableSubtitleProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubtitleProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


