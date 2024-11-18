# ModelClientCapabilitiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to [**[]ModelModelMediaType**](ModelModelMediaType.md) | Gets or sets the list of playable media types. | [optional] 
**SupportedCommands** | Pointer to [**[]ModelModelGeneralCommandType**](ModelModelGeneralCommandType.md) | Gets or sets the list of supported commands. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether session supports media control. | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** | Gets or sets a value indicating whether session supports a persistent identifier. | [optional] 
**DeviceProfile** | Pointer to [**NullableModelModelDeviceProfile**](ModelDeviceProfile.md) | Gets or sets the device profile. | [optional] 
**AppStoreUrl** | Pointer to **NullableString** | Gets or sets the app store url. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon url. | [optional] 

## Methods

### NewModelClientCapabilitiesDto

`func NewModelClientCapabilitiesDto() *ModelClientCapabilitiesDto`

NewModelClientCapabilitiesDto instantiates a new ModelClientCapabilitiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelClientCapabilitiesDtoWithDefaults

`func NewModelClientCapabilitiesDtoWithDefaults() *ModelClientCapabilitiesDto`

NewModelClientCapabilitiesDtoWithDefaults instantiates a new ModelClientCapabilitiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *ModelClientCapabilitiesDto) GetPlayableMediaTypes() []ModelModelMediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ModelClientCapabilitiesDto) GetPlayableMediaTypesOk() (*[]ModelModelMediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ModelClientCapabilitiesDto) SetPlayableMediaTypes(v []ModelModelMediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ModelClientCapabilitiesDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *ModelClientCapabilitiesDto) GetSupportedCommands() []ModelModelGeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ModelClientCapabilitiesDto) GetSupportedCommandsOk() (*[]ModelModelGeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ModelClientCapabilitiesDto) SetSupportedCommands(v []ModelModelGeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ModelClientCapabilitiesDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *ModelClientCapabilitiesDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ModelClientCapabilitiesDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ModelClientCapabilitiesDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ModelClientCapabilitiesDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsPersistentIdentifier

`func (o *ModelClientCapabilitiesDto) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *ModelClientCapabilitiesDto) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *ModelClientCapabilitiesDto) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *ModelClientCapabilitiesDto) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ModelClientCapabilitiesDto) GetDeviceProfile() ModelModelDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ModelClientCapabilitiesDto) GetDeviceProfileOk() (*ModelModelDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ModelClientCapabilitiesDto) SetDeviceProfile(v ModelModelDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ModelClientCapabilitiesDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *ModelClientCapabilitiesDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *ModelClientCapabilitiesDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetAppStoreUrl

`func (o *ModelClientCapabilitiesDto) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *ModelClientCapabilitiesDto) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *ModelClientCapabilitiesDto) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *ModelClientCapabilitiesDto) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrlNil

`func (o *ModelClientCapabilitiesDto) SetAppStoreUrlNil(b bool)`

 SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil

### UnsetAppStoreUrl
`func (o *ModelClientCapabilitiesDto) UnsetAppStoreUrl()`

UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
### GetIconUrl

`func (o *ModelClientCapabilitiesDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ModelClientCapabilitiesDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ModelClientCapabilitiesDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ModelClientCapabilitiesDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ModelClientCapabilitiesDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ModelClientCapabilitiesDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


