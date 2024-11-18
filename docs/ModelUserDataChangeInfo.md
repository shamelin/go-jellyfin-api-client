# ModelUserDataChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserDataList** | Pointer to [**[]ModelModelUserItemDataDto**](ModelModelUserItemDataDto.md) | Gets or sets the user data list. | [optional] 

## Methods

### NewModelUserDataChangeInfo

`func NewModelUserDataChangeInfo() *ModelUserDataChangeInfo`

NewModelUserDataChangeInfo instantiates a new ModelUserDataChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserDataChangeInfoWithDefaults

`func NewModelUserDataChangeInfoWithDefaults() *ModelUserDataChangeInfo`

NewModelUserDataChangeInfoWithDefaults instantiates a new ModelUserDataChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ModelUserDataChangeInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelUserDataChangeInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelUserDataChangeInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelUserDataChangeInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserDataList

`func (o *ModelUserDataChangeInfo) GetUserDataList() []ModelModelUserItemDataDto`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *ModelUserDataChangeInfo) GetUserDataListOk() (*[]ModelModelUserItemDataDto, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *ModelUserDataChangeInfo) SetUserDataList(v []ModelModelUserItemDataDto)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *ModelUserDataChangeInfo) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


