# ModelProfileCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**ModelModelProfileConditionType**](ModelProfileConditionType.md) |  | [optional] 
**Property** | Pointer to [**ModelModelProfileConditionValue**](ModelProfileConditionValue.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelProfileCondition

`func NewModelProfileCondition() *ModelProfileCondition`

NewModelProfileCondition instantiates a new ModelProfileCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelProfileConditionWithDefaults

`func NewModelProfileConditionWithDefaults() *ModelProfileCondition`

NewModelProfileConditionWithDefaults instantiates a new ModelProfileCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *ModelProfileCondition) GetCondition() ModelModelProfileConditionType`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ModelProfileCondition) GetConditionOk() (*ModelModelProfileConditionType, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ModelProfileCondition) SetCondition(v ModelModelProfileConditionType)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ModelProfileCondition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetProperty

`func (o *ModelProfileCondition) GetProperty() ModelModelProfileConditionValue`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ModelProfileCondition) GetPropertyOk() (*ModelModelProfileConditionValue, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ModelProfileCondition) SetProperty(v ModelModelProfileConditionValue)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ModelProfileCondition) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *ModelProfileCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelProfileCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelProfileCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelProfileCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ModelProfileCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ModelProfileCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *ModelProfileCondition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ModelProfileCondition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ModelProfileCondition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ModelProfileCondition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


