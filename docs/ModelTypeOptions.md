# ModelTypeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**MetadataFetchers** | Pointer to **[]string** |  | [optional] 
**MetadataFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageFetchers** | Pointer to **[]string** |  | [optional] 
**ImageFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageOptions** | Pointer to [**[]ModelModelImageOption**](ModelModelImageOption.md) |  | [optional] 

## Methods

### NewModelTypeOptions

`func NewModelTypeOptions() *ModelTypeOptions`

NewModelTypeOptions instantiates a new ModelTypeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTypeOptionsWithDefaults

`func NewModelTypeOptionsWithDefaults() *ModelTypeOptions`

NewModelTypeOptionsWithDefaults instantiates a new ModelTypeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelTypeOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelTypeOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelTypeOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelTypeOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelTypeOptions) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelTypeOptions) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMetadataFetchers

`func (o *ModelTypeOptions) GetMetadataFetchers() []string`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *ModelTypeOptions) GetMetadataFetchersOk() (*[]string, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *ModelTypeOptions) SetMetadataFetchers(v []string)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *ModelTypeOptions) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### SetMetadataFetchersNil

`func (o *ModelTypeOptions) SetMetadataFetchersNil(b bool)`

 SetMetadataFetchersNil sets the value for MetadataFetchers to be an explicit nil

### UnsetMetadataFetchers
`func (o *ModelTypeOptions) UnsetMetadataFetchers()`

UnsetMetadataFetchers ensures that no value is present for MetadataFetchers, not even an explicit nil
### GetMetadataFetcherOrder

`func (o *ModelTypeOptions) GetMetadataFetcherOrder() []string`

GetMetadataFetcherOrder returns the MetadataFetcherOrder field if non-nil, zero value otherwise.

### GetMetadataFetcherOrderOk

`func (o *ModelTypeOptions) GetMetadataFetcherOrderOk() (*[]string, bool)`

GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetcherOrder

`func (o *ModelTypeOptions) SetMetadataFetcherOrder(v []string)`

SetMetadataFetcherOrder sets MetadataFetcherOrder field to given value.

### HasMetadataFetcherOrder

`func (o *ModelTypeOptions) HasMetadataFetcherOrder() bool`

HasMetadataFetcherOrder returns a boolean if a field has been set.

### SetMetadataFetcherOrderNil

`func (o *ModelTypeOptions) SetMetadataFetcherOrderNil(b bool)`

 SetMetadataFetcherOrderNil sets the value for MetadataFetcherOrder to be an explicit nil

### UnsetMetadataFetcherOrder
`func (o *ModelTypeOptions) UnsetMetadataFetcherOrder()`

UnsetMetadataFetcherOrder ensures that no value is present for MetadataFetcherOrder, not even an explicit nil
### GetImageFetchers

`func (o *ModelTypeOptions) GetImageFetchers() []string`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *ModelTypeOptions) GetImageFetchersOk() (*[]string, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *ModelTypeOptions) SetImageFetchers(v []string)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *ModelTypeOptions) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### SetImageFetchersNil

`func (o *ModelTypeOptions) SetImageFetchersNil(b bool)`

 SetImageFetchersNil sets the value for ImageFetchers to be an explicit nil

### UnsetImageFetchers
`func (o *ModelTypeOptions) UnsetImageFetchers()`

UnsetImageFetchers ensures that no value is present for ImageFetchers, not even an explicit nil
### GetImageFetcherOrder

`func (o *ModelTypeOptions) GetImageFetcherOrder() []string`

GetImageFetcherOrder returns the ImageFetcherOrder field if non-nil, zero value otherwise.

### GetImageFetcherOrderOk

`func (o *ModelTypeOptions) GetImageFetcherOrderOk() (*[]string, bool)`

GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetcherOrder

`func (o *ModelTypeOptions) SetImageFetcherOrder(v []string)`

SetImageFetcherOrder sets ImageFetcherOrder field to given value.

### HasImageFetcherOrder

`func (o *ModelTypeOptions) HasImageFetcherOrder() bool`

HasImageFetcherOrder returns a boolean if a field has been set.

### SetImageFetcherOrderNil

`func (o *ModelTypeOptions) SetImageFetcherOrderNil(b bool)`

 SetImageFetcherOrderNil sets the value for ImageFetcherOrder to be an explicit nil

### UnsetImageFetcherOrder
`func (o *ModelTypeOptions) UnsetImageFetcherOrder()`

UnsetImageFetcherOrder ensures that no value is present for ImageFetcherOrder, not even an explicit nil
### GetImageOptions

`func (o *ModelTypeOptions) GetImageOptions() []ModelModelImageOption`

GetImageOptions returns the ImageOptions field if non-nil, zero value otherwise.

### GetImageOptionsOk

`func (o *ModelTypeOptions) GetImageOptionsOk() (*[]ModelModelImageOption, bool)`

GetImageOptionsOk returns a tuple with the ImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptions

`func (o *ModelTypeOptions) SetImageOptions(v []ModelModelImageOption)`

SetImageOptions sets ImageOptions field to given value.

### HasImageOptions

`func (o *ModelTypeOptions) HasImageOptions() bool`

HasImageOptions returns a boolean if a field has been set.

### SetImageOptionsNil

`func (o *ModelTypeOptions) SetImageOptionsNil(b bool)`

 SetImageOptionsNil sets the value for ImageOptions to be an explicit nil

### UnsetImageOptions
`func (o *ModelTypeOptions) UnsetImageOptions()`

UnsetImageOptions ensures that no value is present for ImageOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


