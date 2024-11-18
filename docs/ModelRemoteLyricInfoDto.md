# ModelRemoteLyricInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the id for the lyric. | [optional] 
**ProviderName** | Pointer to **string** | Gets the provider name. | [optional] 
**Lyrics** | Pointer to [**ModelModelLyricDto**](ModelLyricDto.md) | Gets the lyrics. | [optional] 

## Methods

### NewModelRemoteLyricInfoDto

`func NewModelRemoteLyricInfoDto() *ModelRemoteLyricInfoDto`

NewModelRemoteLyricInfoDto instantiates a new ModelRemoteLyricInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteLyricInfoDtoWithDefaults

`func NewModelRemoteLyricInfoDtoWithDefaults() *ModelRemoteLyricInfoDto`

NewModelRemoteLyricInfoDtoWithDefaults instantiates a new ModelRemoteLyricInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelRemoteLyricInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelRemoteLyricInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelRemoteLyricInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelRemoteLyricInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *ModelRemoteLyricInfoDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ModelRemoteLyricInfoDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ModelRemoteLyricInfoDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ModelRemoteLyricInfoDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLyrics

`func (o *ModelRemoteLyricInfoDto) GetLyrics() ModelModelLyricDto`

GetLyrics returns the Lyrics field if non-nil, zero value otherwise.

### GetLyricsOk

`func (o *ModelRemoteLyricInfoDto) GetLyricsOk() (*ModelModelLyricDto, bool)`

GetLyricsOk returns a tuple with the Lyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyrics

`func (o *ModelRemoteLyricInfoDto) SetLyrics(v ModelModelLyricDto)`

SetLyrics sets Lyrics field to given value.

### HasLyrics

`func (o *ModelRemoteLyricInfoDto) HasLyrics() bool`

HasLyrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


