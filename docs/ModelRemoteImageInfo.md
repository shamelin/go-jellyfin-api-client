# ModelRemoteImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | Pointer to **NullableString** | Gets or sets the name of the provider. | [optional] 
**Url** | Pointer to **NullableString** | Gets or sets the URL. | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | Gets or sets a url used for previewing a smaller version. | [optional] 
**Height** | Pointer to **NullableInt32** | Gets or sets the height. | [optional] 
**Width** | Pointer to **NullableInt32** | Gets or sets the width. | [optional] 
**CommunityRating** | Pointer to **NullableFloat64** | Gets or sets the community rating. | [optional] 
**VoteCount** | Pointer to **NullableInt32** | Gets or sets the vote count. | [optional] 
**Language** | Pointer to **NullableString** | Gets or sets the language. | [optional] 
**Type** | Pointer to [**ModelModelImageType**](ModelImageType.md) | Gets or sets the type. | [optional] 
**RatingType** | Pointer to [**ModelModelRatingType**](ModelRatingType.md) | Gets or sets the type of the rating. | [optional] 

## Methods

### NewModelRemoteImageInfo

`func NewModelRemoteImageInfo() *ModelRemoteImageInfo`

NewModelRemoteImageInfo instantiates a new ModelRemoteImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteImageInfoWithDefaults

`func NewModelRemoteImageInfoWithDefaults() *ModelRemoteImageInfo`

NewModelRemoteImageInfoWithDefaults instantiates a new ModelRemoteImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *ModelRemoteImageInfo) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ModelRemoteImageInfo) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ModelRemoteImageInfo) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ModelRemoteImageInfo) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ModelRemoteImageInfo) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ModelRemoteImageInfo) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetUrl

`func (o *ModelRemoteImageInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelRemoteImageInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelRemoteImageInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelRemoteImageInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ModelRemoteImageInfo) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ModelRemoteImageInfo) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetThumbnailUrl

`func (o *ModelRemoteImageInfo) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ModelRemoteImageInfo) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ModelRemoteImageInfo) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ModelRemoteImageInfo) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *ModelRemoteImageInfo) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *ModelRemoteImageInfo) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetHeight

`func (o *ModelRemoteImageInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelRemoteImageInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelRemoteImageInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelRemoteImageInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelRemoteImageInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelRemoteImageInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ModelRemoteImageInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelRemoteImageInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelRemoteImageInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelRemoteImageInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelRemoteImageInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelRemoteImageInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetCommunityRating

`func (o *ModelRemoteImageInfo) GetCommunityRating() float64`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *ModelRemoteImageInfo) GetCommunityRatingOk() (*float64, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *ModelRemoteImageInfo) SetCommunityRating(v float64)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *ModelRemoteImageInfo) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *ModelRemoteImageInfo) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *ModelRemoteImageInfo) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetVoteCount

`func (o *ModelRemoteImageInfo) GetVoteCount() int32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *ModelRemoteImageInfo) GetVoteCountOk() (*int32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *ModelRemoteImageInfo) SetVoteCount(v int32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *ModelRemoteImageInfo) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### SetVoteCountNil

`func (o *ModelRemoteImageInfo) SetVoteCountNil(b bool)`

 SetVoteCountNil sets the value for VoteCount to be an explicit nil

### UnsetVoteCount
`func (o *ModelRemoteImageInfo) UnsetVoteCount()`

UnsetVoteCount ensures that no value is present for VoteCount, not even an explicit nil
### GetLanguage

`func (o *ModelRemoteImageInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ModelRemoteImageInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ModelRemoteImageInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ModelRemoteImageInfo) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ModelRemoteImageInfo) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ModelRemoteImageInfo) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetType

`func (o *ModelRemoteImageInfo) GetType() ModelModelImageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelRemoteImageInfo) GetTypeOk() (*ModelModelImageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelRemoteImageInfo) SetType(v ModelModelImageType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelRemoteImageInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRatingType

`func (o *ModelRemoteImageInfo) GetRatingType() ModelModelRatingType`

GetRatingType returns the RatingType field if non-nil, zero value otherwise.

### GetRatingTypeOk

`func (o *ModelRemoteImageInfo) GetRatingTypeOk() (*ModelModelRatingType, bool)`

GetRatingTypeOk returns a tuple with the RatingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingType

`func (o *ModelRemoteImageInfo) SetRatingType(v ModelModelRatingType)`

SetRatingType sets RatingType field to given value.

### HasRatingType

`func (o *ModelRemoteImageInfo) HasRatingType() bool`

HasRatingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


