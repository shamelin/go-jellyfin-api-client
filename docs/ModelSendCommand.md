# ModelSendCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets the playlist identifier of the playing item. | [optional] 
**When** | Pointer to **time.Time** | Gets or sets the UTC time when to execute the command. | [optional] 
**PositionTicks** | Pointer to **NullableInt64** | Gets the position ticks. | [optional] 
**Command** | Pointer to [**ModelModelSendCommandType**](ModelSendCommandType.md) | Gets the command. | [optional] 
**EmittedAt** | Pointer to **time.Time** | Gets the UTC time when this command has been emitted. | [optional] 

## Methods

### NewModelSendCommand

`func NewModelSendCommand() *ModelSendCommand`

NewModelSendCommand instantiates a new ModelSendCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSendCommandWithDefaults

`func NewModelSendCommandWithDefaults() *ModelSendCommand`

NewModelSendCommandWithDefaults instantiates a new ModelSendCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelSendCommand) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelSendCommand) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelSendCommand) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelSendCommand) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelSendCommand) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelSendCommand) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelSendCommand) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelSendCommand) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetWhen

`func (o *ModelSendCommand) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ModelSendCommand) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ModelSendCommand) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *ModelSendCommand) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPositionTicks

`func (o *ModelSendCommand) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelSendCommand) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelSendCommand) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelSendCommand) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelSendCommand) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelSendCommand) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetCommand

`func (o *ModelSendCommand) GetCommand() ModelModelSendCommandType`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ModelSendCommand) GetCommandOk() (*ModelModelSendCommandType, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ModelSendCommand) SetCommand(v ModelModelSendCommandType)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ModelSendCommand) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEmittedAt

`func (o *ModelSendCommand) GetEmittedAt() time.Time`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *ModelSendCommand) GetEmittedAtOk() (*time.Time, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *ModelSendCommand) SetEmittedAt(v time.Time)`

SetEmittedAt sets EmittedAt field to given value.

### HasEmittedAt

`func (o *ModelSendCommand) HasEmittedAt() bool`

HasEmittedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


