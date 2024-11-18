# ModelAuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**NullableModelModelUserDto**](ModelUserDto.md) | Class UserDto. | [optional] 
**SessionInfo** | Pointer to [**NullableModelModelSessionInfoDto**](ModelSessionInfoDto.md) | Session info DTO. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 

## Methods

### NewModelAuthenticationResult

`func NewModelAuthenticationResult() *ModelAuthenticationResult`

NewModelAuthenticationResult instantiates a new ModelAuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAuthenticationResultWithDefaults

`func NewModelAuthenticationResultWithDefaults() *ModelAuthenticationResult`

NewModelAuthenticationResultWithDefaults instantiates a new ModelAuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ModelAuthenticationResult) GetUser() ModelModelUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelAuthenticationResult) GetUserOk() (*ModelModelUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelAuthenticationResult) SetUser(v ModelModelUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelAuthenticationResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ModelAuthenticationResult) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ModelAuthenticationResult) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSessionInfo

`func (o *ModelAuthenticationResult) GetSessionInfo() ModelModelSessionInfoDto`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *ModelAuthenticationResult) GetSessionInfoOk() (*ModelModelSessionInfoDto, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *ModelAuthenticationResult) SetSessionInfo(v ModelModelSessionInfoDto)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *ModelAuthenticationResult) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### SetSessionInfoNil

`func (o *ModelAuthenticationResult) SetSessionInfoNil(b bool)`

 SetSessionInfoNil sets the value for SessionInfo to be an explicit nil

### UnsetSessionInfo
`func (o *ModelAuthenticationResult) UnsetSessionInfo()`

UnsetSessionInfo ensures that no value is present for SessionInfo, not even an explicit nil
### GetAccessToken

`func (o *ModelAuthenticationResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ModelAuthenticationResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ModelAuthenticationResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ModelAuthenticationResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *ModelAuthenticationResult) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *ModelAuthenticationResult) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetServerId

`func (o *ModelAuthenticationResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelAuthenticationResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelAuthenticationResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelAuthenticationResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ModelAuthenticationResult) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ModelAuthenticationResult) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


