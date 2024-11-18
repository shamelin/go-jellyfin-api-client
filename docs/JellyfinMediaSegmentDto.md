# JellyfinMediaSegmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the id of the media segment. | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the id of the associated item. | [optional] 
**Type** | Pointer to [**JellyfinJellyfinMediaSegmentType**](JellyfinMediaSegmentType.md) | Defines the types of content an individual Jellyfin.Data.Entities.MediaSegment represents. | [optional] 
**StartTicks** | Pointer to **int64** | Gets or sets the start of the segment. | [optional] 
**EndTicks** | Pointer to **int64** | Gets or sets the end of the segment. | [optional] 

## Methods

### NewJellyfinMediaSegmentDto

`func NewJellyfinMediaSegmentDto() *JellyfinMediaSegmentDto`

NewJellyfinMediaSegmentDto instantiates a new JellyfinMediaSegmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinMediaSegmentDtoWithDefaults

`func NewJellyfinMediaSegmentDtoWithDefaults() *JellyfinMediaSegmentDto`

NewJellyfinMediaSegmentDtoWithDefaults instantiates a new JellyfinMediaSegmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JellyfinMediaSegmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JellyfinMediaSegmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JellyfinMediaSegmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JellyfinMediaSegmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *JellyfinMediaSegmentDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *JellyfinMediaSegmentDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *JellyfinMediaSegmentDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *JellyfinMediaSegmentDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetType

`func (o *JellyfinMediaSegmentDto) GetType() JellyfinJellyfinMediaSegmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JellyfinMediaSegmentDto) GetTypeOk() (*JellyfinJellyfinMediaSegmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JellyfinMediaSegmentDto) SetType(v JellyfinJellyfinMediaSegmentType)`

SetType sets Type field to given value.

### HasType

`func (o *JellyfinMediaSegmentDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartTicks

`func (o *JellyfinMediaSegmentDto) GetStartTicks() int64`

GetStartTicks returns the StartTicks field if non-nil, zero value otherwise.

### GetStartTicksOk

`func (o *JellyfinMediaSegmentDto) GetStartTicksOk() (*int64, bool)`

GetStartTicksOk returns a tuple with the StartTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTicks

`func (o *JellyfinMediaSegmentDto) SetStartTicks(v int64)`

SetStartTicks sets StartTicks field to given value.

### HasStartTicks

`func (o *JellyfinMediaSegmentDto) HasStartTicks() bool`

HasStartTicks returns a boolean if a field has been set.

### GetEndTicks

`func (o *JellyfinMediaSegmentDto) GetEndTicks() int64`

GetEndTicks returns the EndTicks field if non-nil, zero value otherwise.

### GetEndTicksOk

`func (o *JellyfinMediaSegmentDto) GetEndTicksOk() (*int64, bool)`

GetEndTicksOk returns a tuple with the EndTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTicks

`func (o *JellyfinMediaSegmentDto) SetEndTicks(v int64)`

SetEndTicks sets EndTicks field to given value.

### HasEndTicks

`func (o *JellyfinMediaSegmentDto) HasEndTicks() bool`

HasEndTicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


