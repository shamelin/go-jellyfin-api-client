/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"encoding/json"
)

// checks if the RecommendationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationDto{}

// RecommendationDto struct for RecommendationDto
type RecommendationDto struct {
	Items []BaseItemDto `json:"Items,omitempty"`
	RecommendationType *RecommendationType `json:"RecommendationType,omitempty"`
	BaselineItemName NullableString `json:"BaselineItemName,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty"`
}

// NewRecommendationDto instantiates a new RecommendationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationDto() *RecommendationDto {
	this := RecommendationDto{}
	return &this
}

// NewRecommendationDtoWithDefaults instantiates a new RecommendationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationDtoWithDefaults() *RecommendationDto {
	this := RecommendationDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationDto) GetItems() []BaseItemDto {
	if o == nil {
		var ret []BaseItemDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationDto) GetItemsOk() ([]BaseItemDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RecommendationDto) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BaseItemDto and assigns it to the Items field.
func (o *RecommendationDto) SetItems(v []BaseItemDto) {
	o.Items = v
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise.
func (o *RecommendationDto) GetRecommendationType() RecommendationType {
	if o == nil || IsNil(o.RecommendationType) {
		var ret RecommendationType
		return ret
	}
	return *o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationDto) GetRecommendationTypeOk() (*RecommendationType, bool) {
	if o == nil || IsNil(o.RecommendationType) {
		return nil, false
	}
	return o.RecommendationType, true
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *RecommendationDto) HasRecommendationType() bool {
	if o != nil && !IsNil(o.RecommendationType) {
		return true
	}

	return false
}

// SetRecommendationType gets a reference to the given RecommendationType and assigns it to the RecommendationType field.
func (o *RecommendationDto) SetRecommendationType(v RecommendationType) {
	o.RecommendationType = &v
}

// GetBaselineItemName returns the BaselineItemName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationDto) GetBaselineItemName() string {
	if o == nil || IsNil(o.BaselineItemName.Get()) {
		var ret string
		return ret
	}
	return *o.BaselineItemName.Get()
}

// GetBaselineItemNameOk returns a tuple with the BaselineItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationDto) GetBaselineItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaselineItemName.Get(), o.BaselineItemName.IsSet()
}

// HasBaselineItemName returns a boolean if a field has been set.
func (o *RecommendationDto) HasBaselineItemName() bool {
	if o != nil && o.BaselineItemName.IsSet() {
		return true
	}

	return false
}

// SetBaselineItemName gets a reference to the given NullableString and assigns it to the BaselineItemName field.
func (o *RecommendationDto) SetBaselineItemName(v string) {
	o.BaselineItemName.Set(&v)
}
// SetBaselineItemNameNil sets the value for BaselineItemName to be an explicit nil
func (o *RecommendationDto) SetBaselineItemNameNil() {
	o.BaselineItemName.Set(nil)
}

// UnsetBaselineItemName ensures that no value is present for BaselineItemName, not even an explicit nil
func (o *RecommendationDto) UnsetBaselineItemName() {
	o.BaselineItemName.Unset()
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *RecommendationDto) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationDto) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *RecommendationDto) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *RecommendationDto) SetCategoryId(v string) {
	o.CategoryId = &v
}

func (o RecommendationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.RecommendationType) {
		toSerialize["RecommendationType"] = o.RecommendationType
	}
	if o.BaselineItemName.IsSet() {
		toSerialize["BaselineItemName"] = o.BaselineItemName.Get()
	}
	if !IsNil(o.CategoryId) {
		toSerialize["CategoryId"] = o.CategoryId
	}
	return toSerialize, nil
}

type NullableRecommendationDto struct {
	value *RecommendationDto
	isSet bool
}

func (v NullableRecommendationDto) Get() *RecommendationDto {
	return v.value
}

func (v *NullableRecommendationDto) Set(val *RecommendationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationDto(val *RecommendationDto) *NullableRecommendationDto {
	return &NullableRecommendationDto{value: val, isSet: true}
}

func (v NullableRecommendationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


