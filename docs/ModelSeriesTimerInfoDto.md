# ModelSeriesTimerInfoDto

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
**RecordAnyTime** | Pointer to **bool** | Gets or sets a value indicating whether [record any time]. | [optional] 
**SkipEpisodesInLibrary** | Pointer to **bool** |  | [optional] 
**RecordAnyChannel** | Pointer to **bool** | Gets or sets a value indicating whether [record any channel]. | [optional] 
**KeepUpTo** | Pointer to **int32** |  | [optional] 
**RecordNewOnly** | Pointer to **bool** | Gets or sets a value indicating whether [record new only]. | [optional] 
**Days** | Pointer to [**[]ModelModelDayOfWeek**](ModelModelDayOfWeek.md) | Gets or sets the days. | [optional] 
**DayPattern** | Pointer to [**NullableModelModelDayPattern**](ModelDayPattern.md) | Gets or sets the day pattern. | [optional] 
**ImageTags** | Pointer to **map[string]string** | Gets or sets the image tags. | [optional] 
**ParentThumbItemId** | Pointer to **NullableString** | Gets or sets the parent thumb item id. | [optional] 
**ParentThumbImageTag** | Pointer to **NullableString** | Gets or sets the parent thumb image tag. | [optional] 
**ParentPrimaryImageItemId** | Pointer to **NullableString** | Gets or sets the parent primary image item identifier. | [optional] 
**ParentPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the parent primary image tag. | [optional] 

## Methods

### NewModelSeriesTimerInfoDto

`func NewModelSeriesTimerInfoDto() *ModelSeriesTimerInfoDto`

NewModelSeriesTimerInfoDto instantiates a new ModelSeriesTimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSeriesTimerInfoDtoWithDefaults

`func NewModelSeriesTimerInfoDtoWithDefaults() *ModelSeriesTimerInfoDto`

NewModelSeriesTimerInfoDtoWithDefaults instantiates a new ModelSeriesTimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelSeriesTimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSeriesTimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSeriesTimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSeriesTimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ModelSeriesTimerInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelSeriesTimerInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *ModelSeriesTimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelSeriesTimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelSeriesTimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelSeriesTimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ModelSeriesTimerInfoDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ModelSeriesTimerInfoDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetServerId

`func (o *ModelSeriesTimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelSeriesTimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelSeriesTimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelSeriesTimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ModelSeriesTimerInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ModelSeriesTimerInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetExternalId

`func (o *ModelSeriesTimerInfoDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ModelSeriesTimerInfoDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ModelSeriesTimerInfoDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ModelSeriesTimerInfoDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ModelSeriesTimerInfoDto) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ModelSeriesTimerInfoDto) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetChannelId

`func (o *ModelSeriesTimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelSeriesTimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelSeriesTimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelSeriesTimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetExternalChannelId

`func (o *ModelSeriesTimerInfoDto) GetExternalChannelId() string`

GetExternalChannelId returns the ExternalChannelId field if non-nil, zero value otherwise.

### GetExternalChannelIdOk

`func (o *ModelSeriesTimerInfoDto) GetExternalChannelIdOk() (*string, bool)`

GetExternalChannelIdOk returns a tuple with the ExternalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelId

`func (o *ModelSeriesTimerInfoDto) SetExternalChannelId(v string)`

SetExternalChannelId sets ExternalChannelId field to given value.

### HasExternalChannelId

`func (o *ModelSeriesTimerInfoDto) HasExternalChannelId() bool`

HasExternalChannelId returns a boolean if a field has been set.

### SetExternalChannelIdNil

`func (o *ModelSeriesTimerInfoDto) SetExternalChannelIdNil(b bool)`

 SetExternalChannelIdNil sets the value for ExternalChannelId to be an explicit nil

### UnsetExternalChannelId
`func (o *ModelSeriesTimerInfoDto) UnsetExternalChannelId()`

UnsetExternalChannelId ensures that no value is present for ExternalChannelId, not even an explicit nil
### GetChannelName

`func (o *ModelSeriesTimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ModelSeriesTimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ModelSeriesTimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ModelSeriesTimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *ModelSeriesTimerInfoDto) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ModelSeriesTimerInfoDto) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *ModelSeriesTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *ModelSeriesTimerInfoDto) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *ModelSeriesTimerInfoDto) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetProgramId

`func (o *ModelSeriesTimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ModelSeriesTimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ModelSeriesTimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ModelSeriesTimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *ModelSeriesTimerInfoDto) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *ModelSeriesTimerInfoDto) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetExternalProgramId

`func (o *ModelSeriesTimerInfoDto) GetExternalProgramId() string`

GetExternalProgramId returns the ExternalProgramId field if non-nil, zero value otherwise.

### GetExternalProgramIdOk

`func (o *ModelSeriesTimerInfoDto) GetExternalProgramIdOk() (*string, bool)`

GetExternalProgramIdOk returns a tuple with the ExternalProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProgramId

`func (o *ModelSeriesTimerInfoDto) SetExternalProgramId(v string)`

SetExternalProgramId sets ExternalProgramId field to given value.

### HasExternalProgramId

`func (o *ModelSeriesTimerInfoDto) HasExternalProgramId() bool`

HasExternalProgramId returns a boolean if a field has been set.

### SetExternalProgramIdNil

`func (o *ModelSeriesTimerInfoDto) SetExternalProgramIdNil(b bool)`

 SetExternalProgramIdNil sets the value for ExternalProgramId to be an explicit nil

### UnsetExternalProgramId
`func (o *ModelSeriesTimerInfoDto) UnsetExternalProgramId()`

UnsetExternalProgramId ensures that no value is present for ExternalProgramId, not even an explicit nil
### GetName

`func (o *ModelSeriesTimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSeriesTimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSeriesTimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSeriesTimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelSeriesTimerInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelSeriesTimerInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOverview

`func (o *ModelSeriesTimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelSeriesTimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelSeriesTimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelSeriesTimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ModelSeriesTimerInfoDto) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ModelSeriesTimerInfoDto) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetStartDate

`func (o *ModelSeriesTimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelSeriesTimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelSeriesTimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelSeriesTimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelSeriesTimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelSeriesTimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelSeriesTimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelSeriesTimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetServiceName

`func (o *ModelSeriesTimerInfoDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ModelSeriesTimerInfoDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ModelSeriesTimerInfoDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ModelSeriesTimerInfoDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ModelSeriesTimerInfoDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ModelSeriesTimerInfoDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetPriority

`func (o *ModelSeriesTimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelSeriesTimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelSeriesTimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelSeriesTimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *ModelSeriesTimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelSeriesTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelSeriesTimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelSeriesTimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelSeriesTimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelSeriesTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelSeriesTimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelSeriesTimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *ModelSeriesTimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *ModelSeriesTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *ModelSeriesTimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *ModelSeriesTimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *ModelSeriesTimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *ModelSeriesTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *ModelSeriesTimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *ModelSeriesTimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *ModelSeriesTimerInfoDto) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *ModelSeriesTimerInfoDto) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *ModelSeriesTimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *ModelSeriesTimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *ModelSeriesTimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *ModelSeriesTimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *ModelSeriesTimerInfoDto) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *ModelSeriesTimerInfoDto) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetIsPostPaddingRequired

`func (o *ModelSeriesTimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *ModelSeriesTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *ModelSeriesTimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *ModelSeriesTimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *ModelSeriesTimerInfoDto) GetKeepUntil() ModelModelKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *ModelSeriesTimerInfoDto) GetKeepUntilOk() (*ModelModelKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *ModelSeriesTimerInfoDto) SetKeepUntil(v ModelModelKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *ModelSeriesTimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetRecordAnyTime

`func (o *ModelSeriesTimerInfoDto) GetRecordAnyTime() bool`

GetRecordAnyTime returns the RecordAnyTime field if non-nil, zero value otherwise.

### GetRecordAnyTimeOk

`func (o *ModelSeriesTimerInfoDto) GetRecordAnyTimeOk() (*bool, bool)`

GetRecordAnyTimeOk returns a tuple with the RecordAnyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyTime

`func (o *ModelSeriesTimerInfoDto) SetRecordAnyTime(v bool)`

SetRecordAnyTime sets RecordAnyTime field to given value.

### HasRecordAnyTime

`func (o *ModelSeriesTimerInfoDto) HasRecordAnyTime() bool`

HasRecordAnyTime returns a boolean if a field has been set.

### GetSkipEpisodesInLibrary

`func (o *ModelSeriesTimerInfoDto) GetSkipEpisodesInLibrary() bool`

GetSkipEpisodesInLibrary returns the SkipEpisodesInLibrary field if non-nil, zero value otherwise.

### GetSkipEpisodesInLibraryOk

`func (o *ModelSeriesTimerInfoDto) GetSkipEpisodesInLibraryOk() (*bool, bool)`

GetSkipEpisodesInLibraryOk returns a tuple with the SkipEpisodesInLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEpisodesInLibrary

`func (o *ModelSeriesTimerInfoDto) SetSkipEpisodesInLibrary(v bool)`

SetSkipEpisodesInLibrary sets SkipEpisodesInLibrary field to given value.

### HasSkipEpisodesInLibrary

`func (o *ModelSeriesTimerInfoDto) HasSkipEpisodesInLibrary() bool`

HasSkipEpisodesInLibrary returns a boolean if a field has been set.

### GetRecordAnyChannel

`func (o *ModelSeriesTimerInfoDto) GetRecordAnyChannel() bool`

GetRecordAnyChannel returns the RecordAnyChannel field if non-nil, zero value otherwise.

### GetRecordAnyChannelOk

`func (o *ModelSeriesTimerInfoDto) GetRecordAnyChannelOk() (*bool, bool)`

GetRecordAnyChannelOk returns a tuple with the RecordAnyChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyChannel

`func (o *ModelSeriesTimerInfoDto) SetRecordAnyChannel(v bool)`

SetRecordAnyChannel sets RecordAnyChannel field to given value.

### HasRecordAnyChannel

`func (o *ModelSeriesTimerInfoDto) HasRecordAnyChannel() bool`

HasRecordAnyChannel returns a boolean if a field has been set.

### GetKeepUpTo

`func (o *ModelSeriesTimerInfoDto) GetKeepUpTo() int32`

GetKeepUpTo returns the KeepUpTo field if non-nil, zero value otherwise.

### GetKeepUpToOk

`func (o *ModelSeriesTimerInfoDto) GetKeepUpToOk() (*int32, bool)`

GetKeepUpToOk returns a tuple with the KeepUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpTo

`func (o *ModelSeriesTimerInfoDto) SetKeepUpTo(v int32)`

SetKeepUpTo sets KeepUpTo field to given value.

### HasKeepUpTo

`func (o *ModelSeriesTimerInfoDto) HasKeepUpTo() bool`

HasKeepUpTo returns a boolean if a field has been set.

### GetRecordNewOnly

`func (o *ModelSeriesTimerInfoDto) GetRecordNewOnly() bool`

GetRecordNewOnly returns the RecordNewOnly field if non-nil, zero value otherwise.

### GetRecordNewOnlyOk

`func (o *ModelSeriesTimerInfoDto) GetRecordNewOnlyOk() (*bool, bool)`

GetRecordNewOnlyOk returns a tuple with the RecordNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNewOnly

`func (o *ModelSeriesTimerInfoDto) SetRecordNewOnly(v bool)`

SetRecordNewOnly sets RecordNewOnly field to given value.

### HasRecordNewOnly

`func (o *ModelSeriesTimerInfoDto) HasRecordNewOnly() bool`

HasRecordNewOnly returns a boolean if a field has been set.

### GetDays

`func (o *ModelSeriesTimerInfoDto) GetDays() []ModelModelDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ModelSeriesTimerInfoDto) GetDaysOk() (*[]ModelModelDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ModelSeriesTimerInfoDto) SetDays(v []ModelModelDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *ModelSeriesTimerInfoDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *ModelSeriesTimerInfoDto) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *ModelSeriesTimerInfoDto) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetDayPattern

`func (o *ModelSeriesTimerInfoDto) GetDayPattern() ModelModelDayPattern`

GetDayPattern returns the DayPattern field if non-nil, zero value otherwise.

### GetDayPatternOk

`func (o *ModelSeriesTimerInfoDto) GetDayPatternOk() (*ModelModelDayPattern, bool)`

GetDayPatternOk returns a tuple with the DayPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayPattern

`func (o *ModelSeriesTimerInfoDto) SetDayPattern(v ModelModelDayPattern)`

SetDayPattern sets DayPattern field to given value.

### HasDayPattern

`func (o *ModelSeriesTimerInfoDto) HasDayPattern() bool`

HasDayPattern returns a boolean if a field has been set.

### SetDayPatternNil

`func (o *ModelSeriesTimerInfoDto) SetDayPatternNil(b bool)`

 SetDayPatternNil sets the value for DayPattern to be an explicit nil

### UnsetDayPattern
`func (o *ModelSeriesTimerInfoDto) UnsetDayPattern()`

UnsetDayPattern ensures that no value is present for DayPattern, not even an explicit nil
### GetImageTags

`func (o *ModelSeriesTimerInfoDto) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *ModelSeriesTimerInfoDto) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *ModelSeriesTimerInfoDto) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *ModelSeriesTimerInfoDto) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### SetImageTagsNil

`func (o *ModelSeriesTimerInfoDto) SetImageTagsNil(b bool)`

 SetImageTagsNil sets the value for ImageTags to be an explicit nil

### UnsetImageTags
`func (o *ModelSeriesTimerInfoDto) UnsetImageTags()`

UnsetImageTags ensures that no value is present for ImageTags, not even an explicit nil
### GetParentThumbItemId

`func (o *ModelSeriesTimerInfoDto) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *ModelSeriesTimerInfoDto) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *ModelSeriesTimerInfoDto) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *ModelSeriesTimerInfoDto) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### SetParentThumbItemIdNil

`func (o *ModelSeriesTimerInfoDto) SetParentThumbItemIdNil(b bool)`

 SetParentThumbItemIdNil sets the value for ParentThumbItemId to be an explicit nil

### UnsetParentThumbItemId
`func (o *ModelSeriesTimerInfoDto) UnsetParentThumbItemId()`

UnsetParentThumbItemId ensures that no value is present for ParentThumbItemId, not even an explicit nil
### GetParentThumbImageTag

`func (o *ModelSeriesTimerInfoDto) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *ModelSeriesTimerInfoDto) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *ModelSeriesTimerInfoDto) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *ModelSeriesTimerInfoDto) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### SetParentThumbImageTagNil

`func (o *ModelSeriesTimerInfoDto) SetParentThumbImageTagNil(b bool)`

 SetParentThumbImageTagNil sets the value for ParentThumbImageTag to be an explicit nil

### UnsetParentThumbImageTag
`func (o *ModelSeriesTimerInfoDto) UnsetParentThumbImageTag()`

UnsetParentThumbImageTag ensures that no value is present for ParentThumbImageTag, not even an explicit nil
### GetParentPrimaryImageItemId

`func (o *ModelSeriesTimerInfoDto) GetParentPrimaryImageItemId() string`

GetParentPrimaryImageItemId returns the ParentPrimaryImageItemId field if non-nil, zero value otherwise.

### GetParentPrimaryImageItemIdOk

`func (o *ModelSeriesTimerInfoDto) GetParentPrimaryImageItemIdOk() (*string, bool)`

GetParentPrimaryImageItemIdOk returns a tuple with the ParentPrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageItemId

`func (o *ModelSeriesTimerInfoDto) SetParentPrimaryImageItemId(v string)`

SetParentPrimaryImageItemId sets ParentPrimaryImageItemId field to given value.

### HasParentPrimaryImageItemId

`func (o *ModelSeriesTimerInfoDto) HasParentPrimaryImageItemId() bool`

HasParentPrimaryImageItemId returns a boolean if a field has been set.

### SetParentPrimaryImageItemIdNil

`func (o *ModelSeriesTimerInfoDto) SetParentPrimaryImageItemIdNil(b bool)`

 SetParentPrimaryImageItemIdNil sets the value for ParentPrimaryImageItemId to be an explicit nil

### UnsetParentPrimaryImageItemId
`func (o *ModelSeriesTimerInfoDto) UnsetParentPrimaryImageItemId()`

UnsetParentPrimaryImageItemId ensures that no value is present for ParentPrimaryImageItemId, not even an explicit nil
### GetParentPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) GetParentPrimaryImageTag() string`

GetParentPrimaryImageTag returns the ParentPrimaryImageTag field if non-nil, zero value otherwise.

### GetParentPrimaryImageTagOk

`func (o *ModelSeriesTimerInfoDto) GetParentPrimaryImageTagOk() (*string, bool)`

GetParentPrimaryImageTagOk returns a tuple with the ParentPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) SetParentPrimaryImageTag(v string)`

SetParentPrimaryImageTag sets ParentPrimaryImageTag field to given value.

### HasParentPrimaryImageTag

`func (o *ModelSeriesTimerInfoDto) HasParentPrimaryImageTag() bool`

HasParentPrimaryImageTag returns a boolean if a field has been set.

### SetParentPrimaryImageTagNil

`func (o *ModelSeriesTimerInfoDto) SetParentPrimaryImageTagNil(b bool)`

 SetParentPrimaryImageTagNil sets the value for ParentPrimaryImageTag to be an explicit nil

### UnsetParentPrimaryImageTag
`func (o *ModelSeriesTimerInfoDto) UnsetParentPrimaryImageTag()`

UnsetParentPrimaryImageTag ensures that no value is present for ParentPrimaryImageTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


