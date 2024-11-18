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

// checks if the JellyfinCodecProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinCodecProfile{}

// JellyfinCodecProfile Defines the MediaBrowser.Model.Dlna.CodecProfile.
type JellyfinCodecProfile struct {
	// Gets or sets the MediaBrowser.Model.Dlna.CodecType which this container must meet.
	Type *JellyfinCodecType `json:"Type,omitempty"`
	// Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition which this profile must meet.
	Conditions []JellyfinProfileCondition `json:"Conditions,omitempty"`
	// Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition to apply if this profile is met.
	ApplyConditions []JellyfinProfileCondition `json:"ApplyConditions,omitempty"`
	// Gets or sets the codec(s) that this profile applies to.
	Codec NullableString `json:"Codec,omitempty"`
	// Gets or sets the container(s) which this profile will be applied to.
	Container NullableString `json:"Container,omitempty"`
	// Gets or sets the sub-container(s) which this profile will be applied to.
	SubContainer NullableString `json:"SubContainer,omitempty"`
}

// NewJellyfinCodecProfile instantiates a new JellyfinCodecProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinCodecProfile() *JellyfinCodecProfile {
	this := JellyfinCodecProfile{}
	return &this
}

// NewJellyfinCodecProfileWithDefaults instantiates a new JellyfinCodecProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinCodecProfileWithDefaults() *JellyfinCodecProfile {
	this := JellyfinCodecProfile{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JellyfinCodecProfile) GetType() JellyfinCodecType {
	if o == nil || IsNil(o.Type) {
		var ret JellyfinCodecType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCodecProfile) GetTypeOk() (*JellyfinCodecType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given JellyfinCodecType and assigns it to the Type field.
func (o *JellyfinCodecProfile) SetType(v JellyfinCodecType) {
	o.Type = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *JellyfinCodecProfile) GetConditions() []JellyfinProfileCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []JellyfinProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCodecProfile) GetConditionsOk() ([]JellyfinProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []JellyfinProfileCondition and assigns it to the Conditions field.
func (o *JellyfinCodecProfile) SetConditions(v []JellyfinProfileCondition) {
	o.Conditions = v
}

// GetApplyConditions returns the ApplyConditions field value if set, zero value otherwise.
func (o *JellyfinCodecProfile) GetApplyConditions() []JellyfinProfileCondition {
	if o == nil || IsNil(o.ApplyConditions) {
		var ret []JellyfinProfileCondition
		return ret
	}
	return o.ApplyConditions
}

// GetApplyConditionsOk returns a tuple with the ApplyConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCodecProfile) GetApplyConditionsOk() ([]JellyfinProfileCondition, bool) {
	if o == nil || IsNil(o.ApplyConditions) {
		return nil, false
	}
	return o.ApplyConditions, true
}

// HasApplyConditions returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasApplyConditions() bool {
	if o != nil && !IsNil(o.ApplyConditions) {
		return true
	}

	return false
}

// SetApplyConditions gets a reference to the given []JellyfinProfileCondition and assigns it to the ApplyConditions field.
func (o *JellyfinCodecProfile) SetApplyConditions(v []JellyfinProfileCondition) {
	o.ApplyConditions = v
}

// GetCodec returns the Codec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinCodecProfile) GetCodec() string {
	if o == nil || IsNil(o.Codec.Get()) {
		var ret string
		return ret
	}
	return *o.Codec.Get()
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinCodecProfile) GetCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codec.Get(), o.Codec.IsSet()
}

// HasCodec returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasCodec() bool {
	if o != nil && o.Codec.IsSet() {
		return true
	}

	return false
}

// SetCodec gets a reference to the given NullableString and assigns it to the Codec field.
func (o *JellyfinCodecProfile) SetCodec(v string) {
	o.Codec.Set(&v)
}
// SetCodecNil sets the value for Codec to be an explicit nil
func (o *JellyfinCodecProfile) SetCodecNil() {
	o.Codec.Set(nil)
}

// UnsetCodec ensures that no value is present for Codec, not even an explicit nil
func (o *JellyfinCodecProfile) UnsetCodec() {
	o.Codec.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinCodecProfile) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinCodecProfile) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *JellyfinCodecProfile) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *JellyfinCodecProfile) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *JellyfinCodecProfile) UnsetContainer() {
	o.Container.Unset()
}

// GetSubContainer returns the SubContainer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinCodecProfile) GetSubContainer() string {
	if o == nil || IsNil(o.SubContainer.Get()) {
		var ret string
		return ret
	}
	return *o.SubContainer.Get()
}

// GetSubContainerOk returns a tuple with the SubContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinCodecProfile) GetSubContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubContainer.Get(), o.SubContainer.IsSet()
}

// HasSubContainer returns a boolean if a field has been set.
func (o *JellyfinCodecProfile) HasSubContainer() bool {
	if o != nil && o.SubContainer.IsSet() {
		return true
	}

	return false
}

// SetSubContainer gets a reference to the given NullableString and assigns it to the SubContainer field.
func (o *JellyfinCodecProfile) SetSubContainer(v string) {
	o.SubContainer.Set(&v)
}
// SetSubContainerNil sets the value for SubContainer to be an explicit nil
func (o *JellyfinCodecProfile) SetSubContainerNil() {
	o.SubContainer.Set(nil)
}

// UnsetSubContainer ensures that no value is present for SubContainer, not even an explicit nil
func (o *JellyfinCodecProfile) UnsetSubContainer() {
	o.SubContainer.Unset()
}

func (o JellyfinCodecProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinCodecProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	if !IsNil(o.ApplyConditions) {
		toSerialize["ApplyConditions"] = o.ApplyConditions
	}
	if o.Codec.IsSet() {
		toSerialize["Codec"] = o.Codec.Get()
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	if o.SubContainer.IsSet() {
		toSerialize["SubContainer"] = o.SubContainer.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinCodecProfile struct {
	value *JellyfinCodecProfile
	isSet bool
}

func (v NullableJellyfinCodecProfile) Get() *JellyfinCodecProfile {
	return v.value
}

func (v *NullableJellyfinCodecProfile) Set(val *JellyfinCodecProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinCodecProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinCodecProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinCodecProfile(val *JellyfinCodecProfile) *NullableJellyfinCodecProfile {
	return &NullableJellyfinCodecProfile{value: val, isSet: true}
}

func (v NullableJellyfinCodecProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinCodecProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

