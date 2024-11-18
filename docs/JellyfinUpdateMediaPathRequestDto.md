# JellyfinUpdateMediaPathRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the library name. | 
**PathInfo** | [**JellyfinJellyfinMediaPathInfo**](JellyfinMediaPathInfo.md) | Gets or sets library folder path information. | 

## Methods

### NewJellyfinUpdateMediaPathRequestDto

`func NewJellyfinUpdateMediaPathRequestDto(name string, pathInfo JellyfinJellyfinMediaPathInfo, ) *JellyfinUpdateMediaPathRequestDto`

NewJellyfinUpdateMediaPathRequestDto instantiates a new JellyfinUpdateMediaPathRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinUpdateMediaPathRequestDtoWithDefaults

`func NewJellyfinUpdateMediaPathRequestDtoWithDefaults() *JellyfinUpdateMediaPathRequestDto`

NewJellyfinUpdateMediaPathRequestDtoWithDefaults instantiates a new JellyfinUpdateMediaPathRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinUpdateMediaPathRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinUpdateMediaPathRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinUpdateMediaPathRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetPathInfo

`func (o *JellyfinUpdateMediaPathRequestDto) GetPathInfo() JellyfinJellyfinMediaPathInfo`

GetPathInfo returns the PathInfo field if non-nil, zero value otherwise.

### GetPathInfoOk

`func (o *JellyfinUpdateMediaPathRequestDto) GetPathInfoOk() (*JellyfinJellyfinMediaPathInfo, bool)`

GetPathInfoOk returns a tuple with the PathInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfo

`func (o *JellyfinUpdateMediaPathRequestDto) SetPathInfo(v JellyfinJellyfinMediaPathInfo)`

SetPathInfo sets PathInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


