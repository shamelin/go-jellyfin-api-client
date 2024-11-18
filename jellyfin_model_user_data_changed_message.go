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

// checks if the JellyfinUserDataChangedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinUserDataChangedMessage{}

// JellyfinUserDataChangedMessage User data changed message.
type JellyfinUserDataChangedMessage struct {
	// Class UserDataChangeInfo.
	Data NullableJellyfinUserDataChangeInfo `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinUserDataChangedMessage instantiates a new JellyfinUserDataChangedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinUserDataChangedMessage() *JellyfinUserDataChangedMessage {
	this := JellyfinUserDataChangedMessage{}
	return &this
}

// NewJellyfinUserDataChangedMessageWithDefaults instantiates a new JellyfinUserDataChangedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinUserDataChangedMessageWithDefaults() *JellyfinUserDataChangedMessage {
	this := JellyfinUserDataChangedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinUserDataChangedMessage) GetData() JellyfinUserDataChangeInfo {
	if o == nil || IsNil(o.Data.Get()) {
		var ret JellyfinUserDataChangeInfo
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinUserDataChangedMessage) GetDataOk() (*JellyfinUserDataChangeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinUserDataChangedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableJellyfinUserDataChangeInfo and assigns it to the Data field.
func (o *JellyfinUserDataChangedMessage) SetData(v JellyfinUserDataChangeInfo) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JellyfinUserDataChangedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JellyfinUserDataChangedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinUserDataChangedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserDataChangedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinUserDataChangedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinUserDataChangedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinUserDataChangedMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinUserDataChangedMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinUserDataChangedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinUserDataChangedMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinUserDataChangedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinUserDataChangedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableJellyfinUserDataChangedMessage struct {
	value *JellyfinUserDataChangedMessage
	isSet bool
}

func (v NullableJellyfinUserDataChangedMessage) Get() *JellyfinUserDataChangedMessage {
	return v.value
}

func (v *NullableJellyfinUserDataChangedMessage) Set(val *JellyfinUserDataChangedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinUserDataChangedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinUserDataChangedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinUserDataChangedMessage(val *JellyfinUserDataChangedMessage) *NullableJellyfinUserDataChangedMessage {
	return &NullableJellyfinUserDataChangedMessage{value: val, isSet: true}
}

func (v NullableJellyfinUserDataChangedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinUserDataChangedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

