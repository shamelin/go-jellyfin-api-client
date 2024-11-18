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

// checks if the JellyfinVirtualFolderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinVirtualFolderInfo{}

// JellyfinVirtualFolderInfo Used to hold information about a user's list of configured virtual folders.
type JellyfinVirtualFolderInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the locations.
	Locations []string `json:"Locations,omitempty"`
	// Gets or sets the type of the collection.
	CollectionType NullableJellyfinCollectionTypeOptions `json:"CollectionType,omitempty"`
	LibraryOptions NullableJellyfinLibraryOptions `json:"LibraryOptions,omitempty"`
	// Gets or sets the item identifier.
	ItemId NullableString `json:"ItemId,omitempty"`
	// Gets or sets the primary image item identifier.
	PrimaryImageItemId NullableString `json:"PrimaryImageItemId,omitempty"`
	RefreshProgress NullableFloat64 `json:"RefreshProgress,omitempty"`
	RefreshStatus NullableString `json:"RefreshStatus,omitempty"`
}

// NewJellyfinVirtualFolderInfo instantiates a new JellyfinVirtualFolderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinVirtualFolderInfo() *JellyfinVirtualFolderInfo {
	this := JellyfinVirtualFolderInfo{}
	return &this
}

// NewJellyfinVirtualFolderInfoWithDefaults instantiates a new JellyfinVirtualFolderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinVirtualFolderInfoWithDefaults() *JellyfinVirtualFolderInfo {
	this := JellyfinVirtualFolderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JellyfinVirtualFolderInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetName() {
	o.Name.Unset()
}

// GetLocations returns the Locations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetLocations() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetLocationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *JellyfinVirtualFolderInfo) SetLocations(v []string) {
	o.Locations = v
}

// GetCollectionType returns the CollectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetCollectionType() JellyfinCollectionTypeOptions {
	if o == nil || IsNil(o.CollectionType.Get()) {
		var ret JellyfinCollectionTypeOptions
		return ret
	}
	return *o.CollectionType.Get()
}

// GetCollectionTypeOk returns a tuple with the CollectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetCollectionTypeOk() (*JellyfinCollectionTypeOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionType.Get(), o.CollectionType.IsSet()
}

// HasCollectionType returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasCollectionType() bool {
	if o != nil && o.CollectionType.IsSet() {
		return true
	}

	return false
}

// SetCollectionType gets a reference to the given NullableJellyfinCollectionTypeOptions and assigns it to the CollectionType field.
func (o *JellyfinVirtualFolderInfo) SetCollectionType(v JellyfinCollectionTypeOptions) {
	o.CollectionType.Set(&v)
}
// SetCollectionTypeNil sets the value for CollectionType to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetCollectionTypeNil() {
	o.CollectionType.Set(nil)
}

// UnsetCollectionType ensures that no value is present for CollectionType, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetCollectionType() {
	o.CollectionType.Unset()
}

// GetLibraryOptions returns the LibraryOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetLibraryOptions() JellyfinLibraryOptions {
	if o == nil || IsNil(o.LibraryOptions.Get()) {
		var ret JellyfinLibraryOptions
		return ret
	}
	return *o.LibraryOptions.Get()
}

// GetLibraryOptionsOk returns a tuple with the LibraryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetLibraryOptionsOk() (*JellyfinLibraryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibraryOptions.Get(), o.LibraryOptions.IsSet()
}

// HasLibraryOptions returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasLibraryOptions() bool {
	if o != nil && o.LibraryOptions.IsSet() {
		return true
	}

	return false
}

// SetLibraryOptions gets a reference to the given NullableJellyfinLibraryOptions and assigns it to the LibraryOptions field.
func (o *JellyfinVirtualFolderInfo) SetLibraryOptions(v JellyfinLibraryOptions) {
	o.LibraryOptions.Set(&v)
}
// SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetLibraryOptionsNil() {
	o.LibraryOptions.Set(nil)
}

// UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetLibraryOptions() {
	o.LibraryOptions.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *JellyfinVirtualFolderInfo) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetItemId() {
	o.ItemId.Unset()
}

// GetPrimaryImageItemId returns the PrimaryImageItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetPrimaryImageItemId() string {
	if o == nil || IsNil(o.PrimaryImageItemId.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryImageItemId.Get()
}

// GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetPrimaryImageItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageItemId.Get(), o.PrimaryImageItemId.IsSet()
}

// HasPrimaryImageItemId returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasPrimaryImageItemId() bool {
	if o != nil && o.PrimaryImageItemId.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageItemId gets a reference to the given NullableString and assigns it to the PrimaryImageItemId field.
func (o *JellyfinVirtualFolderInfo) SetPrimaryImageItemId(v string) {
	o.PrimaryImageItemId.Set(&v)
}
// SetPrimaryImageItemIdNil sets the value for PrimaryImageItemId to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetPrimaryImageItemIdNil() {
	o.PrimaryImageItemId.Set(nil)
}

// UnsetPrimaryImageItemId ensures that no value is present for PrimaryImageItemId, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetPrimaryImageItemId() {
	o.PrimaryImageItemId.Unset()
}

// GetRefreshProgress returns the RefreshProgress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetRefreshProgress() float64 {
	if o == nil || IsNil(o.RefreshProgress.Get()) {
		var ret float64
		return ret
	}
	return *o.RefreshProgress.Get()
}

// GetRefreshProgressOk returns a tuple with the RefreshProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetRefreshProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshProgress.Get(), o.RefreshProgress.IsSet()
}

// HasRefreshProgress returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasRefreshProgress() bool {
	if o != nil && o.RefreshProgress.IsSet() {
		return true
	}

	return false
}

// SetRefreshProgress gets a reference to the given NullableFloat64 and assigns it to the RefreshProgress field.
func (o *JellyfinVirtualFolderInfo) SetRefreshProgress(v float64) {
	o.RefreshProgress.Set(&v)
}
// SetRefreshProgressNil sets the value for RefreshProgress to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetRefreshProgressNil() {
	o.RefreshProgress.Set(nil)
}

// UnsetRefreshProgress ensures that no value is present for RefreshProgress, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetRefreshProgress() {
	o.RefreshProgress.Unset()
}

// GetRefreshStatus returns the RefreshStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinVirtualFolderInfo) GetRefreshStatus() string {
	if o == nil || IsNil(o.RefreshStatus.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshStatus.Get()
}

// GetRefreshStatusOk returns a tuple with the RefreshStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinVirtualFolderInfo) GetRefreshStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshStatus.Get(), o.RefreshStatus.IsSet()
}

// HasRefreshStatus returns a boolean if a field has been set.
func (o *JellyfinVirtualFolderInfo) HasRefreshStatus() bool {
	if o != nil && o.RefreshStatus.IsSet() {
		return true
	}

	return false
}

// SetRefreshStatus gets a reference to the given NullableString and assigns it to the RefreshStatus field.
func (o *JellyfinVirtualFolderInfo) SetRefreshStatus(v string) {
	o.RefreshStatus.Set(&v)
}
// SetRefreshStatusNil sets the value for RefreshStatus to be an explicit nil
func (o *JellyfinVirtualFolderInfo) SetRefreshStatusNil() {
	o.RefreshStatus.Set(nil)
}

// UnsetRefreshStatus ensures that no value is present for RefreshStatus, not even an explicit nil
func (o *JellyfinVirtualFolderInfo) UnsetRefreshStatus() {
	o.RefreshStatus.Unset()
}

func (o JellyfinVirtualFolderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinVirtualFolderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Locations != nil {
		toSerialize["Locations"] = o.Locations
	}
	if o.CollectionType.IsSet() {
		toSerialize["CollectionType"] = o.CollectionType.Get()
	}
	if o.LibraryOptions.IsSet() {
		toSerialize["LibraryOptions"] = o.LibraryOptions.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["ItemId"] = o.ItemId.Get()
	}
	if o.PrimaryImageItemId.IsSet() {
		toSerialize["PrimaryImageItemId"] = o.PrimaryImageItemId.Get()
	}
	if o.RefreshProgress.IsSet() {
		toSerialize["RefreshProgress"] = o.RefreshProgress.Get()
	}
	if o.RefreshStatus.IsSet() {
		toSerialize["RefreshStatus"] = o.RefreshStatus.Get()
	}
	return toSerialize, nil
}

type NullableJellyfinVirtualFolderInfo struct {
	value *JellyfinVirtualFolderInfo
	isSet bool
}

func (v NullableJellyfinVirtualFolderInfo) Get() *JellyfinVirtualFolderInfo {
	return v.value
}

func (v *NullableJellyfinVirtualFolderInfo) Set(val *JellyfinVirtualFolderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinVirtualFolderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinVirtualFolderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinVirtualFolderInfo(val *JellyfinVirtualFolderInfo) *NullableJellyfinVirtualFolderInfo {
	return &NullableJellyfinVirtualFolderInfo{value: val, isSet: true}
}

func (v NullableJellyfinVirtualFolderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinVirtualFolderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

