# ModelMetadataOptions

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

### NewModelMetadataOptions

`func NewModelMetadataOptions() *ModelMetadataOptions`

NewModelMetadataOptions instantiates a new ModelMetadataOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataOptionsWithDefaults

`func NewModelMetadataOptionsWithDefaults() *ModelMetadataOptions`

NewModelMetadataOptionsWithDefaults instantiates a new ModelMetadataOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *ModelMetadataOptions) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ModelMetadataOptions) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ModelMetadataOptions) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ModelMetadataOptions) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *ModelMetadataOptions) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *ModelMetadataOptions) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetDisabledMetadataSavers

`func (o *ModelMetadataOptions) GetDisabledMetadataSavers() []string`

GetDisabledMetadataSavers returns the DisabledMetadataSavers field if non-nil, zero value otherwise.

### GetDisabledMetadataSaversOk

`func (o *ModelMetadataOptions) GetDisabledMetadataSaversOk() (*[]string, bool)`

GetDisabledMetadataSaversOk returns a tuple with the DisabledMetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMetadataSavers

`func (o *ModelMetadataOptions) SetDisabledMetadataSavers(v []string)`

SetDisabledMetadataSavers sets DisabledMetadataSavers field to given value.

### HasDisabledMetadataSavers

`func (o *ModelMetadataOptions) HasDisabledMetadataSavers() bool`

HasDisabledMetadataSavers returns a boolean if a field has been set.

### SetDisabledMetadataSaversNil

`func (o *ModelMetadataOptions) SetDisabledMetadataSaversNil(b bool)`

 SetDisabledMetadataSaversNil sets the value for DisabledMetadataSavers to be an explicit nil

### UnsetDisabledMetadataSavers
`func (o *ModelMetadataOptions) UnsetDisabledMetadataSavers()`

UnsetDisabledMetadataSavers ensures that no value is present for DisabledMetadataSavers, not even an explicit nil
### GetLocalMetadataReaderOrder

`func (o *ModelMetadataOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *ModelMetadataOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *ModelMetadataOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *ModelMetadataOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *ModelMetadataOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *ModelMetadataOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledMetadataFetchers

`func (o *ModelMetadataOptions) GetDisabledMetadataFetchers() []string`

GetDisabledMetadataFetchers returns the DisabledMetadataFetchers field if non-nil, zero value otherwise.

### GetDisabledMetadataFetchersOk

`func (o *ModelMetadataOptions) GetDisabledMetadataFetchersOk() (*[]string, bool)`

GetDisabledMetadataFetchersOk returns a tuple with the DisabledMetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMetadataFetchers

`func (o *ModelMetadataOptions) SetDisabledMetadataFetchers(v []string)`

SetDisabledMetadataFetchers sets DisabledMetadataFetchers field to given value.

### HasDisabledMetadataFetchers

`func (o *ModelMetadataOptions) HasDisabledMetadataFetchers() bool`

HasDisabledMetadataFetchers returns a boolean if a field has been set.

### SetDisabledMetadataFetchersNil

`func (o *ModelMetadataOptions) SetDisabledMetadataFetchersNil(b bool)`

 SetDisabledMetadataFetchersNil sets the value for DisabledMetadataFetchers to be an explicit nil

### UnsetDisabledMetadataFetchers
`func (o *ModelMetadataOptions) UnsetDisabledMetadataFetchers()`

UnsetDisabledMetadataFetchers ensures that no value is present for DisabledMetadataFetchers, not even an explicit nil
### GetMetadataFetcherOrder

`func (o *ModelMetadataOptions) GetMetadataFetcherOrder() []string`

GetMetadataFetcherOrder returns the MetadataFetcherOrder field if non-nil, zero value otherwise.

### GetMetadataFetcherOrderOk

`func (o *ModelMetadataOptions) GetMetadataFetcherOrderOk() (*[]string, bool)`

GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetcherOrder

`func (o *ModelMetadataOptions) SetMetadataFetcherOrder(v []string)`

SetMetadataFetcherOrder sets MetadataFetcherOrder field to given value.

### HasMetadataFetcherOrder

`func (o *ModelMetadataOptions) HasMetadataFetcherOrder() bool`

HasMetadataFetcherOrder returns a boolean if a field has been set.

### SetMetadataFetcherOrderNil

`func (o *ModelMetadataOptions) SetMetadataFetcherOrderNil(b bool)`

 SetMetadataFetcherOrderNil sets the value for MetadataFetcherOrder to be an explicit nil

### UnsetMetadataFetcherOrder
`func (o *ModelMetadataOptions) UnsetMetadataFetcherOrder()`

UnsetMetadataFetcherOrder ensures that no value is present for MetadataFetcherOrder, not even an explicit nil
### GetDisabledImageFetchers

`func (o *ModelMetadataOptions) GetDisabledImageFetchers() []string`

GetDisabledImageFetchers returns the DisabledImageFetchers field if non-nil, zero value otherwise.

### GetDisabledImageFetchersOk

`func (o *ModelMetadataOptions) GetDisabledImageFetchersOk() (*[]string, bool)`

GetDisabledImageFetchersOk returns a tuple with the DisabledImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledImageFetchers

`func (o *ModelMetadataOptions) SetDisabledImageFetchers(v []string)`

SetDisabledImageFetchers sets DisabledImageFetchers field to given value.

### HasDisabledImageFetchers

`func (o *ModelMetadataOptions) HasDisabledImageFetchers() bool`

HasDisabledImageFetchers returns a boolean if a field has been set.

### SetDisabledImageFetchersNil

`func (o *ModelMetadataOptions) SetDisabledImageFetchersNil(b bool)`

 SetDisabledImageFetchersNil sets the value for DisabledImageFetchers to be an explicit nil

### UnsetDisabledImageFetchers
`func (o *ModelMetadataOptions) UnsetDisabledImageFetchers()`

UnsetDisabledImageFetchers ensures that no value is present for DisabledImageFetchers, not even an explicit nil
### GetImageFetcherOrder

`func (o *ModelMetadataOptions) GetImageFetcherOrder() []string`

GetImageFetcherOrder returns the ImageFetcherOrder field if non-nil, zero value otherwise.

### GetImageFetcherOrderOk

`func (o *ModelMetadataOptions) GetImageFetcherOrderOk() (*[]string, bool)`

GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetcherOrder

`func (o *ModelMetadataOptions) SetImageFetcherOrder(v []string)`

SetImageFetcherOrder sets ImageFetcherOrder field to given value.

### HasImageFetcherOrder

`func (o *ModelMetadataOptions) HasImageFetcherOrder() bool`

HasImageFetcherOrder returns a boolean if a field has been set.

### SetImageFetcherOrderNil

`func (o *ModelMetadataOptions) SetImageFetcherOrderNil(b bool)`

 SetImageFetcherOrderNil sets the value for ImageFetcherOrder to be an explicit nil

### UnsetImageFetcherOrder
`func (o *ModelMetadataOptions) UnsetImageFetcherOrder()`

UnsetImageFetcherOrder ensures that no value is present for ImageFetcherOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


