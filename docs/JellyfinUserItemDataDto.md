# JellyfinUserItemDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **NullableFloat64** | Gets or sets the rating. | [optional] 
**PlayedPercentage** | Pointer to **NullableFloat64** | Gets or sets the played percentage. | [optional] 
**UnplayedItemCount** | Pointer to **NullableInt32** | Gets or sets the unplayed item count. | [optional] 
**PlaybackPositionTicks** | Pointer to **int64** | Gets or sets the playback position ticks. | [optional] 
**PlayCount** | Pointer to **int32** | Gets or sets the play count. | [optional] 
**IsFavorite** | Pointer to **bool** | Gets or sets a value indicating whether this instance is favorite. | [optional] 
**Likes** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is likes. | [optional] 
**LastPlayedDate** | Pointer to **NullableTime** | Gets or sets the last played date. | [optional] 
**Played** | Pointer to **bool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played. | [optional] 
**Key** | Pointer to **string** | Gets or sets the key. | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the item identifier. | [optional] 

## Methods

### NewJellyfinUserItemDataDto

`func NewJellyfinUserItemDataDto() *JellyfinUserItemDataDto`

NewJellyfinUserItemDataDto instantiates a new JellyfinUserItemDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUserItemDataDtoWithDefaults

`func NewJellyfinUserItemDataDtoWithDefaults() *JellyfinUserItemDataDto`

NewJellyfinUserItemDataDtoWithDefaults instantiates a new JellyfinUserItemDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *JellyfinUserItemDataDto) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *JellyfinUserItemDataDto) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *JellyfinUserItemDataDto) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *JellyfinUserItemDataDto) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *JellyfinUserItemDataDto) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *JellyfinUserItemDataDto) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetPlayedPercentage

`func (o *JellyfinUserItemDataDto) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *JellyfinUserItemDataDto) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *JellyfinUserItemDataDto) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *JellyfinUserItemDataDto) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *JellyfinUserItemDataDto) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *JellyfinUserItemDataDto) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetUnplayedItemCount

`func (o *JellyfinUserItemDataDto) GetUnplayedItemCount() int32`

GetUnplayedItemCount returns the UnplayedItemCount field if non-nil, zero value otherwise.

### GetUnplayedItemCountOk

`func (o *JellyfinUserItemDataDto) GetUnplayedItemCountOk() (*int32, bool)`

GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplayedItemCount

`func (o *JellyfinUserItemDataDto) SetUnplayedItemCount(v int32)`

SetUnplayedItemCount sets UnplayedItemCount field to given value.

### HasUnplayedItemCount

`func (o *JellyfinUserItemDataDto) HasUnplayedItemCount() bool`

HasUnplayedItemCount returns a boolean if a field has been set.

### SetUnplayedItemCountNil

`func (o *JellyfinUserItemDataDto) SetUnplayedItemCountNil(b bool)`

 SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil

### UnsetUnplayedItemCount
`func (o *JellyfinUserItemDataDto) UnsetUnplayedItemCount()`

UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
### GetPlaybackPositionTicks

`func (o *JellyfinUserItemDataDto) GetPlaybackPositionTicks() int64`

GetPlaybackPositionTicks returns the PlaybackPositionTicks field if non-nil, zero value otherwise.

### GetPlaybackPositionTicksOk

`func (o *JellyfinUserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool)`

GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPositionTicks

`func (o *JellyfinUserItemDataDto) SetPlaybackPositionTicks(v int64)`

SetPlaybackPositionTicks sets PlaybackPositionTicks field to given value.

### HasPlaybackPositionTicks

`func (o *JellyfinUserItemDataDto) HasPlaybackPositionTicks() bool`

HasPlaybackPositionTicks returns a boolean if a field has been set.

### GetPlayCount

`func (o *JellyfinUserItemDataDto) GetPlayCount() int32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *JellyfinUserItemDataDto) GetPlayCountOk() (*int32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *JellyfinUserItemDataDto) SetPlayCount(v int32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *JellyfinUserItemDataDto) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.

### GetIsFavorite

`func (o *JellyfinUserItemDataDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *JellyfinUserItemDataDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *JellyfinUserItemDataDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *JellyfinUserItemDataDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetLikes

`func (o *JellyfinUserItemDataDto) GetLikes() bool`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *JellyfinUserItemDataDto) GetLikesOk() (*bool, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *JellyfinUserItemDataDto) SetLikes(v bool)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *JellyfinUserItemDataDto) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### SetLikesNil

`func (o *JellyfinUserItemDataDto) SetLikesNil(b bool)`

 SetLikesNil sets the value for Likes to be an explicit nil

### UnsetLikes
`func (o *JellyfinUserItemDataDto) UnsetLikes()`

UnsetLikes ensures that no value is present for Likes, not even an explicit nil
### GetLastPlayedDate

`func (o *JellyfinUserItemDataDto) GetLastPlayedDate() time.Time`

GetLastPlayedDate returns the LastPlayedDate field if non-nil, zero value otherwise.

### GetLastPlayedDateOk

`func (o *JellyfinUserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool)`

GetLastPlayedDateOk returns a tuple with the LastPlayedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayedDate

`func (o *JellyfinUserItemDataDto) SetLastPlayedDate(v time.Time)`

SetLastPlayedDate sets LastPlayedDate field to given value.

### HasLastPlayedDate

`func (o *JellyfinUserItemDataDto) HasLastPlayedDate() bool`

HasLastPlayedDate returns a boolean if a field has been set.

### SetLastPlayedDateNil

`func (o *JellyfinUserItemDataDto) SetLastPlayedDateNil(b bool)`

 SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil

### UnsetLastPlayedDate
`func (o *JellyfinUserItemDataDto) UnsetLastPlayedDate()`

UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
### GetPlayed

`func (o *JellyfinUserItemDataDto) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *JellyfinUserItemDataDto) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *JellyfinUserItemDataDto) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *JellyfinUserItemDataDto) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetKey

`func (o *JellyfinUserItemDataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JellyfinUserItemDataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JellyfinUserItemDataDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *JellyfinUserItemDataDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetItemId

`func (o *JellyfinUserItemDataDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinUserItemDataDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinUserItemDataDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinUserItemDataDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


