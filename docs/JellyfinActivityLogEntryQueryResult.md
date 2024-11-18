# JellyfinActivityLogEntryQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]JellyfinJellyfinActivityLogEntry**](JellyfinJellyfinActivityLogEntry.md) | Gets or sets the items. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of records available. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the index of the first record in Items. | [optional] 

## Methods

### NewJellyfinActivityLogEntryQueryResult

`func NewJellyfinActivityLogEntryQueryResult() *JellyfinActivityLogEntryQueryResult`

NewJellyfinActivityLogEntryQueryResult instantiates a new JellyfinActivityLogEntryQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinActivityLogEntryQueryResultWithDefaults

`func NewJellyfinActivityLogEntryQueryResultWithDefaults() *JellyfinActivityLogEntryQueryResult`

NewJellyfinActivityLogEntryQueryResultWithDefaults instantiates a new JellyfinActivityLogEntryQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *JellyfinActivityLogEntryQueryResult) GetItems() []JellyfinJellyfinActivityLogEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JellyfinActivityLogEntryQueryResult) GetItemsOk() (*[]JellyfinJellyfinActivityLogEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JellyfinActivityLogEntryQueryResult) SetItems(v []JellyfinJellyfinActivityLogEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *JellyfinActivityLogEntryQueryResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *JellyfinActivityLogEntryQueryResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *JellyfinActivityLogEntryQueryResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *JellyfinActivityLogEntryQueryResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *JellyfinActivityLogEntryQueryResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *JellyfinActivityLogEntryQueryResult) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *JellyfinActivityLogEntryQueryResult) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *JellyfinActivityLogEntryQueryResult) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *JellyfinActivityLogEntryQueryResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


