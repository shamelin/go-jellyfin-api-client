/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"time"
)

// checks if the JellyfinTaskResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinTaskResult{}

// JellyfinTaskResult Class TaskExecutionInfo.
type JellyfinTaskResult struct {
	// Gets or sets the start time UTC.
	StartTimeUtc *time.Time `json:"StartTimeUtc,omitempty"`
	// Gets or sets the end time UTC.
	EndTimeUtc *time.Time `json:"EndTimeUtc,omitempty"`
	// Gets or sets the status.
	Status *JellyfinTaskCompletionStatus `json:"Status,omitempty"`
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the key.
	Key NullableString `json:"Key,omitempty"`
	// Gets or sets the id.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the error message.
	ErrorMessage NullableString `json:"ErrorMessage,omitempty"`
	// Gets or sets the long error message.
	LongErrorMessage NullableString `json:"LongErrorMessage,omitempty"`
}

// NewJellyfinTaskResult instantiates a new JellyfinTaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinTaskResult() *JellyfinTaskResult {
	this := JellyfinTaskResult{}
	return &this
}

// NewJellyfinTaskResultWithDefaults instantiates a new JellyfinTaskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinTaskResultWithDefaults() *JellyfinTaskResult {
	this := JellyfinTaskResult{}
	return &this
}

// GetStartTimeUtc returns the StartTimeUtc field value if set, zero value otherwise.
func (o *JellyfinTaskResult) GetStartTimeUtc() time.Time {
	if o == nil || IsNil(o.StartTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.StartTimeUtc
}

// GetStartTimeUtcOk returns a tuple with the StartTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTaskResult) GetStartTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTimeUtc) {
		return nil, false
	}
	return o.StartTimeUtc, true
}

// HasStartTimeUtc returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasStartTimeUtc() bool {
	if o != nil && !IsNil(o.StartTimeUtc) {
		return true
	}

	return false
}

// SetStartTimeUtc gets a reference to the given time.Time and assigns it to the StartTimeUtc field.
func (o *JellyfinTaskResult) SetStartTimeUtc(v time.Time) {
	o.StartTimeUtc = &v
}

// GetEndTimeUtc returns the EndTimeUtc field value if set, zero value otherwise.
func (o *JellyfinTaskResult) GetEndTimeUtc() time.Time {
	if o == nil || IsNil(o.EndTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.EndTimeUtc
}

// GetEndTimeUtcOk returns a tuple with the EndTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTaskResult) GetEndTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTimeUtc) {
		return nil, false
	}
	return o.EndTimeUtc, true
}

// HasEndTimeUtc returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasEndTimeUtc() bool {
	if o != nil && !IsNil(o.EndTimeUtc) {
		return true
	}

	return false
}

// SetEndTimeUtc gets a reference to the given time.Time and assigns it to the EndTimeUtc field.
func (o *JellyfinTaskResult) SetEndTimeUtc(v time.Time) {
	o.EndTimeUtc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JellyfinTaskResult) GetStatus() JellyfinTaskCompletionStatus {
	if o == nil || IsNil(o.Status) {
		var ret JellyfinTaskCompletionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinTaskResult) GetStatusOk() (*JellyfinTaskCompletionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given JellyfinTaskCompletionStatus and assigns it to the Status field.
func (o *JellyfinTaskResult) SetStatus(v JellyfinTaskCompletionStatus) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinTaskResult) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinTaskResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JellyfinTaskResult) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JellyfinTaskResult) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JellyfinTaskResult) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinTaskResult) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinTaskResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *JellyfinTaskResult) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *JellyfinTaskResult) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *JellyfinTaskResult) UnsetKey() {
	o.Key.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinTaskResult) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinTaskResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *JellyfinTaskResult) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *JellyfinTaskResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *JellyfinTaskResult) UnsetId() {
	o.Id.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinTaskResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinTaskResult) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *JellyfinTaskResult) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *JellyfinTaskResult) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *JellyfinTaskResult) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetLongErrorMessage returns the LongErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinTaskResult) GetLongErrorMessage() string {
	if o == nil || IsNil(o.LongErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.LongErrorMessage.Get()
}

// GetLongErrorMessageOk returns a tuple with the LongErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinTaskResult) GetLongErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongErrorMessage.Get(), o.LongErrorMessage.IsSet()
}

// HasLongErrorMessage returns a boolean if a field has been set.
func (o *JellyfinTaskResult) HasLongErrorMessage() bool {
	if o != nil && o.LongErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetLongErrorMessage gets a reference to the given NullableString and assigns it to the LongErrorMessage field.
func (o *JellyfinTaskResult) SetLongErrorMessage(v string) {
	o.LongErrorMessage.Set(&v)
}
// SetLongErrorMessageNil sets the value for LongErrorMessage to be an explicit nil
func (o *JellyfinTaskResult) SetLongErrorMessageNil() {
	o.LongErrorMessage.Set(nil)
}

// UnsetLongErrorMessage ensures that no value is present for LongErrorMessage, not even an explicit nil
func (o *JellyfinTaskResult) UnsetLongErrorMessage() {
	o.LongErrorMessage.Unset()
}

func (o JellyfinTaskResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinTaskResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTimeUtc) {
		toSerialize["StartTimeUtc"] = o.StartTimeUtc
	}
	if !IsNil(o.EndTimeUtc) {
		toSerialize["EndTimeUtc"] = o.EndTimeUtc
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["ErrorMessage"] = o.ErrorMessage.Get()
	}
	if o.LongErrorMessage.IsSet() {
		toSerialize["LongErrorMessage"] = o.LongErrorMessage.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinTaskResult struct {
	value *JellyfinTaskResult
	isSet bool
}

func (v NullableJellyfinTaskResult) Get() *JellyfinTaskResult {
	return v.value
}

func (v *NullableJellyfinTaskResult) Set(val *JellyfinTaskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinTaskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinTaskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinTaskResult(val *JellyfinTaskResult) *NullableJellyfinTaskResult {
	return &NullableJellyfinTaskResult{value: val, isSet: true}
}

func (v NullableJellyfinTaskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinTaskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

