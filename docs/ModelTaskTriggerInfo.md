# ModelTaskTriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Gets or sets the type. | [optional] 
**TimeOfDayTicks** | Pointer to **NullableInt64** | Gets or sets the time of day. | [optional] 
**IntervalTicks** | Pointer to **NullableInt64** | Gets or sets the interval. | [optional] 
**DayOfWeek** | Pointer to [**NullableModelModelDayOfWeek**](ModelDayOfWeek.md) | Gets or sets the day of week. | [optional] 
**MaxRuntimeTicks** | Pointer to **NullableInt64** | Gets or sets the maximum runtime ticks. | [optional] 

## Methods

### NewModelTaskTriggerInfo

`func NewModelTaskTriggerInfo() *ModelTaskTriggerInfo`

NewModelTaskTriggerInfo instantiates a new ModelTaskTriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTaskTriggerInfoWithDefaults

`func NewModelTaskTriggerInfoWithDefaults() *ModelTaskTriggerInfo`

NewModelTaskTriggerInfoWithDefaults instantiates a new ModelTaskTriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelTaskTriggerInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelTaskTriggerInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelTaskTriggerInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelTaskTriggerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelTaskTriggerInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelTaskTriggerInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTimeOfDayTicks

`func (o *ModelTaskTriggerInfo) GetTimeOfDayTicks() int64`

GetTimeOfDayTicks returns the TimeOfDayTicks field if non-nil, zero value otherwise.

### GetTimeOfDayTicksOk

`func (o *ModelTaskTriggerInfo) GetTimeOfDayTicksOk() (*int64, bool)`

GetTimeOfDayTicksOk returns a tuple with the TimeOfDayTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayTicks

`func (o *ModelTaskTriggerInfo) SetTimeOfDayTicks(v int64)`

SetTimeOfDayTicks sets TimeOfDayTicks field to given value.

### HasTimeOfDayTicks

`func (o *ModelTaskTriggerInfo) HasTimeOfDayTicks() bool`

HasTimeOfDayTicks returns a boolean if a field has been set.

### SetTimeOfDayTicksNil

`func (o *ModelTaskTriggerInfo) SetTimeOfDayTicksNil(b bool)`

 SetTimeOfDayTicksNil sets the value for TimeOfDayTicks to be an explicit nil

### UnsetTimeOfDayTicks
`func (o *ModelTaskTriggerInfo) UnsetTimeOfDayTicks()`

UnsetTimeOfDayTicks ensures that no value is present for TimeOfDayTicks, not even an explicit nil
### GetIntervalTicks

`func (o *ModelTaskTriggerInfo) GetIntervalTicks() int64`

GetIntervalTicks returns the IntervalTicks field if non-nil, zero value otherwise.

### GetIntervalTicksOk

`func (o *ModelTaskTriggerInfo) GetIntervalTicksOk() (*int64, bool)`

GetIntervalTicksOk returns a tuple with the IntervalTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalTicks

`func (o *ModelTaskTriggerInfo) SetIntervalTicks(v int64)`

SetIntervalTicks sets IntervalTicks field to given value.

### HasIntervalTicks

`func (o *ModelTaskTriggerInfo) HasIntervalTicks() bool`

HasIntervalTicks returns a boolean if a field has been set.

### SetIntervalTicksNil

`func (o *ModelTaskTriggerInfo) SetIntervalTicksNil(b bool)`

 SetIntervalTicksNil sets the value for IntervalTicks to be an explicit nil

### UnsetIntervalTicks
`func (o *ModelTaskTriggerInfo) UnsetIntervalTicks()`

UnsetIntervalTicks ensures that no value is present for IntervalTicks, not even an explicit nil
### GetDayOfWeek

`func (o *ModelTaskTriggerInfo) GetDayOfWeek() ModelModelDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ModelTaskTriggerInfo) GetDayOfWeekOk() (*ModelModelDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ModelTaskTriggerInfo) SetDayOfWeek(v ModelModelDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ModelTaskTriggerInfo) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *ModelTaskTriggerInfo) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *ModelTaskTriggerInfo) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetMaxRuntimeTicks

`func (o *ModelTaskTriggerInfo) GetMaxRuntimeTicks() int64`

GetMaxRuntimeTicks returns the MaxRuntimeTicks field if non-nil, zero value otherwise.

### GetMaxRuntimeTicksOk

`func (o *ModelTaskTriggerInfo) GetMaxRuntimeTicksOk() (*int64, bool)`

GetMaxRuntimeTicksOk returns a tuple with the MaxRuntimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRuntimeTicks

`func (o *ModelTaskTriggerInfo) SetMaxRuntimeTicks(v int64)`

SetMaxRuntimeTicks sets MaxRuntimeTicks field to given value.

### HasMaxRuntimeTicks

`func (o *ModelTaskTriggerInfo) HasMaxRuntimeTicks() bool`

HasMaxRuntimeTicks returns a boolean if a field has been set.

### SetMaxRuntimeTicksNil

`func (o *ModelTaskTriggerInfo) SetMaxRuntimeTicksNil(b bool)`

 SetMaxRuntimeTicksNil sets the value for MaxRuntimeTicks to be an explicit nil

### UnsetMaxRuntimeTicks
`func (o *ModelTaskTriggerInfo) UnsetMaxRuntimeTicks()`

UnsetMaxRuntimeTicks ensures that no value is present for MaxRuntimeTicks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


