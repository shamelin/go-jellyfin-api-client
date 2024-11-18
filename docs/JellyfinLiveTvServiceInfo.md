# JellyfinLiveTvServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**HomePageUrl** | Pointer to **NullableString** | Gets or sets the home page URL. | [optional] 
**Status** | Pointer to [**JellyfinJellyfinLiveTvServiceStatus**](JellyfinLiveTvServiceStatus.md) | Gets or sets the status. | [optional] 
**StatusMessage** | Pointer to **NullableString** | Gets or sets the status message. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the version. | [optional] 
**HasUpdateAvailable** | Pointer to **bool** | Gets or sets a value indicating whether this instance has update available. | [optional] 
**IsVisible** | Pointer to **bool** | Gets or sets a value indicating whether this instance is visible. | [optional] 
**Tuners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJellyfinLiveTvServiceInfo

`func NewJellyfinLiveTvServiceInfo() *JellyfinLiveTvServiceInfo`

NewJellyfinLiveTvServiceInfo instantiates a new JellyfinLiveTvServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinLiveTvServiceInfoWithDefaults

`func NewJellyfinLiveTvServiceInfoWithDefaults() *JellyfinLiveTvServiceInfo`

NewJellyfinLiveTvServiceInfoWithDefaults instantiates a new JellyfinLiveTvServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinLiveTvServiceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinLiveTvServiceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinLiveTvServiceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinLiveTvServiceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinLiveTvServiceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinLiveTvServiceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHomePageUrl

`func (o *JellyfinLiveTvServiceInfo) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *JellyfinLiveTvServiceInfo) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *JellyfinLiveTvServiceInfo) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *JellyfinLiveTvServiceInfo) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *JellyfinLiveTvServiceInfo) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *JellyfinLiveTvServiceInfo) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetStatus

`func (o *JellyfinLiveTvServiceInfo) GetStatus() JellyfinJellyfinLiveTvServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JellyfinLiveTvServiceInfo) GetStatusOk() (*JellyfinJellyfinLiveTvServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JellyfinLiveTvServiceInfo) SetStatus(v JellyfinJellyfinLiveTvServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JellyfinLiveTvServiceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *JellyfinLiveTvServiceInfo) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *JellyfinLiveTvServiceInfo) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *JellyfinLiveTvServiceInfo) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *JellyfinLiveTvServiceInfo) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *JellyfinLiveTvServiceInfo) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *JellyfinLiveTvServiceInfo) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetVersion

`func (o *JellyfinLiveTvServiceInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JellyfinLiveTvServiceInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JellyfinLiveTvServiceInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JellyfinLiveTvServiceInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *JellyfinLiveTvServiceInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *JellyfinLiveTvServiceInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetHasUpdateAvailable

`func (o *JellyfinLiveTvServiceInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *JellyfinLiveTvServiceInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *JellyfinLiveTvServiceInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *JellyfinLiveTvServiceInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetIsVisible

`func (o *JellyfinLiveTvServiceInfo) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *JellyfinLiveTvServiceInfo) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *JellyfinLiveTvServiceInfo) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *JellyfinLiveTvServiceInfo) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetTuners

`func (o *JellyfinLiveTvServiceInfo) GetTuners() []string`

GetTuners returns the Tuners field if non-nil, zero value otherwise.

### GetTunersOk

`func (o *JellyfinLiveTvServiceInfo) GetTunersOk() (*[]string, bool)`

GetTunersOk returns a tuple with the Tuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuners

`func (o *JellyfinLiveTvServiceInfo) SetTuners(v []string)`

SetTuners sets Tuners field to given value.

### HasTuners

`func (o *JellyfinLiveTvServiceInfo) HasTuners() bool`

HasTuners returns a boolean if a field has been set.

### SetTunersNil

`func (o *JellyfinLiveTvServiceInfo) SetTunersNil(b bool)`

 SetTunersNil sets the value for Tuners to be an explicit nil

### UnsetTuners
`func (o *JellyfinLiveTvServiceInfo) UnsetTuners()`

UnsetTuners ensures that no value is present for Tuners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


