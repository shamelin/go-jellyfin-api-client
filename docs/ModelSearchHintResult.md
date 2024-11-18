# ModelSearchHintResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchHints** | Pointer to [**[]ModelModelSearchHint**](ModelModelSearchHint.md) | Gets the search hints. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets the total record count. | [optional] 

## Methods

### NewModelSearchHintResult

`func NewModelSearchHintResult() *ModelSearchHintResult`

NewModelSearchHintResult instantiates a new ModelSearchHintResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSearchHintResultWithDefaults

`func NewModelSearchHintResultWithDefaults() *ModelSearchHintResult`

NewModelSearchHintResultWithDefaults instantiates a new ModelSearchHintResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchHints

`func (o *ModelSearchHintResult) GetSearchHints() []ModelModelSearchHint`

GetSearchHints returns the SearchHints field if non-nil, zero value otherwise.

### GetSearchHintsOk

`func (o *ModelSearchHintResult) GetSearchHintsOk() (*[]ModelModelSearchHint, bool)`

GetSearchHintsOk returns a tuple with the SearchHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchHints

`func (o *ModelSearchHintResult) SetSearchHints(v []ModelModelSearchHint)`

SetSearchHints sets SearchHints field to given value.

### HasSearchHints

`func (o *ModelSearchHintResult) HasSearchHints() bool`

HasSearchHints returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *ModelSearchHintResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *ModelSearchHintResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *ModelSearchHintResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *ModelSearchHintResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


