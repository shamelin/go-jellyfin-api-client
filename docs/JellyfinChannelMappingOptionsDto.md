# JellyfinChannelMappingOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunerChannels** | Pointer to [**[]JellyfinJellyfinTunerChannelMapping**](JellyfinJellyfinTunerChannelMapping.md) | Gets or sets list of tuner channels. | [optional] 
**ProviderChannels** | Pointer to [**[]JellyfinJellyfinNameIdPair**](JellyfinJellyfinNameIdPair.md) | Gets or sets list of provider channels. | [optional] 
**Mappings** | Pointer to [**[]JellyfinJellyfinNameValuePair**](JellyfinJellyfinNameValuePair.md) | Gets or sets list of mappings. | [optional] 
**ProviderName** | Pointer to **NullableString** | Gets or sets provider name. | [optional] 

## Methods

### NewJellyfinChannelMappingOptionsDto

`func NewJellyfinChannelMappingOptionsDto() *JellyfinChannelMappingOptionsDto`

NewJellyfinChannelMappingOptionsDto instantiates a new JellyfinChannelMappingOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinChannelMappingOptionsDtoWithDefaults

`func NewJellyfinChannelMappingOptionsDtoWithDefaults() *JellyfinChannelMappingOptionsDto`

NewJellyfinChannelMappingOptionsDtoWithDefaults instantiates a new JellyfinChannelMappingOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunerChannels

`func (o *JellyfinChannelMappingOptionsDto) GetTunerChannels() []JellyfinJellyfinTunerChannelMapping`

GetTunerChannels returns the TunerChannels field if non-nil, zero value otherwise.

### GetTunerChannelsOk

`func (o *JellyfinChannelMappingOptionsDto) GetTunerChannelsOk() (*[]JellyfinJellyfinTunerChannelMapping, bool)`

GetTunerChannelsOk returns a tuple with the TunerChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerChannels

`func (o *JellyfinChannelMappingOptionsDto) SetTunerChannels(v []JellyfinJellyfinTunerChannelMapping)`

SetTunerChannels sets TunerChannels field to given value.

### HasTunerChannels

`func (o *JellyfinChannelMappingOptionsDto) HasTunerChannels() bool`

HasTunerChannels returns a boolean if a field has been set.

### GetProviderChannels

`func (o *JellyfinChannelMappingOptionsDto) GetProviderChannels() []JellyfinJellyfinNameIdPair`

GetProviderChannels returns the ProviderChannels field if non-nil, zero value otherwise.

### GetProviderChannelsOk

`func (o *JellyfinChannelMappingOptionsDto) GetProviderChannelsOk() (*[]JellyfinJellyfinNameIdPair, bool)`

GetProviderChannelsOk returns a tuple with the ProviderChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderChannels

`func (o *JellyfinChannelMappingOptionsDto) SetProviderChannels(v []JellyfinJellyfinNameIdPair)`

SetProviderChannels sets ProviderChannels field to given value.

### HasProviderChannels

`func (o *JellyfinChannelMappingOptionsDto) HasProviderChannels() bool`

HasProviderChannels returns a boolean if a field has been set.

### GetMappings

`func (o *JellyfinChannelMappingOptionsDto) GetMappings() []JellyfinJellyfinNameValuePair`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *JellyfinChannelMappingOptionsDto) GetMappingsOk() (*[]JellyfinJellyfinNameValuePair, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *JellyfinChannelMappingOptionsDto) SetMappings(v []JellyfinJellyfinNameValuePair)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *JellyfinChannelMappingOptionsDto) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetProviderName

`func (o *JellyfinChannelMappingOptionsDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *JellyfinChannelMappingOptionsDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *JellyfinChannelMappingOptionsDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *JellyfinChannelMappingOptionsDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *JellyfinChannelMappingOptionsDto) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *JellyfinChannelMappingOptionsDto) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


