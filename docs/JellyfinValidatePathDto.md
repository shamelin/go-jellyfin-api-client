# JellyfinValidatePathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateWritable** | Pointer to **bool** | Gets or sets a value indicating whether validate if path is writable. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**IsFile** | Pointer to **NullableBool** | Gets or sets is path file. | [optional] 

## Methods

### NewJellyfinValidatePathDto

`func NewJellyfinValidatePathDto() *JellyfinValidatePathDto`

NewJellyfinValidatePathDto instantiates a new JellyfinValidatePathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinValidatePathDtoWithDefaults

`func NewJellyfinValidatePathDtoWithDefaults() *JellyfinValidatePathDto`

NewJellyfinValidatePathDtoWithDefaults instantiates a new JellyfinValidatePathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidateWritable

`func (o *JellyfinValidatePathDto) GetValidateWritable() bool`

GetValidateWritable returns the ValidateWritable field if non-nil, zero value otherwise.

### GetValidateWritableOk

`func (o *JellyfinValidatePathDto) GetValidateWritableOk() (*bool, bool)`

GetValidateWritableOk returns a tuple with the ValidateWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateWritable

`func (o *JellyfinValidatePathDto) SetValidateWritable(v bool)`

SetValidateWritable sets ValidateWritable field to given value.

### HasValidateWritable

`func (o *JellyfinValidatePathDto) HasValidateWritable() bool`

HasValidateWritable returns a boolean if a field has been set.

### GetPath

`func (o *JellyfinValidatePathDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JellyfinValidatePathDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JellyfinValidatePathDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *JellyfinValidatePathDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *JellyfinValidatePathDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *JellyfinValidatePathDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIsFile

`func (o *JellyfinValidatePathDto) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *JellyfinValidatePathDto) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *JellyfinValidatePathDto) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.

### HasIsFile

`func (o *JellyfinValidatePathDto) HasIsFile() bool`

HasIsFile returns a boolean if a field has been set.

### SetIsFileNil

`func (o *JellyfinValidatePathDto) SetIsFileNil(b bool)`

 SetIsFileNil sets the value for IsFile to be an explicit nil

### UnsetIsFile
`func (o *JellyfinValidatePathDto) UnsetIsFile()`

UnsetIsFile ensures that no value is present for IsFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


