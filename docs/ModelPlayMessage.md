# ModelPlayMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableModelModelPlayRequest**](ModelPlayRequest.md) | Class PlayRequest. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_PLAY]

## Methods

### NewModelPlayMessage

`func NewModelPlayMessage() *ModelPlayMessage`

NewModelPlayMessage instantiates a new ModelPlayMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlayMessageWithDefaults

`func NewModelPlayMessageWithDefaults() *ModelPlayMessage`

NewModelPlayMessageWithDefaults instantiates a new ModelPlayMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelPlayMessage) GetData() ModelModelPlayRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelPlayMessage) GetDataOk() (*ModelModelPlayRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelPlayMessage) SetData(v ModelModelPlayRequest)`

SetData sets Data field to given value.

### HasData

`func (o *ModelPlayMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ModelPlayMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ModelPlayMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *ModelPlayMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelPlayMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelPlayMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelPlayMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelPlayMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelPlayMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelPlayMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelPlayMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


