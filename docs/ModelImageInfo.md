# ModelImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageType** | Pointer to [**ModelModelImageType**](ModelImageType.md) | Gets or sets the type of the image. | [optional] 
**ImageIndex** | Pointer to **NullableInt32** | Gets or sets the index of the image. | [optional] 
**ImageTag** | Pointer to **NullableString** | Gets or sets the image tag. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**BlurHash** | Pointer to **NullableString** | Gets or sets the blurhash. | [optional] 
**Height** | Pointer to **NullableInt32** | Gets or sets the height. | [optional] 
**Width** | Pointer to **NullableInt32** | Gets or sets the width. | [optional] 
**Size** | Pointer to **int64** | Gets or sets the size. | [optional] 

## Methods

### NewModelImageInfo

`func NewModelImageInfo() *ModelImageInfo`

NewModelImageInfo instantiates a new ModelImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelImageInfoWithDefaults

`func NewModelImageInfoWithDefaults() *ModelImageInfo`

NewModelImageInfoWithDefaults instantiates a new ModelImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageType

`func (o *ModelImageInfo) GetImageType() ModelModelImageType`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ModelImageInfo) GetImageTypeOk() (*ModelModelImageType, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ModelImageInfo) SetImageType(v ModelModelImageType)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ModelImageInfo) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetImageIndex

`func (o *ModelImageInfo) GetImageIndex() int32`

GetImageIndex returns the ImageIndex field if non-nil, zero value otherwise.

### GetImageIndexOk

`func (o *ModelImageInfo) GetImageIndexOk() (*int32, bool)`

GetImageIndexOk returns a tuple with the ImageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIndex

`func (o *ModelImageInfo) SetImageIndex(v int32)`

SetImageIndex sets ImageIndex field to given value.

### HasImageIndex

`func (o *ModelImageInfo) HasImageIndex() bool`

HasImageIndex returns a boolean if a field has been set.

### SetImageIndexNil

`func (o *ModelImageInfo) SetImageIndexNil(b bool)`

 SetImageIndexNil sets the value for ImageIndex to be an explicit nil

### UnsetImageIndex
`func (o *ModelImageInfo) UnsetImageIndex()`

UnsetImageIndex ensures that no value is present for ImageIndex, not even an explicit nil
### GetImageTag

`func (o *ModelImageInfo) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ModelImageInfo) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ModelImageInfo) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ModelImageInfo) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### SetImageTagNil

`func (o *ModelImageInfo) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *ModelImageInfo) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
### GetPath

`func (o *ModelImageInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelImageInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelImageInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelImageInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ModelImageInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ModelImageInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetBlurHash

`func (o *ModelImageInfo) GetBlurHash() string`

GetBlurHash returns the BlurHash field if non-nil, zero value otherwise.

### GetBlurHashOk

`func (o *ModelImageInfo) GetBlurHashOk() (*string, bool)`

GetBlurHashOk returns a tuple with the BlurHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlurHash

`func (o *ModelImageInfo) SetBlurHash(v string)`

SetBlurHash sets BlurHash field to given value.

### HasBlurHash

`func (o *ModelImageInfo) HasBlurHash() bool`

HasBlurHash returns a boolean if a field has been set.

### SetBlurHashNil

`func (o *ModelImageInfo) SetBlurHashNil(b bool)`

 SetBlurHashNil sets the value for BlurHash to be an explicit nil

### UnsetBlurHash
`func (o *ModelImageInfo) UnsetBlurHash()`

UnsetBlurHash ensures that no value is present for BlurHash, not even an explicit nil
### GetHeight

`func (o *ModelImageInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelImageInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelImageInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelImageInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelImageInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelImageInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ModelImageInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelImageInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelImageInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelImageInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelImageInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelImageInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetSize

`func (o *ModelImageInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ModelImageInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ModelImageInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ModelImageInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


