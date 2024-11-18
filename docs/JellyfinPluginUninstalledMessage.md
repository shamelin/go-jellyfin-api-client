# JellyfinPluginUninstalledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableJellyfinJellyfinPluginInfo**](JellyfinPluginInfo.md) | This is a serializable stub class that is used by the api to provide information about installed plugins. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to JELLYFINJELLYFINSESSIONMESSAGETYPE_PACKAGE_UNINSTALLED]

## Methods

### NewJellyfinPluginUninstalledMessage

`func NewJellyfinPluginUninstalledMessage() *JellyfinPluginUninstalledMessage`

NewJellyfinPluginUninstalledMessage instantiates a new JellyfinPluginUninstalledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPluginUninstalledMessageWithDefaults

`func NewJellyfinPluginUninstalledMessageWithDefaults() *JellyfinPluginUninstalledMessage`

NewJellyfinPluginUninstalledMessageWithDefaults instantiates a new JellyfinPluginUninstalledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinPluginUninstalledMessage) GetData() JellyfinJellyfinPluginInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinPluginUninstalledMessage) GetDataOk() (*JellyfinJellyfinPluginInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinPluginUninstalledMessage) SetData(v JellyfinJellyfinPluginInfo)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinPluginUninstalledMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinPluginUninstalledMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinPluginUninstalledMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinPluginUninstalledMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinPluginUninstalledMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinPluginUninstalledMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinPluginUninstalledMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinPluginUninstalledMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinPluginUninstalledMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinPluginUninstalledMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinPluginUninstalledMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


