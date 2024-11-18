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

// checks if the JellyfinGroupInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JellyfinGroupInfoDto{}

// JellyfinGroupInfoDto Class GroupInfoDto.
type JellyfinGroupInfoDto struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the group name.
	GroupName *string `json:"GroupName,omitempty"`
	// Gets the group state.
	State *JellyfinGroupStateType `json:"State,omitempty"`
	// Gets the participants.
	Participants []string `json:"Participants,omitempty"`
	// Gets the date when this DTO has been created.
	LastUpdatedAt *time.Time `json:"LastUpdatedAt,omitempty"`
}

// NewJellyfinGroupInfoDto instantiates a new JellyfinGroupInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJellyfinGroupInfoDto() *JellyfinGroupInfoDto {
	this := JellyfinGroupInfoDto{}
	return &this
}

// NewJellyfinGroupInfoDtoWithDefaults instantiates a new JellyfinGroupInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJellyfinGroupInfoDtoWithDefaults() *JellyfinGroupInfoDto {
	this := JellyfinGroupInfoDto{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *JellyfinGroupInfoDto) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinGroupInfoDto) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *JellyfinGroupInfoDto) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *JellyfinGroupInfoDto) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *JellyfinGroupInfoDto) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinGroupInfoDto) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *JellyfinGroupInfoDto) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *JellyfinGroupInfoDto) SetGroupName(v string) {
	o.GroupName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *JellyfinGroupInfoDto) GetState() JellyfinGroupStateType {
	if o == nil || IsNil(o.State) {
		var ret JellyfinGroupStateType
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinGroupInfoDto) GetStateOk() (*JellyfinGroupStateType, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *JellyfinGroupInfoDto) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given JellyfinGroupStateType and assigns it to the State field.
func (o *JellyfinGroupInfoDto) SetState(v JellyfinGroupStateType) {
	o.State = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *JellyfinGroupInfoDto) GetParticipants() []string {
	if o == nil || IsNil(o.Participants) {
		var ret []string
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinGroupInfoDto) GetParticipantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *JellyfinGroupInfoDto) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []string and assigns it to the Participants field.
func (o *JellyfinGroupInfoDto) SetParticipants(v []string) {
	o.Participants = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *JellyfinGroupInfoDto) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JellyfinGroupInfoDto) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *JellyfinGroupInfoDto) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *JellyfinGroupInfoDto) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

func (o JellyfinGroupInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JellyfinGroupInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["GroupName"] = o.GroupName
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Participants) {
		toSerialize["Participants"] = o.Participants
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["LastUpdatedAt"] = o.LastUpdatedAt
	}
	return toSerialize, nil
}

type NullableJellyfinGroupInfoDto struct {
	value *JellyfinGroupInfoDto
	isSet bool
}

func (v NullableJellyfinGroupInfoDto) Get() *JellyfinGroupInfoDto {
	return v.value
}

func (v *NullableJellyfinGroupInfoDto) Set(val *JellyfinGroupInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableJellyfinGroupInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableJellyfinGroupInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJellyfinGroupInfoDto(val *JellyfinGroupInfoDto) *NullableJellyfinGroupInfoDto {
	return &NullableJellyfinGroupInfoDto{value: val, isSet: true}
}

func (v NullableJellyfinGroupInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJellyfinGroupInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

