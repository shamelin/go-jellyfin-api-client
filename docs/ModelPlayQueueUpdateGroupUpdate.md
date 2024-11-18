# ModelPlayQueueUpdateGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**ModelModelGroupUpdateType**](ModelGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**ModelModelPlayQueueUpdate**](ModelPlayQueueUpdate.md) | Gets the update data. | [optional] 

## Methods

### NewModelPlayQueueUpdateGroupUpdate

`func NewModelPlayQueueUpdateGroupUpdate() *ModelPlayQueueUpdateGroupUpdate`

NewModelPlayQueueUpdateGroupUpdate instantiates a new ModelPlayQueueUpdateGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlayQueueUpdateGroupUpdateWithDefaults

`func NewModelPlayQueueUpdateGroupUpdateWithDefaults() *ModelPlayQueueUpdateGroupUpdate`

NewModelPlayQueueUpdateGroupUpdateWithDefaults instantiates a new ModelPlayQueueUpdateGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelPlayQueueUpdateGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelPlayQueueUpdateGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelPlayQueueUpdateGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelPlayQueueUpdateGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *ModelPlayQueueUpdateGroupUpdate) GetType() ModelModelGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelPlayQueueUpdateGroupUpdate) GetTypeOk() (*ModelModelGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelPlayQueueUpdateGroupUpdate) SetType(v ModelModelGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelPlayQueueUpdateGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ModelPlayQueueUpdateGroupUpdate) GetData() ModelModelPlayQueueUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelPlayQueueUpdateGroupUpdate) GetDataOk() (*ModelModelPlayQueueUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelPlayQueueUpdateGroupUpdate) SetData(v ModelModelPlayQueueUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *ModelPlayQueueUpdateGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


