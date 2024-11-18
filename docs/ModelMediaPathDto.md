# ModelMediaPathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the name of the library. | 
**Path** | Pointer to **NullableString** | Gets or sets the path to add. | [optional] 
**PathInfo** | Pointer to [**NullableModelModelMediaPathInfo**](ModelMediaPathInfo.md) | Gets or sets the path info. | [optional] 

## Methods

### NewModelMediaPathDto

`func NewModelMediaPathDto(name string, ) *ModelMediaPathDto`

NewModelMediaPathDto instantiates a new ModelMediaPathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMediaPathDtoWithDefaults

`func NewModelMediaPathDtoWithDefaults() *ModelMediaPathDto`

NewModelMediaPathDtoWithDefaults instantiates a new ModelMediaPathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelMediaPathDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMediaPathDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMediaPathDto) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *ModelMediaPathDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelMediaPathDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelMediaPathDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelMediaPathDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ModelMediaPathDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ModelMediaPathDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPathInfo

`func (o *ModelMediaPathDto) GetPathInfo() ModelModelMediaPathInfo`

GetPathInfo returns the PathInfo field if non-nil, zero value otherwise.

### GetPathInfoOk

`func (o *ModelMediaPathDto) GetPathInfoOk() (*ModelModelMediaPathInfo, bool)`

GetPathInfoOk returns a tuple with the PathInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfo

`func (o *ModelMediaPathDto) SetPathInfo(v ModelModelMediaPathInfo)`

SetPathInfo sets PathInfo field to given value.

### HasPathInfo

`func (o *ModelMediaPathDto) HasPathInfo() bool`

HasPathInfo returns a boolean if a field has been set.

### SetPathInfoNil

`func (o *ModelMediaPathDto) SetPathInfoNil(b bool)`

 SetPathInfoNil sets the value for PathInfo to be an explicit nil

### UnsetPathInfo
`func (o *ModelMediaPathDto) UnsetPathInfo()`

UnsetPathInfo ensures that no value is present for PathInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


