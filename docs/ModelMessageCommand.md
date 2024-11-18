# ModelMessageCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **NullableString** |  | [optional] 
**Text** | **string** |  | 
**TimeoutMs** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewModelMessageCommand

`func NewModelMessageCommand(text string, ) *ModelMessageCommand`

NewModelMessageCommand instantiates a new ModelMessageCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMessageCommandWithDefaults

`func NewModelMessageCommandWithDefaults() *ModelMessageCommand`

NewModelMessageCommandWithDefaults instantiates a new ModelMessageCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *ModelMessageCommand) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ModelMessageCommand) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ModelMessageCommand) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ModelMessageCommand) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ModelMessageCommand) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ModelMessageCommand) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetText

`func (o *ModelMessageCommand) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelMessageCommand) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelMessageCommand) SetText(v string)`

SetText sets Text field to given value.


### GetTimeoutMs

`func (o *ModelMessageCommand) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ModelMessageCommand) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ModelMessageCommand) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *ModelMessageCommand) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### SetTimeoutMsNil

`func (o *ModelMessageCommand) SetTimeoutMsNil(b bool)`

 SetTimeoutMsNil sets the value for TimeoutMs to be an explicit nil

### UnsetTimeoutMs
`func (o *ModelMessageCommand) UnsetTimeoutMs()`

UnsetTimeoutMs ensures that no value is present for TimeoutMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


