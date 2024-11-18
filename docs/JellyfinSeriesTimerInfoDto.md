# JellyfinSeriesTimerInfoDto

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
**KeepUntil** | Pointer to [**JellyfinJellyfinKeepUntil**](JellyfinKeepUntil.md) |  | [optional] 
**RecordAnyTime** | Pointer to **bool** | Gets or sets a value indicating whether [record any time]. | [optional] 
**SkipEpisodesInLibrary** | Pointer to **bool** |  | [optional] 
**RecordAnyChannel** | Pointer to **bool** | Gets or sets a value indicating whether [record any channel]. | [optional] 
**KeepUpTo** | Pointer to **int32** |  | [optional] 
**RecordNewOnly** | Pointer to **bool** | Gets or sets a value indicating whether [record new only]. | [optional] 
**Days** | Pointer to [**[]JellyfinJellyfinDayOfWeek**](JellyfinJellyfinDayOfWeek.md) | Gets or sets the days. | [optional] 
**DayPattern** | Pointer to [**NullableJellyfinJellyfinDayPattern**](JellyfinDayPattern.md) | Gets or sets the day pattern. | [optional] 
**ImageTags** | Pointer to **map[string]string** | Gets or sets the image tags. | [optional] 
**ParentThumbItemId** | Pointer to **NullableString** | Gets or sets the parent thumb item id. | [optional] 
**ParentThumbImageTag** | Pointer to **NullableString** | Gets or sets the parent thumb image tag. | [optional] 
**ParentPrimaryImageItemId** | Pointer to **NullableString** | Gets or sets the parent primary image item identifier. | [optional] 
**ParentPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the parent primary image tag. | [optional] 

## Methods

### NewJellyfinSeriesTimerInfoDto

`func NewJellyfinSeriesTimerInfoDto() *JellyfinSeriesTimerInfoDto`

NewJellyfinSeriesTimerInfoDto instantiates a new JellyfinSeriesTimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSeriesTimerInfoDtoWithDefaults

`func NewJellyfinSeriesTimerInfoDtoWithDefaults() *JellyfinSeriesTimerInfoDto`

NewJellyfinSeriesTimerInfoDtoWithDefaults instantiates a new JellyfinSeriesTimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JellyfinSeriesTimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinSeriesTimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinSeriesTimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JellyfinSeriesTimerInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *JellyfinSeriesTimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinSeriesTimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinSeriesTimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinSeriesTimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JellyfinSeriesTimerInfoDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JellyfinSeriesTimerInfoDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetServerId

`func (o *JellyfinSeriesTimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *JellyfinSeriesTimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *JellyfinSeriesTimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *JellyfinSeriesTimerInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetExternalId

`func (o *JellyfinSeriesTimerInfoDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *JellyfinSeriesTimerInfoDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *JellyfinSeriesTimerInfoDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *JellyfinSeriesTimerInfoDto) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetChannelId

`func (o *JellyfinSeriesTimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *JellyfinSeriesTimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *JellyfinSeriesTimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetExternalChannelId

`func (o *JellyfinSeriesTimerInfoDto) GetExternalChannelId() string`

GetExternalChannelId returns the ExternalChannelId field if non-nil, zero value otherwise.

### GetExternalChannelIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetExternalChannelIdOk() (*string, bool)`

GetExternalChannelIdOk returns a tuple with the ExternalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelId

`func (o *JellyfinSeriesTimerInfoDto) SetExternalChannelId(v string)`

SetExternalChannelId sets ExternalChannelId field to given value.

### HasExternalChannelId

`func (o *JellyfinSeriesTimerInfoDto) HasExternalChannelId() bool`

HasExternalChannelId returns a boolean if a field has been set.

### SetExternalChannelIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetExternalChannelIdNil(b bool)`

 SetExternalChannelIdNil sets the value for ExternalChannelId to be an explicit nil

### UnsetExternalChannelId
`func (o *JellyfinSeriesTimerInfoDto) UnsetExternalChannelId()`

UnsetExternalChannelId ensures that no value is present for ExternalChannelId, not even an explicit nil
### GetChannelName

`func (o *JellyfinSeriesTimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *JellyfinSeriesTimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *JellyfinSeriesTimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *JellyfinSeriesTimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *JellyfinSeriesTimerInfoDto) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *JellyfinSeriesTimerInfoDto) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *JellyfinSeriesTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *JellyfinSeriesTimerInfoDto) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *JellyfinSeriesTimerInfoDto) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetProgramId

`func (o *JellyfinSeriesTimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *JellyfinSeriesTimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *JellyfinSeriesTimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *JellyfinSeriesTimerInfoDto) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetExternalProgramId

`func (o *JellyfinSeriesTimerInfoDto) GetExternalProgramId() string`

GetExternalProgramId returns the ExternalProgramId field if non-nil, zero value otherwise.

### GetExternalProgramIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetExternalProgramIdOk() (*string, bool)`

GetExternalProgramIdOk returns a tuple with the ExternalProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProgramId

`func (o *JellyfinSeriesTimerInfoDto) SetExternalProgramId(v string)`

SetExternalProgramId sets ExternalProgramId field to given value.

### HasExternalProgramId

`func (o *JellyfinSeriesTimerInfoDto) HasExternalProgramId() bool`

HasExternalProgramId returns a boolean if a field has been set.

### SetExternalProgramIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetExternalProgramIdNil(b bool)`

 SetExternalProgramIdNil sets the value for ExternalProgramId to be an explicit nil

### UnsetExternalProgramId
`func (o *JellyfinSeriesTimerInfoDto) UnsetExternalProgramId()`

UnsetExternalProgramId ensures that no value is present for ExternalProgramId, not even an explicit nil
### GetName

`func (o *JellyfinSeriesTimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinSeriesTimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinSeriesTimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinSeriesTimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JellyfinSeriesTimerInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JellyfinSeriesTimerInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOverview

`func (o *JellyfinSeriesTimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *JellyfinSeriesTimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *JellyfinSeriesTimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *JellyfinSeriesTimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *JellyfinSeriesTimerInfoDto) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *JellyfinSeriesTimerInfoDto) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetStartDate

`func (o *JellyfinSeriesTimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *JellyfinSeriesTimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *JellyfinSeriesTimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *JellyfinSeriesTimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *JellyfinSeriesTimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *JellyfinSeriesTimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *JellyfinSeriesTimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *JellyfinSeriesTimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetServiceName

`func (o *JellyfinSeriesTimerInfoDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *JellyfinSeriesTimerInfoDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *JellyfinSeriesTimerInfoDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *JellyfinSeriesTimerInfoDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *JellyfinSeriesTimerInfoDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *JellyfinSeriesTimerInfoDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetPriority

`func (o *JellyfinSeriesTimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JellyfinSeriesTimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JellyfinSeriesTimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JellyfinSeriesTimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *JellyfinSeriesTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *JellyfinSeriesTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *JellyfinSeriesTimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *JellyfinSeriesTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *JellyfinSeriesTimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *JellyfinSeriesTimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *JellyfinSeriesTimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *JellyfinSeriesTimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *JellyfinSeriesTimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *JellyfinSeriesTimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetIsPostPaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *JellyfinSeriesTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *JellyfinSeriesTimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *JellyfinSeriesTimerInfoDto) GetKeepUntil() JellyfinJellyfinKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *JellyfinSeriesTimerInfoDto) GetKeepUntilOk() (*JellyfinJellyfinKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *JellyfinSeriesTimerInfoDto) SetKeepUntil(v JellyfinJellyfinKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *JellyfinSeriesTimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetRecordAnyTime

`func (o *JellyfinSeriesTimerInfoDto) GetRecordAnyTime() bool`

GetRecordAnyTime returns the RecordAnyTime field if non-nil, zero value otherwise.

### GetRecordAnyTimeOk

`func (o *JellyfinSeriesTimerInfoDto) GetRecordAnyTimeOk() (*bool, bool)`

GetRecordAnyTimeOk returns a tuple with the RecordAnyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyTime

`func (o *JellyfinSeriesTimerInfoDto) SetRecordAnyTime(v bool)`

SetRecordAnyTime sets RecordAnyTime field to given value.

### HasRecordAnyTime

`func (o *JellyfinSeriesTimerInfoDto) HasRecordAnyTime() bool`

HasRecordAnyTime returns a boolean if a field has been set.

### GetSkipEpisodesInLibrary

`func (o *JellyfinSeriesTimerInfoDto) GetSkipEpisodesInLibrary() bool`

GetSkipEpisodesInLibrary returns the SkipEpisodesInLibrary field if non-nil, zero value otherwise.

### GetSkipEpisodesInLibraryOk

`func (o *JellyfinSeriesTimerInfoDto) GetSkipEpisodesInLibraryOk() (*bool, bool)`

GetSkipEpisodesInLibraryOk returns a tuple with the SkipEpisodesInLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEpisodesInLibrary

`func (o *JellyfinSeriesTimerInfoDto) SetSkipEpisodesInLibrary(v bool)`

SetSkipEpisodesInLibrary sets SkipEpisodesInLibrary field to given value.

### HasSkipEpisodesInLibrary

`func (o *JellyfinSeriesTimerInfoDto) HasSkipEpisodesInLibrary() bool`

HasSkipEpisodesInLibrary returns a boolean if a field has been set.

### GetRecordAnyChannel

`func (o *JellyfinSeriesTimerInfoDto) GetRecordAnyChannel() bool`

GetRecordAnyChannel returns the RecordAnyChannel field if non-nil, zero value otherwise.

### GetRecordAnyChannelOk

`func (o *JellyfinSeriesTimerInfoDto) GetRecordAnyChannelOk() (*bool, bool)`

GetRecordAnyChannelOk returns a tuple with the RecordAnyChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyChannel

`func (o *JellyfinSeriesTimerInfoDto) SetRecordAnyChannel(v bool)`

SetRecordAnyChannel sets RecordAnyChannel field to given value.

### HasRecordAnyChannel

`func (o *JellyfinSeriesTimerInfoDto) HasRecordAnyChannel() bool`

HasRecordAnyChannel returns a boolean if a field has been set.

### GetKeepUpTo

`func (o *JellyfinSeriesTimerInfoDto) GetKeepUpTo() int32`

GetKeepUpTo returns the KeepUpTo field if non-nil, zero value otherwise.

### GetKeepUpToOk

`func (o *JellyfinSeriesTimerInfoDto) GetKeepUpToOk() (*int32, bool)`

GetKeepUpToOk returns a tuple with the KeepUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpTo

`func (o *JellyfinSeriesTimerInfoDto) SetKeepUpTo(v int32)`

SetKeepUpTo sets KeepUpTo field to given value.

### HasKeepUpTo

`func (o *JellyfinSeriesTimerInfoDto) HasKeepUpTo() bool`

HasKeepUpTo returns a boolean if a field has been set.

### GetRecordNewOnly

`func (o *JellyfinSeriesTimerInfoDto) GetRecordNewOnly() bool`

GetRecordNewOnly returns the RecordNewOnly field if non-nil, zero value otherwise.

### GetRecordNewOnlyOk

`func (o *JellyfinSeriesTimerInfoDto) GetRecordNewOnlyOk() (*bool, bool)`

GetRecordNewOnlyOk returns a tuple with the RecordNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNewOnly

`func (o *JellyfinSeriesTimerInfoDto) SetRecordNewOnly(v bool)`

SetRecordNewOnly sets RecordNewOnly field to given value.

### HasRecordNewOnly

`func (o *JellyfinSeriesTimerInfoDto) HasRecordNewOnly() bool`

HasRecordNewOnly returns a boolean if a field has been set.

### GetDays

`func (o *JellyfinSeriesTimerInfoDto) GetDays() []JellyfinJellyfinDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *JellyfinSeriesTimerInfoDto) GetDaysOk() (*[]JellyfinJellyfinDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *JellyfinSeriesTimerInfoDto) SetDays(v []JellyfinJellyfinDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *JellyfinSeriesTimerInfoDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *JellyfinSeriesTimerInfoDto) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *JellyfinSeriesTimerInfoDto) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetDayPattern

`func (o *JellyfinSeriesTimerInfoDto) GetDayPattern() JellyfinJellyfinDayPattern`

GetDayPattern returns the DayPattern field if non-nil, zero value otherwise.

### GetDayPatternOk

`func (o *JellyfinSeriesTimerInfoDto) GetDayPatternOk() (*JellyfinJellyfinDayPattern, bool)`

GetDayPatternOk returns a tuple with the DayPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayPattern

`func (o *JellyfinSeriesTimerInfoDto) SetDayPattern(v JellyfinJellyfinDayPattern)`

SetDayPattern sets DayPattern field to given value.

### HasDayPattern

`func (o *JellyfinSeriesTimerInfoDto) HasDayPattern() bool`

HasDayPattern returns a boolean if a field has been set.

### SetDayPatternNil

`func (o *JellyfinSeriesTimerInfoDto) SetDayPatternNil(b bool)`

 SetDayPatternNil sets the value for DayPattern to be an explicit nil

### UnsetDayPattern
`func (o *JellyfinSeriesTimerInfoDto) UnsetDayPattern()`

UnsetDayPattern ensures that no value is present for DayPattern, not even an explicit nil
### GetImageTags

`func (o *JellyfinSeriesTimerInfoDto) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *JellyfinSeriesTimerInfoDto) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *JellyfinSeriesTimerInfoDto) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *JellyfinSeriesTimerInfoDto) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### SetImageTagsNil

`func (o *JellyfinSeriesTimerInfoDto) SetImageTagsNil(b bool)`

 SetImageTagsNil sets the value for ImageTags to be an explicit nil

### UnsetImageTags
`func (o *JellyfinSeriesTimerInfoDto) UnsetImageTags()`

UnsetImageTags ensures that no value is present for ImageTags, not even an explicit nil
### GetParentThumbItemId

`func (o *JellyfinSeriesTimerInfoDto) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *JellyfinSeriesTimerInfoDto) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *JellyfinSeriesTimerInfoDto) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### SetParentThumbItemIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentThumbItemIdNil(b bool)`

 SetParentThumbItemIdNil sets the value for ParentThumbItemId to be an explicit nil

### UnsetParentThumbItemId
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentThumbItemId()`

UnsetParentThumbItemId ensures that no value is present for ParentThumbItemId, not even an explicit nil
### GetParentThumbImageTag

`func (o *JellyfinSeriesTimerInfoDto) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *JellyfinSeriesTimerInfoDto) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *JellyfinSeriesTimerInfoDto) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### SetParentThumbImageTagNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentThumbImageTagNil(b bool)`

 SetParentThumbImageTagNil sets the value for ParentThumbImageTag to be an explicit nil

### UnsetParentThumbImageTag
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentThumbImageTag()`

UnsetParentThumbImageTag ensures that no value is present for ParentThumbImageTag, not even an explicit nil
### GetParentPrimaryImageItemId

`func (o *JellyfinSeriesTimerInfoDto) GetParentPrimaryImageItemId() string`

GetParentPrimaryImageItemId returns the ParentPrimaryImageItemId field if non-nil, zero value otherwise.

### GetParentPrimaryImageItemIdOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentPrimaryImageItemIdOk() (*string, bool)`

GetParentPrimaryImageItemIdOk returns a tuple with the ParentPrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageItemId

`func (o *JellyfinSeriesTimerInfoDto) SetParentPrimaryImageItemId(v string)`

SetParentPrimaryImageItemId sets ParentPrimaryImageItemId field to given value.

### HasParentPrimaryImageItemId

`func (o *JellyfinSeriesTimerInfoDto) HasParentPrimaryImageItemId() bool`

HasParentPrimaryImageItemId returns a boolean if a field has been set.

### SetParentPrimaryImageItemIdNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentPrimaryImageItemIdNil(b bool)`

 SetParentPrimaryImageItemIdNil sets the value for ParentPrimaryImageItemId to be an explicit nil

### UnsetParentPrimaryImageItemId
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentPrimaryImageItemId()`

UnsetParentPrimaryImageItemId ensures that no value is present for ParentPrimaryImageItemId, not even an explicit nil
### GetParentPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) GetParentPrimaryImageTag() string`

GetParentPrimaryImageTag returns the ParentPrimaryImageTag field if non-nil, zero value otherwise.

### GetParentPrimaryImageTagOk

`func (o *JellyfinSeriesTimerInfoDto) GetParentPrimaryImageTagOk() (*string, bool)`

GetParentPrimaryImageTagOk returns a tuple with the ParentPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) SetParentPrimaryImageTag(v string)`

SetParentPrimaryImageTag sets ParentPrimaryImageTag field to given value.

### HasParentPrimaryImageTag

`func (o *JellyfinSeriesTimerInfoDto) HasParentPrimaryImageTag() bool`

HasParentPrimaryImageTag returns a boolean if a field has been set.

### SetParentPrimaryImageTagNil

`func (o *JellyfinSeriesTimerInfoDto) SetParentPrimaryImageTagNil(b bool)`

 SetParentPrimaryImageTagNil sets the value for ParentPrimaryImageTag to be an explicit nil

### UnsetParentPrimaryImageTag
`func (o *JellyfinSeriesTimerInfoDto) UnsetParentPrimaryImageTag()`

UnsetParentPrimaryImageTag ensures that no value is present for ParentPrimaryImageTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


