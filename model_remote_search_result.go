/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin-api-client

import (
	"encoding/json"
	"time"
)

// checks if the RemoteSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteSearchResult{}

// RemoteSearchResult struct for RemoteSearchResult
type RemoteSearchResult struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the provider ids.
	ProviderIds map[string]string `json:"ProviderIds,omitempty"`
	// Gets or sets the year.
	ProductionYear NullableInt32 `json:"ProductionYear,omitempty"`
	IndexNumber NullableInt32 `json:"IndexNumber,omitempty"`
	IndexNumberEnd NullableInt32 `json:"IndexNumberEnd,omitempty"`
	ParentIndexNumber NullableInt32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate NullableTime `json:"PremiereDate,omitempty"`
	ImageUrl NullableString `json:"ImageUrl,omitempty"`
	SearchProviderName NullableString `json:"SearchProviderName,omitempty"`
	Overview NullableString `json:"Overview,omitempty"`
	AlbumArtist NullableRemoteSearchResult `json:"AlbumArtist,omitempty"`
	Artists []RemoteSearchResult `json:"Artists,omitempty"`
}

// NewRemoteSearchResult instantiates a new RemoteSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteSearchResult() *RemoteSearchResult {
	this := RemoteSearchResult{}
	return &this
}

// NewRemoteSearchResultWithDefaults instantiates a new RemoteSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteSearchResultWithDefaults() *RemoteSearchResult {
	this := RemoteSearchResult{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RemoteSearchResult) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RemoteSearchResult) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RemoteSearchResult) UnsetName() {
	o.Name.Unset()
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetProviderIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetProviderIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ProviderIds) {
		return nil, false
	}
	return &o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasProviderIds() bool {
	if o != nil && !IsNil(o.ProviderIds) {
		return true
	}

	return false
}

// SetProviderIds gets a reference to the given map[string]string and assigns it to the ProviderIds field.
func (o *RemoteSearchResult) SetProviderIds(v map[string]string) {
	o.ProviderIds = v
}

// GetProductionYear returns the ProductionYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetProductionYear() int32 {
	if o == nil || IsNil(o.ProductionYear.Get()) {
		var ret int32
		return ret
	}
	return *o.ProductionYear.Get()
}

// GetProductionYearOk returns a tuple with the ProductionYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetProductionYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductionYear.Get(), o.ProductionYear.IsSet()
}

// HasProductionYear returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasProductionYear() bool {
	if o != nil && o.ProductionYear.IsSet() {
		return true
	}

	return false
}

// SetProductionYear gets a reference to the given NullableInt32 and assigns it to the ProductionYear field.
func (o *RemoteSearchResult) SetProductionYear(v int32) {
	o.ProductionYear.Set(&v)
}
// SetProductionYearNil sets the value for ProductionYear to be an explicit nil
func (o *RemoteSearchResult) SetProductionYearNil() {
	o.ProductionYear.Set(nil)
}

// UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
func (o *RemoteSearchResult) UnsetProductionYear() {
	o.ProductionYear.Unset()
}

// GetIndexNumber returns the IndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetIndexNumber() int32 {
	if o == nil || IsNil(o.IndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexNumber.Get()
}

// GetIndexNumberOk returns a tuple with the IndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexNumber.Get(), o.IndexNumber.IsSet()
}

// HasIndexNumber returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasIndexNumber() bool {
	if o != nil && o.IndexNumber.IsSet() {
		return true
	}

	return false
}

// SetIndexNumber gets a reference to the given NullableInt32 and assigns it to the IndexNumber field.
func (o *RemoteSearchResult) SetIndexNumber(v int32) {
	o.IndexNumber.Set(&v)
}
// SetIndexNumberNil sets the value for IndexNumber to be an explicit nil
func (o *RemoteSearchResult) SetIndexNumberNil() {
	o.IndexNumber.Set(nil)
}

// UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
func (o *RemoteSearchResult) UnsetIndexNumber() {
	o.IndexNumber.Unset()
}

// GetIndexNumberEnd returns the IndexNumberEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetIndexNumberEnd() int32 {
	if o == nil || IsNil(o.IndexNumberEnd.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexNumberEnd.Get()
}

// GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetIndexNumberEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexNumberEnd.Get(), o.IndexNumberEnd.IsSet()
}

// HasIndexNumberEnd returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasIndexNumberEnd() bool {
	if o != nil && o.IndexNumberEnd.IsSet() {
		return true
	}

	return false
}

// SetIndexNumberEnd gets a reference to the given NullableInt32 and assigns it to the IndexNumberEnd field.
func (o *RemoteSearchResult) SetIndexNumberEnd(v int32) {
	o.IndexNumberEnd.Set(&v)
}
// SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil
func (o *RemoteSearchResult) SetIndexNumberEndNil() {
	o.IndexNumberEnd.Set(nil)
}

// UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
func (o *RemoteSearchResult) UnsetIndexNumberEnd() {
	o.IndexNumberEnd.Unset()
}

// GetParentIndexNumber returns the ParentIndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetParentIndexNumber() int32 {
	if o == nil || IsNil(o.ParentIndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentIndexNumber.Get()
}

// GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetParentIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentIndexNumber.Get(), o.ParentIndexNumber.IsSet()
}

// HasParentIndexNumber returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasParentIndexNumber() bool {
	if o != nil && o.ParentIndexNumber.IsSet() {
		return true
	}

	return false
}

// SetParentIndexNumber gets a reference to the given NullableInt32 and assigns it to the ParentIndexNumber field.
func (o *RemoteSearchResult) SetParentIndexNumber(v int32) {
	o.ParentIndexNumber.Set(&v)
}
// SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil
func (o *RemoteSearchResult) SetParentIndexNumberNil() {
	o.ParentIndexNumber.Set(nil)
}

// UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
func (o *RemoteSearchResult) UnsetParentIndexNumber() {
	o.ParentIndexNumber.Unset()
}

// GetPremiereDate returns the PremiereDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetPremiereDate() time.Time {
	if o == nil || IsNil(o.PremiereDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PremiereDate.Get()
}

// GetPremiereDateOk returns a tuple with the PremiereDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetPremiereDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiereDate.Get(), o.PremiereDate.IsSet()
}

// HasPremiereDate returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasPremiereDate() bool {
	if o != nil && o.PremiereDate.IsSet() {
		return true
	}

	return false
}

// SetPremiereDate gets a reference to the given NullableTime and assigns it to the PremiereDate field.
func (o *RemoteSearchResult) SetPremiereDate(v time.Time) {
	o.PremiereDate.Set(&v)
}
// SetPremiereDateNil sets the value for PremiereDate to be an explicit nil
func (o *RemoteSearchResult) SetPremiereDateNil() {
	o.PremiereDate.Set(nil)
}

// UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
func (o *RemoteSearchResult) UnsetPremiereDate() {
	o.PremiereDate.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *RemoteSearchResult) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *RemoteSearchResult) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *RemoteSearchResult) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.SearchProviderName.Get()
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetSearchProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchProviderName.Get(), o.SearchProviderName.IsSet()
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasSearchProviderName() bool {
	if o != nil && o.SearchProviderName.IsSet() {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given NullableString and assigns it to the SearchProviderName field.
func (o *RemoteSearchResult) SetSearchProviderName(v string) {
	o.SearchProviderName.Set(&v)
}
// SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil
func (o *RemoteSearchResult) SetSearchProviderNameNil() {
	o.SearchProviderName.Set(nil)
}

// UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
func (o *RemoteSearchResult) UnsetSearchProviderName() {
	o.SearchProviderName.Unset()
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetOverview() string {
	if o == nil || IsNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *RemoteSearchResult) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *RemoteSearchResult) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *RemoteSearchResult) UnsetOverview() {
	o.Overview.Unset()
}

// GetAlbumArtist returns the AlbumArtist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetAlbumArtist() RemoteSearchResult {
	if o == nil || IsNil(o.AlbumArtist.Get()) {
		var ret RemoteSearchResult
		return ret
	}
	return *o.AlbumArtist.Get()
}

// GetAlbumArtistOk returns a tuple with the AlbumArtist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetAlbumArtistOk() (*RemoteSearchResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlbumArtist.Get(), o.AlbumArtist.IsSet()
}

// HasAlbumArtist returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasAlbumArtist() bool {
	if o != nil && o.AlbumArtist.IsSet() {
		return true
	}

	return false
}

// SetAlbumArtist gets a reference to the given NullableRemoteSearchResult and assigns it to the AlbumArtist field.
func (o *RemoteSearchResult) SetAlbumArtist(v RemoteSearchResult) {
	o.AlbumArtist.Set(&v)
}
// SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil
func (o *RemoteSearchResult) SetAlbumArtistNil() {
	o.AlbumArtist.Set(nil)
}

// UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
func (o *RemoteSearchResult) UnsetAlbumArtist() {
	o.AlbumArtist.Unset()
}

// GetArtists returns the Artists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSearchResult) GetArtists() []RemoteSearchResult {
	if o == nil {
		var ret []RemoteSearchResult
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSearchResult) GetArtistsOk() ([]RemoteSearchResult, bool) {
	if o == nil || IsNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *RemoteSearchResult) HasArtists() bool {
	if o != nil && !IsNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []RemoteSearchResult and assigns it to the Artists field.
func (o *RemoteSearchResult) SetArtists(v []RemoteSearchResult) {
	o.Artists = v
}

func (o RemoteSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.ProviderIds != nil {
		toSerialize["ProviderIds"] = o.ProviderIds
	}
	if o.ProductionYear.IsSet() {
		toSerialize["ProductionYear"] = o.ProductionYear.Get()
	}
	if o.IndexNumber.IsSet() {
		toSerialize["IndexNumber"] = o.IndexNumber.Get()
	}
	if o.IndexNumberEnd.IsSet() {
		toSerialize["IndexNumberEnd"] = o.IndexNumberEnd.Get()
	}
	if o.ParentIndexNumber.IsSet() {
		toSerialize["ParentIndexNumber"] = o.ParentIndexNumber.Get()
	}
	if o.PremiereDate.IsSet() {
		toSerialize["PremiereDate"] = o.PremiereDate.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["ImageUrl"] = o.ImageUrl.Get()
	}
	if o.SearchProviderName.IsSet() {
		toSerialize["SearchProviderName"] = o.SearchProviderName.Get()
	}
	if o.Overview.IsSet() {
		toSerialize["Overview"] = o.Overview.Get()
	}
	if o.AlbumArtist.IsSet() {
		toSerialize["AlbumArtist"] = o.AlbumArtist.Get()
	}
	if o.Artists != nil {
		toSerialize["Artists"] = o.Artists
	}
	return toSerialize, nil
}

type NullableRemoteSearchResult struct {
	value *RemoteSearchResult
	isSet bool
}

func (v NullableRemoteSearchResult) Get() *RemoteSearchResult {
	return v.value
}

func (v *NullableRemoteSearchResult) Set(val *RemoteSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSearchResult(val *RemoteSearchResult) *NullableRemoteSearchResult {
	return &NullableRemoteSearchResult{value: val, isSet: true}
}

func (v NullableRemoteSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


