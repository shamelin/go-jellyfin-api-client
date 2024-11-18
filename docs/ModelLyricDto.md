# ModelLyricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ModelModelLyricMetadata**](ModelLyricMetadata.md) | Gets or sets Metadata for the lyrics. | [optional] 
**Lyrics** | Pointer to [**[]ModelModelLyricLine**](ModelModelLyricLine.md) | Gets or sets a collection of individual lyric lines. | [optional] 

## Methods

### NewModelLyricDto

`func NewModelLyricDto() *ModelLyricDto`

NewModelLyricDto instantiates a new ModelLyricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLyricDtoWithDefaults

`func NewModelLyricDtoWithDefaults() *ModelLyricDto`

NewModelLyricDtoWithDefaults instantiates a new ModelLyricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ModelLyricDto) GetMetadata() ModelModelLyricMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelLyricDto) GetMetadataOk() (*ModelModelLyricMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelLyricDto) SetMetadata(v ModelModelLyricMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelLyricDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLyrics

`func (o *ModelLyricDto) GetLyrics() []ModelModelLyricLine`

GetLyrics returns the Lyrics field if non-nil, zero value otherwise.

### GetLyricsOk

`func (o *ModelLyricDto) GetLyricsOk() (*[]ModelModelLyricLine, bool)`

GetLyricsOk returns a tuple with the Lyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyrics

`func (o *ModelLyricDto) SetLyrics(v []ModelModelLyricLine)`

SetLyrics sets Lyrics field to given value.

### HasLyrics

`func (o *ModelLyricDto) HasLyrics() bool`

HasLyrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


