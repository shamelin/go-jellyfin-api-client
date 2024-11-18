# ModelAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Gets or sets the identifier. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device identifier. | [optional] 
**AppName** | Pointer to **NullableString** | Gets or sets the name of the application. | [optional] 
**AppVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**IsActive** | Pointer to **bool** | Gets or sets a value indicating whether this instance is active. | [optional] 
**DateCreated** | Pointer to **time.Time** | Gets or sets the date created. | [optional] 
**DateRevoked** | Pointer to **NullableTime** | Gets or sets the date revoked. | [optional] 
**DateLastActivity** | Pointer to **time.Time** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelAuthenticationInfo

`func NewModelAuthenticationInfo() *ModelAuthenticationInfo`

NewModelAuthenticationInfo instantiates a new ModelAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAuthenticationInfoWithDefaults

`func NewModelAuthenticationInfoWithDefaults() *ModelAuthenticationInfo`

NewModelAuthenticationInfoWithDefaults instantiates a new ModelAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelAuthenticationInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAuthenticationInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAuthenticationInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAuthenticationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessToken

`func (o *ModelAuthenticationInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ModelAuthenticationInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ModelAuthenticationInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ModelAuthenticationInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *ModelAuthenticationInfo) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *ModelAuthenticationInfo) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetDeviceId

`func (o *ModelAuthenticationInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelAuthenticationInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelAuthenticationInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelAuthenticationInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *ModelAuthenticationInfo) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *ModelAuthenticationInfo) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetAppName

`func (o *ModelAuthenticationInfo) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ModelAuthenticationInfo) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ModelAuthenticationInfo) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ModelAuthenticationInfo) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *ModelAuthenticationInfo) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *ModelAuthenticationInfo) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *ModelAuthenticationInfo) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *ModelAuthenticationInfo) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *ModelAuthenticationInfo) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *ModelAuthenticationInfo) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *ModelAuthenticationInfo) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *ModelAuthenticationInfo) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetDeviceName

`func (o *ModelAuthenticationInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ModelAuthenticationInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ModelAuthenticationInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ModelAuthenticationInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *ModelAuthenticationInfo) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *ModelAuthenticationInfo) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetUserId

`func (o *ModelAuthenticationInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelAuthenticationInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelAuthenticationInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelAuthenticationInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsActive

`func (o *ModelAuthenticationInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelAuthenticationInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelAuthenticationInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelAuthenticationInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *ModelAuthenticationInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelAuthenticationInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelAuthenticationInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelAuthenticationInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateRevoked

`func (o *ModelAuthenticationInfo) GetDateRevoked() time.Time`

GetDateRevoked returns the DateRevoked field if non-nil, zero value otherwise.

### GetDateRevokedOk

`func (o *ModelAuthenticationInfo) GetDateRevokedOk() (*time.Time, bool)`

GetDateRevokedOk returns a tuple with the DateRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRevoked

`func (o *ModelAuthenticationInfo) SetDateRevoked(v time.Time)`

SetDateRevoked sets DateRevoked field to given value.

### HasDateRevoked

`func (o *ModelAuthenticationInfo) HasDateRevoked() bool`

HasDateRevoked returns a boolean if a field has been set.

### SetDateRevokedNil

`func (o *ModelAuthenticationInfo) SetDateRevokedNil(b bool)`

 SetDateRevokedNil sets the value for DateRevoked to be an explicit nil

### UnsetDateRevoked
`func (o *ModelAuthenticationInfo) UnsetDateRevoked()`

UnsetDateRevoked ensures that no value is present for DateRevoked, not even an explicit nil
### GetDateLastActivity

`func (o *ModelAuthenticationInfo) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *ModelAuthenticationInfo) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *ModelAuthenticationInfo) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *ModelAuthenticationInfo) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### GetUserName

`func (o *ModelAuthenticationInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModelAuthenticationInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModelAuthenticationInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ModelAuthenticationInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ModelAuthenticationInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ModelAuthenticationInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


