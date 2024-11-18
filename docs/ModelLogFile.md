# ModelLogFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Gets or sets the date created. | [optional] 
**DateModified** | Pointer to **time.Time** | Gets or sets the date modified. | [optional] 
**Size** | Pointer to **int64** | Gets or sets the size. | [optional] 
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 

## Methods

### NewModelLogFile

`func NewModelLogFile() *ModelLogFile`

NewModelLogFile instantiates a new ModelLogFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLogFileWithDefaults

`func NewModelLogFileWithDefaults() *ModelLogFile`

NewModelLogFileWithDefaults instantiates a new ModelLogFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *ModelLogFile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelLogFile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelLogFile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelLogFile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *ModelLogFile) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *ModelLogFile) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *ModelLogFile) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *ModelLogFile) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetSize

`func (o *ModelLogFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ModelLogFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ModelLogFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ModelLogFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetName

`func (o *ModelLogFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelLogFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelLogFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelLogFile) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


