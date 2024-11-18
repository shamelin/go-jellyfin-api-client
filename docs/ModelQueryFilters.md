# ModelQueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genres** | Pointer to [**[]ModelModelNameGuidPair**](ModelModelNameGuidPair.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelQueryFilters

`func NewModelQueryFilters() *ModelQueryFilters`

NewModelQueryFilters instantiates a new ModelQueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQueryFiltersWithDefaults

`func NewModelQueryFiltersWithDefaults() *ModelQueryFilters`

NewModelQueryFiltersWithDefaults instantiates a new ModelQueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenres

`func (o *ModelQueryFilters) GetGenres() []ModelModelNameGuidPair`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ModelQueryFilters) GetGenresOk() (*[]ModelModelNameGuidPair, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ModelQueryFilters) SetGenres(v []ModelModelNameGuidPair)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ModelQueryFilters) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *ModelQueryFilters) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *ModelQueryFilters) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetTags

`func (o *ModelQueryFilters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelQueryFilters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelQueryFilters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelQueryFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ModelQueryFilters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ModelQueryFilters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


