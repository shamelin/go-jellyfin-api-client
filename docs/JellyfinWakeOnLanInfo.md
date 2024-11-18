# JellyfinWakeOnLanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **NullableString** | Gets the MAC address of the device. | [optional] 
**Port** | Pointer to **int32** | Gets or sets the wake-on-LAN port. | [optional] 

## Methods

### NewJellyfinWakeOnLanInfo

`func NewJellyfinWakeOnLanInfo() *JellyfinWakeOnLanInfo`

NewJellyfinWakeOnLanInfo instantiates a new JellyfinWakeOnLanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinWakeOnLanInfoWithDefaults

`func NewJellyfinWakeOnLanInfoWithDefaults() *JellyfinWakeOnLanInfo`

NewJellyfinWakeOnLanInfoWithDefaults instantiates a new JellyfinWakeOnLanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *JellyfinWakeOnLanInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *JellyfinWakeOnLanInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *JellyfinWakeOnLanInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *JellyfinWakeOnLanInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *JellyfinWakeOnLanInfo) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *JellyfinWakeOnLanInfo) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetPort

`func (o *JellyfinWakeOnLanInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *JellyfinWakeOnLanInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *JellyfinWakeOnLanInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *JellyfinWakeOnLanInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


