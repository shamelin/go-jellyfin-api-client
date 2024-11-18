# ModelGroupInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] 
**GroupName** | Pointer to **string** | Gets the group name. | [optional] 
**State** | Pointer to [**ModelModelGroupStateType**](ModelGroupStateType.md) | Gets the group state. | [optional] 
**Participants** | Pointer to **[]string** | Gets the participants. | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Gets the date when this DTO has been created. | [optional] 

## Methods

### NewModelGroupInfoDto

`func NewModelGroupInfoDto() *ModelGroupInfoDto`

NewModelGroupInfoDto instantiates a new ModelGroupInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGroupInfoDtoWithDefaults

`func NewModelGroupInfoDtoWithDefaults() *ModelGroupInfoDto`

NewModelGroupInfoDtoWithDefaults instantiates a new ModelGroupInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ModelGroupInfoDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelGroupInfoDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelGroupInfoDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelGroupInfoDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *ModelGroupInfoDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelGroupInfoDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelGroupInfoDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ModelGroupInfoDto) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetState

`func (o *ModelGroupInfoDto) GetState() ModelModelGroupStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelGroupInfoDto) GetStateOk() (*ModelModelGroupStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelGroupInfoDto) SetState(v ModelModelGroupStateType)`

SetState sets State field to given value.

### HasState

`func (o *ModelGroupInfoDto) HasState() bool`

HasState returns a boolean if a field has been set.

### GetParticipants

`func (o *ModelGroupInfoDto) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ModelGroupInfoDto) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ModelGroupInfoDto) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ModelGroupInfoDto) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ModelGroupInfoDto) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ModelGroupInfoDto) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ModelGroupInfoDto) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ModelGroupInfoDto) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


