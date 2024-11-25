# JellyfinRestartRequiredMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to JELLYFINJELLYFINSESSIONMESSAGETYPE_RESTART_REQUIRED]

## Methods

### NewJellyfinRestartRequiredMessage

`func NewJellyfinRestartRequiredMessage() *JellyfinRestartRequiredMessage`

NewJellyfinRestartRequiredMessage instantiates a new JellyfinRestartRequiredMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinRestartRequiredMessageWithDefaults

`func NewJellyfinRestartRequiredMessageWithDefaults() *JellyfinRestartRequiredMessage`

NewJellyfinRestartRequiredMessageWithDefaults instantiates a new JellyfinRestartRequiredMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *JellyfinRestartRequiredMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinRestartRequiredMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinRestartRequiredMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinRestartRequiredMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinRestartRequiredMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinRestartRequiredMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinRestartRequiredMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinRestartRequiredMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


