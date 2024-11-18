# ModelDeviceOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Gets or sets the id. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**CustomName** | Pointer to **NullableString** | Gets or sets the custom name. | [optional] 

## Methods

### NewModelDeviceOptionsDto

`func NewModelDeviceOptionsDto() *ModelDeviceOptionsDto`

NewModelDeviceOptionsDto instantiates a new ModelDeviceOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDeviceOptionsDtoWithDefaults

`func NewModelDeviceOptionsDtoWithDefaults() *ModelDeviceOptionsDto`

NewModelDeviceOptionsDtoWithDefaults instantiates a new ModelDeviceOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelDeviceOptionsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDeviceOptionsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDeviceOptionsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDeviceOptionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ModelDeviceOptionsDto) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelDeviceOptionsDto) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelDeviceOptionsDto) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelDeviceOptionsDto) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *ModelDeviceOptionsDto) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *ModelDeviceOptionsDto) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetCustomName

`func (o *ModelDeviceOptionsDto) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ModelDeviceOptionsDto) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ModelDeviceOptionsDto) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ModelDeviceOptionsDto) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ModelDeviceOptionsDto) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ModelDeviceOptionsDto) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


