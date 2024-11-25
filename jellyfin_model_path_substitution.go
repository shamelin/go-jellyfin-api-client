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

// checks if the JellyfinPathSubstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinPathSubstitution{}

// JellyfinPathSubstitution Defines the MediaBrowser.Model.Configuration.PathSubstitution.
type JellyfinPathSubstitution struct {
	// Gets or sets the value to substitute.
	From *string `json:"From,omitempty"`
	// Gets or sets the value to substitution with.
	To *string `json:"To,omitempty"`
}

// NewJellyfinPathSubstitution instantiates a new JellyfinPathSubstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinPathSubstitution() *JellyfinPathSubstitution {
	this := JellyfinPathSubstitution{}
	return &this
}

// NewJellyfinPathSubstitutionWithDefaults instantiates a new JellyfinPathSubstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinPathSubstitutionWithDefaults() *JellyfinPathSubstitution {
	this := JellyfinPathSubstitution{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *JellyfinPathSubstitution) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPathSubstitution) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *JellyfinPathSubstitution) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *JellyfinPathSubstitution) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *JellyfinPathSubstitution) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPathSubstitution) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *JellyfinPathSubstitution) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *JellyfinPathSubstitution) SetTo(v string) {
	o.To = &v
}

func (o JellyfinPathSubstitution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinPathSubstitution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["From"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["To"] = o.To
	}
	return toSerialize, nil
}

type NullableJellyfinPathSubstitution struct {
	value *JellyfinPathSubstitution
	isSet bool
}

func (v NullableJellyfinPathSubstitution) Get() *JellyfinPathSubstitution {
	return v.value
}

func (v *NullableJellyfinPathSubstitution) Set(val *JellyfinPathSubstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPathSubstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPathSubstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPathSubstitution(val *JellyfinPathSubstitution) *NullableJellyfinPathSubstitution {
	return &NullableJellyfinPathSubstitution{value: val, isSet: true}
}

func (v NullableJellyfinPathSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPathSubstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


