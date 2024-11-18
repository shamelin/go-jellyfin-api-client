# JellyfinPluginInstallingMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableJellyfinJellyfinInstallationInfo**](JellyfinInstallationInfo.md) | Class InstallationInfo. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to PACKAGE_INSTALLING]

## Methods

### NewJellyfinPluginInstallingMessage

`func NewJellyfinPluginInstallingMessage() *JellyfinPluginInstallingMessage`

NewJellyfinPluginInstallingMessage instantiates a new JellyfinPluginInstallingMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPluginInstallingMessageWithDefaults

`func NewJellyfinPluginInstallingMessageWithDefaults() *JellyfinPluginInstallingMessage`

NewJellyfinPluginInstallingMessageWithDefaults instantiates a new JellyfinPluginInstallingMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinPluginInstallingMessage) GetData() JellyfinJellyfinInstallationInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinPluginInstallingMessage) GetDataOk() (*JellyfinJellyfinInstallationInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinPluginInstallingMessage) SetData(v JellyfinJellyfinInstallationInfo)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinPluginInstallingMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinPluginInstallingMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinPluginInstallingMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinPluginInstallingMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinPluginInstallingMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinPluginInstallingMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinPluginInstallingMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinPluginInstallingMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinPluginInstallingMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinPluginInstallingMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinPluginInstallingMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


