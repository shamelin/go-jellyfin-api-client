# JellyfinGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**JellyfinJellyfinGroupUpdateType**](JellyfinGroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**JellyfinJellyfinPlayQueueUpdate**](JellyfinPlayQueueUpdate.md) | Gets the update data. | [optional] 

## Methods

### NewJellyfinGroupUpdate

`func NewJellyfinGroupUpdate() *JellyfinGroupUpdate`

NewJellyfinGroupUpdate instantiates a new JellyfinGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGroupUpdateWithDefaults

`func NewJellyfinGroupUpdateWithDefaults() *JellyfinGroupUpdate`

NewJellyfinGroupUpdateWithDefaults instantiates a new JellyfinGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *JellyfinGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *JellyfinGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *JellyfinGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *JellyfinGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *JellyfinGroupUpdate) GetType() JellyfinJellyfinGroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinGroupUpdate) GetTypeOk() (*JellyfinJellyfinGroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinGroupUpdate) SetType(v JellyfinJellyfinGroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *JellyfinGroupUpdate) GetData() JellyfinJellyfinPlayQueueUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JellyfinGroupUpdate) GetDataOk() (*JellyfinJellyfinPlayQueueUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JellyfinGroupUpdate) SetData(v JellyfinJellyfinPlayQueueUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *JellyfinGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


