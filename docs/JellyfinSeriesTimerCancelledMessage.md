# JellyfinSeriesTimerCancelledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableJellyfinJellyfinTimerEventInfo**](JellyfinTimerEventInfo.md) | Gets or sets the data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**JellyfinJellyfinSessionMessageType**](JellyfinSessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to JELLYFINJELLYFINSESSIONMESSAGETYPE_SERIES_TIMER_CANCELLED]

## Methods

### NewJellyfinSeriesTimerCancelledMessage

`func NewJellyfinSeriesTimerCancelledMessage() *JellyfinSeriesTimerCancelledMessage`

NewJellyfinSeriesTimerCancelledMessage instantiates a new JellyfinSeriesTimerCancelledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSeriesTimerCancelledMessageWithDefaults

`func NewJellyfinSeriesTimerCancelledMessageWithDefaults() *JellyfinSeriesTimerCancelledMessage`

NewJellyfinSeriesTimerCancelledMessageWithDefaults instantiates a new JellyfinSeriesTimerCancelledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JellyfinSeriesTimerCancelledMessage) GetData() JellyfinJellyfinTimerEventInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinSeriesTimerCancelledMessage) GetDataOk() (*JellyfinJellyfinTimerEventInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinSeriesTimerCancelledMessage) SetData(v JellyfinJellyfinTimerEventInfo)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinSeriesTimerCancelledMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JellyfinSeriesTimerCancelledMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JellyfinSeriesTimerCancelledMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *JellyfinSeriesTimerCancelledMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *JellyfinSeriesTimerCancelledMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *JellyfinSeriesTimerCancelledMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *JellyfinSeriesTimerCancelledMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *JellyfinSeriesTimerCancelledMessage) GetMessageType() JellyfinJellyfinSessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *JellyfinSeriesTimerCancelledMessage) GetMessageTypeOk() (*JellyfinJellyfinSessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *JellyfinSeriesTimerCancelledMessage) SetMessageType(v JellyfinJellyfinSessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *JellyfinSeriesTimerCancelledMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


