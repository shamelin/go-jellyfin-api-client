# ModelSessionUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the name of the user. | [optional] 

## Methods

### NewModelSessionUserInfo

`func NewModelSessionUserInfo() *ModelSessionUserInfo`

NewModelSessionUserInfo instantiates a new ModelSessionUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSessionUserInfoWithDefaults

`func NewModelSessionUserInfoWithDefaults() *ModelSessionUserInfo`

NewModelSessionUserInfoWithDefaults instantiates a new ModelSessionUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ModelSessionUserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSessionUserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSessionUserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSessionUserInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ModelSessionUserInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModelSessionUserInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModelSessionUserInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ModelSessionUserInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ModelSessionUserInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ModelSessionUserInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


