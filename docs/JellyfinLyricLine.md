# JellyfinLyricLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Gets the text of this lyric line. | [optional] 
**Start** | Pointer to **NullableInt64** | Gets the start time in ticks. | [optional] 

## Methods

### NewJellyfinLyricLine

`func NewJellyfinLyricLine() *JellyfinLyricLine`

NewJellyfinLyricLine instantiates a new JellyfinLyricLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLyricLineWithDefaults

`func NewJellyfinLyricLineWithDefaults() *JellyfinLyricLine`

NewJellyfinLyricLineWithDefaults instantiates a new JellyfinLyricLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *JellyfinLyricLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *JellyfinLyricLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *JellyfinLyricLine) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *JellyfinLyricLine) HasText() bool`

HasText returns a boolean if a field has been set.

### GetStart

`func (o *JellyfinLyricLine) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *JellyfinLyricLine) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *JellyfinLyricLine) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *JellyfinLyricLine) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *JellyfinLyricLine) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *JellyfinLyricLine) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


