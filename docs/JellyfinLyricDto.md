# JellyfinLyricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**JellyfinJellyfinLyricMetadata**](JellyfinLyricMetadata.md) | Gets or sets Metadata for the lyrics. | [optional] 
**Lyrics** | Pointer to [**[]JellyfinJellyfinLyricLine**](JellyfinJellyfinLyricLine.md) | Gets or sets a collection of individual lyric lines. | [optional] 

## Methods

### NewJellyfinLyricDto

`func NewJellyfinLyricDto() *JellyfinLyricDto`

NewJellyfinLyricDto instantiates a new JellyfinLyricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLyricDtoWithDefaults

`func NewJellyfinLyricDtoWithDefaults() *JellyfinLyricDto`

NewJellyfinLyricDtoWithDefaults instantiates a new JellyfinLyricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *JellyfinLyricDto) GetMetadata() JellyfinJellyfinLyricMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JellyfinLyricDto) GetMetadataOk() (*JellyfinJellyfinLyricMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JellyfinLyricDto) SetMetadata(v JellyfinJellyfinLyricMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *JellyfinLyricDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLyrics

`func (o *JellyfinLyricDto) GetLyrics() []JellyfinJellyfinLyricLine`

GetLyrics returns the Lyrics field if non-nil, zero value otherwise.

### GetLyricsOk

`func (o *JellyfinLyricDto) GetLyricsOk() (*[]JellyfinJellyfinLyricLine, bool)`

GetLyricsOk returns a tuple with the Lyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyrics

`func (o *JellyfinLyricDto) SetLyrics(v []JellyfinJellyfinLyricLine)`

SetLyrics sets Lyrics field to given value.

### HasLyrics

`func (o *JellyfinLyricDto) HasLyrics() bool`

HasLyrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


