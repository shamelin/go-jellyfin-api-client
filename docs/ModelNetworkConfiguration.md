# ModelNetworkConfiguration

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

### NewModelNetworkConfiguration

`func NewModelNetworkConfiguration() *ModelNetworkConfiguration`

NewModelNetworkConfiguration instantiates a new ModelNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNetworkConfigurationWithDefaults

`func NewModelNetworkConfigurationWithDefaults() *ModelNetworkConfiguration`

NewModelNetworkConfigurationWithDefaults instantiates a new ModelNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *ModelNetworkConfiguration) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ModelNetworkConfiguration) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ModelNetworkConfiguration) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ModelNetworkConfiguration) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetEnableHttps

`func (o *ModelNetworkConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *ModelNetworkConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *ModelNetworkConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *ModelNetworkConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetRequireHttps

`func (o *ModelNetworkConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *ModelNetworkConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *ModelNetworkConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *ModelNetworkConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetCertificatePath

`func (o *ModelNetworkConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *ModelNetworkConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *ModelNetworkConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *ModelNetworkConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *ModelNetworkConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *ModelNetworkConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *ModelNetworkConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *ModelNetworkConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetInternalHttpPort

`func (o *ModelNetworkConfiguration) GetInternalHttpPort() int32`

GetInternalHttpPort returns the InternalHttpPort field if non-nil, zero value otherwise.

### GetInternalHttpPortOk

`func (o *ModelNetworkConfiguration) GetInternalHttpPortOk() (*int32, bool)`

GetInternalHttpPortOk returns a tuple with the InternalHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpPort

`func (o *ModelNetworkConfiguration) SetInternalHttpPort(v int32)`

SetInternalHttpPort sets InternalHttpPort field to given value.

### HasInternalHttpPort

`func (o *ModelNetworkConfiguration) HasInternalHttpPort() bool`

HasInternalHttpPort returns a boolean if a field has been set.

### GetInternalHttpsPort

`func (o *ModelNetworkConfiguration) GetInternalHttpsPort() int32`

GetInternalHttpsPort returns the InternalHttpsPort field if non-nil, zero value otherwise.

### GetInternalHttpsPortOk

`func (o *ModelNetworkConfiguration) GetInternalHttpsPortOk() (*int32, bool)`

GetInternalHttpsPortOk returns a tuple with the InternalHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpsPort

`func (o *ModelNetworkConfiguration) SetInternalHttpsPort(v int32)`

SetInternalHttpsPort sets InternalHttpsPort field to given value.

### HasInternalHttpsPort

`func (o *ModelNetworkConfiguration) HasInternalHttpsPort() bool`

HasInternalHttpsPort returns a boolean if a field has been set.

### GetPublicHttpPort

`func (o *ModelNetworkConfiguration) GetPublicHttpPort() int32`

GetPublicHttpPort returns the PublicHttpPort field if non-nil, zero value otherwise.

### GetPublicHttpPortOk

`func (o *ModelNetworkConfiguration) GetPublicHttpPortOk() (*int32, bool)`

GetPublicHttpPortOk returns a tuple with the PublicHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpPort

`func (o *ModelNetworkConfiguration) SetPublicHttpPort(v int32)`

SetPublicHttpPort sets PublicHttpPort field to given value.

### HasPublicHttpPort

`func (o *ModelNetworkConfiguration) HasPublicHttpPort() bool`

HasPublicHttpPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *ModelNetworkConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *ModelNetworkConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *ModelNetworkConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *ModelNetworkConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetAutoDiscovery

`func (o *ModelNetworkConfiguration) GetAutoDiscovery() bool`

GetAutoDiscovery returns the AutoDiscovery field if non-nil, zero value otherwise.

### GetAutoDiscoveryOk

`func (o *ModelNetworkConfiguration) GetAutoDiscoveryOk() (*bool, bool)`

GetAutoDiscoveryOk returns a tuple with the AutoDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscovery

`func (o *ModelNetworkConfiguration) SetAutoDiscovery(v bool)`

SetAutoDiscovery sets AutoDiscovery field to given value.

### HasAutoDiscovery

`func (o *ModelNetworkConfiguration) HasAutoDiscovery() bool`

HasAutoDiscovery returns a boolean if a field has been set.

### GetEnableUPnP

`func (o *ModelNetworkConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *ModelNetworkConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *ModelNetworkConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *ModelNetworkConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetEnableIPv4

`func (o *ModelNetworkConfiguration) GetEnableIPv4() bool`

GetEnableIPv4 returns the EnableIPv4 field if non-nil, zero value otherwise.

### GetEnableIPv4Ok

`func (o *ModelNetworkConfiguration) GetEnableIPv4Ok() (*bool, bool)`

GetEnableIPv4Ok returns a tuple with the EnableIPv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv4

`func (o *ModelNetworkConfiguration) SetEnableIPv4(v bool)`

SetEnableIPv4 sets EnableIPv4 field to given value.

### HasEnableIPv4

`func (o *ModelNetworkConfiguration) HasEnableIPv4() bool`

HasEnableIPv4 returns a boolean if a field has been set.

### GetEnableIPv6

`func (o *ModelNetworkConfiguration) GetEnableIPv6() bool`

GetEnableIPv6 returns the EnableIPv6 field if non-nil, zero value otherwise.

### GetEnableIPv6Ok

`func (o *ModelNetworkConfiguration) GetEnableIPv6Ok() (*bool, bool)`

GetEnableIPv6Ok returns a tuple with the EnableIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv6

`func (o *ModelNetworkConfiguration) SetEnableIPv6(v bool)`

SetEnableIPv6 sets EnableIPv6 field to given value.

### HasEnableIPv6

`func (o *ModelNetworkConfiguration) HasEnableIPv6() bool`

HasEnableIPv6 returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *ModelNetworkConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *ModelNetworkConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *ModelNetworkConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *ModelNetworkConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *ModelNetworkConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *ModelNetworkConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *ModelNetworkConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *ModelNetworkConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *ModelNetworkConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *ModelNetworkConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *ModelNetworkConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *ModelNetworkConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetKnownProxies

`func (o *ModelNetworkConfiguration) GetKnownProxies() []string`

GetKnownProxies returns the KnownProxies field if non-nil, zero value otherwise.

### GetKnownProxiesOk

`func (o *ModelNetworkConfiguration) GetKnownProxiesOk() (*[]string, bool)`

GetKnownProxiesOk returns a tuple with the KnownProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownProxies

`func (o *ModelNetworkConfiguration) SetKnownProxies(v []string)`

SetKnownProxies sets KnownProxies field to given value.

### HasKnownProxies

`func (o *ModelNetworkConfiguration) HasKnownProxies() bool`

HasKnownProxies returns a boolean if a field has been set.

### GetIgnoreVirtualInterfaces

`func (o *ModelNetworkConfiguration) GetIgnoreVirtualInterfaces() bool`

GetIgnoreVirtualInterfaces returns the IgnoreVirtualInterfaces field if non-nil, zero value otherwise.

### GetIgnoreVirtualInterfacesOk

`func (o *ModelNetworkConfiguration) GetIgnoreVirtualInterfacesOk() (*bool, bool)`

GetIgnoreVirtualInterfacesOk returns a tuple with the IgnoreVirtualInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreVirtualInterfaces

`func (o *ModelNetworkConfiguration) SetIgnoreVirtualInterfaces(v bool)`

SetIgnoreVirtualInterfaces sets IgnoreVirtualInterfaces field to given value.

### HasIgnoreVirtualInterfaces

`func (o *ModelNetworkConfiguration) HasIgnoreVirtualInterfaces() bool`

HasIgnoreVirtualInterfaces returns a boolean if a field has been set.

### GetVirtualInterfaceNames

`func (o *ModelNetworkConfiguration) GetVirtualInterfaceNames() []string`

GetVirtualInterfaceNames returns the VirtualInterfaceNames field if non-nil, zero value otherwise.

### GetVirtualInterfaceNamesOk

`func (o *ModelNetworkConfiguration) GetVirtualInterfaceNamesOk() (*[]string, bool)`

GetVirtualInterfaceNamesOk returns a tuple with the VirtualInterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaceNames

`func (o *ModelNetworkConfiguration) SetVirtualInterfaceNames(v []string)`

SetVirtualInterfaceNames sets VirtualInterfaceNames field to given value.

### HasVirtualInterfaceNames

`func (o *ModelNetworkConfiguration) HasVirtualInterfaceNames() bool`

HasVirtualInterfaceNames returns a boolean if a field has been set.

### GetEnablePublishedServerUriByRequest

`func (o *ModelNetworkConfiguration) GetEnablePublishedServerUriByRequest() bool`

GetEnablePublishedServerUriByRequest returns the EnablePublishedServerUriByRequest field if non-nil, zero value otherwise.

### GetEnablePublishedServerUriByRequestOk

`func (o *ModelNetworkConfiguration) GetEnablePublishedServerUriByRequestOk() (*bool, bool)`

GetEnablePublishedServerUriByRequestOk returns a tuple with the EnablePublishedServerUriByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublishedServerUriByRequest

`func (o *ModelNetworkConfiguration) SetEnablePublishedServerUriByRequest(v bool)`

SetEnablePublishedServerUriByRequest sets EnablePublishedServerUriByRequest field to given value.

### HasEnablePublishedServerUriByRequest

`func (o *ModelNetworkConfiguration) HasEnablePublishedServerUriByRequest() bool`

HasEnablePublishedServerUriByRequest returns a boolean if a field has been set.

### GetPublishedServerUriBySubnet

`func (o *ModelNetworkConfiguration) GetPublishedServerUriBySubnet() []string`

GetPublishedServerUriBySubnet returns the PublishedServerUriBySubnet field if non-nil, zero value otherwise.

### GetPublishedServerUriBySubnetOk

`func (o *ModelNetworkConfiguration) GetPublishedServerUriBySubnetOk() (*[]string, bool)`

GetPublishedServerUriBySubnetOk returns a tuple with the PublishedServerUriBySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServerUriBySubnet

`func (o *ModelNetworkConfiguration) SetPublishedServerUriBySubnet(v []string)`

SetPublishedServerUriBySubnet sets PublishedServerUriBySubnet field to given value.

### HasPublishedServerUriBySubnet

`func (o *ModelNetworkConfiguration) HasPublishedServerUriBySubnet() bool`

HasPublishedServerUriBySubnet returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *ModelNetworkConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *ModelNetworkConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *ModelNetworkConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *ModelNetworkConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *ModelNetworkConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *ModelNetworkConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *ModelNetworkConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *ModelNetworkConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


