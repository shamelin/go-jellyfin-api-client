# JellyfinGroupStateUpdateGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**JellyfinJellyfinGroupUpdateType**](JellyfinGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**JellyfinJellyfinGroupStateUpdate**](JellyfinGroupStateUpdate.md) | Gets the update data. | [optional] 

## Methods

### NewJellyfinGroupStateUpdateGroupUpdate

`func NewJellyfinGroupStateUpdateGroupUpdate() *JellyfinGroupStateUpdateGroupUpdate`

NewJellyfinGroupStateUpdateGroupUpdate instantiates a new JellyfinGroupStateUpdateGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGroupStateUpdateGroupUpdateWithDefaults

`func NewJellyfinGroupStateUpdateGroupUpdateWithDefaults() *JellyfinGroupStateUpdateGroupUpdate`

NewJellyfinGroupStateUpdateGroupUpdateWithDefaults instantiates a new JellyfinGroupStateUpdateGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *JellyfinGroupStateUpdateGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *JellyfinGroupStateUpdateGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetType() JellyfinJellyfinGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetTypeOk() (*JellyfinJellyfinGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinGroupStateUpdateGroupUpdate) SetType(v JellyfinJellyfinGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinGroupStateUpdateGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetData() JellyfinJellyfinGroupStateUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinGroupStateUpdateGroupUpdate) GetDataOk() (*JellyfinJellyfinGroupStateUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinGroupStateUpdateGroupUpdate) SetData(v JellyfinJellyfinGroupStateUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinGroupStateUpdateGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

