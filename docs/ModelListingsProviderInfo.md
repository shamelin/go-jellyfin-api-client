# ModelListingsProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**ListingsId** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**EnabledTuners** | Pointer to **[]string** |  | [optional] 
**EnableAllTuners** | Pointer to **bool** |  | [optional] 
**NewsCategories** | Pointer to **[]string** |  | [optional] 
**SportsCategories** | Pointer to **[]string** |  | [optional] 
**KidsCategories** | Pointer to **[]string** |  | [optional] 
**MovieCategories** | Pointer to **[]string** |  | [optional] 
**ChannelMappings** | Pointer to [**[]ModelModelNameValuePair**](ModelModelNameValuePair.md) |  | [optional] 
**MoviePrefix** | Pointer to **NullableString** |  | [optional] 
**PreferredLanguage** | Pointer to **NullableString** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelListingsProviderInfo

`func NewModelListingsProviderInfo() *ModelListingsProviderInfo`

NewModelListingsProviderInfo instantiates a new ModelListingsProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelListingsProviderInfoWithDefaults

`func NewModelListingsProviderInfoWithDefaults() *ModelListingsProviderInfo`

NewModelListingsProviderInfoWithDefaults instantiates a new ModelListingsProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelListingsProviderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelListingsProviderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelListingsProviderInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelListingsProviderInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelListingsProviderInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelListingsProviderInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *ModelListingsProviderInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelListingsProviderInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelListingsProviderInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelListingsProviderInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelListingsProviderInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelListingsProviderInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsername

`func (o *ModelListingsProviderInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelListingsProviderInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelListingsProviderInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelListingsProviderInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ModelListingsProviderInfo) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ModelListingsProviderInfo) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ModelListingsProviderInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelListingsProviderInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelListingsProviderInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelListingsProviderInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ModelListingsProviderInfo) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ModelListingsProviderInfo) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetListingsId

`func (o *ModelListingsProviderInfo) GetListingsId() string`

GetListingsId returns the ListingsId field if non-nil, zero value otherwise.

### GetListingsIdOk

`func (o *ModelListingsProviderInfo) GetListingsIdOk() (*string, bool)`

GetListingsIdOk returns a tuple with the ListingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsId

`func (o *ModelListingsProviderInfo) SetListingsId(v string)`

SetListingsId sets ListingsId field to given value.

### HasListingsId

`func (o *ModelListingsProviderInfo) HasListingsId() bool`

HasListingsId returns a boolean if a field has been set.

### SetListingsIdNil

`func (o *ModelListingsProviderInfo) SetListingsIdNil(b bool)`

 SetListingsIdNil sets the value for ListingsId to be an explicit nil

### UnsetListingsId
`func (o *ModelListingsProviderInfo) UnsetListingsId()`

UnsetListingsId ensures that no value is present for ListingsId, not even an explicit nil
### GetZipCode

`func (o *ModelListingsProviderInfo) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ModelListingsProviderInfo) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ModelListingsProviderInfo) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ModelListingsProviderInfo) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ModelListingsProviderInfo) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ModelListingsProviderInfo) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCountry

`func (o *ModelListingsProviderInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ModelListingsProviderInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ModelListingsProviderInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ModelListingsProviderInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ModelListingsProviderInfo) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ModelListingsProviderInfo) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPath

`func (o *ModelListingsProviderInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelListingsProviderInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelListingsProviderInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelListingsProviderInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ModelListingsProviderInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ModelListingsProviderInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEnabledTuners

`func (o *ModelListingsProviderInfo) GetEnabledTuners() []string`

GetEnabledTuners returns the EnabledTuners field if non-nil, zero value otherwise.

### GetEnabledTunersOk

`func (o *ModelListingsProviderInfo) GetEnabledTunersOk() (*[]string, bool)`

GetEnabledTunersOk returns a tuple with the EnabledTuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTuners

`func (o *ModelListingsProviderInfo) SetEnabledTuners(v []string)`

SetEnabledTuners sets EnabledTuners field to given value.

### HasEnabledTuners

`func (o *ModelListingsProviderInfo) HasEnabledTuners() bool`

HasEnabledTuners returns a boolean if a field has been set.

### SetEnabledTunersNil

`func (o *ModelListingsProviderInfo) SetEnabledTunersNil(b bool)`

 SetEnabledTunersNil sets the value for EnabledTuners to be an explicit nil

### UnsetEnabledTuners
`func (o *ModelListingsProviderInfo) UnsetEnabledTuners()`

UnsetEnabledTuners ensures that no value is present for EnabledTuners, not even an explicit nil
### GetEnableAllTuners

`func (o *ModelListingsProviderInfo) GetEnableAllTuners() bool`

GetEnableAllTuners returns the EnableAllTuners field if non-nil, zero value otherwise.

### GetEnableAllTunersOk

`func (o *ModelListingsProviderInfo) GetEnableAllTunersOk() (*bool, bool)`

GetEnableAllTunersOk returns a tuple with the EnableAllTuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllTuners

`func (o *ModelListingsProviderInfo) SetEnableAllTuners(v bool)`

SetEnableAllTuners sets EnableAllTuners field to given value.

### HasEnableAllTuners

`func (o *ModelListingsProviderInfo) HasEnableAllTuners() bool`

HasEnableAllTuners returns a boolean if a field has been set.

### GetNewsCategories

`func (o *ModelListingsProviderInfo) GetNewsCategories() []string`

GetNewsCategories returns the NewsCategories field if non-nil, zero value otherwise.

### GetNewsCategoriesOk

`func (o *ModelListingsProviderInfo) GetNewsCategoriesOk() (*[]string, bool)`

GetNewsCategoriesOk returns a tuple with the NewsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsCategories

`func (o *ModelListingsProviderInfo) SetNewsCategories(v []string)`

SetNewsCategories sets NewsCategories field to given value.

### HasNewsCategories

`func (o *ModelListingsProviderInfo) HasNewsCategories() bool`

HasNewsCategories returns a boolean if a field has been set.

### SetNewsCategoriesNil

`func (o *ModelListingsProviderInfo) SetNewsCategoriesNil(b bool)`

 SetNewsCategoriesNil sets the value for NewsCategories to be an explicit nil

### UnsetNewsCategories
`func (o *ModelListingsProviderInfo) UnsetNewsCategories()`

UnsetNewsCategories ensures that no value is present for NewsCategories, not even an explicit nil
### GetSportsCategories

`func (o *ModelListingsProviderInfo) GetSportsCategories() []string`

GetSportsCategories returns the SportsCategories field if non-nil, zero value otherwise.

### GetSportsCategoriesOk

`func (o *ModelListingsProviderInfo) GetSportsCategoriesOk() (*[]string, bool)`

GetSportsCategoriesOk returns a tuple with the SportsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportsCategories

`func (o *ModelListingsProviderInfo) SetSportsCategories(v []string)`

SetSportsCategories sets SportsCategories field to given value.

### HasSportsCategories

`func (o *ModelListingsProviderInfo) HasSportsCategories() bool`

HasSportsCategories returns a boolean if a field has been set.

### SetSportsCategoriesNil

`func (o *ModelListingsProviderInfo) SetSportsCategoriesNil(b bool)`

 SetSportsCategoriesNil sets the value for SportsCategories to be an explicit nil

### UnsetSportsCategories
`func (o *ModelListingsProviderInfo) UnsetSportsCategories()`

UnsetSportsCategories ensures that no value is present for SportsCategories, not even an explicit nil
### GetKidsCategories

`func (o *ModelListingsProviderInfo) GetKidsCategories() []string`

GetKidsCategories returns the KidsCategories field if non-nil, zero value otherwise.

### GetKidsCategoriesOk

`func (o *ModelListingsProviderInfo) GetKidsCategoriesOk() (*[]string, bool)`

GetKidsCategoriesOk returns a tuple with the KidsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKidsCategories

`func (o *ModelListingsProviderInfo) SetKidsCategories(v []string)`

SetKidsCategories sets KidsCategories field to given value.

### HasKidsCategories

`func (o *ModelListingsProviderInfo) HasKidsCategories() bool`

HasKidsCategories returns a boolean if a field has been set.

### SetKidsCategoriesNil

`func (o *ModelListingsProviderInfo) SetKidsCategoriesNil(b bool)`

 SetKidsCategoriesNil sets the value for KidsCategories to be an explicit nil

### UnsetKidsCategories
`func (o *ModelListingsProviderInfo) UnsetKidsCategories()`

UnsetKidsCategories ensures that no value is present for KidsCategories, not even an explicit nil
### GetMovieCategories

`func (o *ModelListingsProviderInfo) GetMovieCategories() []string`

GetMovieCategories returns the MovieCategories field if non-nil, zero value otherwise.

### GetMovieCategoriesOk

`func (o *ModelListingsProviderInfo) GetMovieCategoriesOk() (*[]string, bool)`

GetMovieCategoriesOk returns a tuple with the MovieCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieCategories

`func (o *ModelListingsProviderInfo) SetMovieCategories(v []string)`

SetMovieCategories sets MovieCategories field to given value.

### HasMovieCategories

`func (o *ModelListingsProviderInfo) HasMovieCategories() bool`

HasMovieCategories returns a boolean if a field has been set.

### SetMovieCategoriesNil

`func (o *ModelListingsProviderInfo) SetMovieCategoriesNil(b bool)`

 SetMovieCategoriesNil sets the value for MovieCategories to be an explicit nil

### UnsetMovieCategories
`func (o *ModelListingsProviderInfo) UnsetMovieCategories()`

UnsetMovieCategories ensures that no value is present for MovieCategories, not even an explicit nil
### GetChannelMappings

`func (o *ModelListingsProviderInfo) GetChannelMappings() []ModelModelNameValuePair`

GetChannelMappings returns the ChannelMappings field if non-nil, zero value otherwise.

### GetChannelMappingsOk

`func (o *ModelListingsProviderInfo) GetChannelMappingsOk() (*[]ModelModelNameValuePair, bool)`

GetChannelMappingsOk returns a tuple with the ChannelMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelMappings

`func (o *ModelListingsProviderInfo) SetChannelMappings(v []ModelModelNameValuePair)`

SetChannelMappings sets ChannelMappings field to given value.

### HasChannelMappings

`func (o *ModelListingsProviderInfo) HasChannelMappings() bool`

HasChannelMappings returns a boolean if a field has been set.

### SetChannelMappingsNil

`func (o *ModelListingsProviderInfo) SetChannelMappingsNil(b bool)`

 SetChannelMappingsNil sets the value for ChannelMappings to be an explicit nil

### UnsetChannelMappings
`func (o *ModelListingsProviderInfo) UnsetChannelMappings()`

UnsetChannelMappings ensures that no value is present for ChannelMappings, not even an explicit nil
### GetMoviePrefix

`func (o *ModelListingsProviderInfo) GetMoviePrefix() string`

GetMoviePrefix returns the MoviePrefix field if non-nil, zero value otherwise.

### GetMoviePrefixOk

`func (o *ModelListingsProviderInfo) GetMoviePrefixOk() (*string, bool)`

GetMoviePrefixOk returns a tuple with the MoviePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoviePrefix

`func (o *ModelListingsProviderInfo) SetMoviePrefix(v string)`

SetMoviePrefix sets MoviePrefix field to given value.

### HasMoviePrefix

`func (o *ModelListingsProviderInfo) HasMoviePrefix() bool`

HasMoviePrefix returns a boolean if a field has been set.

### SetMoviePrefixNil

`func (o *ModelListingsProviderInfo) SetMoviePrefixNil(b bool)`

 SetMoviePrefixNil sets the value for MoviePrefix to be an explicit nil

### UnsetMoviePrefix
`func (o *ModelListingsProviderInfo) UnsetMoviePrefix()`

UnsetMoviePrefix ensures that no value is present for MoviePrefix, not even an explicit nil
### GetPreferredLanguage

`func (o *ModelListingsProviderInfo) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *ModelListingsProviderInfo) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *ModelListingsProviderInfo) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *ModelListingsProviderInfo) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *ModelListingsProviderInfo) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *ModelListingsProviderInfo) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetUserAgent

`func (o *ModelListingsProviderInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ModelListingsProviderInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ModelListingsProviderInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ModelListingsProviderInfo) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ModelListingsProviderInfo) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ModelListingsProviderInfo) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


