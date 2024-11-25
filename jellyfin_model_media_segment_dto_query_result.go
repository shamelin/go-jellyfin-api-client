/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
)

// checks if the JellyfinMediaSegmentDtoQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinMediaSegmentDtoQueryResult{}

// JellyfinMediaSegmentDtoQueryResult Query result container.
type JellyfinMediaSegmentDtoQueryResult struct {
	// Gets or sets the items.
	Items []JellyfinMediaSegmentDto `json:"Items,omitempty"`
	// Gets or sets the total number of records available.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the index of the first record in Items.
	StartIndex *int32 `json:"StartIndex,omitempty"`
}

// NewJellyfinMediaSegmentDtoQueryResult instantiates a new JellyfinMediaSegmentDtoQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinMediaSegmentDtoQueryResult() *JellyfinMediaSegmentDtoQueryResult {
	this := JellyfinMediaSegmentDtoQueryResult{}
	return &this
}

// NewJellyfinMediaSegmentDtoQueryResultWithDefaults instantiates a new JellyfinMediaSegmentDtoQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinMediaSegmentDtoQueryResultWithDefaults() *JellyfinMediaSegmentDtoQueryResult {
	this := JellyfinMediaSegmentDtoQueryResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *JellyfinMediaSegmentDtoQueryResult) GetItems() []JellyfinMediaSegmentDto {
	if o == nil || IsNil(o.Items) {
		var ret []JellyfinMediaSegmentDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) GetItemsOk() ([]JellyfinMediaSegmentDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []JellyfinMediaSegmentDto and assigns it to the Items field.
func (o *JellyfinMediaSegmentDtoQueryResult) SetItems(v []JellyfinMediaSegmentDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *JellyfinMediaSegmentDtoQueryResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *JellyfinMediaSegmentDtoQueryResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *JellyfinMediaSegmentDtoQueryResult) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *JellyfinMediaSegmentDtoQueryResult) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *JellyfinMediaSegmentDtoQueryResult) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o JellyfinMediaSegmentDtoQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinMediaSegmentDtoQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	if !IsNil(o.StartIndex) {
		toSerialize["StartIndex"] = o.StartIndex
	}
	return toSerialize, nil
}

type NullableJellyfinMediaSegmentDtoQueryResult struct {
	value *JellyfinMediaSegmentDtoQueryResult
	isSet bool
}

func (v NullableJellyfinMediaSegmentDtoQueryResult) Get() *JellyfinMediaSegmentDtoQueryResult {
	return v.value
}

func (v *NullableJellyfinMediaSegmentDtoQueryResult) Set(val *JellyfinMediaSegmentDtoQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinMediaSegmentDtoQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinMediaSegmentDtoQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinMediaSegmentDtoQueryResult(val *JellyfinMediaSegmentDtoQueryResult) *NullableJellyfinMediaSegmentDtoQueryResult {
	return &NullableJellyfinMediaSegmentDtoQueryResult{value: val, isSet: true}
}

func (v NullableJellyfinMediaSegmentDtoQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinMediaSegmentDtoQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


