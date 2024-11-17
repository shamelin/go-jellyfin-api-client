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

// checks if the QueueRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueRequestDto{}

// QueueRequestDto Class QueueRequestDto.
type QueueRequestDto struct {
	// Gets or sets the items to enqueue.
	ItemIds []string `json:"ItemIds,omitempty"`
	// Gets or sets the mode in which to add the new items.
	Mode *GroupQueueMode `json:"Mode,omitempty"`
}

// NewQueueRequestDto instantiates a new QueueRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueRequestDto() *QueueRequestDto {
	this := QueueRequestDto{}
	return &this
}

// NewQueueRequestDtoWithDefaults instantiates a new QueueRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueRequestDtoWithDefaults() *QueueRequestDto {
	this := QueueRequestDto{}
	return &this
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *QueueRequestDto) GetItemIds() []string {
	if o == nil || IsNil(o.ItemIds) {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueRequestDto) GetItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemIds) {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *QueueRequestDto) HasItemIds() bool {
	if o != nil && !IsNil(o.ItemIds) {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *QueueRequestDto) SetItemIds(v []string) {
	o.ItemIds = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *QueueRequestDto) GetMode() GroupQueueMode {
	if o == nil || IsNil(o.Mode) {
		var ret GroupQueueMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueRequestDto) GetModeOk() (*GroupQueueMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *QueueRequestDto) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given GroupQueueMode and assigns it to the Mode field.
func (o *QueueRequestDto) SetMode(v GroupQueueMode) {
	o.Mode = &v
}

func (o QueueRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemIds) {
		toSerialize["ItemIds"] = o.ItemIds
	}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableQueueRequestDto struct {
	value *QueueRequestDto
	isSet bool
}

func (v NullableQueueRequestDto) Get() *QueueRequestDto {
	return v.value
}

func (v *NullableQueueRequestDto) Set(val *QueueRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueRequestDto(val *QueueRequestDto) *NullableQueueRequestDto {
	return &NullableQueueRequestDto{value: val, isSet: true}
}

func (v NullableQueueRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


