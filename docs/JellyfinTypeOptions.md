# JellyfinTypeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**MetadataFetchers** | Pointer to **[]string** |  | [optional] 
**MetadataFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageFetchers** | Pointer to **[]string** |  | [optional] 
**ImageFetcherOrder** | Pointer to **[]string** |  | [optional] 
**ImageOptions** | Pointer to [**[]JellyfinJellyfinImageOption**](JellyfinJellyfinImageOption.md) |  | [optional] 

## Methods

### NewJellyfinTypeOptions

`func NewJellyfinTypeOptions() *JellyfinTypeOptions`

NewJellyfinTypeOptions instantiates a new JellyfinTypeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinTypeOptionsWithDefaults

`func NewJellyfinTypeOptionsWithDefaults() *JellyfinTypeOptions`

NewJellyfinTypeOptionsWithDefaults instantiates a new JellyfinTypeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JellyfinTypeOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinTypeOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinTypeOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinTypeOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JellyfinTypeOptions) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JellyfinTypeOptions) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMetadataFetchers

`func (o *JellyfinTypeOptions) GetMetadataFetchers() []string`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *JellyfinTypeOptions) GetMetadataFetchersOk() (*[]string, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *JellyfinTypeOptions) SetMetadataFetchers(v []string)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *JellyfinTypeOptions) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### SetMetadataFetchersNil

`func (o *JellyfinTypeOptions) SetMetadataFetchersNil(b bool)`

 SetMetadataFetchersNil sets the value for MetadataFetchers to be an explicit nil

### UnsetMetadataFetchers
`func (o *JellyfinTypeOptions) UnsetMetadataFetchers()`

UnsetMetadataFetchers ensures that no value is present for MetadataFetchers, not even an explicit nil
### GetMetadataFetcherOrder

`func (o *JellyfinTypeOptions) GetMetadataFetcherOrder() []string`

GetMetadataFetcherOrder returns the MetadataFetcherOrder field if non-nil, zero value otherwise.

### GetMetadataFetcherOrderOk

`func (o *JellyfinTypeOptions) GetMetadataFetcherOrderOk() (*[]string, bool)`

GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetcherOrder

`func (o *JellyfinTypeOptions) SetMetadataFetcherOrder(v []string)`

SetMetadataFetcherOrder sets MetadataFetcherOrder field to given value.

### HasMetadataFetcherOrder

`func (o *JellyfinTypeOptions) HasMetadataFetcherOrder() bool`

HasMetadataFetcherOrder returns a boolean if a field has been set.

### SetMetadataFetcherOrderNil

`func (o *JellyfinTypeOptions) SetMetadataFetcherOrderNil(b bool)`

 SetMetadataFetcherOrderNil sets the value for MetadataFetcherOrder to be an explicit nil

### UnsetMetadataFetcherOrder
`func (o *JellyfinTypeOptions) UnsetMetadataFetcherOrder()`

UnsetMetadataFetcherOrder ensures that no value is present for MetadataFetcherOrder, not even an explicit nil
### GetImageFetchers

`func (o *JellyfinTypeOptions) GetImageFetchers() []string`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *JellyfinTypeOptions) GetImageFetchersOk() (*[]string, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *JellyfinTypeOptions) SetImageFetchers(v []string)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *JellyfinTypeOptions) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### SetImageFetchersNil

`func (o *JellyfinTypeOptions) SetImageFetchersNil(b bool)`

 SetImageFetchersNil sets the value for ImageFetchers to be an explicit nil

### UnsetImageFetchers
`func (o *JellyfinTypeOptions) UnsetImageFetchers()`

UnsetImageFetchers ensures that no value is present for ImageFetchers, not even an explicit nil
### GetImageFetcherOrder

`func (o *JellyfinTypeOptions) GetImageFetcherOrder() []string`

GetImageFetcherOrder returns the ImageFetcherOrder field if non-nil, zero value otherwise.

### GetImageFetcherOrderOk

`func (o *JellyfinTypeOptions) GetImageFetcherOrderOk() (*[]string, bool)`

GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetcherOrder

`func (o *JellyfinTypeOptions) SetImageFetcherOrder(v []string)`

SetImageFetcherOrder sets ImageFetcherOrder field to given value.

### HasImageFetcherOrder

`func (o *JellyfinTypeOptions) HasImageFetcherOrder() bool`

HasImageFetcherOrder returns a boolean if a field has been set.

### SetImageFetcherOrderNil

`func (o *JellyfinTypeOptions) SetImageFetcherOrderNil(b bool)`

 SetImageFetcherOrderNil sets the value for ImageFetcherOrder to be an explicit nil

### UnsetImageFetcherOrder
`func (o *JellyfinTypeOptions) UnsetImageFetcherOrder()`

UnsetImageFetcherOrder ensures that no value is present for ImageFetcherOrder, not even an explicit nil
### GetImageOptions

`func (o *JellyfinTypeOptions) GetImageOptions() []JellyfinJellyfinImageOption`

GetImageOptions returns the ImageOptions field if non-nil, zero value otherwise.

### GetImageOptionsOk

`func (o *JellyfinTypeOptions) GetImageOptionsOk() (*[]JellyfinJellyfinImageOption, bool)`

GetImageOptionsOk returns a tuple with the ImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptions

`func (o *JellyfinTypeOptions) SetImageOptions(v []JellyfinJellyfinImageOption)`

SetImageOptions sets ImageOptions field to given value.

### HasImageOptions

`func (o *JellyfinTypeOptions) HasImageOptions() bool`

HasImageOptions returns a boolean if a field has been set.

### SetImageOptionsNil

`func (o *JellyfinTypeOptions) SetImageOptionsNil(b bool)`

 SetImageOptionsNil sets the value for ImageOptions to be an explicit nil

### UnsetImageOptions
`func (o *JellyfinTypeOptions) UnsetImageOptions()`

UnsetImageOptions ensures that no value is present for ImageOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


