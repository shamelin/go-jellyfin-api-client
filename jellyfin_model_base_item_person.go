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

// checks if the JellyfinBaseItemPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinBaseItemPerson{}

// JellyfinBaseItemPerson This is used by the api to get information about a Person within a BaseItem.
type JellyfinBaseItemPerson struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the identifier.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the role.
	Role NullableString `json:"Role,omitempty"`
	// The person kind.
	Type *JellyfinPersonKind `json:"Type,omitempty"`
	// Gets or sets the primary image tag.
	PrimaryImageTag NullableString `json:"PrimaryImageTag,omitempty"`
	ImageBlurHashes NullableJellyfinBaseItemPersonImageBlurHashes `json:"ImageBlurHashes,omitempty"`
}

// NewJellyfinBaseItemPerson instantiates a new JellyfinBaseItemPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinBaseItemPerson() *JellyfinBaseItemPerson {
	this := JellyfinBaseItemPerson{}
	return &this
}

// NewJellyfinBaseItemPersonWithDefaults instantiates a new JellyfinBaseItemPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinBaseItemPersonWithDefaults() *JellyfinBaseItemPerson {
	this := JellyfinBaseItemPerson{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinBaseItemPerson) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinBaseItemPerson) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JellyfinBaseItemPerson) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JellyfinBaseItemPerson) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JellyfinBaseItemPerson) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JellyfinBaseItemPerson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinBaseItemPerson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JellyfinBaseItemPerson) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinBaseItemPerson) GetRole() string {
	if o == nil || IsNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinBaseItemPerson) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *JellyfinBaseItemPerson) SetRole(v string) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *JellyfinBaseItemPerson) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *JellyfinBaseItemPerson) UnsetRole() {
	o.Role.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JellyfinBaseItemPerson) GetType() JellyfinPersonKind {
	if o == nil || IsNil(o.Type) {
		var ret JellyfinPersonKind
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinBaseItemPerson) GetTypeOk() (*JellyfinPersonKind, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given JellyfinPersonKind and assigns it to the Type field.
func (o *JellyfinBaseItemPerson) SetType(v JellyfinPersonKind) {
	o.Type = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinBaseItemPerson) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag.Get()
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinBaseItemPerson) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageTag.Get(), o.PrimaryImageTag.IsSet()
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasPrimaryImageTag() bool {
	if o != nil && o.PrimaryImageTag.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given NullableString and assigns it to the PrimaryImageTag field.
func (o *JellyfinBaseItemPerson) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag.Set(&v)
}
// SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil
func (o *JellyfinBaseItemPerson) SetPrimaryImageTagNil() {
	o.PrimaryImageTag.Set(nil)
}

// UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
func (o *JellyfinBaseItemPerson) UnsetPrimaryImageTag() {
	o.PrimaryImageTag.Unset()
}

// GetImageBlurHashes returns the ImageBlurHashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinBaseItemPerson) GetImageBlurHashes() JellyfinBaseItemPersonImageBlurHashes {
	if o == nil || IsNil(o.ImageBlurHashes.Get()) {
		var ret JellyfinBaseItemPersonImageBlurHashes
		return ret
	}
	return *o.ImageBlurHashes.Get()
}

// GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinBaseItemPerson) GetImageBlurHashesOk() (*JellyfinBaseItemPersonImageBlurHashes, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageBlurHashes.Get(), o.ImageBlurHashes.IsSet()
}

// HasImageBlurHashes returns a boolean if a field has been set.
func (o *JellyfinBaseItemPerson) HasImageBlurHashes() bool {
	if o != nil && o.ImageBlurHashes.IsSet() {
		return true
	}

	return false
}

// SetImageBlurHashes gets a reference to the given NullableJellyfinBaseItemPersonImageBlurHashes and assigns it to the ImageBlurHashes field.
func (o *JellyfinBaseItemPerson) SetImageBlurHashes(v JellyfinBaseItemPersonImageBlurHashes) {
	o.ImageBlurHashes.Set(&v)
}
// SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil
func (o *JellyfinBaseItemPerson) SetImageBlurHashesNil() {
	o.ImageBlurHashes.Set(nil)
}

// UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil
func (o *JellyfinBaseItemPerson) UnsetImageBlurHashes() {
	o.ImageBlurHashes.Unset()
}

func (o JellyfinBaseItemPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinBaseItemPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Role.IsSet() {
		toSerialize["Role"] = o.Role.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.PrimaryImageTag.IsSet() {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag.Get()
	}
	if o.ImageBlurHashes.IsSet() {
		toSerialize["ImageBlurHashes"] = o.ImageBlurHashes.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinBaseItemPerson struct {
	value *JellyfinBaseItemPerson
	isSet bool
}

func (v NullableJellyfinBaseItemPerson) Get() *JellyfinBaseItemPerson {
	return v.value
}

func (v *NullableJellyfinBaseItemPerson) Set(val *JellyfinBaseItemPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinBaseItemPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinBaseItemPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinBaseItemPerson(val *JellyfinBaseItemPerson) *NullableJellyfinBaseItemPerson {
	return &NullableJellyfinBaseItemPerson{value: val, isSet: true}
}

func (v NullableJellyfinBaseItemPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinBaseItemPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


