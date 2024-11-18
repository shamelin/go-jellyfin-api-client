# JellyfinPlayQueueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to [**JellyfinJellyfinPlayQueueUpdateReason**](JellyfinPlayQueueUpdateReason.md) | Gets the request type that originated this update. | [optional] 
**LastUpdate** | Pointer to **time.Time** | Gets the UTC time of the last change to the playing queue. | [optional] 
**Playlist** | Pointer to [**[]JellyfinJellyfinSyncPlayQueueItem**](JellyfinJellyfinSyncPlayQueueItem.md) | Gets the playlist. | [optional] 
**PlayingItemIndex** | Pointer to **int32** | Gets the playing item index in the playlist. | [optional] 
**StartPositionTicks** | Pointer to **int64** | Gets the start position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets a value indicating whether the current item is playing. | [optional] 
**ShuffleMode** | Pointer to [**JellyfinJellyfinGroupShuffleMode**](JellyfinGroupShuffleMode.md) | Gets the shuffle mode. | [optional] 
**RepeatMode** | Pointer to [**JellyfinJellyfinGroupRepeatMode**](JellyfinGroupRepeatMode.md) | Gets the repeat mode. | [optional] 

## Methods

### NewJellyfinPlayQueueUpdate

`func NewJellyfinPlayQueueUpdate() *JellyfinPlayQueueUpdate`

NewJellyfinPlayQueueUpdate instantiates a new JellyfinPlayQueueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPlayQueueUpdateWithDefaults

`func NewJellyfinPlayQueueUpdateWithDefaults() *JellyfinPlayQueueUpdate`

NewJellyfinPlayQueueUpdateWithDefaults instantiates a new JellyfinPlayQueueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *JellyfinPlayQueueUpdate) GetReason() JellyfinJellyfinPlayQueueUpdateReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JellyfinPlayQueueUpdate) GetReasonOk() (*JellyfinJellyfinPlayQueueUpdateReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JellyfinPlayQueueUpdate) SetReason(v JellyfinJellyfinPlayQueueUpdateReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JellyfinPlayQueueUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLastUpdate

`func (o *JellyfinPlayQueueUpdate) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *JellyfinPlayQueueUpdate) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *JellyfinPlayQueueUpdate) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *JellyfinPlayQueueUpdate) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetPlaylist

`func (o *JellyfinPlayQueueUpdate) GetPlaylist() []JellyfinJellyfinSyncPlayQueueItem`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *JellyfinPlayQueueUpdate) GetPlaylistOk() (*[]JellyfinJellyfinSyncPlayQueueItem, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *JellyfinPlayQueueUpdate) SetPlaylist(v []JellyfinJellyfinSyncPlayQueueItem)`

SetPlaylist sets Playlist field to given value.

### HasPlaylist

`func (o *JellyfinPlayQueueUpdate) HasPlaylist() bool`

HasPlaylist returns a boolean if a field has been set.

### GetPlayingItemIndex

`func (o *JellyfinPlayQueueUpdate) GetPlayingItemIndex() int32`

GetPlayingItemIndex returns the PlayingItemIndex field if non-nil, zero value otherwise.

### GetPlayingItemIndexOk

`func (o *JellyfinPlayQueueUpdate) GetPlayingItemIndexOk() (*int32, bool)`

GetPlayingItemIndexOk returns a tuple with the PlayingItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingItemIndex

`func (o *JellyfinPlayQueueUpdate) SetPlayingItemIndex(v int32)`

SetPlayingItemIndex sets PlayingItemIndex field to given value.

### HasPlayingItemIndex

`func (o *JellyfinPlayQueueUpdate) HasPlayingItemIndex() bool`

HasPlayingItemIndex returns a boolean if a field has been set.

### GetStartPositionTicks

`func (o *JellyfinPlayQueueUpdate) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *JellyfinPlayQueueUpdate) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *JellyfinPlayQueueUpdate) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *JellyfinPlayQueueUpdate) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *JellyfinPlayQueueUpdate) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *JellyfinPlayQueueUpdate) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *JellyfinPlayQueueUpdate) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *JellyfinPlayQueueUpdate) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetShuffleMode

`func (o *JellyfinPlayQueueUpdate) GetShuffleMode() JellyfinJellyfinGroupShuffleMode`

GetShuffleMode returns the ShuffleMode field if non-nil, zero value otherwise.

### GetShuffleModeOk

`func (o *JellyfinPlayQueueUpdate) GetShuffleModeOk() (*JellyfinJellyfinGroupShuffleMode, bool)`

GetShuffleModeOk returns a tuple with the ShuffleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffleMode

`func (o *JellyfinPlayQueueUpdate) SetShuffleMode(v JellyfinJellyfinGroupShuffleMode)`

SetShuffleMode sets ShuffleMode field to given value.

### HasShuffleMode

`func (o *JellyfinPlayQueueUpdate) HasShuffleMode() bool`

HasShuffleMode returns a boolean if a field has been set.

### GetRepeatMode

`func (o *JellyfinPlayQueueUpdate) GetRepeatMode() JellyfinJellyfinGroupRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *JellyfinPlayQueueUpdate) GetRepeatModeOk() (*JellyfinJellyfinGroupRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *JellyfinPlayQueueUpdate) SetRepeatMode(v JellyfinJellyfinGroupRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *JellyfinPlayQueueUpdate) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


