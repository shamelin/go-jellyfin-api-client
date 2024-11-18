# JellyfinScheduledTasksInfoMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]JellyfinJellyfinTaskInfo**](JellyfinJellyfinTaskInfo.md) | Gets or sets the data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SCHEDULED_TASKS_INFO]

## Methods

### NewJellyfinScheduledTasksInfoMessage

`func NewJellyfinScheduledTasksInfoMessage() *JellyfinScheduledTasksInfoMessage`

NewJellyfinScheduledTasksInfoMessage instantiates a new JellyfinScheduledTasksInfoMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinScheduledTasksInfoMessageWithDefaults

`func NewJellyfinScheduledTasksInfoMessageWithDefaults() *JellyfinScheduledTasksInfoMessage`

NewJellyfinScheduledTasksInfoMessageWithDefaults instantiates a new JellyfinScheduledTasksInfoMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinScheduledTasksInfoMessage) GetData() []JellyfinJellyfinTaskInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinScheduledTasksInfoMessage) GetDataOk() (*[]JellyfinJellyfinTaskInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinScheduledTasksInfoMessage) SetData(v []JellyfinJellyfinTaskInfo)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinScheduledTasksInfoMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinScheduledTasksInfoMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinScheduledTasksInfoMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinScheduledTasksInfoMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinScheduledTasksInfoMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinScheduledTasksInfoMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinScheduledTasksInfoMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinScheduledTasksInfoMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinScheduledTasksInfoMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinScheduledTasksInfoMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinScheduledTasksInfoMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


