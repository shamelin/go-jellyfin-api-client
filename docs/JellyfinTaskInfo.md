# JellyfinTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**State** | Pointer to [**JellyfinJellyfinTaskState**](JellyfinTaskState.md) | Gets or sets the state of the task. | [optional] 
**CurrentProgressPercentage** | Pointer to **NullableFloat64** | Gets or sets the progress. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**LastExecutionResult** | Pointer to [**NullableJellyfinJellyfinTaskResult**](JellyfinTaskResult.md) | Gets or sets the last execution result. | [optional] 
**Triggers** | Pointer to [**[]JellyfinJellyfinTaskTriggerInfo**](JellyfinJellyfinTaskTriggerInfo.md) | Gets or sets the triggers. | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the description. | [optional] 
**Category** | Pointer to **NullableString** | Gets or sets the category. | [optional] 
**IsHidden** | Pointer to **bool** | Gets or sets a value indicating whether this instance is hidden. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 

## Methods

### NewJellyfinTaskInfo

`func NewJellyfinTaskInfo() *JellyfinTaskInfo`

NewJellyfinTaskInfo instantiates a new JellyfinTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinTaskInfoWithDefaults

`func NewJellyfinTaskInfoWithDefaults() *JellyfinTaskInfo`

NewJellyfinTaskInfoWithDefaults instantiates a new JellyfinTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinTaskInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinTaskInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *JellyfinTaskInfo) GetState() JellyfinJellyfinTaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JellyfinTaskInfo) GetStateOk() (*JellyfinJellyfinTaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JellyfinTaskInfo) SetState(v JellyfinJellyfinTaskState)`

SetState sets State field to given value.

### HasState

`func (o *JellyfinTaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentProgressPercentage

`func (o *JellyfinTaskInfo) GetCurrentProgressPercentage() float64`

GetCurrentProgressPercentage returns the CurrentProgressPercentage field if non-nil, zero value otherwise.

### GetCurrentProgressPercentageOk

`func (o *JellyfinTaskInfo) GetCurrentProgressPercentageOk() (*float64, bool)`

GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgressPercentage

`func (o *JellyfinTaskInfo) SetCurrentProgressPercentage(v float64)`

SetCurrentProgressPercentage sets CurrentProgressPercentage field to given value.

### HasCurrentProgressPercentage

`func (o *JellyfinTaskInfo) HasCurrentProgressPercentage() bool`

HasCurrentProgressPercentage returns a boolean if a field has been set.

### SetCurrentProgressPercentageNil

`func (o *JellyfinTaskInfo) SetCurrentProgressPercentageNil(b bool)`

 SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil

### UnsetCurrentProgressPercentage
`func (o *JellyfinTaskInfo) UnsetCurrentProgressPercentage()`

UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
### GetId

`func (o *JellyfinTaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinTaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinTaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinTaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinTaskInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinTaskInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastExecutionResult

`func (o *JellyfinTaskInfo) GetLastExecutionResult() JellyfinJellyfinTaskResult`

GetLastExecutionResult returns the LastExecutionResult field if non-nil, zero value otherwise.

### GetLastExecutionResultOk

`func (o *JellyfinTaskInfo) GetLastExecutionResultOk() (*JellyfinJellyfinTaskResult, bool)`

GetLastExecutionResultOk returns a tuple with the LastExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionResult

`func (o *JellyfinTaskInfo) SetLastExecutionResult(v JellyfinJellyfinTaskResult)`

SetLastExecutionResult sets LastExecutionResult field to given value.

### HasLastExecutionResult

`func (o *JellyfinTaskInfo) HasLastExecutionResult() bool`

HasLastExecutionResult returns a boolean if a field has been set.

### SetLastExecutionResultNil

`func (o *JellyfinTaskInfo) SetLastExecutionResultNil(b bool)`

 SetLastExecutionResultNil sets the value for LastExecutionResult to be an explicit nil

### UnsetLastExecutionResult
`func (o *JellyfinTaskInfo) UnsetLastExecutionResult()`

UnsetLastExecutionResult ensures that no value is present for LastExecutionResult, not even an explicit nil
### GetTriggers

`func (o *JellyfinTaskInfo) GetTriggers() []JellyfinJellyfinTaskTriggerInfo`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *JellyfinTaskInfo) GetTriggersOk() (*[]JellyfinJellyfinTaskTriggerInfo, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *JellyfinTaskInfo) SetTriggers(v []JellyfinJellyfinTaskTriggerInfo)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *JellyfinTaskInfo) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *JellyfinTaskInfo) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *JellyfinTaskInfo) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetDescription

`func (o *JellyfinTaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JellyfinTaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JellyfinTaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JellyfinTaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JellyfinTaskInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JellyfinTaskInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *JellyfinTaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *JellyfinTaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *JellyfinTaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *JellyfinTaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *JellyfinTaskInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *JellyfinTaskInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsHidden

`func (o *JellyfinTaskInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *JellyfinTaskInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *JellyfinTaskInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *JellyfinTaskInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetKey

`func (o *JellyfinTaskInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JellyfinTaskInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JellyfinTaskInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *JellyfinTaskInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *JellyfinTaskInfo) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *JellyfinTaskInfo) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


