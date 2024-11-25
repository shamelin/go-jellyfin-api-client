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

// checks if the JellyfinUserDeletedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinUserDeletedMessage{}

// JellyfinUserDeletedMessage User deleted message.
type JellyfinUserDeletedMessage struct {
	// Gets or sets the data.
	Data *string `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinUserDeletedMessage instantiates a new JellyfinUserDeletedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinUserDeletedMessage() *JellyfinUserDeletedMessage {
	this := JellyfinUserDeletedMessage{}
	return &this
}

// NewJellyfinUserDeletedMessageWithDefaults instantiates a new JellyfinUserDeletedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinUserDeletedMessageWithDefaults() *JellyfinUserDeletedMessage {
	this := JellyfinUserDeletedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *JellyfinUserDeletedMessage) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserDeletedMessage) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinUserDeletedMessage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *JellyfinUserDeletedMessage) SetData(v string) {
	o.Data = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinUserDeletedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserDeletedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinUserDeletedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinUserDeletedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinUserDeletedMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserDeletedMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinUserDeletedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinUserDeletedMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinUserDeletedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinUserDeletedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableJellyfinUserDeletedMessage struct {
	value *JellyfinUserDeletedMessage
	isSet bool
}

func (v NullableJellyfinUserDeletedMessage) Get() *JellyfinUserDeletedMessage {
	return v.value
}

func (v *NullableJellyfinUserDeletedMessage) Set(val *JellyfinUserDeletedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinUserDeletedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinUserDeletedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinUserDeletedMessage(val *JellyfinUserDeletedMessage) *NullableJellyfinUserDeletedMessage {
	return &NullableJellyfinUserDeletedMessage{value: val, isSet: true}
}

func (v NullableJellyfinUserDeletedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinUserDeletedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


