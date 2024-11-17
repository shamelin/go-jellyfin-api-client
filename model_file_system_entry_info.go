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

// checks if the FileSystemEntryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSystemEntryInfo{}

// FileSystemEntryInfo Class FileSystemEntryInfo.
type FileSystemEntryInfo struct {
	// Gets the name.
	Name *string `json:"Name,omitempty"`
	// Gets the path.
	Path *string `json:"Path,omitempty"`
	// Gets the type.
	Type *FileSystemEntryType `json:"Type,omitempty"`
}

// NewFileSystemEntryInfo instantiates a new FileSystemEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSystemEntryInfo() *FileSystemEntryInfo {
	this := FileSystemEntryInfo{}
	return &this
}

// NewFileSystemEntryInfoWithDefaults instantiates a new FileSystemEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSystemEntryInfoWithDefaults() *FileSystemEntryInfo {
	this := FileSystemEntryInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileSystemEntryInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemEntryInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileSystemEntryInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileSystemEntryInfo) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FileSystemEntryInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemEntryInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FileSystemEntryInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FileSystemEntryInfo) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileSystemEntryInfo) GetType() FileSystemEntryType {
	if o == nil || IsNil(o.Type) {
		var ret FileSystemEntryType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemEntryInfo) GetTypeOk() (*FileSystemEntryType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileSystemEntryInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FileSystemEntryType and assigns it to the Type field.
func (o *FileSystemEntryInfo) SetType(v FileSystemEntryType) {
	o.Type = &v
}

func (o FileSystemEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSystemEntryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFileSystemEntryInfo struct {
	value *FileSystemEntryInfo
	isSet bool
}

func (v NullableFileSystemEntryInfo) Get() *FileSystemEntryInfo {
	return v.value
}

func (v *NullableFileSystemEntryInfo) Set(val *FileSystemEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSystemEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSystemEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSystemEntryInfo(val *FileSystemEntryInfo) *NullableFileSystemEntryInfo {
	return &NullableFileSystemEntryInfo{value: val, isSet: true}
}

func (v NullableFileSystemEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSystemEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


