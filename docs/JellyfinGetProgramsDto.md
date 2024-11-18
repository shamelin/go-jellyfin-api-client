# JellyfinGetProgramsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelIds** | Pointer to **[]string** | Gets or sets the channels to return guide information for. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets optional. Filter by user id. | [optional] 
**MinStartDate** | Pointer to **NullableTime** | Gets or sets the minimum premiere start date. | [optional] 
**HasAired** | Pointer to **NullableBool** | Gets or sets filter by programs that have completed airing, or not. | [optional] 
**IsAiring** | Pointer to **NullableBool** | Gets or sets filter by programs that are currently airing, or not. | [optional] 
**MaxStartDate** | Pointer to **NullableTime** | Gets or sets the maximum premiere start date. | [optional] 
**MinEndDate** | Pointer to **NullableTime** | Gets or sets the minimum premiere end date. | [optional] 
**MaxEndDate** | Pointer to **NullableTime** | Gets or sets the maximum premiere end date. | [optional] 
**IsMovie** | Pointer to **NullableBool** | Gets or sets filter for movies. | [optional] 
**IsSeries** | Pointer to **NullableBool** | Gets or sets filter for series. | [optional] 
**IsNews** | Pointer to **NullableBool** | Gets or sets filter for news. | [optional] 
**IsKids** | Pointer to **NullableBool** | Gets or sets filter for kids. | [optional] 
**IsSports** | Pointer to **NullableBool** | Gets or sets filter for sports. | [optional] 
**StartIndex** | Pointer to **NullableInt32** | Gets or sets the record index to start at. All items with a lower index will be dropped from the results. | [optional] 
**Limit** | Pointer to **NullableInt32** | Gets or sets the maximum number of records to return. | [optional] 
**SortBy** | Pointer to [**[]JellyfinJellyfinItemSortBy**](JellyfinJellyfinItemSortBy.md) | Gets or sets specify one or more sort orders, comma delimited. Options: Name, StartDate. | [optional] 
**SortOrder** | Pointer to [**[]JellyfinJellyfinSortOrder**](JellyfinJellyfinSortOrder.md) | Gets or sets sort order. | [optional] 
**Genres** | Pointer to **[]string** | Gets or sets the genres to return guide information for. | [optional] 
**GenreIds** | Pointer to **[]string** | Gets or sets the genre ids to return guide information for. | [optional] 
**EnableImages** | Pointer to **NullableBool** | Gets or sets include image information in output. | [optional] 
**EnableTotalRecordCount** | Pointer to **bool** | Gets or sets a value indicating whether retrieve total record count. | [optional] [default to true]
**ImageTypeLimit** | Pointer to **NullableInt32** | Gets or sets the max number of images to return, per image type. | [optional] 
**EnableImageTypes** | Pointer to [**[]JellyfinJellyfinImageType**](JellyfinJellyfinImageType.md) | Gets or sets the image types to include in the output. | [optional] 
**EnableUserData** | Pointer to **NullableBool** | Gets or sets include user data. | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets filter by series timer id. | [optional] 
**LibrarySeriesId** | Pointer to **NullableString** | Gets or sets filter by library series id. | [optional] 
**Fields** | Pointer to [**[]JellyfinJellyfinItemFields**](JellyfinJellyfinItemFields.md) | Gets or sets specify additional fields of information to return in the output. | [optional] 

## Methods

### NewJellyfinGetProgramsDto

`func NewJellyfinGetProgramsDto() *JellyfinGetProgramsDto`

NewJellyfinGetProgramsDto instantiates a new JellyfinGetProgramsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGetProgramsDtoWithDefaults

`func NewJellyfinGetProgramsDtoWithDefaults() *JellyfinGetProgramsDto`

NewJellyfinGetProgramsDtoWithDefaults instantiates a new JellyfinGetProgramsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelIds

`func (o *JellyfinGetProgramsDto) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *JellyfinGetProgramsDto) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *JellyfinGetProgramsDto) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *JellyfinGetProgramsDto) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *JellyfinGetProgramsDto) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *JellyfinGetProgramsDto) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil
### GetUserId

`func (o *JellyfinGetProgramsDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinGetProgramsDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinGetProgramsDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinGetProgramsDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *JellyfinGetProgramsDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *JellyfinGetProgramsDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMinStartDate

`func (o *JellyfinGetProgramsDto) GetMinStartDate() time.Time`

GetMinStartDate returns the MinStartDate field if non-nil, zero value otherwise.

### GetMinStartDateOk

`func (o *JellyfinGetProgramsDto) GetMinStartDateOk() (*time.Time, bool)`

GetMinStartDateOk returns a tuple with the MinStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStartDate

`func (o *JellyfinGetProgramsDto) SetMinStartDate(v time.Time)`

SetMinStartDate sets MinStartDate field to given value.

### HasMinStartDate

`func (o *JellyfinGetProgramsDto) HasMinStartDate() bool`

HasMinStartDate returns a boolean if a field has been set.

### SetMinStartDateNil

`func (o *JellyfinGetProgramsDto) SetMinStartDateNil(b bool)`

 SetMinStartDateNil sets the value for MinStartDate to be an explicit nil

### UnsetMinStartDate
`func (o *JellyfinGetProgramsDto) UnsetMinStartDate()`

UnsetMinStartDate ensures that no value is present for MinStartDate, not even an explicit nil
### GetHasAired

`func (o *JellyfinGetProgramsDto) GetHasAired() bool`

GetHasAired returns the HasAired field if non-nil, zero value otherwise.

### GetHasAiredOk

`func (o *JellyfinGetProgramsDto) GetHasAiredOk() (*bool, bool)`

GetHasAiredOk returns a tuple with the HasAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAired

`func (o *JellyfinGetProgramsDto) SetHasAired(v bool)`

SetHasAired sets HasAired field to given value.

### HasHasAired

`func (o *JellyfinGetProgramsDto) HasHasAired() bool`

HasHasAired returns a boolean if a field has been set.

### SetHasAiredNil

`func (o *JellyfinGetProgramsDto) SetHasAiredNil(b bool)`

 SetHasAiredNil sets the value for HasAired to be an explicit nil

### UnsetHasAired
`func (o *JellyfinGetProgramsDto) UnsetHasAired()`

UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil
### GetIsAiring

`func (o *JellyfinGetProgramsDto) GetIsAiring() bool`

GetIsAiring returns the IsAiring field if non-nil, zero value otherwise.

### GetIsAiringOk

`func (o *JellyfinGetProgramsDto) GetIsAiringOk() (*bool, bool)`

GetIsAiringOk returns a tuple with the IsAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiring

`func (o *JellyfinGetProgramsDto) SetIsAiring(v bool)`

SetIsAiring sets IsAiring field to given value.

### HasIsAiring

`func (o *JellyfinGetProgramsDto) HasIsAiring() bool`

HasIsAiring returns a boolean if a field has been set.

### SetIsAiringNil

`func (o *JellyfinGetProgramsDto) SetIsAiringNil(b bool)`

 SetIsAiringNil sets the value for IsAiring to be an explicit nil

### UnsetIsAiring
`func (o *JellyfinGetProgramsDto) UnsetIsAiring()`

UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
### GetMaxStartDate

`func (o *JellyfinGetProgramsDto) GetMaxStartDate() time.Time`

GetMaxStartDate returns the MaxStartDate field if non-nil, zero value otherwise.

### GetMaxStartDateOk

`func (o *JellyfinGetProgramsDto) GetMaxStartDateOk() (*time.Time, bool)`

GetMaxStartDateOk returns a tuple with the MaxStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDate

`func (o *JellyfinGetProgramsDto) SetMaxStartDate(v time.Time)`

SetMaxStartDate sets MaxStartDate field to given value.

### HasMaxStartDate

`func (o *JellyfinGetProgramsDto) HasMaxStartDate() bool`

HasMaxStartDate returns a boolean if a field has been set.

### SetMaxStartDateNil

`func (o *JellyfinGetProgramsDto) SetMaxStartDateNil(b bool)`

 SetMaxStartDateNil sets the value for MaxStartDate to be an explicit nil

### UnsetMaxStartDate
`func (o *JellyfinGetProgramsDto) UnsetMaxStartDate()`

UnsetMaxStartDate ensures that no value is present for MaxStartDate, not even an explicit nil
### GetMinEndDate

`func (o *JellyfinGetProgramsDto) GetMinEndDate() time.Time`

GetMinEndDate returns the MinEndDate field if non-nil, zero value otherwise.

### GetMinEndDateOk

`func (o *JellyfinGetProgramsDto) GetMinEndDateOk() (*time.Time, bool)`

GetMinEndDateOk returns a tuple with the MinEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEndDate

`func (o *JellyfinGetProgramsDto) SetMinEndDate(v time.Time)`

SetMinEndDate sets MinEndDate field to given value.

### HasMinEndDate

`func (o *JellyfinGetProgramsDto) HasMinEndDate() bool`

HasMinEndDate returns a boolean if a field has been set.

### SetMinEndDateNil

`func (o *JellyfinGetProgramsDto) SetMinEndDateNil(b bool)`

 SetMinEndDateNil sets the value for MinEndDate to be an explicit nil

### UnsetMinEndDate
`func (o *JellyfinGetProgramsDto) UnsetMinEndDate()`

UnsetMinEndDate ensures that no value is present for MinEndDate, not even an explicit nil
### GetMaxEndDate

`func (o *JellyfinGetProgramsDto) GetMaxEndDate() time.Time`

GetMaxEndDate returns the MaxEndDate field if non-nil, zero value otherwise.

### GetMaxEndDateOk

`func (o *JellyfinGetProgramsDto) GetMaxEndDateOk() (*time.Time, bool)`

GetMaxEndDateOk returns a tuple with the MaxEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndDate

`func (o *JellyfinGetProgramsDto) SetMaxEndDate(v time.Time)`

SetMaxEndDate sets MaxEndDate field to given value.

### HasMaxEndDate

`func (o *JellyfinGetProgramsDto) HasMaxEndDate() bool`

HasMaxEndDate returns a boolean if a field has been set.

### SetMaxEndDateNil

`func (o *JellyfinGetProgramsDto) SetMaxEndDateNil(b bool)`

 SetMaxEndDateNil sets the value for MaxEndDate to be an explicit nil

### UnsetMaxEndDate
`func (o *JellyfinGetProgramsDto) UnsetMaxEndDate()`

UnsetMaxEndDate ensures that no value is present for MaxEndDate, not even an explicit nil
### GetIsMovie

`func (o *JellyfinGetProgramsDto) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *JellyfinGetProgramsDto) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *JellyfinGetProgramsDto) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *JellyfinGetProgramsDto) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *JellyfinGetProgramsDto) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *JellyfinGetProgramsDto) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSeries

`func (o *JellyfinGetProgramsDto) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *JellyfinGetProgramsDto) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *JellyfinGetProgramsDto) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *JellyfinGetProgramsDto) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *JellyfinGetProgramsDto) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *JellyfinGetProgramsDto) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsNews

`func (o *JellyfinGetProgramsDto) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *JellyfinGetProgramsDto) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *JellyfinGetProgramsDto) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *JellyfinGetProgramsDto) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *JellyfinGetProgramsDto) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *JellyfinGetProgramsDto) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *JellyfinGetProgramsDto) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *JellyfinGetProgramsDto) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *JellyfinGetProgramsDto) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *JellyfinGetProgramsDto) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *JellyfinGetProgramsDto) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *JellyfinGetProgramsDto) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsSports

`func (o *JellyfinGetProgramsDto) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *JellyfinGetProgramsDto) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *JellyfinGetProgramsDto) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *JellyfinGetProgramsDto) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *JellyfinGetProgramsDto) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *JellyfinGetProgramsDto) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetStartIndex

`func (o *JellyfinGetProgramsDto) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *JellyfinGetProgramsDto) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *JellyfinGetProgramsDto) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *JellyfinGetProgramsDto) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### SetStartIndexNil

`func (o *JellyfinGetProgramsDto) SetStartIndexNil(b bool)`

 SetStartIndexNil sets the value for StartIndex to be an explicit nil

### UnsetStartIndex
`func (o *JellyfinGetProgramsDto) UnsetStartIndex()`

UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
### GetLimit

`func (o *JellyfinGetProgramsDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JellyfinGetProgramsDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JellyfinGetProgramsDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *JellyfinGetProgramsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *JellyfinGetProgramsDto) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *JellyfinGetProgramsDto) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetSortBy

`func (o *JellyfinGetProgramsDto) GetSortBy() []JellyfinJellyfinItemSortBy`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *JellyfinGetProgramsDto) GetSortByOk() (*[]JellyfinJellyfinItemSortBy, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *JellyfinGetProgramsDto) SetSortBy(v []JellyfinJellyfinItemSortBy)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *JellyfinGetProgramsDto) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *JellyfinGetProgramsDto) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *JellyfinGetProgramsDto) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *JellyfinGetProgramsDto) GetSortOrder() []JellyfinJellyfinSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *JellyfinGetProgramsDto) GetSortOrderOk() (*[]JellyfinJellyfinSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *JellyfinGetProgramsDto) SetSortOrder(v []JellyfinJellyfinSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *JellyfinGetProgramsDto) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *JellyfinGetProgramsDto) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *JellyfinGetProgramsDto) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetGenres

`func (o *JellyfinGetProgramsDto) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *JellyfinGetProgramsDto) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *JellyfinGetProgramsDto) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *JellyfinGetProgramsDto) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *JellyfinGetProgramsDto) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *JellyfinGetProgramsDto) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetGenreIds

`func (o *JellyfinGetProgramsDto) GetGenreIds() []string`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *JellyfinGetProgramsDto) GetGenreIdsOk() (*[]string, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *JellyfinGetProgramsDto) SetGenreIds(v []string)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *JellyfinGetProgramsDto) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### SetGenreIdsNil

`func (o *JellyfinGetProgramsDto) SetGenreIdsNil(b bool)`

 SetGenreIdsNil sets the value for GenreIds to be an explicit nil

### UnsetGenreIds
`func (o *JellyfinGetProgramsDto) UnsetGenreIds()`

UnsetGenreIds ensures that no value is present for GenreIds, not even an explicit nil
### GetEnableImages

`func (o *JellyfinGetProgramsDto) GetEnableImages() bool`

GetEnableImages returns the EnableImages field if non-nil, zero value otherwise.

### GetEnableImagesOk

`func (o *JellyfinGetProgramsDto) GetEnableImagesOk() (*bool, bool)`

GetEnableImagesOk returns a tuple with the EnableImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImages

`func (o *JellyfinGetProgramsDto) SetEnableImages(v bool)`

SetEnableImages sets EnableImages field to given value.

### HasEnableImages

`func (o *JellyfinGetProgramsDto) HasEnableImages() bool`

HasEnableImages returns a boolean if a field has been set.

### SetEnableImagesNil

`func (o *JellyfinGetProgramsDto) SetEnableImagesNil(b bool)`

 SetEnableImagesNil sets the value for EnableImages to be an explicit nil

### UnsetEnableImages
`func (o *JellyfinGetProgramsDto) UnsetEnableImages()`

UnsetEnableImages ensures that no value is present for EnableImages, not even an explicit nil
### GetEnableTotalRecordCount

`func (o *JellyfinGetProgramsDto) GetEnableTotalRecordCount() bool`

GetEnableTotalRecordCount returns the EnableTotalRecordCount field if non-nil, zero value otherwise.

### GetEnableTotalRecordCountOk

`func (o *JellyfinGetProgramsDto) GetEnableTotalRecordCountOk() (*bool, bool)`

GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotalRecordCount

`func (o *JellyfinGetProgramsDto) SetEnableTotalRecordCount(v bool)`

SetEnableTotalRecordCount sets EnableTotalRecordCount field to given value.

### HasEnableTotalRecordCount

`func (o *JellyfinGetProgramsDto) HasEnableTotalRecordCount() bool`

HasEnableTotalRecordCount returns a boolean if a field has been set.

### GetImageTypeLimit

`func (o *JellyfinGetProgramsDto) GetImageTypeLimit() int32`

GetImageTypeLimit returns the ImageTypeLimit field if non-nil, zero value otherwise.

### GetImageTypeLimitOk

`func (o *JellyfinGetProgramsDto) GetImageTypeLimitOk() (*int32, bool)`

GetImageTypeLimitOk returns a tuple with the ImageTypeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTypeLimit

`func (o *JellyfinGetProgramsDto) SetImageTypeLimit(v int32)`

SetImageTypeLimit sets ImageTypeLimit field to given value.

### HasImageTypeLimit

`func (o *JellyfinGetProgramsDto) HasImageTypeLimit() bool`

HasImageTypeLimit returns a boolean if a field has been set.

### SetImageTypeLimitNil

`func (o *JellyfinGetProgramsDto) SetImageTypeLimitNil(b bool)`

 SetImageTypeLimitNil sets the value for ImageTypeLimit to be an explicit nil

### UnsetImageTypeLimit
`func (o *JellyfinGetProgramsDto) UnsetImageTypeLimit()`

UnsetImageTypeLimit ensures that no value is present for ImageTypeLimit, not even an explicit nil
### GetEnableImageTypes

`func (o *JellyfinGetProgramsDto) GetEnableImageTypes() []JellyfinJellyfinImageType`

GetEnableImageTypes returns the EnableImageTypes field if non-nil, zero value otherwise.

### GetEnableImageTypesOk

`func (o *JellyfinGetProgramsDto) GetEnableImageTypesOk() (*[]JellyfinJellyfinImageType, bool)`

GetEnableImageTypesOk returns a tuple with the EnableImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImageTypes

`func (o *JellyfinGetProgramsDto) SetEnableImageTypes(v []JellyfinJellyfinImageType)`

SetEnableImageTypes sets EnableImageTypes field to given value.

### HasEnableImageTypes

`func (o *JellyfinGetProgramsDto) HasEnableImageTypes() bool`

HasEnableImageTypes returns a boolean if a field has been set.

### SetEnableImageTypesNil

`func (o *JellyfinGetProgramsDto) SetEnableImageTypesNil(b bool)`

 SetEnableImageTypesNil sets the value for EnableImageTypes to be an explicit nil

### UnsetEnableImageTypes
`func (o *JellyfinGetProgramsDto) UnsetEnableImageTypes()`

UnsetEnableImageTypes ensures that no value is present for EnableImageTypes, not even an explicit nil
### GetEnableUserData

`func (o *JellyfinGetProgramsDto) GetEnableUserData() bool`

GetEnableUserData returns the EnableUserData field if non-nil, zero value otherwise.

### GetEnableUserDataOk

`func (o *JellyfinGetProgramsDto) GetEnableUserDataOk() (*bool, bool)`

GetEnableUserDataOk returns a tuple with the EnableUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserData

`func (o *JellyfinGetProgramsDto) SetEnableUserData(v bool)`

SetEnableUserData sets EnableUserData field to given value.

### HasEnableUserData

`func (o *JellyfinGetProgramsDto) HasEnableUserData() bool`

HasEnableUserData returns a boolean if a field has been set.

### SetEnableUserDataNil

`func (o *JellyfinGetProgramsDto) SetEnableUserDataNil(b bool)`

 SetEnableUserDataNil sets the value for EnableUserData to be an explicit nil

### UnsetEnableUserData
`func (o *JellyfinGetProgramsDto) UnsetEnableUserData()`

UnsetEnableUserData ensures that no value is present for EnableUserData, not even an explicit nil
### GetSeriesTimerId

`func (o *JellyfinGetProgramsDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *JellyfinGetProgramsDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *JellyfinGetProgramsDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *JellyfinGetProgramsDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *JellyfinGetProgramsDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *JellyfinGetProgramsDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetLibrarySeriesId

`func (o *JellyfinGetProgramsDto) GetLibrarySeriesId() string`

GetLibrarySeriesId returns the LibrarySeriesId field if non-nil, zero value otherwise.

### GetLibrarySeriesIdOk

`func (o *JellyfinGetProgramsDto) GetLibrarySeriesIdOk() (*string, bool)`

GetLibrarySeriesIdOk returns a tuple with the LibrarySeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrarySeriesId

`func (o *JellyfinGetProgramsDto) SetLibrarySeriesId(v string)`

SetLibrarySeriesId sets LibrarySeriesId field to given value.

### HasLibrarySeriesId

`func (o *JellyfinGetProgramsDto) HasLibrarySeriesId() bool`

HasLibrarySeriesId returns a boolean if a field has been set.

### SetLibrarySeriesIdNil

`func (o *JellyfinGetProgramsDto) SetLibrarySeriesIdNil(b bool)`

 SetLibrarySeriesIdNil sets the value for LibrarySeriesId to be an explicit nil

### UnsetLibrarySeriesId
`func (o *JellyfinGetProgramsDto) UnsetLibrarySeriesId()`

UnsetLibrarySeriesId ensures that no value is present for LibrarySeriesId, not even an explicit nil
### GetFields

`func (o *JellyfinGetProgramsDto) GetFields() []JellyfinJellyfinItemFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *JellyfinGetProgramsDto) GetFieldsOk() (*[]JellyfinJellyfinItemFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *JellyfinGetProgramsDto) SetFields(v []JellyfinJellyfinItemFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *JellyfinGetProgramsDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *JellyfinGetProgramsDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *JellyfinGetProgramsDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


