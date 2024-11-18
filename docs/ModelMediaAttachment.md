# ModelMediaAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to **NullableString** | Gets or sets the codec. | [optional] 
**CodecTag** | Pointer to **NullableString** | Gets or sets the codec tag. | [optional] 
**Comment** | Pointer to **NullableString** | Gets or sets the comment. | [optional] 
**Index** | Pointer to **int32** | Gets or sets the index. | [optional] 
**FileName** | Pointer to **NullableString** | Gets or sets the filename. | [optional] 
**MimeType** | Pointer to **NullableString** | Gets or sets the MIME type. | [optional] 
**DeliveryUrl** | Pointer to **NullableString** | Gets or sets the delivery URL. | [optional] 

## Methods

### NewModelMediaAttachment

`func NewModelMediaAttachment() *ModelMediaAttachment`

NewModelMediaAttachment instantiates a new ModelMediaAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMediaAttachmentWithDefaults

`func NewModelMediaAttachmentWithDefaults() *ModelMediaAttachment`

NewModelMediaAttachmentWithDefaults instantiates a new ModelMediaAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *ModelMediaAttachment) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *ModelMediaAttachment) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *ModelMediaAttachment) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *ModelMediaAttachment) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *ModelMediaAttachment) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *ModelMediaAttachment) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetCodecTag

`func (o *ModelMediaAttachment) GetCodecTag() string`

GetCodecTag returns the CodecTag field if non-nil, zero value otherwise.

### GetCodecTagOk

`func (o *ModelMediaAttachment) GetCodecTagOk() (*string, bool)`

GetCodecTagOk returns a tuple with the CodecTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTag

`func (o *ModelMediaAttachment) SetCodecTag(v string)`

SetCodecTag sets CodecTag field to given value.

### HasCodecTag

`func (o *ModelMediaAttachment) HasCodecTag() bool`

HasCodecTag returns a boolean if a field has been set.

### SetCodecTagNil

`func (o *ModelMediaAttachment) SetCodecTagNil(b bool)`

 SetCodecTagNil sets the value for CodecTag to be an explicit nil

### UnsetCodecTag
`func (o *ModelMediaAttachment) UnsetCodecTag()`

UnsetCodecTag ensures that no value is present for CodecTag, not even an explicit nil
### GetComment

`func (o *ModelMediaAttachment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelMediaAttachment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelMediaAttachment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelMediaAttachment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelMediaAttachment) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelMediaAttachment) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIndex

`func (o *ModelMediaAttachment) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ModelMediaAttachment) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ModelMediaAttachment) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ModelMediaAttachment) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetFileName

`func (o *ModelMediaAttachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ModelMediaAttachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ModelMediaAttachment) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ModelMediaAttachment) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ModelMediaAttachment) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ModelMediaAttachment) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetMimeType

`func (o *ModelMediaAttachment) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ModelMediaAttachment) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ModelMediaAttachment) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ModelMediaAttachment) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *ModelMediaAttachment) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *ModelMediaAttachment) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetDeliveryUrl

`func (o *ModelMediaAttachment) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ModelMediaAttachment) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ModelMediaAttachment) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ModelMediaAttachment) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### SetDeliveryUrlNil

`func (o *ModelMediaAttachment) SetDeliveryUrlNil(b bool)`

 SetDeliveryUrlNil sets the value for DeliveryUrl to be an explicit nil

### UnsetDeliveryUrl
`func (o *ModelMediaAttachment) UnsetDeliveryUrl()`

UnsetDeliveryUrl ensures that no value is present for DeliveryUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


