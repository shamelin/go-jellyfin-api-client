# ModelLyricLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Gets the text of this lyric line. | [optional] 
**Start** | Pointer to **NullableInt64** | Gets the start time in ticks. | [optional] 

## Methods

### NewModelLyricLine

`func NewModelLyricLine() *ModelLyricLine`

NewModelLyricLine instantiates a new ModelLyricLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLyricLineWithDefaults

`func NewModelLyricLineWithDefaults() *ModelLyricLine`

NewModelLyricLineWithDefaults instantiates a new ModelLyricLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ModelLyricLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ModelLyricLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ModelLyricLine) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ModelLyricLine) HasText() bool`

HasText returns a boolean if a field has been set.

### GetStart

`func (o *ModelLyricLine) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ModelLyricLine) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ModelLyricLine) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ModelLyricLine) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *ModelLyricLine) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *ModelLyricLine) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


