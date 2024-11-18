# ModelActivityLogEntryStopMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_ACTIVITY_LOG_ENTRY_STOP]

## Methods

### NewModelActivityLogEntryStopMessage

`func NewModelActivityLogEntryStopMessage() *ModelActivityLogEntryStopMessage`

NewModelActivityLogEntryStopMessage instantiates a new ModelActivityLogEntryStopMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelActivityLogEntryStopMessageWithDefaults

`func NewModelActivityLogEntryStopMessageWithDefaults() *ModelActivityLogEntryStopMessage`

NewModelActivityLogEntryStopMessageWithDefaults instantiates a new ModelActivityLogEntryStopMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *ModelActivityLogEntryStopMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelActivityLogEntryStopMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelActivityLogEntryStopMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelActivityLogEntryStopMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

