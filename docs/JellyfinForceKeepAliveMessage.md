# JellyfinForceKeepAliveMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **int32** | Gets or sets the data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to JELLYFINJELLYFINSESSIONMESSAGETYPE_FORCE_KEEP_ALIVE]

## Methods

### NewJellyfinForceKeepAliveMessage

`func NewJellyfinForceKeepAliveMessage() *JellyfinForceKeepAliveMessage`

NewJellyfinForceKeepAliveMessage instantiates a new JellyfinForceKeepAliveMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinForceKeepAliveMessageWithDefaults

`func NewJellyfinForceKeepAliveMessageWithDefaults() *JellyfinForceKeepAliveMessage`

NewJellyfinForceKeepAliveMessageWithDefaults instantiates a new JellyfinForceKeepAliveMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinForceKeepAliveMessage) GetData() int32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinForceKeepAliveMessage) GetDataOk() (*int32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinForceKeepAliveMessage) SetData(v int32)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinForceKeepAliveMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessageId

`func (o *JellyfinForceKeepAliveMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinForceKeepAliveMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinForceKeepAliveMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinForceKeepAliveMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinForceKeepAliveMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinForceKeepAliveMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinForceKeepAliveMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinForceKeepAliveMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


