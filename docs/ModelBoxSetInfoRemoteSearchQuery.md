# ModelBoxSetInfoRemoteSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchInfo** | Pointer to [**NullableModelModelBoxSetInfo**](ModelBoxSetInfo.md) |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**SearchProviderName** | Pointer to **NullableString** | Gets or sets the provider name to search within if set. | [optional] 
**IncludeDisabledProviders** | Pointer to **bool** | Gets or sets a value indicating whether disabled providers should be included. | [optional] 

## Methods

### NewModelBoxSetInfoRemoteSearchQuery

`func NewModelBoxSetInfoRemoteSearchQuery() *ModelBoxSetInfoRemoteSearchQuery`

NewModelBoxSetInfoRemoteSearchQuery instantiates a new ModelBoxSetInfoRemoteSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBoxSetInfoRemoteSearchQueryWithDefaults

`func NewModelBoxSetInfoRemoteSearchQueryWithDefaults() *ModelBoxSetInfoRemoteSearchQuery`

NewModelBoxSetInfoRemoteSearchQueryWithDefaults instantiates a new ModelBoxSetInfoRemoteSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchInfo

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetSearchInfo() ModelModelBoxSetInfo`

GetSearchInfo returns the SearchInfo field if non-nil, zero value otherwise.

### GetSearchInfoOk

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetSearchInfoOk() (*ModelModelBoxSetInfo, bool)`

GetSearchInfoOk returns a tuple with the SearchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInfo

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetSearchInfo(v ModelModelBoxSetInfo)`

SetSearchInfo sets SearchInfo field to given value.

### HasSearchInfo

`func (o *ModelBoxSetInfoRemoteSearchQuery) HasSearchInfo() bool`

HasSearchInfo returns a boolean if a field has been set.

### SetSearchInfoNil

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetSearchInfoNil(b bool)`

 SetSearchInfoNil sets the value for SearchInfo to be an explicit nil

### UnsetSearchInfo
`func (o *ModelBoxSetInfoRemoteSearchQuery) UnsetSearchInfo()`

UnsetSearchInfo ensures that no value is present for SearchInfo, not even an explicit nil
### GetItemId

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelBoxSetInfoRemoteSearchQuery) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSearchProviderName

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *ModelBoxSetInfoRemoteSearchQuery) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### SetSearchProviderNameNil

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetSearchProviderNameNil(b bool)`

 SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil

### UnsetSearchProviderName
`func (o *ModelBoxSetInfoRemoteSearchQuery) UnsetSearchProviderName()`

UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
### GetIncludeDisabledProviders

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetIncludeDisabledProviders() bool`

GetIncludeDisabledProviders returns the IncludeDisabledProviders field if non-nil, zero value otherwise.

### GetIncludeDisabledProvidersOk

`func (o *ModelBoxSetInfoRemoteSearchQuery) GetIncludeDisabledProvidersOk() (*bool, bool)`

GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabledProviders

`func (o *ModelBoxSetInfoRemoteSearchQuery) SetIncludeDisabledProviders(v bool)`

SetIncludeDisabledProviders sets IncludeDisabledProviders field to given value.

### HasIncludeDisabledProviders

`func (o *ModelBoxSetInfoRemoteSearchQuery) HasIncludeDisabledProviders() bool`

HasIncludeDisabledProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


