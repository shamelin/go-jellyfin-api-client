# ModelOutboundKeepAliveMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_KEEP_ALIVE]

## Methods

### NewModelOutboundKeepAliveMessage

`func NewModelOutboundKeepAliveMessage() *ModelOutboundKeepAliveMessage`

NewModelOutboundKeepAliveMessage instantiates a new ModelOutboundKeepAliveMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelOutboundKeepAliveMessageWithDefaults

`func NewModelOutboundKeepAliveMessageWithDefaults() *ModelOutboundKeepAliveMessage`

NewModelOutboundKeepAliveMessageWithDefaults instantiates a new ModelOutboundKeepAliveMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *ModelOutboundKeepAliveMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelOutboundKeepAliveMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelOutboundKeepAliveMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelOutboundKeepAliveMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelOutboundKeepAliveMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelOutboundKeepAliveMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelOutboundKeepAliveMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelOutboundKeepAliveMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

