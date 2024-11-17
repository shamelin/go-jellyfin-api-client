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

// checks if the GroupStateUpdateGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupStateUpdateGroupUpdate{}

// GroupStateUpdateGroupUpdate Class GroupUpdate.
type GroupStateUpdateGroupUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update type.
	Type *GroupUpdateType `json:"Type,omitempty"`
	// Gets the update data.
	Data *GroupStateUpdate `json:"Data,omitempty"`
}

// NewGroupStateUpdateGroupUpdate instantiates a new GroupStateUpdateGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStateUpdateGroupUpdate() *GroupStateUpdateGroupUpdate {
	this := GroupStateUpdateGroupUpdate{}
	return &this
}

// NewGroupStateUpdateGroupUpdateWithDefaults instantiates a new GroupStateUpdateGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStateUpdateGroupUpdateWithDefaults() *GroupStateUpdateGroupUpdate {
	this := GroupStateUpdateGroupUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupStateUpdateGroupUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStateUpdateGroupUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupStateUpdateGroupUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupStateUpdateGroupUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupStateUpdateGroupUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStateUpdateGroupUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupStateUpdateGroupUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *GroupStateUpdateGroupUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupStateUpdateGroupUpdate) GetData() GroupStateUpdate {
	if o == nil || IsNil(o.Data) {
		var ret GroupStateUpdate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStateUpdateGroupUpdate) GetDataOk() (*GroupStateUpdate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupStateUpdateGroupUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GroupStateUpdate and assigns it to the Data field.
func (o *GroupStateUpdateGroupUpdate) SetData(v GroupStateUpdate) {
	o.Data = &v
}

func (o GroupStateUpdateGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupStateUpdateGroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGroupStateUpdateGroupUpdate struct {
	value *GroupStateUpdateGroupUpdate
	isSet bool
}

func (v NullableGroupStateUpdateGroupUpdate) Get() *GroupStateUpdateGroupUpdate {
	return v.value
}

func (v *NullableGroupStateUpdateGroupUpdate) Set(val *GroupStateUpdateGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupStateUpdateGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupStateUpdateGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupStateUpdateGroupUpdate(val *GroupStateUpdateGroupUpdate) *NullableGroupStateUpdateGroupUpdate {
	return &NullableGroupStateUpdateGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupStateUpdateGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupStateUpdateGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


