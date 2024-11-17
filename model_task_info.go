/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TaskInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskInfo{}

// TaskInfo Class TaskInfo.
type TaskInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the state of the task.
	State *TaskState `json:"State,omitempty"`
	// Gets or sets the progress.
	CurrentProgressPercentage NullableFloat64 `json:"CurrentProgressPercentage,omitempty"`
	// Gets or sets the id.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the last execution result.
	LastExecutionResult NullableTaskResult `json:"LastExecutionResult,omitempty"`
	// Gets or sets the triggers.
	Triggers []TaskTriggerInfo `json:"Triggers,omitempty"`
	// Gets or sets the description.
	Description NullableString `json:"Description,omitempty"`
	// Gets or sets the category.
	Category NullableString `json:"Category,omitempty"`
	// Gets or sets a value indicating whether this instance is hidden.
	IsHidden *bool `json:"IsHidden,omitempty"`
	// Gets or sets the key.
	Key NullableString `json:"Key,omitempty"`
}

// NewTaskInfo instantiates a new TaskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskInfo() *TaskInfo {
	this := TaskInfo{}
	return &this
}

// NewTaskInfoWithDefaults instantiates a new TaskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskInfoWithDefaults() *TaskInfo {
	this := TaskInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TaskInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TaskInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TaskInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TaskInfo) UnsetName() {
	o.Name.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TaskInfo) GetState() TaskState {
	if o == nil || IsNil(o.State) {
		var ret TaskState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInfo) GetStateOk() (*TaskState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TaskInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given TaskState and assigns it to the State field.
func (o *TaskInfo) SetState(v TaskState) {
	o.State = &v
}

// GetCurrentProgressPercentage returns the CurrentProgressPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetCurrentProgressPercentage() float64 {
	if o == nil || IsNil(o.CurrentProgressPercentage.Get()) {
		var ret float64
		return ret
	}
	return *o.CurrentProgressPercentage.Get()
}

// GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetCurrentProgressPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentProgressPercentage.Get(), o.CurrentProgressPercentage.IsSet()
}

// HasCurrentProgressPercentage returns a boolean if a field has been set.
func (o *TaskInfo) HasCurrentProgressPercentage() bool {
	if o != nil && o.CurrentProgressPercentage.IsSet() {
		return true
	}

	return false
}

// SetCurrentProgressPercentage gets a reference to the given NullableFloat64 and assigns it to the CurrentProgressPercentage field.
func (o *TaskInfo) SetCurrentProgressPercentage(v float64) {
	o.CurrentProgressPercentage.Set(&v)
}
// SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil
func (o *TaskInfo) SetCurrentProgressPercentageNil() {
	o.CurrentProgressPercentage.Set(nil)
}

// UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
func (o *TaskInfo) UnsetCurrentProgressPercentage() {
	o.CurrentProgressPercentage.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TaskInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TaskInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TaskInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TaskInfo) UnsetId() {
	o.Id.Unset()
}

// GetLastExecutionResult returns the LastExecutionResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetLastExecutionResult() TaskResult {
	if o == nil || IsNil(o.LastExecutionResult.Get()) {
		var ret TaskResult
		return ret
	}
	return *o.LastExecutionResult.Get()
}

// GetLastExecutionResultOk returns a tuple with the LastExecutionResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetLastExecutionResultOk() (*TaskResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionResult.Get(), o.LastExecutionResult.IsSet()
}

// HasLastExecutionResult returns a boolean if a field has been set.
func (o *TaskInfo) HasLastExecutionResult() bool {
	if o != nil && o.LastExecutionResult.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionResult gets a reference to the given NullableTaskResult and assigns it to the LastExecutionResult field.
func (o *TaskInfo) SetLastExecutionResult(v TaskResult) {
	o.LastExecutionResult.Set(&v)
}
// SetLastExecutionResultNil sets the value for LastExecutionResult to be an explicit nil
func (o *TaskInfo) SetLastExecutionResultNil() {
	o.LastExecutionResult.Set(nil)
}

// UnsetLastExecutionResult ensures that no value is present for LastExecutionResult, not even an explicit nil
func (o *TaskInfo) UnsetLastExecutionResult() {
	o.LastExecutionResult.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetTriggers() []TaskTriggerInfo {
	if o == nil {
		var ret []TaskTriggerInfo
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetTriggersOk() ([]TaskTriggerInfo, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *TaskInfo) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []TaskTriggerInfo and assigns it to the Triggers field.
func (o *TaskInfo) SetTriggers(v []TaskTriggerInfo) {
	o.Triggers = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TaskInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TaskInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TaskInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TaskInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *TaskInfo) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *TaskInfo) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *TaskInfo) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *TaskInfo) UnsetCategory() {
	o.Category.Unset()
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *TaskInfo) GetIsHidden() bool {
	if o == nil || IsNil(o.IsHidden) {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInfo) GetIsHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHidden) {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *TaskInfo) HasIsHidden() bool {
	if o != nil && !IsNil(o.IsHidden) {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *TaskInfo) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInfo) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInfo) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *TaskInfo) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *TaskInfo) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *TaskInfo) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *TaskInfo) UnsetKey() {
	o.Key.Unset()
}

func (o TaskInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if o.CurrentProgressPercentage.IsSet() {
		toSerialize["CurrentProgressPercentage"] = o.CurrentProgressPercentage.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.LastExecutionResult.IsSet() {
		toSerialize["LastExecutionResult"] = o.LastExecutionResult.Get()
	}
	if o.Triggers != nil {
		toSerialize["Triggers"] = o.Triggers
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Category.IsSet() {
		toSerialize["Category"] = o.Category.Get()
	}
	if !IsNil(o.IsHidden) {
		toSerialize["IsHidden"] = o.IsHidden
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableTaskInfo struct {
	value *TaskInfo
	isSet bool
}

func (v NullableTaskInfo) Get() *TaskInfo {
	return v.value
}

func (v *NullableTaskInfo) Set(val *TaskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskInfo(val *TaskInfo) *NullableTaskInfo {
	return &NullableTaskInfo{value: val, isSet: true}
}

func (v NullableTaskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

