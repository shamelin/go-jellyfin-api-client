# ModelGroupStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**ModelModelGroupStateType**](ModelGroupStateType.md) | Gets the state of the group. | [optional] 
**Reason** | Pointer to [**ModelModelPlaybackRequestType**](ModelPlaybackRequestType.md) | Gets the reason of the state change. | [optional] 

## Methods

### NewModelGroupStateUpdate

`func NewModelGroupStateUpdate() *ModelGroupStateUpdate`

NewModelGroupStateUpdate instantiates a new ModelGroupStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGroupStateUpdateWithDefaults

`func NewModelGroupStateUpdateWithDefaults() *ModelGroupStateUpdate`

NewModelGroupStateUpdateWithDefaults instantiates a new ModelGroupStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ModelGroupStateUpdate) GetState() ModelModelGroupStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelGroupStateUpdate) GetStateOk() (*ModelModelGroupStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelGroupStateUpdate) SetState(v ModelModelGroupStateType)`

SetState sets State field to given value.

### HasState

`func (o *ModelGroupStateUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *ModelGroupStateUpdate) GetReason() ModelModelPlaybackRequestType`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelGroupStateUpdate) GetReasonOk() (*ModelModelPlaybackRequestType, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelGroupStateUpdate) SetReason(v ModelModelPlaybackRequestType)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ModelGroupStateUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


