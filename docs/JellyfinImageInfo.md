# JellyfinImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageType** | Pointer to [**JellyfinJellyfinImageType**](JellyfinImageType.md) | Gets or sets the type of the image. | [optional] 
**ImageIndex** | Pointer to **NullableInt32** | Gets or sets the index of the image. | [optional] 
**ImageTag** | Pointer to **NullableString** | Gets or sets the image tag. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**BlurHash** | Pointer to **NullableString** | Gets or sets the blurhash. | [optional] 
**Height** | Pointer to **NullableInt32** | Gets or sets the height. | [optional] 
**Width** | Pointer to **NullableInt32** | Gets or sets the width. | [optional] 
**Size** | Pointer to **int64** | Gets or sets the size. | [optional] 

## Methods

### NewJellyfinImageInfo

`func NewJellyfinImageInfo() *JellyfinImageInfo`

NewJellyfinImageInfo instantiates a new JellyfinImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinImageInfoWithDefaults

`func NewJellyfinImageInfoWithDefaults() *JellyfinImageInfo`

NewJellyfinImageInfoWithDefaults instantiates a new JellyfinImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageType

`func (o *JellyfinImageInfo) GetImageType() JellyfinJellyfinImageType`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *JellyfinImageInfo) GetImageTypeOk() (*JellyfinJellyfinImageType, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *JellyfinImageInfo) SetImageType(v JellyfinJellyfinImageType)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *JellyfinImageInfo) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetImageIndex

`func (o *JellyfinImageInfo) GetImageIndex() int32`

GetImageIndex returns the ImageIndex field if non-nil, zero value otherwise.

### GetImageIndexOk

`func (o *JellyfinImageInfo) GetImageIndexOk() (*int32, bool)`

GetImageIndexOk returns a tuple with the ImageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIndex

`func (o *JellyfinImageInfo) SetImageIndex(v int32)`

SetImageIndex sets ImageIndex field to given value.

### HasImageIndex

`func (o *JellyfinImageInfo) HasImageIndex() bool`

HasImageIndex returns a boolean if a field has been set.

### SetImageIndexNil

`func (o *JellyfinImageInfo) SetImageIndexNil(b bool)`

 SetImageIndexNil sets the value for ImageIndex to be an explicit nil

### UnsetImageIndex
`func (o *JellyfinImageInfo) UnsetImageIndex()`

UnsetImageIndex ensures that no value is present for ImageIndex, not even an explicit nil
### GetImageTag

`func (o *JellyfinImageInfo) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *JellyfinImageInfo) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *JellyfinImageInfo) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *JellyfinImageInfo) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### SetImageTagNil

`func (o *JellyfinImageInfo) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *JellyfinImageInfo) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
### GetPath

`func (o *JellyfinImageInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JellyfinImageInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JellyfinImageInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *JellyfinImageInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *JellyfinImageInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *JellyfinImageInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetBlurHash

`func (o *JellyfinImageInfo) GetBlurHash() string`

GetBlurHash returns the BlurHash field if non-nil, zero value otherwise.

### GetBlurHashOk

`func (o *JellyfinImageInfo) GetBlurHashOk() (*string, bool)`

GetBlurHashOk returns a tuple with the BlurHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlurHash

`func (o *JellyfinImageInfo) SetBlurHash(v string)`

SetBlurHash sets BlurHash field to given value.

### HasBlurHash

`func (o *JellyfinImageInfo) HasBlurHash() bool`

HasBlurHash returns a boolean if a field has been set.

### SetBlurHashNil

`func (o *JellyfinImageInfo) SetBlurHashNil(b bool)`

 SetBlurHashNil sets the value for BlurHash to be an explicit nil

### UnsetBlurHash
`func (o *JellyfinImageInfo) UnsetBlurHash()`

UnsetBlurHash ensures that no value is present for BlurHash, not even an explicit nil
### GetHeight

`func (o *JellyfinImageInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *JellyfinImageInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *JellyfinImageInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *JellyfinImageInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *JellyfinImageInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *JellyfinImageInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *JellyfinImageInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *JellyfinImageInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *JellyfinImageInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *JellyfinImageInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *JellyfinImageInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *JellyfinImageInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetSize

`func (o *JellyfinImageInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *JellyfinImageInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *JellyfinImageInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *JellyfinImageInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


