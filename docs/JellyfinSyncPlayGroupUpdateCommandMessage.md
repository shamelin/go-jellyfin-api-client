# JellyfinSyncPlayGroupUpdateCommandMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableJellyfinJellyfinGroupUpdate**](JellyfinGroupUpdate.md) | Group update without data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SYNC_PLAY_GROUP_UPDATE]

## Methods

### NewJellyfinSyncPlayGroupUpdateCommandMessage

`func NewJellyfinSyncPlayGroupUpdateCommandMessage() *JellyfinSyncPlayGroupUpdateCommandMessage`

NewJellyfinSyncPlayGroupUpdateCommandMessage instantiates a new JellyfinSyncPlayGroupUpdateCommandMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSyncPlayGroupUpdateCommandMessageWithDefaults

`func NewJellyfinSyncPlayGroupUpdateCommandMessageWithDefaults() *JellyfinSyncPlayGroupUpdateCommandMessage`

NewJellyfinSyncPlayGroupUpdateCommandMessageWithDefaults instantiates a new JellyfinSyncPlayGroupUpdateCommandMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetData() JellyfinJellyfinGroupUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetDataOk() (*JellyfinJellyfinGroupUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) SetData(v JellyfinJellyfinGroupUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinSyncPlayGroupUpdateCommandMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


