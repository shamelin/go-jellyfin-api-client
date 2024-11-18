# ModelGroupStateUpdateGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**ModelModelGroupUpdateType**](ModelGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**ModelModelGroupStateUpdate**](ModelGroupStateUpdate.md) | Gets the update data. | [optional] 

## Methods

### NewModelGroupStateUpdateGroupUpdate

`func NewModelGroupStateUpdateGroupUpdate() *ModelGroupStateUpdateGroupUpdate`

NewModelGroupStateUpdateGroupUpdate instantiates a new ModelGroupStateUpdateGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGroupStateUpdateGroupUpdateWithDefaults

`func NewModelGroupStateUpdateGroupUpdateWithDefaults() *ModelGroupStateUpdateGroupUpdate`

NewModelGroupStateUpdateGroupUpdateWithDefaults instantiates a new ModelGroupStateUpdateGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelGroupStateUpdateGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelGroupStateUpdateGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelGroupStateUpdateGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelGroupStateUpdateGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *ModelGroupStateUpdateGroupUpdate) GetType() ModelModelGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelGroupStateUpdateGroupUpdate) GetTypeOk() (*ModelModelGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelGroupStateUpdateGroupUpdate) SetType(v ModelModelGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelGroupStateUpdateGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ModelGroupStateUpdateGroupUpdate) GetData() ModelModelGroupStateUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelGroupStateUpdateGroupUpdate) GetDataOk() (*ModelModelGroupStateUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelGroupStateUpdateGroupUpdate) SetData(v ModelModelGroupStateUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *ModelGroupStateUpdateGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


