# ModelGetProgramsDto

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
**SortBy** | Pointer to [**[]ModelModelItemSortBy**](ModelModelItemSortBy.md) | Gets or sets specify one or more sort orders, comma delimited. Options: Name, StartDate. | [optional] 
**SortOrder** | Pointer to [**[]ModelModelSortOrder**](ModelModelSortOrder.md) | Gets or sets sort order. | [optional] 
**Genres** | Pointer to **[]string** | Gets or sets the genres to return guide information for. | [optional] 
**GenreIds** | Pointer to **[]string** | Gets or sets the genre ids to return guide information for. | [optional] 
**EnableImages** | Pointer to **NullableBool** | Gets or sets include image information in output. | [optional] 
**EnableTotalRecordCount** | Pointer to **bool** | Gets or sets a value indicating whether retrieve total record count. | [optional] [default to true]
**ImageTypeLimit** | Pointer to **NullableInt32** | Gets or sets the max number of images to return, per image type. | [optional] 
**EnableImageTypes** | Pointer to [**[]ModelModelImageType**](ModelModelImageType.md) | Gets or sets the image types to include in the output. | [optional] 
**EnableUserData** | Pointer to **NullableBool** | Gets or sets include user data. | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets filter by series timer id. | [optional] 
**LibrarySeriesId** | Pointer to **NullableString** | Gets or sets filter by library series id. | [optional] 
**Fields** | Pointer to [**[]ModelModelItemFields**](ModelModelItemFields.md) | Gets or sets specify additional fields of information to return in the output. | [optional] 

## Methods

### NewModelGetProgramsDto

`func NewModelGetProgramsDto() *ModelGetProgramsDto`

NewModelGetProgramsDto instantiates a new ModelGetProgramsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGetProgramsDtoWithDefaults

`func NewModelGetProgramsDtoWithDefaults() *ModelGetProgramsDto`

NewModelGetProgramsDtoWithDefaults instantiates a new ModelGetProgramsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelIds

`func (o *ModelGetProgramsDto) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *ModelGetProgramsDto) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *ModelGetProgramsDto) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *ModelGetProgramsDto) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### SetChannelIdsNil

`func (o *ModelGetProgramsDto) SetChannelIdsNil(b bool)`

 SetChannelIdsNil sets the value for ChannelIds to be an explicit nil

### UnsetChannelIds
`func (o *ModelGetProgramsDto) UnsetChannelIds()`

UnsetChannelIds ensures that no value is present for ChannelIds, not even an explicit nil
### GetUserId

`func (o *ModelGetProgramsDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelGetProgramsDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelGetProgramsDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelGetProgramsDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ModelGetProgramsDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ModelGetProgramsDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMinStartDate

`func (o *ModelGetProgramsDto) GetMinStartDate() time.Time`

GetMinStartDate returns the MinStartDate field if non-nil, zero value otherwise.

### GetMinStartDateOk

`func (o *ModelGetProgramsDto) GetMinStartDateOk() (*time.Time, bool)`

GetMinStartDateOk returns a tuple with the MinStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStartDate

`func (o *ModelGetProgramsDto) SetMinStartDate(v time.Time)`

SetMinStartDate sets MinStartDate field to given value.

### HasMinStartDate

`func (o *ModelGetProgramsDto) HasMinStartDate() bool`

HasMinStartDate returns a boolean if a field has been set.

### SetMinStartDateNil

`func (o *ModelGetProgramsDto) SetMinStartDateNil(b bool)`

 SetMinStartDateNil sets the value for MinStartDate to be an explicit nil

### UnsetMinStartDate
`func (o *ModelGetProgramsDto) UnsetMinStartDate()`

UnsetMinStartDate ensures that no value is present for MinStartDate, not even an explicit nil
### GetHasAired

`func (o *ModelGetProgramsDto) GetHasAired() bool`

GetHasAired returns the HasAired field if non-nil, zero value otherwise.

### GetHasAiredOk

`func (o *ModelGetProgramsDto) GetHasAiredOk() (*bool, bool)`

GetHasAiredOk returns a tuple with the HasAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAired

`func (o *ModelGetProgramsDto) SetHasAired(v bool)`

SetHasAired sets HasAired field to given value.

### HasHasAired

`func (o *ModelGetProgramsDto) HasHasAired() bool`

HasHasAired returns a boolean if a field has been set.

### SetHasAiredNil

`func (o *ModelGetProgramsDto) SetHasAiredNil(b bool)`

 SetHasAiredNil sets the value for HasAired to be an explicit nil

### UnsetHasAired
`func (o *ModelGetProgramsDto) UnsetHasAired()`

UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil
### GetIsAiring

`func (o *ModelGetProgramsDto) GetIsAiring() bool`

GetIsAiring returns the IsAiring field if non-nil, zero value otherwise.

### GetIsAiringOk

`func (o *ModelGetProgramsDto) GetIsAiringOk() (*bool, bool)`

GetIsAiringOk returns a tuple with the IsAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiring

`func (o *ModelGetProgramsDto) SetIsAiring(v bool)`

SetIsAiring sets IsAiring field to given value.

### HasIsAiring

`func (o *ModelGetProgramsDto) HasIsAiring() bool`

HasIsAiring returns a boolean if a field has been set.

### SetIsAiringNil

`func (o *ModelGetProgramsDto) SetIsAiringNil(b bool)`

 SetIsAiringNil sets the value for IsAiring to be an explicit nil

### UnsetIsAiring
`func (o *ModelGetProgramsDto) UnsetIsAiring()`

UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
### GetMaxStartDate

`func (o *ModelGetProgramsDto) GetMaxStartDate() time.Time`

GetMaxStartDate returns the MaxStartDate field if non-nil, zero value otherwise.

### GetMaxStartDateOk

`func (o *ModelGetProgramsDto) GetMaxStartDateOk() (*time.Time, bool)`

GetMaxStartDateOk returns a tuple with the MaxStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDate

`func (o *ModelGetProgramsDto) SetMaxStartDate(v time.Time)`

SetMaxStartDate sets MaxStartDate field to given value.

### HasMaxStartDate

`func (o *ModelGetProgramsDto) HasMaxStartDate() bool`

HasMaxStartDate returns a boolean if a field has been set.

### SetMaxStartDateNil

`func (o *ModelGetProgramsDto) SetMaxStartDateNil(b bool)`

 SetMaxStartDateNil sets the value for MaxStartDate to be an explicit nil

### UnsetMaxStartDate
`func (o *ModelGetProgramsDto) UnsetMaxStartDate()`

UnsetMaxStartDate ensures that no value is present for MaxStartDate, not even an explicit nil
### GetMinEndDate

`func (o *ModelGetProgramsDto) GetMinEndDate() time.Time`

GetMinEndDate returns the MinEndDate field if non-nil, zero value otherwise.

### GetMinEndDateOk

`func (o *ModelGetProgramsDto) GetMinEndDateOk() (*time.Time, bool)`

GetMinEndDateOk returns a tuple with the MinEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEndDate

`func (o *ModelGetProgramsDto) SetMinEndDate(v time.Time)`

SetMinEndDate sets MinEndDate field to given value.

### HasMinEndDate

`func (o *ModelGetProgramsDto) HasMinEndDate() bool`

HasMinEndDate returns a boolean if a field has been set.

### SetMinEndDateNil

`func (o *ModelGetProgramsDto) SetMinEndDateNil(b bool)`

 SetMinEndDateNil sets the value for MinEndDate to be an explicit nil

### UnsetMinEndDate
`func (o *ModelGetProgramsDto) UnsetMinEndDate()`

UnsetMinEndDate ensures that no value is present for MinEndDate, not even an explicit nil
### GetMaxEndDate

`func (o *ModelGetProgramsDto) GetMaxEndDate() time.Time`

GetMaxEndDate returns the MaxEndDate field if non-nil, zero value otherwise.

### GetMaxEndDateOk

`func (o *ModelGetProgramsDto) GetMaxEndDateOk() (*time.Time, bool)`

GetMaxEndDateOk returns a tuple with the MaxEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndDate

`func (o *ModelGetProgramsDto) SetMaxEndDate(v time.Time)`

SetMaxEndDate sets MaxEndDate field to given value.

### HasMaxEndDate

`func (o *ModelGetProgramsDto) HasMaxEndDate() bool`

HasMaxEndDate returns a boolean if a field has been set.

### SetMaxEndDateNil

`func (o *ModelGetProgramsDto) SetMaxEndDateNil(b bool)`

 SetMaxEndDateNil sets the value for MaxEndDate to be an explicit nil

### UnsetMaxEndDate
`func (o *ModelGetProgramsDto) UnsetMaxEndDate()`

UnsetMaxEndDate ensures that no value is present for MaxEndDate, not even an explicit nil
### GetIsMovie

`func (o *ModelGetProgramsDto) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *ModelGetProgramsDto) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *ModelGetProgramsDto) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *ModelGetProgramsDto) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *ModelGetProgramsDto) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *ModelGetProgramsDto) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSeries

`func (o *ModelGetProgramsDto) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *ModelGetProgramsDto) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *ModelGetProgramsDto) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *ModelGetProgramsDto) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *ModelGetProgramsDto) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *ModelGetProgramsDto) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsNews

`func (o *ModelGetProgramsDto) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *ModelGetProgramsDto) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *ModelGetProgramsDto) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *ModelGetProgramsDto) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *ModelGetProgramsDto) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *ModelGetProgramsDto) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *ModelGetProgramsDto) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *ModelGetProgramsDto) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *ModelGetProgramsDto) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *ModelGetProgramsDto) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *ModelGetProgramsDto) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *ModelGetProgramsDto) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsSports

`func (o *ModelGetProgramsDto) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *ModelGetProgramsDto) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *ModelGetProgramsDto) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *ModelGetProgramsDto) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *ModelGetProgramsDto) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *ModelGetProgramsDto) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetStartIndex

`func (o *ModelGetProgramsDto) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ModelGetProgramsDto) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ModelGetProgramsDto) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *ModelGetProgramsDto) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### SetStartIndexNil

`func (o *ModelGetProgramsDto) SetStartIndexNil(b bool)`

 SetStartIndexNil sets the value for StartIndex to be an explicit nil

### UnsetStartIndex
`func (o *ModelGetProgramsDto) UnsetStartIndex()`

UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
### GetLimit

`func (o *ModelGetProgramsDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ModelGetProgramsDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ModelGetProgramsDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ModelGetProgramsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ModelGetProgramsDto) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ModelGetProgramsDto) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetSortBy

`func (o *ModelGetProgramsDto) GetSortBy() []ModelModelItemSortBy`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ModelGetProgramsDto) GetSortByOk() (*[]ModelModelItemSortBy, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ModelGetProgramsDto) SetSortBy(v []ModelModelItemSortBy)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ModelGetProgramsDto) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *ModelGetProgramsDto) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *ModelGetProgramsDto) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *ModelGetProgramsDto) GetSortOrder() []ModelModelSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ModelGetProgramsDto) GetSortOrderOk() (*[]ModelModelSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ModelGetProgramsDto) SetSortOrder(v []ModelModelSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ModelGetProgramsDto) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ModelGetProgramsDto) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ModelGetProgramsDto) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetGenres

`func (o *ModelGetProgramsDto) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ModelGetProgramsDto) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ModelGetProgramsDto) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ModelGetProgramsDto) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *ModelGetProgramsDto) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *ModelGetProgramsDto) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetGenreIds

`func (o *ModelGetProgramsDto) GetGenreIds() []string`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *ModelGetProgramsDto) GetGenreIdsOk() (*[]string, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *ModelGetProgramsDto) SetGenreIds(v []string)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *ModelGetProgramsDto) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### SetGenreIdsNil

`func (o *ModelGetProgramsDto) SetGenreIdsNil(b bool)`

 SetGenreIdsNil sets the value for GenreIds to be an explicit nil

### UnsetGenreIds
`func (o *ModelGetProgramsDto) UnsetGenreIds()`

UnsetGenreIds ensures that no value is present for GenreIds, not even an explicit nil
### GetEnableImages

`func (o *ModelGetProgramsDto) GetEnableImages() bool`

GetEnableImages returns the EnableImages field if non-nil, zero value otherwise.

### GetEnableImagesOk

`func (o *ModelGetProgramsDto) GetEnableImagesOk() (*bool, bool)`

GetEnableImagesOk returns a tuple with the EnableImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImages

`func (o *ModelGetProgramsDto) SetEnableImages(v bool)`

SetEnableImages sets EnableImages field to given value.

### HasEnableImages

`func (o *ModelGetProgramsDto) HasEnableImages() bool`

HasEnableImages returns a boolean if a field has been set.

### SetEnableImagesNil

`func (o *ModelGetProgramsDto) SetEnableImagesNil(b bool)`

 SetEnableImagesNil sets the value for EnableImages to be an explicit nil

### UnsetEnableImages
`func (o *ModelGetProgramsDto) UnsetEnableImages()`

UnsetEnableImages ensures that no value is present for EnableImages, not even an explicit nil
### GetEnableTotalRecordCount

`func (o *ModelGetProgramsDto) GetEnableTotalRecordCount() bool`

GetEnableTotalRecordCount returns the EnableTotalRecordCount field if non-nil, zero value otherwise.

### GetEnableTotalRecordCountOk

`func (o *ModelGetProgramsDto) GetEnableTotalRecordCountOk() (*bool, bool)`

GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotalRecordCount

`func (o *ModelGetProgramsDto) SetEnableTotalRecordCount(v bool)`

SetEnableTotalRecordCount sets EnableTotalRecordCount field to given value.

### HasEnableTotalRecordCount

`func (o *ModelGetProgramsDto) HasEnableTotalRecordCount() bool`

HasEnableTotalRecordCount returns a boolean if a field has been set.

### GetImageTypeLimit

`func (o *ModelGetProgramsDto) GetImageTypeLimit() int32`

GetImageTypeLimit returns the ImageTypeLimit field if non-nil, zero value otherwise.

### GetImageTypeLimitOk

`func (o *ModelGetProgramsDto) GetImageTypeLimitOk() (*int32, bool)`

GetImageTypeLimitOk returns a tuple with the ImageTypeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTypeLimit

`func (o *ModelGetProgramsDto) SetImageTypeLimit(v int32)`

SetImageTypeLimit sets ImageTypeLimit field to given value.

### HasImageTypeLimit

`func (o *ModelGetProgramsDto) HasImageTypeLimit() bool`

HasImageTypeLimit returns a boolean if a field has been set.

### SetImageTypeLimitNil

`func (o *ModelGetProgramsDto) SetImageTypeLimitNil(b bool)`

 SetImageTypeLimitNil sets the value for ImageTypeLimit to be an explicit nil

### UnsetImageTypeLimit
`func (o *ModelGetProgramsDto) UnsetImageTypeLimit()`

UnsetImageTypeLimit ensures that no value is present for ImageTypeLimit, not even an explicit nil
### GetEnableImageTypes

`func (o *ModelGetProgramsDto) GetEnableImageTypes() []ModelModelImageType`

GetEnableImageTypes returns the EnableImageTypes field if non-nil, zero value otherwise.

### GetEnableImageTypesOk

`func (o *ModelGetProgramsDto) GetEnableImageTypesOk() (*[]ModelModelImageType, bool)`

GetEnableImageTypesOk returns a tuple with the EnableImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImageTypes

`func (o *ModelGetProgramsDto) SetEnableImageTypes(v []ModelModelImageType)`

SetEnableImageTypes sets EnableImageTypes field to given value.

### HasEnableImageTypes

`func (o *ModelGetProgramsDto) HasEnableImageTypes() bool`

HasEnableImageTypes returns a boolean if a field has been set.

### SetEnableImageTypesNil

`func (o *ModelGetProgramsDto) SetEnableImageTypesNil(b bool)`

 SetEnableImageTypesNil sets the value for EnableImageTypes to be an explicit nil

### UnsetEnableImageTypes
`func (o *ModelGetProgramsDto) UnsetEnableImageTypes()`

UnsetEnableImageTypes ensures that no value is present for EnableImageTypes, not even an explicit nil
### GetEnableUserData

`func (o *ModelGetProgramsDto) GetEnableUserData() bool`

GetEnableUserData returns the EnableUserData field if non-nil, zero value otherwise.

### GetEnableUserDataOk

`func (o *ModelGetProgramsDto) GetEnableUserDataOk() (*bool, bool)`

GetEnableUserDataOk returns a tuple with the EnableUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserData

`func (o *ModelGetProgramsDto) SetEnableUserData(v bool)`

SetEnableUserData sets EnableUserData field to given value.

### HasEnableUserData

`func (o *ModelGetProgramsDto) HasEnableUserData() bool`

HasEnableUserData returns a boolean if a field has been set.

### SetEnableUserDataNil

`func (o *ModelGetProgramsDto) SetEnableUserDataNil(b bool)`

 SetEnableUserDataNil sets the value for EnableUserData to be an explicit nil

### UnsetEnableUserData
`func (o *ModelGetProgramsDto) UnsetEnableUserData()`

UnsetEnableUserData ensures that no value is present for EnableUserData, not even an explicit nil
### GetSeriesTimerId

`func (o *ModelGetProgramsDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *ModelGetProgramsDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *ModelGetProgramsDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *ModelGetProgramsDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *ModelGetProgramsDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *ModelGetProgramsDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetLibrarySeriesId

`func (o *ModelGetProgramsDto) GetLibrarySeriesId() string`

GetLibrarySeriesId returns the LibrarySeriesId field if non-nil, zero value otherwise.

### GetLibrarySeriesIdOk

`func (o *ModelGetProgramsDto) GetLibrarySeriesIdOk() (*string, bool)`

GetLibrarySeriesIdOk returns a tuple with the LibrarySeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrarySeriesId

`func (o *ModelGetProgramsDto) SetLibrarySeriesId(v string)`

SetLibrarySeriesId sets LibrarySeriesId field to given value.

### HasLibrarySeriesId

`func (o *ModelGetProgramsDto) HasLibrarySeriesId() bool`

HasLibrarySeriesId returns a boolean if a field has been set.

### SetLibrarySeriesIdNil

`func (o *ModelGetProgramsDto) SetLibrarySeriesIdNil(b bool)`

 SetLibrarySeriesIdNil sets the value for LibrarySeriesId to be an explicit nil

### UnsetLibrarySeriesId
`func (o *ModelGetProgramsDto) UnsetLibrarySeriesId()`

UnsetLibrarySeriesId ensures that no value is present for LibrarySeriesId, not even an explicit nil
### GetFields

`func (o *ModelGetProgramsDto) GetFields() []ModelModelItemFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ModelGetProgramsDto) GetFieldsOk() (*[]ModelModelItemFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ModelGetProgramsDto) SetFields(v []ModelModelItemFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ModelGetProgramsDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ModelGetProgramsDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ModelGetProgramsDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


