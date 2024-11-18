# ModelUpdateMediaPathRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the library name. | 
**PathInfo** | [**ModelModelMediaPathInfo**](ModelMediaPathInfo.md) | Gets or sets library folder path information. | 

## Methods

### NewModelUpdateMediaPathRequestDto

`func NewModelUpdateMediaPathRequestDto(name string, pathInfo ModelModelMediaPathInfo, ) *ModelUpdateMediaPathRequestDto`

NewModelUpdateMediaPathRequestDto instantiates a new ModelUpdateMediaPathRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpdateMediaPathRequestDtoWithDefaults

`func NewModelUpdateMediaPathRequestDtoWithDefaults() *ModelUpdateMediaPathRequestDto`

NewModelUpdateMediaPathRequestDtoWithDefaults instantiates a new ModelUpdateMediaPathRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelUpdateMediaPathRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelUpdateMediaPathRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelUpdateMediaPathRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetPathInfo

`func (o *ModelUpdateMediaPathRequestDto) GetPathInfo() ModelModelMediaPathInfo`

GetPathInfo returns the PathInfo field if non-nil, zero value otherwise.

### GetPathInfoOk

`func (o *ModelUpdateMediaPathRequestDto) GetPathInfoOk() (*ModelModelMediaPathInfo, bool)`

GetPathInfoOk returns a tuple with the PathInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfo

`func (o *ModelUpdateMediaPathRequestDto) SetPathInfo(v ModelModelMediaPathInfo)`

SetPathInfo sets PathInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


