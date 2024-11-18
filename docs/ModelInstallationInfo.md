# ModelInstallationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Gets or sets the Id. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the version. | [optional] 
**Changelog** | Pointer to **NullableString** | Gets or sets the changelog for this version. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Gets or sets the source URL. | [optional] 
**Checksum** | Pointer to **NullableString** | Gets or sets a checksum for the binary. | [optional] 
**PackageInfo** | Pointer to [**NullableModelModelPackageInfo**](ModelPackageInfo.md) | Gets or sets package information for the installation. | [optional] 

## Methods

### NewModelInstallationInfo

`func NewModelInstallationInfo() *ModelInstallationInfo`

NewModelInstallationInfo instantiates a new ModelInstallationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInstallationInfoWithDefaults

`func NewModelInstallationInfoWithDefaults() *ModelInstallationInfo`

NewModelInstallationInfoWithDefaults instantiates a new ModelInstallationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *ModelInstallationInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelInstallationInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelInstallationInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelInstallationInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *ModelInstallationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelInstallationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelInstallationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelInstallationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelInstallationInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelInstallationInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *ModelInstallationInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelInstallationInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelInstallationInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelInstallationInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelInstallationInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelInstallationInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetChangelog

`func (o *ModelInstallationInfo) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *ModelInstallationInfo) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *ModelInstallationInfo) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *ModelInstallationInfo) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### SetChangelogNil

`func (o *ModelInstallationInfo) SetChangelogNil(b bool)`

 SetChangelogNil sets the value for Changelog to be an explicit nil

### UnsetChangelog
`func (o *ModelInstallationInfo) UnsetChangelog()`

UnsetChangelog ensures that no value is present for Changelog, not even an explicit nil
### GetSourceUrl

`func (o *ModelInstallationInfo) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ModelInstallationInfo) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ModelInstallationInfo) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ModelInstallationInfo) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *ModelInstallationInfo) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *ModelInstallationInfo) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetChecksum

`func (o *ModelInstallationInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ModelInstallationInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ModelInstallationInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ModelInstallationInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *ModelInstallationInfo) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *ModelInstallationInfo) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetPackageInfo

`func (o *ModelInstallationInfo) GetPackageInfo() ModelModelPackageInfo`

GetPackageInfo returns the PackageInfo field if non-nil, zero value otherwise.

### GetPackageInfoOk

`func (o *ModelInstallationInfo) GetPackageInfoOk() (*ModelModelPackageInfo, bool)`

GetPackageInfoOk returns a tuple with the PackageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageInfo

`func (o *ModelInstallationInfo) SetPackageInfo(v ModelModelPackageInfo)`

SetPackageInfo sets PackageInfo field to given value.

### HasPackageInfo

`func (o *ModelInstallationInfo) HasPackageInfo() bool`

HasPackageInfo returns a boolean if a field has been set.

### SetPackageInfoNil

`func (o *ModelInstallationInfo) SetPackageInfoNil(b bool)`

 SetPackageInfoNil sets the value for PackageInfo to be an explicit nil

### UnsetPackageInfo
`func (o *ModelInstallationInfo) UnsetPackageInfo()`

UnsetPackageInfo ensures that no value is present for PackageInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


