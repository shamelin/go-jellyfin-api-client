# ModelLiveTvInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]ModelModelLiveTvServiceInfo**](ModelModelLiveTvServiceInfo.md) | Gets or sets the services. | [optional] 
**IsEnabled** | Pointer to **bool** | Gets or sets a value indicating whether this instance is enabled. | [optional] 
**EnabledUsers** | Pointer to **[]string** | Gets or sets the enabled users. | [optional] 

## Methods

### NewModelLiveTvInfo

`func NewModelLiveTvInfo() *ModelLiveTvInfo`

NewModelLiveTvInfo instantiates a new ModelLiveTvInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvInfoWithDefaults

`func NewModelLiveTvInfoWithDefaults() *ModelLiveTvInfo`

NewModelLiveTvInfoWithDefaults instantiates a new ModelLiveTvInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ModelLiveTvInfo) GetServices() []ModelModelLiveTvServiceInfo`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ModelLiveTvInfo) GetServicesOk() (*[]ModelModelLiveTvServiceInfo, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ModelLiveTvInfo) SetServices(v []ModelModelLiveTvServiceInfo)`

SetServices sets Services field to given value.

### HasServices

`func (o *ModelLiveTvInfo) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ModelLiveTvInfo) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ModelLiveTvInfo) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ModelLiveTvInfo) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ModelLiveTvInfo) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetEnabledUsers

`func (o *ModelLiveTvInfo) GetEnabledUsers() []string`

GetEnabledUsers returns the EnabledUsers field if non-nil, zero value otherwise.

### GetEnabledUsersOk

`func (o *ModelLiveTvInfo) GetEnabledUsersOk() (*[]string, bool)`

GetEnabledUsersOk returns a tuple with the EnabledUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledUsers

`func (o *ModelLiveTvInfo) SetEnabledUsers(v []string)`

SetEnabledUsers sets EnabledUsers field to given value.

### HasEnabledUsers

`func (o *ModelLiveTvInfo) HasEnabledUsers() bool`

HasEnabledUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


