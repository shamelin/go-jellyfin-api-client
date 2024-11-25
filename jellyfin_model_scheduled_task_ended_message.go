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

// checks if the JellyfinScheduledTaskEndedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinScheduledTaskEndedMessage{}

// JellyfinScheduledTaskEndedMessage Scheduled task ended message.
type JellyfinScheduledTaskEndedMessage struct {
	// Class TaskExecutionInfo.
	Data NullableJellyfinTaskResult `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinScheduledTaskEndedMessage instantiates a new JellyfinScheduledTaskEndedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinScheduledTaskEndedMessage() *JellyfinScheduledTaskEndedMessage {
	this := JellyfinScheduledTaskEndedMessage{}
	return &this
}

// NewJellyfinScheduledTaskEndedMessageWithDefaults instantiates a new JellyfinScheduledTaskEndedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinScheduledTaskEndedMessageWithDefaults() *JellyfinScheduledTaskEndedMessage {
	this := JellyfinScheduledTaskEndedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinScheduledTaskEndedMessage) GetData() JellyfinTaskResult {
	if o == nil || IsNil(o.Data.Get()) {
		var ret JellyfinTaskResult
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinScheduledTaskEndedMessage) GetDataOk() (*JellyfinTaskResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinScheduledTaskEndedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableJellyfinTaskResult and assigns it to the Data field.
func (o *JellyfinScheduledTaskEndedMessage) SetData(v JellyfinTaskResult) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JellyfinScheduledTaskEndedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JellyfinScheduledTaskEndedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinScheduledTaskEndedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinScheduledTaskEndedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinScheduledTaskEndedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinScheduledTaskEndedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinScheduledTaskEndedMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinScheduledTaskEndedMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinScheduledTaskEndedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinScheduledTaskEndedMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinScheduledTaskEndedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinScheduledTaskEndedMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableJellyfinScheduledTaskEndedMessage struct {
	value *JellyfinScheduledTaskEndedMessage
	isSet bool
}

func (v NullableJellyfinScheduledTaskEndedMessage) Get() *JellyfinScheduledTaskEndedMessage {
	return v.value
}

func (v *NullableJellyfinScheduledTaskEndedMessage) Set(val *JellyfinScheduledTaskEndedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinScheduledTaskEndedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinScheduledTaskEndedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinScheduledTaskEndedMessage(val *JellyfinScheduledTaskEndedMessage) *NullableJellyfinScheduledTaskEndedMessage {
	return &NullableJellyfinScheduledTaskEndedMessage{value: val, isSet: true}
}

func (v NullableJellyfinScheduledTaskEndedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinScheduledTaskEndedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


