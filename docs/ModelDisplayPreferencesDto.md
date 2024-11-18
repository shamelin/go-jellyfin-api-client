# ModelDisplayPreferencesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**ViewType** | Pointer to **NullableString** | Gets or sets the type of the view. | [optional] 
**SortBy** | Pointer to **NullableString** | Gets or sets the sort by. | [optional] 
**IndexBy** | Pointer to **NullableString** | Gets or sets the index by. | [optional] 
**RememberIndexing** | Pointer to **bool** | Gets or sets a value indicating whether [remember indexing]. | [optional] 
**PrimaryImageHeight** | Pointer to **int32** | Gets or sets the height of the primary image. | [optional] 
**PrimaryImageWidth** | Pointer to **int32** | Gets or sets the width of the primary image. | [optional] 
**CustomPrefs** | Pointer to **map[string]string** | Gets or sets the custom prefs. | [optional] 
**ScrollDirection** | Pointer to [**ModelModelScrollDirection**](ModelScrollDirection.md) | An enum representing the axis that should be scrolled. | [optional] 
**ShowBackdrop** | Pointer to **bool** | Gets or sets a value indicating whether to show backdrops on this item. | [optional] 
**RememberSorting** | Pointer to **bool** | Gets or sets a value indicating whether [remember sorting]. | [optional] 
**SortOrder** | Pointer to [**ModelModelSortOrder**](ModelSortOrder.md) | An enum representing the sorting order. | [optional] 
**ShowSidebar** | Pointer to **bool** | Gets or sets a value indicating whether [show sidebar]. | [optional] 
**Client** | Pointer to **NullableString** | Gets or sets the client. | [optional] 

## Methods

### NewModelDisplayPreferencesDto

`func NewModelDisplayPreferencesDto() *ModelDisplayPreferencesDto`

NewModelDisplayPreferencesDto instantiates a new ModelDisplayPreferencesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDisplayPreferencesDtoWithDefaults

`func NewModelDisplayPreferencesDtoWithDefaults() *ModelDisplayPreferencesDto`

NewModelDisplayPreferencesDtoWithDefaults instantiates a new ModelDisplayPreferencesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelDisplayPreferencesDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDisplayPreferencesDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDisplayPreferencesDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDisplayPreferencesDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelDisplayPreferencesDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelDisplayPreferencesDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetViewType

`func (o *ModelDisplayPreferencesDto) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *ModelDisplayPreferencesDto) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *ModelDisplayPreferencesDto) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *ModelDisplayPreferencesDto) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### SetViewTypeNil

`func (o *ModelDisplayPreferencesDto) SetViewTypeNil(b bool)`

 SetViewTypeNil sets the value for ViewType to be an explicit nil

### UnsetViewType
`func (o *ModelDisplayPreferencesDto) UnsetViewType()`

UnsetViewType ensures that no value is present for ViewType, not even an explicit nil
### GetSortBy

`func (o *ModelDisplayPreferencesDto) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ModelDisplayPreferencesDto) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ModelDisplayPreferencesDto) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ModelDisplayPreferencesDto) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *ModelDisplayPreferencesDto) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *ModelDisplayPreferencesDto) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetIndexBy

`func (o *ModelDisplayPreferencesDto) GetIndexBy() string`

GetIndexBy returns the IndexBy field if non-nil, zero value otherwise.

### GetIndexByOk

`func (o *ModelDisplayPreferencesDto) GetIndexByOk() (*string, bool)`

GetIndexByOk returns a tuple with the IndexBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBy

`func (o *ModelDisplayPreferencesDto) SetIndexBy(v string)`

SetIndexBy sets IndexBy field to given value.

### HasIndexBy

`func (o *ModelDisplayPreferencesDto) HasIndexBy() bool`

HasIndexBy returns a boolean if a field has been set.

### SetIndexByNil

`func (o *ModelDisplayPreferencesDto) SetIndexByNil(b bool)`

 SetIndexByNil sets the value for IndexBy to be an explicit nil

### UnsetIndexBy
`func (o *ModelDisplayPreferencesDto) UnsetIndexBy()`

UnsetIndexBy ensures that no value is present for IndexBy, not even an explicit nil
### GetRememberIndexing

`func (o *ModelDisplayPreferencesDto) GetRememberIndexing() bool`

GetRememberIndexing returns the RememberIndexing field if non-nil, zero value otherwise.

### GetRememberIndexingOk

`func (o *ModelDisplayPreferencesDto) GetRememberIndexingOk() (*bool, bool)`

GetRememberIndexingOk returns a tuple with the RememberIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberIndexing

`func (o *ModelDisplayPreferencesDto) SetRememberIndexing(v bool)`

SetRememberIndexing sets RememberIndexing field to given value.

### HasRememberIndexing

`func (o *ModelDisplayPreferencesDto) HasRememberIndexing() bool`

HasRememberIndexing returns a boolean if a field has been set.

### GetPrimaryImageHeight

`func (o *ModelDisplayPreferencesDto) GetPrimaryImageHeight() int32`

GetPrimaryImageHeight returns the PrimaryImageHeight field if non-nil, zero value otherwise.

### GetPrimaryImageHeightOk

`func (o *ModelDisplayPreferencesDto) GetPrimaryImageHeightOk() (*int32, bool)`

GetPrimaryImageHeightOk returns a tuple with the PrimaryImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageHeight

`func (o *ModelDisplayPreferencesDto) SetPrimaryImageHeight(v int32)`

SetPrimaryImageHeight sets PrimaryImageHeight field to given value.

### HasPrimaryImageHeight

`func (o *ModelDisplayPreferencesDto) HasPrimaryImageHeight() bool`

HasPrimaryImageHeight returns a boolean if a field has been set.

### GetPrimaryImageWidth

`func (o *ModelDisplayPreferencesDto) GetPrimaryImageWidth() int32`

GetPrimaryImageWidth returns the PrimaryImageWidth field if non-nil, zero value otherwise.

### GetPrimaryImageWidthOk

`func (o *ModelDisplayPreferencesDto) GetPrimaryImageWidthOk() (*int32, bool)`

GetPrimaryImageWidthOk returns a tuple with the PrimaryImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageWidth

`func (o *ModelDisplayPreferencesDto) SetPrimaryImageWidth(v int32)`

SetPrimaryImageWidth sets PrimaryImageWidth field to given value.

### HasPrimaryImageWidth

`func (o *ModelDisplayPreferencesDto) HasPrimaryImageWidth() bool`

HasPrimaryImageWidth returns a boolean if a field has been set.

### GetCustomPrefs

`func (o *ModelDisplayPreferencesDto) GetCustomPrefs() map[string]string`

GetCustomPrefs returns the CustomPrefs field if non-nil, zero value otherwise.

### GetCustomPrefsOk

`func (o *ModelDisplayPreferencesDto) GetCustomPrefsOk() (*map[string]string, bool)`

GetCustomPrefsOk returns a tuple with the CustomPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrefs

`func (o *ModelDisplayPreferencesDto) SetCustomPrefs(v map[string]string)`

SetCustomPrefs sets CustomPrefs field to given value.

### HasCustomPrefs

`func (o *ModelDisplayPreferencesDto) HasCustomPrefs() bool`

HasCustomPrefs returns a boolean if a field has been set.

### GetScrollDirection

`func (o *ModelDisplayPreferencesDto) GetScrollDirection() ModelModelScrollDirection`

GetScrollDirection returns the ScrollDirection field if non-nil, zero value otherwise.

### GetScrollDirectionOk

`func (o *ModelDisplayPreferencesDto) GetScrollDirectionOk() (*ModelModelScrollDirection, bool)`

GetScrollDirectionOk returns a tuple with the ScrollDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollDirection

`func (o *ModelDisplayPreferencesDto) SetScrollDirection(v ModelModelScrollDirection)`

SetScrollDirection sets ScrollDirection field to given value.

### HasScrollDirection

`func (o *ModelDisplayPreferencesDto) HasScrollDirection() bool`

HasScrollDirection returns a boolean if a field has been set.

### GetShowBackdrop

`func (o *ModelDisplayPreferencesDto) GetShowBackdrop() bool`

GetShowBackdrop returns the ShowBackdrop field if non-nil, zero value otherwise.

### GetShowBackdropOk

`func (o *ModelDisplayPreferencesDto) GetShowBackdropOk() (*bool, bool)`

GetShowBackdropOk returns a tuple with the ShowBackdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBackdrop

`func (o *ModelDisplayPreferencesDto) SetShowBackdrop(v bool)`

SetShowBackdrop sets ShowBackdrop field to given value.

### HasShowBackdrop

`func (o *ModelDisplayPreferencesDto) HasShowBackdrop() bool`

HasShowBackdrop returns a boolean if a field has been set.

### GetRememberSorting

`func (o *ModelDisplayPreferencesDto) GetRememberSorting() bool`

GetRememberSorting returns the RememberSorting field if non-nil, zero value otherwise.

### GetRememberSortingOk

`func (o *ModelDisplayPreferencesDto) GetRememberSortingOk() (*bool, bool)`

GetRememberSortingOk returns a tuple with the RememberSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSorting

`func (o *ModelDisplayPreferencesDto) SetRememberSorting(v bool)`

SetRememberSorting sets RememberSorting field to given value.

### HasRememberSorting

`func (o *ModelDisplayPreferencesDto) HasRememberSorting() bool`

HasRememberSorting returns a boolean if a field has been set.

### GetSortOrder

`func (o *ModelDisplayPreferencesDto) GetSortOrder() ModelModelSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ModelDisplayPreferencesDto) GetSortOrderOk() (*ModelModelSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ModelDisplayPreferencesDto) SetSortOrder(v ModelModelSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ModelDisplayPreferencesDto) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetShowSidebar

`func (o *ModelDisplayPreferencesDto) GetShowSidebar() bool`

GetShowSidebar returns the ShowSidebar field if non-nil, zero value otherwise.

### GetShowSidebarOk

`func (o *ModelDisplayPreferencesDto) GetShowSidebarOk() (*bool, bool)`

GetShowSidebarOk returns a tuple with the ShowSidebar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSidebar

`func (o *ModelDisplayPreferencesDto) SetShowSidebar(v bool)`

SetShowSidebar sets ShowSidebar field to given value.

### HasShowSidebar

`func (o *ModelDisplayPreferencesDto) HasShowSidebar() bool`

HasShowSidebar returns a boolean if a field has been set.

### GetClient

`func (o *ModelDisplayPreferencesDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ModelDisplayPreferencesDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ModelDisplayPreferencesDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ModelDisplayPreferencesDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *ModelDisplayPreferencesDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *ModelDisplayPreferencesDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


