# JellyfinUserDataChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserDataList** | Pointer to [**[]JellyfinJellyfinUserItemDataDto**](JellyfinJellyfinUserItemDataDto.md) | Gets or sets the user data list. | [optional] 

## Methods

### NewJellyfinUserDataChangeInfo

`func NewJellyfinUserDataChangeInfo() *JellyfinUserDataChangeInfo`

NewJellyfinUserDataChangeInfo instantiates a new JellyfinUserDataChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUserDataChangeInfoWithDefaults

`func NewJellyfinUserDataChangeInfoWithDefaults() *JellyfinUserDataChangeInfo`

NewJellyfinUserDataChangeInfoWithDefaults instantiates a new JellyfinUserDataChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *JellyfinUserDataChangeInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinUserDataChangeInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinUserDataChangeInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinUserDataChangeInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserDataList

`func (o *JellyfinUserDataChangeInfo) GetUserDataList() []JellyfinJellyfinUserItemDataDto`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *JellyfinUserDataChangeInfo) GetUserDataListOk() (*[]JellyfinJellyfinUserItemDataDto, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *JellyfinUserDataChangeInfo) SetUserDataList(v []JellyfinJellyfinUserItemDataDto)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *JellyfinUserDataChangeInfo) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


