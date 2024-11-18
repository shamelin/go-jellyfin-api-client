# ModelValidatePathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateWritable** | Pointer to **bool** | Gets or sets a value indicating whether validate if path is writable. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**IsFile** | Pointer to **NullableBool** | Gets or sets is path file. | [optional] 

## Methods

### NewModelValidatePathDto

`func NewModelValidatePathDto() *ModelValidatePathDto`

NewModelValidatePathDto instantiates a new ModelValidatePathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelValidatePathDtoWithDefaults

`func NewModelValidatePathDtoWithDefaults() *ModelValidatePathDto`

NewModelValidatePathDtoWithDefaults instantiates a new ModelValidatePathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidateWritable

`func (o *ModelValidatePathDto) GetValidateWritable() bool`

GetValidateWritable returns the ValidateWritable field if non-nil, zero value otherwise.

### GetValidateWritableOk

`func (o *ModelValidatePathDto) GetValidateWritableOk() (*bool, bool)`

GetValidateWritableOk returns a tuple with the ValidateWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateWritable

`func (o *ModelValidatePathDto) SetValidateWritable(v bool)`

SetValidateWritable sets ValidateWritable field to given value.

### HasValidateWritable

`func (o *ModelValidatePathDto) HasValidateWritable() bool`

HasValidateWritable returns a boolean if a field has been set.

### GetPath

`func (o *ModelValidatePathDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelValidatePathDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelValidatePathDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelValidatePathDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ModelValidatePathDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ModelValidatePathDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIsFile

`func (o *ModelValidatePathDto) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *ModelValidatePathDto) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *ModelValidatePathDto) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.

### HasIsFile

`func (o *ModelValidatePathDto) HasIsFile() bool`

HasIsFile returns a boolean if a field has been set.

### SetIsFileNil

`func (o *ModelValidatePathDto) SetIsFileNil(b bool)`

 SetIsFileNil sets the value for IsFile to be an explicit nil

### UnsetIsFile
`func (o *ModelValidatePathDto) UnsetIsFile()`

UnsetIsFile ensures that no value is present for IsFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


