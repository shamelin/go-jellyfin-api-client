# JellyfinMetadataEditorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentalRatingOptions** | Pointer to [**[]JellyfinJellyfinParentalRating**](JellyfinJellyfinParentalRating.md) |  | [optional] 
**Countries** | Pointer to [**[]JellyfinJellyfinCountryInfo**](JellyfinJellyfinCountryInfo.md) |  | [optional] 
**Cultures** | Pointer to [**[]JellyfinJellyfinCultureDto**](JellyfinJellyfinCultureDto.md) |  | [optional] 
**ExternalIdInfos** | Pointer to [**[]JellyfinJellyfinExternalIdInfo**](JellyfinJellyfinExternalIdInfo.md) |  | [optional] 
**ContentType** | Pointer to [**NullableJellyfinJellyfinCollectionType**](JellyfinCollectionType.md) |  | [optional] 
**ContentTypeOptions** | Pointer to [**[]JellyfinJellyfinNameValuePair**](JellyfinJellyfinNameValuePair.md) |  | [optional] 

## Methods

### NewJellyfinMetadataEditorInfo

`func NewJellyfinMetadataEditorInfo() *JellyfinMetadataEditorInfo`

NewJellyfinMetadataEditorInfo instantiates a new JellyfinMetadataEditorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinMetadataEditorInfoWithDefaults

`func NewJellyfinMetadataEditorInfoWithDefaults() *JellyfinMetadataEditorInfo`

NewJellyfinMetadataEditorInfoWithDefaults instantiates a new JellyfinMetadataEditorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentalRatingOptions

`func (o *JellyfinMetadataEditorInfo) GetParentalRatingOptions() []JellyfinJellyfinParentalRating`

GetParentalRatingOptions returns the ParentalRatingOptions field if non-nil, zero value otherwise.

### GetParentalRatingOptionsOk

`func (o *JellyfinMetadataEditorInfo) GetParentalRatingOptionsOk() (*[]JellyfinJellyfinParentalRating, bool)`

GetParentalRatingOptionsOk returns a tuple with the ParentalRatingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentalRatingOptions

`func (o *JellyfinMetadataEditorInfo) SetParentalRatingOptions(v []JellyfinJellyfinParentalRating)`

SetParentalRatingOptions sets ParentalRatingOptions field to given value.

### HasParentalRatingOptions

`func (o *JellyfinMetadataEditorInfo) HasParentalRatingOptions() bool`

HasParentalRatingOptions returns a boolean if a field has been set.

### GetCountries

`func (o *JellyfinMetadataEditorInfo) GetCountries() []JellyfinJellyfinCountryInfo`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *JellyfinMetadataEditorInfo) GetCountriesOk() (*[]JellyfinJellyfinCountryInfo, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *JellyfinMetadataEditorInfo) SetCountries(v []JellyfinJellyfinCountryInfo)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *JellyfinMetadataEditorInfo) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCultures

`func (o *JellyfinMetadataEditorInfo) GetCultures() []JellyfinJellyfinCultureDto`

GetCultures returns the Cultures field if non-nil, zero value otherwise.

### GetCulturesOk

`func (o *JellyfinMetadataEditorInfo) GetCulturesOk() (*[]JellyfinJellyfinCultureDto, bool)`

GetCulturesOk returns a tuple with the Cultures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultures

`func (o *JellyfinMetadataEditorInfo) SetCultures(v []JellyfinJellyfinCultureDto)`

SetCultures sets Cultures field to given value.

### HasCultures

`func (o *JellyfinMetadataEditorInfo) HasCultures() bool`

HasCultures returns a boolean if a field has been set.

### GetExternalIdInfos

`func (o *JellyfinMetadataEditorInfo) GetExternalIdInfos() []JellyfinJellyfinExternalIdInfo`

GetExternalIdInfos returns the ExternalIdInfos field if non-nil, zero value otherwise.

### GetExternalIdInfosOk

`func (o *JellyfinMetadataEditorInfo) GetExternalIdInfosOk() (*[]JellyfinJellyfinExternalIdInfo, bool)`

GetExternalIdInfosOk returns a tuple with the ExternalIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdInfos

`func (o *JellyfinMetadataEditorInfo) SetExternalIdInfos(v []JellyfinJellyfinExternalIdInfo)`

SetExternalIdInfos sets ExternalIdInfos field to given value.

### HasExternalIdInfos

`func (o *JellyfinMetadataEditorInfo) HasExternalIdInfos() bool`

HasExternalIdInfos returns a boolean if a field has been set.

### GetContentType

`func (o *JellyfinMetadataEditorInfo) GetContentType() JellyfinJellyfinCollectionType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *JellyfinMetadataEditorInfo) GetContentTypeOk() (*JellyfinJellyfinCollectionType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *JellyfinMetadataEditorInfo) SetContentType(v JellyfinJellyfinCollectionType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *JellyfinMetadataEditorInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *JellyfinMetadataEditorInfo) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *JellyfinMetadataEditorInfo) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentTypeOptions

`func (o *JellyfinMetadataEditorInfo) GetContentTypeOptions() []JellyfinJellyfinNameValuePair`

GetContentTypeOptions returns the ContentTypeOptions field if non-nil, zero value otherwise.

### GetContentTypeOptionsOk

`func (o *JellyfinMetadataEditorInfo) GetContentTypeOptionsOk() (*[]JellyfinJellyfinNameValuePair, bool)`

GetContentTypeOptionsOk returns a tuple with the ContentTypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeOptions

`func (o *JellyfinMetadataEditorInfo) SetContentTypeOptions(v []JellyfinJellyfinNameValuePair)`

SetContentTypeOptions sets ContentTypeOptions field to given value.

### HasContentTypeOptions

`func (o *JellyfinMetadataEditorInfo) HasContentTypeOptions() bool`

HasContentTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


