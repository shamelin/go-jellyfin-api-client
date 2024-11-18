# ModelLibraryTypeOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Gets or sets the type. | [optional] 
**MetadataFetchers** | Pointer to [**[]ModelModelLibraryOptionInfoDto**](ModelModelLibraryOptionInfoDto.md) | Gets or sets the metadata fetchers. | [optional] 
**ImageFetchers** | Pointer to [**[]ModelModelLibraryOptionInfoDto**](ModelModelLibraryOptionInfoDto.md) | Gets or sets the image fetchers. | [optional] 
**SupportedImageTypes** | Pointer to [**[]ModelModelImageType**](ModelModelImageType.md) | Gets or sets the supported image types. | [optional] 
**DefaultImageOptions** | Pointer to [**[]ModelModelImageOption**](ModelModelImageOption.md) | Gets or sets the default image options. | [optional] 

## Methods

### NewModelLibraryTypeOptionsDto

`func NewModelLibraryTypeOptionsDto() *ModelLibraryTypeOptionsDto`

NewModelLibraryTypeOptionsDto instantiates a new ModelLibraryTypeOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLibraryTypeOptionsDtoWithDefaults

`func NewModelLibraryTypeOptionsDtoWithDefaults() *ModelLibraryTypeOptionsDto`

NewModelLibraryTypeOptionsDtoWithDefaults instantiates a new ModelLibraryTypeOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelLibraryTypeOptionsDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelLibraryTypeOptionsDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelLibraryTypeOptionsDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelLibraryTypeOptionsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelLibraryTypeOptionsDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelLibraryTypeOptionsDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMetadataFetchers

`func (o *ModelLibraryTypeOptionsDto) GetMetadataFetchers() []ModelModelLibraryOptionInfoDto`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *ModelLibraryTypeOptionsDto) GetMetadataFetchersOk() (*[]ModelModelLibraryOptionInfoDto, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *ModelLibraryTypeOptionsDto) SetMetadataFetchers(v []ModelModelLibraryOptionInfoDto)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *ModelLibraryTypeOptionsDto) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetImageFetchers

`func (o *ModelLibraryTypeOptionsDto) GetImageFetchers() []ModelModelLibraryOptionInfoDto`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *ModelLibraryTypeOptionsDto) GetImageFetchersOk() (*[]ModelModelLibraryOptionInfoDto, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *ModelLibraryTypeOptionsDto) SetImageFetchers(v []ModelModelLibraryOptionInfoDto)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *ModelLibraryTypeOptionsDto) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetSupportedImageTypes

`func (o *ModelLibraryTypeOptionsDto) GetSupportedImageTypes() []ModelModelImageType`

GetSupportedImageTypes returns the SupportedImageTypes field if non-nil, zero value otherwise.

### GetSupportedImageTypesOk

`func (o *ModelLibraryTypeOptionsDto) GetSupportedImageTypesOk() (*[]ModelModelImageType, bool)`

GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImageTypes

`func (o *ModelLibraryTypeOptionsDto) SetSupportedImageTypes(v []ModelModelImageType)`

SetSupportedImageTypes sets SupportedImageTypes field to given value.

### HasSupportedImageTypes

`func (o *ModelLibraryTypeOptionsDto) HasSupportedImageTypes() bool`

HasSupportedImageTypes returns a boolean if a field has been set.

### GetDefaultImageOptions

`func (o *ModelLibraryTypeOptionsDto) GetDefaultImageOptions() []ModelModelImageOption`

GetDefaultImageOptions returns the DefaultImageOptions field if non-nil, zero value otherwise.

### GetDefaultImageOptionsOk

`func (o *ModelLibraryTypeOptionsDto) GetDefaultImageOptionsOk() (*[]ModelModelImageOption, bool)`

GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageOptions

`func (o *ModelLibraryTypeOptionsDto) SetDefaultImageOptions(v []ModelModelImageOption)`

SetDefaultImageOptions sets DefaultImageOptions field to given value.

### HasDefaultImageOptions

`func (o *ModelLibraryTypeOptionsDto) HasDefaultImageOptions() bool`

HasDefaultImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


