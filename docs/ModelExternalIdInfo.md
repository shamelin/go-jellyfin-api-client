# ModelExternalIdInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the display name of the external id provider (IE: IMDB, MusicBrainz, etc). | [optional] 
**Key** | Pointer to **string** | Gets or sets the unique key for this id. This key should be unique across all providers. | [optional] 
**Type** | Pointer to [**NullableModelModelExternalIdMediaType**](ModelExternalIdMediaType.md) | Gets or sets the specific media type for this id. This is used to distinguish between the different  external id types for providers with multiple ids.  A null value indicates there is no specific media type associated with the external id, or this is the  default id for the external provider so there is no need to specify a type. | [optional] 
**UrlFormatString** | Pointer to **NullableString** | Gets or sets the URL format string. | [optional] 

## Methods

### NewModelExternalIdInfo

`func NewModelExternalIdInfo() *ModelExternalIdInfo`

NewModelExternalIdInfo instantiates a new ModelExternalIdInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelExternalIdInfoWithDefaults

`func NewModelExternalIdInfoWithDefaults() *ModelExternalIdInfo`

NewModelExternalIdInfoWithDefaults instantiates a new ModelExternalIdInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelExternalIdInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelExternalIdInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelExternalIdInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelExternalIdInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *ModelExternalIdInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelExternalIdInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelExternalIdInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelExternalIdInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *ModelExternalIdInfo) GetType() ModelModelExternalIdMediaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelExternalIdInfo) GetTypeOk() (*ModelModelExternalIdMediaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelExternalIdInfo) SetType(v ModelModelExternalIdMediaType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelExternalIdInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelExternalIdInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelExternalIdInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrlFormatString

`func (o *ModelExternalIdInfo) GetUrlFormatString() string`

GetUrlFormatString returns the UrlFormatString field if non-nil, zero value otherwise.

### GetUrlFormatStringOk

`func (o *ModelExternalIdInfo) GetUrlFormatStringOk() (*string, bool)`

GetUrlFormatStringOk returns a tuple with the UrlFormatString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFormatString

`func (o *ModelExternalIdInfo) SetUrlFormatString(v string)`

SetUrlFormatString sets UrlFormatString field to given value.

### HasUrlFormatString

`func (o *ModelExternalIdInfo) HasUrlFormatString() bool`

HasUrlFormatString returns a boolean if a field has been set.

### SetUrlFormatStringNil

`func (o *ModelExternalIdInfo) SetUrlFormatStringNil(b bool)`

 SetUrlFormatStringNil sets the value for UrlFormatString to be an explicit nil

### UnsetUrlFormatString
`func (o *ModelExternalIdInfo) UnsetUrlFormatString()`

UnsetUrlFormatString ensures that no value is present for UrlFormatString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


