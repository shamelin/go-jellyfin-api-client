# JellyfinDeviceInfoDto

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
**Capabilities** | Pointer to [**JellyfinJellyfinClientCapabilitiesDto**](JellyfinClientCapabilitiesDto.md) | Gets or sets the capabilities. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon URL. | [optional] 

## Methods

### NewJellyfinDeviceInfoDto

`func NewJellyfinDeviceInfoDto() *JellyfinDeviceInfoDto`

NewJellyfinDeviceInfoDto instantiates a new JellyfinDeviceInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinDeviceInfoDtoWithDefaults

`func NewJellyfinDeviceInfoDtoWithDefaults() *JellyfinDeviceInfoDto`

NewJellyfinDeviceInfoDtoWithDefaults instantiates a new JellyfinDeviceInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinDeviceInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinDeviceInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinDeviceInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinDeviceInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinDeviceInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinDeviceInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomName

`func (o *JellyfinDeviceInfoDto) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *JellyfinDeviceInfoDto) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *JellyfinDeviceInfoDto) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *JellyfinDeviceInfoDto) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *JellyfinDeviceInfoDto) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *JellyfinDeviceInfoDto) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetAccessToken

`func (o *JellyfinDeviceInfoDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *JellyfinDeviceInfoDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *JellyfinDeviceInfoDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *JellyfinDeviceInfoDto) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *JellyfinDeviceInfoDto) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *JellyfinDeviceInfoDto) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetId

`func (o *JellyfinDeviceInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinDeviceInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinDeviceInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinDeviceInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinDeviceInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinDeviceInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastUserName

`func (o *JellyfinDeviceInfoDto) GetLastUserName() string`

GetLastUserName returns the LastUserName field if non-nil, zero value otherwise.

### GetLastUserNameOk

`func (o *JellyfinDeviceInfoDto) GetLastUserNameOk() (*string, bool)`

GetLastUserNameOk returns a tuple with the LastUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserName

`func (o *JellyfinDeviceInfoDto) SetLastUserName(v string)`

SetLastUserName sets LastUserName field to given value.

### HasLastUserName

`func (o *JellyfinDeviceInfoDto) HasLastUserName() bool`

HasLastUserName returns a boolean if a field has been set.

### SetLastUserNameNil

`func (o *JellyfinDeviceInfoDto) SetLastUserNameNil(b bool)`

 SetLastUserNameNil sets the value for LastUserName to be an explicit nil

### UnsetLastUserName
`func (o *JellyfinDeviceInfoDto) UnsetLastUserName()`

UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
### GetAppName

`func (o *JellyfinDeviceInfoDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *JellyfinDeviceInfoDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *JellyfinDeviceInfoDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *JellyfinDeviceInfoDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *JellyfinDeviceInfoDto) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *JellyfinDeviceInfoDto) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *JellyfinDeviceInfoDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *JellyfinDeviceInfoDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *JellyfinDeviceInfoDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *JellyfinDeviceInfoDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *JellyfinDeviceInfoDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *JellyfinDeviceInfoDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetLastUserId

`func (o *JellyfinDeviceInfoDto) GetLastUserId() string`

GetLastUserId returns the LastUserId field if non-nil, zero value otherwise.

### GetLastUserIdOk

`func (o *JellyfinDeviceInfoDto) GetLastUserIdOk() (*string, bool)`

GetLastUserIdOk returns a tuple with the LastUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserId

`func (o *JellyfinDeviceInfoDto) SetLastUserId(v string)`

SetLastUserId sets LastUserId field to given value.

### HasLastUserId

`func (o *JellyfinDeviceInfoDto) HasLastUserId() bool`

HasLastUserId returns a boolean if a field has been set.

### SetLastUserIdNil

`func (o *JellyfinDeviceInfoDto) SetLastUserIdNil(b bool)`

 SetLastUserIdNil sets the value for LastUserId to be an explicit nil

### UnsetLastUserId
`func (o *JellyfinDeviceInfoDto) UnsetLastUserId()`

UnsetLastUserId ensures that no value is present for LastUserId, not even an explicit nil
### GetDateLastActivity

`func (o *JellyfinDeviceInfoDto) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *JellyfinDeviceInfoDto) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *JellyfinDeviceInfoDto) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *JellyfinDeviceInfoDto) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### SetDateLastActivityNil

`func (o *JellyfinDeviceInfoDto) SetDateLastActivityNil(b bool)`

 SetDateLastActivityNil sets the value for DateLastActivity to be an explicit nil

### UnsetDateLastActivity
`func (o *JellyfinDeviceInfoDto) UnsetDateLastActivity()`

UnsetDateLastActivity ensures that no value is present for DateLastActivity, not even an explicit nil
### GetCapabilities

`func (o *JellyfinDeviceInfoDto) GetCapabilities() JellyfinJellyfinClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *JellyfinDeviceInfoDto) GetCapabilitiesOk() (*JellyfinJellyfinClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *JellyfinDeviceInfoDto) SetCapabilities(v JellyfinJellyfinClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *JellyfinDeviceInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetIconUrl

`func (o *JellyfinDeviceInfoDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *JellyfinDeviceInfoDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *JellyfinDeviceInfoDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *JellyfinDeviceInfoDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *JellyfinDeviceInfoDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *JellyfinDeviceInfoDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


