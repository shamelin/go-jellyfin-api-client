# JellyfinVirtualFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Locations** | Pointer to **[]string** | Gets or sets the locations. | [optional] 
**CollectionType** | Pointer to [**NullableJellyfinJellyfinCollectionTypeOptions**](JellyfinCollectionTypeOptions.md) | Gets or sets the type of the collection. | [optional] 
**LibraryOptions** | Pointer to [**NullableJellyfinJellyfinLibraryOptions**](JellyfinLibraryOptions.md) |  | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 
**PrimaryImageItemId** | Pointer to **NullableString** | Gets or sets the primary image item identifier. | [optional] 
**RefreshProgress** | Pointer to **NullableFloat64** |  | [optional] 
**RefreshStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJellyfinVirtualFolderInfo

`func NewJellyfinVirtualFolderInfo() *JellyfinVirtualFolderInfo`

NewJellyfinVirtualFolderInfo instantiates a new JellyfinVirtualFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinVirtualFolderInfoWithDefaults

`func NewJellyfinVirtualFolderInfoWithDefaults() *JellyfinVirtualFolderInfo`

NewJellyfinVirtualFolderInfoWithDefaults instantiates a new JellyfinVirtualFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinVirtualFolderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinVirtualFolderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinVirtualFolderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinVirtualFolderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinVirtualFolderInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinVirtualFolderInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocations

`func (o *JellyfinVirtualFolderInfo) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *JellyfinVirtualFolderInfo) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *JellyfinVirtualFolderInfo) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *JellyfinVirtualFolderInfo) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *JellyfinVirtualFolderInfo) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *JellyfinVirtualFolderInfo) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetCollectionType

`func (o *JellyfinVirtualFolderInfo) GetCollectionType() JellyfinJellyfinCollectionTypeOptions`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *JellyfinVirtualFolderInfo) GetCollectionTypeOk() (*JellyfinJellyfinCollectionTypeOptions, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *JellyfinVirtualFolderInfo) SetCollectionType(v JellyfinJellyfinCollectionTypeOptions)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *JellyfinVirtualFolderInfo) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### SetCollectionTypeNil

`func (o *JellyfinVirtualFolderInfo) SetCollectionTypeNil(b bool)`

 SetCollectionTypeNil sets the value for CollectionType to be an explicit nil

### UnsetCollectionType
`func (o *JellyfinVirtualFolderInfo) UnsetCollectionType()`

UnsetCollectionType ensures that no value is present for CollectionType, not even an explicit nil
### GetLibraryOptions

`func (o *JellyfinVirtualFolderInfo) GetLibraryOptions() JellyfinJellyfinLibraryOptions`

GetLibraryOptions returns the LibraryOptions field if non-nil, zero value otherwise.

### GetLibraryOptionsOk

`func (o *JellyfinVirtualFolderInfo) GetLibraryOptionsOk() (*JellyfinJellyfinLibraryOptions, bool)`

GetLibraryOptionsOk returns a tuple with the LibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryOptions

`func (o *JellyfinVirtualFolderInfo) SetLibraryOptions(v JellyfinJellyfinLibraryOptions)`

SetLibraryOptions sets LibraryOptions field to given value.

### HasLibraryOptions

`func (o *JellyfinVirtualFolderInfo) HasLibraryOptions() bool`

HasLibraryOptions returns a boolean if a field has been set.

### SetLibraryOptionsNil

`func (o *JellyfinVirtualFolderInfo) SetLibraryOptionsNil(b bool)`

 SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil

### UnsetLibraryOptions
`func (o *JellyfinVirtualFolderInfo) UnsetLibraryOptions()`

UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil
### GetItemId

`func (o *JellyfinVirtualFolderInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinVirtualFolderInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinVirtualFolderInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinVirtualFolderInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *JellyfinVirtualFolderInfo) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *JellyfinVirtualFolderInfo) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetPrimaryImageItemId

`func (o *JellyfinVirtualFolderInfo) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *JellyfinVirtualFolderInfo) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *JellyfinVirtualFolderInfo) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *JellyfinVirtualFolderInfo) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### SetPrimaryImageItemIdNil

`func (o *JellyfinVirtualFolderInfo) SetPrimaryImageItemIdNil(b bool)`

 SetPrimaryImageItemIdNil sets the value for PrimaryImageItemId to be an explicit nil

### UnsetPrimaryImageItemId
`func (o *JellyfinVirtualFolderInfo) UnsetPrimaryImageItemId()`

UnsetPrimaryImageItemId ensures that no value is present for PrimaryImageItemId, not even an explicit nil
### GetRefreshProgress

`func (o *JellyfinVirtualFolderInfo) GetRefreshProgress() float64`

GetRefreshProgress returns the RefreshProgress field if non-nil, zero value otherwise.

### GetRefreshProgressOk

`func (o *JellyfinVirtualFolderInfo) GetRefreshProgressOk() (*float64, bool)`

GetRefreshProgressOk returns a tuple with the RefreshProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshProgress

`func (o *JellyfinVirtualFolderInfo) SetRefreshProgress(v float64)`

SetRefreshProgress sets RefreshProgress field to given value.

### HasRefreshProgress

`func (o *JellyfinVirtualFolderInfo) HasRefreshProgress() bool`

HasRefreshProgress returns a boolean if a field has been set.

### SetRefreshProgressNil

`func (o *JellyfinVirtualFolderInfo) SetRefreshProgressNil(b bool)`

 SetRefreshProgressNil sets the value for RefreshProgress to be an explicit nil

### UnsetRefreshProgress
`func (o *JellyfinVirtualFolderInfo) UnsetRefreshProgress()`

UnsetRefreshProgress ensures that no value is present for RefreshProgress, not even an explicit nil
### GetRefreshStatus

`func (o *JellyfinVirtualFolderInfo) GetRefreshStatus() string`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *JellyfinVirtualFolderInfo) GetRefreshStatusOk() (*string, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *JellyfinVirtualFolderInfo) SetRefreshStatus(v string)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *JellyfinVirtualFolderInfo) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.

### SetRefreshStatusNil

`func (o *JellyfinVirtualFolderInfo) SetRefreshStatusNil(b bool)`

 SetRefreshStatusNil sets the value for RefreshStatus to be an explicit nil

### UnsetRefreshStatus
`func (o *JellyfinVirtualFolderInfo) UnsetRefreshStatus()`

UnsetRefreshStatus ensures that no value is present for RefreshStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


