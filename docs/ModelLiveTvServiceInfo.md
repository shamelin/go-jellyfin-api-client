# ModelLiveTvServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**HomePageUrl** | Pointer to **NullableString** | Gets or sets the home page URL. | [optional] 
**Status** | Pointer to [**ModelModelLiveTvServiceStatus**](ModelLiveTvServiceStatus.md) | Gets or sets the status. | [optional] 
**StatusMessage** | Pointer to **NullableString** | Gets or sets the status message. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the version. | [optional] 
**HasUpdateAvailable** | Pointer to **bool** | Gets or sets a value indicating whether this instance has update available. | [optional] 
**IsVisible** | Pointer to **bool** | Gets or sets a value indicating whether this instance is visible. | [optional] 
**Tuners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelLiveTvServiceInfo

`func NewModelLiveTvServiceInfo() *ModelLiveTvServiceInfo`

NewModelLiveTvServiceInfo instantiates a new ModelLiveTvServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvServiceInfoWithDefaults

`func NewModelLiveTvServiceInfoWithDefaults() *ModelLiveTvServiceInfo`

NewModelLiveTvServiceInfoWithDefaults instantiates a new ModelLiveTvServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelLiveTvServiceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelLiveTvServiceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelLiveTvServiceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelLiveTvServiceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelLiveTvServiceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelLiveTvServiceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHomePageUrl

`func (o *ModelLiveTvServiceInfo) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ModelLiveTvServiceInfo) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ModelLiveTvServiceInfo) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *ModelLiveTvServiceInfo) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *ModelLiveTvServiceInfo) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *ModelLiveTvServiceInfo) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetStatus

`func (o *ModelLiveTvServiceInfo) GetStatus() ModelModelLiveTvServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelLiveTvServiceInfo) GetStatusOk() (*ModelModelLiveTvServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelLiveTvServiceInfo) SetStatus(v ModelModelLiveTvServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelLiveTvServiceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ModelLiveTvServiceInfo) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ModelLiveTvServiceInfo) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ModelLiveTvServiceInfo) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ModelLiveTvServiceInfo) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ModelLiveTvServiceInfo) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ModelLiveTvServiceInfo) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetVersion

`func (o *ModelLiveTvServiceInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelLiveTvServiceInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelLiveTvServiceInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelLiveTvServiceInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ModelLiveTvServiceInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ModelLiveTvServiceInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetHasUpdateAvailable

`func (o *ModelLiveTvServiceInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *ModelLiveTvServiceInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *ModelLiveTvServiceInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *ModelLiveTvServiceInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetIsVisible

`func (o *ModelLiveTvServiceInfo) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *ModelLiveTvServiceInfo) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *ModelLiveTvServiceInfo) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *ModelLiveTvServiceInfo) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetTuners

`func (o *ModelLiveTvServiceInfo) GetTuners() []string`

GetTuners returns the Tuners field if non-nil, zero value otherwise.

### GetTunersOk

`func (o *ModelLiveTvServiceInfo) GetTunersOk() (*[]string, bool)`

GetTunersOk returns a tuple with the Tuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuners

`func (o *ModelLiveTvServiceInfo) SetTuners(v []string)`

SetTuners sets Tuners field to given value.

### HasTuners

`func (o *ModelLiveTvServiceInfo) HasTuners() bool`

HasTuners returns a boolean if a field has been set.

### SetTunersNil

`func (o *ModelLiveTvServiceInfo) SetTunersNil(b bool)`

 SetTunersNil sets the value for Tuners to be an explicit nil

### UnsetTuners
`func (o *ModelLiveTvServiceInfo) UnsetTuners()`

UnsetTuners ensures that no value is present for Tuners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


