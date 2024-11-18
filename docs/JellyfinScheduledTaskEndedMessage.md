# JellyfinScheduledTaskEndedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableJellyfinJellyfinTaskResult**](JellyfinTaskResult.md) | Class TaskExecutionInfo. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to JELLYFINJELLYFINSESSIONMESSAGETYPE_SCHEDULED_TASK_ENDED]

## Methods

### NewJellyfinScheduledTaskEndedMessage

`func NewJellyfinScheduledTaskEndedMessage() *JellyfinScheduledTaskEndedMessage`

NewJellyfinScheduledTaskEndedMessage instantiates a new JellyfinScheduledTaskEndedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinScheduledTaskEndedMessageWithDefaults

`func NewJellyfinScheduledTaskEndedMessageWithDefaults() *JellyfinScheduledTaskEndedMessage`

NewJellyfinScheduledTaskEndedMessageWithDefaults instantiates a new JellyfinScheduledTaskEndedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinScheduledTaskEndedMessage) GetData() JellyfinJellyfinTaskResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinScheduledTaskEndedMessage) GetDataOk() (*JellyfinJellyfinTaskResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinScheduledTaskEndedMessage) SetData(v JellyfinJellyfinTaskResult)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinScheduledTaskEndedMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinScheduledTaskEndedMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinScheduledTaskEndedMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinScheduledTaskEndedMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinScheduledTaskEndedMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinScheduledTaskEndedMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinScheduledTaskEndedMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinScheduledTaskEndedMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinScheduledTaskEndedMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinScheduledTaskEndedMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinScheduledTaskEndedMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


