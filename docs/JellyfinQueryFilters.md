# JellyfinQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genres** | Pointer to [**[]JellyfinJellyfinNameGuidPair**](JellyfinJellyfinNameGuidPair.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJellyfinQueryFilters

`func NewJellyfinQueryFilters() *JellyfinQueryFilters`

NewJellyfinQueryFilters instantiates a new JellyfinQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinQueryFiltersWithDefaults

`func NewJellyfinQueryFiltersWithDefaults() *JellyfinQueryFilters`

NewJellyfinQueryFiltersWithDefaults instantiates a new JellyfinQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenres

`func (o *JellyfinQueryFilters) GetGenres() []JellyfinJellyfinNameGuidPair`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *JellyfinQueryFilters) GetGenresOk() (*[]JellyfinJellyfinNameGuidPair, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *JellyfinQueryFilters) SetGenres(v []JellyfinJellyfinNameGuidPair)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *JellyfinQueryFilters) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *JellyfinQueryFilters) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *JellyfinQueryFilters) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetTags

`func (o *JellyfinQueryFilters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *JellyfinQueryFilters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *JellyfinQueryFilters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *JellyfinQueryFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *JellyfinQueryFilters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *JellyfinQueryFilters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


