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

// checks if the ForceKeepAliveMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForceKeepAliveMessage{}

// ForceKeepAliveMessage Force keep alive websocket messages.
type ForceKeepAliveMessage struct {
	// Gets or sets the data.
	Data *int32 `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewForceKeepAliveMessage instantiates a new ForceKeepAliveMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForceKeepAliveMessage() *ForceKeepAliveMessage {
	this := ForceKeepAliveMessage{}
	return &this
}

// NewForceKeepAliveMessageWithDefaults instantiates a new ForceKeepAliveMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForceKeepAliveMessageWithDefaults() *ForceKeepAliveMessage {
	this := ForceKeepAliveMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ForceKeepAliveMessage) GetData() int32 {
	if o == nil || IsNil(o.Data) {
		var ret int32
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForceKeepAliveMessage) GetDataOk() (*int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ForceKeepAliveMessage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given int32 and assigns it to the Data field.
func (o *ForceKeepAliveMessage) SetData(v int32) {
	o.Data = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ForceKeepAliveMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForceKeepAliveMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ForceKeepAliveMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ForceKeepAliveMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ForceKeepAliveMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForceKeepAliveMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ForceKeepAliveMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ForceKeepAliveMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ForceKeepAliveMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForceKeepAliveMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableForceKeepAliveMessage struct {
	value *ForceKeepAliveMessage
	isSet bool
}

func (v NullableForceKeepAliveMessage) Get() *ForceKeepAliveMessage {
	return v.value
}

func (v *NullableForceKeepAliveMessage) Set(val *ForceKeepAliveMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableForceKeepAliveMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableForceKeepAliveMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForceKeepAliveMessage(val *ForceKeepAliveMessage) *NullableForceKeepAliveMessage {
	return &NullableForceKeepAliveMessage{value: val, isSet: true}
}

func (v NullableForceKeepAliveMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForceKeepAliveMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


