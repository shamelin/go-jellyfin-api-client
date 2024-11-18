# ModelInboundKeepAliveMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_KEEP_ALIVE]

## Methods

### NewModelInboundKeepAliveMessage

`func NewModelInboundKeepAliveMessage() *ModelInboundKeepAliveMessage`

NewModelInboundKeepAliveMessage instantiates a new ModelInboundKeepAliveMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInboundKeepAliveMessageWithDefaults

`func NewModelInboundKeepAliveMessageWithDefaults() *ModelInboundKeepAliveMessage`

NewModelInboundKeepAliveMessageWithDefaults instantiates a new ModelInboundKeepAliveMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *ModelInboundKeepAliveMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelInboundKeepAliveMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelInboundKeepAliveMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelInboundKeepAliveMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


