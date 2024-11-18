# JellyfinPluginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Version** | Pointer to **string** | Gets or sets the version. | [optional] 
**ConfigurationFileName** | Pointer to **NullableString** | Gets or sets the name of the configuration file. | [optional] 
**Description** | Pointer to **string** | Gets or sets the description. | [optional] 
**Id** | Pointer to **string** | Gets or sets the unique id. | [optional] 
**CanUninstall** | Pointer to **bool** | Gets or sets a value indicating whether the plugin can be uninstalled. | [optional] 
**HasImage** | Pointer to **bool** | Gets or sets a value indicating whether this plugin has a valid image. | [optional] 
**Status** | Pointer to [**JellyfinJellyfinPluginStatus**](JellyfinPluginStatus.md) | Gets or sets a value indicating the status of the plugin. | [optional] 

## Methods

### NewJellyfinPluginInfo

`func NewJellyfinPluginInfo() *JellyfinPluginInfo`

NewJellyfinPluginInfo instantiates a new JellyfinPluginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinPluginInfoWithDefaults

`func NewJellyfinPluginInfoWithDefaults() *JellyfinPluginInfo`

NewJellyfinPluginInfoWithDefaults instantiates a new JellyfinPluginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinPluginInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinPluginInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinPluginInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinPluginInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *JellyfinPluginInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JellyfinPluginInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JellyfinPluginInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JellyfinPluginInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConfigurationFileName

`func (o *JellyfinPluginInfo) GetConfigurationFileName() string`

GetConfigurationFileName returns the ConfigurationFileName field if non-nil, zero value otherwise.

### GetConfigurationFileNameOk

`func (o *JellyfinPluginInfo) GetConfigurationFileNameOk() (*string, bool)`

GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFileName

`func (o *JellyfinPluginInfo) SetConfigurationFileName(v string)`

SetConfigurationFileName sets ConfigurationFileName field to given value.

### HasConfigurationFileName

`func (o *JellyfinPluginInfo) HasConfigurationFileName() bool`

HasConfigurationFileName returns a boolean if a field has been set.

### SetConfigurationFileNameNil

`func (o *JellyfinPluginInfo) SetConfigurationFileNameNil(b bool)`

 SetConfigurationFileNameNil sets the value for ConfigurationFileName to be an explicit nil

### UnsetConfigurationFileName
`func (o *JellyfinPluginInfo) UnsetConfigurationFileName()`

UnsetConfigurationFileName ensures that no value is present for ConfigurationFileName, not even an explicit nil
### GetDescription

`func (o *JellyfinPluginInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JellyfinPluginInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JellyfinPluginInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JellyfinPluginInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *JellyfinPluginInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinPluginInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinPluginInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinPluginInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCanUninstall

`func (o *JellyfinPluginInfo) GetCanUninstall() bool`

GetCanUninstall returns the CanUninstall field if non-nil, zero value otherwise.

### GetCanUninstallOk

`func (o *JellyfinPluginInfo) GetCanUninstallOk() (*bool, bool)`

GetCanUninstallOk returns a tuple with the CanUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUninstall

`func (o *JellyfinPluginInfo) SetCanUninstall(v bool)`

SetCanUninstall sets CanUninstall field to given value.

### HasCanUninstall

`func (o *JellyfinPluginInfo) HasCanUninstall() bool`

HasCanUninstall returns a boolean if a field has been set.

### GetHasImage

`func (o *JellyfinPluginInfo) GetHasImage() bool`

GetHasImage returns the HasImage field if non-nil, zero value otherwise.

### GetHasImageOk

`func (o *JellyfinPluginInfo) GetHasImageOk() (*bool, bool)`

GetHasImageOk returns a tuple with the HasImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImage

`func (o *JellyfinPluginInfo) SetHasImage(v bool)`

SetHasImage sets HasImage field to given value.

### HasHasImage

`func (o *JellyfinPluginInfo) HasHasImage() bool`

HasHasImage returns a boolean if a field has been set.

### GetStatus

`func (o *JellyfinPluginInfo) GetStatus() JellyfinJellyfinPluginStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JellyfinPluginInfo) GetStatusOk() (*JellyfinJellyfinPluginStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JellyfinPluginInfo) SetStatus(v JellyfinJellyfinPluginStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JellyfinPluginInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


