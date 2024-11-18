# ModelGroupInfoDtoGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**ModelModelGroupUpdateType**](ModelGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**ModelModelGroupInfoDto**](ModelGroupInfoDto.md) | Gets the update data. | [optional] 

## Methods

### NewModelGroupInfoDtoGroupUpdate

`func NewModelGroupInfoDtoGroupUpdate() *ModelGroupInfoDtoGroupUpdate`

NewModelGroupInfoDtoGroupUpdate instantiates a new ModelGroupInfoDtoGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGroupInfoDtoGroupUpdateWithDefaults

`func NewModelGroupInfoDtoGroupUpdateWithDefaults() *ModelGroupInfoDtoGroupUpdate`

NewModelGroupInfoDtoGroupUpdateWithDefaults instantiates a new ModelGroupInfoDtoGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelGroupInfoDtoGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelGroupInfoDtoGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelGroupInfoDtoGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelGroupInfoDtoGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *ModelGroupInfoDtoGroupUpdate) GetType() ModelModelGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelGroupInfoDtoGroupUpdate) GetTypeOk() (*ModelModelGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelGroupInfoDtoGroupUpdate) SetType(v ModelModelGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelGroupInfoDtoGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ModelGroupInfoDtoGroupUpdate) GetData() ModelModelGroupInfoDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelGroupInfoDtoGroupUpdate) GetDataOk() (*ModelModelGroupInfoDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelGroupInfoDtoGroupUpdate) SetData(v ModelModelGroupInfoDto)`

SetData sets Data field to given value.

### HasData

`func (o *ModelGroupInfoDtoGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


