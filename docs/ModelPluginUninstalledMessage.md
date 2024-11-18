# ModelPluginUninstalledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableModelModelPluginInfo**](ModelPluginInfo.md) | This is a serializable stub class that is used by the api to provide information about installed plugins. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_PACKAGE_UNINSTALLED]

## Methods

### NewModelPluginUninstalledMessage

`func NewModelPluginUninstalledMessage() *ModelPluginUninstalledMessage`

NewModelPluginUninstalledMessage instantiates a new ModelPluginUninstalledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPluginUninstalledMessageWithDefaults

`func NewModelPluginUninstalledMessageWithDefaults() *ModelPluginUninstalledMessage`

NewModelPluginUninstalledMessageWithDefaults instantiates a new ModelPluginUninstalledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelPluginUninstalledMessage) GetData() ModelModelPluginInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelPluginUninstalledMessage) GetDataOk() (*ModelModelPluginInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelPluginUninstalledMessage) SetData(v ModelModelPluginInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ModelPluginUninstalledMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ModelPluginUninstalledMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ModelPluginUninstalledMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *ModelPluginUninstalledMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelPluginUninstalledMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelPluginUninstalledMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelPluginUninstalledMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelPluginUninstalledMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelPluginUninstalledMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelPluginUninstalledMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelPluginUninstalledMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


