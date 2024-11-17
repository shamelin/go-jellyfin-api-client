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

// checks if the SongInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SongInfo{}

// SongInfo struct for SongInfo
type SongInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the original title.
	OriginalTitle NullableString `json:"OriginalTitle,omitempty"`
	// Gets or sets the path.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets the metadata language.
	MetadataLanguage NullableString `json:"MetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode NullableString `json:"MetadataCountryCode,omitempty"`
	// Gets or sets the provider ids.
	ProviderIds map[string]string `json:"ProviderIds,omitempty"`
	// Gets or sets the year.
	Year NullableInt32 `json:"Year,omitempty"`
	IndexNumber NullableInt32 `json:"IndexNumber,omitempty"`
	ParentIndexNumber NullableInt32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate NullableTime `json:"PremiereDate,omitempty"`
	IsAutomated *bool `json:"IsAutomated,omitempty"`
	AlbumArtists []string `json:"AlbumArtists,omitempty"`
	Album NullableString `json:"Album,omitempty"`
	Artists []string `json:"Artists,omitempty"`
}

// NewSongInfo instantiates a new SongInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSongInfo() *SongInfo {
	this := SongInfo{}
	return &this
}

// NewSongInfoWithDefaults instantiates a new SongInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSongInfoWithDefaults() *SongInfo {
	this := SongInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SongInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SongInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SongInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SongInfo) UnsetName() {
	o.Name.Unset()
}

// GetOriginalTitle returns the OriginalTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetOriginalTitle() string {
	if o == nil || IsNil(o.OriginalTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalTitle.Get()
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetOriginalTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalTitle.Get(), o.OriginalTitle.IsSet()
}

// HasOriginalTitle returns a boolean if a field has been set.
func (o *SongInfo) HasOriginalTitle() bool {
	if o != nil && o.OriginalTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginalTitle gets a reference to the given NullableString and assigns it to the OriginalTitle field.
func (o *SongInfo) SetOriginalTitle(v string) {
	o.OriginalTitle.Set(&v)
}
// SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil
func (o *SongInfo) SetOriginalTitleNil() {
	o.OriginalTitle.Set(nil)
}

// UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
func (o *SongInfo) UnsetOriginalTitle() {
	o.OriginalTitle.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *SongInfo) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *SongInfo) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *SongInfo) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *SongInfo) UnsetPath() {
	o.Path.Unset()
}

// GetMetadataLanguage returns the MetadataLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetMetadataLanguage() string {
	if o == nil || IsNil(o.MetadataLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataLanguage.Get()
}

// GetMetadataLanguageOk returns a tuple with the MetadataLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetMetadataLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataLanguage.Get(), o.MetadataLanguage.IsSet()
}

// HasMetadataLanguage returns a boolean if a field has been set.
func (o *SongInfo) HasMetadataLanguage() bool {
	if o != nil && o.MetadataLanguage.IsSet() {
		return true
	}

	return false
}

// SetMetadataLanguage gets a reference to the given NullableString and assigns it to the MetadataLanguage field.
func (o *SongInfo) SetMetadataLanguage(v string) {
	o.MetadataLanguage.Set(&v)
}
// SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil
func (o *SongInfo) SetMetadataLanguageNil() {
	o.MetadataLanguage.Set(nil)
}

// UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
func (o *SongInfo) UnsetMetadataLanguage() {
	o.MetadataLanguage.Unset()
}

// GetMetadataCountryCode returns the MetadataCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetMetadataCountryCode() string {
	if o == nil || IsNil(o.MetadataCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataCountryCode.Get()
}

// GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetMetadataCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataCountryCode.Get(), o.MetadataCountryCode.IsSet()
}

// HasMetadataCountryCode returns a boolean if a field has been set.
func (o *SongInfo) HasMetadataCountryCode() bool {
	if o != nil && o.MetadataCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMetadataCountryCode gets a reference to the given NullableString and assigns it to the MetadataCountryCode field.
func (o *SongInfo) SetMetadataCountryCode(v string) {
	o.MetadataCountryCode.Set(&v)
}
// SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil
func (o *SongInfo) SetMetadataCountryCodeNil() {
	o.MetadataCountryCode.Set(nil)
}

// UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
func (o *SongInfo) UnsetMetadataCountryCode() {
	o.MetadataCountryCode.Unset()
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetProviderIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetProviderIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ProviderIds) {
		return nil, false
	}
	return &o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *SongInfo) HasProviderIds() bool {
	if o != nil && !IsNil(o.ProviderIds) {
		return true
	}

	return false
}

// SetProviderIds gets a reference to the given map[string]string and assigns it to the ProviderIds field.
func (o *SongInfo) SetProviderIds(v map[string]string) {
	o.ProviderIds = v
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetYear() int32 {
	if o == nil || IsNil(o.Year.Get()) {
		var ret int32
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *SongInfo) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableInt32 and assigns it to the Year field.
func (o *SongInfo) SetYear(v int32) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *SongInfo) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *SongInfo) UnsetYear() {
	o.Year.Unset()
}

// GetIndexNumber returns the IndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetIndexNumber() int32 {
	if o == nil || IsNil(o.IndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexNumber.Get()
}

// GetIndexNumberOk returns a tuple with the IndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexNumber.Get(), o.IndexNumber.IsSet()
}

// HasIndexNumber returns a boolean if a field has been set.
func (o *SongInfo) HasIndexNumber() bool {
	if o != nil && o.IndexNumber.IsSet() {
		return true
	}

	return false
}

// SetIndexNumber gets a reference to the given NullableInt32 and assigns it to the IndexNumber field.
func (o *SongInfo) SetIndexNumber(v int32) {
	o.IndexNumber.Set(&v)
}
// SetIndexNumberNil sets the value for IndexNumber to be an explicit nil
func (o *SongInfo) SetIndexNumberNil() {
	o.IndexNumber.Set(nil)
}

// UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
func (o *SongInfo) UnsetIndexNumber() {
	o.IndexNumber.Unset()
}

// GetParentIndexNumber returns the ParentIndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetParentIndexNumber() int32 {
	if o == nil || IsNil(o.ParentIndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentIndexNumber.Get()
}

// GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetParentIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentIndexNumber.Get(), o.ParentIndexNumber.IsSet()
}

// HasParentIndexNumber returns a boolean if a field has been set.
func (o *SongInfo) HasParentIndexNumber() bool {
	if o != nil && o.ParentIndexNumber.IsSet() {
		return true
	}

	return false
}

// SetParentIndexNumber gets a reference to the given NullableInt32 and assigns it to the ParentIndexNumber field.
func (o *SongInfo) SetParentIndexNumber(v int32) {
	o.ParentIndexNumber.Set(&v)
}
// SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil
func (o *SongInfo) SetParentIndexNumberNil() {
	o.ParentIndexNumber.Set(nil)
}

// UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
func (o *SongInfo) UnsetParentIndexNumber() {
	o.ParentIndexNumber.Unset()
}

// GetPremiereDate returns the PremiereDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetPremiereDate() time.Time {
	if o == nil || IsNil(o.PremiereDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PremiereDate.Get()
}

// GetPremiereDateOk returns a tuple with the PremiereDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetPremiereDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiereDate.Get(), o.PremiereDate.IsSet()
}

// HasPremiereDate returns a boolean if a field has been set.
func (o *SongInfo) HasPremiereDate() bool {
	if o != nil && o.PremiereDate.IsSet() {
		return true
	}

	return false
}

// SetPremiereDate gets a reference to the given NullableTime and assigns it to the PremiereDate field.
func (o *SongInfo) SetPremiereDate(v time.Time) {
	o.PremiereDate.Set(&v)
}
// SetPremiereDateNil sets the value for PremiereDate to be an explicit nil
func (o *SongInfo) SetPremiereDateNil() {
	o.PremiereDate.Set(nil)
}

// UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
func (o *SongInfo) UnsetPremiereDate() {
	o.PremiereDate.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *SongInfo) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated) {
		var ret bool
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SongInfo) GetIsAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *SongInfo) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given bool and assigns it to the IsAutomated field.
func (o *SongInfo) SetIsAutomated(v bool) {
	o.IsAutomated = &v
}

// GetAlbumArtists returns the AlbumArtists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetAlbumArtists() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AlbumArtists
}

// GetAlbumArtistsOk returns a tuple with the AlbumArtists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetAlbumArtistsOk() ([]string, bool) {
	if o == nil || IsNil(o.AlbumArtists) {
		return nil, false
	}
	return o.AlbumArtists, true
}

// HasAlbumArtists returns a boolean if a field has been set.
func (o *SongInfo) HasAlbumArtists() bool {
	if o != nil && !IsNil(o.AlbumArtists) {
		return true
	}

	return false
}

// SetAlbumArtists gets a reference to the given []string and assigns it to the AlbumArtists field.
func (o *SongInfo) SetAlbumArtists(v []string) {
	o.AlbumArtists = v
}

// GetAlbum returns the Album field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetAlbum() string {
	if o == nil || IsNil(o.Album.Get()) {
		var ret string
		return ret
	}
	return *o.Album.Get()
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetAlbumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Album.Get(), o.Album.IsSet()
}

// HasAlbum returns a boolean if a field has been set.
func (o *SongInfo) HasAlbum() bool {
	if o != nil && o.Album.IsSet() {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given NullableString and assigns it to the Album field.
func (o *SongInfo) SetAlbum(v string) {
	o.Album.Set(&v)
}
// SetAlbumNil sets the value for Album to be an explicit nil
func (o *SongInfo) SetAlbumNil() {
	o.Album.Set(nil)
}

// UnsetAlbum ensures that no value is present for Album, not even an explicit nil
func (o *SongInfo) UnsetAlbum() {
	o.Album.Unset()
}

// GetArtists returns the Artists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SongInfo) GetArtists() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SongInfo) GetArtistsOk() ([]string, bool) {
	if o == nil || IsNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *SongInfo) HasArtists() bool {
	if o != nil && !IsNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []string and assigns it to the Artists field.
func (o *SongInfo) SetArtists(v []string) {
	o.Artists = v
}

func (o SongInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SongInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.OriginalTitle.IsSet() {
		toSerialize["OriginalTitle"] = o.OriginalTitle.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.MetadataLanguage.IsSet() {
		toSerialize["MetadataLanguage"] = o.MetadataLanguage.Get()
	}
	if o.MetadataCountryCode.IsSet() {
		toSerialize["MetadataCountryCode"] = o.MetadataCountryCode.Get()
	}
	if o.ProviderIds != nil {
		toSerialize["ProviderIds"] = o.ProviderIds
	}
	if o.Year.IsSet() {
		toSerialize["Year"] = o.Year.Get()
	}
	if o.IndexNumber.IsSet() {
		toSerialize["IndexNumber"] = o.IndexNumber.Get()
	}
	if o.ParentIndexNumber.IsSet() {
		toSerialize["ParentIndexNumber"] = o.ParentIndexNumber.Get()
	}
	if o.PremiereDate.IsSet() {
		toSerialize["PremiereDate"] = o.PremiereDate.Get()
	}
	if !IsNil(o.IsAutomated) {
		toSerialize["IsAutomated"] = o.IsAutomated
	}
	if o.AlbumArtists != nil {
		toSerialize["AlbumArtists"] = o.AlbumArtists
	}
	if o.Album.IsSet() {
		toSerialize["Album"] = o.Album.Get()
	}
	if o.Artists != nil {
		toSerialize["Artists"] = o.Artists
	}
	return toSerialize, nil
}

type NullableSongInfo struct {
	value *SongInfo
	isSet bool
}

func (v NullableSongInfo) Get() *SongInfo {
	return v.value
}

func (v *NullableSongInfo) Set(val *SongInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSongInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSongInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSongInfo(val *SongInfo) *NullableSongInfo {
	return &NullableSongInfo{value: val, isSet: true}
}

func (v NullableSongInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSongInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


