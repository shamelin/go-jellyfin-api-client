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

// checks if the JellyfinSeriesTimerCreatedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinSeriesTimerCreatedMessage{}

// JellyfinSeriesTimerCreatedMessage Series timer created message.
type JellyfinSeriesTimerCreatedMessage struct {
	// Gets or sets the data.
	Data NullableJellyfinTimerEventInfo `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *JellyfinSessionMessageType `json:"MessageType,omitempty"`
}

// NewJellyfinSeriesTimerCreatedMessage instantiates a new JellyfinSeriesTimerCreatedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinSeriesTimerCreatedMessage() *JellyfinSeriesTimerCreatedMessage {
	this := JellyfinSeriesTimerCreatedMessage{}
	return &this
}

// NewJellyfinSeriesTimerCreatedMessageWithDefaults instantiates a new JellyfinSeriesTimerCreatedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinSeriesTimerCreatedMessageWithDefaults() *JellyfinSeriesTimerCreatedMessage {
	this := JellyfinSeriesTimerCreatedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinSeriesTimerCreatedMessage) GetData() JellyfinTimerEventInfo {
	if o == nil || IsNil(o.Data.Get()) {
		var ret JellyfinTimerEventInfo
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinSeriesTimerCreatedMessage) GetDataOk() (*JellyfinTimerEventInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *JellyfinSeriesTimerCreatedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableJellyfinTimerEventInfo and assigns it to the Data field.
func (o *JellyfinSeriesTimerCreatedMessage) SetData(v JellyfinTimerEventInfo) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *JellyfinSeriesTimerCreatedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *JellyfinSeriesTimerCreatedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *JellyfinSeriesTimerCreatedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSeriesTimerCreatedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *JellyfinSeriesTimerCreatedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *JellyfinSeriesTimerCreatedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *JellyfinSeriesTimerCreatedMessage) GetMessageType() JellyfinSessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret JellyfinSessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinSeriesTimerCreatedMessage) GetMessageTypeOk() (*JellyfinSessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *JellyfinSeriesTimerCreatedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given JellyfinSessionMessageType and assigns it to the MessageType field.
func (o *JellyfinSeriesTimerCreatedMessage) SetMessageType(v JellyfinSessionMessageType) {
	o.MessageType = &v
}

func (o JellyfinSeriesTimerCreatedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinSeriesTimerCreatedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableJellyfinSeriesTimerCreatedMessage struct {
	value *JellyfinSeriesTimerCreatedMessage
	isSet bool
}

func (v NullableJellyfinSeriesTimerCreatedMessage) Get() *JellyfinSeriesTimerCreatedMessage {
	return v.value
}

func (v *NullableJellyfinSeriesTimerCreatedMessage) Set(val *JellyfinSeriesTimerCreatedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinSeriesTimerCreatedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinSeriesTimerCreatedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinSeriesTimerCreatedMessage(val *JellyfinSeriesTimerCreatedMessage) *NullableJellyfinSeriesTimerCreatedMessage {
	return &NullableJellyfinSeriesTimerCreatedMessage{value: val, isSet: true}
}

func (v NullableJellyfinSeriesTimerCreatedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinSeriesTimerCreatedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


