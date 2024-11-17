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

// checks if the PlayQueueUpdateGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayQueueUpdateGroupUpdate{}

// PlayQueueUpdateGroupUpdate Class GroupUpdate.
type PlayQueueUpdateGroupUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update type.
	Type *GroupUpdateType `json:"Type,omitempty"`
	// Gets the update data.
	Data *PlayQueueUpdate `json:"Data,omitempty"`
}

// NewPlayQueueUpdateGroupUpdate instantiates a new PlayQueueUpdateGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayQueueUpdateGroupUpdate() *PlayQueueUpdateGroupUpdate {
	this := PlayQueueUpdateGroupUpdate{}
	return &this
}

// NewPlayQueueUpdateGroupUpdateWithDefaults instantiates a new PlayQueueUpdateGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayQueueUpdateGroupUpdateWithDefaults() *PlayQueueUpdateGroupUpdate {
	this := PlayQueueUpdateGroupUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PlayQueueUpdateGroupUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdateGroupUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PlayQueueUpdateGroupUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PlayQueueUpdateGroupUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlayQueueUpdateGroupUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdateGroupUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlayQueueUpdateGroupUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *PlayQueueUpdateGroupUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PlayQueueUpdateGroupUpdate) GetData() PlayQueueUpdate {
	if o == nil || IsNil(o.Data) {
		var ret PlayQueueUpdate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayQueueUpdateGroupUpdate) GetDataOk() (*PlayQueueUpdate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PlayQueueUpdateGroupUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PlayQueueUpdate and assigns it to the Data field.
func (o *PlayQueueUpdateGroupUpdate) SetData(v PlayQueueUpdate) {
	o.Data = &v
}

func (o PlayQueueUpdateGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayQueueUpdateGroupUpdate) ToMap() (map[string]interface{}, error) {
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

type NullablePlayQueueUpdateGroupUpdate struct {
	value *PlayQueueUpdateGroupUpdate
	isSet bool
}

func (v NullablePlayQueueUpdateGroupUpdate) Get() *PlayQueueUpdateGroupUpdate {
	return v.value
}

func (v *NullablePlayQueueUpdateGroupUpdate) Set(val *PlayQueueUpdateGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayQueueUpdateGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayQueueUpdateGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayQueueUpdateGroupUpdate(val *PlayQueueUpdateGroupUpdate) *NullablePlayQueueUpdateGroupUpdate {
	return &NullablePlayQueueUpdateGroupUpdate{value: val, isSet: true}
}

func (v NullablePlayQueueUpdateGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayQueueUpdateGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


