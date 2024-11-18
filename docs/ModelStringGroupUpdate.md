# ModelStringGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**ModelModelGroupUpdateType**](ModelGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to **string** | Gets the update data. | [optional] 

## Methods

### NewModelStringGroupUpdate

`func NewModelStringGroupUpdate() *ModelStringGroupUpdate`

NewModelStringGroupUpdate instantiates a new ModelStringGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelStringGroupUpdateWithDefaults

`func NewModelStringGroupUpdateWithDefaults() *ModelStringGroupUpdate`

NewModelStringGroupUpdateWithDefaults instantiates a new ModelStringGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelStringGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelStringGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelStringGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelStringGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *ModelStringGroupUpdate) GetType() ModelModelGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelStringGroupUpdate) GetTypeOk() (*ModelModelGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelStringGroupUpdate) SetType(v ModelModelGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelStringGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ModelStringGroupUpdate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelStringGroupUpdate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelStringGroupUpdate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ModelStringGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


