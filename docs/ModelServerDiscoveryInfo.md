# ModelServerDiscoveryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Gets the address. | [optional] 
**Id** | Pointer to **string** | Gets the server identifier. | [optional] 
**Name** | Pointer to **string** | Gets the name. | [optional] 
**EndpointAddress** | Pointer to **NullableString** | Gets the endpoint address. | [optional] 

## Methods

### NewModelServerDiscoveryInfo

`func NewModelServerDiscoveryInfo() *ModelServerDiscoveryInfo`

NewModelServerDiscoveryInfo instantiates a new ModelServerDiscoveryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelServerDiscoveryInfoWithDefaults

`func NewModelServerDiscoveryInfoWithDefaults() *ModelServerDiscoveryInfo`

NewModelServerDiscoveryInfoWithDefaults instantiates a new ModelServerDiscoveryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ModelServerDiscoveryInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ModelServerDiscoveryInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ModelServerDiscoveryInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ModelServerDiscoveryInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *ModelServerDiscoveryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelServerDiscoveryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelServerDiscoveryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelServerDiscoveryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelServerDiscoveryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelServerDiscoveryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelServerDiscoveryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelServerDiscoveryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpointAddress

`func (o *ModelServerDiscoveryInfo) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### GetEndpointAddressOk

`func (o *ModelServerDiscoveryInfo) GetEndpointAddressOk() (*string, bool)`

GetEndpointAddressOk returns a tuple with the EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAddress

`func (o *ModelServerDiscoveryInfo) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### HasEndpointAddress

`func (o *ModelServerDiscoveryInfo) HasEndpointAddress() bool`

HasEndpointAddress returns a boolean if a field has been set.

### SetEndpointAddressNil

`func (o *ModelServerDiscoveryInfo) SetEndpointAddressNil(b bool)`

 SetEndpointAddressNil sets the value for EndpointAddress to be an explicit nil

### UnsetEndpointAddress
`func (o *ModelServerDiscoveryInfo) UnsetEndpointAddress()`

UnsetEndpointAddress ensures that no value is present for EndpointAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

