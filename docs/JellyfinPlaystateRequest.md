# JellyfinPlaystateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to [**JellyfinJellyfinPlaystateCommand**](JellyfinPlaystateCommand.md) | Enum PlaystateCommand. | [optional] 
**SeekPositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**ControllingUserId** | Pointer to **NullableString** | Gets or sets the controlling user identifier. | [optional] 

## Methods

### NewJellyfinPlaystateRequest

`func NewJellyfinPlaystateRequest() *JellyfinPlaystateRequest`

NewJellyfinPlaystateRequest instantiates a new JellyfinPlaystateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlaystateRequestWithDefaults

`func NewJellyfinPlaystateRequestWithDefaults() *JellyfinPlaystateRequest`

NewJellyfinPlaystateRequestWithDefaults instantiates a new JellyfinPlaystateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *JellyfinPlaystateRequest) GetCommand() JellyfinJellyfinPlaystateCommand`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *JellyfinPlaystateRequest) GetCommandOk() (*JellyfinJellyfinPlaystateCommand, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *JellyfinPlaystateRequest) SetCommand(v JellyfinJellyfinPlaystateCommand)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *JellyfinPlaystateRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSeekPositionTicks

`func (o *JellyfinPlaystateRequest) GetSeekPositionTicks() int64`

GetSeekPositionTicks returns the SeekPositionTicks field if non-nil, zero value otherwise.

### GetSeekPositionTicksOk

`func (o *JellyfinPlaystateRequest) GetSeekPositionTicksOk() (*int64, bool)`

GetSeekPositionTicksOk returns a tuple with the SeekPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeekPositionTicks

`func (o *JellyfinPlaystateRequest) SetSeekPositionTicks(v int64)`

SetSeekPositionTicks sets SeekPositionTicks field to given value.

### HasSeekPositionTicks

`func (o *JellyfinPlaystateRequest) HasSeekPositionTicks() bool`

HasSeekPositionTicks returns a boolean if a field has been set.

### SetSeekPositionTicksNil

`func (o *JellyfinPlaystateRequest) SetSeekPositionTicksNil(b bool)`

 SetSeekPositionTicksNil sets the value for SeekPositionTicks to be an explicit nil

### UnsetSeekPositionTicks
`func (o *JellyfinPlaystateRequest) UnsetSeekPositionTicks()`

UnsetSeekPositionTicks ensures that no value is present for SeekPositionTicks, not even an explicit nil
### GetControllingUserId

`func (o *JellyfinPlaystateRequest) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *JellyfinPlaystateRequest) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *JellyfinPlaystateRequest) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *JellyfinPlaystateRequest) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### SetControllingUserIdNil

`func (o *JellyfinPlaystateRequest) SetControllingUserIdNil(b bool)`

 SetControllingUserIdNil sets the value for ControllingUserId to be an explicit nil

### UnsetControllingUserId
`func (o *JellyfinPlaystateRequest) UnsetControllingUserId()`

UnsetControllingUserId ensures that no value is present for ControllingUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


