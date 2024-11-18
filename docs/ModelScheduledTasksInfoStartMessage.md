# ModelScheduledTasksInfoStartMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO_START]

## Methods

### NewModelScheduledTasksInfoStartMessage

`func NewModelScheduledTasksInfoStartMessage() *ModelScheduledTasksInfoStartMessage`

NewModelScheduledTasksInfoStartMessage instantiates a new ModelScheduledTasksInfoStartMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScheduledTasksInfoStartMessageWithDefaults

`func NewModelScheduledTasksInfoStartMessageWithDefaults() *ModelScheduledTasksInfoStartMessage`

NewModelScheduledTasksInfoStartMessageWithDefaults instantiates a new ModelScheduledTasksInfoStartMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelScheduledTasksInfoStartMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelScheduledTasksInfoStartMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelScheduledTasksInfoStartMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ModelScheduledTasksInfoStartMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ModelScheduledTasksInfoStartMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ModelScheduledTasksInfoStartMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *ModelScheduledTasksInfoStartMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelScheduledTasksInfoStartMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelScheduledTasksInfoStartMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelScheduledTasksInfoStartMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


