# ModelChannelMappingOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunerChannels** | Pointer to [**[]ModelModelTunerChannelMapping**](ModelModelTunerChannelMapping.md) | Gets or sets list of tuner channels. | [optional] 
**ProviderChannels** | Pointer to [**[]ModelModelNameIdPair**](ModelModelNameIdPair.md) | Gets or sets list of provider channels. | [optional] 
**Mappings** | Pointer to [**[]ModelModelNameValuePair**](ModelModelNameValuePair.md) | Gets or sets list of mappings. | [optional] 
**ProviderName** | Pointer to **NullableString** | Gets or sets provider name. | [optional] 

## Methods

### NewModelChannelMappingOptionsDto

`func NewModelChannelMappingOptionsDto() *ModelChannelMappingOptionsDto`

NewModelChannelMappingOptionsDto instantiates a new ModelChannelMappingOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelChannelMappingOptionsDtoWithDefaults

`func NewModelChannelMappingOptionsDtoWithDefaults() *ModelChannelMappingOptionsDto`

NewModelChannelMappingOptionsDtoWithDefaults instantiates a new ModelChannelMappingOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunerChannels

`func (o *ModelChannelMappingOptionsDto) GetTunerChannels() []ModelModelTunerChannelMapping`

GetTunerChannels returns the TunerChannels field if non-nil, zero value otherwise.

### GetTunerChannelsOk

`func (o *ModelChannelMappingOptionsDto) GetTunerChannelsOk() (*[]ModelModelTunerChannelMapping, bool)`

GetTunerChannelsOk returns a tuple with the TunerChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerChannels

`func (o *ModelChannelMappingOptionsDto) SetTunerChannels(v []ModelModelTunerChannelMapping)`

SetTunerChannels sets TunerChannels field to given value.

### HasTunerChannels

`func (o *ModelChannelMappingOptionsDto) HasTunerChannels() bool`

HasTunerChannels returns a boolean if a field has been set.

### GetProviderChannels

`func (o *ModelChannelMappingOptionsDto) GetProviderChannels() []ModelModelNameIdPair`

GetProviderChannels returns the ProviderChannels field if non-nil, zero value otherwise.

### GetProviderChannelsOk

`func (o *ModelChannelMappingOptionsDto) GetProviderChannelsOk() (*[]ModelModelNameIdPair, bool)`

GetProviderChannelsOk returns a tuple with the ProviderChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderChannels

`func (o *ModelChannelMappingOptionsDto) SetProviderChannels(v []ModelModelNameIdPair)`

SetProviderChannels sets ProviderChannels field to given value.

### HasProviderChannels

`func (o *ModelChannelMappingOptionsDto) HasProviderChannels() bool`

HasProviderChannels returns a boolean if a field has been set.

### GetMappings

`func (o *ModelChannelMappingOptionsDto) GetMappings() []ModelModelNameValuePair`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ModelChannelMappingOptionsDto) GetMappingsOk() (*[]ModelModelNameValuePair, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ModelChannelMappingOptionsDto) SetMappings(v []ModelModelNameValuePair)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ModelChannelMappingOptionsDto) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetProviderName

`func (o *ModelChannelMappingOptionsDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ModelChannelMappingOptionsDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ModelChannelMappingOptionsDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ModelChannelMappingOptionsDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ModelChannelMappingOptionsDto) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ModelChannelMappingOptionsDto) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


