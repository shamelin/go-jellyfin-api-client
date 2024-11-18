# JellyfinAuthenticationInfoQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]JellyfinJellyfinAuthenticationInfo**](JellyfinJellyfinAuthenticationInfo.md) | Gets or sets the items. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of records available. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the index of the first record in Items. | [optional] 

## Methods

### NewJellyfinAuthenticationInfoQueryResult

`func NewJellyfinAuthenticationInfoQueryResult() *JellyfinAuthenticationInfoQueryResult`

NewJellyfinAuthenticationInfoQueryResult instantiates a new JellyfinAuthenticationInfoQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinAuthenticationInfoQueryResultWithDefaults

`func NewJellyfinAuthenticationInfoQueryResultWithDefaults() *JellyfinAuthenticationInfoQueryResult`

NewJellyfinAuthenticationInfoQueryResultWithDefaults instantiates a new JellyfinAuthenticationInfoQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *JellyfinAuthenticationInfoQueryResult) GetItems() []JellyfinJellyfinAuthenticationInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JellyfinAuthenticationInfoQueryResult) GetItemsOk() (*[]JellyfinJellyfinAuthenticationInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JellyfinAuthenticationInfoQueryResult) SetItems(v []JellyfinJellyfinAuthenticationInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *JellyfinAuthenticationInfoQueryResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *JellyfinAuthenticationInfoQueryResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *JellyfinAuthenticationInfoQueryResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *JellyfinAuthenticationInfoQueryResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *JellyfinAuthenticationInfoQueryResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *JellyfinAuthenticationInfoQueryResult) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *JellyfinAuthenticationInfoQueryResult) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *JellyfinAuthenticationInfoQueryResult) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *JellyfinAuthenticationInfoQueryResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


