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

// checks if the ServerShuttingDownMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerShuttingDownMessage{}

// ServerShuttingDownMessage Server shutting down message.
type ServerShuttingDownMessage struct {
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewServerShuttingDownMessage instantiates a new ServerShuttingDownMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerShuttingDownMessage() *ServerShuttingDownMessage {
	this := ServerShuttingDownMessage{}
	return &this
}

// NewServerShuttingDownMessageWithDefaults instantiates a new ServerShuttingDownMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerShuttingDownMessageWithDefaults() *ServerShuttingDownMessage {
	this := ServerShuttingDownMessage{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ServerShuttingDownMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerShuttingDownMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ServerShuttingDownMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ServerShuttingDownMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ServerShuttingDownMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerShuttingDownMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ServerShuttingDownMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ServerShuttingDownMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ServerShuttingDownMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerShuttingDownMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableServerShuttingDownMessage struct {
	value *ServerShuttingDownMessage
	isSet bool
}

func (v NullableServerShuttingDownMessage) Get() *ServerShuttingDownMessage {
	return v.value
}

func (v *NullableServerShuttingDownMessage) Set(val *ServerShuttingDownMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServerShuttingDownMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServerShuttingDownMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerShuttingDownMessage(val *ServerShuttingDownMessage) *NullableServerShuttingDownMessage {
	return &NullableServerShuttingDownMessage{value: val, isSet: true}
}

func (v NullableServerShuttingDownMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerShuttingDownMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


