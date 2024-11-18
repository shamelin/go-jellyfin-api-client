# JellyfinTaskTriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Gets or sets the type. | [optional] 
**TimeOfDayTicks** | Pointer to **NullableInt64** | Gets or sets the time of day. | [optional] 
**IntervalTicks** | Pointer to **NullableInt64** | Gets or sets the interval. | [optional] 
**DayOfWeek** | Pointer to [**NullableJellyfinJellyfinDayOfWeek**](JellyfinDayOfWeek.md) | Gets or sets the day of week. | [optional] 
**MaxRuntimeTicks** | Pointer to **NullableInt64** | Gets or sets the maximum runtime ticks. | [optional] 

## Methods

### NewJellyfinTaskTriggerInfo

`func NewJellyfinTaskTriggerInfo() *JellyfinTaskTriggerInfo`

NewJellyfinTaskTriggerInfo instantiates a new JellyfinTaskTriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinTaskTriggerInfoWithDefaults

`func NewJellyfinTaskTriggerInfoWithDefaults() *JellyfinTaskTriggerInfo`

NewJellyfinTaskTriggerInfoWithDefaults instantiates a new JellyfinTaskTriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JellyfinTaskTriggerInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinTaskTriggerInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinTaskTriggerInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinTaskTriggerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JellyfinTaskTriggerInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JellyfinTaskTriggerInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTimeOfDayTicks

`func (o *JellyfinTaskTriggerInfo) GetTimeOfDayTicks() int64`

GetTimeOfDayTicks returns the TimeOfDayTicks field if non-nil, zero value otherwise.

### GetTimeOfDayTicksOk

`func (o *JellyfinTaskTriggerInfo) GetTimeOfDayTicksOk() (*int64, bool)`

GetTimeOfDayTicksOk returns a tuple with the TimeOfDayTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayTicks

`func (o *JellyfinTaskTriggerInfo) SetTimeOfDayTicks(v int64)`

SetTimeOfDayTicks sets TimeOfDayTicks field to given value.

### HasTimeOfDayTicks

`func (o *JellyfinTaskTriggerInfo) HasTimeOfDayTicks() bool`

HasTimeOfDayTicks returns a boolean if a field has been set.

### SetTimeOfDayTicksNil

`func (o *JellyfinTaskTriggerInfo) SetTimeOfDayTicksNil(b bool)`

 SetTimeOfDayTicksNil sets the value for TimeOfDayTicks to be an explicit nil

### UnsetTimeOfDayTicks
`func (o *JellyfinTaskTriggerInfo) UnsetTimeOfDayTicks()`

UnsetTimeOfDayTicks ensures that no value is present for TimeOfDayTicks, not even an explicit nil
### GetIntervalTicks

`func (o *JellyfinTaskTriggerInfo) GetIntervalTicks() int64`

GetIntervalTicks returns the IntervalTicks field if non-nil, zero value otherwise.

### GetIntervalTicksOk

`func (o *JellyfinTaskTriggerInfo) GetIntervalTicksOk() (*int64, bool)`

GetIntervalTicksOk returns a tuple with the IntervalTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalTicks

`func (o *JellyfinTaskTriggerInfo) SetIntervalTicks(v int64)`

SetIntervalTicks sets IntervalTicks field to given value.

### HasIntervalTicks

`func (o *JellyfinTaskTriggerInfo) HasIntervalTicks() bool`

HasIntervalTicks returns a boolean if a field has been set.

### SetIntervalTicksNil

`func (o *JellyfinTaskTriggerInfo) SetIntervalTicksNil(b bool)`

 SetIntervalTicksNil sets the value for IntervalTicks to be an explicit nil

### UnsetIntervalTicks
`func (o *JellyfinTaskTriggerInfo) UnsetIntervalTicks()`

UnsetIntervalTicks ensures that no value is present for IntervalTicks, not even an explicit nil
### GetDayOfWeek

`func (o *JellyfinTaskTriggerInfo) GetDayOfWeek() JellyfinJellyfinDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *JellyfinTaskTriggerInfo) GetDayOfWeekOk() (*JellyfinJellyfinDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *JellyfinTaskTriggerInfo) SetDayOfWeek(v JellyfinJellyfinDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *JellyfinTaskTriggerInfo) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *JellyfinTaskTriggerInfo) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *JellyfinTaskTriggerInfo) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetMaxRuntimeTicks

`func (o *JellyfinTaskTriggerInfo) GetMaxRuntimeTicks() int64`

GetMaxRuntimeTicks returns the MaxRuntimeTicks field if non-nil, zero value otherwise.

### GetMaxRuntimeTicksOk

`func (o *JellyfinTaskTriggerInfo) GetMaxRuntimeTicksOk() (*int64, bool)`

GetMaxRuntimeTicksOk returns a tuple with the MaxRuntimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRuntimeTicks

`func (o *JellyfinTaskTriggerInfo) SetMaxRuntimeTicks(v int64)`

SetMaxRuntimeTicks sets MaxRuntimeTicks field to given value.

### HasMaxRuntimeTicks

`func (o *JellyfinTaskTriggerInfo) HasMaxRuntimeTicks() bool`

HasMaxRuntimeTicks returns a boolean if a field has been set.

### SetMaxRuntimeTicksNil

`func (o *JellyfinTaskTriggerInfo) SetMaxRuntimeTicksNil(b bool)`

 SetMaxRuntimeTicksNil sets the value for MaxRuntimeTicks to be an explicit nil

### UnsetMaxRuntimeTicks
`func (o *JellyfinTaskTriggerInfo) UnsetMaxRuntimeTicks()`

UnsetMaxRuntimeTicks ensures that no value is present for MaxRuntimeTicks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


