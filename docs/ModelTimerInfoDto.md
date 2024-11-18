# ModelTimerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the Id of the recording. | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server identifier. | [optional] 
**ExternalId** | Pointer to **NullableString** | Gets or sets the external identifier. | [optional] 
**ChannelId** | Pointer to **string** | Gets or sets the channel id of the recording. | [optional] 
**ExternalChannelId** | Pointer to **NullableString** | Gets or sets the external channel identifier. | [optional] 
**ChannelName** | Pointer to **NullableString** | Gets or sets the channel name of the recording. | [optional] 
**ChannelPrimaryImageTag** | Pointer to **NullableString** |  | [optional] 
**ProgramId** | Pointer to **NullableString** | Gets or sets the program identifier. | [optional] 
**ExternalProgramId** | Pointer to **NullableString** | Gets or sets the external program identifier. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name of the recording. | [optional] 
**Overview** | Pointer to **NullableString** | Gets or sets the description of the recording. | [optional] 
**StartDate** | Pointer to **time.Time** | Gets or sets the start date of the recording, in UTC. | [optional] 
**EndDate** | Pointer to **time.Time** | Gets or sets the end date of the recording, in UTC. | [optional] 
**ServiceName** | Pointer to **NullableString** | Gets or sets the name of the service. | [optional] 
**Priority** | Pointer to **int32** | Gets or sets the priority. | [optional] 
**PrePaddingSeconds** | Pointer to **int32** | Gets or sets the pre padding seconds. | [optional] 
**PostPaddingSeconds** | Pointer to **int32** | Gets or sets the post padding seconds. | [optional] 
**IsPrePaddingRequired** | Pointer to **bool** | Gets or sets a value indicating whether this instance is pre padding required. | [optional] 
**ParentBackdropItemId** | Pointer to **NullableString** | Gets or sets the Id of the Parent that has a backdrop if the item does not have one. | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** | Gets or sets the parent backdrop image tags. | [optional] 
**IsPostPaddingRequired** | Pointer to **bool** | Gets or sets a value indicating whether this instance is post padding required. | [optional] 
**KeepUntil** | Pointer to [**ModelModelKeepUntil**](ModelKeepUntil.md) |  | [optional] 
**Status** | Pointer to [**ModelModelRecordingStatus**](ModelRecordingStatus.md) | Gets or sets the status. | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets the series timer identifier. | [optional] 
**ExternalSeriesTimerId** | Pointer to **NullableString** | Gets or sets the external series timer identifier. | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** | Gets or sets the run time ticks. | [optional] 
**ProgramInfo** | Pointer to [**NullableModelModelBaseItemDto**](ModelBaseItemDto.md) | Gets or sets the program information. | [optional] 

## Methods

### NewModelTimerInfoDto

`func NewModelTimerInfoDto() *ModelTimerInfoDto`

NewModelTimerInfoDto instantiates a new ModelTimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTimerInfoDtoWithDefaults

`func NewModelTimerInfoDtoWithDefaults() *ModelTimerInfoDto`

NewModelTimerInfoDtoWithDefaults instantiates a new ModelTimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelTimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelTimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelTimerInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelTimerInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *ModelTimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelTimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelTimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelTimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelTimerInfoDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelTimerInfoDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetServerId

`func (o *ModelTimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelTimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelTimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelTimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ModelTimerInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ModelTimerInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetExternalId

`func (o *ModelTimerInfoDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ModelTimerInfoDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ModelTimerInfoDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ModelTimerInfoDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ModelTimerInfoDto) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ModelTimerInfoDto) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetChannelId

`func (o *ModelTimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelTimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelTimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelTimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetExternalChannelId

`func (o *ModelTimerInfoDto) GetExternalChannelId() string`

GetExternalChannelId returns the ExternalChannelId field if non-nil, zero value otherwise.

### GetExternalChannelIdOk

`func (o *ModelTimerInfoDto) GetExternalChannelIdOk() (*string, bool)`

GetExternalChannelIdOk returns a tuple with the ExternalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelId

`func (o *ModelTimerInfoDto) SetExternalChannelId(v string)`

SetExternalChannelId sets ExternalChannelId field to given value.

### HasExternalChannelId

`func (o *ModelTimerInfoDto) HasExternalChannelId() bool`

HasExternalChannelId returns a boolean if a field has been set.

### SetExternalChannelIdNil

`func (o *ModelTimerInfoDto) SetExternalChannelIdNil(b bool)`

 SetExternalChannelIdNil sets the value for ExternalChannelId to be an explicit nil

### UnsetExternalChannelId
`func (o *ModelTimerInfoDto) UnsetExternalChannelId()`

UnsetExternalChannelId ensures that no value is present for ExternalChannelId, not even an explicit nil
### GetChannelName

`func (o *ModelTimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ModelTimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ModelTimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ModelTimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *ModelTimerInfoDto) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ModelTimerInfoDto) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *ModelTimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *ModelTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *ModelTimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *ModelTimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *ModelTimerInfoDto) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *ModelTimerInfoDto) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetProgramId

`func (o *ModelTimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ModelTimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ModelTimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ModelTimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *ModelTimerInfoDto) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *ModelTimerInfoDto) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetExternalProgramId

`func (o *ModelTimerInfoDto) GetExternalProgramId() string`

GetExternalProgramId returns the ExternalProgramId field if non-nil, zero value otherwise.

### GetExternalProgramIdOk

`func (o *ModelTimerInfoDto) GetExternalProgramIdOk() (*string, bool)`

GetExternalProgramIdOk returns a tuple with the ExternalProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProgramId

`func (o *ModelTimerInfoDto) SetExternalProgramId(v string)`

SetExternalProgramId sets ExternalProgramId field to given value.

### HasExternalProgramId

`func (o *ModelTimerInfoDto) HasExternalProgramId() bool`

HasExternalProgramId returns a boolean if a field has been set.

### SetExternalProgramIdNil

`func (o *ModelTimerInfoDto) SetExternalProgramIdNil(b bool)`

 SetExternalProgramIdNil sets the value for ExternalProgramId to be an explicit nil

### UnsetExternalProgramId
`func (o *ModelTimerInfoDto) UnsetExternalProgramId()`

UnsetExternalProgramId ensures that no value is present for ExternalProgramId, not even an explicit nil
### GetName

`func (o *ModelTimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelTimerInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelTimerInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOverview

`func (o *ModelTimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelTimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelTimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelTimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ModelTimerInfoDto) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ModelTimerInfoDto) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetStartDate

`func (o *ModelTimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelTimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelTimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelTimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelTimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelTimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelTimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelTimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetServiceName

`func (o *ModelTimerInfoDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ModelTimerInfoDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ModelTimerInfoDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ModelTimerInfoDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ModelTimerInfoDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ModelTimerInfoDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetPriority

`func (o *ModelTimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelTimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelTimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelTimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *ModelTimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelTimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelTimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelTimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelTimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelTimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *ModelTimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *ModelTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *ModelTimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *ModelTimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *ModelTimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *ModelTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *ModelTimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *ModelTimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *ModelTimerInfoDto) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *ModelTimerInfoDto) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *ModelTimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *ModelTimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *ModelTimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *ModelTimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *ModelTimerInfoDto) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *ModelTimerInfoDto) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetIsPostPaddingRequired

`func (o *ModelTimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *ModelTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *ModelTimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *ModelTimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *ModelTimerInfoDto) GetKeepUntil() ModelModelKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *ModelTimerInfoDto) GetKeepUntilOk() (*ModelModelKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *ModelTimerInfoDto) SetKeepUntil(v ModelModelKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *ModelTimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetStatus

`func (o *ModelTimerInfoDto) GetStatus() ModelModelRecordingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelTimerInfoDto) GetStatusOk() (*ModelModelRecordingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelTimerInfoDto) SetStatus(v ModelModelRecordingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelTimerInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeriesTimerId

`func (o *ModelTimerInfoDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *ModelTimerInfoDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *ModelTimerInfoDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *ModelTimerInfoDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *ModelTimerInfoDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *ModelTimerInfoDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetExternalSeriesTimerId

`func (o *ModelTimerInfoDto) GetExternalSeriesTimerId() string`

GetExternalSeriesTimerId returns the ExternalSeriesTimerId field if non-nil, zero value otherwise.

### GetExternalSeriesTimerIdOk

`func (o *ModelTimerInfoDto) GetExternalSeriesTimerIdOk() (*string, bool)`

GetExternalSeriesTimerIdOk returns a tuple with the ExternalSeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSeriesTimerId

`func (o *ModelTimerInfoDto) SetExternalSeriesTimerId(v string)`

SetExternalSeriesTimerId sets ExternalSeriesTimerId field to given value.

### HasExternalSeriesTimerId

`func (o *ModelTimerInfoDto) HasExternalSeriesTimerId() bool`

HasExternalSeriesTimerId returns a boolean if a field has been set.

### SetExternalSeriesTimerIdNil

`func (o *ModelTimerInfoDto) SetExternalSeriesTimerIdNil(b bool)`

 SetExternalSeriesTimerIdNil sets the value for ExternalSeriesTimerId to be an explicit nil

### UnsetExternalSeriesTimerId
`func (o *ModelTimerInfoDto) UnsetExternalSeriesTimerId()`

UnsetExternalSeriesTimerId ensures that no value is present for ExternalSeriesTimerId, not even an explicit nil
### GetRunTimeTicks

`func (o *ModelTimerInfoDto) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelTimerInfoDto) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelTimerInfoDto) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelTimerInfoDto) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelTimerInfoDto) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelTimerInfoDto) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetProgramInfo

`func (o *ModelTimerInfoDto) GetProgramInfo() ModelModelBaseItemDto`

GetProgramInfo returns the ProgramInfo field if non-nil, zero value otherwise.

### GetProgramInfoOk

`func (o *ModelTimerInfoDto) GetProgramInfoOk() (*ModelModelBaseItemDto, bool)`

GetProgramInfoOk returns a tuple with the ProgramInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramInfo

`func (o *ModelTimerInfoDto) SetProgramInfo(v ModelModelBaseItemDto)`

SetProgramInfo sets ProgramInfo field to given value.

### HasProgramInfo

`func (o *ModelTimerInfoDto) HasProgramInfo() bool`

HasProgramInfo returns a boolean if a field has been set.

### SetProgramInfoNil

`func (o *ModelTimerInfoDto) SetProgramInfoNil(b bool)`

 SetProgramInfoNil sets the value for ProgramInfo to be an explicit nil

### UnsetProgramInfo
`func (o *ModelTimerInfoDto) UnsetProgramInfo()`

UnsetProgramInfo ensures that no value is present for ProgramInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


