# JellyfinUpdateUserItemDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **NullableFloat64** | Gets or sets the rating. | [optional] 
**PlayedPercentage** | Pointer to **NullableFloat64** | Gets or sets the played percentage. | [optional] 
**UnplayedItemCount** | Pointer to **NullableInt32** | Gets or sets the unplayed item count. | [optional] 
**PlaybackPositionTicks** | Pointer to **NullableInt64** | Gets or sets the playback position ticks. | [optional] 
**PlayCount** | Pointer to **NullableInt32** | Gets or sets the play count. | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is favorite. | [optional] 
**Likes** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UpdateUserItemDataDto is likes. | [optional] 
**LastPlayedDate** | Pointer to **NullableTime** | Gets or sets the last played date. | [optional] 
**Played** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 

## Methods

### NewJellyfinUpdateUserItemDataDto

`func NewJellyfinUpdateUserItemDataDto() *JellyfinUpdateUserItemDataDto`

NewJellyfinUpdateUserItemDataDto instantiates a new JellyfinUpdateUserItemDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUpdateUserItemDataDtoWithDefaults

`func NewJellyfinUpdateUserItemDataDtoWithDefaults() *JellyfinUpdateUserItemDataDto`

NewJellyfinUpdateUserItemDataDtoWithDefaults instantiates a new JellyfinUpdateUserItemDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *JellyfinUpdateUserItemDataDto) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *JellyfinUpdateUserItemDataDto) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *JellyfinUpdateUserItemDataDto) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *JellyfinUpdateUserItemDataDto) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *JellyfinUpdateUserItemDataDto) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *JellyfinUpdateUserItemDataDto) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetPlayedPercentage

`func (o *JellyfinUpdateUserItemDataDto) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *JellyfinUpdateUserItemDataDto) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *JellyfinUpdateUserItemDataDto) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *JellyfinUpdateUserItemDataDto) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *JellyfinUpdateUserItemDataDto) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *JellyfinUpdateUserItemDataDto) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetUnplayedItemCount

`func (o *JellyfinUpdateUserItemDataDto) GetUnplayedItemCount() int32`

GetUnplayedItemCount returns the UnplayedItemCount field if non-nil, zero value otherwise.

### GetUnplayedItemCountOk

`func (o *JellyfinUpdateUserItemDataDto) GetUnplayedItemCountOk() (*int32, bool)`

GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplayedItemCount

`func (o *JellyfinUpdateUserItemDataDto) SetUnplayedItemCount(v int32)`

SetUnplayedItemCount sets UnplayedItemCount field to given value.

### HasUnplayedItemCount

`func (o *JellyfinUpdateUserItemDataDto) HasUnplayedItemCount() bool`

HasUnplayedItemCount returns a boolean if a field has been set.

### SetUnplayedItemCountNil

`func (o *JellyfinUpdateUserItemDataDto) SetUnplayedItemCountNil(b bool)`

 SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil

### UnsetUnplayedItemCount
`func (o *JellyfinUpdateUserItemDataDto) UnsetUnplayedItemCount()`

UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
### GetPlaybackPositionTicks

`func (o *JellyfinUpdateUserItemDataDto) GetPlaybackPositionTicks() int64`

GetPlaybackPositionTicks returns the PlaybackPositionTicks field if non-nil, zero value otherwise.

### GetPlaybackPositionTicksOk

`func (o *JellyfinUpdateUserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool)`

GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPositionTicks

`func (o *JellyfinUpdateUserItemDataDto) SetPlaybackPositionTicks(v int64)`

SetPlaybackPositionTicks sets PlaybackPositionTicks field to given value.

### HasPlaybackPositionTicks

`func (o *JellyfinUpdateUserItemDataDto) HasPlaybackPositionTicks() bool`

HasPlaybackPositionTicks returns a boolean if a field has been set.

### SetPlaybackPositionTicksNil

`func (o *JellyfinUpdateUserItemDataDto) SetPlaybackPositionTicksNil(b bool)`

 SetPlaybackPositionTicksNil sets the value for PlaybackPositionTicks to be an explicit nil

### UnsetPlaybackPositionTicks
`func (o *JellyfinUpdateUserItemDataDto) UnsetPlaybackPositionTicks()`

UnsetPlaybackPositionTicks ensures that no value is present for PlaybackPositionTicks, not even an explicit nil
### GetPlayCount

`func (o *JellyfinUpdateUserItemDataDto) GetPlayCount() int32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *JellyfinUpdateUserItemDataDto) GetPlayCountOk() (*int32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *JellyfinUpdateUserItemDataDto) SetPlayCount(v int32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *JellyfinUpdateUserItemDataDto) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.

### SetPlayCountNil

`func (o *JellyfinUpdateUserItemDataDto) SetPlayCountNil(b bool)`

 SetPlayCountNil sets the value for PlayCount to be an explicit nil

### UnsetPlayCount
`func (o *JellyfinUpdateUserItemDataDto) UnsetPlayCount()`

UnsetPlayCount ensures that no value is present for PlayCount, not even an explicit nil
### GetIsFavorite

`func (o *JellyfinUpdateUserItemDataDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *JellyfinUpdateUserItemDataDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *JellyfinUpdateUserItemDataDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *JellyfinUpdateUserItemDataDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *JellyfinUpdateUserItemDataDto) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *JellyfinUpdateUserItemDataDto) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetLikes

`func (o *JellyfinUpdateUserItemDataDto) GetLikes() bool`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *JellyfinUpdateUserItemDataDto) GetLikesOk() (*bool, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *JellyfinUpdateUserItemDataDto) SetLikes(v bool)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *JellyfinUpdateUserItemDataDto) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### SetLikesNil

`func (o *JellyfinUpdateUserItemDataDto) SetLikesNil(b bool)`

 SetLikesNil sets the value for Likes to be an explicit nil

### UnsetLikes
`func (o *JellyfinUpdateUserItemDataDto) UnsetLikes()`

UnsetLikes ensures that no value is present for Likes, not even an explicit nil
### GetLastPlayedDate

`func (o *JellyfinUpdateUserItemDataDto) GetLastPlayedDate() time.Time`

GetLastPlayedDate returns the LastPlayedDate field if non-nil, zero value otherwise.

### GetLastPlayedDateOk

`func (o *JellyfinUpdateUserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool)`

GetLastPlayedDateOk returns a tuple with the LastPlayedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayedDate

`func (o *JellyfinUpdateUserItemDataDto) SetLastPlayedDate(v time.Time)`

SetLastPlayedDate sets LastPlayedDate field to given value.

### HasLastPlayedDate

`func (o *JellyfinUpdateUserItemDataDto) HasLastPlayedDate() bool`

HasLastPlayedDate returns a boolean if a field has been set.

### SetLastPlayedDateNil

`func (o *JellyfinUpdateUserItemDataDto) SetLastPlayedDateNil(b bool)`

 SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil

### UnsetLastPlayedDate
`func (o *JellyfinUpdateUserItemDataDto) UnsetLastPlayedDate()`

UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
### GetPlayed

`func (o *JellyfinUpdateUserItemDataDto) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *JellyfinUpdateUserItemDataDto) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *JellyfinUpdateUserItemDataDto) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *JellyfinUpdateUserItemDataDto) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### SetPlayedNil

`func (o *JellyfinUpdateUserItemDataDto) SetPlayedNil(b bool)`

 SetPlayedNil sets the value for Played to be an explicit nil

### UnsetPlayed
`func (o *JellyfinUpdateUserItemDataDto) UnsetPlayed()`

UnsetPlayed ensures that no value is present for Played, not even an explicit nil
### GetKey

`func (o *JellyfinUpdateUserItemDataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JellyfinUpdateUserItemDataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JellyfinUpdateUserItemDataDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *JellyfinUpdateUserItemDataDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *JellyfinUpdateUserItemDataDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *JellyfinUpdateUserItemDataDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetItemId

`func (o *JellyfinUpdateUserItemDataDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinUpdateUserItemDataDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinUpdateUserItemDataDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinUpdateUserItemDataDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *JellyfinUpdateUserItemDataDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *JellyfinUpdateUserItemDataDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


