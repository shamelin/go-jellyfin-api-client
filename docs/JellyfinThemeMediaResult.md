# JellyfinThemeMediaResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]JellyfinJellyfinBaseItemDto**](JellyfinJellyfinBaseItemDto.md) | Gets or sets the items. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of records available. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the index of the first record in Items. | [optional] 
**OwnerId** | Pointer to **string** | Gets or sets the owner id. | [optional] 

## Methods

### NewJellyfinThemeMediaResult

`func NewJellyfinThemeMediaResult() *JellyfinThemeMediaResult`

NewJellyfinThemeMediaResult instantiates a new JellyfinThemeMediaResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinThemeMediaResultWithDefaults

`func NewJellyfinThemeMediaResultWithDefaults() *JellyfinThemeMediaResult`

NewJellyfinThemeMediaResultWithDefaults instantiates a new JellyfinThemeMediaResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *JellyfinThemeMediaResult) GetItems() []JellyfinJellyfinBaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JellyfinThemeMediaResult) GetItemsOk() (*[]JellyfinJellyfinBaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JellyfinThemeMediaResult) SetItems(v []JellyfinJellyfinBaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *JellyfinThemeMediaResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *JellyfinThemeMediaResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *JellyfinThemeMediaResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *JellyfinThemeMediaResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *JellyfinThemeMediaResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *JellyfinThemeMediaResult) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *JellyfinThemeMediaResult) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *JellyfinThemeMediaResult) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *JellyfinThemeMediaResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetOwnerId

`func (o *JellyfinThemeMediaResult) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *JellyfinThemeMediaResult) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *JellyfinThemeMediaResult) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *JellyfinThemeMediaResult) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


