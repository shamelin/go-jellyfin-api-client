# ModelPlayQueueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to [**ModelModelPlayQueueUpdateReason**](ModelPlayQueueUpdateReason.md) | Gets the request type that originated this update. | [optional] 
**LastUpdate** | Pointer to **time.Time** | Gets the UTC time of the last change to the playing queue. | [optional] 
**Playlist** | Pointer to [**[]ModelModelSyncPlayQueueItem**](ModelModelSyncPlayQueueItem.md) | Gets the playlist. | [optional] 
**PlayingItemIndex** | Pointer to **int32** | Gets the playing item index in the playlist. | [optional] 
**StartPositionTicks** | Pointer to **int64** | Gets the start position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets a value indicating whether the current item is playing. | [optional] 
**ShuffleMode** | Pointer to [**ModelModelGroupShuffleMode**](ModelGroupShuffleMode.md) | Gets the shuffle mode. | [optional] 
**RepeatMode** | Pointer to [**ModelModelGroupRepeatMode**](ModelGroupRepeatMode.md) | Gets the repeat mode. | [optional] 

## Methods

### NewModelPlayQueueUpdate

`func NewModelPlayQueueUpdate() *ModelPlayQueueUpdate`

NewModelPlayQueueUpdate instantiates a new ModelPlayQueueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlayQueueUpdateWithDefaults

`func NewModelPlayQueueUpdateWithDefaults() *ModelPlayQueueUpdate`

NewModelPlayQueueUpdateWithDefaults instantiates a new ModelPlayQueueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ModelPlayQueueUpdate) GetReason() ModelModelPlayQueueUpdateReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelPlayQueueUpdate) GetReasonOk() (*ModelModelPlayQueueUpdateReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelPlayQueueUpdate) SetReason(v ModelModelPlayQueueUpdateReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ModelPlayQueueUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ModelPlayQueueUpdate) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ModelPlayQueueUpdate) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ModelPlayQueueUpdate) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ModelPlayQueueUpdate) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetPlaylist

`func (o *ModelPlayQueueUpdate) GetPlaylist() []ModelModelSyncPlayQueueItem`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *ModelPlayQueueUpdate) GetPlaylistOk() (*[]ModelModelSyncPlayQueueItem, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *ModelPlayQueueUpdate) SetPlaylist(v []ModelModelSyncPlayQueueItem)`

SetPlaylist sets Playlist field to given value.

### HasPlaylist

`func (o *ModelPlayQueueUpdate) HasPlaylist() bool`

HasPlaylist returns a boolean if a field has been set.

### GetPlayingItemIndex

`func (o *ModelPlayQueueUpdate) GetPlayingItemIndex() int32`

GetPlayingItemIndex returns the PlayingItemIndex field if non-nil, zero value otherwise.

### GetPlayingItemIndexOk

`func (o *ModelPlayQueueUpdate) GetPlayingItemIndexOk() (*int32, bool)`

GetPlayingItemIndexOk returns a tuple with the PlayingItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingItemIndex

`func (o *ModelPlayQueueUpdate) SetPlayingItemIndex(v int32)`

SetPlayingItemIndex sets PlayingItemIndex field to given value.

### HasPlayingItemIndex

`func (o *ModelPlayQueueUpdate) HasPlayingItemIndex() bool`

HasPlayingItemIndex returns a boolean if a field has been set.

### GetStartPositionTicks

`func (o *ModelPlayQueueUpdate) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *ModelPlayQueueUpdate) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *ModelPlayQueueUpdate) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *ModelPlayQueueUpdate) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *ModelPlayQueueUpdate) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *ModelPlayQueueUpdate) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *ModelPlayQueueUpdate) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *ModelPlayQueueUpdate) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetShuffleMode

`func (o *ModelPlayQueueUpdate) GetShuffleMode() ModelModelGroupShuffleMode`

GetShuffleMode returns the ShuffleMode field if non-nil, zero value otherwise.

### GetShuffleModeOk

`func (o *ModelPlayQueueUpdate) GetShuffleModeOk() (*ModelModelGroupShuffleMode, bool)`

GetShuffleModeOk returns a tuple with the ShuffleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffleMode

`func (o *ModelPlayQueueUpdate) SetShuffleMode(v ModelModelGroupShuffleMode)`

SetShuffleMode sets ShuffleMode field to given value.

### HasShuffleMode

`func (o *ModelPlayQueueUpdate) HasShuffleMode() bool`

HasShuffleMode returns a boolean if a field has been set.

### GetRepeatMode

`func (o *ModelPlayQueueUpdate) GetRepeatMode() ModelModelGroupRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *ModelPlayQueueUpdate) GetRepeatModeOk() (*ModelModelGroupRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *ModelPlayQueueUpdate) SetRepeatMode(v ModelModelGroupRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *ModelPlayQueueUpdate) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


