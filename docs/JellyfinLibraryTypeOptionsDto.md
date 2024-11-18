# JellyfinLibraryTypeOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Gets or sets the type. | [optional] 
**MetadataFetchers** | Pointer to [**[]JellyfinJellyfinLibraryOptionInfoDto**](JellyfinJellyfinLibraryOptionInfoDto.md) | Gets or sets the metadata fetchers. | [optional] 
**ImageFetchers** | Pointer to [**[]JellyfinJellyfinLibraryOptionInfoDto**](JellyfinJellyfinLibraryOptionInfoDto.md) | Gets or sets the image fetchers. | [optional] 
**SupportedImageTypes** | Pointer to [**[]JellyfinJellyfinImageType**](JellyfinJellyfinImageType.md) | Gets or sets the supported image types. | [optional] 
**DefaultImageOptions** | Pointer to [**[]JellyfinJellyfinImageOption**](JellyfinJellyfinImageOption.md) | Gets or sets the default image options. | [optional] 

## Methods

### NewJellyfinLibraryTypeOptionsDto

`func NewJellyfinLibraryTypeOptionsDto() *JellyfinLibraryTypeOptionsDto`

NewJellyfinLibraryTypeOptionsDto instantiates a new JellyfinLibraryTypeOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLibraryTypeOptionsDtoWithDefaults

`func NewJellyfinLibraryTypeOptionsDtoWithDefaults() *JellyfinLibraryTypeOptionsDto`

NewJellyfinLibraryTypeOptionsDtoWithDefaults instantiates a new JellyfinLibraryTypeOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JellyfinLibraryTypeOptionsDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinLibraryTypeOptionsDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinLibraryTypeOptionsDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinLibraryTypeOptionsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JellyfinLibraryTypeOptionsDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JellyfinLibraryTypeOptionsDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMetadataFetchers

`func (o *JellyfinLibraryTypeOptionsDto) GetMetadataFetchers() []JellyfinJellyfinLibraryOptionInfoDto`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *JellyfinLibraryTypeOptionsDto) GetMetadataFetchersOk() (*[]JellyfinJellyfinLibraryOptionInfoDto, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *JellyfinLibraryTypeOptionsDto) SetMetadataFetchers(v []JellyfinJellyfinLibraryOptionInfoDto)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *JellyfinLibraryTypeOptionsDto) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetImageFetchers

`func (o *JellyfinLibraryTypeOptionsDto) GetImageFetchers() []JellyfinJellyfinLibraryOptionInfoDto`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *JellyfinLibraryTypeOptionsDto) GetImageFetchersOk() (*[]JellyfinJellyfinLibraryOptionInfoDto, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *JellyfinLibraryTypeOptionsDto) SetImageFetchers(v []JellyfinJellyfinLibraryOptionInfoDto)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *JellyfinLibraryTypeOptionsDto) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetSupportedImageTypes

`func (o *JellyfinLibraryTypeOptionsDto) GetSupportedImageTypes() []JellyfinJellyfinImageType`

GetSupportedImageTypes returns the SupportedImageTypes field if non-nil, zero value otherwise.

### GetSupportedImageTypesOk

`func (o *JellyfinLibraryTypeOptionsDto) GetSupportedImageTypesOk() (*[]JellyfinJellyfinImageType, bool)`

GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImageTypes

`func (o *JellyfinLibraryTypeOptionsDto) SetSupportedImageTypes(v []JellyfinJellyfinImageType)`

SetSupportedImageTypes sets SupportedImageTypes field to given value.

### HasSupportedImageTypes

`func (o *JellyfinLibraryTypeOptionsDto) HasSupportedImageTypes() bool`

HasSupportedImageTypes returns a boolean if a field has been set.

### GetDefaultImageOptions

`func (o *JellyfinLibraryTypeOptionsDto) GetDefaultImageOptions() []JellyfinJellyfinImageOption`

GetDefaultImageOptions returns the DefaultImageOptions field if non-nil, zero value otherwise.

### GetDefaultImageOptionsOk

`func (o *JellyfinLibraryTypeOptionsDto) GetDefaultImageOptionsOk() (*[]JellyfinJellyfinImageOption, bool)`

GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageOptions

`func (o *JellyfinLibraryTypeOptionsDto) SetDefaultImageOptions(v []JellyfinJellyfinImageOption)`

SetDefaultImageOptions sets DefaultImageOptions field to given value.

### HasDefaultImageOptions

`func (o *JellyfinLibraryTypeOptionsDto) HasDefaultImageOptions() bool`

HasDefaultImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


