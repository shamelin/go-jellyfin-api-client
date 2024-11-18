# JellyfinGroupInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] 
**GroupName** | Pointer to **string** | Gets the group name. | [optional] 
**State** | Pointer to [**JellyfinJellyfinGroupStateType**](JellyfinGroupStateType.md) | Gets the group state. | [optional] 
**Participants** | Pointer to **[]string** | Gets the participants. | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Gets the date when this DTO has been created. | [optional] 

## Methods

### NewJellyfinGroupInfoDto

`func NewJellyfinGroupInfoDto() *JellyfinGroupInfoDto`

NewJellyfinGroupInfoDto instantiates a new JellyfinGroupInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinGroupInfoDtoWithDefaults

`func NewJellyfinGroupInfoDtoWithDefaults() *JellyfinGroupInfoDto`

NewJellyfinGroupInfoDtoWithDefaults instantiates a new JellyfinGroupInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *JellyfinGroupInfoDto) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *JellyfinGroupInfoDto) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *JellyfinGroupInfoDto) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *JellyfinGroupInfoDto) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *JellyfinGroupInfoDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *JellyfinGroupInfoDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *JellyfinGroupInfoDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *JellyfinGroupInfoDto) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetState

`func (o *JellyfinGroupInfoDto) GetState() JellyfinJellyfinGroupStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JellyfinGroupInfoDto) GetStateOk() (*JellyfinJellyfinGroupStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JellyfinGroupInfoDto) SetState(v JellyfinJellyfinGroupStateType)`

SetState sets State field to given value.

### HasState

`func (o *JellyfinGroupInfoDto) HasState() bool`

HasState returns a boolean if a field has been set.

### GetParticipants

`func (o *JellyfinGroupInfoDto) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *JellyfinGroupInfoDto) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *JellyfinGroupInfoDto) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *JellyfinGroupInfoDto) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *JellyfinGroupInfoDto) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *JellyfinGroupInfoDto) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *JellyfinGroupInfoDto) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *JellyfinGroupInfoDto) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


