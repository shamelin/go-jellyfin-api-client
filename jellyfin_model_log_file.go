/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyfin

import (
	"encoding/json"
	"time"
)

// checks if the JellyfinLogFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinLogFile{}

// JellyfinLogFile struct for JellyfinLogFile
type JellyfinLogFile struct {
	// Gets or sets the date created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Gets or sets the date modified.
	DateModified *time.Time `json:"DateModified,omitempty"`
	// Gets or sets the size.
	Size *int64 `json:"Size,omitempty"`
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
}

// NewJellyfinLogFile instantiates a new JellyfinLogFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinLogFile() *JellyfinLogFile {
	this := JellyfinLogFile{}
	return &this
}

// NewJellyfinLogFileWithDefaults instantiates a new JellyfinLogFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinLogFileWithDefaults() *JellyfinLogFile {
	this := JellyfinLogFile{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *JellyfinLogFile) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLogFile) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *JellyfinLogFile) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *JellyfinLogFile) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateModified returns the DateModified field value if set, zero value otherwise.
func (o *JellyfinLogFile) GetDateModified() time.Time {
	if o == nil || IsNil(o.DateModified) {
		var ret time.Time
		return ret
	}
	return *o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLogFile) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateModified) {
		return nil, false
	}
	return o.DateModified, true
}

// HasDateModified returns a boolean if a field has been set.
func (o *JellyfinLogFile) HasDateModified() bool {
	if o != nil && !IsNil(o.DateModified) {
		return true
	}

	return false
}

// SetDateModified gets a reference to the given time.Time and assigns it to the DateModified field.
func (o *JellyfinLogFile) SetDateModified(v time.Time) {
	o.DateModified = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *JellyfinLogFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLogFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *JellyfinLogFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *JellyfinLogFile) SetSize(v int64) {
	o.Size = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JellyfinLogFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinLogFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinLogFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JellyfinLogFile) SetName(v string) {
	o.Name = &v
}

func (o JellyfinLogFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinLogFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if !IsNil(o.DateModified) {
		toSerialize["DateModified"] = o.DateModified
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableJellyfinLogFile struct {
	value *JellyfinLogFile
	isSet bool
}

func (v NullableJellyfinLogFile) Get() *JellyfinLogFile {
	return v.value
}

func (v *NullableJellyfinLogFile) Set(val *JellyfinLogFile) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinLogFile) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinLogFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinLogFile(val *JellyfinLogFile) *NullableJellyfinLogFile {
	return &NullableJellyfinLogFile{value: val, isSet: true}
}

func (v NullableJellyfinLogFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinLogFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


