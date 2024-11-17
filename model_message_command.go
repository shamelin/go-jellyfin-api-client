/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MessageCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCommand{}

// MessageCommand struct for MessageCommand
type MessageCommand struct {
	Header NullableString `json:"Header,omitempty"`
	Text string `json:"Text"`
	TimeoutMs NullableInt64 `json:"TimeoutMs,omitempty"`
}

type _MessageCommand MessageCommand

// NewMessageCommand instantiates a new MessageCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCommand(text string) *MessageCommand {
	this := MessageCommand{}
	this.Text = text
	return &this
}

// NewMessageCommandWithDefaults instantiates a new MessageCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCommandWithDefaults() *MessageCommand {
	this := MessageCommand{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCommand) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCommand) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *MessageCommand) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *MessageCommand) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *MessageCommand) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *MessageCommand) UnsetHeader() {
	o.Header.Unset()
}

// GetText returns the Text field value
func (o *MessageCommand) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageCommand) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageCommand) SetText(v string) {
	o.Text = v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCommand) GetTimeoutMs() int64 {
	if o == nil || IsNil(o.TimeoutMs.Get()) {
		var ret int64
		return ret
	}
	return *o.TimeoutMs.Get()
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCommand) GetTimeoutMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeoutMs.Get(), o.TimeoutMs.IsSet()
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *MessageCommand) HasTimeoutMs() bool {
	if o != nil && o.TimeoutMs.IsSet() {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given NullableInt64 and assigns it to the TimeoutMs field.
func (o *MessageCommand) SetTimeoutMs(v int64) {
	o.TimeoutMs.Set(&v)
}
// SetTimeoutMsNil sets the value for TimeoutMs to be an explicit nil
func (o *MessageCommand) SetTimeoutMsNil() {
	o.TimeoutMs.Set(nil)
}

// UnsetTimeoutMs ensures that no value is present for TimeoutMs, not even an explicit nil
func (o *MessageCommand) UnsetTimeoutMs() {
	o.TimeoutMs.Unset()
}

func (o MessageCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Header.IsSet() {
		toSerialize["Header"] = o.Header.Get()
	}
	toSerialize["Text"] = o.Text
	if o.TimeoutMs.IsSet() {
		toSerialize["TimeoutMs"] = o.TimeoutMs.Get()
	}
	return toSerialize, nil
}

func (o *MessageCommand) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Text",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMessageCommand := _MessageCommand{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageCommand)

	if err != nil {
		return err
	}

	*o = MessageCommand(varMessageCommand)

	return err
}

type NullableMessageCommand struct {
	value *MessageCommand
	isSet bool
}

func (v NullableMessageCommand) Get() *MessageCommand {
	return v.value
}

func (v *NullableMessageCommand) Set(val *MessageCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCommand(val *MessageCommand) *NullableMessageCommand {
	return &NullableMessageCommand{value: val, isSet: true}
}

func (v NullableMessageCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


