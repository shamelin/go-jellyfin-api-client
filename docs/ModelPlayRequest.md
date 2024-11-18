# ModelPlayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemIds** | Pointer to **[]string** | Gets or sets the item ids. | [optional] 
**StartPositionTicks** | Pointer to **NullableInt64** | Gets or sets the start position ticks that the first item should be played at. | [optional] 
**PlayCommand** | Pointer to [**ModelModelPlayCommand**](ModelPlayCommand.md) | Gets or sets the play command. | [optional] 
**ControllingUserId** | Pointer to **string** | Gets or sets the controlling user identifier. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MediaSourceId** | Pointer to **NullableString** |  | [optional] 
**StartIndex** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewModelPlayRequest

`func NewModelPlayRequest() *ModelPlayRequest`

NewModelPlayRequest instantiates a new ModelPlayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlayRequestWithDefaults

`func NewModelPlayRequestWithDefaults() *ModelPlayRequest`

NewModelPlayRequestWithDefaults instantiates a new ModelPlayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemIds

`func (o *ModelPlayRequest) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ModelPlayRequest) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ModelPlayRequest) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ModelPlayRequest) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### SetItemIdsNil

`func (o *ModelPlayRequest) SetItemIdsNil(b bool)`

 SetItemIdsNil sets the value for ItemIds to be an explicit nil

### UnsetItemIds
`func (o *ModelPlayRequest) UnsetItemIds()`

UnsetItemIds ensures that no value is present for ItemIds, not even an explicit nil
### GetStartPositionTicks

`func (o *ModelPlayRequest) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *ModelPlayRequest) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *ModelPlayRequest) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *ModelPlayRequest) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.

### SetStartPositionTicksNil

`func (o *ModelPlayRequest) SetStartPositionTicksNil(b bool)`

 SetStartPositionTicksNil sets the value for StartPositionTicks to be an explicit nil

### UnsetStartPositionTicks
`func (o *ModelPlayRequest) UnsetStartPositionTicks()`

UnsetStartPositionTicks ensures that no value is present for StartPositionTicks, not even an explicit nil
### GetPlayCommand

`func (o *ModelPlayRequest) GetPlayCommand() ModelModelPlayCommand`

GetPlayCommand returns the PlayCommand field if non-nil, zero value otherwise.

### GetPlayCommandOk

`func (o *ModelPlayRequest) GetPlayCommandOk() (*ModelModelPlayCommand, bool)`

GetPlayCommandOk returns a tuple with the PlayCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCommand

`func (o *ModelPlayRequest) SetPlayCommand(v ModelModelPlayCommand)`

SetPlayCommand sets PlayCommand field to given value.

### HasPlayCommand

`func (o *ModelPlayRequest) HasPlayCommand() bool`

HasPlayCommand returns a boolean if a field has been set.

### GetControllingUserId

`func (o *ModelPlayRequest) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *ModelPlayRequest) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *ModelPlayRequest) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *ModelPlayRequest) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### GetSubtitleStreamIndex

`func (o *ModelPlayRequest) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelPlayRequest) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelPlayRequest) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelPlayRequest) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelPlayRequest) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelPlayRequest) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetAudioStreamIndex

`func (o *ModelPlayRequest) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelPlayRequest) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelPlayRequest) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelPlayRequest) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelPlayRequest) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelPlayRequest) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetMediaSourceId

`func (o *ModelPlayRequest) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlayRequest) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlayRequest) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlayRequest) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *ModelPlayRequest) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *ModelPlayRequest) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetStartIndex

`func (o *ModelPlayRequest) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ModelPlayRequest) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ModelPlayRequest) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ModelPlayRequest) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### SetStartIndexNil

`func (o *ModelPlayRequest) SetStartIndexNil(b bool)`

 SetStartIndexNil sets the value for StartIndex to be an explicit nil

### UnsetStartIndex
`func (o *ModelPlayRequest) UnsetStartIndex()`

UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


