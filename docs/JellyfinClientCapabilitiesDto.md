# JellyfinClientCapabilitiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to [**[]JellyfinJellyfinMediaType**](JellyfinJellyfinMediaType.md) | Gets or sets the list of playable media types. | [optional] 
**SupportedCommands** | Pointer to [**[]JellyfinJellyfinGeneralCommandType**](JellyfinJellyfinGeneralCommandType.md) | Gets or sets the list of supported commands. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether session supports media control. | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** | Gets or sets a value indicating whether session supports a persistent identifier. | [optional] 
**DeviceProfile** | Pointer to [**NullableJellyfinJellyfinDeviceProfile**](JellyfinDeviceProfile.md) | Gets or sets the device profile. | [optional] 
**AppStoreUrl** | Pointer to **NullableString** | Gets or sets the app store url. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon url. | [optional] 

## Methods

### NewJellyfinClientCapabilitiesDto

`func NewJellyfinClientCapabilitiesDto() *JellyfinClientCapabilitiesDto`

NewJellyfinClientCapabilitiesDto instantiates a new JellyfinClientCapabilitiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinClientCapabilitiesDtoWithDefaults

`func NewJellyfinClientCapabilitiesDtoWithDefaults() *JellyfinClientCapabilitiesDto`

NewJellyfinClientCapabilitiesDtoWithDefaults instantiates a new JellyfinClientCapabilitiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *JellyfinClientCapabilitiesDto) GetPlayableMediaTypes() []JellyfinJellyfinMediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *JellyfinClientCapabilitiesDto) GetPlayableMediaTypesOk() (*[]JellyfinJellyfinMediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *JellyfinClientCapabilitiesDto) SetPlayableMediaTypes(v []JellyfinJellyfinMediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *JellyfinClientCapabilitiesDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *JellyfinClientCapabilitiesDto) GetSupportedCommands() []JellyfinJellyfinGeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *JellyfinClientCapabilitiesDto) GetSupportedCommandsOk() (*[]JellyfinJellyfinGeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *JellyfinClientCapabilitiesDto) SetSupportedCommands(v []JellyfinJellyfinGeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *JellyfinClientCapabilitiesDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *JellyfinClientCapabilitiesDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *JellyfinClientCapabilitiesDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *JellyfinClientCapabilitiesDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *JellyfinClientCapabilitiesDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsPersistentIdentifier

`func (o *JellyfinClientCapabilitiesDto) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *JellyfinClientCapabilitiesDto) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *JellyfinClientCapabilitiesDto) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *JellyfinClientCapabilitiesDto) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *JellyfinClientCapabilitiesDto) GetDeviceProfile() JellyfinJellyfinDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *JellyfinClientCapabilitiesDto) GetDeviceProfileOk() (*JellyfinJellyfinDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *JellyfinClientCapabilitiesDto) SetDeviceProfile(v JellyfinJellyfinDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *JellyfinClientCapabilitiesDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *JellyfinClientCapabilitiesDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *JellyfinClientCapabilitiesDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetAppStoreUrl

`func (o *JellyfinClientCapabilitiesDto) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *JellyfinClientCapabilitiesDto) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *JellyfinClientCapabilitiesDto) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *JellyfinClientCapabilitiesDto) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrlNil

`func (o *JellyfinClientCapabilitiesDto) SetAppStoreUrlNil(b bool)`

 SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil

### UnsetAppStoreUrl
`func (o *JellyfinClientCapabilitiesDto) UnsetAppStoreUrl()`

UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
### GetIconUrl

`func (o *JellyfinClientCapabilitiesDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *JellyfinClientCapabilitiesDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *JellyfinClientCapabilitiesDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *JellyfinClientCapabilitiesDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *JellyfinClientCapabilitiesDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *JellyfinClientCapabilitiesDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


