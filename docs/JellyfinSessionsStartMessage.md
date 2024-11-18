# JellyfinSessionsStartMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONS_START]

## Methods

### NewJellyfinSessionsStartMessage

`func NewJellyfinSessionsStartMessage() *JellyfinSessionsStartMessage`

NewJellyfinSessionsStartMessage instantiates a new JellyfinSessionsStartMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSessionsStartMessageWithDefaults

`func NewJellyfinSessionsStartMessageWithDefaults() *JellyfinSessionsStartMessage`

NewJellyfinSessionsStartMessageWithDefaults instantiates a new JellyfinSessionsStartMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinSessionsStartMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinSessionsStartMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinSessionsStartMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinSessionsStartMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinSessionsStartMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinSessionsStartMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *JellyfinSessionsStartMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinSessionsStartMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinSessionsStartMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinSessionsStartMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


