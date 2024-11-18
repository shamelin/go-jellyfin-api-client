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

// checks if the JellyfinAuthenticateUserByName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinAuthenticateUserByName{}

// JellyfinAuthenticateUserByName The authenticate user by name request body.
type JellyfinAuthenticateUserByName struct {
	// Gets or sets the username.
	Username NullableString `json:"Username,omitempty"`
	// Gets or sets the plain text password.
	Pw NullableString `json:"Pw,omitempty"`
}

// NewJellyfinAuthenticateUserByName instantiates a new JellyfinAuthenticateUserByName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinAuthenticateUserByName() *JellyfinAuthenticateUserByName {
	this := JellyfinAuthenticateUserByName{}
	return &this
}

// NewJellyfinAuthenticateUserByNameWithDefaults instantiates a new JellyfinAuthenticateUserByName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinAuthenticateUserByNameWithDefaults() *JellyfinAuthenticateUserByName {
	this := JellyfinAuthenticateUserByName{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinAuthenticateUserByName) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinAuthenticateUserByName) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *JellyfinAuthenticateUserByName) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *JellyfinAuthenticateUserByName) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *JellyfinAuthenticateUserByName) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *JellyfinAuthenticateUserByName) UnsetUsername() {
	o.Username.Unset()
}

// GetPw returns the Pw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinAuthenticateUserByName) GetPw() string {
	if o == nil || IsNil(o.Pw.Get()) {
		var ret string
		return ret
	}
	return *o.Pw.Get()
}

// GetPwOk returns a tuple with the Pw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinAuthenticateUserByName) GetPwOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pw.Get(), o.Pw.IsSet()
}

// HasPw returns a boolean if a field has been set.
func (o *JellyfinAuthenticateUserByName) HasPw() bool {
	if o != nil && o.Pw.IsSet() {
		return true
	}

	return false
}

// SetPw gets a reference to the given NullableString and assigns it to the Pw field.
func (o *JellyfinAuthenticateUserByName) SetPw(v string) {
	o.Pw.Set(&v)
}
// SetPwNil sets the value for Pw to be an explicit nil
func (o *JellyfinAuthenticateUserByName) SetPwNil() {
	o.Pw.Set(nil)
}

// UnsetPw ensures that no value is present for Pw, not even an explicit nil
func (o *JellyfinAuthenticateUserByName) UnsetPw() {
	o.Pw.Unset()
}

func (o JellyfinAuthenticateUserByName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinAuthenticateUserByName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["Username"] = o.Username.Get()
	}
	if o.Pw.IsSet() {
		toSerialize["Pw"] = o.Pw.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinAuthenticateUserByName struct {
	value *JellyfinAuthenticateUserByName
	isSet bool
}

func (v NullableJellyfinAuthenticateUserByName) Get() *JellyfinAuthenticateUserByName {
	return v.value
}

func (v *NullableJellyfinAuthenticateUserByName) Set(val *JellyfinAuthenticateUserByName) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinAuthenticateUserByName) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinAuthenticateUserByName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinAuthenticateUserByName(val *JellyfinAuthenticateUserByName) *NullableJellyfinAuthenticateUserByName {
	return &NullableJellyfinAuthenticateUserByName{value: val, isSet: true}
}

func (v NullableJellyfinAuthenticateUserByName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinAuthenticateUserByName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

