# JellyfinAuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**NullableJellyfinJellyfinUserDto**](JellyfinUserDto.md) | Class UserDto. | [optional] 
**SessionInfo** | Pointer to [**NullableJellyfinJellyfinSessionInfoDto**](JellyfinSessionInfoDto.md) | Session info DTO. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 

## Methods

### NewJellyfinAuthenticationResult

`func NewJellyfinAuthenticationResult() *JellyfinAuthenticationResult`

NewJellyfinAuthenticationResult instantiates a new JellyfinAuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinAuthenticationResultWithDefaults

`func NewJellyfinAuthenticationResultWithDefaults() *JellyfinAuthenticationResult`

NewJellyfinAuthenticationResultWithDefaults instantiates a new JellyfinAuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *JellyfinAuthenticationResult) GetUser() JellyfinJellyfinUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JellyfinAuthenticationResult) GetUserOk() (*JellyfinJellyfinUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JellyfinAuthenticationResult) SetUser(v JellyfinJellyfinUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *JellyfinAuthenticationResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *JellyfinAuthenticationResult) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *JellyfinAuthenticationResult) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSessionInfo

`func (o *JellyfinAuthenticationResult) GetSessionInfo() JellyfinJellyfinSessionInfoDto`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *JellyfinAuthenticationResult) GetSessionInfoOk() (*JellyfinJellyfinSessionInfoDto, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *JellyfinAuthenticationResult) SetSessionInfo(v JellyfinJellyfinSessionInfoDto)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *JellyfinAuthenticationResult) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### SetSessionInfoNil

`func (o *JellyfinAuthenticationResult) SetSessionInfoNil(b bool)`

 SetSessionInfoNil sets the value for SessionInfo to be an explicit nil

### UnsetSessionInfo
`func (o *JellyfinAuthenticationResult) UnsetSessionInfo()`

UnsetSessionInfo ensures that no value is present for SessionInfo, not even an explicit nil
### GetAccessToken

`func (o *JellyfinAuthenticationResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *JellyfinAuthenticationResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *JellyfinAuthenticationResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *JellyfinAuthenticationResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *JellyfinAuthenticationResult) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *JellyfinAuthenticationResult) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetServerId

`func (o *JellyfinAuthenticationResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *JellyfinAuthenticationResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *JellyfinAuthenticationResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *JellyfinAuthenticationResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *JellyfinAuthenticationResult) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *JellyfinAuthenticationResult) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


