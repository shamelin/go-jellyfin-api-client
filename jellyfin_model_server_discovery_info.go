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

// checks if the JellyfinServerDiscoveryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinServerDiscoveryInfo{}

// JellyfinServerDiscoveryInfo The server discovery info model.
type JellyfinServerDiscoveryInfo struct {
	// Gets the address.
	Address *string `json:"Address,omitempty"`
	// Gets the server identifier.
	Id *string `json:"Id,omitempty"`
	// Gets the name.
	Name *string `json:"Name,omitempty"`
	// Gets the endpoint address.
	EndpointAddress NullableString `json:"EndpointAddress,omitempty"`
}

// NewJellyfinServerDiscoveryInfo instantiates a new JellyfinServerDiscoveryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinServerDiscoveryInfo() *JellyfinServerDiscoveryInfo {
	this := JellyfinServerDiscoveryInfo{}
	return &this
}

// NewJellyfinServerDiscoveryInfoWithDefaults instantiates a new JellyfinServerDiscoveryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinServerDiscoveryInfoWithDefaults() *JellyfinServerDiscoveryInfo {
	this := JellyfinServerDiscoveryInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *JellyfinServerDiscoveryInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinServerDiscoveryInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *JellyfinServerDiscoveryInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *JellyfinServerDiscoveryInfo) SetAddress(v string) {
	o.Address = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JellyfinServerDiscoveryInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinServerDiscoveryInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JellyfinServerDiscoveryInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JellyfinServerDiscoveryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JellyfinServerDiscoveryInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinServerDiscoveryInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinServerDiscoveryInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JellyfinServerDiscoveryInfo) SetName(v string) {
	o.Name = &v
}

// GetEndpointAddress returns the EndpointAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinServerDiscoveryInfo) GetEndpointAddress() string {
	if o == nil || IsNil(o.EndpointAddress.Get()) {
		var ret string
		return ret
	}
	return *o.EndpointAddress.Get()
}

// GetEndpointAddressOk returns a tuple with the EndpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinServerDiscoveryInfo) GetEndpointAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointAddress.Get(), o.EndpointAddress.IsSet()
}

// HasEndpointAddress returns a boolean if a field has been set.
func (o *JellyfinServerDiscoveryInfo) HasEndpointAddress() bool {
	if o != nil && o.EndpointAddress.IsSet() {
		return true
	}

	return false
}

// SetEndpointAddress gets a reference to the given NullableString and assigns it to the EndpointAddress field.
func (o *JellyfinServerDiscoveryInfo) SetEndpointAddress(v string) {
	o.EndpointAddress.Set(&v)
}
// SetEndpointAddressNil sets the value for EndpointAddress to be an explicit nil
func (o *JellyfinServerDiscoveryInfo) SetEndpointAddressNil() {
	o.EndpointAddress.Set(nil)
}

// UnsetEndpointAddress ensures that no value is present for EndpointAddress, not even an explicit nil
func (o *JellyfinServerDiscoveryInfo) UnsetEndpointAddress() {
	o.EndpointAddress.Unset()
}

func (o JellyfinServerDiscoveryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinServerDiscoveryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.EndpointAddress.IsSet() {
		toSerialize["EndpointAddress"] = o.EndpointAddress.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinServerDiscoveryInfo struct {
	value *JellyfinServerDiscoveryInfo
	isSet bool
}

func (v NullableJellyfinServerDiscoveryInfo) Get() *JellyfinServerDiscoveryInfo {
	return v.value
}

func (v *NullableJellyfinServerDiscoveryInfo) Set(val *JellyfinServerDiscoveryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinServerDiscoveryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinServerDiscoveryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinServerDiscoveryInfo(val *JellyfinServerDiscoveryInfo) *NullableJellyfinServerDiscoveryInfo {
	return &NullableJellyfinServerDiscoveryInfo{value: val, isSet: true}
}

func (v NullableJellyfinServerDiscoveryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinServerDiscoveryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


