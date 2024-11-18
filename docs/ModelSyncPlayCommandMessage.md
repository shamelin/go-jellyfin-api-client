# ModelSyncPlayCommandMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableModelModelSendCommand**](ModelSendCommand.md) | Class SendCommand. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_SYNC_PLAY_COMMAND]

## Methods

### NewModelSyncPlayCommandMessage

`func NewModelSyncPlayCommandMessage() *ModelSyncPlayCommandMessage`

NewModelSyncPlayCommandMessage instantiates a new ModelSyncPlayCommandMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncPlayCommandMessageWithDefaults

`func NewModelSyncPlayCommandMessageWithDefaults() *ModelSyncPlayCommandMessage`

NewModelSyncPlayCommandMessageWithDefaults instantiates a new ModelSyncPlayCommandMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelSyncPlayCommandMessage) GetData() ModelModelSendCommand`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelSyncPlayCommandMessage) GetDataOk() (*ModelModelSendCommand, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelSyncPlayCommandMessage) SetData(v ModelModelSendCommand)`

SetData sets Data field to given value.

### HasData

`func (o *ModelSyncPlayCommandMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ModelSyncPlayCommandMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ModelSyncPlayCommandMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *ModelSyncPlayCommandMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelSyncPlayCommandMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelSyncPlayCommandMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelSyncPlayCommandMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelSyncPlayCommandMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelSyncPlayCommandMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelSyncPlayCommandMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelSyncPlayCommandMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


