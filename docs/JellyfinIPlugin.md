# JellyfinIPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets the name of the plugin. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Gets the Description. | [optional] [readonly] 
**Id** | Pointer to **string** | Gets the unique id. | [optional] [readonly] 
**Version** | Pointer to **NullableString** | Gets the plugin version. | [optional] [readonly] 
**AssemblyFilePath** | Pointer to **NullableString** | Gets the path to the assembly file. | [optional] [readonly] 
**CanUninstall** | Pointer to **bool** | Gets a value indicating whether the plugin can be uninstalled. | [optional] [readonly] 
**DataFolderPath** | Pointer to **NullableString** | Gets the full path to the data folder, where the plugin can store any miscellaneous files needed. | [optional] [readonly] 

## Methods

### NewJellyfinIPlugin

`func NewJellyfinIPlugin() *JellyfinIPlugin`

NewJellyfinIPlugin instantiates a new JellyfinIPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinIPluginWithDefaults

`func NewJellyfinIPluginWithDefaults() *JellyfinIPlugin`

NewJellyfinIPluginWithDefaults instantiates a new JellyfinIPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinIPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinIPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinIPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinIPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinIPlugin) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinIPlugin) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *JellyfinIPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JellyfinIPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JellyfinIPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JellyfinIPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JellyfinIPlugin) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JellyfinIPlugin) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *JellyfinIPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinIPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinIPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinIPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *JellyfinIPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JellyfinIPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JellyfinIPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JellyfinIPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *JellyfinIPlugin) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *JellyfinIPlugin) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAssemblyFilePath

`func (o *JellyfinIPlugin) GetAssemblyFilePath() string`

GetAssemblyFilePath returns the AssemblyFilePath field if non-nil, zero value otherwise.

### GetAssemblyFilePathOk

`func (o *JellyfinIPlugin) GetAssemblyFilePathOk() (*string, bool)`

GetAssemblyFilePathOk returns a tuple with the AssemblyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyFilePath

`func (o *JellyfinIPlugin) SetAssemblyFilePath(v string)`

SetAssemblyFilePath sets AssemblyFilePath field to given value.

### HasAssemblyFilePath

`func (o *JellyfinIPlugin) HasAssemblyFilePath() bool`

HasAssemblyFilePath returns a boolean if a field has been set.

### SetAssemblyFilePathNil

`func (o *JellyfinIPlugin) SetAssemblyFilePathNil(b bool)`

 SetAssemblyFilePathNil sets the value for AssemblyFilePath to be an explicit nil

### UnsetAssemblyFilePath
`func (o *JellyfinIPlugin) UnsetAssemblyFilePath()`

UnsetAssemblyFilePath ensures that no value is present for AssemblyFilePath, not even an explicit nil
### GetCanUninstall

`func (o *JellyfinIPlugin) GetCanUninstall() bool`

GetCanUninstall returns the CanUninstall field if non-nil, zero value otherwise.

### GetCanUninstallOk

`func (o *JellyfinIPlugin) GetCanUninstallOk() (*bool, bool)`

GetCanUninstallOk returns a tuple with the CanUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUninstall

`func (o *JellyfinIPlugin) SetCanUninstall(v bool)`

SetCanUninstall sets CanUninstall field to given value.

### HasCanUninstall

`func (o *JellyfinIPlugin) HasCanUninstall() bool`

HasCanUninstall returns a boolean if a field has been set.

### GetDataFolderPath

`func (o *JellyfinIPlugin) GetDataFolderPath() string`

GetDataFolderPath returns the DataFolderPath field if non-nil, zero value otherwise.

### GetDataFolderPathOk

`func (o *JellyfinIPlugin) GetDataFolderPathOk() (*string, bool)`

GetDataFolderPathOk returns a tuple with the DataFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFolderPath

`func (o *JellyfinIPlugin) SetDataFolderPath(v string)`

SetDataFolderPath sets DataFolderPath field to given value.

### HasDataFolderPath

`func (o *JellyfinIPlugin) HasDataFolderPath() bool`

HasDataFolderPath returns a boolean if a field has been set.

### SetDataFolderPathNil

`func (o *JellyfinIPlugin) SetDataFolderPathNil(b bool)`

 SetDataFolderPathNil sets the value for DataFolderPath to be an explicit nil

### UnsetDataFolderPath
`func (o *JellyfinIPlugin) UnsetDataFolderPath()`

UnsetDataFolderPath ensures that no value is present for DataFolderPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


