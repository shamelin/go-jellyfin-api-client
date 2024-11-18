# ModelFileSystemEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets the name. | [optional] 
**Path** | Pointer to **string** | Gets the path. | [optional] 
**Type** | Pointer to [**ModelModelFileSystemEntryType**](ModelFileSystemEntryType.md) | Gets the type. | [optional] 

## Methods

### NewModelFileSystemEntryInfo

`func NewModelFileSystemEntryInfo() *ModelFileSystemEntryInfo`

NewModelFileSystemEntryInfo instantiates a new ModelFileSystemEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFileSystemEntryInfoWithDefaults

`func NewModelFileSystemEntryInfoWithDefaults() *ModelFileSystemEntryInfo`

NewModelFileSystemEntryInfoWithDefaults instantiates a new ModelFileSystemEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelFileSystemEntryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelFileSystemEntryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelFileSystemEntryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelFileSystemEntryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ModelFileSystemEntryInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelFileSystemEntryInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelFileSystemEntryInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelFileSystemEntryInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *ModelFileSystemEntryInfo) GetType() ModelModelFileSystemEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelFileSystemEntryInfo) GetTypeOk() (*ModelModelFileSystemEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelFileSystemEntryInfo) SetType(v ModelModelFileSystemEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelFileSystemEntryInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


