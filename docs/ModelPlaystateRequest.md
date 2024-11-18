# ModelPlaystateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to [**ModelModelPlaystateCommand**](ModelPlaystateCommand.md) | Enum PlaystateCommand. | [optional] 
**SeekPositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**ControllingUserId** | Pointer to **NullableString** | Gets or sets the controlling user identifier. | [optional] 

## Methods

### NewModelPlaystateRequest

`func NewModelPlaystateRequest() *ModelPlaystateRequest`

NewModelPlaystateRequest instantiates a new ModelPlaystateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaystateRequestWithDefaults

`func NewModelPlaystateRequestWithDefaults() *ModelPlaystateRequest`

NewModelPlaystateRequestWithDefaults instantiates a new ModelPlaystateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ModelPlaystateRequest) GetCommand() ModelModelPlaystateCommand`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ModelPlaystateRequest) GetCommandOk() (*ModelModelPlaystateCommand, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ModelPlaystateRequest) SetCommand(v ModelModelPlaystateCommand)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ModelPlaystateRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSeekPositionTicks

`func (o *ModelPlaystateRequest) GetSeekPositionTicks() int64`

GetSeekPositionTicks returns the SeekPositionTicks field if non-nil, zero value otherwise.

### GetSeekPositionTicksOk

`func (o *ModelPlaystateRequest) GetSeekPositionTicksOk() (*int64, bool)`

GetSeekPositionTicksOk returns a tuple with the SeekPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeekPositionTicks

`func (o *ModelPlaystateRequest) SetSeekPositionTicks(v int64)`

SetSeekPositionTicks sets SeekPositionTicks field to given value.

### HasSeekPositionTicks

`func (o *ModelPlaystateRequest) HasSeekPositionTicks() bool`

HasSeekPositionTicks returns a boolean if a field has been set.

### SetSeekPositionTicksNil

`func (o *ModelPlaystateRequest) SetSeekPositionTicksNil(b bool)`

 SetSeekPositionTicksNil sets the value for SeekPositionTicks to be an explicit nil

### UnsetSeekPositionTicks
`func (o *ModelPlaystateRequest) UnsetSeekPositionTicks()`

UnsetSeekPositionTicks ensures that no value is present for SeekPositionTicks, not even an explicit nil
### GetControllingUserId

`func (o *ModelPlaystateRequest) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *ModelPlaystateRequest) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *ModelPlaystateRequest) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *ModelPlaystateRequest) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### SetControllingUserIdNil

`func (o *ModelPlaystateRequest) SetControllingUserIdNil(b bool)`

 SetControllingUserIdNil sets the value for ControllingUserId to be an explicit nil

### UnsetControllingUserId
`func (o *ModelPlaystateRequest) UnsetControllingUserId()`

UnsetControllingUserId ensures that no value is present for ControllingUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


