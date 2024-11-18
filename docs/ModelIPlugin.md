# ModelIPlugin

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

### NewModelIPlugin

`func NewModelIPlugin() *ModelIPlugin`

NewModelIPlugin instantiates a new ModelIPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIPluginWithDefaults

`func NewModelIPluginWithDefaults() *ModelIPlugin`

NewModelIPluginWithDefaults instantiates a new ModelIPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelIPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelIPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelIPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelIPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelIPlugin) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelIPlugin) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ModelIPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelIPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelIPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelIPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelIPlugin) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelIPlugin) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ModelIPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelIPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelIPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelIPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *ModelIPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelIPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelIPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelIPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelIPlugin) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelIPlugin) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAssemblyFilePath

`func (o *ModelIPlugin) GetAssemblyFilePath() string`

GetAssemblyFilePath returns the AssemblyFilePath field if non-nil, zero value otherwise.

### GetAssemblyFilePathOk

`func (o *ModelIPlugin) GetAssemblyFilePathOk() (*string, bool)`

GetAssemblyFilePathOk returns a tuple with the AssemblyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyFilePath

`func (o *ModelIPlugin) SetAssemblyFilePath(v string)`

SetAssemblyFilePath sets AssemblyFilePath field to given value.

### HasAssemblyFilePath

`func (o *ModelIPlugin) HasAssemblyFilePath() bool`

HasAssemblyFilePath returns a boolean if a field has been set.

### SetAssemblyFilePathNil

`func (o *ModelIPlugin) SetAssemblyFilePathNil(b bool)`

 SetAssemblyFilePathNil sets the value for AssemblyFilePath to be an explicit nil

### UnsetAssemblyFilePath
`func (o *ModelIPlugin) UnsetAssemblyFilePath()`

UnsetAssemblyFilePath ensures that no value is present for AssemblyFilePath, not even an explicit nil
### GetCanUninstall

`func (o *ModelIPlugin) GetCanUninstall() bool`

GetCanUninstall returns the CanUninstall field if non-nil, zero value otherwise.

### GetCanUninstallOk

`func (o *ModelIPlugin) GetCanUninstallOk() (*bool, bool)`

GetCanUninstallOk returns a tuple with the CanUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUninstall

`func (o *ModelIPlugin) SetCanUninstall(v bool)`

SetCanUninstall sets CanUninstall field to given value.

### HasCanUninstall

`func (o *ModelIPlugin) HasCanUninstall() bool`

HasCanUninstall returns a boolean if a field has been set.

### GetDataFolderPath

`func (o *ModelIPlugin) GetDataFolderPath() string`

GetDataFolderPath returns the DataFolderPath field if non-nil, zero value otherwise.

### GetDataFolderPathOk

`func (o *ModelIPlugin) GetDataFolderPathOk() (*string, bool)`

GetDataFolderPathOk returns a tuple with the DataFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFolderPath

`func (o *ModelIPlugin) SetDataFolderPath(v string)`

SetDataFolderPath sets DataFolderPath field to given value.

### HasDataFolderPath

`func (o *ModelIPlugin) HasDataFolderPath() bool`

HasDataFolderPath returns a boolean if a field has been set.

### SetDataFolderPathNil

`func (o *ModelIPlugin) SetDataFolderPathNil(b bool)`

 SetDataFolderPathNil sets the value for DataFolderPath to be an explicit nil

### UnsetDataFolderPath
`func (o *ModelIPlugin) UnsetDataFolderPath()`

UnsetDataFolderPath ensures that no value is present for DataFolderPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


