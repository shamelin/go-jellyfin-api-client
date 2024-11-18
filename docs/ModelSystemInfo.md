# ModelSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalAddress** | Pointer to **NullableString** | Gets or sets the local address. | [optional] 
**ServerName** | Pointer to **NullableString** | Gets or sets the name of the server. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the server version. | [optional] 
**ProductName** | Pointer to **NullableString** | Gets or sets the product name. This is the AssemblyProduct name. | [optional] 
**OperatingSystem** | Pointer to **NullableString** | Gets or sets the operating system. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**StartupWizardCompleted** | Pointer to **NullableBool** | Gets or sets a value indicating whether the startup wizard is completed. | [optional] 
**OperatingSystemDisplayName** | Pointer to **NullableString** | Gets or sets the display name of the operating system. | [optional] 
**PackageName** | Pointer to **NullableString** | Gets or sets the package name. | [optional] 
**HasPendingRestart** | Pointer to **bool** | Gets or sets a value indicating whether this instance has pending restart. | [optional] 
**IsShuttingDown** | Pointer to **bool** |  | [optional] 
**SupportsLibraryMonitor** | Pointer to **bool** | Gets or sets a value indicating whether [supports library monitor]. | [optional] 
**WebSocketPortNumber** | Pointer to **int32** | Gets or sets the web socket port number. | [optional] 
**CompletedInstallations** | Pointer to [**[]ModelModelInstallationInfo**](ModelModelInstallationInfo.md) | Gets or sets the completed installations. | [optional] 
**CanSelfRestart** | Pointer to **bool** | Gets or sets a value indicating whether this instance can self restart. | [optional] [default to true]
**CanLaunchWebBrowser** | Pointer to **bool** |  | [optional] [default to false]
**ProgramDataPath** | Pointer to **NullableString** | Gets or sets the program data path. | [optional] 
**WebPath** | Pointer to **NullableString** | Gets or sets the web UI resources path. | [optional] 
**ItemsByNamePath** | Pointer to **NullableString** | Gets or sets the items by name path. | [optional] 
**CachePath** | Pointer to **NullableString** | Gets or sets the cache path. | [optional] 
**LogPath** | Pointer to **NullableString** | Gets or sets the log path. | [optional] 
**InternalMetadataPath** | Pointer to **NullableString** | Gets or sets the internal metadata path. | [optional] 
**TranscodingTempPath** | Pointer to **NullableString** | Gets or sets the transcode path. | [optional] 
**CastReceiverApplications** | Pointer to [**[]ModelModelCastReceiverApplication**](ModelModelCastReceiverApplication.md) | Gets or sets the list of cast receiver applications. | [optional] 
**HasUpdateAvailable** | Pointer to **bool** | Gets or sets a value indicating whether this instance has update available. | [optional] [default to false]
**EncoderLocation** | Pointer to **NullableString** |  | [optional] [default to "System"]
**SystemArchitecture** | Pointer to **NullableString** |  | [optional] [default to "X64"]

## Methods

### NewModelSystemInfo

`func NewModelSystemInfo() *ModelSystemInfo`

NewModelSystemInfo instantiates a new ModelSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSystemInfoWithDefaults

`func NewModelSystemInfoWithDefaults() *ModelSystemInfo`

NewModelSystemInfoWithDefaults instantiates a new ModelSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalAddress

`func (o *ModelSystemInfo) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *ModelSystemInfo) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *ModelSystemInfo) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *ModelSystemInfo) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### SetLocalAddressNil

`func (o *ModelSystemInfo) SetLocalAddressNil(b bool)`

 SetLocalAddressNil sets the value for LocalAddress to be an explicit nil

### UnsetLocalAddress
`func (o *ModelSystemInfo) UnsetLocalAddress()`

UnsetLocalAddress ensures that no value is present for LocalAddress, not even an explicit nil
### GetServerName

`func (o *ModelSystemInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ModelSystemInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ModelSystemInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ModelSystemInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *ModelSystemInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *ModelSystemInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetVersion

`func (o *ModelSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelSystemInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelSystemInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetProductName

`func (o *ModelSystemInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ModelSystemInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ModelSystemInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ModelSystemInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *ModelSystemInfo) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *ModelSystemInfo) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetOperatingSystem

`func (o *ModelSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ModelSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ModelSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ModelSystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *ModelSystemInfo) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *ModelSystemInfo) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetId

`func (o *ModelSystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSystemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSystemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelSystemInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelSystemInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartupWizardCompleted

`func (o *ModelSystemInfo) GetStartupWizardCompleted() bool`

GetStartupWizardCompleted returns the StartupWizardCompleted field if non-nil, zero value otherwise.

### GetStartupWizardCompletedOk

`func (o *ModelSystemInfo) GetStartupWizardCompletedOk() (*bool, bool)`

GetStartupWizardCompletedOk returns a tuple with the StartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupWizardCompleted

`func (o *ModelSystemInfo) SetStartupWizardCompleted(v bool)`

SetStartupWizardCompleted sets StartupWizardCompleted field to given value.

### HasStartupWizardCompleted

`func (o *ModelSystemInfo) HasStartupWizardCompleted() bool`

HasStartupWizardCompleted returns a boolean if a field has been set.

### SetStartupWizardCompletedNil

`func (o *ModelSystemInfo) SetStartupWizardCompletedNil(b bool)`

 SetStartupWizardCompletedNil sets the value for StartupWizardCompleted to be an explicit nil

### UnsetStartupWizardCompleted
`func (o *ModelSystemInfo) UnsetStartupWizardCompleted()`

UnsetStartupWizardCompleted ensures that no value is present for StartupWizardCompleted, not even an explicit nil
### GetOperatingSystemDisplayName

`func (o *ModelSystemInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *ModelSystemInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *ModelSystemInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *ModelSystemInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### SetOperatingSystemDisplayNameNil

`func (o *ModelSystemInfo) SetOperatingSystemDisplayNameNil(b bool)`

 SetOperatingSystemDisplayNameNil sets the value for OperatingSystemDisplayName to be an explicit nil

### UnsetOperatingSystemDisplayName
`func (o *ModelSystemInfo) UnsetOperatingSystemDisplayName()`

UnsetOperatingSystemDisplayName ensures that no value is present for OperatingSystemDisplayName, not even an explicit nil
### GetPackageName

`func (o *ModelSystemInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ModelSystemInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ModelSystemInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ModelSystemInfo) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### SetPackageNameNil

`func (o *ModelSystemInfo) SetPackageNameNil(b bool)`

 SetPackageNameNil sets the value for PackageName to be an explicit nil

### UnsetPackageName
`func (o *ModelSystemInfo) UnsetPackageName()`

UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
### GetHasPendingRestart

`func (o *ModelSystemInfo) GetHasPendingRestart() bool`

GetHasPendingRestart returns the HasPendingRestart field if non-nil, zero value otherwise.

### GetHasPendingRestartOk

`func (o *ModelSystemInfo) GetHasPendingRestartOk() (*bool, bool)`

GetHasPendingRestartOk returns a tuple with the HasPendingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingRestart

`func (o *ModelSystemInfo) SetHasPendingRestart(v bool)`

SetHasPendingRestart sets HasPendingRestart field to given value.

### HasHasPendingRestart

`func (o *ModelSystemInfo) HasHasPendingRestart() bool`

HasHasPendingRestart returns a boolean if a field has been set.

### GetIsShuttingDown

`func (o *ModelSystemInfo) GetIsShuttingDown() bool`

GetIsShuttingDown returns the IsShuttingDown field if non-nil, zero value otherwise.

### GetIsShuttingDownOk

`func (o *ModelSystemInfo) GetIsShuttingDownOk() (*bool, bool)`

GetIsShuttingDownOk returns a tuple with the IsShuttingDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuttingDown

`func (o *ModelSystemInfo) SetIsShuttingDown(v bool)`

SetIsShuttingDown sets IsShuttingDown field to given value.

### HasIsShuttingDown

`func (o *ModelSystemInfo) HasIsShuttingDown() bool`

HasIsShuttingDown returns a boolean if a field has been set.

### GetSupportsLibraryMonitor

`func (o *ModelSystemInfo) GetSupportsLibraryMonitor() bool`

GetSupportsLibraryMonitor returns the SupportsLibraryMonitor field if non-nil, zero value otherwise.

### GetSupportsLibraryMonitorOk

`func (o *ModelSystemInfo) GetSupportsLibraryMonitorOk() (*bool, bool)`

GetSupportsLibraryMonitorOk returns a tuple with the SupportsLibraryMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLibraryMonitor

`func (o *ModelSystemInfo) SetSupportsLibraryMonitor(v bool)`

SetSupportsLibraryMonitor sets SupportsLibraryMonitor field to given value.

### HasSupportsLibraryMonitor

`func (o *ModelSystemInfo) HasSupportsLibraryMonitor() bool`

HasSupportsLibraryMonitor returns a boolean if a field has been set.

### GetWebSocketPortNumber

`func (o *ModelSystemInfo) GetWebSocketPortNumber() int32`

GetWebSocketPortNumber returns the WebSocketPortNumber field if non-nil, zero value otherwise.

### GetWebSocketPortNumberOk

`func (o *ModelSystemInfo) GetWebSocketPortNumberOk() (*int32, bool)`

GetWebSocketPortNumberOk returns a tuple with the WebSocketPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSocketPortNumber

`func (o *ModelSystemInfo) SetWebSocketPortNumber(v int32)`

SetWebSocketPortNumber sets WebSocketPortNumber field to given value.

### HasWebSocketPortNumber

`func (o *ModelSystemInfo) HasWebSocketPortNumber() bool`

HasWebSocketPortNumber returns a boolean if a field has been set.

### GetCompletedInstallations

`func (o *ModelSystemInfo) GetCompletedInstallations() []ModelModelInstallationInfo`

GetCompletedInstallations returns the CompletedInstallations field if non-nil, zero value otherwise.

### GetCompletedInstallationsOk

`func (o *ModelSystemInfo) GetCompletedInstallationsOk() (*[]ModelModelInstallationInfo, bool)`

GetCompletedInstallationsOk returns a tuple with the CompletedInstallations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInstallations

`func (o *ModelSystemInfo) SetCompletedInstallations(v []ModelModelInstallationInfo)`

SetCompletedInstallations sets CompletedInstallations field to given value.

### HasCompletedInstallations

`func (o *ModelSystemInfo) HasCompletedInstallations() bool`

HasCompletedInstallations returns a boolean if a field has been set.

### SetCompletedInstallationsNil

`func (o *ModelSystemInfo) SetCompletedInstallationsNil(b bool)`

 SetCompletedInstallationsNil sets the value for CompletedInstallations to be an explicit nil

### UnsetCompletedInstallations
`func (o *ModelSystemInfo) UnsetCompletedInstallations()`

UnsetCompletedInstallations ensures that no value is present for CompletedInstallations, not even an explicit nil
### GetCanSelfRestart

`func (o *ModelSystemInfo) GetCanSelfRestart() bool`

GetCanSelfRestart returns the CanSelfRestart field if non-nil, zero value otherwise.

### GetCanSelfRestartOk

`func (o *ModelSystemInfo) GetCanSelfRestartOk() (*bool, bool)`

GetCanSelfRestartOk returns a tuple with the CanSelfRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfRestart

`func (o *ModelSystemInfo) SetCanSelfRestart(v bool)`

SetCanSelfRestart sets CanSelfRestart field to given value.

### HasCanSelfRestart

`func (o *ModelSystemInfo) HasCanSelfRestart() bool`

HasCanSelfRestart returns a boolean if a field has been set.

### GetCanLaunchWebBrowser

`func (o *ModelSystemInfo) GetCanLaunchWebBrowser() bool`

GetCanLaunchWebBrowser returns the CanLaunchWebBrowser field if non-nil, zero value otherwise.

### GetCanLaunchWebBrowserOk

`func (o *ModelSystemInfo) GetCanLaunchWebBrowserOk() (*bool, bool)`

GetCanLaunchWebBrowserOk returns a tuple with the CanLaunchWebBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchWebBrowser

`func (o *ModelSystemInfo) SetCanLaunchWebBrowser(v bool)`

SetCanLaunchWebBrowser sets CanLaunchWebBrowser field to given value.

### HasCanLaunchWebBrowser

`func (o *ModelSystemInfo) HasCanLaunchWebBrowser() bool`

HasCanLaunchWebBrowser returns a boolean if a field has been set.

### GetProgramDataPath

`func (o *ModelSystemInfo) GetProgramDataPath() string`

GetProgramDataPath returns the ProgramDataPath field if non-nil, zero value otherwise.

### GetProgramDataPathOk

`func (o *ModelSystemInfo) GetProgramDataPathOk() (*string, bool)`

GetProgramDataPathOk returns a tuple with the ProgramDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDataPath

`func (o *ModelSystemInfo) SetProgramDataPath(v string)`

SetProgramDataPath sets ProgramDataPath field to given value.

### HasProgramDataPath

`func (o *ModelSystemInfo) HasProgramDataPath() bool`

HasProgramDataPath returns a boolean if a field has been set.

### SetProgramDataPathNil

`func (o *ModelSystemInfo) SetProgramDataPathNil(b bool)`

 SetProgramDataPathNil sets the value for ProgramDataPath to be an explicit nil

### UnsetProgramDataPath
`func (o *ModelSystemInfo) UnsetProgramDataPath()`

UnsetProgramDataPath ensures that no value is present for ProgramDataPath, not even an explicit nil
### GetWebPath

`func (o *ModelSystemInfo) GetWebPath() string`

GetWebPath returns the WebPath field if non-nil, zero value otherwise.

### GetWebPathOk

`func (o *ModelSystemInfo) GetWebPathOk() (*string, bool)`

GetWebPathOk returns a tuple with the WebPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPath

`func (o *ModelSystemInfo) SetWebPath(v string)`

SetWebPath sets WebPath field to given value.

### HasWebPath

`func (o *ModelSystemInfo) HasWebPath() bool`

HasWebPath returns a boolean if a field has been set.

### SetWebPathNil

`func (o *ModelSystemInfo) SetWebPathNil(b bool)`

 SetWebPathNil sets the value for WebPath to be an explicit nil

### UnsetWebPath
`func (o *ModelSystemInfo) UnsetWebPath()`

UnsetWebPath ensures that no value is present for WebPath, not even an explicit nil
### GetItemsByNamePath

`func (o *ModelSystemInfo) GetItemsByNamePath() string`

GetItemsByNamePath returns the ItemsByNamePath field if non-nil, zero value otherwise.

### GetItemsByNamePathOk

`func (o *ModelSystemInfo) GetItemsByNamePathOk() (*string, bool)`

GetItemsByNamePathOk returns a tuple with the ItemsByNamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsByNamePath

`func (o *ModelSystemInfo) SetItemsByNamePath(v string)`

SetItemsByNamePath sets ItemsByNamePath field to given value.

### HasItemsByNamePath

`func (o *ModelSystemInfo) HasItemsByNamePath() bool`

HasItemsByNamePath returns a boolean if a field has been set.

### SetItemsByNamePathNil

`func (o *ModelSystemInfo) SetItemsByNamePathNil(b bool)`

 SetItemsByNamePathNil sets the value for ItemsByNamePath to be an explicit nil

### UnsetItemsByNamePath
`func (o *ModelSystemInfo) UnsetItemsByNamePath()`

UnsetItemsByNamePath ensures that no value is present for ItemsByNamePath, not even an explicit nil
### GetCachePath

`func (o *ModelSystemInfo) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ModelSystemInfo) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ModelSystemInfo) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ModelSystemInfo) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### SetCachePathNil

`func (o *ModelSystemInfo) SetCachePathNil(b bool)`

 SetCachePathNil sets the value for CachePath to be an explicit nil

### UnsetCachePath
`func (o *ModelSystemInfo) UnsetCachePath()`

UnsetCachePath ensures that no value is present for CachePath, not even an explicit nil
### GetLogPath

`func (o *ModelSystemInfo) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *ModelSystemInfo) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *ModelSystemInfo) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *ModelSystemInfo) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### SetLogPathNil

`func (o *ModelSystemInfo) SetLogPathNil(b bool)`

 SetLogPathNil sets the value for LogPath to be an explicit nil

### UnsetLogPath
`func (o *ModelSystemInfo) UnsetLogPath()`

UnsetLogPath ensures that no value is present for LogPath, not even an explicit nil
### GetInternalMetadataPath

`func (o *ModelSystemInfo) GetInternalMetadataPath() string`

GetInternalMetadataPath returns the InternalMetadataPath field if non-nil, zero value otherwise.

### GetInternalMetadataPathOk

`func (o *ModelSystemInfo) GetInternalMetadataPathOk() (*string, bool)`

GetInternalMetadataPathOk returns a tuple with the InternalMetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadataPath

`func (o *ModelSystemInfo) SetInternalMetadataPath(v string)`

SetInternalMetadataPath sets InternalMetadataPath field to given value.

### HasInternalMetadataPath

`func (o *ModelSystemInfo) HasInternalMetadataPath() bool`

HasInternalMetadataPath returns a boolean if a field has been set.

### SetInternalMetadataPathNil

`func (o *ModelSystemInfo) SetInternalMetadataPathNil(b bool)`

 SetInternalMetadataPathNil sets the value for InternalMetadataPath to be an explicit nil

### UnsetInternalMetadataPath
`func (o *ModelSystemInfo) UnsetInternalMetadataPath()`

UnsetInternalMetadataPath ensures that no value is present for InternalMetadataPath, not even an explicit nil
### GetTranscodingTempPath

`func (o *ModelSystemInfo) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *ModelSystemInfo) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *ModelSystemInfo) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *ModelSystemInfo) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### SetTranscodingTempPathNil

`func (o *ModelSystemInfo) SetTranscodingTempPathNil(b bool)`

 SetTranscodingTempPathNil sets the value for TranscodingTempPath to be an explicit nil

### UnsetTranscodingTempPath
`func (o *ModelSystemInfo) UnsetTranscodingTempPath()`

UnsetTranscodingTempPath ensures that no value is present for TranscodingTempPath, not even an explicit nil
### GetCastReceiverApplications

`func (o *ModelSystemInfo) GetCastReceiverApplications() []ModelModelCastReceiverApplication`

GetCastReceiverApplications returns the CastReceiverApplications field if non-nil, zero value otherwise.

### GetCastReceiverApplicationsOk

`func (o *ModelSystemInfo) GetCastReceiverApplicationsOk() (*[]ModelModelCastReceiverApplication, bool)`

GetCastReceiverApplicationsOk returns a tuple with the CastReceiverApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReceiverApplications

`func (o *ModelSystemInfo) SetCastReceiverApplications(v []ModelModelCastReceiverApplication)`

SetCastReceiverApplications sets CastReceiverApplications field to given value.

### HasCastReceiverApplications

`func (o *ModelSystemInfo) HasCastReceiverApplications() bool`

HasCastReceiverApplications returns a boolean if a field has been set.

### SetCastReceiverApplicationsNil

`func (o *ModelSystemInfo) SetCastReceiverApplicationsNil(b bool)`

 SetCastReceiverApplicationsNil sets the value for CastReceiverApplications to be an explicit nil

### UnsetCastReceiverApplications
`func (o *ModelSystemInfo) UnsetCastReceiverApplications()`

UnsetCastReceiverApplications ensures that no value is present for CastReceiverApplications, not even an explicit nil
### GetHasUpdateAvailable

`func (o *ModelSystemInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *ModelSystemInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *ModelSystemInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *ModelSystemInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetEncoderLocation

`func (o *ModelSystemInfo) GetEncoderLocation() string`

GetEncoderLocation returns the EncoderLocation field if non-nil, zero value otherwise.

### GetEncoderLocationOk

`func (o *ModelSystemInfo) GetEncoderLocationOk() (*string, bool)`

GetEncoderLocationOk returns a tuple with the EncoderLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderLocation

`func (o *ModelSystemInfo) SetEncoderLocation(v string)`

SetEncoderLocation sets EncoderLocation field to given value.

### HasEncoderLocation

`func (o *ModelSystemInfo) HasEncoderLocation() bool`

HasEncoderLocation returns a boolean if a field has been set.

### SetEncoderLocationNil

`func (o *ModelSystemInfo) SetEncoderLocationNil(b bool)`

 SetEncoderLocationNil sets the value for EncoderLocation to be an explicit nil

### UnsetEncoderLocation
`func (o *ModelSystemInfo) UnsetEncoderLocation()`

UnsetEncoderLocation ensures that no value is present for EncoderLocation, not even an explicit nil
### GetSystemArchitecture

`func (o *ModelSystemInfo) GetSystemArchitecture() string`

GetSystemArchitecture returns the SystemArchitecture field if non-nil, zero value otherwise.

### GetSystemArchitectureOk

`func (o *ModelSystemInfo) GetSystemArchitectureOk() (*string, bool)`

GetSystemArchitectureOk returns a tuple with the SystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArchitecture

`func (o *ModelSystemInfo) SetSystemArchitecture(v string)`

SetSystemArchitecture sets SystemArchitecture field to given value.

### HasSystemArchitecture

`func (o *ModelSystemInfo) HasSystemArchitecture() bool`

HasSystemArchitecture returns a boolean if a field has been set.

### SetSystemArchitectureNil

`func (o *ModelSystemInfo) SetSystemArchitectureNil(b bool)`

 SetSystemArchitectureNil sets the value for SystemArchitecture to be an explicit nil

### UnsetSystemArchitecture
`func (o *ModelSystemInfo) UnsetSystemArchitecture()`

UnsetSystemArchitecture ensures that no value is present for SystemArchitecture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


