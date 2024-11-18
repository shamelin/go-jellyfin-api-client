# ModelQueueRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemIds** | Pointer to **[]string** | Gets or sets the items to enqueue. | [optional] 
**Mode** | Pointer to [**ModelModelGroupQueueMode**](ModelGroupQueueMode.md) | Gets or sets the mode in which to add the new items. | [optional] 

## Methods

### NewModelQueueRequestDto

`func NewModelQueueRequestDto() *ModelQueueRequestDto`

NewModelQueueRequestDto instantiates a new ModelQueueRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQueueRequestDtoWithDefaults

`func NewModelQueueRequestDtoWithDefaults() *ModelQueueRequestDto`

NewModelQueueRequestDtoWithDefaults instantiates a new ModelQueueRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemIds

`func (o *ModelQueueRequestDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ModelQueueRequestDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ModelQueueRequestDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ModelQueueRequestDto) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetMode

`func (o *ModelQueueRequestDto) GetMode() ModelModelGroupQueueMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ModelQueueRequestDto) GetModeOk() (*ModelModelGroupQueueMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ModelQueueRequestDto) SetMode(v ModelModelGroupQueueMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ModelQueueRequestDto) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


