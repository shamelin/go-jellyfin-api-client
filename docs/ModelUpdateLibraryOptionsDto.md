# ModelUpdateLibraryOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the library item id. | [optional] 
**LibraryOptions** | Pointer to [**NullableModelModelLibraryOptions**](ModelLibraryOptions.md) | Gets or sets library options. | [optional] 

## Methods

### NewModelUpdateLibraryOptionsDto

`func NewModelUpdateLibraryOptionsDto() *ModelUpdateLibraryOptionsDto`

NewModelUpdateLibraryOptionsDto instantiates a new ModelUpdateLibraryOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpdateLibraryOptionsDtoWithDefaults

`func NewModelUpdateLibraryOptionsDtoWithDefaults() *ModelUpdateLibraryOptionsDto`

NewModelUpdateLibraryOptionsDtoWithDefaults instantiates a new ModelUpdateLibraryOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelUpdateLibraryOptionsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelUpdateLibraryOptionsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelUpdateLibraryOptionsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelUpdateLibraryOptionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibraryOptions

`func (o *ModelUpdateLibraryOptionsDto) GetLibraryOptions() ModelModelLibraryOptions`

GetLibraryOptions returns the LibraryOptions field if non-nil, zero value otherwise.

### GetLibraryOptionsOk

`func (o *ModelUpdateLibraryOptionsDto) GetLibraryOptionsOk() (*ModelModelLibraryOptions, bool)`

GetLibraryOptionsOk returns a tuple with the LibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryOptions

`func (o *ModelUpdateLibraryOptionsDto) SetLibraryOptions(v ModelModelLibraryOptions)`

SetLibraryOptions sets LibraryOptions field to given value.

### HasLibraryOptions

`func (o *ModelUpdateLibraryOptionsDto) HasLibraryOptions() bool`

HasLibraryOptions returns a boolean if a field has been set.

### SetLibraryOptionsNil

`func (o *ModelUpdateLibraryOptionsDto) SetLibraryOptionsNil(b bool)`

 SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil

### UnsetLibraryOptions
`func (o *ModelUpdateLibraryOptionsDto) UnsetLibraryOptions()`

UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


