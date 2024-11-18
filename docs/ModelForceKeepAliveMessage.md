# ModelForceKeepAliveMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **int32** | Gets or sets the data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_FORCE_KEEP_ALIVE]

## Methods

### NewModelForceKeepAliveMessage

`func NewModelForceKeepAliveMessage() *ModelForceKeepAliveMessage`

NewModelForceKeepAliveMessage instantiates a new ModelForceKeepAliveMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelForceKeepAliveMessageWithDefaults

`func NewModelForceKeepAliveMessageWithDefaults() *ModelForceKeepAliveMessage`

NewModelForceKeepAliveMessageWithDefaults instantiates a new ModelForceKeepAliveMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelForceKeepAliveMessage) GetData() int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelForceKeepAliveMessage) GetDataOk() (*int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelForceKeepAliveMessage) SetData(v int32)`

SetData sets Data field to given value.

### HasData

`func (o *ModelForceKeepAliveMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessageId

`func (o *ModelForceKeepAliveMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelForceKeepAliveMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelForceKeepAliveMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelForceKeepAliveMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelForceKeepAliveMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelForceKeepAliveMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelForceKeepAliveMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelForceKeepAliveMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


