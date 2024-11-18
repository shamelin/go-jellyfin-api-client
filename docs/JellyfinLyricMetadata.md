# JellyfinLyricMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artist** | Pointer to **NullableString** | Gets or sets the song artist. | [optional] 
**Album** | Pointer to **NullableString** | Gets or sets the album this song is on. | [optional] 
**Title** | Pointer to **NullableString** | Gets or sets the title of the song. | [optional] 
**Author** | Pointer to **NullableString** | Gets or sets the author of the lyric data. | [optional] 
**Length** | Pointer to **NullableInt64** | Gets or sets the length of the song in ticks. | [optional] 
**By** | Pointer to **NullableString** | Gets or sets who the LRC file was created by. | [optional] 
**Offset** | Pointer to **NullableInt64** | Gets or sets the lyric offset compared to audio in ticks. | [optional] 
**Creator** | Pointer to **NullableString** | Gets or sets the software used to create the LRC file. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the version of the creator used. | [optional] 
**IsSynced** | Pointer to **NullableBool** | Gets or sets a value indicating whether this lyric is synced. | [optional] 

## Methods

### NewJellyfinLyricMetadata

`func NewJellyfinLyricMetadata() *JellyfinLyricMetadata`

NewJellyfinLyricMetadata instantiates a new JellyfinLyricMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLyricMetadataWithDefaults

`func NewJellyfinLyricMetadataWithDefaults() *JellyfinLyricMetadata`

NewJellyfinLyricMetadataWithDefaults instantiates a new JellyfinLyricMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtist

`func (o *JellyfinLyricMetadata) GetArtist() string`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *JellyfinLyricMetadata) GetArtistOk() (*string, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *JellyfinLyricMetadata) SetArtist(v string)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *JellyfinLyricMetadata) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### SetArtistNil

`func (o *JellyfinLyricMetadata) SetArtistNil(b bool)`

 SetArtistNil sets the value for Artist to be an explicit nil

### UnsetArtist
`func (o *JellyfinLyricMetadata) UnsetArtist()`

UnsetArtist ensures that no value is present for Artist, not even an explicit nil
### GetAlbum

`func (o *JellyfinLyricMetadata) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *JellyfinLyricMetadata) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *JellyfinLyricMetadata) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *JellyfinLyricMetadata) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *JellyfinLyricMetadata) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *JellyfinLyricMetadata) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetTitle

`func (o *JellyfinLyricMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JellyfinLyricMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JellyfinLyricMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *JellyfinLyricMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *JellyfinLyricMetadata) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *JellyfinLyricMetadata) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *JellyfinLyricMetadata) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *JellyfinLyricMetadata) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *JellyfinLyricMetadata) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *JellyfinLyricMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *JellyfinLyricMetadata) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *JellyfinLyricMetadata) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetLength

`func (o *JellyfinLyricMetadata) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *JellyfinLyricMetadata) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *JellyfinLyricMetadata) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *JellyfinLyricMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *JellyfinLyricMetadata) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *JellyfinLyricMetadata) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetBy

`func (o *JellyfinLyricMetadata) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *JellyfinLyricMetadata) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *JellyfinLyricMetadata) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *JellyfinLyricMetadata) HasBy() bool`

HasBy returns a boolean if a field has been set.

### SetByNil

`func (o *JellyfinLyricMetadata) SetByNil(b bool)`

 SetByNil sets the value for By to be an explicit nil

### UnsetBy
`func (o *JellyfinLyricMetadata) UnsetBy()`

UnsetBy ensures that no value is present for By, not even an explicit nil
### GetOffset

`func (o *JellyfinLyricMetadata) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *JellyfinLyricMetadata) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *JellyfinLyricMetadata) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *JellyfinLyricMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *JellyfinLyricMetadata) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *JellyfinLyricMetadata) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetCreator

`func (o *JellyfinLyricMetadata) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *JellyfinLyricMetadata) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *JellyfinLyricMetadata) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *JellyfinLyricMetadata) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *JellyfinLyricMetadata) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *JellyfinLyricMetadata) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetVersion

`func (o *JellyfinLyricMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JellyfinLyricMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JellyfinLyricMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JellyfinLyricMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *JellyfinLyricMetadata) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *JellyfinLyricMetadata) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIsSynced

`func (o *JellyfinLyricMetadata) GetIsSynced() bool`

GetIsSynced returns the IsSynced field if non-nil, zero value otherwise.

### GetIsSyncedOk

`func (o *JellyfinLyricMetadata) GetIsSyncedOk() (*bool, bool)`

GetIsSyncedOk returns a tuple with the IsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynced

`func (o *JellyfinLyricMetadata) SetIsSynced(v bool)`

SetIsSynced sets IsSynced field to given value.

### HasIsSynced

`func (o *JellyfinLyricMetadata) HasIsSynced() bool`

HasIsSynced returns a boolean if a field has been set.

### SetIsSyncedNil

`func (o *JellyfinLyricMetadata) SetIsSyncedNil(b bool)`

 SetIsSyncedNil sets the value for IsSynced to be an explicit nil

### UnsetIsSynced
`func (o *JellyfinLyricMetadata) UnsetIsSynced()`

UnsetIsSynced ensures that no value is present for IsSynced, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


