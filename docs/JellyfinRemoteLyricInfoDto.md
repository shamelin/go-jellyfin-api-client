# JellyfinRemoteLyricInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the id for the lyric. | [optional] 
**ProviderName** | Pointer to **string** | Gets the provider name. | [optional] 
**Lyrics** | Pointer to [**JellyfinJellyfinLyricDto**](JellyfinLyricDto.md) | Gets the lyrics. | [optional] 

## Methods

### NewJellyfinRemoteLyricInfoDto

`func NewJellyfinRemoteLyricInfoDto() *JellyfinRemoteLyricInfoDto`

NewJellyfinRemoteLyricInfoDto instantiates a new JellyfinRemoteLyricInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinRemoteLyricInfoDtoWithDefaults

`func NewJellyfinRemoteLyricInfoDtoWithDefaults() *JellyfinRemoteLyricInfoDto`

NewJellyfinRemoteLyricInfoDtoWithDefaults instantiates a new JellyfinRemoteLyricInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JellyfinRemoteLyricInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinRemoteLyricInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinRemoteLyricInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinRemoteLyricInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *JellyfinRemoteLyricInfoDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *JellyfinRemoteLyricInfoDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *JellyfinRemoteLyricInfoDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *JellyfinRemoteLyricInfoDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetLyrics

`func (o *JellyfinRemoteLyricInfoDto) GetLyrics() JellyfinJellyfinLyricDto`

GetLyrics returns the Lyrics field if non-nil, zero value otherwise.

### GetLyricsOk

`func (o *JellyfinRemoteLyricInfoDto) GetLyricsOk() (*JellyfinJellyfinLyricDto, bool)`

GetLyricsOk returns a tuple with the Lyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyrics

`func (o *JellyfinRemoteLyricInfoDto) SetLyrics(v JellyfinJellyfinLyricDto)`

SetLyrics sets Lyrics field to given value.

### HasLyrics

`func (o *JellyfinRemoteLyricInfoDto) HasLyrics() bool`

HasLyrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


