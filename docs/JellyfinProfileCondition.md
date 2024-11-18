# JellyfinProfileCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**JellyfinJellyfinProfileConditionType**](JellyfinProfileConditionType.md) |  | [optional] 
**Property** | Pointer to [**JellyfinJellyfinProfileConditionValue**](JellyfinProfileConditionValue.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewJellyfinProfileCondition

`func NewJellyfinProfileCondition() *JellyfinProfileCondition`

NewJellyfinProfileCondition instantiates a new JellyfinProfileCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinProfileConditionWithDefaults

`func NewJellyfinProfileConditionWithDefaults() *JellyfinProfileCondition`

NewJellyfinProfileConditionWithDefaults instantiates a new JellyfinProfileCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *JellyfinProfileCondition) GetCondition() JellyfinJellyfinProfileConditionType`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *JellyfinProfileCondition) GetConditionOk() (*JellyfinJellyfinProfileConditionType, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *JellyfinProfileCondition) SetCondition(v JellyfinJellyfinProfileConditionType)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *JellyfinProfileCondition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetProperty

`func (o *JellyfinProfileCondition) GetProperty() JellyfinJellyfinProfileConditionValue`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *JellyfinProfileCondition) GetPropertyOk() (*JellyfinJellyfinProfileConditionValue, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *JellyfinProfileCondition) SetProperty(v JellyfinJellyfinProfileConditionValue)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *JellyfinProfileCondition) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *JellyfinProfileCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JellyfinProfileCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JellyfinProfileCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *JellyfinProfileCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *JellyfinProfileCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JellyfinProfileCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *JellyfinProfileCondition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *JellyfinProfileCondition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *JellyfinProfileCondition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *JellyfinProfileCondition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


