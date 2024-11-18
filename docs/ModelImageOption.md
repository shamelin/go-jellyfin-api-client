# ModelImageOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModelModelImageType**](ModelImageType.md) | Gets or sets the type. | [optional] 
**Limit** | Pointer to **int32** | Gets or sets the limit. | [optional] 
**MinWidth** | Pointer to **int32** | Gets or sets the minimum width. | [optional] 

## Methods

### NewModelImageOption

`func NewModelImageOption() *ModelImageOption`

NewModelImageOption instantiates a new ModelImageOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelImageOptionWithDefaults

`func NewModelImageOptionWithDefaults() *ModelImageOption`

NewModelImageOptionWithDefaults instantiates a new ModelImageOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelImageOption) GetType() ModelModelImageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelImageOption) GetTypeOk() (*ModelModelImageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelImageOption) SetType(v ModelModelImageType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelImageOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimit

`func (o *ModelImageOption) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ModelImageOption) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ModelImageOption) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ModelImageOption) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMinWidth

`func (o *ModelImageOption) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ModelImageOption) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ModelImageOption) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ModelImageOption) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


