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

// checks if the JellyfinCreatePlaylistDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinCreatePlaylistDto{}

// JellyfinCreatePlaylistDto Create new playlist dto.
type JellyfinCreatePlaylistDto struct {
	// Gets or sets the name of the new playlist.
	Name *string `json:"Name,omitempty"`
	// Gets or sets item ids to add to the playlist.
	Ids []string `json:"Ids,omitempty"`
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the media type.
	MediaType NullableJellyfinMediaType `json:"MediaType,omitempty"`
	// Gets or sets the playlist users.
	Users []JellyfinPlaylistUserPermissions `json:"Users,omitempty"`
	// Gets or sets a value indicating whether the playlist is public.
	IsPublic *bool `json:"IsPublic,omitempty"`
}

// NewJellyfinCreatePlaylistDto instantiates a new JellyfinCreatePlaylistDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinCreatePlaylistDto() *JellyfinCreatePlaylistDto {
	this := JellyfinCreatePlaylistDto{}
	return &this
}

// NewJellyfinCreatePlaylistDtoWithDefaults instantiates a new JellyfinCreatePlaylistDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinCreatePlaylistDtoWithDefaults() *JellyfinCreatePlaylistDto {
	this := JellyfinCreatePlaylistDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JellyfinCreatePlaylistDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCreatePlaylistDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JellyfinCreatePlaylistDto) SetName(v string) {
	o.Name = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *JellyfinCreatePlaylistDto) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCreatePlaylistDto) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *JellyfinCreatePlaylistDto) SetIds(v []string) {
	o.Ids = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinCreatePlaylistDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinCreatePlaylistDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *JellyfinCreatePlaylistDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *JellyfinCreatePlaylistDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *JellyfinCreatePlaylistDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JellyfinCreatePlaylistDto) GetMediaType() JellyfinMediaType {
	if o == nil || IsNil(o.MediaType.Get()) {
		var ret JellyfinMediaType
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JellyfinCreatePlaylistDto) GetMediaTypeOk() (*JellyfinMediaType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasMediaType() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableJellyfinMediaType and assigns it to the MediaType field.
func (o *JellyfinCreatePlaylistDto) SetMediaType(v JellyfinMediaType) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *JellyfinCreatePlaylistDto) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *JellyfinCreatePlaylistDto) UnsetMediaType() {
	o.MediaType.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *JellyfinCreatePlaylistDto) GetUsers() []JellyfinPlaylistUserPermissions {
	if o == nil || IsNil(o.Users) {
		var ret []JellyfinPlaylistUserPermissions
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCreatePlaylistDto) GetUsersOk() ([]JellyfinPlaylistUserPermissions, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []JellyfinPlaylistUserPermissions and assigns it to the Users field.
func (o *JellyfinCreatePlaylistDto) SetUsers(v []JellyfinPlaylistUserPermissions) {
	o.Users = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *JellyfinCreatePlaylistDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinCreatePlaylistDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *JellyfinCreatePlaylistDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *JellyfinCreatePlaylistDto) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o JellyfinCreatePlaylistDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinCreatePlaylistDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Ids) {
		toSerialize["Ids"] = o.Ids
	}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["MediaType"] = o.MediaType.Get()
	}
	if !IsNil(o.Users) {
		toSerialize["Users"] = o.Users
	}
	if !IsNil(o.IsPublic) {
		toSerialize["IsPublic"] = o.IsPublic
	}
	return toSerialize, nil
}

type NullableJellyfinCreatePlaylistDto struct {
	value *JellyfinCreatePlaylistDto
	isSet bool
}

func (v NullableJellyfinCreatePlaylistDto) Get() *JellyfinCreatePlaylistDto {
	return v.value
}

func (v *NullableJellyfinCreatePlaylistDto) Set(val *JellyfinCreatePlaylistDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinCreatePlaylistDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinCreatePlaylistDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinCreatePlaylistDto(val *JellyfinCreatePlaylistDto) *NullableJellyfinCreatePlaylistDto {
	return &NullableJellyfinCreatePlaylistDto{value: val, isSet: true}
}

func (v NullableJellyfinCreatePlaylistDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinCreatePlaylistDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

