# JellyfinActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Gets or sets the identifier. | [optional] 
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Overview** | Pointer to **NullableString** | Gets or sets the overview. | [optional] 
**ShortOverview** | Pointer to **NullableString** | Gets or sets the short overview. | [optional] 
**Type** | Pointer to **string** | Gets or sets the type. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 
**Date** | Pointer to **time.Time** | Gets or sets the date. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the user primary image tag. | [optional] 
**Severity** | Pointer to [**JellyfinJellyfinLogLevel**](JellyfinLogLevel.md) | Gets or sets the log severity. | [optional] 

## Methods

### NewJellyfinActivityLogEntry

`func NewJellyfinActivityLogEntry() *JellyfinActivityLogEntry`

NewJellyfinActivityLogEntry instantiates a new JellyfinActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinActivityLogEntryWithDefaults

`func NewJellyfinActivityLogEntryWithDefaults() *JellyfinActivityLogEntry`

NewJellyfinActivityLogEntryWithDefaults instantiates a new JellyfinActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JellyfinActivityLogEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinActivityLogEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinActivityLogEntry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinActivityLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JellyfinActivityLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinActivityLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinActivityLogEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinActivityLogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *JellyfinActivityLogEntry) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *JellyfinActivityLogEntry) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *JellyfinActivityLogEntry) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *JellyfinActivityLogEntry) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *JellyfinActivityLogEntry) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *JellyfinActivityLogEntry) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetShortOverview

`func (o *JellyfinActivityLogEntry) GetShortOverview() string`

GetShortOverview returns the ShortOverview field if non-nil, zero value otherwise.

### GetShortOverviewOk

`func (o *JellyfinActivityLogEntry) GetShortOverviewOk() (*string, bool)`

GetShortOverviewOk returns a tuple with the ShortOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOverview

`func (o *JellyfinActivityLogEntry) SetShortOverview(v string)`

SetShortOverview sets ShortOverview field to given value.

### HasShortOverview

`func (o *JellyfinActivityLogEntry) HasShortOverview() bool`

HasShortOverview returns a boolean if a field has been set.

### SetShortOverviewNil

`func (o *JellyfinActivityLogEntry) SetShortOverviewNil(b bool)`

 SetShortOverviewNil sets the value for ShortOverview to be an explicit nil

### UnsetShortOverview
`func (o *JellyfinActivityLogEntry) UnsetShortOverview()`

UnsetShortOverview ensures that no value is present for ShortOverview, not even an explicit nil
### GetType

`func (o *JellyfinActivityLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinActivityLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinActivityLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinActivityLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemId

`func (o *JellyfinActivityLogEntry) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinActivityLogEntry) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinActivityLogEntry) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinActivityLogEntry) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *JellyfinActivityLogEntry) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *JellyfinActivityLogEntry) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetDate

`func (o *JellyfinActivityLogEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *JellyfinActivityLogEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *JellyfinActivityLogEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *JellyfinActivityLogEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUserId

`func (o *JellyfinActivityLogEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *JellyfinActivityLogEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *JellyfinActivityLogEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *JellyfinActivityLogEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *JellyfinActivityLogEntry) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *JellyfinActivityLogEntry) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *JellyfinActivityLogEntry) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *JellyfinActivityLogEntry) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *JellyfinActivityLogEntry) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *JellyfinActivityLogEntry) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSeverity

`func (o *JellyfinActivityLogEntry) GetSeverity() JellyfinJellyfinLogLevel`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *JellyfinActivityLogEntry) GetSeverityOk() (*JellyfinJellyfinLogLevel, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *JellyfinActivityLogEntry) SetSeverity(v JellyfinJellyfinLogLevel)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *JellyfinActivityLogEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


