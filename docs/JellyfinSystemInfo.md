# JellyfinSystemInfo

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
**CompletedInstallations** | Pointer to [**[]JellyfinJellyfinInstallationInfo**](JellyfinJellyfinInstallationInfo.md) | Gets or sets the completed installations. | [optional] 
**CanSelfRestart** | Pointer to **bool** | Gets or sets a value indicating whether this instance can self restart. | [optional] [default to true]
**CanLaunchWebBrowser** | Pointer to **bool** |  | [optional] [default to false]
**ProgramDataPath** | Pointer to **NullableString** | Gets or sets the program data path. | [optional] 
**WebPath** | Pointer to **NullableString** | Gets or sets the web UI resources path. | [optional] 
**ItemsByNamePath** | Pointer to **NullableString** | Gets or sets the items by name path. | [optional] 
**CachePath** | Pointer to **NullableString** | Gets or sets the cache path. | [optional] 
**LogPath** | Pointer to **NullableString** | Gets or sets the log path. | [optional] 
**InternalMetadataPath** | Pointer to **NullableString** | Gets or sets the internal metadata path. | [optional] 
**TranscodingTempPath** | Pointer to **NullableString** | Gets or sets the transcode path. | [optional] 
**CastReceiverApplications** | Pointer to [**[]JellyfinJellyfinCastReceiverApplication**](JellyfinJellyfinCastReceiverApplication.md) | Gets or sets the list of cast receiver applications. | [optional] 
**HasUpdateAvailable** | Pointer to **bool** | Gets or sets a value indicating whether this instance has update available. | [optional] [default to false]
**EncoderLocation** | Pointer to **NullableString** |  | [optional] [default to "System"]
**SystemArchitecture** | Pointer to **NullableString** |  | [optional] [default to "X64"]

## Methods

### NewJellyfinSystemInfo

`func NewJellyfinSystemInfo() *JellyfinSystemInfo`

NewJellyfinSystemInfo instantiates a new JellyfinSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSystemInfoWithDefaults

`func NewJellyfinSystemInfoWithDefaults() *JellyfinSystemInfo`

NewJellyfinSystemInfoWithDefaults instantiates a new JellyfinSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalAddress

`func (o *JellyfinSystemInfo) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *JellyfinSystemInfo) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *JellyfinSystemInfo) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *JellyfinSystemInfo) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### SetLocalAddressNil

`func (o *JellyfinSystemInfo) SetLocalAddressNil(b bool)`

 SetLocalAddressNil sets the value for LocalAddress to be an explicit nil

### UnsetLocalAddress
`func (o *JellyfinSystemInfo) UnsetLocalAddress()`

UnsetLocalAddress ensures that no value is present for LocalAddress, not even an explicit nil
### GetServerName

`func (o *JellyfinSystemInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *JellyfinSystemInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *JellyfinSystemInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *JellyfinSystemInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *JellyfinSystemInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *JellyfinSystemInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetVersion

`func (o *JellyfinSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JellyfinSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JellyfinSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JellyfinSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *JellyfinSystemInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *JellyfinSystemInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetProductName

`func (o *JellyfinSystemInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *JellyfinSystemInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *JellyfinSystemInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *JellyfinSystemInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *JellyfinSystemInfo) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *JellyfinSystemInfo) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetOperatingSystem

`func (o *JellyfinSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *JellyfinSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *JellyfinSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *JellyfinSystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *JellyfinSystemInfo) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *JellyfinSystemInfo) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetId

`func (o *JellyfinSystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinSystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinSystemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinSystemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinSystemInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinSystemInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartupWizardCompleted

`func (o *JellyfinSystemInfo) GetStartupWizardCompleted() bool`

GetStartupWizardCompleted returns the StartupWizardCompleted field if non-nil, zero value otherwise.

### GetStartupWizardCompletedOk

`func (o *JellyfinSystemInfo) GetStartupWizardCompletedOk() (*bool, bool)`

GetStartupWizardCompletedOk returns a tuple with the StartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupWizardCompleted

`func (o *JellyfinSystemInfo) SetStartupWizardCompleted(v bool)`

SetStartupWizardCompleted sets StartupWizardCompleted field to given value.

### HasStartupWizardCompleted

`func (o *JellyfinSystemInfo) HasStartupWizardCompleted() bool`

HasStartupWizardCompleted returns a boolean if a field has been set.

### SetStartupWizardCompletedNil

`func (o *JellyfinSystemInfo) SetStartupWizardCompletedNil(b bool)`

 SetStartupWizardCompletedNil sets the value for StartupWizardCompleted to be an explicit nil

### UnsetStartupWizardCompleted
`func (o *JellyfinSystemInfo) UnsetStartupWizardCompleted()`

UnsetStartupWizardCompleted ensures that no value is present for StartupWizardCompleted, not even an explicit nil
### GetOperatingSystemDisplayName

`func (o *JellyfinSystemInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *JellyfinSystemInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *JellyfinSystemInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *JellyfinSystemInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### SetOperatingSystemDisplayNameNil

`func (o *JellyfinSystemInfo) SetOperatingSystemDisplayNameNil(b bool)`

 SetOperatingSystemDisplayNameNil sets the value for OperatingSystemDisplayName to be an explicit nil

### UnsetOperatingSystemDisplayName
`func (o *JellyfinSystemInfo) UnsetOperatingSystemDisplayName()`

UnsetOperatingSystemDisplayName ensures that no value is present for OperatingSystemDisplayName, not even an explicit nil
### GetPackageName

`func (o *JellyfinSystemInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *JellyfinSystemInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *JellyfinSystemInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *JellyfinSystemInfo) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### SetPackageNameNil

`func (o *JellyfinSystemInfo) SetPackageNameNil(b bool)`

 SetPackageNameNil sets the value for PackageName to be an explicit nil

### UnsetPackageName
`func (o *JellyfinSystemInfo) UnsetPackageName()`

UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
### GetHasPendingRestart

`func (o *JellyfinSystemInfo) GetHasPendingRestart() bool`

GetHasPendingRestart returns the HasPendingRestart field if non-nil, zero value otherwise.

### GetHasPendingRestartOk

`func (o *JellyfinSystemInfo) GetHasPendingRestartOk() (*bool, bool)`

GetHasPendingRestartOk returns a tuple with the HasPendingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingRestart

`func (o *JellyfinSystemInfo) SetHasPendingRestart(v bool)`

SetHasPendingRestart sets HasPendingRestart field to given value.

### HasHasPendingRestart

`func (o *JellyfinSystemInfo) HasHasPendingRestart() bool`

HasHasPendingRestart returns a boolean if a field has been set.

### GetIsShuttingDown

`func (o *JellyfinSystemInfo) GetIsShuttingDown() bool`

GetIsShuttingDown returns the IsShuttingDown field if non-nil, zero value otherwise.

### GetIsShuttingDownOk

`func (o *JellyfinSystemInfo) GetIsShuttingDownOk() (*bool, bool)`

GetIsShuttingDownOk returns a tuple with the IsShuttingDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuttingDown

`func (o *JellyfinSystemInfo) SetIsShuttingDown(v bool)`

SetIsShuttingDown sets IsShuttingDown field to given value.

### HasIsShuttingDown

`func (o *JellyfinSystemInfo) HasIsShuttingDown() bool`

HasIsShuttingDown returns a boolean if a field has been set.

### GetSupportsLibraryMonitor

`func (o *JellyfinSystemInfo) GetSupportsLibraryMonitor() bool`

GetSupportsLibraryMonitor returns the SupportsLibraryMonitor field if non-nil, zero value otherwise.

### GetSupportsLibraryMonitorOk

`func (o *JellyfinSystemInfo) GetSupportsLibraryMonitorOk() (*bool, bool)`

GetSupportsLibraryMonitorOk returns a tuple with the SupportsLibraryMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLibraryMonitor

`func (o *JellyfinSystemInfo) SetSupportsLibraryMonitor(v bool)`

SetSupportsLibraryMonitor sets SupportsLibraryMonitor field to given value.

### HasSupportsLibraryMonitor

`func (o *JellyfinSystemInfo) HasSupportsLibraryMonitor() bool`

HasSupportsLibraryMonitor returns a boolean if a field has been set.

### GetWebSocketPortNumber

`func (o *JellyfinSystemInfo) GetWebSocketPortNumber() int32`

GetWebSocketPortNumber returns the WebSocketPortNumber field if non-nil, zero value otherwise.

### GetWebSocketPortNumberOk

`func (o *JellyfinSystemInfo) GetWebSocketPortNumberOk() (*int32, bool)`

GetWebSocketPortNumberOk returns a tuple with the WebSocketPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSocketPortNumber

`func (o *JellyfinSystemInfo) SetWebSocketPortNumber(v int32)`

SetWebSocketPortNumber sets WebSocketPortNumber field to given value.

### HasWebSocketPortNumber

`func (o *JellyfinSystemInfo) HasWebSocketPortNumber() bool`

HasWebSocketPortNumber returns a boolean if a field has been set.

### GetCompletedInstallations

`func (o *JellyfinSystemInfo) GetCompletedInstallations() []JellyfinJellyfinInstallationInfo`

GetCompletedInstallations returns the CompletedInstallations field if non-nil, zero value otherwise.

### GetCompletedInstallationsOk

`func (o *JellyfinSystemInfo) GetCompletedInstallationsOk() (*[]JellyfinJellyfinInstallationInfo, bool)`

GetCompletedInstallationsOk returns a tuple with the CompletedInstallations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInstallations

`func (o *JellyfinSystemInfo) SetCompletedInstallations(v []JellyfinJellyfinInstallationInfo)`

SetCompletedInstallations sets CompletedInstallations field to given value.

### HasCompletedInstallations

`func (o *JellyfinSystemInfo) HasCompletedInstallations() bool`

HasCompletedInstallations returns a boolean if a field has been set.

### SetCompletedInstallationsNil

`func (o *JellyfinSystemInfo) SetCompletedInstallationsNil(b bool)`

 SetCompletedInstallationsNil sets the value for CompletedInstallations to be an explicit nil

### UnsetCompletedInstallations
`func (o *JellyfinSystemInfo) UnsetCompletedInstallations()`

UnsetCompletedInstallations ensures that no value is present for CompletedInstallations, not even an explicit nil
### GetCanSelfRestart

`func (o *JellyfinSystemInfo) GetCanSelfRestart() bool`

GetCanSelfRestart returns the CanSelfRestart field if non-nil, zero value otherwise.

### GetCanSelfRestartOk

`func (o *JellyfinSystemInfo) GetCanSelfRestartOk() (*bool, bool)`

GetCanSelfRestartOk returns a tuple with the CanSelfRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfRestart

`func (o *JellyfinSystemInfo) SetCanSelfRestart(v bool)`

SetCanSelfRestart sets CanSelfRestart field to given value.

### HasCanSelfRestart

`func (o *JellyfinSystemInfo) HasCanSelfRestart() bool`

HasCanSelfRestart returns a boolean if a field has been set.

### GetCanLaunchWebBrowser

`func (o *JellyfinSystemInfo) GetCanLaunchWebBrowser() bool`

GetCanLaunchWebBrowser returns the CanLaunchWebBrowser field if non-nil, zero value otherwise.

### GetCanLaunchWebBrowserOk

`func (o *JellyfinSystemInfo) GetCanLaunchWebBrowserOk() (*bool, bool)`

GetCanLaunchWebBrowserOk returns a tuple with the CanLaunchWebBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchWebBrowser

`func (o *JellyfinSystemInfo) SetCanLaunchWebBrowser(v bool)`

SetCanLaunchWebBrowser sets CanLaunchWebBrowser field to given value.

### HasCanLaunchWebBrowser

`func (o *JellyfinSystemInfo) HasCanLaunchWebBrowser() bool`

HasCanLaunchWebBrowser returns a boolean if a field has been set.

### GetProgramDataPath

`func (o *JellyfinSystemInfo) GetProgramDataPath() string`

GetProgramDataPath returns the ProgramDataPath field if non-nil, zero value otherwise.

### GetProgramDataPathOk

`func (o *JellyfinSystemInfo) GetProgramDataPathOk() (*string, bool)`

GetProgramDataPathOk returns a tuple with the ProgramDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDataPath

`func (o *JellyfinSystemInfo) SetProgramDataPath(v string)`

SetProgramDataPath sets ProgramDataPath field to given value.

### HasProgramDataPath

`func (o *JellyfinSystemInfo) HasProgramDataPath() bool`

HasProgramDataPath returns a boolean if a field has been set.

### SetProgramDataPathNil

`func (o *JellyfinSystemInfo) SetProgramDataPathNil(b bool)`

 SetProgramDataPathNil sets the value for ProgramDataPath to be an explicit nil

### UnsetProgramDataPath
`func (o *JellyfinSystemInfo) UnsetProgramDataPath()`

UnsetProgramDataPath ensures that no value is present for ProgramDataPath, not even an explicit nil
### GetWebPath

`func (o *JellyfinSystemInfo) GetWebPath() string`

GetWebPath returns the WebPath field if non-nil, zero value otherwise.

### GetWebPathOk

`func (o *JellyfinSystemInfo) GetWebPathOk() (*string, bool)`

GetWebPathOk returns a tuple with the WebPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPath

`func (o *JellyfinSystemInfo) SetWebPath(v string)`

SetWebPath sets WebPath field to given value.

### HasWebPath

`func (o *JellyfinSystemInfo) HasWebPath() bool`

HasWebPath returns a boolean if a field has been set.

### SetWebPathNil

`func (o *JellyfinSystemInfo) SetWebPathNil(b bool)`

 SetWebPathNil sets the value for WebPath to be an explicit nil

### UnsetWebPath
`func (o *JellyfinSystemInfo) UnsetWebPath()`

UnsetWebPath ensures that no value is present for WebPath, not even an explicit nil
### GetItemsByNamePath

`func (o *JellyfinSystemInfo) GetItemsByNamePath() string`

GetItemsByNamePath returns the ItemsByNamePath field if non-nil, zero value otherwise.

### GetItemsByNamePathOk

`func (o *JellyfinSystemInfo) GetItemsByNamePathOk() (*string, bool)`

GetItemsByNamePathOk returns a tuple with the ItemsByNamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsByNamePath

`func (o *JellyfinSystemInfo) SetItemsByNamePath(v string)`

SetItemsByNamePath sets ItemsByNamePath field to given value.

### HasItemsByNamePath

`func (o *JellyfinSystemInfo) HasItemsByNamePath() bool`

HasItemsByNamePath returns a boolean if a field has been set.

### SetItemsByNamePathNil

`func (o *JellyfinSystemInfo) SetItemsByNamePathNil(b bool)`

 SetItemsByNamePathNil sets the value for ItemsByNamePath to be an explicit nil

### UnsetItemsByNamePath
`func (o *JellyfinSystemInfo) UnsetItemsByNamePath()`

UnsetItemsByNamePath ensures that no value is present for ItemsByNamePath, not even an explicit nil
### GetCachePath

`func (o *JellyfinSystemInfo) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *JellyfinSystemInfo) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *JellyfinSystemInfo) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *JellyfinSystemInfo) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### SetCachePathNil

`func (o *JellyfinSystemInfo) SetCachePathNil(b bool)`

 SetCachePathNil sets the value for CachePath to be an explicit nil

### UnsetCachePath
`func (o *JellyfinSystemInfo) UnsetCachePath()`

UnsetCachePath ensures that no value is present for CachePath, not even an explicit nil
### GetLogPath

`func (o *JellyfinSystemInfo) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *JellyfinSystemInfo) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *JellyfinSystemInfo) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *JellyfinSystemInfo) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### SetLogPathNil

`func (o *JellyfinSystemInfo) SetLogPathNil(b bool)`

 SetLogPathNil sets the value for LogPath to be an explicit nil

### UnsetLogPath
`func (o *JellyfinSystemInfo) UnsetLogPath()`

UnsetLogPath ensures that no value is present for LogPath, not even an explicit nil
### GetInternalMetadataPath

`func (o *JellyfinSystemInfo) GetInternalMetadataPath() string`

GetInternalMetadataPath returns the InternalMetadataPath field if non-nil, zero value otherwise.

### GetInternalMetadataPathOk

`func (o *JellyfinSystemInfo) GetInternalMetadataPathOk() (*string, bool)`

GetInternalMetadataPathOk returns a tuple with the InternalMetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadataPath

`func (o *JellyfinSystemInfo) SetInternalMetadataPath(v string)`

SetInternalMetadataPath sets InternalMetadataPath field to given value.

### HasInternalMetadataPath

`func (o *JellyfinSystemInfo) HasInternalMetadataPath() bool`

HasInternalMetadataPath returns a boolean if a field has been set.

### SetInternalMetadataPathNil

`func (o *JellyfinSystemInfo) SetInternalMetadataPathNil(b bool)`

 SetInternalMetadataPathNil sets the value for InternalMetadataPath to be an explicit nil

### UnsetInternalMetadataPath
`func (o *JellyfinSystemInfo) UnsetInternalMetadataPath()`

UnsetInternalMetadataPath ensures that no value is present for InternalMetadataPath, not even an explicit nil
### GetTranscodingTempPath

`func (o *JellyfinSystemInfo) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *JellyfinSystemInfo) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *JellyfinSystemInfo) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *JellyfinSystemInfo) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### SetTranscodingTempPathNil

`func (o *JellyfinSystemInfo) SetTranscodingTempPathNil(b bool)`

 SetTranscodingTempPathNil sets the value for TranscodingTempPath to be an explicit nil

### UnsetTranscodingTempPath
`func (o *JellyfinSystemInfo) UnsetTranscodingTempPath()`

UnsetTranscodingTempPath ensures that no value is present for TranscodingTempPath, not even an explicit nil
### GetCastReceiverApplications

`func (o *JellyfinSystemInfo) GetCastReceiverApplications() []JellyfinJellyfinCastReceiverApplication`

GetCastReceiverApplications returns the CastReceiverApplications field if non-nil, zero value otherwise.

### GetCastReceiverApplicationsOk

`func (o *JellyfinSystemInfo) GetCastReceiverApplicationsOk() (*[]JellyfinJellyfinCastReceiverApplication, bool)`

GetCastReceiverApplicationsOk returns a tuple with the CastReceiverApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReceiverApplications

`func (o *JellyfinSystemInfo) SetCastReceiverApplications(v []JellyfinJellyfinCastReceiverApplication)`

SetCastReceiverApplications sets CastReceiverApplications field to given value.

### HasCastReceiverApplications

`func (o *JellyfinSystemInfo) HasCastReceiverApplications() bool`

HasCastReceiverApplications returns a boolean if a field has been set.

### SetCastReceiverApplicationsNil

`func (o *JellyfinSystemInfo) SetCastReceiverApplicationsNil(b bool)`

 SetCastReceiverApplicationsNil sets the value for CastReceiverApplications to be an explicit nil

### UnsetCastReceiverApplications
`func (o *JellyfinSystemInfo) UnsetCastReceiverApplications()`

UnsetCastReceiverApplications ensures that no value is present for CastReceiverApplications, not even an explicit nil
### GetHasUpdateAvailable

`func (o *JellyfinSystemInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *JellyfinSystemInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *JellyfinSystemInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *JellyfinSystemInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetEncoderLocation

`func (o *JellyfinSystemInfo) GetEncoderLocation() string`

GetEncoderLocation returns the EncoderLocation field if non-nil, zero value otherwise.

### GetEncoderLocationOk

`func (o *JellyfinSystemInfo) GetEncoderLocationOk() (*string, bool)`

GetEncoderLocationOk returns a tuple with the EncoderLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderLocation

`func (o *JellyfinSystemInfo) SetEncoderLocation(v string)`

SetEncoderLocation sets EncoderLocation field to given value.

### HasEncoderLocation

`func (o *JellyfinSystemInfo) HasEncoderLocation() bool`

HasEncoderLocation returns a boolean if a field has been set.

### SetEncoderLocationNil

`func (o *JellyfinSystemInfo) SetEncoderLocationNil(b bool)`

 SetEncoderLocationNil sets the value for EncoderLocation to be an explicit nil

### UnsetEncoderLocation
`func (o *JellyfinSystemInfo) UnsetEncoderLocation()`

UnsetEncoderLocation ensures that no value is present for EncoderLocation, not even an explicit nil
### GetSystemArchitecture

`func (o *JellyfinSystemInfo) GetSystemArchitecture() string`

GetSystemArchitecture returns the SystemArchitecture field if non-nil, zero value otherwise.

### GetSystemArchitectureOk

`func (o *JellyfinSystemInfo) GetSystemArchitectureOk() (*string, bool)`

GetSystemArchitectureOk returns a tuple with the SystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemArchitecture

`func (o *JellyfinSystemInfo) SetSystemArchitecture(v string)`

SetSystemArchitecture sets SystemArchitecture field to given value.

### HasSystemArchitecture

`func (o *JellyfinSystemInfo) HasSystemArchitecture() bool`

HasSystemArchitecture returns a boolean if a field has been set.

### SetSystemArchitectureNil

`func (o *JellyfinSystemInfo) SetSystemArchitectureNil(b bool)`

 SetSystemArchitectureNil sets the value for SystemArchitecture to be an explicit nil

### UnsetSystemArchitecture
`func (o *JellyfinSystemInfo) UnsetSystemArchitecture()`

UnsetSystemArchitecture ensures that no value is present for SystemArchitecture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


