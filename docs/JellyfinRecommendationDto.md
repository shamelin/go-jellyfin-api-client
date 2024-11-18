# JellyfinRecommendationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]JellyfinJellyfinBaseItemDto**](JellyfinJellyfinBaseItemDto.md) |  | [optional] 
**RecommendationType** | Pointer to [**JellyfinJellyfinRecommendationType**](JellyfinRecommendationType.md) |  | [optional] 
**BaselineItemName** | Pointer to **NullableString** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 

## Methods

### NewJellyfinRecommendationDto

`func NewJellyfinRecommendationDto() *JellyfinRecommendationDto`

NewJellyfinRecommendationDto instantiates a new JellyfinRecommendationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinRecommendationDtoWithDefaults

`func NewJellyfinRecommendationDtoWithDefaults() *JellyfinRecommendationDto`

NewJellyfinRecommendationDtoWithDefaults instantiates a new JellyfinRecommendationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *JellyfinRecommendationDto) GetItems() []JellyfinJellyfinBaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JellyfinRecommendationDto) GetItemsOk() (*[]JellyfinJellyfinBaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JellyfinRecommendationDto) SetItems(v []JellyfinJellyfinBaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *JellyfinRecommendationDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *JellyfinRecommendationDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *JellyfinRecommendationDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetRecommendationType

`func (o *JellyfinRecommendationDto) GetRecommendationType() JellyfinJellyfinRecommendationType`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *JellyfinRecommendationDto) GetRecommendationTypeOk() (*JellyfinJellyfinRecommendationType, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *JellyfinRecommendationDto) SetRecommendationType(v JellyfinJellyfinRecommendationType)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *JellyfinRecommendationDto) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetBaselineItemName

`func (o *JellyfinRecommendationDto) GetBaselineItemName() string`

GetBaselineItemName returns the BaselineItemName field if non-nil, zero value otherwise.

### GetBaselineItemNameOk

`func (o *JellyfinRecommendationDto) GetBaselineItemNameOk() (*string, bool)`

GetBaselineItemNameOk returns a tuple with the BaselineItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineItemName

`func (o *JellyfinRecommendationDto) SetBaselineItemName(v string)`

SetBaselineItemName sets BaselineItemName field to given value.

### HasBaselineItemName

`func (o *JellyfinRecommendationDto) HasBaselineItemName() bool`

HasBaselineItemName returns a boolean if a field has been set.

### SetBaselineItemNameNil

`func (o *JellyfinRecommendationDto) SetBaselineItemNameNil(b bool)`

 SetBaselineItemNameNil sets the value for BaselineItemName to be an explicit nil

### UnsetBaselineItemName
`func (o *JellyfinRecommendationDto) UnsetBaselineItemName()`

UnsetBaselineItemName ensures that no value is present for BaselineItemName, not even an explicit nil
### GetCategoryId

`func (o *JellyfinRecommendationDto) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *JellyfinRecommendationDto) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *JellyfinRecommendationDto) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *JellyfinRecommendationDto) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


