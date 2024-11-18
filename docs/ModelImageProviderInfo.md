# ModelImageProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets the name. | [optional] 
**SupportedImages** | Pointer to [**[]ModelModelImageType**](ModelModelImageType.md) | Gets the supported image types. | [optional] 

## Methods

### NewModelImageProviderInfo

`func NewModelImageProviderInfo() *ModelImageProviderInfo`

NewModelImageProviderInfo instantiates a new ModelImageProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelImageProviderInfoWithDefaults

`func NewModelImageProviderInfoWithDefaults() *ModelImageProviderInfo`

NewModelImageProviderInfoWithDefaults instantiates a new ModelImageProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelImageProviderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelImageProviderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelImageProviderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelImageProviderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSupportedImages

`func (o *ModelImageProviderInfo) GetSupportedImages() []ModelModelImageType`

GetSupportedImages returns the SupportedImages field if non-nil, zero value otherwise.

### GetSupportedImagesOk

`func (o *ModelImageProviderInfo) GetSupportedImagesOk() (*[]ModelModelImageType, bool)`

GetSupportedImagesOk returns a tuple with the SupportedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImages

`func (o *ModelImageProviderInfo) SetSupportedImages(v []ModelModelImageType)`

SetSupportedImages sets SupportedImages field to given value.

### HasSupportedImages

`func (o *ModelImageProviderInfo) HasSupportedImages() bool`

HasSupportedImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


