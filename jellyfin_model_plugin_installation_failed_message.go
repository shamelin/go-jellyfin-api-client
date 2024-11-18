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

// checks if the JellyfinPluginInstallationFailedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinPluginInstallationFailedMessage{}

// JellyfinPluginInstallationFailedMessage Plugin installation failed message.
type JellyfinPluginInstallationFailedMessage struct {
	// Class InstallationInfo.
	Data NullableJellyfinInstallationInfo `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinPluginInstallationFailedMessage instantiates a new JellyfinPluginInstallationFailedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinPluginInstallationFailedMessage() *JellyfinPluginInstallationFailedMessage {
	this := JellyfinPluginInstallationFailedMessage{}
	return &this
}

// NewJellyfinPluginInstallationFailedMessageWithDefaults instantiates a new JellyfinPluginInstallationFailedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinPluginInstallationFailedMessageWithDefaults() *JellyfinPluginInstallationFailedMessage {
	this := JellyfinPluginInstallationFailedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinPluginInstallationFailedMessage) GetData() JellyfinInstallationInfo {
	if o == nil || IsNil(o.Data.Get()) {
		var ret JellyfinInstallationInfo
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinPluginInstallationFailedMessage) GetDataOk() (*JellyfinInstallationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinPluginInstallationFailedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableJellyfinInstallationInfo and assigns it to the Data field.
func (o *JellyfinPluginInstallationFailedMessage) SetData(v JellyfinInstallationInfo) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JellyfinPluginInstallationFailedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JellyfinPluginInstallationFailedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinPluginInstallationFailedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPluginInstallationFailedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinPluginInstallationFailedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinPluginInstallationFailedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinPluginInstallationFailedMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinPluginInstallationFailedMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinPluginInstallationFailedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinPluginInstallationFailedMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinPluginInstallationFailedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinPluginInstallationFailedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableJellyfinPluginInstallationFailedMessage struct {
	value *JellyfinPluginInstallationFailedMessage
	isSet bool
}

func (v NullableJellyfinPluginInstallationFailedMessage) Get() *JellyfinPluginInstallationFailedMessage {
	return v.value
}

func (v *NullableJellyfinPluginInstallationFailedMessage) Set(val *JellyfinPluginInstallationFailedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinPluginInstallationFailedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinPluginInstallationFailedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinPluginInstallationFailedMessage(val *JellyfinPluginInstallationFailedMessage) *NullableJellyfinPluginInstallationFailedMessage {
	return &NullableJellyfinPluginInstallationFailedMessage{value: val, isSet: true}
}

func (v NullableJellyfinPluginInstallationFailedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinPluginInstallationFailedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

