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

// checks if the JellyfinRestartRequiredMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinRestartRequiredMessage{}

// JellyfinRestartRequiredMessage Restart required.
type JellyfinRestartRequiredMessage struct {
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinRestartRequiredMessage instantiates a new JellyfinRestartRequiredMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinRestartRequiredMessage() *JellyfinRestartRequiredMessage {
	this := JellyfinRestartRequiredMessage{}
	return &this
}

// NewJellyfinRestartRequiredMessageWithDefaults instantiates a new JellyfinRestartRequiredMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinRestartRequiredMessageWithDefaults() *JellyfinRestartRequiredMessage {
	this := JellyfinRestartRequiredMessage{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinRestartRequiredMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinRestartRequiredMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinRestartRequiredMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinRestartRequiredMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinRestartRequiredMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinRestartRequiredMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinRestartRequiredMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinRestartRequiredMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinRestartRequiredMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinRestartRequiredMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableJellyfinRestartRequiredMessage struct {
	value *JellyfinRestartRequiredMessage
	isSet bool
}

func (v NullableJellyfinRestartRequiredMessage) Get() *JellyfinRestartRequiredMessage {
	return v.value
}

func (v *NullableJellyfinRestartRequiredMessage) Set(val *JellyfinRestartRequiredMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinRestartRequiredMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinRestartRequiredMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinRestartRequiredMessage(val *JellyfinRestartRequiredMessage) *NullableJellyfinRestartRequiredMessage {
	return &NullableJellyfinRestartRequiredMessage{value: val, isSet: true}
}

func (v NullableJellyfinRestartRequiredMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinRestartRequiredMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

