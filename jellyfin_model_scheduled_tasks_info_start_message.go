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

// checks if the JellyfinScheduledTasksInfoStartMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinScheduledTasksInfoStartMessage{}

// JellyfinScheduledTasksInfoStartMessage Scheduled tasks info start message.  Data is the timing data encoded as \"$initialDelay,$interval\" in ms.
type JellyfinScheduledTasksInfoStartMessage struct {
	// Gets or sets the data.
	Data NullableString `json:"Data,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinScheduledTasksInfoStartMessage instantiates a new JellyfinScheduledTasksInfoStartMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinScheduledTasksInfoStartMessage() *JellyfinScheduledTasksInfoStartMessage {
	this := JellyfinScheduledTasksInfoStartMessage{}
	return &this
}

// NewJellyfinScheduledTasksInfoStartMessageWithDefaults instantiates a new JellyfinScheduledTasksInfoStartMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinScheduledTasksInfoStartMessageWithDefaults() *JellyfinScheduledTasksInfoStartMessage {
	this := JellyfinScheduledTasksInfoStartMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinScheduledTasksInfoStartMessage) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinScheduledTasksInfoStartMessage) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinScheduledTasksInfoStartMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *JellyfinScheduledTasksInfoStartMessage) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JellyfinScheduledTasksInfoStartMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JellyfinScheduledTasksInfoStartMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinScheduledTasksInfoStartMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinScheduledTasksInfoStartMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinScheduledTasksInfoStartMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinScheduledTasksInfoStartMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinScheduledTasksInfoStartMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinScheduledTasksInfoStartMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableJellyfinScheduledTasksInfoStartMessage struct {
	value *JellyfinScheduledTasksInfoStartMessage
	isSet bool
}

func (v NullableJellyfinScheduledTasksInfoStartMessage) Get() *JellyfinScheduledTasksInfoStartMessage {
	return v.value
}

func (v *NullableJellyfinScheduledTasksInfoStartMessage) Set(val *JellyfinScheduledTasksInfoStartMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinScheduledTasksInfoStartMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinScheduledTasksInfoStartMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinScheduledTasksInfoStartMessage(val *JellyfinScheduledTasksInfoStartMessage) *NullableJellyfinScheduledTasksInfoStartMessage {
	return &NullableJellyfinScheduledTasksInfoStartMessage{value: val, isSet: true}
}

func (v NullableJellyfinScheduledTasksInfoStartMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinScheduledTasksInfoStartMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


