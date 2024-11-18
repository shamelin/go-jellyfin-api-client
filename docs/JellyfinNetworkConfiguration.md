# JellyfinNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | Gets or sets a value used to specify the URL prefix that your Jellyfin instance can be accessed at. | [optional] 
**EnableHttps** | Pointer to **bool** | Gets or sets a value indicating whether to use HTTPS. | [optional] 
**RequireHttps** | Pointer to **bool** | Gets or sets a value indicating whether the server should force connections over HTTPS. | [optional] 
**CertificatePath** | Pointer to **string** | Gets or sets the filesystem path of an X.509 certificate to use for SSL. | [optional] 
**CertificatePassword** | Pointer to **string** | Gets or sets the password required to access the X.509 certificate data in the file specified by MediaBrowser.Common.Net.NetworkConfiguration.CertificatePath. | [optional] 
**InternalHttpPort** | Pointer to **int32** | Gets or sets the internal HTTP server port. | [optional] 
**InternalHttpsPort** | Pointer to **int32** | Gets or sets the internal HTTPS server port. | [optional] 
**PublicHttpPort** | Pointer to **int32** | Gets or sets the public HTTP port. | [optional] 
**PublicHttpsPort** | Pointer to **int32** | Gets or sets the public HTTPS port. | [optional] 
**AutoDiscovery** | Pointer to **bool** | Gets or sets a value indicating whether Autodiscovery is enabled. | [optional] 
**EnableUPnP** | Pointer to **bool** | Gets or sets a value indicating whether to enable automatic port forwarding. | [optional] 
**EnableIPv4** | Pointer to **bool** | Gets or sets a value indicating whether IPv6 is enabled. | [optional] 
**EnableIPv6** | Pointer to **bool** | Gets or sets a value indicating whether IPv6 is enabled. | [optional] 
**EnableRemoteAccess** | Pointer to **bool** | Gets or sets a value indicating whether access from outside of the LAN is permitted. | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** | Gets or sets the subnets that are deemed to make up the LAN. | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** | Gets or sets the interface addresses which Jellyfin will bind to. If empty, all interfaces will be used. | [optional] 
**KnownProxies** | Pointer to **[]string** | Gets or sets the known proxies. | [optional] 
**IgnoreVirtualInterfaces** | Pointer to **bool** | Gets or sets a value indicating whether address names that match MediaBrowser.Common.Net.NetworkConfiguration.VirtualInterfaceNames should be ignored for the purposes of binding. | [optional] 
**VirtualInterfaceNames** | Pointer to **[]string** | Gets or sets a value indicating the interface name prefixes that should be ignored. The list can be comma separated and values are case-insensitive. &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.IgnoreVirtualInterfaces\&quot; /&gt;. | [optional] 
**EnablePublishedServerUriByRequest** | Pointer to **bool** | Gets or sets a value indicating whether the published server uri is based on information in HTTP requests. | [optional] 
**PublishedServerUriBySubnet** | Pointer to **[]string** | Gets or sets the PublishedServerUriBySubnet  Gets or sets PublishedServerUri to advertise for specific subnets. | [optional] 
**RemoteIPFilter** | Pointer to **[]string** | Gets or sets the filter for remote IP connectivity. Used in conjunction with &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.IsRemoteIPFilterBlacklist\&quot; /&gt;. | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** | Gets or sets a value indicating whether &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.RemoteIPFilter\&quot; /&gt; contains a blacklist or a whitelist. Default is a whitelist. | [optional] 

## Methods

### NewJellyfinNetworkConfiguration

`func NewJellyfinNetworkConfiguration() *JellyfinNetworkConfiguration`

NewJellyfinNetworkConfiguration instantiates a new JellyfinNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinNetworkConfigurationWithDefaults

`func NewJellyfinNetworkConfigurationWithDefaults() *JellyfinNetworkConfiguration`

NewJellyfinNetworkConfigurationWithDefaults instantiates a new JellyfinNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *JellyfinNetworkConfiguration) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *JellyfinNetworkConfiguration) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *JellyfinNetworkConfiguration) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *JellyfinNetworkConfiguration) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetEnableHttps

`func (o *JellyfinNetworkConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *JellyfinNetworkConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *JellyfinNetworkConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *JellyfinNetworkConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetRequireHttps

`func (o *JellyfinNetworkConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *JellyfinNetworkConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *JellyfinNetworkConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *JellyfinNetworkConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetCertificatePath

`func (o *JellyfinNetworkConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *JellyfinNetworkConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *JellyfinNetworkConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *JellyfinNetworkConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *JellyfinNetworkConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *JellyfinNetworkConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *JellyfinNetworkConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *JellyfinNetworkConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetInternalHttpPort

`func (o *JellyfinNetworkConfiguration) GetInternalHttpPort() int32`

GetInternalHttpPort returns the InternalHttpPort field if non-nil, zero value otherwise.

### GetInternalHttpPortOk

`func (o *JellyfinNetworkConfiguration) GetInternalHttpPortOk() (*int32, bool)`

GetInternalHttpPortOk returns a tuple with the InternalHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpPort

`func (o *JellyfinNetworkConfiguration) SetInternalHttpPort(v int32)`

SetInternalHttpPort sets InternalHttpPort field to given value.

### HasInternalHttpPort

`func (o *JellyfinNetworkConfiguration) HasInternalHttpPort() bool`

HasInternalHttpPort returns a boolean if a field has been set.

### GetInternalHttpsPort

`func (o *JellyfinNetworkConfiguration) GetInternalHttpsPort() int32`

GetInternalHttpsPort returns the InternalHttpsPort field if non-nil, zero value otherwise.

### GetInternalHttpsPortOk

`func (o *JellyfinNetworkConfiguration) GetInternalHttpsPortOk() (*int32, bool)`

GetInternalHttpsPortOk returns a tuple with the InternalHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpsPort

`func (o *JellyfinNetworkConfiguration) SetInternalHttpsPort(v int32)`

SetInternalHttpsPort sets InternalHttpsPort field to given value.

### HasInternalHttpsPort

`func (o *JellyfinNetworkConfiguration) HasInternalHttpsPort() bool`

HasInternalHttpsPort returns a boolean if a field has been set.

### GetPublicHttpPort

`func (o *JellyfinNetworkConfiguration) GetPublicHttpPort() int32`

GetPublicHttpPort returns the PublicHttpPort field if non-nil, zero value otherwise.

### GetPublicHttpPortOk

`func (o *JellyfinNetworkConfiguration) GetPublicHttpPortOk() (*int32, bool)`

GetPublicHttpPortOk returns a tuple with the PublicHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpPort

`func (o *JellyfinNetworkConfiguration) SetPublicHttpPort(v int32)`

SetPublicHttpPort sets PublicHttpPort field to given value.

### HasPublicHttpPort

`func (o *JellyfinNetworkConfiguration) HasPublicHttpPort() bool`

HasPublicHttpPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *JellyfinNetworkConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *JellyfinNetworkConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *JellyfinNetworkConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *JellyfinNetworkConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetAutoDiscovery

`func (o *JellyfinNetworkConfiguration) GetAutoDiscovery() bool`

GetAutoDiscovery returns the AutoDiscovery field if non-nil, zero value otherwise.

### GetAutoDiscoveryOk

`func (o *JellyfinNetworkConfiguration) GetAutoDiscoveryOk() (*bool, bool)`

GetAutoDiscoveryOk returns a tuple with the AutoDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscovery

`func (o *JellyfinNetworkConfiguration) SetAutoDiscovery(v bool)`

SetAutoDiscovery sets AutoDiscovery field to given value.

### HasAutoDiscovery

`func (o *JellyfinNetworkConfiguration) HasAutoDiscovery() bool`

HasAutoDiscovery returns a boolean if a field has been set.

### GetEnableUPnP

`func (o *JellyfinNetworkConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *JellyfinNetworkConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *JellyfinNetworkConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *JellyfinNetworkConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetEnableIPv4

`func (o *JellyfinNetworkConfiguration) GetEnableIPv4() bool`

GetEnableIPv4 returns the EnableIPv4 field if non-nil, zero value otherwise.

### GetEnableIPv4Ok

`func (o *JellyfinNetworkConfiguration) GetEnableIPv4Ok() (*bool, bool)`

GetEnableIPv4Ok returns a tuple with the EnableIPv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv4

`func (o *JellyfinNetworkConfiguration) SetEnableIPv4(v bool)`

SetEnableIPv4 sets EnableIPv4 field to given value.

### HasEnableIPv4

`func (o *JellyfinNetworkConfiguration) HasEnableIPv4() bool`

HasEnableIPv4 returns a boolean if a field has been set.

### GetEnableIPv6

`func (o *JellyfinNetworkConfiguration) GetEnableIPv6() bool`

GetEnableIPv6 returns the EnableIPv6 field if non-nil, zero value otherwise.

### GetEnableIPv6Ok

`func (o *JellyfinNetworkConfiguration) GetEnableIPv6Ok() (*bool, bool)`

GetEnableIPv6Ok returns a tuple with the EnableIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv6

`func (o *JellyfinNetworkConfiguration) SetEnableIPv6(v bool)`

SetEnableIPv6 sets EnableIPv6 field to given value.

### HasEnableIPv6

`func (o *JellyfinNetworkConfiguration) HasEnableIPv6() bool`

HasEnableIPv6 returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *JellyfinNetworkConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *JellyfinNetworkConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *JellyfinNetworkConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *JellyfinNetworkConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *JellyfinNetworkConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *JellyfinNetworkConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *JellyfinNetworkConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *JellyfinNetworkConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *JellyfinNetworkConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *JellyfinNetworkConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *JellyfinNetworkConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *JellyfinNetworkConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetKnownProxies

`func (o *JellyfinNetworkConfiguration) GetKnownProxies() []string`

GetKnownProxies returns the KnownProxies field if non-nil, zero value otherwise.

### GetKnownProxiesOk

`func (o *JellyfinNetworkConfiguration) GetKnownProxiesOk() (*[]string, bool)`

GetKnownProxiesOk returns a tuple with the KnownProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownProxies

`func (o *JellyfinNetworkConfiguration) SetKnownProxies(v []string)`

SetKnownProxies sets KnownProxies field to given value.

### HasKnownProxies

`func (o *JellyfinNetworkConfiguration) HasKnownProxies() bool`

HasKnownProxies returns a boolean if a field has been set.

### GetIgnoreVirtualInterfaces

`func (o *JellyfinNetworkConfiguration) GetIgnoreVirtualInterfaces() bool`

GetIgnoreVirtualInterfaces returns the IgnoreVirtualInterfaces field if non-nil, zero value otherwise.

### GetIgnoreVirtualInterfacesOk

`func (o *JellyfinNetworkConfiguration) GetIgnoreVirtualInterfacesOk() (*bool, bool)`

GetIgnoreVirtualInterfacesOk returns a tuple with the IgnoreVirtualInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreVirtualInterfaces

`func (o *JellyfinNetworkConfiguration) SetIgnoreVirtualInterfaces(v bool)`

SetIgnoreVirtualInterfaces sets IgnoreVirtualInterfaces field to given value.

### HasIgnoreVirtualInterfaces

`func (o *JellyfinNetworkConfiguration) HasIgnoreVirtualInterfaces() bool`

HasIgnoreVirtualInterfaces returns a boolean if a field has been set.

### GetVirtualInterfaceNames

`func (o *JellyfinNetworkConfiguration) GetVirtualInterfaceNames() []string`

GetVirtualInterfaceNames returns the VirtualInterfaceNames field if non-nil, zero value otherwise.

### GetVirtualInterfaceNamesOk

`func (o *JellyfinNetworkConfiguration) GetVirtualInterfaceNamesOk() (*[]string, bool)`

GetVirtualInterfaceNamesOk returns a tuple with the VirtualInterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaceNames

`func (o *JellyfinNetworkConfiguration) SetVirtualInterfaceNames(v []string)`

SetVirtualInterfaceNames sets VirtualInterfaceNames field to given value.

### HasVirtualInterfaceNames

`func (o *JellyfinNetworkConfiguration) HasVirtualInterfaceNames() bool`

HasVirtualInterfaceNames returns a boolean if a field has been set.

### GetEnablePublishedServerUriByRequest

`func (o *JellyfinNetworkConfiguration) GetEnablePublishedServerUriByRequest() bool`

GetEnablePublishedServerUriByRequest returns the EnablePublishedServerUriByRequest field if non-nil, zero value otherwise.

### GetEnablePublishedServerUriByRequestOk

`func (o *JellyfinNetworkConfiguration) GetEnablePublishedServerUriByRequestOk() (*bool, bool)`

GetEnablePublishedServerUriByRequestOk returns a tuple with the EnablePublishedServerUriByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublishedServerUriByRequest

`func (o *JellyfinNetworkConfiguration) SetEnablePublishedServerUriByRequest(v bool)`

SetEnablePublishedServerUriByRequest sets EnablePublishedServerUriByRequest field to given value.

### HasEnablePublishedServerUriByRequest

`func (o *JellyfinNetworkConfiguration) HasEnablePublishedServerUriByRequest() bool`

HasEnablePublishedServerUriByRequest returns a boolean if a field has been set.

### GetPublishedServerUriBySubnet

`func (o *JellyfinNetworkConfiguration) GetPublishedServerUriBySubnet() []string`

GetPublishedServerUriBySubnet returns the PublishedServerUriBySubnet field if non-nil, zero value otherwise.

### GetPublishedServerUriBySubnetOk

`func (o *JellyfinNetworkConfiguration) GetPublishedServerUriBySubnetOk() (*[]string, bool)`

GetPublishedServerUriBySubnetOk returns a tuple with the PublishedServerUriBySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServerUriBySubnet

`func (o *JellyfinNetworkConfiguration) SetPublishedServerUriBySubnet(v []string)`

SetPublishedServerUriBySubnet sets PublishedServerUriBySubnet field to given value.

### HasPublishedServerUriBySubnet

`func (o *JellyfinNetworkConfiguration) HasPublishedServerUriBySubnet() bool`

HasPublishedServerUriBySubnet returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *JellyfinNetworkConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *JellyfinNetworkConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *JellyfinNetworkConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *JellyfinNetworkConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *JellyfinNetworkConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *JellyfinNetworkConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *JellyfinNetworkConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *JellyfinNetworkConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


