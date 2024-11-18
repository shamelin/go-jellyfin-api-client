# ModelDeviceInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**CustomName** | Pointer to **NullableString** | Gets or sets the custom name. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the identifier. | [optional] 
**LastUserName** | Pointer to **NullableString** | Gets or sets the last name of the user. | [optional] 
**AppName** | Pointer to **NullableString** | Gets or sets the name of the application. | [optional] 
**AppVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**LastUserId** | Pointer to **NullableString** | Gets or sets the last user identifier. | [optional] 
**DateLastActivity** | Pointer to **NullableTime** | Gets or sets the date last modified. | [optional] 
**Capabilities** | Pointer to [**ModelModelClientCapabilitiesDto**](ModelClientCapabilitiesDto.md) | Gets or sets the capabilities. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon URL. | [optional] 

## Methods

### NewModelDeviceInfoDto

`func NewModelDeviceInfoDto() *ModelDeviceInfoDto`

NewModelDeviceInfoDto instantiates a new ModelDeviceInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDeviceInfoDtoWithDefaults

`func NewModelDeviceInfoDtoWithDefaults() *ModelDeviceInfoDto`

NewModelDeviceInfoDtoWithDefaults instantiates a new ModelDeviceInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelDeviceInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelDeviceInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelDeviceInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelDeviceInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelDeviceInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelDeviceInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomName

`func (o *ModelDeviceInfoDto) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ModelDeviceInfoDto) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ModelDeviceInfoDto) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ModelDeviceInfoDto) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ModelDeviceInfoDto) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ModelDeviceInfoDto) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetAccessToken

`func (o *ModelDeviceInfoDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ModelDeviceInfoDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ModelDeviceInfoDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ModelDeviceInfoDto) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *ModelDeviceInfoDto) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *ModelDeviceInfoDto) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetId

`func (o *ModelDeviceInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDeviceInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDeviceInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDeviceInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelDeviceInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelDeviceInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastUserName

`func (o *ModelDeviceInfoDto) GetLastUserName() string`

GetLastUserName returns the LastUserName field if non-nil, zero value otherwise.

### GetLastUserNameOk

`func (o *ModelDeviceInfoDto) GetLastUserNameOk() (*string, bool)`

GetLastUserNameOk returns a tuple with the LastUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserName

`func (o *ModelDeviceInfoDto) SetLastUserName(v string)`

SetLastUserName sets LastUserName field to given value.

### HasLastUserName

`func (o *ModelDeviceInfoDto) HasLastUserName() bool`

HasLastUserName returns a boolean if a field has been set.

### SetLastUserNameNil

`func (o *ModelDeviceInfoDto) SetLastUserNameNil(b bool)`

 SetLastUserNameNil sets the value for LastUserName to be an explicit nil

### UnsetLastUserName
`func (o *ModelDeviceInfoDto) UnsetLastUserName()`

UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
### GetAppName

`func (o *ModelDeviceInfoDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelDeviceInfoDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelDeviceInfoDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelDeviceInfoDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *ModelDeviceInfoDto) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *ModelDeviceInfoDto) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *ModelDeviceInfoDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *ModelDeviceInfoDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *ModelDeviceInfoDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *ModelDeviceInfoDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *ModelDeviceInfoDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *ModelDeviceInfoDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetLastUserId

`func (o *ModelDeviceInfoDto) GetLastUserId() string`

GetLastUserId returns the LastUserId field if non-nil, zero value otherwise.

### GetLastUserIdOk

`func (o *ModelDeviceInfoDto) GetLastUserIdOk() (*string, bool)`

GetLastUserIdOk returns a tuple with the LastUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserId

`func (o *ModelDeviceInfoDto) SetLastUserId(v string)`

SetLastUserId sets LastUserId field to given value.

### HasLastUserId

`func (o *ModelDeviceInfoDto) HasLastUserId() bool`

HasLastUserId returns a boolean if a field has been set.

### SetLastUserIdNil

`func (o *ModelDeviceInfoDto) SetLastUserIdNil(b bool)`

 SetLastUserIdNil sets the value for LastUserId to be an explicit nil

### UnsetLastUserId
`func (o *ModelDeviceInfoDto) UnsetLastUserId()`

UnsetLastUserId ensures that no value is present for LastUserId, not even an explicit nil
### GetDateLastActivity

`func (o *ModelDeviceInfoDto) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *ModelDeviceInfoDto) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *ModelDeviceInfoDto) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *ModelDeviceInfoDto) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### SetDateLastActivityNil

`func (o *ModelDeviceInfoDto) SetDateLastActivityNil(b bool)`

 SetDateLastActivityNil sets the value for DateLastActivity to be an explicit nil

### UnsetDateLastActivity
`func (o *ModelDeviceInfoDto) UnsetDateLastActivity()`

UnsetDateLastActivity ensures that no value is present for DateLastActivity, not even an explicit nil
### GetCapabilities

`func (o *ModelDeviceInfoDto) GetCapabilities() ModelModelClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelDeviceInfoDto) GetCapabilitiesOk() (*ModelModelClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelDeviceInfoDto) SetCapabilities(v ModelModelClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelDeviceInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetIconUrl

`func (o *ModelDeviceInfoDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ModelDeviceInfoDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ModelDeviceInfoDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ModelDeviceInfoDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ModelDeviceInfoDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ModelDeviceInfoDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


