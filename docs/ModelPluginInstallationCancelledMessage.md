# ModelPluginInstallationCancelledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableModelModelInstallationInfo**](ModelInstallationInfo.md) | Class InstallationInfo. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**ModelModelSessionMessageType**](ModelSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to MODELMODELSESSIONMESSAGETYPE_PACKAGE_INSTALLATION_CANCELLED]

## Methods

### NewModelPluginInstallationCancelledMessage

`func NewModelPluginInstallationCancelledMessage() *ModelPluginInstallationCancelledMessage`

NewModelPluginInstallationCancelledMessage instantiates a new ModelPluginInstallationCancelledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPluginInstallationCancelledMessageWithDefaults

`func NewModelPluginInstallationCancelledMessageWithDefaults() *ModelPluginInstallationCancelledMessage`

NewModelPluginInstallationCancelledMessageWithDefaults instantiates a new ModelPluginInstallationCancelledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelPluginInstallationCancelledMessage) GetData() ModelModelInstallationInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelPluginInstallationCancelledMessage) GetDataOk() (*ModelModelInstallationInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelPluginInstallationCancelledMessage) SetData(v ModelModelInstallationInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ModelPluginInstallationCancelledMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ModelPluginInstallationCancelledMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ModelPluginInstallationCancelledMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *ModelPluginInstallationCancelledMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ModelPluginInstallationCancelledMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ModelPluginInstallationCancelledMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ModelPluginInstallationCancelledMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *ModelPluginInstallationCancelledMessage) GetMessageType() ModelModelSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ModelPluginInstallationCancelledMessage) GetMessageTypeOk() (*ModelModelSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ModelPluginInstallationCancelledMessage) SetMessageType(v ModelModelSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ModelPluginInstallationCancelledMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


