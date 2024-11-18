# JellyfinScheduledTasksInfoStartMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SCHEDULED_TASKS_INFO_START]

## Methods

### NewJellyfinScheduledTasksInfoStartMessage

`func NewJellyfinScheduledTasksInfoStartMessage() *JellyfinScheduledTasksInfoStartMessage`

NewJellyfinScheduledTasksInfoStartMessage instantiates a new JellyfinScheduledTasksInfoStartMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinScheduledTasksInfoStartMessageWithDefaults

`func NewJellyfinScheduledTasksInfoStartMessageWithDefaults() *JellyfinScheduledTasksInfoStartMessage`

NewJellyfinScheduledTasksInfoStartMessageWithDefaults instantiates a new JellyfinScheduledTasksInfoStartMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinScheduledTasksInfoStartMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinScheduledTasksInfoStartMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinScheduledTasksInfoStartMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinScheduledTasksInfoStartMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinScheduledTasksInfoStartMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinScheduledTasksInfoStartMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *JellyfinScheduledTasksInfoStartMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinScheduledTasksInfoStartMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinScheduledTasksInfoStartMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinScheduledTasksInfoStartMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


