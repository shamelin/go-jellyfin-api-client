# JellyfinMetadataOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | Pointer to **NullableString** |  | [optional] 
**DisabledMetadataSavers** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledMetadataFetchers** | Pointer to **[]string** |  | [optional] 
**MetadataFetcherOrder** | Pointer to **[]string** |  | [optional] 
**DisabledImageFetchers** | Pointer to **[]string** |  | [optional] 
**ImageFetcherOrder** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJellyfinMetadataOptions

`func NewJellyfinMetadataOptions() *JellyfinMetadataOptions`

NewJellyfinMetadataOptions instantiates a new JellyfinMetadataOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinMetadataOptionsWithDefaults

`func NewJellyfinMetadataOptionsWithDefaults() *JellyfinMetadataOptions`

NewJellyfinMetadataOptionsWithDefaults instantiates a new JellyfinMetadataOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *JellyfinMetadataOptions) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *JellyfinMetadataOptions) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *JellyfinMetadataOptions) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *JellyfinMetadataOptions) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *JellyfinMetadataOptions) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *JellyfinMetadataOptions) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetDisabledMetadataSavers

`func (o *JellyfinMetadataOptions) GetDisabledMetadataSavers() []string`

GetDisabledMetadataSavers returns the DisabledMetadataSavers field if non-nil, zero value otherwise.

### GetDisabledMetadataSaversOk

`func (o *JellyfinMetadataOptions) GetDisabledMetadataSaversOk() (*[]string, bool)`

GetDisabledMetadataSaversOk returns a tuple with the DisabledMetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMetadataSavers

`func (o *JellyfinMetadataOptions) SetDisabledMetadataSavers(v []string)`

SetDisabledMetadataSavers sets DisabledMetadataSavers field to given value.

### HasDisabledMetadataSavers

`func (o *JellyfinMetadataOptions) HasDisabledMetadataSavers() bool`

HasDisabledMetadataSavers returns a boolean if a field has been set.

### SetDisabledMetadataSaversNil

`func (o *JellyfinMetadataOptions) SetDisabledMetadataSaversNil(b bool)`

 SetDisabledMetadataSaversNil sets the value for DisabledMetadataSavers to be an explicit nil

### UnsetDisabledMetadataSavers
`func (o *JellyfinMetadataOptions) UnsetDisabledMetadataSavers()`

UnsetDisabledMetadataSavers ensures that no value is present for DisabledMetadataSavers, not even an explicit nil
### GetLocalMetadataReaderOrder

`func (o *JellyfinMetadataOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *JellyfinMetadataOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *JellyfinMetadataOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *JellyfinMetadataOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *JellyfinMetadataOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *JellyfinMetadataOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledMetadataFetchers

`func (o *JellyfinMetadataOptions) GetDisabledMetadataFetchers() []string`

GetDisabledMetadataFetchers returns the DisabledMetadataFetchers field if non-nil, zero value otherwise.

### GetDisabledMetadataFetchersOk

`func (o *JellyfinMetadataOptions) GetDisabledMetadataFetchersOk() (*[]string, bool)`

GetDisabledMetadataFetchersOk returns a tuple with the DisabledMetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMetadataFetchers

`func (o *JellyfinMetadataOptions) SetDisabledMetadataFetchers(v []string)`

SetDisabledMetadataFetchers sets DisabledMetadataFetchers field to given value.

### HasDisabledMetadataFetchers

`func (o *JellyfinMetadataOptions) HasDisabledMetadataFetchers() bool`

HasDisabledMetadataFetchers returns a boolean if a field has been set.

### SetDisabledMetadataFetchersNil

`func (o *JellyfinMetadataOptions) SetDisabledMetadataFetchersNil(b bool)`

 SetDisabledMetadataFetchersNil sets the value for DisabledMetadataFetchers to be an explicit nil

### UnsetDisabledMetadataFetchers
`func (o *JellyfinMetadataOptions) UnsetDisabledMetadataFetchers()`

UnsetDisabledMetadataFetchers ensures that no value is present for DisabledMetadataFetchers, not even an explicit nil
### GetMetadataFetcherOrder

`func (o *JellyfinMetadataOptions) GetMetadataFetcherOrder() []string`

GetMetadataFetcherOrder returns the MetadataFetcherOrder field if non-nil, zero value otherwise.

### GetMetadataFetcherOrderOk

`func (o *JellyfinMetadataOptions) GetMetadataFetcherOrderOk() (*[]string, bool)`

GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetcherOrder

`func (o *JellyfinMetadataOptions) SetMetadataFetcherOrder(v []string)`

SetMetadataFetcherOrder sets MetadataFetcherOrder field to given value.

### HasMetadataFetcherOrder

`func (o *JellyfinMetadataOptions) HasMetadataFetcherOrder() bool`

HasMetadataFetcherOrder returns a boolean if a field has been set.

### SetMetadataFetcherOrderNil

`func (o *JellyfinMetadataOptions) SetMetadataFetcherOrderNil(b bool)`

 SetMetadataFetcherOrderNil sets the value for MetadataFetcherOrder to be an explicit nil

### UnsetMetadataFetcherOrder
`func (o *JellyfinMetadataOptions) UnsetMetadataFetcherOrder()`

UnsetMetadataFetcherOrder ensures that no value is present for MetadataFetcherOrder, not even an explicit nil
### GetDisabledImageFetchers

`func (o *JellyfinMetadataOptions) GetDisabledImageFetchers() []string`

GetDisabledImageFetchers returns the DisabledImageFetchers field if non-nil, zero value otherwise.

### GetDisabledImageFetchersOk

`func (o *JellyfinMetadataOptions) GetDisabledImageFetchersOk() (*[]string, bool)`

GetDisabledImageFetchersOk returns a tuple with the DisabledImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledImageFetchers

`func (o *JellyfinMetadataOptions) SetDisabledImageFetchers(v []string)`

SetDisabledImageFetchers sets DisabledImageFetchers field to given value.

### HasDisabledImageFetchers

`func (o *JellyfinMetadataOptions) HasDisabledImageFetchers() bool`

HasDisabledImageFetchers returns a boolean if a field has been set.

### SetDisabledImageFetchersNil

`func (o *JellyfinMetadataOptions) SetDisabledImageFetchersNil(b bool)`

 SetDisabledImageFetchersNil sets the value for DisabledImageFetchers to be an explicit nil

### UnsetDisabledImageFetchers
`func (o *JellyfinMetadataOptions) UnsetDisabledImageFetchers()`

UnsetDisabledImageFetchers ensures that no value is present for DisabledImageFetchers, not even an explicit nil
### GetImageFetcherOrder

`func (o *JellyfinMetadataOptions) GetImageFetcherOrder() []string`

GetImageFetcherOrder returns the ImageFetcherOrder field if non-nil, zero value otherwise.

### GetImageFetcherOrderOk

`func (o *JellyfinMetadataOptions) GetImageFetcherOrderOk() (*[]string, bool)`

GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetcherOrder

`func (o *JellyfinMetadataOptions) SetImageFetcherOrder(v []string)`

SetImageFetcherOrder sets ImageFetcherOrder field to given value.

### HasImageFetcherOrder

`func (o *JellyfinMetadataOptions) HasImageFetcherOrder() bool`

HasImageFetcherOrder returns a boolean if a field has been set.

### SetImageFetcherOrderNil

`func (o *JellyfinMetadataOptions) SetImageFetcherOrderNil(b bool)`

 SetImageFetcherOrderNil sets the value for ImageFetcherOrder to be an explicit nil

### UnsetImageFetcherOrder
`func (o *JellyfinMetadataOptions) UnsetImageFetcherOrder()`

UnsetImageFetcherOrder ensures that no value is present for ImageFetcherOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


