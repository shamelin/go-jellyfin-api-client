# JellyfinGroupStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**JellyfinJellyfinGroupStateType**](JellyfinGroupStateType.md) | Gets the state of the group. | [optional] 
**Reason** | Pointer to [**JellyfinJellyfinPlaybackRequestType**](JellyfinPlaybackRequestType.md) | Gets the reason of the state change. | [optional] 

## Methods

### NewJellyfinGroupStateUpdate

`func NewJellyfinGroupStateUpdate() *JellyfinGroupStateUpdate`

NewJellyfinGroupStateUpdate instantiates a new JellyfinGroupStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGroupStateUpdateWithDefaults

`func NewJellyfinGroupStateUpdateWithDefaults() *JellyfinGroupStateUpdate`

NewJellyfinGroupStateUpdateWithDefaults instantiates a new JellyfinGroupStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *JellyfinGroupStateUpdate) GetState() JellyfinJellyfinGroupStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JellyfinGroupStateUpdate) GetStateOk() (*JellyfinJellyfinGroupStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JellyfinGroupStateUpdate) SetState(v JellyfinJellyfinGroupStateType)`

SetState sets State field to given value.

### HasState

`func (o *JellyfinGroupStateUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *JellyfinGroupStateUpdate) GetReason() JellyfinJellyfinPlaybackRequestType`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JellyfinGroupStateUpdate) GetReasonOk() (*JellyfinJellyfinPlaybackRequestType, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JellyfinGroupStateUpdate) SetReason(v JellyfinJellyfinPlaybackRequestType)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JellyfinGroupStateUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


